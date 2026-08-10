# LoRa Still-Image Uplink (Experimental)

**Version:** 1.0
**Date:** August 2025
**Applies to:** `drone_purchase_list.md` §3 (Dual-Band 915MHz Comms System), `drone_build_guide.md` §10 (FPV Capability)
**Status:** Experimental / DIY — not a parts-list line item, no off-the-shelf support in this build's software stack
**Companion Documents:** `drone_build_guide.md`, `drone_purchase_list.md`, `drone_frame_merge_guide.md`, `drone_kit_alternatives.md`

---

## Why This Exists

`drone_build_guide.md` §10 covers real-time FPV video (analog or digital HD), but that video link's range is limited compared to the MicoAir LR900-F LoRa telemetry link this build already carries — the LR900-F is rated for line-of-sight range up to 30km, far beyond where any 5.8GHz analog or digital FPV video signal would have already dropped out.

That gap raises a reasonable question: can the existing LoRa telemetry radio *also* carry some form of image data, so there's still a rough visual check-in once you've flown beyond FPV video range but are still within telemetry range? The answer is yes, for periodic low-resolution stills — not for anything resembling real video.

## The Real Numbers

The MicoAir LR900-F's datasheet gives concrete throughput figures for the exact radio this build uses:

| Mode | Throughput | Notes |
|---|---|---|
| Default | 2.1 KB/s | Standard operating mode |
| Max (FHSS) | 3.2 KB/s down / 1.6 KB/s up | Best case, likely short range |
| Robust/long-range | 0.4–1.1 KB/s | Lower speed, more reliable at extended range |

Max single packet length is 20KB, over a 3.3V TTL UART running at 57600 baud by default.

## What That Buys You

- A heavily compressed, small JPEG (roughly 160×120 to 320×240, aggressive quality reduction) runs about **3–15KB**.
- At the default 2.1 KB/s — after subtracting whatever bandwidth normal MAVLink telemetry (GPS, battery, attitude) is already using — that works out to **roughly one thumbnail-quality still every 10–30 seconds**.
- At long range, where the more robust 0.4–1.1 KB/s modes are needed to hold the link, that stretches to **every 15–60+ seconds**.
- True video, even at 1 fps, isn't realistic except at very close range with telemetry traffic minimized. At that point you're just describing a faster still-image stream, not motion video — don't expect anything resembling a live feed.

## This Is a Known Technique, Not a New Idea

This is essentially **SSDV (Slow Scan Digital Video)** — the technique the high-altitude balloon and amateur radio community uses to send images down low-bandwidth links, including LoRa specifically. SSDV packetizes a JPEG so a partial or lossy image is still viewable even with dropped packets, which matters a lot at long range. See the [UKHAS Wiki SSDV guide](https://ukhas.org.uk/doku.php?id=guides:ssdv) for the reference implementation this concept borrows from.

## What It Would Actually Take to Build

This is not a settings toggle on the existing hardware — it's a real DIY project layered on top of the current build:

1. **Extra hardware.** The LR900-F is a transparent serial bridge carrying the flight controller's MAVLink stream — it has no camera input of its own. You'd need a companion computer (e.g., an ESP32-CAM or a Raspberry Pi Zero with a camera module) to capture and JPEG-compress a frame.
2. **Extra software, both ends.** Neither ArduPilot/PX4 nor QGroundControl has SSDV-style image decoding built in — that's a ham-radio/HAB-ecosystem tool, not a standard GCS feature. You'd write custom encode logic on the companion computer and either adapt an existing SSDV decoder or write a custom one for the ground side.
3. **Bandwidth contention.** Whatever bandwidth you spend on images comes out of the same link this build already relies on for GPS/battery/failsafe telemetry (`drone_build_guide.md` §7 step 5). Rate-limit image transfers carefully so they can't delay a low-battery RTL trigger or other failsafe data.

## Recommendation

Treat this as a "nice to have if you're already comfortable with companion-computer development," not a build requirement. If you just want real-time situational video, the FPV path in `drone_build_guide.md` §10 is the right tool — this uplink is specifically for the case where you've flown past FPV video range but are still within LoRa telemetry range and want a rough idea of what the camera is pointed at.

---

**Document Version:** 1.0
**Last Updated:** August 2025
**Companion Documents:** `drone_build_guide.md`, `drone_purchase_list.md`, `drone_frame_merge_guide.md`, `drone_kit_alternatives.md`
