# Homelab Dashboard — Remaining Setup

Decision made: **Homepage** (dashboard) fronted by **Authelia** forward-auth,
with **Redis** for Authelia session storage. Those services are now merged
into [`arrstack-compose.yml`](./arrstack-compose.yml) — this is the
checklist of what's left before bringing the stack up.

## Steps

1. **Env file**
   ```
   cp .env.example .env
   ```
   Fill in `DASHBOARD_DOMAIN`, `ROOT_DOMAIN`, `TZ`.

2. **Generate Authelia secrets** — `secrets/` currently only has a
   `.gitkeep`; the three files below still need to be created:
   ```
   authelia crypto rand --length 64 --charset alphanumeric > secrets/jwt_secret.txt
   authelia crypto rand --length 64 --charset alphanumeric > secrets/session_secret.txt
   authelia crypto rand --length 64 --charset alphanumeric > secrets/storage_encryption_key.txt
   ```
   (No local Authelia binary? `docker run --rm authelia/authelia:latest authelia crypto rand --length 64 --charset alphanumeric`.)

3. **Set your login** — `config/authelia/users_database.yml` still has the
   placeholder `$argon2id$REPLACE_WITH_GENERATED_HASH`. Generate a real hash
   and drop it in:
   ```
   docker run --rm authelia/authelia:latest authelia crypto hash generate argon2 --password 'yourpassword'
   ```

4. **Fill in `config/homepage/services.yaml`** — still has
   `REPLACE_ME_*_API_KEY` placeholders for Sonarr/Radarr/Prowlarr/qBittorrent
   (Settings → General in each app).

   Note: those *arr containers run with `network_mode: service:gluetun` in
   `arrstack-compose.yml`, so they aren't attached to the `homelab` network
   and won't resolve by container name (`http://sonarr:8989`) from Homepage.
   Either point the widget `url:` fields at gluetun's published host ports
   (e.g. `http://<home-server-ip>:8989`) or attach `gluetun` to the
   `homelab` network as well.

5. **Bring the stack up**:
   ```
   docker compose -f arrstack-compose.yml up -d
   ```
   Homepage listens on host port `3001` (`3000` was taken locally), Authelia
   on host port `9092` (`9091` was already claimed by gluetun/transmission).

6. **Wire up NPM** — two proxy hosts on the VPS pointing at the home
   server's WireGuard IP:
   - `dash.snvgglebear.com` → `http://<wg-ip>:3001` (Homepage)
   - `auth.snvgglebear.com` → `http://<wg-ip>:9092` (Authelia portal)

   Then add the forward-auth `location` blocks (see
   [`README.md`](./README.md#nginx-proxy-manager-wiring)) to each gated
   proxy host's **Advanced** tab, updating `$upstream_authelia` to
   `http://<wg-ip>:9092/api/verify`.

7. **Verify one-factor login works**, then register TOTP for your account
   and bump `access_control` policy in `config/authelia/configuration.yml`
   from `one_factor` to `two_factor`.
