# DIY Projects

Documentation and research for DIY projects.

📖 **Browse as a website:** <https://snvgglebear.github.io/DIY-projects/>

## Projects

### Ender 5 Plus / Mercury One.1

An Ender 5 Plus converted to CoreXY with the ZeroG Mercury One.1 kit, running a
BTT Octopus V1.1, BTT Pi V1.2 and TFT35-E3 V3.0.1.

- [Electronics enclosure options](ender5-mercury-one/electronics_enclosure_options.md) — official ZeroG enclosure and configurator, community base/skirt and frame-mount alternatives, DIN rail interiors, screen mounting, hardware list
- [Toolhead printed parts](ender5-mercury-one/toolhead_printed_parts.md) — EVA 2.4 BOM for E3D V6 + LGX Lite, ZeroG's toolhead configurator, CAN vs direct wiring
- [Multi-Z (quad Z) mod](ender5-mercury-one/quad_z_mod.md) — four-motor bed tramming: brackets, ZeroG Hydra, Klipper `z_tilt` and Marlin `G34` references
- [Enclosure research summary](ender5-mercury-one/enclosure_research_summary.md) — short-form research notes

### Drone build

Compact camera drone with a 3D printed propeller-guard frame, dual-band 915 MHz
comms and a smartphone ground station.

- [Build guide](drone-build/drone_build_guide.md) — full assembly, wiring, configuration and first-flight checklist
- [Purchase list](drone-build/drone_purchase_list.md) — parts, costs and sourcing
- [Kit alternatives](drone-build/drone_kit_alternatives.md) — hybrid kit-vs-parts approach for the core electronics
- [Frame merge guide (Onshape)](drone-build/drone_frame_merge_guide.md) — merging a multi-part frame design into fewer large prints
- [LoRa still-image uplink](drone-build/drone_lora_still_image_uplink.md) — experimental low-bandwidth image downlink

### Homelab

Dashboard stack for a home server running the *arr stack behind a VPS and
WireGuard tunnel.

- [Dashboard stack](homelab/README.md) — Homepage + Authelia compose setup
- [Dashboard options](homelab/dashboard_options.md) — comparison and recommendation

### Link collection bot

A Discord/Twilio bot that files a submitted link + category into a links
page on this site, prompting to confirm before creating a new category.

- [Plan](link-bot/plan.md) — architecture, data model, category confirmation flow, deployment

## Building the site locally

The site is [MkDocs](https://www.mkdocs.org/) with the
[Material](https://squidfunk.github.io/mkdocs-material/) theme. Markdown files
stay where they are in the repo — the `same-dir` plugin points MkDocs at the
repo root instead of a `docs/` directory, so links work both on GitHub and on
the site.

```bash
python3 -m venv venv
./venv/bin/pip install -r requirements-docs.txt
./venv/bin/mkdocs serve      # http://127.0.0.1:8000
```

To add a page: drop the Markdown file anywhere in the repo and add it to the
`nav:` list in [`mkdocs.yml`](https://github.com/snvgglebear/DIY-projects/blob/main/mkdocs.yml).
Builds run with `--strict`, so a broken internal link or a page missing from
`nav` fails CI rather than shipping quietly.

Pushes to `main` deploy via
[`.github/workflows/docs.yml`](https://github.com/snvgglebear/DIY-projects/blob/main/.github/workflows/docs.yml);
pull requests build without deploying.
