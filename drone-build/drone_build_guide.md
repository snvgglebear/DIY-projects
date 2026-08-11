# Open-Source Camera Drone — Complete Build Guide

**Version:** 1.7
**Date:** August 2026
**Build Time:** 8–12 hours (excluding 3D print time)
**Skill Level:** Intermediate (basic soldering and assembly required)

---

## Table of Contents

1. [Overview](#overview)
2. [Prerequisites](#prerequisites)
3. [3D Printed Frame Assembly](#3d-printed-frame-assembly)
4. [Electronics Installation](#electronics-installation)
5. [Dual-Band 915MHz Configuration](#dual-band-915mhz-configuration)
6. [Smartphone Ground Station Setup](#smartphone-ground-station-setup)
7. [Flight Controller Configuration](#flight-controller-configuration)
8. [Testing & Calibration](#testing--calibration)
9. [First Flight Checklist](#first-flight-checklist)
10. [FPV Capability (Optional Add-On)](#fpv-capability-optional-add-on)
11. [Flight Time Optimization (Optional)](#flight-time-optimization-optional)
12. [Troubleshooting](#troubleshooting)

---

## Overview

This guide walks through building a fully open-source, 5" camera drone with an enclosed 3D printed propeller-guard frame, dual-band 915MHz communications (ExpressLRS for control, LoRa for backup telemetry), and a smartphone + Xbox controller ground station running QGroundControl. See the companion `drone_purchase_list.md` for the full parts list and `drone_requirements_summary.md` for design rationale.

**Core platform:**
- Flight controller: Holybro Pixhawk 6C/6X
- Firmware: ArduPilot or PX4
- Frame: 5" enclosed cage, PETG or Nylon-CF
- Motors/ESC: T-Motor Velox 2207 1700KV + 55A BLHeli_32 (lower-KV variant available for extended flight time — see §11)
- Battery: 6S 1300–1500mAh LiPo (Li-ion swap available for extended flight time — see §11)
- Camera: Action cam (GoPro Session / DJI Osmo Action), optional Holybro A8 Mini gimbal (skippable in favor of a lighter fixed mount — see §11)
- Optional: FPV camera + VTX + goggles for a live first-person video feed (see §10)

---

## Prerequisites

**Skills needed:**
- Basic soldering (ESC-to-motor and power wiring)
- Comfort with small hardware (M2/M3 screws, standoffs)
- Basic command-line/app-based firmware flashing

**Tools required** (see purchase list §7):
- Soldering iron + solder
- Hex driver set
- Multimeter
- Heat shrink tubing
- Threadlocker

**Before you start:**
- Download and slice the frame STL files (see §3)
- Install QGroundControl on your Android phone
- Install the ExpressLRS Configurator on a laptop
- Charge at least one 6S LiPo battery
- If adding FPV, decide analog vs. digital HD before ordering hardware (see §10) — it changes the camera, VTX, and goggle purchases together as a matched set
- If optimizing for flight time, decide on battery chemistry, motor/prop, and gimbal-vs-fixed-mount choices before ordering (see §11) — these affect the parts list, not just assembly

---

## 3D Printed Frame Assembly

> ✅ **Printer note:** With a 600×600×600mm build volume, a 5" frame's full part set (arms, ducts, center plate, and any add-on brackets) fits easily on a single build plate, so everything prints in one job regardless of which design you pick. Earlier versions of this guide described the frame as a single-piece "unibody" print — that didn't hold up against the real caged/ducted 5" frame designs verified in purchase list §2. All three options there are multi-part (arms, center plate, and duct sections that bolt together), not one continuous piece; what the large bed buys you is printing all of those parts together in one job instead of splitting the run or outsourcing.

1. **Pick and download a frame design** from purchase list §2 — Option A or B (both replicas of the Lumenier QAV-PRO Whoop 5" caged cinewhoop frame, modeled independently by two different designers) or Option C (a distinct, free 3-part ducted design). All three fit comfortably within your printer's bed with room to spare; lay out every part of whichever design you pick on one plate to print it in a single job.
   - **Fewer, larger parts?** A search specifically for a real, enclosed/ducted 5" cage design purpose-built with fewer, larger parts didn't turn up a match — the closest hit, ARS-5, is a well-known "few parts" printable 5" frame, but it's an open racing frame without propeller guards, so it doesn't meet this build's enclosed-cage spec. Option C is already the leanest of the three verified options at just 3 parts. If you want to go further, `drone_frame_merge_guide.md` walks through merging Option A's pieces yourself in Onshape (Option A is used there since it includes a STEP file — a much more reliable starting point for a boolean merge than Option C's STL-only files).
2. **Print settings:** PETG or Nylon-CF, 3–4 perimeter walls, 25–40% infill for the duct/arm structure, 100% infill for motor mount bosses. Printing every part of the frame together in one job takes roughly the same total time (8–12 hours) as printing them across separate jobs, but runs as one continuous job — start it early so it's not the critical path.
3. **Bed adhesion for large multi-part layouts:** Use a brim or raft given how much of the bed is covered, and make sure your bed is leveled across its full working area — large flat sections and tightly-packed parts are more prone to edge warping than a single small part would be.
4. **Post-processing:** Remove supports, clean up motor mount holes, and lightly sand any layer lines. Test-fit the arm-to-center-plate and duct-section joints before final assembly — these are multi-part designs, so joint fit matters here, unlike it would for a true single-piece print.
5. **Assemble the cage:** Bolt the arms, center plate, and duct sections together using your chosen design's joint hardware (see purchase list §2), then attach the top plate (if separate) and any add-on brackets (camera mount, GPS mast). Leave the top plate off until electronics are installed.
   - **Retaining the gimbal mount:** Options A and B are replicas of a frame with a modular cinema camera mount confirmed compatible with gimbals — good odds it also fits the Holybro A8 Mini (55×55×70mm, four M2.5×8mm mounting screws), but test-fit before committing since replica prints can vary slightly from the original. Option C has no confirmed gimbal provision in its design. If your chosen frame doesn't have a compatible mount point and you want to keep the gimbal rather than go to a fixed mount (§11), print a simple custom adapter plate — drilled for the A8 Mini's 4× M2.5 mounting holes on one side and your frame's existing standoff pattern on the other — instead of losing gimbal capability.
6. **Mount motors:** Thread motors onto each arm mount using the included screws + threadlocker. Confirm motor rotation direction against your ESC/firmware motor map before final tightening.

---

## Electronics Installation

1. **Mount the flight controller** on vibration-dampening standoffs at the center of the frame, with the arrow on the FC case pointing toward the front of the drone.
2. **Install the 4-in-1 ESC** beneath or above the FC (per your stack configuration) and connect motor leads (swap any two leads later if rotation direction is wrong — don't guess in advance).
3. **Wire power:** Solder the XT60 battery lead to the power distribution board / ESC input. Add an inline fuse or power module per Pixhawk 6C wiring diagram.
4. **Mount the GPS module** on a small mast or the top plate, away from ESC/motor wiring to minimize compass interference.
5. **Install the camera mount bracket** on the front of the frame (or gimbal, if used) and route a clean power tap for the action camera (5V rail or dedicated small battery).
6. **Cable management:** Zip-tie and route all wiring away from propeller arcs; use foam tape to secure the FC and any loose modules.

---

## Dual-Band 915MHz Configuration

This build uses two independent 915MHz links: **ExpressLRS (ELRS)** for low-latency RC control, and **LoRa (MicoAir LR900-F)** for long-range backup telemetry.

### ELRS Setup
1. Flash both the transmitter module and receiver with matching firmware using the [ExpressLRS Configurator](https://www.expresslrs.org).
2. Bind the receiver to the transmitter (hold bind button or use bind phrase method).
3. Connect the ELRS receiver to the Pixhawk's RC input (SBUS/CRSF port depending on receiver type).
4. In your firmware (ArduPilot/PX4), configure the RC protocol to CRSF and verify channel mapping.

### LoRa Telemetry Setup
1. Connect the air-side MicoAir LR900-F unit to the Pixhawk's telemetry (TELEM2) port.
2. Connect the ground-side LR900-F unit to your Android phone via the USB-C OTG adapter and LoRa USB dongle.
3. Confirm both units are set to the same baud rate and frequency channel (check MicoAir documentation for regional 915MHz channel plan).
4. Verify telemetry link in QGroundControl once paired (see §6).

> ⚠️ Since both ELRS and LoRa share the 915MHz band, use different channel/frequency-hopping settings for each to avoid interference — refer to the ExpressLRS and MicoAir documentation for coexistence guidance.

> 📷 The LR900-F's telemetry link has far more range than any FPV video link (§10) will ever reach. If you want a rough visual check-in beyond FPV range rather than just numeric telemetry, see `drone_lora_still_image_uplink.md` — an experimental, DIY-only approach for sending periodic low-res stills over this same link.

---

## Smartphone Ground Station Setup

1. Install **QGroundControl** on your Android phone: [qgroundcontrol.com](https://qgroundcontrol.com).
2. Connect the ground-side LoRa dongle to the phone via USB-C OTG adapter.
3. Pair the **Xbox Wireless Controller** to the phone via Bluetooth, or connect via USB-C OTG if using a wired controller.
4. Open QGroundControl → confirm the telemetry link connects automatically (or add it manually under **Application Settings → Comm Links**).
5. Configure joystick mapping in QGroundControl (**Application Settings → General → Virtual Joystick / Joystick**) to map Xbox controller sticks to roll/pitch/throttle/yaw.
6. Verify live telemetry (battery voltage, GPS lock, altitude) appears on the QGC dashboard.

---

## Flight Controller Configuration

1. Connect the Pixhawk to a laptop via USB and open **Mission Planner** (ArduPilot) or **QGroundControl** (PX4) for initial setup.
2. Select and flash your chosen firmware (ArduPilot or PX4) — see [ardupilot.org](https://ardupilot.org) or [docs.px4.io](https://docs.px4.io).
3. Run the **frame setup wizard** and select a 5" quad/X configuration matching your build.
4. Complete **accelerometer, compass, and radio calibration** wizards in order.
5. Set up **failsafe behavior** (RTL or land) for RC signal loss and low battery — critical if you're not running FPV, since there's no live video feed as a backup situational cue. If you're using a Li-ion pack (§11), set the low-battery voltage thresholds for Li-ion's discharge curve, not LiPo's — see §11 for details.
6. Configure **motor output test** to confirm correct spin direction for each motor before ever installing propellers.

---

## Testing & Calibration

1. **Propeller-off motor test:** Arm the drone (props removed) and verify each motor spins in the correct direction and responds to stick input.
2. **ESC calibration:** Follow your ESC's calibration procedure (throttle range endpoints) if motors are unresponsive or uneven.
3. **Compass/GPS check:** Take the drone outdoors, wait for GPS lock (green/solid LED per Pixhawk status), and confirm heading accuracy by comparing to a known direction.
4. **Range test:** With props off, walk the ELRS transmitter away from the aircraft to confirm link quality (RSSI) remains strong at your intended operating range.
5. **Telemetry range test:** Similarly verify the LoRa telemetry link holds at extended range using QGroundControl's signal indicators.
6. **Camera test:** Power on the action camera via the drone's rail and confirm recording function and mount stability (no vibration blur).

---

## First Flight Checklist

- [ ] Propellers installed and tightened, correct rotation direction confirmed
- [ ] Battery fully charged and securely mounted
- [ ] GPS lock acquired (solid lock indicator)
- [ ] ELRS link bound and showing good RSSI
- [ ] LoRa telemetry connected in QGroundControl
- [ ] Failsafe settings verified (RTL/land on signal loss or low battery; voltage thresholds match your battery chemistry — see §11 if using Li-ion)
- [ ] Camera recording and mounted securely
- [ ] FPV video link confirmed clean and goggles paired, if FPV installed (see §10)
- [ ] FPV camera angle checked for prop-guard obstruction, if FPV installed
- [ ] Open outdoor area, clear of people/obstacles, legal to fly (see regulatory notes in requirements summary)
- [ ] First flight in manual/stabilize mode only — confirm control response before enabling autonomous modes

---

## FPV Capability (Optional Add-On)

This build is designed around a smartphone ground station with QGroundControl telemetry — the action camera records locally rather than streaming live. Adding FPV (first-person view) layers in a real-time video feed from the aircraft's perspective, which is useful for threading tighter spaces, judging distance/obstacles more precisely, or simply flying "through the lens" instead of watching from the ground. It's an add-on video path, independent of the existing action-cam recording and QGC telemetry link.

### Choosing a System

| Type | Latency | Image Quality | Cost | Notes |
|---|---|---|---|---|
| Analog (5.8GHz) | Lowest (~ms-level) | Low-res, degrades with noise at range | $ | Cheapest, simplest to diagnose in the field, most rugged |
| Digital HD (e.g., DJI O4, HDZero, Walksnail Avatar) | Low (roughly 20–40ms) | HD, stays clean at range | $$$ | Best image quality; camera + VTX usually bundled as one "air unit" |

For a camera-focused build like this one, digital HD is the more natural fit since it matches the "good picture" goal of the action-cam payload. Analog remains the budget-friendly, easy-to-repair option.

### Hardware Needed (see purchase list §8)
- Dedicated FPV camera (separate from the action cam — FPV cameras are tuned for latency and dynamic range, not recording quality)
- Video transmitter (VTX): 5.8GHz analog or a digital HD system
- FPV goggles (analog goggles, or digital goggles paired with the matching HD system)
- 5.8GHz antenna (circular polarized recommended over linear — more resistant to multipath interference at range)

### Installation Steps
1. Mount the FPV camera at the front of the frame, angled slightly down (roughly 20–30°), with a clear line of sight past the prop guards — test-fit first, since an enclosed cage frame can clip into the camera's field of view.
2. Mount the VTX away from the GPS module and ELRS/LoRa antennas to reduce RF interference; keep its antenna vertical and clear of carbon/metal frame parts.
3. Wire VTX power from the same battery/PDB rail as the flight electronics — confirm the VTX's input voltage range first (most accept 2S–6S directly, but don't assume).
4. Connect camera video-out to VTX video-in; set the VTX's output power and channel/band according to local regulations (5.8GHz FPV power limits are generally higher in the US than in the EU — check your local rules before transmitting).
5. Bind the goggles to the VTX (analog: tune to a matching channel; digital: pair via the system's bind procedure, similar in spirit to the ELRS binding in §5).
6. No coexistence tuning is needed between FPV video and the existing ELRS/LoRa links — 5.8GHz video and 915MHz control/telemetry sit on separate bands.

### Testing
- Bench-test the video feed before the first flight: power the drone with props off, put on the goggles, and confirm a clean, correctly oriented, in-focus image.
- Range-test the video link the same way you range-test ELRS (§8): walk away from the aircraft with props off and watch for signal degradation or dropouts.
- Re-check propeller balance and vibration (§8) with the FPV camera mounted — it's added mass on the frame and can introduce its own vibration if not secured tightly.

---

## Flight Time Optimization (Optional)

The baseline spec (6S 1300–1500mAh LiPo, 1700KV motors, gimbal-mounted camera) prioritizes agility and image stabilization over endurance. If flight time matters more than punch-out performance for your use case, the following changes are worth considering — independently or together. See purchase list §9 for the associated parts.

### 1. Battery: Li-ion Instead of High-C LiPo
Swap the 6S1P LiPo pack for a 6S1P Li-ion pack built from high-drain 21700 cells (Molicel P45B preferred, P42A acceptable — see purchase list §5b for cell specs, pack geometry, and dimensions). Li-ion has roughly double the energy density of LiPo at the discharge rates this build actually needs (cruise flight, not racing) — a P45B-based pack runs ~97 Wh at ~450–470g versus ~29 Wh at ~220–230g for the baseline 1300mAh LiPo — so flight time can extend from ~6–8 minutes to ~12–18 minutes.
- **Tradeoffs:** lower max discharge current, slower charge times, and higher upfront cost (~$60–100/pack vs. $30–40 for LiPo).
- **Current limit is not optional, it's required.** A 6S1P pack has one cell's rating as its ceiling — 45A for the P45B/P42A, for the whole aircraft — with no parallel cells to share load. Hover draw (~10–15A) is well within that, but a hard throttle punch on four 2207 motors at 6S can exceed 45A, which is also above the PM02 V3/PM06 V2 lead's 60A burst rating (purchase list §1). Set a firmware current limit (ArduPilot: `MOT_BAT_CURR_MAX`, tuned to something below the pack's continuous rating with margin) before the first Li-ion flight, and fly it as a cruiser rather than a puncher.
- **Failsafe reconfiguration required:** Li-ion's voltage curve is flatter across most of the discharge range but drops sharply near cutoff, unlike LiPo's more gradual sag. Update the flight controller's low-battery voltage failsafe thresholds (`BATT_LOW_VOLT`, `BATT_CRT_VOLT` in ArduPilot, or the equivalent Mission Planner/QGC battery monitor fields) for Li-ion before its first flight — reusing LiPo thresholds risks either a false-early RTL or a battery over-discharged before failsafe triggers. Also set `BATT_FS_VOLTSRC` to use sag-corrected voltage rather than raw voltage, so a hard throttle input doesn't trip a false failsafe on Li-ion's steeper sag curve.
- Confirm your charger supports a Li-ion charge profile, not just LiPo — most modern 6S balance chargers do, but check before assuming.
- **Mounting:** the Li-ion pack's geometry doesn't match the LiPo bay — see purchase list §5b for why it's mounted externally on the top plate rather than in the printed bay, and the CG check that comes with it.

### 2. Motor & Propeller: Lower KV, Bi-Blade Prop, Pitch Matched to KV
Replace the 1700KV motors with a lower-KV option (roughly 1400–1500KV) and swap the baseline 5040 **tri-blade** propellers for a 5" **bi-blade**. This trades peak thrust and agility for efficiency — hover current draw drops noticeably, which extends flight time at effectively no added weight and no added cost versus the baseline motor spec.

**The metric that matters is grams of thrust per watt (g/W) at hover, not peak thrust.** Manufacturer thrust tables usually headline max thrust; the number to read is the one at roughly a quarter of max, which is where a cruising camera platform actually spends the entire flight. Three physical properties drive that figure, in descending order of impact:

1. **Disc area (diameter) — dominant, but fixed on this build.** Induced power scales roughly as thrust^1.5 / √(disc area), so more disc area for the same thrust is the single largest efficiency win available to any multirotor. It is *not* available here: all three frame options in §3 are 5" ducted cages, and the duct sets tip clearance. Going larger means rescaling the frame, not buying a different prop.
2. **Blade count — the biggest lever you actually have.** Two blades beat three at the same diameter and pitch, typically by around 10–15% in hover g/W, because each blade isn't flying through the wake of the one ahead of it. Since diameter is unchanged, a tri-blade → bi-blade swap needs no duct clearance re-check, costs nothing extra, and is the cheapest endurance gain in the build.
3. **Pitch — a motor-loading match, not an efficiency lever.** At a fixed diameter, lower pitch is generally the *more* efficient hover prop; high pitch buys top speed and punch-out, which is exactly what this optimization is trading away. What justifies adding pitch here is reloading the slower motor, not efficiency in itself — see below.

**Why the KV change forces a pitch decision:** at 6S, dropping 1700KV → 1400KV takes unloaded RPM from roughly 37,700 to 31,100, about −17%. Static thrust scales with RPM², so on an unchanged propeller you would lose roughly 30% of peak thrust and hover throttle would climb accordingly. Pitch is how you recover that thrust — but adding pitch spends part of the efficiency you just bought. Start at the low end and add pitch only if the hover throttle check below comes out high.

**What to buy:** 5" diameter (fixed by the duct), **2 blades**, pitch roughly 3.5–4.5" (pitch-to-diameter ratio ≈ 0.7–0.9), light stiff polycarbonate. Skip heavy reinforced racing props — the added rotational mass and blade thickness both work against you here. HQProp's 5x4.3x2 and 5x3.5x2 are reasonable starting candidates in this class, but treat any specific model as a candidate to verify against its own thrust data at the hover point rather than a settled answer.

**Verify with the hover throttle check.** After a hover flight, read ArduPilot's `MOT_THST_HOVER`:
- **0.35–0.45** — the target band. Good efficiency with enough authority left for wind and attitude correction.
- **Below ~0.25** — overpropped. You're carrying thrust you never use; that mass would have been better spent on battery.
- **Above ~0.5** — too little margin. Reduce weight or add pitch.

Aim for roughly 2–2.5:1 total thrust-to-weight for an endurance cruiser, versus the 5:1 or more a racing build wants.

**Then measure instead of guessing.** The power module (§1) gives you calibrated current sensing, so you can A/B propellers empirically without a thrust stand: hover each candidate for 60 seconds at identical takeoff weight and compare mAh consumed from the battery monitor log. That single test accounts for your duct, your actual all-up weight, and your ESC — all things a spec sheet cannot. Buy two or three candidate props and fly the comparison.

- **Tradeoffs:** reduced punch-out/acceleration performance — not a concern for a cruising camera platform, but noticeable if you ever want aggressive maneuvers.
- Any propeller change still needs a duct/cage clearance check against §3 before ordering. A bi-blade swap at the same diameter is clearance-neutral; anything that changes diameter is not.
- **Expectation setting on the cage:** a shrouded rotor *can* outperform an open one, but only with tight tip clearance (roughly 1–2% of radius) and a proper diffuser profile. Printed cinewhoop ducts generally have looser clearance than that and add mass, so an enclosed 5" is usually somewhat *less* efficient than an open-prop quad on the same propellers. Propeller choice here is worth on the order of 10–15%; the battery change in #1 and the payload decision in #3 are the larger numbers. The enclosed cage is a stated requirement of this build (see `drone_requirements_summary.md`), so this is a constraint to plan around, not a reason to change course — just don't expect propellers alone to move flight time much.

### 3. Payload: Skip the Gimbal, Use the Fixed Camera Mount — or Retain It
Mount the action camera (GoPro Session / DJI Osmo Action) directly on the rigid 3D-printed camera bracket (already included in the frame print, §3, and in the parts list at purchase list §6) instead of the Holybro A8 Mini 3-axis gimbal. This removes the gimbal's ~130–180g and its full cost from the build.

> **Note on the mount:** with the gimbal skipped, footage stabilization falls back entirely to the action camera's built-in electronic stabilization (GoPro HyperSmooth / DJI RockSteady) rather than mechanical gimbal stabilization. This is noticeably softer for smooth panning/tilting shots but is fine for straight cruise footage — a reasonable tradeoff if flight time and build simplicity matter more than gimbal-smooth video.

**If you'd rather keep the gimbal, that's still the default** — the Holybro A8 Mini stays in the parts list (purchase list §6) with nothing else to change. Gimbal-mount compatibility depends on which frame option you picked in §3: Options A and B (Lumenier QAV-PRO Whoop 5" replicas) have a modular cinema mount with good odds of fitting the A8 Mini directly — test-fit first. Option C has no confirmed gimbal provision, so retaining the gimbal on that frame means printing the custom adapter plate described in §3 step 5. Either way, this is a parts-list/mount decision, not a frame redesign.

Combined, these three changes stack: a lighter payload, a more efficient motor/prop draw, and a higher energy-density battery. None require frame changes beyond the prop clearance check in #2.

---

## Troubleshooting

| Symptom | Likely Cause | Fix |
|---|---|---|
| Motor spins wrong direction | Incorrect ESC wiring or motor map | Swap any two of the three motor-to-ESC leads for that motor |
| No GPS lock | Poor sky view, compass interference | Move outdoors, away from metal; check GPS mast placement |
| ELRS won't bind | Mismatched firmware versions | Re-flash TX and RX with matching ExpressLRS firmware version |
| LoRa telemetry drops out | Frequency/baud mismatch, antenna orientation | Confirm matching channel/baud on both LR900-F units; check antenna orientation |
| QGroundControl doesn't detect telemetry | OTG adapter/dongle not recognized | Try a different USB-C OTG adapter; confirm phone supports USB host mode |
| Xbox controller not responding in QGC | Joystick not enabled in settings | Enable joystick input under Application Settings → General |
| Excessive vibration in camera footage | Loose frame standoffs or unbalanced props | Re-tighten all standoffs; balance propellers before flight |
| FPV image blank / no signal | VTX/goggle band-channel mismatch, VTX not powered | Confirm matching band and channel on both ends; check VTX power LED |
| FPV image noisy or dropping out at range | Antenna polarization mismatch, VTX power too low, RF interference | Match antenna polarization on both ends; increase VTX power within legal limits; try a different channel |
| Voltage sag / early failsafe trip under throttle (Li-ion pack) | Li-ion has lower max continuous discharge current than LiPo | Avoid aggressive throttle stabs; if sag persists, use a pack with more parallel cells for higher continuous amperage |
| Noticeably less punch/climb after motor-prop swap | Lower-KV motor + efficiency prop trades thrust for endurance (§11) | Expected behavior — adjust throttle curve/expo in FC settings if more stick throw is needed |

---

## Reference Links

| Resource | Link |
|---|---|
| ArduPilot Firmware | [ardupilot.org](https://ardupilot.org) |
| PX4 Firmware | [docs.px4.io](https://docs.px4.io) |
| QGroundControl | [qgroundcontrol.com](https://qgroundcontrol.com) |
| ExpressLRS Docs | [expresslrs.org](https://www.expresslrs.org) |
| ArduPilot Forum | [discuss.ardupilot.org](https://discuss.ardupilot.org) |
| PX4 Forum | [discuss.px4.io](https://discuss.px4.io) |
| Printables (STL files) | [printables.com](https://printables.com) |
| Cults3D (STL files) | [cults3d.com](https://cults3d.com) |
| Thingiverse (STL files) | [thingiverse.com](https://thingiverse.com) |
| HDZero (digital FPV) | [hd-zero.com](https://www.hd-zero.com) |

---

**Document Version:** 1.7
**Last Updated:** August 2026
**Companion Documents:** `drone_purchase_list.md`, `drone_frame_merge_guide.md`, `drone_kit_alternatives.md`, `drone_lora_still_image_uplink.md`, `drone_requirements_summary.md`
