# Homelab dashboard stack

Implementation of the recommendation in [`dashboard_options.md`](./dashboard_options.md):
[Homepage](https://github.com/gethomepage/homepage) for the dashboard, fronted
by [Authelia](https://www.authelia.com/) forward-auth, with Redis for
Authelia session storage.

This is a **standalone** `docker-compose.yml` — merge the `services`,
`secrets`, and `volumes` blocks into your existing *arr-stack compose file
when you're ready, or run it side-by-side on the same `homelab` Docker
network so Homepage can reach the other containers by name.

## Files

```
homelab/
├── docker-compose.yml
├── .env.example
├── secrets/                       # gitignored — generate locally
└── config/
    ├── homepage/
    │   ├── settings.yaml
    │   ├── services.yaml          # edit: your actual app hosts/API keys
    │   ├── docker.yaml            # docker.sock auto-discovery
    │   └── widgets.yaml
    └── authelia/
        ├── configuration.yml
        └── users_database.yml     # edit: your account + password hash
```

## Setup

1. **Env file**
   ```
   cp .env.example .env
   ```
   Fill in `DASHBOARD_DOMAIN` / `ROOT_DOMAIN`.

2. **Generate Authelia secrets** (three random 64-char strings):
   ```
   mkdir -p secrets
   authelia crypto rand --length 64 --charset alphanumeric > secrets/jwt_secret.txt
   authelia crypto rand --length 64 --charset alphanumeric > secrets/session_secret.txt
   authelia crypto rand --length 64 --charset alphanumeric > secrets/storage_encryption_key.txt
   ```
   (No local Authelia binary? Run it via the container:
   `docker run --rm authelia/authelia:latest authelia crypto rand --length 64 --charset alphanumeric`.)

3. **Set your login**: generate a password hash and drop it into
   `config/authelia/users_database.yml`:
   ```
   docker run --rm authelia/authelia:latest authelia crypto hash generate argon2 --password 'yourpassword'
   ```

4. **Edit `config/homepage/services.yaml`** with your real Sonarr/Radarr/
   Prowlarr/qBittorrent hosts and API keys (Settings → General in each app).

5. **Bring it up**:
   ```
   docker compose up -d
   ```
   Homepage listens on `:3000`, Authelia on `:9091`.

## Nginx Proxy Manager wiring

Create two proxy hosts on the VPS, both pointing at the home server's
WireGuard IP:

- `dash.snvgglebear.com` → `http://<wg-ip>:3000` (Homepage)
- `auth.snvgglebear.com` → `http://<wg-ip>:9091` (Authelia portal)

Then, in NPM's **Advanced** tab for `dash.snvgglebear.com` (and any other
subdomain you want gated), add:

```nginx
location /authelia {
    internal;
    set $upstream_authelia http://<wg-ip>:9091/api/verify;
    proxy_pass_request_body off;
    proxy_pass $upstream_authelia;
    proxy_set_header Content-Length "";
    proxy_set_header X-Original-URL $scheme://$http_host$request_uri;
    proxy_set_header X-Original-Method $request_method;
    proxy_set_header X-Forwarded-Method $request_method;
    proxy_set_header X-Forwarded-Proto $scheme;
    proxy_set_header X-Forwarded-Host $http_host;
    proxy_set_header X-Forwarded-Uri $request_uri;
    proxy_set_header X-Forwarded-For $remote_addr;
    proxy_set_header Content-Type "";
    proxy_buffers 4 32k;
    proxy_buffer_size 32k;
}

location / {
    auth_request /authelia;
    auth_request_set $target_url $scheme://$http_host$request_uri;
    auth_request_set $user $upstream_http_remote_user;
    auth_request_set $groups $upstream_http_remote_groups;
    auth_request_set $name $upstream_http_remote_name;
    auth_request_set $email $upstream_http_remote_email;
    proxy_set_header Remote-User $user;
    proxy_set_header Remote-Groups $groups;
    proxy_set_header Remote-Name $name;
    proxy_set_header Remote-Email $email;

    error_page 401 =302 https://auth.snvgglebear.com/?rd=$target_url;

    # existing proxy_pass to Homepage / the app stays below this
}
```

Repeat the same two `location` blocks on any other app's proxy host you want
gated by the same login (e.g. `radarr.snvgglebear.com`), pointing
`$upstream_authelia` at the same Authelia instance each time.

## Notes

- `secrets/*.txt`, `.env`, and Authelia's runtime `db.sqlite3` /
  `notification.txt` are gitignored — never commit real secrets or the
  password database.
- `access_control` in `configuration.yml` defaults every `*.snvgglebear.com`
  host to `one_factor`. Register TOTP for your account and bump the policy
  to `two_factor` once you've confirmed login works.
- The `notifier.filesystem` block writes password-reset/2FA links to a local
  file — fine to start, but swap in `smtp:` once you have a mail relay so
  you can actually receive those links remotely.
