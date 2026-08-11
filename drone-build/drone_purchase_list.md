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
| 8 | Propellers (endurance alternative — pair with lower-KV motors, see §9) | 5" **bi-blade** (2-blade), pitch ~3.5–4.5" (P/D ≈ 0.7–0.9), 6S-compatible, 2 sets incl. spares. Diameter stays 5" — the duct fixes it. Blade count is the real lever; pitch is a motor-loading match, not an efficiency gain — see build guide §11.2 | $8–14/set | [GetFPV](https://www.getfpv.com) • [RaceDayQuads](https://www.racedayquads.com) |
| 1 | Power Distribution Board | Matches ESC stack (often bundled) | $0–20 | Bundled with FC stack or [Holybro Store](https://holybro.com) |

**Subtotal: ~$275–325 (excluding spares; endurance alternative motors/props are a like-for-like swap, not an add-on)**

---

## 2. Frame & Propeller Guards (3D Printed)

> ✅ **Owner has a large-format printer (600×600×600mm build volume).** This comfortably exceeds a 5" frame's footprint. Note: genuinely single-piece "the whole cage prints as one part" designs are uncommon for 5" caged/ducted frames — the real designs below, like most enclosed-cage frames, print as a handful of parts (arms, center plate, duct sections) that bolt together, not one continuous piece. What the large bed actually buys you is printing all of those parts together in a single job instead of splitting the run across multiple prints or outsourcing. Outsourced printing is no longer needed — that line item has been removed below.

| Qty | Item | Model / Specs | Price | Where to Buy |
|---|---|---|---|---|
| 1 | 5" Caged/Ducted Frame — Option A | Replica of the Lumenier QAV-PRO Whoop 5" caged cinewhoop frame; full propeller-guard duct cage; files include STL + STEP + DXF; **gimbal mount:** modular cinema mount, good odds it fits the Holybro A8 Mini — test-fit first | Free–$20 (confirm current price on listing) | [virk3D — Printables](https://www.printables.com/model/1511864-cinewhoop-5-inch-frame-replica-of-lumenier-qav-pro) |
| 1 | 5" Caged/Ducted Frame — Option B | Same Lumenier QAV-PRO Whoop 5" replica design, independently modeled by a different designer; **gimbal mount:** same modular cinema mount as Option A — test-fit first | ~$57–59 (paid) | [msnikon1995 — Cults3D](https://cults3d.com/en/3d-model/various/frame-cinewhoop-5inch-replica-lumenier-qav-pro-whoop-5-msnikon1995) |
| 1 | 5" Ducted Frame — Option C | Distinct 3-part ducted 5" frame design (not a Lumenier replica) — **already the leanest part count of the three options here**; designer recommends PETG or ABS for strength (PLA+ also reported to work), TPU for landing feet; **gimbal mount:** none confirmed in the design — see note below | Free | [carnoforge3D — Cults3D](https://cults3d.com/en/3d-model/various/5-inch-fpv-quadcopter-frame-with-duted-props-freebie) |
| — | Filament | PETG or Nylon-CF, ~500g–1kg depending on which option above you print | $25–60 | [Prusament](https://prusament.com) • [Amazon](https://www.amazon.com) |
| 1 set | Hardware (motor screws, FC/GPS standoffs, plus small frame-joint screws — these are multi-part designs, not a single continuous piece) | M2/M3 assorted kit | $8–12 | [Amazon](https://www.amazon.com) • local hardware store |

**Subtotal: ~$35–130** (as low as ~$35 if you pick the free Option C; higher if you go with a paid frame file — all parts for any option print together in a single build-plate job on this bed)

> 🎥 **Retaining the gimbal:** the Holybro A8 Mini gimbal (§6) stays in the parts list by default regardless of which frame option you pick. If you go with Option C (no confirmed gimbal mount), print a small custom adapter plate — drilled for the A8 Mini's 4× M2.5×8mm mounting screws on one side and your frame's standoff pattern on the other — using filament already budgeted above. See build guide §3 step 5 for the full note.

> 🧩 **Fewer, larger parts?** A follow-up search specifically for a 5" *enclosed/ducted cage* design purpose-built with fewer, larger parts (not just Option A/B's usual arms+center+duct split) didn't turn up a real match — the closest hit, ARS-5, is a well-known "few parts" printable 5" frame, but it's an open racing frame without propeller guards, so it doesn't meet this build's enclosed-cage spec and wasn't substituted in. Option C's 3-part count above is already the leanest of the verified enclosed-cage options. If you want fewer parts than that, see `drone_frame_merge_guide.md` for a guide to merging Option A's pieces yourself in CAD (Option A is used there because it includes a STEP file — Option C is STL-only, which is a much less reliable starting point for a boolean merge).

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
| 2–3 | Flight Batteries (baseline) | 6S 1300–1500mAh LiPo, 90–100C+ — see §5a for specific models and dimensions | $30–40 each | [GetFPV](https://www.getfpv.com) • [RaceDayQuads](https://www.racedayquads.com) |
| 2–3 | Flight Batteries (endurance alternative — see §9) | 6S1P Li-ion pack, high-drain 21700 cells (Molicel P45B preferred, P42A acceptable) — see §5b | $60–100 each | [GetFPV](https://www.getfpv.com) • [Amazon](https://www.amazon.com) |
| 1 | Charger | 6S-capable balance charger with a **Li-ion charge profile** (e.g., ISDT Q6/Q8, SkyRC) — needs XT60 output and JST-XH 7-pin balance support | $40–70 | [GetFPV](https://www.getfpv.com) • [Amazon](https://www.amazon.com) |
| 1 | LiPo/Li-ion Safe Bag | Charging/storage safety bag | $8–12 | [Amazon](https://www.amazon.com) |

**Subtotal: ~$75–100 (LiPo baseline, 2 batteries)** — swapping to Li-ion (§9) runs ~$120–200 for 2 packs; add ~$60–100 per extra Li-ion pack vs. ~$35 per extra LiPo pack.

### 5a. Baseline LiPo — Specific Models & Dimensions

All three below are 6S1P, XT60 discharge lead, JST-XH 7-pin balance lead. Pick on price and availability; there is no meaningful endurance difference between them at this size.

| Model | Capacity / C | Dimensions (L × W × H) | Weight | Energy |
|---|---|---|---|---|
| CNHL Black Series 6S 1300mAh | 1300mAh, 100C (130C V2) | 76 × 35 × 47 mm | 230 g (incl. wire + connector) | ~28.9 Wh |
| Tattu R-Line V3 6S 1300mAh | 1300mAh, 120C | 74 × 36 × 43.5 mm | 218.5 g | ~28.9 Wh |
| Tattu FunFly 6S 1300mAh | 1300mAh, 100C | 74 × 35 × 45 mm | 225 g | ~28.9 Wh |

**Design envelope for the LiPo bay: 76 × 36 × 47 mm, 205–235 g.** Add 2–3 mm clearance per axis plus strap thickness. A 1500mAh pack in this class runs a few mm longer and ~20–30 g heavier (~33.3 Wh) — if you want the option of either capacity, size the bay to ~82 × 38 × 50 mm.

> ⚠️ **C ratings on LiPo are marketing.** A "100C" 1300mAh pack does not deliver 130A. For this build it doesn't matter — a cruising camera platform never approaches these packs' real limits — so buy on capacity, weight, and dimensions rather than chasing the highest C number.

### 5b. Endurance Li-ion — Cells, Pack Geometry & Dimensions

**Cell specifications** (both are 21700 format, flat top):

| Cell | Capacity | Nominal V | Max continuous discharge | Dimensions | Weight | Energy |
|---|---|---|---|---|---|---|
| **Molicel INR21700-P45B** (preferred) | 4500 mAh | 3.6 V | 45 A | 21.55 × 70.15 mm (max) | ~70 g | 16.2 Wh |
| Molicel INR21700-P42A | 4200 mAh | 3.6 V | 45 A | 21.7 × 70.2 mm (max) | ~70 g | 15.5 Wh |

**Assembled 6S1P P45B pack:** 4500 mAh at 21.6 V nominal (25.2 V full, 6× 4.2 V) = **~97 Wh**, **450–470 g** including nickel strip, wrap, and leads. That is ~3.4× the energy of the 1300mAh LiPo at ~2.1× the weight — the trade that produces the flight-time gain in §9.

**Pack geometry — this is the part that affects your frame print.** A 6S1P of 21700 cells can be assembled in two common shapes, and they are not interchangeable in a 5" cage:

| Geometry | Dimensions | Notes |
|---|---|---|
| **Flat / long** (e.g., Flywoo Explorer P45B 6S1P 4500mAh) | ~152 × 43 × 40 mm (L × W × H), ~455 g | Commercially available, XT60 + JST-XH 7-pin balance. Long axis runs fore-aft along the top plate — the usual long-range mounting approach. **Confirm exact dimensions on the listing before printing.** |
| **3 × 2 brick** (custom / spot-welded) | ~65 × 43 × 70 mm bare cells, ~67 × 45 × 72 mm wrapped (derived from cell dims above) | More compact footprint but ~70 mm tall, which is taller than most cinewhoop stacks allow above the top plate. Requires spot welding — do not solder directly to cell terminals. |

> 🔋 **Recommended approach:** design the §2 frame's battery bay around the **LiPo envelope in §5a** and mount the Li-ion pack externally on the top plate with a strap, rather than trying to size one bay for both. The two form factors are too different to share a bay, and the flat/long Li-ion geometry is designed to be strapped on top anyway. Check CG fore-aft after mounting — a 152 mm pack has real leverage on a ~200 mm airframe.

> ⚠️ **The Li-ion pack becomes the current bottleneck of the entire power path.** A 6S1P pack has no parallel cells to share load, so its ceiling is one cell's rating: **45 A for the whole aircraft**. Hover draw (~10–15 A) is comfortable, but four 2207 motors at 6S can exceed 45 A on an aggressive climb or full-throttle input — below even the power module's 60 A burst lead rating. This is not optional tuning: set a current limit in firmware before the first Li-ion flight (see build guide §11.1) and fly it as a cruiser.

> ⚠️ **Higher-capacity ≠ better here.** Packs built on 5000mAh cells (e.g., Samsung INR21700-50S, sold as GNB 6S 5000mAh "10C") trade continuous-current capability for capacity, and the "10C" figure on Li-ion packs is optimistic. If you consider one, verify the *cell's* datasheet continuous rating — not the pack's marketing number — against the current headroom warning above.

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
| Li-ion battery instead of LiPo | +$45–100 for 2 packs vs. baseline LiPo | +~220–250g per pack (450–470g vs. 220–230g), offset by ~3.4× the energy — see §5b | §5 |
| Lower-KV motor + 5" bi-blade prop (pitch matched to KV) | Roughly cost-neutral (like-for-like swap) | Neutral to slightly lower (one less blade per prop) | §1 |
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
- If switching to Li-ion batteries, also set a firmware current limit before the first flight — a 6S1P pack's discharge ceiling is one cell's rating (45A for the recommended Molicel cells), with no parallel cells to share load. See §5b and build guide §11.1.

---

**Document Version:** 1.7
**Last Updated:** August 2026
**Companion Documents:** `drone_build_guide.md`, `drone_frame_merge_guide.md`, `drone_kit_alternatives.md`, `drone_lora_still_image_uplink.md`, `drone_requirements_summary.md`
