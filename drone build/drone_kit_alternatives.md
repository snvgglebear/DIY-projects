# Drone Kit Alternatives — Hybrid Approach

**Version:** 1.0
**Date:** August 2025
**Applies to:** `drone_purchase_list.md` §1 (Flight Controller & Core Electronics)
**Companion Documents:** `drone_build_guide.md`, `drone_purchase_list.md`, `drone_frame_merge_guide.md`

---

## Why Hybrid

No commercial kit currently bundles this build's two defining requirements together: an enclosed propeller-guard cage *and* Pixhawk/ArduPilot autonomy (waypoints, GPS RTL, QGroundControl monitoring). The closest match, the **Holybro QAV250 Kit**, gets the Pixhawk/ArduPilot electronics right but ships with a bare, open racing frame — no prop guards. The hybrid approach takes the QAV250 kit's electronics bundle for the cost savings, and keeps everything else (frame, comms, camera payload, ground station) from the existing custom parts list.

## What to Buy From the Kit

| Kit Item | Replaces (purchase list ref) | Notes |
|---|---|---|
| Pixhawk 6C **Mini** | FC Stack (§1) | Smaller form factor than the standard Pixhawk 6C — fewer ports; check port count against what §4/§5 of the build guide actually wires up (GPS, telemetry, ELRS receiver, ESC signal) before committing |
| M10 GPS module | GPS Module (§1) | Same GPS this build already specs — no functional change |
| PM06 V2 power module | Power Distribution Board (§1) | A power module (regulated 5V + voltage/current sensing to the FC), not a full PDB — see integration note below |
| 4× 2207 KV1950 motors | Motors (§1) | Higher KV than the baseline 1700KV spec — see conflict note below |
| 4× 20A ESCs (individual) | ESC 4-in-1 (§1) | Individual ESCs, not a 4-in-1 — more wiring, and well under the 55A 4-in-1 headroom already specced |
| *(Complete kit only)* 5.8GHz analog VTX combo | FPV System (`drone_build_guide.md` §10), optional | Bonus if you're adding FPV anyway — covers the analog path for free; the digital HD path in §10 is still the better image-quality option if you want it |
| *(Kit)* SiK telemetry radio, 915MHz variant | Telemetry Radio Pair (§3), optional | See the telemetry sub-decision below — this can replace or complement the MicoAir LR900-F line item |

**Choose Basic (~$142.59) or Complete (~$288):** Complete adds the analog VTX combo; skip it if you're going digital HD for FPV, since it'd be unused.

## What to Keep From the Custom Build

| Item | Purchase list ref | Why it stays |
|---|---|---|
| Frame (Options A/B/C) | §2 | The kit's frame is open/exposed with no propeller guards — this build's core safety requirement. The kit's own frame goes unused. |
| ExpressLRS RX + TX module | §3 | The kit includes no RC receiver at all — this is needed regardless of kit vs. custom electronics. |
| MicoAir LR900-F LoRa telemetry (probably) | §3 | See telemetry sub-decision below. |
| Smartphone + Xbox controller ground station | §4 | Unaffected by the electronics swap. |
| Battery & charger | §5 | Unaffected, but re-verify 6S compatibility against the kit's 20A ESCs before assuming it's a drop-in swap — see integration notes. |
| Camera, mount, gimbal | §6 | Unaffected. |
| Propellers | §1 | The kit's props (if any come with it) are sized for its own open racing frame — you'll still need props matched to whichever cage/duct clearance you printed. |

### Telemetry sub-decision

The kit's bundled SiK radio and the purchase list's MicoAir LR900-F both do the same job — MAVLink telemetry to your ground station — but they aren't equivalent. SiK is the older, simpler, shorter-range option; LR900-F was chosen for this build specifically for better range/reliability. Two paths:
- **Keep LR900-F** (buy it separately per §3, skip using the kit's SiK radio) — preserves the build's original telemetry range spec.
- **Use the kit's SiK radio instead** — saves the §3 telemetry line item cost, at the cost of shorter-range/less-reliable telemetry than originally speced.

## Integration Notes

- **PDB gap:** switching from a 4-in-1 ESC (which typically includes its own power distribution) to 4 individual ESCs means the PM06 power module alone likely isn't enough to distribute battery power to all four ESCs — budget for a small separate PDB or plan to wire the ESCs in parallel directly off the battery leads.
- **ESC headroom:** 20A per motor is well below the 55A 4-in-1 originally specced. That's normally fine for a light open racer, but this build carries a heavier enclosed cage plus a camera/gimbal payload — verify the 20A rating actually covers your motor's max draw at 6S before flying, not just "6S compatible" on the label.
- **Motor KV conflict:** the kit's 2207 KV1950 motors are *higher* KV than this build's 1700KV baseline — the opposite direction from the lower-KV endurance optimization in `drone_build_guide.md` §11. If flight time matters to you, the kit's bundled motors work against that goal; you'd need to source lower-KV motors separately, giving up some of the kit's cost advantage.
- **Prop clearance:** whichever motor/prop combo you end up with (kit's or your own), re-check it against the duct clearance of whichever frame option (A/B/C) you printed, per `drone_build_guide.md` §11 item 2.

## Cost Comparison

| | FC + GPS + Motors + ESC + PDB, priced separately | Holybro QAV250 Kit (Basic) |
|---|---|---|
| Cost | ~$390–549 | ~$142.59 |
| Savings | — | **~$250–400**, minus ~$10–15 for a supplemental PDB if needed |

This is a savings estimate for that specific slice of the parts list only — frame, comms, battery, camera payload, and ground station costs are unchanged either way.

## Tradeoffs of the Hybrid Approach Specifically

- **Wasted kit components:** you're paying for the kit's frame (and its 5.8GHz VTX in the Complete version, if unused) without using them — the effective savings are smaller than the sticker price suggests once you back that out.
- **Mismatched, unvalidated system:** the whole point of a kit is that the vendor tested the frame/motor/ESC/FC/firmware combination together. Once you swap the kit's electronics into a different (3D-printed) frame with different motors/props than the kit intended, you've recreated a custom, unvalidated build — you lose the "known good combination" benefit that's normally the point of buying a kit at all.
- **Split support/warranty:** troubleshooting now spans two purchase paths (Holybro's kit electronics, your separately-sourced frame/comms/etc.) instead of one vendor or one fully custom build.

---

**Document Version:** 1.0
**Last Updated:** August 2025
**Companion Documents:** `drone_build_guide.md`, `drone_purchase_list.md`, `drone_frame_merge_guide.md`
