# Mercury One.1 Toolhead — Printed Parts

Checking a Mercury One.1 toolhead parts kit (MPP) against the official BOM, for an
**E3D V6 hotend + LGX Lite extruder** on the EVA platform.

Companion docs: [`electronics_enclosure_options.md`](./electronics_enclosure_options.md) ·
[`enclosure_research_summary.md`](./enclosure_research_summary.md) (§4 is the source for this page)

> **Verification status:** `2.eva-3d.page` is blocked by this environment's egress proxy, so the
> per-part BOM below could not be checked against the official EVA pages and is carried over from
> the research summary. GitHub *is* reachable, so everything about ZeroG's own configurator and the
> E34M1 alternative was read from source repos. Treat the part names as a checklist to verify, not
> as confirmed truth.

---

## Use ZeroG's configurator, not the generic EVA BOM

The summary points at [2.eva-3d.page](https://2.eva-3d.page/) for the parts list. That's the right
upstream source for EVA generally, but for this printer there's a better one:

**ZeroG ships its own EVA 2.4 configurator**, on the same page as the Mercury configurator:
<https://docs.zerog.one/manual/build/mercury_eva/printed_files>

It generates a zip of all the STLs for your print head, with options for **hotend, extruder and
probe** — which is exactly the three-way choice the summary's checklist enumerates by hand. Per
ZeroG's docs, the files you get "are made by ZeroG to fit the Mercury 1.1 or by Paweł Kucmus of
EVA," meaning **some parts are Mercury-specific variants rather than stock EVA**. Generating the
pack from the configurator gets the right variant automatically; working from the generic EVA BOM
risks printing a stock part where the Mercury one differs.

Two configurators live on that page — don't mix them up:

| Configurator | Generates | Options |
|---|---|---|
| Mercury One.1 | Frame/motion STLs | Z-adapters, toothed idler, tension plate (tapped or T-nut) |
| EVA 2.4 | Toolhead STLs | Hotend, extruder, probe |

**Frame size is not a configurator option.** Ender 5 (Pro) and Ender 5 Plus are handled as separate
downloadable STEP files on that page — check you're pulling the **Plus** set.

---

## Parts checklist

From the research summary, sourced from the official EVA 2.4 docs. Cross-check against what's in
the kit box.

### Drive module — LGX Lite (CoreXY, MGN12)

<https://2.eva-3d.page/drives/lgx_lite/>

| Qty | Printed part |
|---|---|
| 1 | `top_mgn12` |
| 1 | `universal_face` |
| 2 | `face_belt_grabber` |
| 1 | `bottom_mgn12_short_duct` |
| 1 | TriHorn Duct (default) |
| 1 | `back_corexy` |
| 2 | `tension_slider_9mm_belt_M5` |
| 1 | `cable_holder` |
| 1 | `universal_cable_mount` |

### Hotend module — E3D V6

<https://2.eva-3d.page/hotends/v6/>

| Qty | Printed part |
|---|---|
| 1 | `v6_face` |
| 1 | `v6_face_clamp` |
| 1 | `v6_support` |

### Probe module — pick one

The MPP kit ships both BLTouch and inductive mounts; you only need the one matching your probe.

| Probe | Parts | Source |
|---|---|---|
| BL-Touch | `bl_touch_mount`, plus `bl_touch_mount_alt` as an alternate orientation — print one | <https://2.eva-3d.page/addons/probes/bl_touch/> |
| Inductive | Check the page for the equivalent mount name and count | <https://2.eva-3d.page/addons/probes/inductive/> |

**Total unique toolhead printed parts: ~12–13**, counting the doubled parts once.

---

## Easy things to miss

- **The TriHorn duct and the cable-management parts** (`cable_holder`, `universal_cable_mount`) are
  the usual omissions — they don't look like "core" hotend pieces, so check them first if the kit
  feels short.
- **Non-printed hardware is part of the BOM too:** a 4010 fan (hotend), a 5015 fan (part cooling),
  M3 screws in 8/10/12/20/25/35 mm, M3 hex nuts, M5×45 mm screws with nyloc nuts for the tension
  sliders, and PTFE tube. Cut the PTFE as the **sum** of the drive and hotend BOM lengths — a single
  piece, not two.
- **A remote toolhead board mount is not in the EVA BOM.** If you run a BTT EBB36/EBB42 on CAN
  instead of wiring the toolhead back to the Octopus, that mount is a community addon and has to be
  sourced separately.

---

## The CAN decision affects the enclosure

Worth settling before you print either the toolhead or the electronics bay, because it changes both:

- **Direct wiring to the Octopus** means the full toolhead loom — steppers, heater, thermistor,
  fans, probe — runs from the bay up a rear upright to the top gantry. That's a thick bundle to
  route and strain-relieve, and it sets how much cable-entry room the enclosure needs. See the
  Mercury-specific notes in
  [`electronics_enclosure_options.md`](./electronics_enclosure_options.md#hardware-constraints).
- **CAN toolhead (EBB36/EBB42)** reduces that to a four-wire umbilical, which is a real
  simplification on a CoreXY conversion where the gantry moves in both axes. Cost: an extra board, a
  CAN transceiver on the Octopus side, and a toolhead mount that isn't in the EVA BOM.

---

## EVA 3 alternative: E34M1

- Docs: <https://jon-harper.github.io/E34M1/>
- Repo: <https://github.com/jon-harper/E34M1>
- Printables: <https://www.printables.com/model/386043-e34m1-eva-3-for-mercury-one1>

E34M1 is an overhaul of **EVA 3** reworked to fit the Mercury One.1, with its own assembly
documentation. ZeroG's own toolhead is EVA **2.4**, so this is a deliberate departure from the kit
rather than a drop-in — but it's the actively developed path if you'd rather be on EVA 3.

Only worth considering **before** you print the 2.4 parts. If your MPP kit is already EVA 2.4 and
complete, finishing that build is the cheaper move; revisit E34M1 at your next toolhead rebuild.

---

## Sources

- [EVA 2.4.2 documentation](https://2.eva-3d.page/) — [LGX Lite drive](https://2.eva-3d.page/drives/lgx_lite/) · [V6 hotend](https://2.eva-3d.page/hotends/v6/) · [BL-Touch probe](https://2.eva-3d.page/addons/probes/bl_touch/) · [inductive probe](https://2.eva-3d.page/addons/probes/inductive/)
- [Mercury One.1 printed files + both configurators — ZeroG Documentation](https://docs.zerog.one/manual/build/mercury_eva/printed_files)
- [Mercury One.1 build instructions — ZeroG Documentation](https://docs.zerog.one/manual/build/mercury_eva/build_instruction)
- [ZeroG docs source (printed_files page read from here)](https://github.com/ZeroGDesign/docs/blob/gh-pages/docs/pages/manual/build/mercury_one_1/printed_files.md)
- [E34M1 — EVA 3 for Mercury One.1 (docs)](https://jon-harper.github.io/E34M1/) · [repo](https://github.com/jon-harper/E34M1)
- [Zero G Mercury 1.1 printed parts by MPP — Fabreeko](https://www.fabreeko.com/products/zero-g-mercury-1-1-printed-parts-by-honeybadger)
