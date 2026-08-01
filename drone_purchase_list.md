# Open-Source Camera Drone — Purchase List

**Build Goal:** Compact camera-capable drone with 3D printed propeller-guard frame, dual-band 915MHz comms (ELRS + LoRa), smartphone/Xbox ground control station
**Estimated Total:** ~$900–1,250 (depending on smartphone availability); +$135–970 if adding optional FPV (§8); flight-time optimizations in §9 can add or save cost depending on which are chosen
**Last Updated:** August 2025
**Note:** Frame cost reduced — owner has a 600×600×600mm printer, so the full caged/ducted frame (all parts) is printed in-house in one job (no outsourced print service needed).

---

## Quick Summary Table

| Category | Estimated Cost |
|---|---|
| Flight Controller & Core Electronics | $275–325 |
| Frame & Propeller Guards (3D Printed, home-printed, 3 file options — see §2) | $35–130 |
| Dual-Band 915MHz Comms System | $145–185 |
| Ground Control Station (Smartphone) | $0–165 |
| Battery & Charging | $75–100 |
| Camera Payload | $100–340 |
| Tools & Misc | $100–120 |
| **TOTAL** | **~$730–1,365** |
| FPV System (Optional Add-On, not in TOTAL above) | +$135–970 |
| Li-ion Battery Upgrade (Optional, replaces LiPo cost in Battery & Charging above) | +$45–100 vs. baseline LiPo (2 packs) |
| Skip Gimbal for Fixed Mount (Optional, reduces Camera Payload above) | −$130–180 |

---

## 1. Flight Controller & Core Electronics

| Qty | Item | Model / Specs | Price | Where to Buy |
|---|---|---|---|---|
| 1 | Flight Controller Stack | Holybro Pixhawk 6C (with plastic case) | $220–299 | [Holybro Store](https://holybro.com/products/pixhawk-6c) • [ReadyMadeRC](https://www.readymaderc.com/products/details/86931-holybro-pixhawk-6c-with-plastic-case) • [PyroDrone](https://pyrodrone.com/products/holybro-pixhawk-6c-plastic-case) |
| 1 | GPS Module | Holybro M8N (budget) or M10 (better accuracy) | $25–45 | [Holybro Store](https://holybro.com/collections/gps) |
| 4 | Motors | T-Motor Velox 2207 1700KV | $22–28 each (~$90–110 set) | [T-Motor Store](https://www.tmotor.com) • [GetFPV](https://www.getfpv.com) |
| 4 | Motors (endurance alternative — see §9) | T-Motor Velox 2207 ~1400–1500KV | $22–30 each (~$90–120 set) | [T-Motor Store](https://www.tmotor.com) • [GetFPV](https://www.getfpv.com) |
| 1 | ESC (4-in-1) | 55A BLHeli_32 4-in-1 ESC | $55–75 | [GetFPV](https://www.getfpv.com) • [PyroDrone](https://pyrodrone.com) |
| 8 | Propellers | 5040 tri-blade, 6S-compatible (2 sets, incl. spares) | $8–12/set | [GetFPV](https://www.getfpv.com) • [RaceDayQuads](https://www.racedayquads.com) |
| 8 | Propellers (endurance alternative — pair with lower-KV motors, see §9) | Larger diameter or higher-pitch 6S-compatible, 2 sets incl. spares | $8–14/set | [GetFPV](https://www.getfpv.com) • [RaceDayQuads](https://www.racedayquads.com) |
| 1 | Power Distribution Board | Matches ESC stack (often bundled) | $0–20 | Bundled with FC stack or [Holybro Store](https://holybro.com) |

**Subtotal: ~$275–325 (excluding spares; endurance alternative motors/props are a like-for-like swap, not an add-on)**

---

## 2. Frame & Propeller Guards (3D Printed)

> ✅ **Owner has a large-format printer (600×600×600mm build volume).** This comfortably exceeds a 5" frame's footprint. Note: genuinely single-piece "the whole cage prints as one part" designs are uncommon for 5" caged/ducted frames — the real designs below, like most enclosed-cage frames, print as a handful of parts (arms, center plate, duct sections) that bolt together, not one continuous piece. What the large bed actually buys you is printing all of those parts together in a single job instead of splitting the run across multiple prints or outsourcing. Outsourced printing is no longer needed — that line item has been removed below.

| Qty | Item | Model / Specs | Price | Where to Buy |
|---|---|---|---|---|
| 1 | 5" Caged/Ducted Frame — Option A | Replica of the Lumenier QAV-PRO Whoop 5" caged cinewhoop frame; full propeller-guard duct cage; files include STL + STEP + DXF; **gimbal mount:** modular cinema mount, good odds it fits the Holybro A8 Mini — test-fit first | Free–$20 (confirm current price on listing) | [virk3D — Printables](https://www.printables.com/model/1511864-cinewhoop-5-inch-frame-replica-of-lumenier-qav-pro) |
| 1 | 5" Caged/Ducted Frame — Option B | Same Lumenier QAV-PRO Whoop 5" replica design, independently modeled by a different designer; **gimbal mount:** same modular cinema mount as Option A — test-fit first | ~$57–59 (paid) | [msnikon1995 — Cults3D](https://cults3d.com/en/3d-model/various/frame-cinewhoop-5inch-replica-lumenier-qav-pro-whoop-5-msnikon1995) |
| 1 | 5" Ducted Frame — Option C | Distinct 3-part ducted 5" frame design (not a Lumenier replica); designer recommends PETG or ABS for strength (PLA+ also reported to work), TPU for landing feet; **gimbal mount:** none confirmed in the design — see note below | Free | [carnoforge3D — Cults3D](https://cults3d.com/en/3d-model/various/5-inch-fpv-quadcopter-frame-with-duted-props-freebie) |
| — | Filament | PETG or Nylon-CF, ~500g–1kg depending on which option above you print | $25–60 | [Prusament](https://prusament.com) • [Amazon](https://www.amazon.com) |
| 1 set | Hardware (motor screws, FC/GPS standoffs, plus small frame-joint screws — these are multi-part designs, not a single continuous piece) | M2/M3 assorted kit | $8–12 | [Amazon](https://www.amazon.com) • local hardware store |

**Subtotal: ~$35–130** (as low as ~$35 if you pick the free Option C; higher if you go with a paid frame file — all parts for any option print together in a single build-plate job on this bed)

> 🎥 **Retaining the gimbal:** the Holybro A8 Mini gimbal (§6) stays in the parts list by default regardless of which frame option you pick. If you go with Option C (no confirmed gimbal mount), print a small custom adapter plate — drilled for the A8 Mini's 4× M2.5×8mm mounting screws on one side and your frame's standoff pattern on the other — using filament already budgeted above. See build guide §3 step 5 for the full note.

> 💡 With 600×600×600mm of build volume, you also have headroom to print an **oversized duct/cage variant** (thicker guard rings, integrated camera mount, or a slightly larger prop clearance margin) if you want extra impact protection — search Printables/Cults3D for "5 inch cinewhoop caged frame" or "5 inch ducted frame" (searching "unibody" mostly surfaces single-arm designs, not full cages), or scale/remix one of the three options above.

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
| 2–3 | Flight Batteries (baseline) | 6S 1300–1500mAh LiPo, 90–100C | $30–40 each | [GetFPV](https://www.getfpv.com) • [RaceDayQuads](https://www.racedayquads.com) |
| 2–3 | Flight Batteries (endurance alternative — see §9) | 6S Li-ion pack, high-drain 21700 cells (e.g., Molicel P42A/P45B) | $60–100 each | [GetFPV](https://www.getfpv.com) • [Amazon](https://www.amazon.com) |
| 1 | Charger | 6S-capable balance charger with Li-ion charge profile (e.g., ISDT or SkyRC) | $40–70 | [GetFPV](https://www.getfpv.com) • [Amazon](https://www.amazon.com) |
| 1 | LiPo/Li-ion Safe Bag | Charging/storage safety bag | $8–12 | [Amazon](https://www.amazon.com) |

**Subtotal: ~$75–100 (LiPo baseline, 2 batteries)** — swapping to Li-ion (§9) runs ~$120–200 for 2 packs; add ~$60–100 per extra Li-ion pack vs. ~$35 per extra LiPo pack.

---

## 6. Camera Payload

| Qty | Item | Model / Specs | Price | Where to Buy |
|---|---|---|---|---|
| 1 | Action Camera | GoPro Session (budget) or DJI Osmo Action (better stabilization) | $150–300 | [GoPro Store](https://gopro.com) • [DJI Store](https://www.dji.com) |
| 1 | Camera Mount Bracket | Custom 3D printed mount (STL, printed with frame) | Included in frame filament cost | [Printables](https://printables.com) |
| 1 (optional) | Gimbal | Holybro A8 Mini 3-axis gimbal | $130–180 | [Holybro Store](https://holybro.com/products/a8-mini) |
| 1 | MicroSD Card | 64–128GB, U3/V30 rated | $12–20 | [Amazon](https://www.amazon.com) |

**Subtotal: ~$100–340** (higher end includes gimbal; low end assumes gimbal skipped)

> 💡 Skipping the gimbal and mounting the action camera on the fixed 3D-printed bracket above removes this entire $130–180 line item and ~130–180g of flight weight. See build guide §11 for the stabilization tradeoff this involves.

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

## 8. FPV System (Optional Add-On)

> This section is optional — it's not included in the TOTAL in the summary table above. See the build guide §10 for installation and setup steps. Pick analog or digital HD as a matched set rather than mixing components across systems.

| Qty | Item | Model / Specs | Price | Where to Buy |
|---|---|---|---|---|
| 1 | FPV Camera | 1200TVL analog (budget), or digital HD camera module bundled with a system below | $15–90 | [GetFPV](https://www.getfpv.com) • [Foxeer](https://www.foxeer.com) |
| 1 | Video Transmitter (VTX) | 5.8GHz analog 25–800mW, or digital HD air unit (DJI O4 / HDZero / Walksnail Avatar) | $25–250 | [GetFPV](https://www.getfpv.com) • [DJI Store](https://www.dji.com/fpv) • [HDZero](https://www.hd-zero.com) |
| 1 | FPV Goggles | Analog goggles (budget), or digital HD goggles matched to the VTX system | $80–599 | [GetFPV](https://www.getfpv.com) • [DJI Store](https://www.dji.com/fpv) • [HDZero](https://www.hd-zero.com) |
| 1 set | 5.8GHz Antennas (air + goggle) | Circular polarized recommended over linear | $15–30 | [GetFPV](https://www.getfpv.com) |

**Subtotal: ~$135–970** (wide range: budget analog kit vs. full digital HD system)

> 💡 Digital HD systems (DJI O4, HDZero, Walksnail) typically bundle the camera and VTX into one "air unit," which simplifies wiring compared to sourcing a separate analog camera and VTX.

---

## 9. Flight Time Optimization (Optional)

> This section summarizes the parts-list impact of the endurance tradeoffs described in build guide §11. None of these are new categories — they modify or replace line items already listed in §1, §5, and §6 above.

| Change | Cost Impact | Weight Impact | Where Listed |
|---|---|---|---|
| Li-ion battery instead of LiPo | +$45–100 for 2 packs vs. baseline LiPo | Roughly neutral (similar pack weight, much higher energy density) | §5 |
| Lower-KV motor + larger/higher-pitch prop | Roughly cost-neutral (like-for-like swap) | No change | §1 |
| Skip gimbal, use fixed 3D-printed mount | −$130–180 (gimbal removed entirely) | −130–180g | §6 |

**Net effect:** all three together typically save money (the gimbal removal outweighs the Li-ion upgrade cost) while extending flight time — see build guide §11 for the performance/stabilization tradeoffs each one carries.

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
- The wide cost range reflects optional upgrades (gimbal, new smartphone, outsourced 3D printing, FPV system) vs. the leanest build using owned hardware.
- Buy 2–3 extra propeller sets — they are the most common consumable/crash casualty.
- Confirm 915MHz band legality in your country before purchasing ELRS/LoRa hardware (see Requirements Summary for details).
- If adding FPV, confirm local 5.8GHz video transmitter power limits before buying a high-power analog or digital system — limits vary by country and are generally more permissive in the US than in the EU.
- If switching to Li-ion batteries (§9), confirm your charger has a Li-ion charge profile before relying on it — Li-ion charge voltage/current curves differ from LiPo, and not all "6S-capable" chargers support both chemistries out of the box.

---

**Document Version:** 1.5
**Last Updated:** August 2025
**Companion Documents:** `drone_build_guide.md`, `drone_requirements_summary.md`
