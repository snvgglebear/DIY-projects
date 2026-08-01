# Open-Source Camera Drone — Purchase List

**Build Goal:** Compact camera-capable drone with 3D printed propeller-guard frame, dual-band 915MHz comms (ELRS + LoRa), smartphone/Xbox ground control station
**Estimated Total:** ~$900–1,250 (depending on smartphone availability)
**Last Updated:** August 2025
**Note:** Frame cost reduced — owner has a 600×600×600mm printer, so the full unibody cage frame is printed in-house (no outsourced print service needed).

---

## Quick Summary Table

| Category | Estimated Cost |
|---|---|
| Flight Controller & Core Electronics | $275–325 |
| Frame & Propeller Guards (3D Printed, unibody, home-printed) | $35–90 |
| Dual-Band 915MHz Comms System | $145–185 |
| Ground Control Station (Smartphone) | $0–165 |
| Battery & Charging | $75–100 |
| Camera Payload | $100–340 |
| Tools & Misc | $100–120 |
| **TOTAL** | **~$730–1,325** |

---

## 1. Flight Controller & Core Electronics

| Qty | Item | Model / Specs | Price | Where to Buy |
|---|---|---|---|---|
| 1 | Flight Controller Stack | Holybro Pixhawk 6C (with plastic case) | $220–299 | [Holybro Store](https://holybro.com/products/pixhawk-6c) • [ReadyMadeRC](https://www.readymaderc.com/products/details/86931-holybro-pixhawk-6c-with-plastic-case) • [PyroDrone](https://pyrodrone.com/products/holybro-pixhawk-6c-plastic-case) |
| 1 | GPS Module | Holybro M8N (budget) or M10 (better accuracy) | $25–45 | [Holybro Store](https://holybro.com/collections/gps) |
| 4 | Motors | T-Motor Velox 2207 1700KV | $22–28 each (~$90–110 set) | [T-Motor Store](https://www.tmotor.com) • [GetFPV](https://www.getfpv.com) |
| 1 | ESC (4-in-1) | 55A BLHeli_32 4-in-1 ESC | $55–75 | [GetFPV](https://www.getfpv.com) • [PyroDrone](https://pyrodrone.com) |
| 8 | Propellers | 5040 tri-blade, 6S-compatible (2 sets, incl. spares) | $8–12/set | [GetFPV](https://www.getfpv.com) • [RaceDayQuads](https://www.racedayquads.com) |
| 1 | Power Distribution Board | Matches ESC stack (often bundled) | $0–20 | Bundled with FC stack or [Holybro Store](https://holybro.com) |

**Subtotal: ~$275–325 (excluding spares)**

---

## 2. Frame & Propeller Guards (3D Printed)

> ✅ **Owner has a large-format printer (600×600×600mm build volume).** This comfortably exceeds a 5" frame's footprint, so the entire cage/duct frame can be printed as a single unibody piece (no split arms, no multi-part joins/standoffs between duct sections). Outsourced printing is no longer needed — that line item has been removed below.

| Qty | Item | Model / Specs | Price | Where to Buy |
|---|---|---|---|---|
| 1 | 5" Unibody Cage/Ducted Frame STL | Single-piece enclosed rotor-guard design (print in one piece using your printer's full bed) | Free–$20 (file only) | [Printables](https://printables.com) • [Thingiverse](https://thingiverse.com) |
| — | Filament | PETG or Nylon-CF, ~500g–1kg (single-piece prints use slightly more filament than split designs due to fewer optimized nesting cuts, but save on joint hardware) | $25–60 | [Prusament](https://prusament.com) • [Amazon](https://www.amazon.com) |
| 1 set | Hardware (motor screws, FC/GPS standoffs — no frame-joint hardware needed with a unibody print) | M2/M3 assorted kit | $8–12 | [Amazon](https://www.amazon.com) • local hardware store |

**Subtotal: ~$35–90** (no outsourcing cost; single build-plate print run)

> 💡 With 600×600×600mm of build volume, you also have headroom to print an **oversized duct/cage variant** (thicker guard rings, integrated camera mount, or a slightly larger prop clearance margin) in the same single print if you want extra impact protection — search Printables/Thingiverse for "5 inch unibody enclosed frame" or scale/remix a split design into one plate.

---

## 3. Dual-Band 915MHz Comms System

| Qty | Item | Model / Specs | Price | Where to Buy |
|---|---|---|---|---|
| 1 | RC Receiver | ExpressLRS (ELRS) 915MHz Nano/Diversity RX | $15–25 | [ExpressLRS-compatible vendors](https://www.expresslrs.org/hardware/) • [GetFPV](https://www.getfpv.com) |
| 1 | RC Transmitter Module | ELRS 915MHz TX module (if not using an existing radio) | $40–60 | [BetaFPV](https://betafpv.com) • [RadioMaster](https://www.radiomasterrc.com) |
| 1 | Telemetry Radio Pair | MicoAir LR900-F 915MHz LoRa telemetry set (air + ground unit) | $70–90 | [MicoAir Store](https://micoair.com) • [Holybro Store](https://holybro.com) |
| 1 | Ground-Side LoRa USB Dongle | USB adapter for ground LoRa unit | Included with LR900-F set (or $10–15 separate) | [MicoAir Store](https://micoair.com) |
| 1 | 915MHz Antenna Set (air + ground) | Duck/whip antennas, SMA/IPEX | $10–15 | [GetFPV](https://www.getfpv.com) |

**Subtotal: ~$145–185**

> ⚠️ 915MHz is legal for RC/telemetry use in North America and Australia. In Europe use 868MHz-band equivalents; check local regulations elsewhere.

---

## 4. Ground Control Station (Smartphone)

| Qty | Item | Model / Specs | Price | Where to Buy |
|---|---|---|---|---|
| 1 | Android Smartphone | Android 5.0+ (reused/old phone acceptable) | $0 (if reused) – $150 (budget new/used) | Existing device or [Swappa](https://swappa.com) / [Amazon Renewed](https://www.amazon.com) |
| 1 | USB-C OTG Adapter | For connecting LoRa dongle + Xbox controller | $8–12 | [Amazon](https://www.amazon.com) |
| 1 | Xbox Wireless Controller | USB-C or Bluetooth model | $0 (if owned) – $60 | [Microsoft Store](https://www.xbox.com/en-us/accessories/controllers) • [Amazon](https://www.amazon.com) |
| 1 | QGroundControl App | Free, open source GCS software | $0 | [qgroundcontrol.com](https://qgroundcontrol.com) |
| 1 | Phone Mount (optional, for field use) | Handheld or tripod mount | $10–15 | [Amazon](https://www.amazon.com) |

**Subtotal: ~$0–165** (depends heavily on whether phone/controller are already owned)

---

## 5. Battery & Charging

| Qty | Item | Model / Specs | Price | Where to Buy |
|---|---|---|---|---|
| 2–3 | Flight Batteries | 6S 1300–1500mAh LiPo, 90–100C | $30–40 each | [GetFPV](https://www.getfpv.com) • [RaceDayQuads](https://www.racedayquads.com) |
| 1 | Charger | 6S-capable balance charger (e.g., ISDT or SkyRC) | $40–70 | [GetFPV](https://www.getfpv.com) • [Amazon](https://www.amazon.com) |
| 1 | LiPo Safe Bag | Charging/storage safety bag | $8–12 | [Amazon](https://www.amazon.com) |

**Subtotal: ~$75–100** (with 2 batteries; add ~$35 per extra battery)

---

## 6. Camera Payload

| Qty | Item | Model / Specs | Price | Where to Buy |
|---|---|---|---|---|
| 1 | Action Camera | GoPro Session (budget) or DJI Osmo Action (better stabilization) | $150–300 | [GoPro Store](https://gopro.com) • [DJI Store](https://www.dji.com) |
| 1 | Camera Mount Bracket | Custom 3D printed mount (STL, printed with frame) | Included in frame filament cost | [Printables](https://printables.com) |
| 1 (optional) | Gimbal | Holybro A8 Mini 3-axis gimbal | $130–180 | [Holybro Store](https://holybro.com/products/a8-mini) |
| 1 | MicroSD Card | 64–128GB, U3/V30 rated | $12–20 | [Amazon](https://www.amazon.com) |

**Subtotal: ~$100–340** (higher end includes gimbal)

---

## 7. Tools & Misc

| Qty | Item | Notes | Price | Where to Buy |
|---|---|---|---|---|
| 1 | Soldering Iron + Solder | Temperature-controlled, for ESC/FC wiring | $30–60 | [Amazon](https://www.amazon.com) |
| 1 | Hex Driver Set | Metric, for frame/motor screws | $10–15 | [Amazon](https://www.amazon.com) |
| 1 | Multimeter | For continuity/voltage checks | $15–25 | [Amazon](https://www.amazon.com) |
| 1 | Heat Shrink Tubing Assortment | Wire connections | $8–10 | [Amazon](https://www.amazon.com) |
| 1 | XT60/XT30 Connectors | Battery/power connections, assorted pack | $8–10 | [Amazon](https://www.amazon.com) |
| 1 | Threadlocker (blue) | For motor screws | $6–8 | [Amazon](https://www.amazon.com) |
| — | Zip Ties, Foam Tape, Wire | General assembly consumables | $15–20 | [Amazon](https://www.amazon.com) |

**Subtotal: ~$100–120**

---

## Firmware & Software (Free)

| Item | Purpose | Link |
|---|---|---|
| ArduPilot or PX4 | Open-source flight controller firmware | [ardupilot.org](https://ardupilot.org) • [docs.px4.io](https://docs.px4.io) |
| QGroundControl | Ground control station software (Android) | [qgroundcontrol.com](https://qgroundcontrol.com) |
| ExpressLRS Configurator | ELRS firmware flashing/config | [expresslrs.org](https://www.expresslrs.org) |

---

## Notes

- Prices are estimates as of August 2025 and will vary by region and vendor stock.
- The wide cost range reflects optional upgrades (gimbal, new smartphone, outsourced 3D printing) vs. the leanest build using owned hardware.
- Buy 2–3 extra propeller sets — they are the most common consumable/crash casualty.
- Confirm 915MHz band legality in your country before purchasing ELRS/LoRa hardware (see Requirements Summary for details).

---

**Document Version:** 1.1
**Last Updated:** August 2025
**Companion Documents:** `drone_build_guide.md`, `drone_requirements_summary.md`
