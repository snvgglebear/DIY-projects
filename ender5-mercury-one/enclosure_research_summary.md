# Electronics Enclosure Options — Ender 5 Plus / Mercury One / ZeroG

**Build context:** Ender 5 Plus with Mercury One CoreXY conversion (ZeroG)
**Boards:** BIGTREETECH Pi V1.2 (Quad-Core, WiFi, 1GB RAM), BIGTREETECH Octopus V1.1
**Screen:** BIGTREETECH TFT35 E3 V3.0.1 Touch Screen

---

## 1. Official ZeroG Electronics Enclosure (Configurator)

The ZeroG docs provide a configurator that generates a custom STL pack matched to your specific board combination (Octopus, BTT Pi, DIN rail mounts, etc.), rather than a generic box.

- **Best fit for:** Octopus V1.1 + BTT Pi setup — actively maintained by the ZeroG team
- Link: https://docs.zerog.one/manual/build/electronics_enclosure/printed_files
- General docs / build info: https://docs.zerog.one/manual/build/electronics_enclosure/introduction

---

## 2. "Base" Style Enclosures (double as frame support/feet)

These support the entire printer frame while also enclosing the electronics — matches what you described seeing.

| Design | Notes | Link |
|---|---|---|
| Ender 5 Plus Mercury One.1 Enclosure Top-Hat | Designed specifically for the Plus frame | https://www.printables.com/model/1099998 |
| syph3rd's ZeroG-Enclosure (GitHub) | Full enclosure repo for Ender 5 Plus/Pro Mercury builds | https://github.com/syph3rd/ZeroG-Enclosure |
| Bottom Panel Enclosure (Capsopie design, modded) | Modified for the Ender 5 Pro/Plus base | https://www.printables.com/model/896979 |

---

## 3. TFT35 Screen Mounting

**Heads up:** Most ZeroG enclosure designs assume a BTT Pi running KlipperScreen directly on its own display — not a separate TFT35 panel. You'll likely need an add-on mount if you want to keep the TFT35.

- Magnetic TFT35-E3 mount (Kazi Toad, Thangs) — mounts independently of the enclosure: https://thangs.com/designer/Kazi%20Toad/3d-model/Magnetic%20screen%20mount%20for%20BigTreeTech%20BTT%20TFT35-E3-408340
- Reference build using a custom TFT35 SPI mount alongside the standard ZeroG enclosure: https://www.3docity.com.au/blogs/3dprinting/mercury-zerog-ender-5-mod-build-info-klipper-slicer-profiles

**Consideration:** If running the BTT Pi as the main controller with KlipperScreen, the TFT35 may be redundant — worth deciding before committing to a mount design.

---

## Open Questions / Next Steps
- [ ] Decide: keep TFT35 or rely on BTT Pi's own touchscreen output via KlipperScreen
- [ ] Confirm frame dimensions/extrusion sizing against chosen base enclosure design
- [ ] Check DIN rail mount compatibility for Octopus V1.1 (ZeroG configurator should handle this)
