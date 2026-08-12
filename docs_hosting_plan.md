# Self-host the MkDocs site on bear-server, replacing GitHub Pages

## Context

`DIY-projects` docs currently build via `.github/workflows/docs.yml`
(`mkdocs build --strict` → GitHub Pages) at `snvgglebear.github.io/DIY-projects/`.
Carl wants to move hosting to his home server (`bear-server` — this session
is running directly on it) and reach it at `projects.snvgglebear.com`,
proxied through the same VPS/Nginx Proxy Manager + WireGuard tunnel pattern
already used for `dash.snvgglebear.com` / `auth.snvgglebear.com`.

Per Carl's answers: GitHub Pages is replaced (not mirrored), the site
rebuilds via a systemd timer polling `git`, and the docs service lives in
its own `docs-compose.yml` rather than being folded into `arrstack-compose.yml`.

**This document is a proposal only — not yet implemented.**

## Facts gathered from bear-server

- WireGuard tunnel to the VPS is `wg1`, IP `10.1.0.2` — this is what NPM's
  new proxy host will point at.
- Host ports currently in use (checked via `ss -tlnp` + `docker ps -a`):
  `22, 53, 631, 1716, 3000, 3389, 3478, 5055, 5355, 5454, 6969, 7878, 8000,
  8081, 8096, 8191, 8686, 8920, 8989, 9091, 9696, 11000, 51143, 51413`.
  **8081 is already taken by `nextcloud-aio-mastercontainer`** — the docs
  container will use **8082** instead.
- `requirements-docs.txt` already pins `mkdocs-material==9.5.44` and
  `mkdocs-same-dir==0.1.3` — reuse these exact versions in the Docker build
  so the home-server build matches what CI validates.
- The live `arrstack-compose.yml` deployment is a separately-maintained copy
  at `/var/home/carl/docker_apps/arrstack/` (not deployed from the git
  checkout) — confirms the convention here: **deployed stacks live under
  `/var/home/carl/docker_apps/<name>/`**, decoupled from Carl's dev checkout
  at `/var/home/carl/GitHub/DIY-projects`. The docs deploy will follow this
  same convention with its own dedicated clone, so the systemd timer never
  touches Carl's working checkout/branch.
- `mkdocs.yml` has a stale reference (`not_in_nav: homelab/docker-compose.yml`)
  left over from the docker-compose.yml → arrstack-compose.yml merge earlier
  this session, plus a comment mentioning the old filename — both need
  fixing regardless of this task.

## Implementation

### 1. `Dockerfile.docs` (repo root)

Multi-stage build — build the static site with the pinned toolchain, serve
it with plain nginx (no Python/mkdocs runtime in the final image):

```dockerfile
FROM python:3.12-slim AS build
WORKDIR /docs
COPY requirements-docs.txt .
RUN pip install --no-cache-dir -r requirements-docs.txt
COPY . .
RUN mkdocs build --strict

FROM nginx:alpine
COPY --from=build /docs/site /usr/share/nginx/html
```

### 2. `docs-compose.yml` (repo root)

```yaml
services:
  docs:
    build:
      context: .
      dockerfile: Dockerfile.docs
    container_name: diy-projects-docs
    restart: unless-stopped
    ports:
      - "8082:80"
```

No env vars, secrets, or volumes — the image is fully static, rebuilt from
source each time.

### 3. `mkdocs.yml` fixes

- `site_url`: `https://snvgglebear.github.io/DIY-projects/` →
  `https://projects.snvgglebear.com/`
- `not_in_nav`: `homelab/docker-compose.yml` → `homelab/arrstack-compose.yml`
- Update the comment above `exclude_docs` (currently says "docker-compose.yml
  stays in...") to reference `arrstack-compose.yml`

### 4. `.github/workflows/docs.yml`

Drop the `deploy` job (and the now-unused `pages: write` / `id-token: write`
permissions, `environment:` block, and `concurrency` group tied to Pages).
Keep the `build` job's `mkdocs build --strict` step — it stays as a
push/PR gate so broken docs still fail CI, they just no longer publish
anywhere from GitHub.

### 5. Disable GitHub Pages in repo settings

`gh api -X DELETE repos/snvgglebear/DIY-projects/pages` (or Settings → Pages
→ set source to "None") once the home-server site is confirmed working —
do this last, after step 8 verification, so there's no gap with nothing
live.

### 6. Dedicated deploy clone + systemd timer on bear-server

Separate from Carl's dev checkout, matching the `docker_apps/` convention:

```
git clone https://github.com/snvgglebear/DIY-projects.git \
  /var/home/carl/docker_apps/docs-site/DIY-projects
```

`/var/home/carl/docker_apps/docs-site/update.sh`:
```bash
#!/bin/bash
set -euo pipefail
cd /var/home/carl/docker_apps/docs-site/DIY-projects
git fetch origin main
before=$(git rev-parse HEAD)
git reset --hard origin/main
after=$(git rev-parse HEAD)
if [ "$before" != "$after" ]; then
  docker compose -f docs-compose.yml build docs
  docker compose -f docs-compose.yml up -d docs
fi
```

`/etc/systemd/system/docs-update.service`:
```ini
[Unit]
Description=Pull DIY-projects docs and rebuild the docs container if changed

[Service]
Type=oneshot
User=carl
ExecStart=/usr/bin/flock -n /tmp/docs-update.lock /var/home/carl/docker_apps/docs-site/update.sh
```

`/etc/systemd/system/docs-update.timer`:
```ini
[Unit]
Description=Periodically check DIY-projects main for docs updates

[Timer]
OnBootSec=5min
OnUnitActiveSec=15min
Persistent=true

[Install]
WantedBy=timers.target
```

Enable: `chmod +x update.sh && sudo systemctl daemon-reload && sudo systemctl enable --now docs-update.timer`

### 7. NPM proxy host (manual, on the VPS)

New proxy host: `projects.snvgglebear.com` → `http://10.1.0.2:8082`, with
Let's Encrypt TLS via NPM (same as the existing proxy hosts). No forward-auth
needed — this is the same public content GitHub Pages was already serving.

### 8. Verification

- `docker compose -f docs-compose.yml build docs && docker compose -f docs-compose.yml up -d docs`, then `curl -I http://localhost:8082` — confirm 200 and that nav/search/theme render correctly in a browser.
- Confirm `mkdocs build --strict` still exits 0 after the `mkdocs.yml` edits (catches the `not_in_nav` path fix being wrong).
- Manually run `systemctl start docs-update.service`, check `journalctl -u docs-update.service` — confirm it no-ops when `main` hasn't moved, and pulls+rebuilds when it has (push a trivial doc change to `main` and wait for the timer, or trigger manually).
- Once `https://projects.snvgglebear.com` serves correctly through NPM from outside the LAN, delete the GitHub Pages site (step 5) and confirm the old `github.io` URL stops resolving/serving.
