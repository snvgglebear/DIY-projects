# Homelab Dashboard Options

Context: home server running the *arr stack (Sonarr/Radarr/etc.), reachable
remotely through a VPS (`snvgglebear.com`) over a WireGuard tunnel, with
Nginx Proxy Manager (NPM) on the VPS handling reverse proxy + TLS. Goal is a
single dashboard to log into and jump off to any app on the home server.

## Requirements to design around

- **Auth**: the dashboard (or a layer in front of it) needs to gate access
  since it'll be reachable from the public domain. Most dashboard apps below
  have weak or no built-in auth — plan for a proper auth layer regardless of
  which dashboard you pick (see [Authentication](#authentication) below).
- **Network path**: `client -> NPM (VPS) -> WireGuard tunnel -> home server`.
  The dashboard just needs one internal hostname/port on the home server for
  NPM to proxy to; the tunnel makes the home server LAN reachable from the
  VPS as if it were local.
- **Service discovery**: with Docker-based *arr stack, auto-discovery via the
  Docker socket is a big win (adds/removes app tiles without manual config).

## Options

### 1. Homepage (gethomepage/homepage) — recommended

- **What it is**: YAML-config dashboard, very popular, actively developed,
  purpose-built for homelabs.
- **Pros**:
  - First-class widgets for Sonarr, Radarr, Prowlarr, qBittorrent, SABnzbd,
    etc. — shows queue/download stats directly on the tile, not just a link.
  - Docker label-based auto-discovery (`homepage.group`, `homepage.name`,
    etc.) so new containers can appear automatically.
  - Config as YAML/code (`services.yaml`, `settings.yaml`, `docker.yaml`),
    easy to keep in git.
  - Supports bookmarks, weather, search bar, resource monitors (CPU/RAM/disk
    via `glances` or built-in).
- **Cons**: no built-in user auth (must front with Authelia/Authentik or NPM
  access lists).
- **Deploy**: single Docker container, mount config dir + optionally the
  Docker socket (read-only) for auto-discovery.
- **Links**: [Site / docs](https://gethomepage.dev/) · [GitHub](https://github.com/gethomepage/homepage)

### 2. Homarr

- **What it is**: Next.js dashboard with a drag-and-drop web UI (no YAML
  editing required) and a growing set of integrations.
- **Pros**:
  - Built-in **user accounts and login** (unlike Homepage) — closer to
    "log in and see your dashboard" out of the box.
  - Docker integration panel, app widgets (similar set to Homepage:
    Sonarr/Radarr/qBittorrent/etc.), customizable boards per user.
  - Nice UI for less config-file-inclined setups.
- **Cons**: heavier (Postgres/SQLite + separate services in newer versions),
  history of breaking changes between major versions, less "config as code"
  than Homepage.
- **Deploy**: Docker container(s) + a database volume.
- **Links**: [Site / docs](https://homarr.dev/) · [GitHub](https://github.com/homarr-labs/homarr)

### 3. Dashy

- **What it is**: highly themeable, feature-rich dashboard with a visual
  editor and YAML config.
- **Pros**: tons of themes/icons, status-checking (ping/HTTP) per item,
  built-in optional auth (`auth.users` in config with hashed passwords) and
  KeyCloak/basic auth support, widgets for various services.
- **Cons**: fewer *arr-specific "rich" widgets than Homepage (mostly
  status/uptime rather than queue details), UI can feel busier.
- **Deploy**: single Docker container, single `conf.yml`.
- **Links**: [Site / docs](https://dashy.to/) · [GitHub](https://github.com/Lissy93/dashy)

### 4. Heimdall

- **What it is**: one of the older, simpler "application dashboard" projects
  (LinuxServer.io image), just tiles + links, some app-specific widgets
  (Sonarr/Radarr/PiHole/etc.) but a smaller, less-maintained set than Homepage.
- **Pros**: dead simple, low resource use, SQLite-backed UI for adding tiles
  by hand, has a "login required" toggle.
- **Cons**: development has slowed noticeably vs. Homepage/Dashy; fewer
  integrations; UI-driven config only (no config-as-code).
- **Deploy**: single Docker container (`linuxserver/heimdall`).
- **Links**: [Site](https://heimdall.site/) · [GitHub](https://github.com/linuxserver/Heimdall)

### 5. Organizr

- **What it is**: dashboard + iframe-based "tabs" app, historically popular
  specifically for *arr-stack homelabs because it can embed each app's full
  UI in a tab plus has **built-in multi-user auth with per-tab permissions**.
- **Pros**: real user management (roles, per-app access), 2FA support,
  embeds full app UIs (not just links) if you don't mind iframes.
- **Cons**: iframes break for apps that set `X-Frame-Options`/CSP headers
  (may need NPM to strip those headers per-service), UI feels dated, project
  momentum has slowed relative to Homepage.
- **Deploy**: single Docker container.
- **Links**: [Site](https://organizr.app/) · [GitHub](https://github.com/causefx/Organizr)

### 6. Flame

- **What it is**: lightweight self-hosted start page with a simple built-in
  app/bookmark editor and basic Docker integration.
- **Pros**: simple, small footprint, has a basic PIN/password auth.
- **Cons**: minimal widget support compared to Homepage/Homarr, less active
  development.
- **Deploy**: single Docker container.
- **Links**: [GitHub](https://github.com/pawelmalak/flame)

## Comparison at a glance

| Dashboard | Config style        | Built-in auth   | *arr widgets | Docker auto-discovery | Activity |
|-----------|---------------------|-----------------|--------------|------------------------|----------|
| Homepage  | YAML                | No              | Rich         | Yes                    | High |
| Homarr    | UI (drag/drop)       | Yes (accounts)  | Rich         | Yes                    | High |
| Dashy     | YAML + UI editor     | Optional (basic)| Basic status | Limited                | Medium |
| Heimdall  | UI only              | Optional        | Basic        | No                     | Low |
| Organizr  | UI only              | Yes (full RBAC) | Iframe embed | No                     | Low-Medium |
| Flame     | UI + YAML            | Basic PIN       | Minimal      | Yes                    | Low |

## Recommendation

**Homepage**, fronted by **Authelia** (or Authentik) at the NPM layer, gets
you the best mix of rich *arr integrations, git-trackable config, and
auto-discovery — the missing piece (real login) is solved once, at the proxy,
rather than per-app.

If you'd rather have per-user accounts baked into the dashboard itself with
less extra infrastructure, **Homarr** is the next best pick.

A ready-to-adapt implementation of this recommendation (Homepage + Authelia +
Redis, as a standalone `docker-compose.yml`) is in this folder — see
[`docker-compose.yml`](./docker-compose.yml) and [`README.md`](./README.md).

## Authentication

None of the lightweight dashboards should be trusted as your only gate on a
box exposed via a public domain. Two common patterns:

1. **Forward-auth via [Authelia](https://www.authelia.com/) ([GitHub](https://github.com/authelia/authelia)) or [Authentik](https://goauthentik.io/) ([GitHub](https://github.com/goauthentik/authentik)) in front of NPM**
   - Run Authelia (or Authentik) as another container on the home server (or
     VPS), configure NPM's "Advanced" tab per proxy host with the
     `auth_request`/forward-auth snippet, and every request to
     `dash.snvgglebear.com` (and optionally each app's subdomain) requires
     login before it ever reaches the container.
   - Gives you 2FA, session timeouts, and one login for every app, not just
     the dashboard tiles.
   - Slightly more setup (Authelia config file, a Redis instance for session
     storage, and the NPM custom-location snippet).

2. **NPM Access Lists**
   - NPM has built-in "Access Lists" (basic auth + IP allow/deny) you can
     attach directly to a proxy host — much simpler, no extra containers.
   - Weaker than Authelia (no 2FA, single shared credential unless you define
     multiple), but fine as a stopgap or paired with a dashboard that has its
     own login (Homarr/Organizr).

Recommended minimum: put an Authelia forward-auth (or at least an NPM access
list) in front of the dashboard's public hostname, even if the dashboard also
has its own login — defense in depth for anything sitting on a public domain.

## Reverse proxy / networking notes

- Give the dashboard its own subdomain, e.g. `dash.snvgglebear.com`, as a
  separate NPM proxy host pointing at `http://<home-server-wg-ip>:<port>`.
- Individual *arr apps can either:
  - stay WireGuard-only / LAN-only and be reached by clicking through the
    dashboard while on the tunnel, or
  - get their own NPM proxy hosts + subdomains (`radarr.snvgglebear.com`,
    etc.) if you want direct external links from the dashboard tiles. If you
    do this, put each one behind the same Authelia forward-auth so a leaked
    dashboard link doesn't bypass login.
- For Organizr specifically, embedding apps in iframes over HTTPS proxy
  hosts may require adding `proxy_hide_header X-Frame-Options;` and a
  relaxed `Content-Security-Policy` in NPM's per-host "Advanced" config,
  since several *arr apps set frame-blocking headers by default.
- Widgets that hit each app's API (Homepage/Homarr/Dashy) need API keys from
  each *arr app; store these in the dashboard's env/config, not hardcoded in
  a file you'd commit to a public repo.
