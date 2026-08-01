# Open-Source Camera Drone — Complete Build Guide

**Version:** 1.3
**Date:** August 2025
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

> ✅ **Printer note:** With a 600×600×600mm build volume, a 5" frame's full footprint (arms, ducts, and center plate together) fits easily on a single build plate. This build uses a **unibody frame print** rather than separate arm/duct sections joined with standoffs — fewer parts, no joint hardware, and no need for outsourced printing.

1. **Download a unibody STL** from Printables or Thingiverse (search "5 inch unibody enclosed cage drone frame") — or remix a split-arm design into one plate in your slicer/CAD tool, since your bed size supports it. Save/print the following STL file: [Sample 5" unibody enclosed cage frame — Printables](https://printables.com).
2. **Print settings:** PETG or Nylon-CF, 3–4 perimeter walls, 25–40% infill for the duct/arm structure, 100% infill for motor mount bosses. A single-piece print typically takes similar total time (8–12 hours) to a multi-part print but runs as one continuous job — start it early so it's not the critical path.
3. **Bed adhesion for large single-piece prints:** Use a brim or raft given the part's larger footprint, and make sure your bed is leveled across its full working area — large flat sections are more prone to warping at the edges than small split parts were.
4. **Post-processing:** Remove supports, clean up motor mount holes, and lightly sand any layer lines. Since arms and center plate are one piece, there's no need to test-fit joints — just verify mounting hole alignment against your FC/motor hardware.
5. **Assemble the cage:** With the frame already unibody, this step is now just attaching the top plate (if separate) and any add-on brackets (camera mount, GPS mast). Leave the top plate off until electronics are installed.
6. **Mount motors:** Thread motors onto each integrated arm mount using the included screws + threadlocker. Confirm motor rotation direction against your ESC/firmware motor map before final tightening.

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
Swap the 6S LiPo pack for a 6S Li-ion pack built from high-drain cells (e.g., Molicel P42A/P45B 21700). Li-ion has roughly double the energy density of LiPo at the discharge rates this build actually needs (cruise flight, not racing), so a similarly-sized/weighted pack can extend flight time from ~6–8 minutes to ~12–18 minutes.
- **Tradeoffs:** lower max discharge current (a non-issue for a cruising camera platform, but avoid punchy throttle stabs), slower charge times, and higher upfront cost (~$60–100/pack vs. $30–40 for LiPo).
- **Failsafe reconfiguration required:** Li-ion's voltage curve is flatter across most of the discharge range but drops sharply near cutoff, unlike LiPo's more gradual sag. Update the flight controller's low-battery voltage failsafe thresholds (Mission Planner/QGC battery monitor settings) for Li-ion before its first flight — reusing LiPo thresholds risks either a false-early RTL or a battery over-discharged before failsafe triggers.
- Confirm your charger supports a Li-ion charge profile, not just LiPo — most modern 6S balance chargers do, but check before assuming.

### 2. Motor & Propeller: Lower KV, Larger/Higher-Pitch Prop
Replace the 1700KV motors with a lower-KV option (roughly 1400–1500KV) and pair with a slightly larger or higher-pitch propeller. This trades peak thrust and agility for efficiency — hover current draw drops noticeably, which extends flight time at effectively no added weight and no added cost versus the baseline motor spec.
- **Tradeoffs:** reduced punch-out/acceleration performance — not a concern for a cruising camera platform, but noticeable if you ever want aggressive maneuvers.
- Confirm the new propeller still clears the frame's duct/cage geometry before ordering (§3) — a larger prop may not fit the baseline duct clearance.

### 3. Payload: Skip the Gimbal, Use the Fixed Camera Mount
Mount the action camera (GoPro Session / DJI Osmo Action) directly on the rigid 3D-printed camera bracket (already included in the frame print, §3, and in the parts list at purchase list §6) instead of the Holybro A8 Mini 3-axis gimbal. This removes the gimbal's ~130–180g and its full cost from the build.

> **Note on the mount:** with the gimbal skipped, footage stabilization falls back entirely to the action camera's built-in electronic stabilization (GoPro HyperSmooth / DJI RockSteady) rather than mechanical gimbal stabilization. This is noticeably softer for smooth panning/tilting shots but is fine for straight cruise footage — a reasonable tradeoff if flight time and build simplicity matter more than gimbal-smooth video. The fixed mount bolts to the same front bracket point as the gimbal would, so this is a parts-list decision, not a frame redesign.

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
| Thingiverse (STL files) | [thingiverse.com](https://thingiverse.com) |
| HDZero (digital FPV) | [hd-zero.com](https://www.hd-zero.com) |

---

**Document Version:** 1.3
**Last Updated:** August 2025
**Companion Documents:** `drone_purchase_list.md`, `drone_requirements_summary.md`
