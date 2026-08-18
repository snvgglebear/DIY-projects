# Link collection bot — plan

A bot that takes a link + category over chat (Discord or SMS) and files it
into a links page on this docs site. If the category doesn't exist yet, the
bot asks before creating it.

## User flow

1. User sends the bot a message with a URL and a category, e.g.
   `!link https://example.com 3d-printing`.
2. Bot parses the URL and category out of the message.
3. Bot checks whether the category already exists:
   - **Exists** — link is filed under it immediately.
   - **New** — bot replies asking for confirmation before creating the
     category. Submission is held in a short-lived pending state, keyed by
     `(channel, submitter)`, until the user confirms, rejects, or the
     confirmation times out (e.g. 10 minutes).
4. On confirmation, the bot writes the change and the docs site picks it up
   on the next deploy.

Keep the command grammar identical across both channels (`<url> <category>`)
so the parsing logic is shared rather than duplicated per adapter.

## Architecture

- One small service with two inbound adapters that normalize incoming
  messages into a shared `LinkSubmission { url, category, submitter, source }`
  and hand off to a common core module:
  - **Discord adapter** — listens for messages/a slash command in a
    designated channel or DMs from allowlisted users.
  - **Twilio adapter** — small HTTP server exposing the inbound-SMS webhook
    endpoint; needs a public HTTPS URL.
- **Core module** — validates the URL, loads the link data store, checks
  category existence, runs the confirmation flow, and on confirmation writes
  the new entry and regenerates the docs page.

## Data storage

Links live in one structured data file, not hand-edited Markdown:

```
link-bot/data/links.yml
```

```yaml
3d-printing:
  - url: https://example.com
    title: Example
    added_by: discord:123456
    added_at: 2026-08-16T00:00:00Z
```

A generator script renders this file into `link-bot/links.md`: one `##`
heading per category, a bullet list of links underneath. Using one page with
category sections (rather than one page per category) means new categories
never require a `mkdocs.yml` nav edit — only `links.md` itself needs a single,
one-time nav entry. That matters because MkDocs here builds with `--strict`
and fails CI on anything missing from `nav`.

## Publishing back to the repo

The generator runs, then the bot commits and pushes the updated `links.yml`
+ `links.md` straight to `main` using a deploy key/PAT. The existing
`docs.yml` GitHub Actions workflow already deploys on push to `main`, so
there's no separate deploy step on the bot's side.

A PR-per-link flow was considered and rejected as too much friction for a
single-user quick-capture tool; direct push is fine at this scale.

## Category confirmation

- **Discord** — reply in-channel/thread: *"Category '`x`' doesn't exist yet.
  React ✅ to create it and file the link, or ❌ to cancel."* Reaction or a
  plain "yes"/"no" reply both work.
- **Twilio SMS** — SMS has no threading, so state must live server-side:
  *"Category '`x`' is new. Reply YES to create it, anything else cancels."*
  Match the pending confirmation to the sender's phone number.
- Both adapters share one pending-confirmation store (in-memory is fine to
  start; TTL per entry) so the confirm/cancel/timeout logic isn't duplicated.

## Auth / abuse prevention

- Discord: restrict to a specific server + role, or DM-only from an
  allowlist of Discord user IDs.
- Twilio: allowlist of phone numbers — anyone who has the Twilio number can
  otherwise text it.

## Deployment

Runs as another container in `homelab/docker-compose.yml`, alongside the
existing dashboard/*arr stack, behind the VPS + WireGuard tunnel already
documented in `homelab/README.md`. The Twilio webhook needs a public HTTPS
endpoint — reuse the VPS reverse proxy that already fronts the homelab
rather than standing up a new one.

Config via environment variables:

- `DISCORD_BOT_TOKEN`, `DISCORD_ALLOWED_USER_IDS`
- `TWILIO_ACCOUNT_SID`, `TWILIO_AUTH_TOKEN`, `TWILIO_ALLOWED_NUMBERS`
- Git deploy key/PAT for pushing the generated pages

## Open questions

- Auto-fetch the page `<title>` for a submitted URL, or leave the title
  blank/optional?
- Dedupe a URL that's already filed under the target category?
- Should SMS support a `categories` command to list existing categories,
  since a phone can't casually browse the docs site mid-conversation?
- Language/framework for the service — no strong constraint yet.

## Milestones

1. Core module + data model + Markdown generator, exercised by hand-editing
   `links.yml` — no bot wiring yet.
2. Discord adapter (develop against the gateway locally, no public endpoint
   needed).
3. Twilio adapter (needs the public webhook endpoint).
4. Deploy into the homelab compose stack; wire up the git auto-push.
5. Harden auth/allowlisting.
