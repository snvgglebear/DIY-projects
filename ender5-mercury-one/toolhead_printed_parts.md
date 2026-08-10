# Mercury One.1 Toolhead — Printed Parts

Checking a Mercury One.1 toolhead parts kit (MPP) against the official BOM, for an
**E3D V6 hotend + LGX Lite extruder** on the EVA platform.

Companion docs: [`electronics_enclosure_options.md`](./electronics_enclosure_options.md) ·
[`enclosure_research_summary.md`](./enclosure_research_summary.md) (§4 is the source for this page)

> **Verification status:** `2.eva-3d.page` is blocked by this environment's egress proxy, so the
> per-part BOM below could not be checked against the official EVA pages and is carried over from
> the research summary. GitHub *is* reachable, so everything about ZeroG's own configurator and the
> E34M1 alternative was read from source repos, the build manual was read from the PDF, and the
> hole geometry in [How the toolhead mounts](#how-the-toolhead-mounts-to-the-x-rail) was measured
> directly from the STLs. Treat the part names as a checklist to verify, not as confirmed truth.

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

## Where the STL files actually live

Neither configurator covers everything, and the docs pages that *explain* parts don't host them.
Three distinct sources, and knowing which is which saves a lot of hunting:

| Source | Holds | Use it for |
|---|---|---|
| [`ZeroGDesign/MercuryOne/STLs/Toolhead/`](https://github.com/ZeroGDesign/MercuryOne/tree/main/STLs/Toolhead) | `Extruders/`, `Hotends/`, `Probes/`, `Fan_Ducts/`, `Universal/` | **The best checklist for this printer** — ZeroG's own Mercury-correct parts |
| EVA module pages, e.g. `2.eva-3d.page/drives/lgx_lite/stls/<part>.stl` | Per-page STLs for that module | Addons and anything ZeroG doesn't ship |
| [`ZeroGDesign/docs/.../assets/stl/eva2_4`](https://github.com/ZeroGDesign/docs/tree/gh-pages/docs/assets/stl/eva2_4) | ZeroG's mirror of the EVA 2.4 set | Manual download instead of the configurator |

**Parts Explained has no download links** — it is conceptual only. On EVA's site each page carries
the STLs for the parts *it* varies: the **drive page** holds the whole carriage set (top, bottom,
back, universal face, duct, tension sliders, cable parts), the **hotend page** holds just the three
adapter pieces, and addons live on their own pages.

### Two parts that fall through the cracks

Neither is a configurator option, so both can be missing from a kit with nothing obviously wrong:

- **Back/front plates** — `Universal/a_Eva_Backplate.stl` and `a_Eva_FrontPlate.stl` in the
  MercuryOne repo. ZeroG uses these *instead of* stock EVA's `back_corexy` and `universal_face`.
  The same folder holds `Belt_clamp_x2.stl` (the belt grabbers), `Eva_RearCableArm.stl`, and
  `a_X_Limit_Stop_Block.stl`.
- **Hotend fan shroud** — not in the Mercury repo at all; `Hotends/` has only `v6_face`,
  `v6_face_clamp`, `v6_support`. The shroud is an EVA **addon**:
  `2.eva-3d.page/addons/shrouds/stls/v6_shroud.stl`
  ([mirror](https://github.com/EVA-3D/eva-2/tree/master/addons/shrouds/stls)). Variants there
  include `shroud_adxl.stl` with an ADXL345 mount built in — worth taking if input shaping is
  coming later.

Layer-cooling duct alternates (TriHorn Default / High Narrow / UHF / Safe) are likewise on an addon
page, though ZeroG ships its own `Fan_Ducts/EVA2_4_Trihorn_2-6-narrow.stl` and `-wide.stl` — prefer
those on the ZeroG plates.

> The `a_X_Limit_Stop_Block.stl` is not optional. ZeroG's manual (p149): *"Your toolhead should
> include a mechanism to trigger the X axis endstop, often called a stop block."*

---

## How the toolhead mounts to the X rail

**The top part bolts to the MGN12H carriage.** For this build that's `Lgx_Lite_Top.stl` (stock EVA
calls it `top_mgn12`). Not the back plate — the back varies by *motion system*, the top is the part
"specified by the preferred drive **and MGN carriage type and size**," which is why the rail size is
in its filename.

Measured from `Lgx_Lite_Top.stl` — four holes in the MGN12H bolt pattern:

| | |
|---|---|
| Hole centres | `(-10, -2)` `(10, -2)` `(-10, 18)` `(10, 18)` mm |
| Spacing | **20.0 × 20.0 mm**, 28.3 mm diagonal (= 20√2) |
| Bore | 3.2–3.3 mm — M3 clearance |
| Counterbore | 6.0 mm, for M3 socket-head cap screws |

`a_Eva_Backplate.stl` has no matching pattern (its holes sit at 30 / 34 / 48 / 58 / 62 mm), which
confirms it isn't the part meeting the rail.

Resulting stack:

```
MGN12H carriage
  └── Lgx_Lite_Top          ← 4× M3 into the carriage's own threads
        ├── LGX Lite extruder on top
        ├── front plate → v6_face + v6_face_clamp + v6_support
        ├── back plate behind → belt attachment
        └── bottom + duct underneath
```

Belts never touch the carriage. They terminate on the toolhead: belt end into the **belt catch**
ribs-facing-out (top-left and bottom-left), a printed **belt grabber** pressed over both, one
**M3×8 mm bolt** each — finger-tighten, align the belt ends, then tighten fully. So belt tension
loads those four carriage screws.

### Heat-set inserts

Measured hole profiles through `Lgx_Lite_Top.stl`:

| Feature | Profile | Verdict |
|---|---|---|
| 4× carriage holes | 6.0 mm counterbore, then 3.2–3.3 mm through | **No insert** — screws thread into the steel carriage |
| 2× long cross holes | 6.18 mm recess, then 3.5 mm clearance ~26.5 mm deep | **No insert** — clearance for a through-screw |
| 1× blind bore | 4.15 mm dia × 7 mm deep, does not break through | **One M3 heat-set insert** |

So the part takes **one** insert, and it is *not* part of the carriage mount. The bolts holding the
toolhead to the rail pass through clean and bite the carriage's own M3 threads.

Caveat: this scan covered bores aligned to the three principal axes, so an angled or very small
pocket could have been missed. Check the physical part before heating anything.

### Assembly gotchas

- **ZeroG's manual does not cover toolhead assembly.** Page 149: *"You'll need to assemble your
  toolhead of choice until the belts are ready for installation. Once that's done, we'll pick up
  from there!"* It resumes at belt installation. EVA's own guide is the authority for the bolt-down.
- **Screw length differs between the two hole pairs.** The `y = -2` pair has ~11 mm of counterbore
  above ~4 mm of clearance; the `y = 18` pair sits in much thinner material. Measure per pair rather
  than buying one length for all four, and don't bottom out in the carriage block.
- **Proud heat-set inserts cause gaps.** The manual raises this for the X joints (p125): an insert
  standing above the surface holds the printed part off the MGN12H carriage. Same hazard here — if
  the top won't sit flat, fix the insert rather than torquing it down.
- **Mount the toolhead before threading belts,** then put long M3 screws through the free holes
  either side of the carriage to hold the gantry square while routing — threading pulls hard on the
  toolhead (p156).

The full build manual is worth having open —
[187-page PDF](https://docs.zerog.one/assets/mercury_one_1_instruction_18-02-2024.pdf); belt routing
reference diagrams are p152–153.

---

## Assembly instructions: where to look

**There is no official EVA 2.4 assembly manual.** The site's navigation is Introduction, Features
and Roadmap, Parts Explained, Printing, Hotends, Drives, Addons, STEP Files and Licence, Community
Contributions, Queen Shroud — no assembly section anywhere. Each module page carries exactly three
blocks: Description, Bill of Materials, Links. No steps, no exploded views, no video.

ZeroG's manual p149 shows a link captioned *"EVA 2.4 Docs & Assembly Guide"*, but its actual target
is `docs.zerog.one/manual/build/mercury_eva/printed_files#eva-24-toolhead-configurator` — their own
configurator anchor. Following it lands back at the file downloads. It is a documentation dead end,
not something you've failed to find.

What to use instead, roughly in order of usefulness:

| Source | What it gives | Caveat |
|---|---|---|
| **OnShape CAD assembly** — linked from each EVA module page's Links block. LGX Lite: [`cad.onshape.com/documents/afc780380bd2317241ddf83e/…`](https://cad.onshape.com/documents/afc780380bd2317241ddf83e/w/e822a01920872e9f27d38958/e/37c09e63f33fc6f60a845257) | The live model — rotate, section, see exactly where each fastener lands and which faces mate | EVA's own warning: bleeding-edge version, may differ from released STLs |
| **Rat Rig Dozuki guides** — `ratrig.dozuki.com`, "07. EVA Assembly" (guide 117), "08. EVA 2 Carriage: BMG + E3D V6" (guide 72) | The de facto illustrated manual; Rat Rig ships EVA on the V-Core 3 | Written for BMG + V6 on a V-Core — drive-specific steps differ, carriage stack-up is identical |
| **Older EVA docs** — `pkucmus.github.io/EVA`, per-combination *assemblies* pages (`hemera-mgn12corexy`, `bowden-mgn12corexyv6`, `bmg-mgn12v6corexy`) | Assembly walkthroughs for the general stack-up | Earlier EVA generation |
| **[E34M1 docs](https://jon-harper.github.io/E34M1/)** | Ships a real assembly guide, and it's for this exact printer | EVA 3, not 2.4 — orientation only unless you switch platforms |
| **ZeroG Discord** — linked from their docs' introduction | Fastest route for Mercury-specific questions | — |

### The gap none of these close

**No EVA guide covers `a_Eva_Backplate` or `a_Eva_FrontPlate`** — those are ZeroG's own Mercury
parts, not upstream EVA, and the readme in `MercuryOne/STLs/Toolhead/` is still marked "Soon™". For
those plates specifically, the OnShape model, the Discord, or a community Mercury build video is
all there is.

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
- [MERCURY ONE.1 build manual, 187-page PDF — ZeroG](https://docs.zerog.one/assets/mercury_one_1_instruction_18-02-2024.pdf) — toolhead handoff p149, belt prep p155–156, X joint insert warning p125
- [ZeroGDesign/MercuryOne — official Mercury STL repo](https://github.com/ZeroGDesign/MercuryOne/tree/main/STLs/Toolhead)
- [EVA-3D/eva-2 — built EVA 2.4 docs site, per-module STLs](https://github.com/EVA-3D/eva-2)
- [EVA 2.4 Parts Explained](https://2.eva-3d.page/parts_explained/)
- [LGX Lite live CAD assembly — OnShape](https://cad.onshape.com/documents/afc780380bd2317241ddf83e/w/e822a01920872e9f27d38958/e/37c09e63f33fc6f60a845257)
- [Rat Rig EVA assembly guides — Dozuki](https://ratrig.dozuki.com/Guide/07.+EVA+Assembly/117)
- [Older EVA docs with per-combination assembly pages](https://pkucmus.github.io/EVA/)
- [Mercury One.1 build instructions — ZeroG Documentation](https://docs.zerog.one/manual/build/mercury_eva/build_instruction)
- [ZeroG docs source (printed_files page read from here)](https://github.com/ZeroGDesign/docs/blob/gh-pages/docs/pages/manual/build/mercury_one_1/printed_files.md)
- [E34M1 — EVA 3 for Mercury One.1 (docs)](https://jon-harper.github.io/E34M1/) · [repo](https://github.com/jon-harper/E34M1)
- [Zero G Mercury 1.1 printed parts by MPP — Fabreeko](https://www.fabreeko.com/products/zero-g-mercury-1-1-printed-parts-by-honeybadger)
