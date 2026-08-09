# Ender 5 Plus + ZeroG Mercury One.1 — Electronics Enclosure Options

Research notes for housing a BTT Octopus V1.1 + BTT Pi V1.2 + TFT35-E3 V3.0.1 on an
Ender 5 Plus converted to CoreXY with the ZeroG Mercury One.1 kit.

> **Verification status:** all model links below were found through search and the titles/URLs
> are real, but the hosting sites (Printables, Thingiverse, Cults3D, docs.zerog.one) are blocked
> from this environment, so I could not open each page and confirm dimensions from the STL
> descriptions myself. Details attributed to a model come from its listing text. **Measure
> against the numbers in [Hardware constraints](#hardware-constraints) before printing 20+ hours
> of panels.**

---

## Hardware constraints

These are the numbers every option has to satisfy.

| Component | Footprint | Mounting | Notes |
|---|---|---|---|
| BTT Octopus V1.1 | 160 × 100 mm PCB | 150 × 90 mm hole pattern, M3 | Add ~25–30 mm height for stepper drivers + heatsinks, plus connector/wire bend room on all four edges |
| BTT Pi V1.2 | 85 × 56 mm | 64 × 49.4 mm, M2.5 — Raspberry Pi positions, but only 3 of 4 holes truly line up | Does **not** drop into most Raspberry Pi 4 cases despite the matching outline. Use a BTT-Pi-specific mount |
| TFT35-E3 V3.0.1 | ~93 × 87 mm face | Ender 3 style bracket pattern | Mounted on the frame, not inside the bay |
| Stock E5+ PSU | 230 × 127 × 40.5 mm | Stock frame brackets | The single biggest item; decide early whether it moves into the new enclosure or stays put |

**Practical minimum internal bay:** ~200 × 150 mm floor area and ~55 mm clear height for the
Octopus alone with cables dressed. Add the PSU and you need roughly 380 × 200 mm of floor.

**Mercury One.1 specifics that change the problem:**

- The conversion moves X/Y motion to a top gantry, so the motor and endstop harness now runs
  **up a rear upright** instead of across the bed. Plan a cable path from the bay to the top rear
  corner, and buy longer stepper extensions (or go CAN toolhead) before mounting anything.
- Z and the bed are unchanged, so the space under the base frame is still the natural home for
  electronics — but the **bed descends into that area at max Z**. Check clearance at Z-max before
  committing to a tall base.
- Mercury speeds mean higher sustained driver current. A sealed printed box around an Octopus is a
  thermal problem; every option below needs at least one 40–60 mm intake and a vented exhaust.
- If you later add the [Mercury One.1 enclosure top-hat](https://www.printables.com/model/1099998-ender-5-plus-mercury-one1-enclosure-top-hat),
  keeping the electronics in the base keeps them *outside* the heated chamber. That's the right
  call — but the base still soaks heat from above, so don't skip the fans.

---

## Approach A — Base / skirt that lifts the whole frame

This is the style you described: a printed plinth that the printer stands on, with the electronics
living in the void it creates. Best-looking option, most printing, and the one that most naturally
fits an Octopus + Pi + PSU together.

### A1. Ender 5 Plus Skirting and Lower Enclosure (Base Case) — *primary recommendation*

<https://www.printables.com/model/557066-ender-5-plus-skirting-and-lower-enclosure-base-cas>
(also listed as "Gen 1.0" at
<https://www.printables.com/model/557066-ender-5-plus-skirting-and-lower-enclosure-gen-10-b>)

By Big Dog Custom Creations. Raises the Ender 5 Plus frame by **70 mm** and puts the electronics in
that space, explicitly to eliminate the original Creality control box. Includes perimeter skirt
panels, screw feet (Screw Foot V3) and a display mount variant for a 7" RPi screen. Listing shows a
June 2025 update, so it's actively maintained.

- **Why it fits your build:** 70 mm of lift minus panel thickness leaves roughly 60 mm usable
  height — enough for a flat-laid Octopus with drivers, and enough for the 40.5 mm PSU lying flat.
  It's designed around the E5+ base footprint, so the Mercury conversion (which doesn't touch the
  base) doesn't invalidate it.
- **What you'll need to adapt:** it is not an Octopus/BTT-Pi-specific design. Plan on adding your
  own mounting plate inside — see [Approach C](#approach-c--din-rail-bay-inside-whatever-shell-you-pick).
- **Check before printing:** bed-at-Z-max clearance, and whether the 70 mm lift pushes the top
  gantry beyond your ceiling/shelf height. The E5+ is already ~619 mm tall stock.

### A2. Ender 5 Electronics Skirt

<https://www.printables.com/model/1518720-ender-5-electronics-skirt>

A skirt that houses all of the Ender 5's electronics in very few parts, designed to print on an
Ender 5-sized bed. **Ender 5 / 5 Pro footprint, not Plus** — the E5+ base is larger, so this needs
scaling or a remix of the corner pieces. Worth reading for the layout idea even if you don't print it.

### A3. Zero G Mercury One.1 Bottom Panel Electronics Enclosure

<https://www.printables.com/model/896979-zero-g-mercury-one1-for-ender-5-pro-bottom-panel-e>

The only model found that is purpose-built for a Mercury One.1 conversion. Caveat: it's for the
**Ender 5 Pro**, not the Plus. Useful as a reference for how others solved the Mercury cable routing,
but the panel dimensions will not transfer to the E5+ frame without rework.

---

## Approach B — External box mounted to the frame

Less printing, far better cooling, easier access to the SD card and USB. Trades looks for
serviceability.

### B1. Ender 5 Plus Electronics Enclosure and Display for Raspberry Pi and Octopus Pro

<https://www.printables.com/model/781678-ender-5-plus-electronics-enclosure-and-display-for>

By Alancm81. Purpose-built for exactly your class of hardware: BTT Octopus Pro mainboard, a Pi, a
BTT display, and the original Creality bits, mounted to the frame with M4 T-nuts. This is the
closest match to your parts list of anything found.

- **Octopus Pro vs Octopus V1.1:** the two share the 160 × 100 mm outline and the 150 × 90 mm hole
  pattern, so the mounting plate should carry over. Connector positions differ slightly — confirm
  the cutouts clear your V1.1's terminal blocks.
- **BTT Pi vs Raspberry Pi 3B:** same 85 × 56 mm outline, and the listing's Pi is a 3B. Expect the
  fourth standoff not to line up (BTT Pi has only 3 of 4 RPi holes aligned). Three M2.5 standoffs
  hold it fine.

### B2. Ender 5 Plus Electronics Box Stand

<https://www.thingiverse.com/thing:4874853> · mirror: <https://cults3d.com/en/3d-model/tool/ender-5-plus-electronics-box-stand>

By spollock28269. The designer moved the control box out of the enclosure specifically to fix heat
issues; the stand sits it beside the printer with wire clearance, USB access, and room for a Pi box
underneath. Print twice, one rotated solid-side-out. Includes an over-the-top tool caddy, because
loose items on the E5+ frame rattle.

Good "get it running today" option, and the heat rationale applies double to a Mercury build.

### B3. Ender 5 Plus BTT Octopus + MOSFET Mounting Adapter

<https://www.printables.com/model/171559-ender-5-plus-btt-octopus-mosfet-mounting-adapter>

By NoWarrenty. Not an enclosure — an adapter that drops an Octopus into the **stock Creality control
box** using the existing screw holes. Cheapest path if you want to reuse what's already on the
printer. The stock bay is tight for a 160 × 100 mm board plus a Pi, and it's the worst option
thermally, but it's a few hours of printing.

### B4. Ender 5 Plus Modular Electronics Box ("The Pizza Box")

<https://www.thingiverse.com/thing:4165805>

By SteinerSE. Two-part modular case that sits on the E5+ frame, sized to take different mainboards
plus a Pi and MOSFETs, with cooling slots. Older design; the modularity is the appeal.

---

## Approach C — DIN rail bay inside whatever shell you pick

**This is what I'd actually do**, combined with A1.

Print or buy a shell for the looks, then mount a length of 35 mm DIN rail inside it and hang
everything off Voron-style printed clips. You get to re-lay out the bay later without reprinting a
single panel, which matters on a conversion where the wiring plan changes twice.

Mounts that exist for your exact boards:

| Part | Model |
|---|---|
| Octopus, DIN rail + fan shroud | <https://www.printables.com/model/137664-btt-octopus-din-rail-mount-with-fan-option> |
| Octopus, metal DIN clips + 90/120 mm shroud | <https://www.printables.com/model/146363-btt-octopus-mount-for-metal-din-rail-clips-90-and-> |
| Octopus, plain DIN bracket (horizontal or vertical) | <https://www.printables.com/model/547490-din-rail-btt-octopus-mount> |
| **BTT Pi V1.2**, Voron DIN mount | <https://www.printables.com/model/750071-voron-din-rail-mount-for-btt-pi-v12> |
| BTT Pi V1.2, alternate Voron bracket | <https://www.printables.com/model/616538-voron-btt-pi-bracket-for-din-rail-by-tom> |
| Voron user-mod mirrors | <https://www.teamfdm.com/files/file/759-voron-btt-pi-v12-din-rail-mount/> · <https://www.teamfdm.com/files/file/301-octopus-din-rail-bracket/> |

Note the BTT-Pi-specific mounts exist because the board *doesn't* fit generic Pi hardware — use one
of those rather than a Raspberry Pi bracket.

---

## Screen: TFT35-E3 V3.0.1

Mount it on the frame, not on the enclosure — you want it at eye level and away from the bay's heat.

- Ender 5 Plus external mount for the TFT35 V3:
  <https://www.printables.com/model/750018-ender-5-plus-btt-tft35-v3-external-screen-mount>
- TFT35-E3 case for the Ender 5 Plus:
  <https://www.printables.com/model/242410-tft35-e3-case-for-ender-5-plus>

**Firmware caveat worth knowing before you plan the wiring:** under Klipper the TFT35-E3 works in
**12864 emulation mode** out of the box (wired to the Octopus's EXP1/EXP2 headers). Its native touch
mode is a Marlin feature — running touch mode under Klipper needs a host-side bridge add-on that
forwards the display's gcode to Klipper, and users report partial functionality (progress/ETA
misbehave). Plan on EXP1/EXP2 + KlipperScreen or Mainsail on a tablet if you want the good UI.

- Bridge write-up: <https://oldhui.wordpress.com/2024/02/02/using-btt-tft35-in-touch-screen-mode-with-a-klipper-add-on/>
- Klipper forum thread on TFT35-E3 V3.0.1 config: <https://klipper.discourse.group/t/btt-tft35-e3-v3-0-1-help-with-printer-cfg/20638/4>

---

## Recommended build

1. **Shell:** A1 — the Big Dog skirting/lower enclosure (70 mm lift). It's the "base that holds up
   the frame" concept you were after, it's E5+-specific, and it's recently maintained.
2. **Interior:** a 35 mm DIN rail across the bay, Octopus on the fan-shroud clip, BTT Pi on the
   Voron BTT-Pi clip, PSU left in its stock frame position for now (moving 40.5 mm of PSU into a
   70 mm bay eats most of your air gap).
3. **Cooling:** one 60 mm intake at the front skirt behind a printed grille, exhaust vents at the
   opposite rear corner so air crosses the Octopus drivers. Filter the intake — this bay sits at
   floor level.
4. **Screen:** TFT35-E3 on the frame-mount above, wired EXP1/EXP2 in 12864 mode.
5. **Fallback if the bay gets tight:** B2's side stand. Zero risk, better cooling, uglier.

### Hardware to order alongside

- 35 mm DIN rail, ~300–400 mm, plus rail end stops
- M3 brass standoffs/screws for the Octopus (150 × 90 pattern), M2.5 for the BTT Pi
- M5 T-nuts + button-head screws for frame mounting — **verify your slot width first**; sources
  disagree on the E5+ profile mix (2020 on X, 2040 on Y/Z, heavier profile in the base), so measure
  rather than trusting a parts list
- 60 mm 24 V fan + printed grille/filter
- Stepper extension cables long enough to reach the Mercury top gantry, or a CAN toolhead kit
- Ferrules for every screw terminal on the Octopus, and a proper strain-relief point where the
  loom leaves the bay

### Safety

If mains-side wiring (PSU input, IEC inlet, switch) ends up inside a printed enclosure: use a fused
IEC inlet, keep mains physically separated from the low-voltage side, cover every mains terminal
with a printed or heat-shrink shroud, and don't print the mains-adjacent parts in PLA — the bay
under a Mercury build runs warm. PETG or ABS/ASA for anything touching or enclosing mains.

---

## Sources

- [Ender 5 Plus Skirting and Lower Enclosure (Base Case) — Printables](https://www.printables.com/model/557066-ender-5-plus-skirting-and-lower-enclosure-base-cas)
- [Ender 5 Plus Electronics Enclosure and Display for Raspberry Pi and Octopus Pro — Printables](https://www.printables.com/model/781678-ender-5-plus-electronics-enclosure-and-display-for)
- [Zero G Mercury One.1 for Ender 5 Pro Bottom Panel Electronics Enclosure — Printables](https://www.printables.com/model/896979-zero-g-mercury-one1-for-ender-5-pro-bottom-panel-e)
- [Ender 5 Electronics Skirt — Printables](https://www.printables.com/model/1518720-ender-5-electronics-skirt)
- [Ender 5 Plus Mercury One.1 Enclosure Top-Hat — Printables](https://www.printables.com/model/1099998-ender-5-plus-mercury-one1-enclosure-top-hat)
- [Ender 5 Plus Electronics Box Stand — Thingiverse](https://www.thingiverse.com/thing:4874853)
- [Ender 5 Plus modular electronics box "The Pizza Box" — Thingiverse](https://www.thingiverse.com/thing:4165805)
- [Ender 5 Plus BTT Octopus + MOSFET Mounting Adapter — Printables](https://www.printables.com/model/171559-ender-5-plus-btt-octopus-mosfet-mounting-adapter)
- [BTT Octopus DIN rail mount with fan option — Printables](https://www.printables.com/model/137664-btt-octopus-din-rail-mount-with-fan-option)
- [Voron DIN rail mount for BTT Pi v1.2 — Printables](https://www.printables.com/model/750071-voron-din-rail-mount-for-btt-pi-v12)
- [Ender 5 Plus BTT TFT35 V3 External screen mount — Printables](https://www.printables.com/model/750018-ender-5-plus-btt-tft35-v3-external-screen-mount)
- [TFT35-E3 case for Ender 5 Plus — Printables](https://www.printables.com/model/242410-tft35-e3-case-for-ender-5-plus)
- [Mercury One.1 — ZeroG Documentation](https://docs.zerog.one/manual/build/mercury_eva)
- [Ender-5-Plus Octopus V1.1 + Klipper install notes — GitHub](https://github.com/juldaani/Ender-5-plus)
- [Ender 5 Plus Mercury One.1 Klipper/Mainsail config — GitHub](https://github.com/ethomasgt/Ender-5-Plus-Mercury-One-Klipper-Mainsail)
- [Zero G Mercury One.1 Build thread — V1E Forum](https://forum.v1e.com/t/zero-g-mercury-one-1-build/42041)
- [BIGTREETECH Pi v1.2 specifications — CNX Software](https://www.cnx-software.com/2023/05/01/bigtreetech-pi-v1-2-a-raspberry-pi-sized-allwinner-h616-sbc-for-3d-printers/)
- [Octopus — BIGTREETECH Wiki](https://global.bttwiki.com/Octopus.html)
- [TFT35 E3 — BIGTREETECH Wiki](https://global.bttwiki.com/TFT35%20E3.html)
- [Using BTT TFT35 in Touch Screen Mode with a Klipper Add-on](https://oldhui.wordpress.com/2024/02/02/using-btt-tft35-in-touch-screen-mode-with-a-klipper-add-on/)
