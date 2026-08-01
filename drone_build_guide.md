# Open-Source Camera Drone — Complete Build Guide

**Version:** 1.1
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
10. [Troubleshooting](#troubleshooting)

---

## Overview

This guide walks through building a fully open-source, 5" camera drone with an enclosed 3D printed propeller-guard frame, dual-band 915MHz communications (ExpressLRS for control, LoRa for backup telemetry), and a smartphone + Xbox controller ground station running QGroundControl. See the companion `drone_purchase_list.md` for the full parts list and `drone_requirements_summary.md` for design rationale.

**Core platform:**
- Flight controller: Holybro Pixhawk 6C/6X
- Firmware: ArduPilot or PX4
- Frame: 5" enclosed cage, PETG or Nylon-CF
- Motors/ESC: T-Motor Velox 2207 1700KV + 55A BLHeli_32
- Battery: 6S 1300–1500mAh LiPo
- Camera: Action cam (GoPro Session / DJI Osmo Action), optional Holybro A8 Mini gimbal

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
5. Set up **failsafe behavior** (RTL or land) for RC signal loss and low battery — critical given this build has no live FPV feed.
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
- [ ] Failsafe settings verified (RTL/land on signal loss or low battery)
- [ ] Camera recording and mounted securely
- [ ] Open outdoor area, clear of people/obstacles, legal to fly (see regulatory notes in requirements summary)
- [ ] First flight in manual/stabilize mode only — confirm control response before enabling autonomous modes

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

---

**Document Version:** 1.1
**Last Updated:** August 2025
**Companion Documents:** `drone_purchase_list.md`, `drone_requirements_summary.md`
