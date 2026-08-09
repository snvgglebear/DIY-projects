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

---

## 4. Toolhead Printed Parts Checklist (E3D V6 + LGX Lite, EVA 2.4 / MPP Kit)

ZeroG's Mercury 1.1 toolhead is built on the **EVA 2.4** modular platform. The toolhead is assembled from separate modules (Drive, Hotend, Probe), each with its own printed-parts list. Below is the full official BOM for your combo — cross-check this against what's in your MPP kit box to spot what's missing.

Source: EVA 2.4 official docs — https://2.eva-3d.page/

### Drive module — LGX Lite (CoreXY, MGN12)
Source: https://2.eva-3d.page/drives/lgx_lite/

| Qty | Printed Part |
|---|---|
| 1 | top_mgn12 |
| 1 | universal_face |
| 2 | face_belt_grabber |
| 1 | bottom_mgn12_short_duct |
| 1 | TriHorn Duct (Default) |
| 1 | back_corexy |
| 2 | tension_slider_9mm_belt_M5 |
| 1 | cable_holder |
| 1 | universal_cable_mount |

### Hotend module — E3D V6
Source: https://2.eva-3d.page/hotends/v6/

| Qty | Printed Part |
|---|---|
| 1 | v6_face |
| 1 | v6_face_clamp |
| 1 | v6_support |

### Probe module — BL-Touch (if applicable)
Source: https://2.eva-3d.page/addons/probes/bl_touch/
*(Your kit ships with both BLTouch and inductive mounts per the MPP listing — you only need one)*

| Qty | Printed Part |
|---|---|
| 1 | bl_touch_mount |
| 1 | bl_touch_mount_alt (alternate orientation — pick one) |

### Probe module — Inductive (if applicable)
Check: https://2.eva-3d.page/addons/probes/inductive/ for the equivalent inductive mount part name/count.

### Total unique toolhead printed parts: ~12–13 pieces
(top_mgn12, universal_face, face_belt_grabber ×2, bottom_mgn12_short_duct, TriHorn duct, back_corexy, tension_slider ×2, cable_holder, universal_cable_mount, v6_face, v6_face_clamp, v6_support, + 1 probe mount)

### Notes
- **Non-printed hardware also required per BOM:** 4010 fan (hotend), 5015 fan (part cooling), M3 screws in various lengths (8/10/12/20/25/35mm), M3 hex nuts, M5×45mm screws + nyloc nuts (tension sliders), PTFE tube (sum of drive + hotend BOM lengths — do not cut into two pieces).
- The **TriHorn duct** and **cable management parts** (cable_holder, universal_cable_mount) are easy to overlook since they're not "core" front/hotend pieces — check these first if something feels missing.
- If you're running a remote toolhead board (BTT EBB36/EBB42, etc.) instead of wiring straight to the Octopus, that mount is **not part of the standard EVA BOM** — it's a community addon and would need to be sourced separately (see prior search results for community EBB36/EBB42 LGX Lite mounts).
