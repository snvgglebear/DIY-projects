# Ender 5 Plus — Multi-Z (Quad Z) Self-Levelling Mod

Reference list for converting the Ender 5 Plus bed to **four independently driven Z motors** so the
printer can tram its own bed. Written for a Mercury One.1 CoreXY conversion running Klipper on a
BTT Octopus V1.1, with the motors and hardware already on hand.

Companion docs: [`electronics_enclosure_options.md`](./electronics_enclosure_options.md) ·
[`toolhead_printed_parts.md`](./toolhead_printed_parts.md)

> **Verification status:** Printables, `klipper3d.org`, `marlinfw.org` and `docs.zerog.one` are all
> blocked by this environment's egress proxy. Model listings were confirmed to exist via search but
> not opened. Everything quoted from Klipper and Marlin was read from their **GitHub sources**
> (`docs/Config_Reference.md`, `docs/G-Codes.md`, `Marlin/Configuration_adv.h`), and the ZeroG Hydra
> details from ZeroG's own docs repo. Treat model specifics as leads to check; treat the firmware
> quotes as solid.

---

## Read this first: it's `z_tilt`, not `quad_gantry_level`

On the Ender 5 Plus the **bed moves in Z** and the gantry is fixed — the Mercury One.1 conversion
doesn't change that. So even with four Z motors, this is **not** quad gantry levelling. Klipper's
own reference is explicit:

> `[quad_gantry_level]` — Moving gantry leveling using 4 independently controlled Z motors. Corrects
> hyperbolic parabola effects (potato chip) on moving gantry which is more flexible.
> **WARNING: Using this on a moving bed may lead to undesirable results.**

QGL is a Voron 2.4-style feature for printers whose *gantry* hangs from four belts. For a bed
mover with 4 Z screws you want **`[z_tilt]`** and the `Z_TILT_ADJUST` command, which works with any
number of Z steppers. This is the single most common mistake when people copy a "quad Z" config
from a Voron.

---

## Consider 3 Z motors before 4

You have parts for four, so this is a note rather than an objection — but it's worth five minutes
of thought before drilling anything.

**Three points define a plane.** A 3-motor bed is kinematically determinate: each motor has exactly
one job, and `Z_TILT_ADJUST` converges cleanly. **Four points over-constrain the bed** — the
correction can fight itself, and a bed that isn't perfectly flat gets twisted rather than levelled.
Four Z is standard on *gantry* movers (where the frame is the rigid element) and much rarer on bed
movers for exactly this reason.

Four still works, and plenty of people run it. Just expect to care more about bed flatness and about
`retry_tolerance` convergence than a 3-motor build would.

Notably, ZeroG's own supported multi-Z upgrade for this printer — **Hydra**, below — is **three**
motors, not four.

---

## Mechanical: parts and models

### Ender 5 Plus Build Plate Mounting Bracket for Dual/Triple/Quad Z

<https://www.printables.com/model/484885-ender-5-plus-build-plate-mounting-bracket-for-dual>

The key part, and Ender 5 Plus-specific. Near drop-in replacement for the stock bracket, with a
stiffening ridge added so it can be printed in any filament, plus a spacer that raises the plate
about 10 mm. **For quad Z it uses two Z motors at the outer positions on each side.** The lead screw
fastens as per the original bracket, or — if you're running anti-backlash nuts — screws pass up from
below with nuts on top. Bearer mounts take M4 heat-set inserts or M4 screws and nuts.

Its listing also carries lead-screw positions for each configuration, which is what you need for
`z_positions` below.

### Ender 5 Triple Z (E5TZ v2) — for CoreXY-modded Ender 5

<https://www.printables.com/model/475994-ender-5-triple-z-e5tz-v2>

A triple-Z design explicitly aimed at CoreXY-converted Ender 5s, so it's the closest community
analogue to a Mercury build. Useful reference even if you commit to four.

### ZeroG Hydra — the official multi-Z upgrade

- Docs: <https://docs.zerog.one/manual/build/hydra/introduction>
- STL repo: <https://github.com/ZeroGDesign/Hydra>
- Docs source: <https://github.com/ZeroGDesign/docs/tree/gh-pages/docs/pages/manual/build/Hydra>
- Video series: [Part 1 — Teardown & Skirts](https://www.youtube.com/watch?v=30WTpSBm1eg) ·
  [Part 2 — Installing Hydra & Bed](https://www.youtube.com/watch?v=pW03FeI4y7Q) ·
  [Part 3 — Wiring & Firmware](https://www.youtube.com/watch?v=XkIQxIgUmeY)
- Community remix: [Hydra remix for Oldham couplers](https://www.printables.com/model/451834-zero-g-hydra-remix-for-use-of-oldham-couplers)

ZeroG's own Z-axis overhaul for the Ender 5 (Pro) and **Ender 5 Plus**, built around three arms —
the docs are split into `left_arm.md`, `right_arm.md`, `rear_arm.md` plus `bed_drawings.md`. Its
bill of materials calls for:

| Qty | Item |
|---|---|
| 3 | NEMA17 steppers |
| 3 | TR8 lead screws — 350–370 mm (Ender 5 Pro), **450–470 mm (Plus)** |
| 3 | Linear rails, MGN9C or MGN12C/H depending on model |
| 3 | Shaft couplings |
| ~119 | Fasteners — M3/M4/M5 bolts, T-nuts, heat-set inserts, washers, spacers |
| — | Aluminium bed plate, heater, solid-state relay |

It exists to kill Z wobble as much as to enable auto-tramming, and it needs a board with more
drivers than stock — your Octopus V1.1 is fine. The Hydra also has a matching deck panel in the
ZeroG electronics enclosure (`STLs/Plus/Hydra_Deck_Panel/`), so the two upgrades are designed to
coexist.

*Inferred from the BOM and the docs file structure — ZeroG's introduction page doesn't state the
system's purpose in words, and the full docs pages were not reachable from here.*

---

## Klipper configuration

### Extra steppers

Each additional Z motor is its own section. From Klipper's `Config_Reference.md`:

> **`[stepper_z1]`** — Multi-stepper axes. On a cartesian style printer, the stepper controlling a
> given axis may have additional config blocks defining steppers that should be stepped in concert
> with the primary stepper. One may define any number of sections with a numeric suffix starting at
> 1 (for example, "stepper_z1", "stepper_z2", etc.).

So four Z motors = `[stepper_z]` + `[stepper_z1]` + `[stepper_z2]` + `[stepper_z3]`. Each takes
`step_pin`, `dir_pin`, `enable_pin`, `microsteps`, `rotation_distance`, and optionally its own
`endstop_pin`:

> If an endstop_pin is defined for the additional stepper then the stepper will home until the
> endstop is triggered. Otherwise, the stepper will home until the endstop on the primary stepper
> for the axis is triggered.

Give the extra steppers `endstop_pin` **only** if you're fitting four independent Z endstops.
Otherwise leave it out and let all four home together on the single Z endstop, then let
`Z_TILT_ADJUST` do the tramming.

### `[z_tilt]`

```
[z_tilt]
z_positions:
#   A list of X, Y coordinates (one per line; subsequent lines indented) describing
#   the location of each bed "pivot point" — the point where the bed attaches to the
#   given Z stepper. Described using nozzle coordinates. The first entry corresponds
#   to stepper_z, the second to stepper_z1, the third to stepper_z2, etc.
points:
#   A list of X, Y coordinates that should be probed during a Z_TILT_ADJUST command.
speed: 50
horizontal_move_z: 5
retries: 0
retry_tolerance: 0
```

Key points from the reference:

- **`z_positions` are the screw/pivot locations expressed in nozzle coordinates** — where the nozzle
  would sit if it could travel directly above each lead screw. They are usually **off the bed**
  (negative, or beyond the bed size), which is expected and correct.
- **Order matters**: entry 1 = `stepper_z`, entry 2 = `stepper_z1`, and so on. Getting this wrong
  makes the bed diverge instead of converge.
- `points` are real probe locations on the bed, and must be reachable by the *probe*, not the nozzle.
- Set `retries: 5` or so with a `retry_tolerance` around `0.01`–`0.02` once it's working.

A community quad-Z example of the `z_positions` shape (2 left, 2 right):

```
z_positions:
    -18.4, 74.4
    -18.4, 285.6
    378.4, 285.6
    378.4, 74.4
```

Use your own bracket's screw coordinates — the numbers above are illustrative, not Ender 5 Plus
values.

### The command

> **`Z_TILT_ADJUST [RETRIES=<value>] [RETRY_TOLERANCE=<value>] [HORIZONTAL_MOVE_Z=<value>]`** — This
> command will probe the points specified in the config and then make independent adjustments to
> each Z stepper to compensate for tilt.

Call it after `G28` in `PRINT_START`, before any bed mesh.

### Driver budget on the Octopus V1.1

Eight drivers. Four Z + two XY + one extruder = **seven**, so a quad-Z Mercury fits with one slot
spare (or two if you move the extruder to a CAN toolhead board). Use identical motors on all four Z
positions and set the same current for each.

---

## Marlin configuration

If you're on Marlin rather than Klipper. Quoted from `Marlin/Configuration_adv.h` (bugfix-2.1.x):

- **`NUM_Z_STEPPER_DRIVERS`** — set to 1–4.
- **`Z_STEPPER_AUTO_ALIGN`** — *"Add the G34 command to align multiple Z steppers using a bed probe."*
- **`Z_STEPPER_ALIGN_XY`** — probe X/Y positions for Z1, Z2 [, Z3 [, Z4]]. *"These positions are
  machine-relative and do not shift with the M206 home offset! If not defined, probe limits will be
  used. Override with `M422 S<index> X<pos> Y<pos>`."*
- **`Z_STEPPERS_ORIENTATION`** — only if you don't define `Z_STEPPER_ALIGN_XY`. The header's own
  4-stepper diagram:

```
 * 4 Steppers:  (0)     (1)     (2)     (3)
 *               | 4   3 | 1   4 | 2   1 | 3   2 |
 *               |       |       |       |       |
 *               | 1   2 | 2   3 | 3   4 | 4   1 |
```

- **`Z_STEPPER_ALIGN_STEPPER_XY`** — *"Z Stepper positions for more rapid convergence in bed
  alignment. Requires 3 or 4 Z steppers. Define Stepper XY positions for Z1, Z2, Z3... corresponding
  to the screw positions in the bed carriage, with one position per Z stepper in stepper driver
  order."* This is Marlin's equivalent of Klipper's `z_positions`, and it materially speeds up
  convergence on a 4-motor bed.
- **`G34`** runs the alignment; **`M422`** sets positions at runtime.

⚠ **Known Marlin bug with your board:**
[MarlinFirmware/Marlin#23012](https://github.com/MarlinFirmware/Marlin/issues/23012) —
*"Settings NUM_Z_STEPPER_DRIVERS=4 on BTT_OCTOPUS_V1 will not include Z3."* Check this is resolved
in whatever version you build before committing to Marlin for a quad-Z Octopus setup. Klipper avoids
the issue entirely.

---

## Reference list

**Mechanical**

- [Ender 5 Plus Build Plate Mounting Bracket for Dual/Triple/Quad Z — Printables](https://www.printables.com/model/484885-ender-5-plus-build-plate-mounting-bracket-for-dual)
- [Ender 5 Triple Z (E5TZ v2), CoreXY-modded Ender 5 — Printables](https://www.printables.com/model/475994-ender-5-triple-z-e5tz-v2)
- [ZeroG Hydra — STL repo](https://github.com/ZeroGDesign/Hydra) · [docs](https://docs.zerog.one/manual/build/hydra/introduction) · [docs source](https://github.com/ZeroGDesign/docs/tree/gh-pages/docs/pages/manual/build/Hydra)
- [Hydra remix for Oldham couplers — Printables](https://www.printables.com/model/451834-zero-g-hydra-remix-for-use-of-oldham-couplers)
- [Ender 5 Plus linear rail mods — Printables](https://www.printables.com/model/73393-ender-5-plus-linear-rail-mods)

**Hydra build videos**

- [Part 1 — Teardown & Skirts](https://www.youtube.com/watch?v=30WTpSBm1eg)
- [Part 2 — Installing Hydra & Bed](https://www.youtube.com/watch?v=pW03FeI4y7Q)
- [Part 3 — Wiring & Firmware](https://www.youtube.com/watch?v=XkIQxIgUmeY)

**Klipper**

- [Config reference — `[z_tilt]`, `[stepper_z1]`, `[quad_gantry_level]`](https://www.klipper3d.org/Config_Reference.html) · [source on GitHub](https://github.com/Klipper3d/klipper/blob/master/docs/Config_Reference.md)
- [G-Codes — `Z_TILT_ADJUST`](https://www.klipper3d.org/G-Codes.html) · [source](https://github.com/Klipper3d/klipper/blob/master/docs/G-Codes.md)
- [Klipper Discourse — "Z_TILT for 4 stepper_z"](https://klipper.discourse.group/t/z-tilt-for-4-stepper-z/17515)
- [Klipper Discourse — storing Z adjustments between prints](https://klipper.discourse.group/t/z-tilt-adjust-store-z-adjustments/7329)
- [klipper-z-tramming — macro helper](https://github.com/Jomik/klipper-z-tramming)

**Ender 5 Plus Klipper configs to crib from**

- [RandieBarsteward/Ender-5-Plus-Klipper](https://github.com/RandieBarsteward/Ender-5-Plus-Klipper) — includes an SKR 3 EZ variant
- [CyberMODE/Klipper---Ender-5-Plus](https://github.com/CyberMODE/Klipper---Ender-5-Plus)
- [SaturnsVoid/Ender5PlusKlipperConfig](https://github.com/SaturnsVoid/Ender5PlusKlipperConfig)
- [ethomasgt/Ender-5-Plus-Mercury-One-Klipper-Mainsail](https://github.com/ethomasgt/Ender-5-Plus-Mercury-One-Klipper-Mainsail) — Mercury One.1 conversion config

**Marlin**

- [G34 Z Steppers Auto-Alignment](https://marlinfw.org/docs/gcode/G034-zsaa.html)
- [M422 Set Z Motor XY](https://marlinfw.org/docs/gcode/M422.html)
- [`Configuration_adv.h` source — `Z_STEPPER_AUTO_ALIGN` block](https://github.com/MarlinFirmware/Marlin/blob/bugfix-2.1.x/Marlin/Configuration_adv.h)
- [Marlin #23012 — `NUM_Z_STEPPER_DRIVERS=4` on BTT Octopus V1 omits Z3](https://github.com/MarlinFirmware/Marlin/issues/23012)

---

## Open questions / next steps

- [ ] Decide 3 vs 4 Z motors — Hydra (3, ZeroG-supported, includes bed and rails) vs the quad-Z
      bracket (4, uses the parts already on hand)
- [ ] Measure the actual lead-screw positions on the chosen bracket and convert to nozzle
      coordinates for `z_positions`
- [ ] Confirm lead screw length for the Plus — Hydra's BOM calls for 450–470 mm TR8
- [ ] Check Z motor current and that all four are identical parts
- [ ] Decide single shared Z endstop vs four independent endstops before wiring
- [ ] Add `Z_TILT_ADJUST` to `PRINT_START` ahead of bed mesh, with `retries` and `retry_tolerance` set
