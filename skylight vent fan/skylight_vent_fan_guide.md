# Skylight-Linked Bathroom Vent Fan — Design Guide

**Version:** 1.0
**Date:** August 2026
**Subject:** 24" × 32" chain-operated venting skylight over a bathroom light well (Philadelphia row house)
**Goal:** Mount an exhaust fan in the ceiling opening below the skylight, fill the rest of the opening with clear plastic so daylight still comes through, and have the skylight open automatically when the fan runs and close when it stops.
**Companion Document:** `skylight_vent_fan_purchase_list.md`

---

## Table of Contents

1. [Identifying Your Skylight](#1-identifying-your-skylight)
2. [The Central Design Question: Where Does the Air Actually Go?](#2-the-central-design-question-where-does-the-air-actually-go)
3. [Recommended Design](#3-recommended-design)
4. [The Clear Ceiling Panel](#4-the-clear-ceiling-panel)
5. [Motorizing the Skylight](#5-motorizing-the-skylight)
6. [Control Wiring — Three Tiers](#6-control-wiring--three-tiers)
7. [Weather, Winter, and Condensation](#7-weather-winter-and-condensation)
8. [Install Sequence](#8-install-sequence)
9. [Things That Will Bite You](#9-things-that-will-bite-you)
10. [Code Notes](#10-code-notes)

---

## 1. Identifying Your Skylight

### The short answer

I can't identify it with certainty from a description — you'll need to look at it (see "How to confirm" below). But the description is specific enough to narrow it a lot.

**Most likely: a Ventarama venting skylight**, or a close functional equivalent using a Truth Hardware / AmesburyTruth chain operator.

### Why that's the best guess

Three details in your description point the same direction:

1. **Two chains hanging down that you pull to open and close.** This is not two independent chains — it's almost certainly **one continuous chain loop running over a sprocket**, with both strands hanging into the room. Pull the front strand and the sprocket winds the sash open; pull the back strand and it winds it closed. That chain-and-sprocket operator is the signature of this whole class of skylight. Truth Hardware's manual skylight operator line is built exactly this way — [hardened steel sprocket and chain with an acetal case liner](https://www.truth.com/products/Hinged-Window-Hardware/Skylights/Manual-Operators), with the chain detachable at the sash — and Truth operators were the OEM hardware in a lot of these units.

2. **A hinged sash over a light well in a row-house bathroom.** This exact configuration is described in [US Patent 3,350,819, "Ventilating skylight with two simultaneously operated closures"](https://patents.google.com/patent/US3350819A/en) — a skylight serving a bathroom in the *central* portion of a row-type home, reached by a duct or well running from the bathroom ceiling up through the roof, with a roller-and-chain operated hinged panel. That's your house, and that patent's era (mid-1960s) is roughly the era of a lot of Philly row-house skylight retrofits.

3. **It's in Philadelphia.** Ventarama (Northport, NY) sold heavily into the Northeast row-house market — Philadelphia, Baltimore, NYC — for decades. Local roofers still call these "pull chain skylights" as a generic term.

Other candidates worth ruling out when you look at it: **CrystaLite**, which builds a [venting skylight with a PVC leg dropping into the light well as a mounting surface for a chain-drive Truth opener](https://crystaliteinc.com/pages/products/skylights/venting.php) — a very close match to your description; and any generic curb-mount vent unit that a previous owner fitted with an aftermarket Truth operator.

### How to confirm

Do this before you buy anything — the answer changes which of the paths in §5 makes sense.

- **Look for a nameplate.** Usually a foil or stamped label on the sash frame or the inside of the curb liner, visible when the sash is cranked fully open. Ventarama units are typically labeled.
- **Photograph the operator housing.** The plastic or metal box the chain runs out of, up at the top of the well. Truth/AmesburyTruth parts carry a molded part number. Post the photo to [SWISCO's discussion forum](https://www.swisco.com/discussions) — they specialize in obsolete window and skylight hardware and are good at IDing these from pictures. Several Ventarama threads already exist there.
- **Confirm it's one loop, not two chains.** Grab both strands and pull gently in opposite directions. If they're a loop over a sprocket, they'll fight each other. If they move independently, you have a different (rarer) two-line pulley design and the retrofit in §5 needs rethinking.
- **Measure the curb, not the glass.** Get the *inside* dimension of the curb (the wood box the skylight sits on) and its *outside* dimension. Your "24 × 32" is probably one of those, and which one matters enormously for replacement.

### The bad news, if it is a Ventarama

**Ventarama Skylight Corp. appears to be out of business.** Replacement parts still circulate: SWISCO stocks the [39-214 metal hook](https://www.swisco.com/Metal-Hook/pd/Replacement-Window-Operator-Accessories/39-214) and a [replacement skylight operator with angle drive](https://www.swisco.com/Skylight-operator/pd/Awning-Window-Replacement-Operators/39-379), and at least one vendor advertises aftermarket replacement tops for Ventarama units. But don't plan a design that depends on getting a specific factory part.

### On the 24 × 32 size

Note that **24 × 32 is not a VELUX standard curb size.** VELUX curb-mount sizes go 2222 (22½ × 22½ inside curb), 2230 (22½ × 30½), 2234 (22½ × 34½), 2246, 3030, 3046 — with the outside curb 3" larger in each direction. So if you ever replace the unit outright, you're either **building a new curb** to a VELUX size or **ordering a custom-size unit** (Wasco and Sun-Tek both do made-to-order). Factor that into §5 Path C.

---

## 2. The Central Design Question: Where Does the Air Actually Go?

This is the part worth thinking hardest about, because the obvious version of your idea has a real problem.

**The obvious version:** fan in the ceiling panel, blowing straight up into the light well. Skylight opens, moist air leaves through it. Clean, simple, no ductwork.

**The problem:** the light well is not a duct, and the code doesn't treat it as one. [IRC M1507.2](https://www.jaspector.com/codes/irc-2024/ch15-exhaust-systems/bathroom-exhaust-duct-requirements-irc-2024/) requires bathroom exhaust to discharge **directly to the outdoors** and prohibits discharging "into an attic, crawl space or other areas inside the building." A framed light well with a closed skylight on top is squarely an "area inside the building." An inspector will call it, and more importantly:

- **When the skylight doesn't open, you've built a moisture chamber.** Actuator fails, rain sensor holds it shut, power blip, chain jumps the sprocket — and now you're pumping shower steam into a sealed insulated box with cold glass on top. That's a mold factory in the joist bay, and you won't see it happening.
- **Even when it works, you get condensation.** Warm wet air hits 20°F glazing, condenses, and runs back down the well onto your clear panel. In a Philadelphia winter this is not hypothetical.
- **You're heating the neighborhood.** Every fan cycle opens a 24×32 hole in the roof.

So there are three honest paths. I'd build **Path B**.

| | **Path A — Well as plenum** | **Path B — Ducted (recommended)** | **Path C — Replace the skylight** |
|---|---|---|---|
| Fan exhausts into | The open light well | A real duct to a roof cap | A real duct to a roof cap |
| Skylight motorized? | Required — it's the only exit | Optional bonus / boost vent | Yes, factory electric unit |
| Code | Fails M1507.2 as written | Compliant | Compliant |
| Fails safe? | **No** — stuck-shut = trapped steam | Yes — fan works regardless | Yes |
| Cost | $ | $$ | $$$$ |
| Effort | Low | Medium (roof penetration) | High (roofer required) |

**Path B keeps everything you asked for** — fan in the ceiling opening, clear plastic around it, skylight opens when the fan runs — it just doesn't make the skylight *load-bearing* for the moisture. The skylight opening becomes a genuine bonus: a big passive stack vent that dumps heat and humidity fast, on top of a fan that would work fine on its own.

---

## 3. Recommended Design

```
                    ┌─── open skylight sash ───┐
   roof line   ═════╧══════════════════════════╧═════  ← existing curb
                    │                          │
                    │   light well / shaft     │   ┌── 4" insulated duct
                    │                          │   │
                    │  ┌────────────────────┐  │   │   ┌─ Broan 636 roof cap
                    │  │  inline fan in a   │──┼───┘   │  (separate penetration
                    │  │  joist bay beside  │  │       │   through the roof)
                    │  │  the shaft         │  │       │
                    │  └─────────┬──────────┘  │
                    │            │ 6" duct     │
   ceiling ─────────┴────────────┴─────────────┴──────
                 clear polycarbonate panel with a
                 small grille where the duct drops in
```

**Key move: use a remote inline fan, not a ceiling fan housing.**

A conventional ceiling-mount fan (Panasonic WhisperCeiling, Broan) is a big white box that has to sit *in* the ceiling plane — right in the middle of your clear panel, blocking the daylight you're trying to preserve, and needing structural support the polycarbonate can't provide.

An inline fan like the [Panasonic WhisperLine FV-20NLF1](https://www.amazon.com/Panasonic-FV-20NLF1-WhisperLine-Line-6-Inch/dp/B000EDUIX2) mounts remotely — in an adjacent joist bay or hung beside the shaft — and all that appears at the ceiling is a **small inlet grille**. Your clear panel stays almost entirely clear. Bonus: the motor noise is 8 feet away from your head instead of directly above it, and the FV-20NLF1's housing is insulated specifically to stop condensation and noise.

The FV-20NLF1 is 240 CFM, which is a lot for one bathroom. Options:
- Use it and put it on a speed control, or
- Use the smaller FV-10NLF1 (160 CFM) from the same [WhisperLine family](https://na.panasonic.com/us/home-and-building-solutions/ventilation-indoor-air-quality/ventilation-fans/whisperliner-remote-mount-line-fan-240-cfm), or
- Use a single-inlet install kit ([FV-NLF06G](https://www.amazon.com/Panasonic-FV-NLF06G-WhisperLine-6-Inch-Grille/dp/B00084ZQA2)) and accept the overkill — a bathroom under a skylight well has more volume than a normal bathroom, and oversizing here mostly costs you a little heat.

**Sizing sanity check:** the common rule is ~1 CFM per square foot of floor area, with **50 CFM the code minimum** for intermittent bathroom exhaust. A typical row-house bathroom (say 6' × 8' = 48 sq ft) wants 50–80 CFM nominal. Add capacity for the light well volume and duct length and 80–160 CFM is a reasonable target.

**Makeup air:** whatever CFM you exhaust has to come back in. Undercut the bathroom door ¾" or fit a door with a grille, or the fan will just stall against its own static pressure and you'll wonder why it's not clearing the mirror.

---

## 4. The Clear Ceiling Panel

**Material: polycarbonate, not acrylic.** Acrylic (Plexiglas) is cheaper and slightly clearer, but it's brittle, it crazes, and it's a worse actor in a fire. Polycarbonate (Lexan) is effectively unbreakable at this thickness and handles the temperature swing in a light well without complaint. Get the **UV-coated** version — it's sitting under direct sky all day and uncoated polycarbonate yellows.

- Home Depot stocks a [24" × 36" × 0.093" UV-coated Lexan sheet](https://www.homedepot.com/p/LEXAN-24-in-x-36-in-x-0-093-in-Shatter-Resistant-Clear-UV-Coated-Polycarbonate-Sheet-11N43101/333295898) (~$50) that covers a 24×32 opening with trim to spare.
- For a panel this size spanning unsupported, **0.093" is thin — it will sag and oil-can.** Either step up to ¼" ([24×36 available](https://www.amazon.com/Clear-Polycarbonate-Lexan-Sheet-24/dp/B06XH5T1G3)) or build a light frame (see below). ¼" is the safer call.

**Build it as a removable framed panel, not a glued-in sheet.** You will need to get back up there — to clear the chain, service the actuator, wipe condensation off the glass, or fish a wire. Design for that on day one.

Suggested construction:
1. Rip a **1× or ¾" plywood frame** to the inside dimensions of the ceiling opening, with a rabbet or a stop bead to receive the panel.
2. Drop the polycarbonate in, secure with a **removable stop bead** and screws — no adhesive, no silicone bead you'll have to cut out later.
3. Cut the grille hole with a hole saw at **slow speed with light pressure**; polycarbonate melts and grabs if you rush it. Leave the protective film on until the very end.
4. **Never drill polycarbonate tight to a fastener.** Oversize every screw hole by 1/16" so the sheet can move — it expands and contracts noticeably across a 32" span, and a pinned sheet will crack or bow.
5. Gasket the panel perimeter with closed-cell foam tape so the well is reasonably sealed from the bathroom. You want the fan pulling air through the *grille*, not around the edges of the panel.

**Aesthetic suggestion:** run a thin white or paint-matched trim ring around the grille so it reads as a deliberate fixture rather than a hole punched in a window. And consider putting the grille **off-center**, toward the shower end of the panel — it captures steam better at the source and looks less like a bullseye.

---

## 5. Motorizing the Skylight

Three paths, matching how much you want to disturb the existing unit.

### Path 1 — Add a chain actuator, keep the manual operator (least invasive)

Mount a **24V electric chain actuator** in the light well, attached to the curb liner, with its chain bracket fixed to the sash. The existing pull-chain operator stays in place as a manual backup — which is genuinely valuable the first time the electronics misbehave.

Typical units: [24V chain actuator, 300–800mm adjustable stroke](https://www.amazon.com/Electric-Actuator-Automatic-Casement-Skylights/dp/B0CGNM2T6N), roughly 10mm/s, with wall switch and remote. For a 24×32 top-hinged sash a **300–400mm stroke** gives a generous opening; more than that and you're fighting the geometry.

**Check the force rating against your sash weight before ordering.** A 24×32 glazed sash with an aluminum frame is plausibly 25–40 lb, and the actuator only sees a fraction of that depending on hinge geometry — but the cheap actuators are rated around 200–300N (45–67 lbf) and the [120V hardwired units advertise ~20 lb lift](https://www.amazon.com/Skylight-Control-Hard-Wired-Adjustable-Skylights/dp/B0DF2QG1MM). If your sash is heavy glass, size up or use two actuators. If the existing manual operator is stiff or gritty, **service or replace it first** — an actuator will happily destroy itself pushing against a seized hinge.

### Path 2 — Replace the manual operator with a factory motorized operator

If the existing operator turns out to be a Truth/AmesburyTruth unit, there may be a direct electric replacement in the same family. Worth a phone call to SWISCO with your part photo before you go the aftermarket-actuator route.

### Path 3 — Replace the whole skylight with a factory electric unit (most reliable, most expensive)

This is the "do it once" answer. A factory electric venting skylight arrives with a **motor, a rain sensor, an insect screen, and a warranty**, all engineered together.

- **VELUX VCE** (electric curb mount) or **VCS** (solar powered) — both include an [integrated rain sensor that closes the unit at the first sign of weather](https://www.solarskylights.com/velux-curb-mounted-electric-skylights/). VCS is attractive here because it's solar-charged and needs **no wiring run to the roof at all**. But remember from §1: **24 × 32 is not a VELUX size** — you'd be rebuilding the curb to 2230 (22½ × 30½ inside) or 2234.
- **If you want to keep the existing 24 × 32 curb**, go custom: Wasco's [CUSTOMEVMS-EML custom-size solar motorized venting skylight](https://www.iqskylights.com/wasco-customevms-eml-ultraseal-self-flashing-custom-size-solar-motorized-venting-glass-skylight.asp) is ordered to your exact rough opening and ships with the motor, solar panel, battery, rain sensor, and screen pre-installed. Expect a 3–4 week lead time and a signed shop-drawing approval.

**Whichever path: an insect screen is not optional.** An open skylight on a summer night in Philadelphia is an invitation.

---

## 6. Control Wiring — Three Tiers

The requirement is simple to state — *fan on ⇒ skylight open, fan off ⇒ skylight closed* — and has a few sharp edges in practice.

### Tier 1 — No smart home, relay logic only

```
120V switched hot (fan) ──► coil of a 120VAC DPDT relay
                                    │
24V PSU (always powered) ───► DPDT contacts ───► actuator + / −
                              (relay energized  = OPEN polarity)
                              (relay at rest    = CLOSE polarity)
```

The 24V supply is **always on**, and the fan's switched hot only drives the relay coil. That's what lets the skylight still close when the fan is off — if you powered the actuator from the fan circuit, it would lose power the instant you needed it to close.

Two things to get right:

- **Does your actuator cut its own motor at the end of travel?** Most have internal limit switches or stall detection, so holding polarity at the limit is harmless. **Verify this in the manual.** If it doesn't, add an off-delay timer relay so drive power is removed ~60 seconds after each movement, or you'll cook the motor.
- **The fan starts instantly; the skylight takes 30–60 seconds to open.** Under Path A (§2) that matters a lot. Add a **delay-on-make timer** so the fan waits for the skylight. Under the recommended Path B it's harmless — the fan is ducted and doesn't care.

### Tier 2 — Smart, and what I'd actually build

Use a **[Shelly 2PM Gen3 in cover mode](https://kb.shelly.cloud/knowledge-base/cover-mode)** to drive the actuator, plus a Shelly on the fan (or just read the fan's switch state), and tie them together in Home Assistant or in a local Shelly script.

Cover mode is purpose-built for exactly this motor pattern — two interlocked outputs, travel-time calibration, position tracking, and a [direction-reverse setting](https://kb.shelly.cloud/knowledge-base/shelly-plus-2-pm-device-smart-control-cover-mode) if it opens when it should close.

**Wiring caveat:** Shelly relay outputs are normally fed from the module's own line input. Switching a *separate* 24V supply requires the potential-free / dry-contact wiring configuration — **check your specific model's manual**, and if it isn't supported, have the Shelly drive an external DPDT relay instead. Alternatively, sidestep the issue entirely by buying a **[120V hardwired skylight motor](https://www.amazon.com/Skylight-Control-Hard-Wired-Adjustable-Skylights/dp/B0DF2QG1MM)** so the Shelly is switching line voltage, which is what it's designed for.

What the smart tier buys you that Tier 1 can't:
- **Sequencing** — open the skylight, *then* start the fan.
- **Run-on timer** — hold the skylight open 10 minutes after the fan stops to finish drying the room out.
- **Humidity trigger** — open and run on humidity, not just on the wall switch. This is the version you'll actually appreciate.
- **Unconditional close** — a nightly "close no matter what" automation, and a close-on-rain that overrides everything.
- **Failure visibility** — an alert when the cover doesn't reach its expected position, which is the difference between noticing a stuck skylight and finding out in the spring.

### Tier 3 — Factory skylight + gateway

If you went Path 3 in §5, a VELUX unit pairs to a **KLF 200** interface, which brings it into Home Assistant cleanly. Same automation logic as Tier 2, with far better hardware underneath.

### Rain sensor — mandatory on every tier

If the skylight can open by itself, it must be able to close by itself. Either buy an actuator with the sensor built in, or add a wired one such as the [Koncey 24V rain/wind sensor](https://www.amazon.com/Koncey-Automatic-Electric-Automatically-Windows/dp/B0DF1PKZ8F). Wire it so that **rain forces close and overrides the fan interlock**, not the other way around. Under Path B this is a non-event — the fan keeps working with the skylight shut. Under Path A, rain-close means the fan should also shut down, which is another argument for Path B.

---

## 7. Weather, Winter, and Condensation

- **Insulate the duct.** Any duct running through the unconditioned zone between ceiling and roof must be insulated, or it condenses internally and drips back down through your fan. This is explicitly required where the duct passes through unconditioned space, and it's the single most common cause of "my bathroom fan is leaking."
- **Slope the duct toward the roof cap** so any condensate that does form runs outward, not back into the fan.
- **Expect condensation on the skylight glazing itself** in winter regardless of what you do. If the existing unit is single-pane — very common in older row-house skylights — this will be dramatic. A modern double- or triple-glazed replacement essentially eliminates it, which is a real argument for Path 3 beyond just the automation.
- **Winter behavior:** consider a Home Assistant condition that skips the skylight-open step below some outdoor temperature (say 35°F) and just runs the fan. You lose the stack-vent boost and keep the heat. Trivial in Tier 2, impossible in Tier 1.
- **Air-seal the clear panel.** A leaky panel turns the light well into a chimney that bleeds conditioned air 24/7, fan or no fan.

---

## 8. Install Sequence

1. **Identify and inspect** (§1). Photograph the nameplate and the operator. Get the curb dimensions, inside and out.
2. **Exercise the existing mechanism.** Open and close it by hand a dozen times. Note any binding, grit, or slop. Service the hinges and operator now, before anything gets motorized.
3. **Decide Path A / B / C** (§2). This locks in everything downstream.
4. **Run power to the well** — a switched leg for the fan and a constant leg for the 24V supply and actuator. Put the PSU in a proper junction box, not loose in the joist bay.
5. **Cut and flash the roof penetration** for the duct and set the roof cap. This is the one step where hiring a roofer is usually the right call in a row house — flat and low-slope roofs are unforgiving, and a bad flashing detail on a Philly rowhome roof is an expensive lesson.
6. **Mount the inline fan** in the adjacent joist bay. Suspend it on the included brackets; don't hard-mount it to framing if you can avoid it — the isolation is worth it acoustically.
7. **Duct it up**, insulated, sloped to the cap.
8. **Mount the actuator** in the well. Attach to the curb liner, bracket to the sash, set the stroke, and cycle it **by hand power first** at low duty to confirm the geometry doesn't bind at either end of travel.
9. **Wire the interlock** (§6). Test with the panel still off so you can see everything move.
10. **Build and fit the clear panel** (§4) last, with the grille cut to match the duct drop.
11. **Commission it:** run a hot shower with the door closed and watch. Mirror should clear in a few minutes. Check the duct for condensation after the first cold night. Trigger the rain sensor with a wet finger and confirm the skylight closes.

---

## 9. Things That Will Bite You

| Symptom | Likely cause |
|---|---|
| Skylight opens but fan barely moves air | No makeup air — undercut the bathroom door |
| Water dripping from the ceiling grille | Uninsulated duct condensing, or duct sloped back toward the fan |
| Actuator buzzes and won't complete travel | Underspecced force rating, or a binding hinge/operator you didn't service first |
| Skylight opens when it should close | Reversed polarity — flip the leads, or use the Shelly cover-mode reverse setting |
| Clear panel bows or cracks near screws | Screw holes not oversized for thermal movement, or sheet too thin for the span |
| Panel fogs on the inside | Panel not sealed to the ceiling; humid air migrating around its edges into the well |
| Actuator motor eventually burns out | Drive polarity held continuously at the travel limit on a unit with no internal cutoff |
| Everything works, room still steams up | Fan undersized for the added light-well volume — this space is bigger than it looks |

---

## 10. Code Notes

- **[IRC M1507.2](https://www.jaspector.com/codes/irc-2018/ch15-exhaust-systems/bathroom-exhaust-fan-outdoor-termination-irc-2018/)**: bathroom exhaust must discharge directly outdoors; it may not discharge into an attic, crawl space, or other interior area. A light well counts as interior. This is the whole reason for Path B.
- **50 CFM** is the standard intermittent-exhaust minimum for a bathroom (20 CFM if run continuously).
- **Ducts through unconditioned space must be insulated** to prevent condensation.
- **Philadelphia L&I**: cutting a new roof penetration and running new branch circuits generally means permits. Check with [Philadelphia Licenses & Inspections](https://www.phila.gov/departments/department-of-licenses-and-inspections/) before you start, especially if your block is in a historic district — row-house roof work sometimes carries extra review.
- **Party walls**: in a row house your roof penetration needs to respect the party-wall setback. Don't site the roof cap near the property line without checking.

---

## Summary Recommendation

Build **Path B, Tier 2, §5 Path 1**:

- Keep the existing skylight; add a 24V chain actuator alongside the working manual pull-chain.
- Remote inline fan in an adjacent joist bay, ducted through an insulated 6"→4" run to its own roof cap.
- Clear ¼" UV-coated polycarbonate panel in a removable wood frame, small off-center inlet grille.
- Shelly 2PM in cover mode + Home Assistant: skylight opens, then fan starts; fan stops, skylight holds 10 minutes, then closes; rain and cold-weather overrides.

That gets you exactly the behavior you described, keeps the daylight, keeps a manual fallback, and — critically — the bathroom still ventilates correctly on the day the actuator dies.

See `skylight_vent_fan_purchase_list.md` for parts and links.
