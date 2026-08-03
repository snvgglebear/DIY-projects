# Skylight-Linked Bathroom Vent Fan — Design Guide

**Version:** 2.0
**Date:** August 2026
**Subject:** 24" × 32" glazed roof ventilator with a chain-operated side louver, over a bathroom light well (Philadelphia row house)
**Goal:** Mount an exhaust fan in the ceiling opening below the skylight, fill the rest of the opening with clear plastic so daylight still comes through, and have the louver open automatically when the fan runs and close when it stops.
**Companion Document:** `skylight_vent_fan_purchase_list.md`

> **v2.0 changelog** — Revised after learning the actual configuration: a **fixed** frosted glass top on a vertical metal box, with a **hinged metal louver in the vertical side wall** as the operable element, and confirmed by testing that rain does not enter when open. This is a different (and considerably more favorable) machine than the hinged-sash skylight v1.0 assumed. Three consequences:
> - **The rain sensor is gone.** It was the most fragile part of the design.
> - **The roof penetration is probably gone too** — the vertical metal box gives you an above-roof wall to terminate through, which is much easier and safer than cutting the roof.
> - **The control system collapses to one part.** A spring-return damper actuator does exactly what you asked for with no relay logic at all.

---

## Table of Contents

1. [What You Actually Have](#1-what-you-actually-have)
2. [Why This Configuration Is Good News](#2-why-this-configuration-is-good-news)
3. [The Design Question: Where Does the Air Terminate?](#3-the-design-question-where-does-the-air-terminate)
4. [Recommended Design](#4-recommended-design)
5. [The Clear Ceiling Panel](#5-the-clear-ceiling-panel)
6. [Motorizing the Louver](#6-motorizing-the-louver)
7. [Control Wiring](#7-control-wiring)
8. [One Thing to Check: The Glass](#8-one-thing-to-check-the-glass)
9. [Condensation and Winter](#9-condensation-and-winter)
10. [Install Sequence](#10-install-sequence)
11. [Things That Will Bite You](#11-things-that-will-bite-you)
12. [Code Notes](#12-code-notes)

---

## 1. What You Actually Have

### The honest answer on identification

**I don't think you're going to find a brand and model number for this, and it probably doesn't matter.**

What you're describing is a **glazed roof ventilator** — a vertical sheet-metal box (riser/curb) standing proud of the roof, capped by a *fixed* sloped light-transmitting panel, with a hinged louver or damper in one vertical face, chain-operated from the room below. In commercial ventilation terms this is a small **louvered penthouse gravity ventilator** with a glazed roof instead of a solid one — the same idea Greenheck and [United Enertech](https://unitedenertech.com/louvered-penthouses) still build at industrial scale, where louvered penthouses serve "as gravity ventilators, fan discharge caps, fresh air intake caps, [and] pressure relief ventilators."

In Philadelphia and Baltimore row houses this was usually **local sheet-metal shop work**, not a catalog product. Galvanized or terne-coated steel, soldered seams, wire or obscure glass, fabricated to fit whatever hole was in that particular roof. That's why it's 24 × 32 — a non-standard size that matches no manufacturer's catalog. The [Philadelphia Rowhouse Manual](https://www.phila.gov/media/20190521124726/Philadelphia_Rowhouse_Manual.pdf) is the best general reference on this housing stock's roof assemblies.

**This is why my v1.0 guess (Ventarama) was wrong.** Ventarama and the VELUX/CrystaLite family all use a *hinged glazed sash* — the glass itself lifts. Yours doesn't; the glass is fixed and a separate metal louver does the work. Different machine entirely.

### What's still true from v1.0

**The two chains are almost certainly one continuous loop over a sprocket or pulley.** Pull one strand, the sprocket rotates and drives the louver's linkage arm open; pull the other, it drives closed. Confirm by grabbing both strands and pulling in opposite directions — if they fight each other, it's a loop. That detail matters for §6, because a rotating shaft is exactly what a damper actuator wants to grab.

### What to look at before you buy anything

- **The louver's operating shaft or linkage arm.** This is the single most important thing to photograph. Is there a round shaft the chain sprocket sits on? What diameter? Or is it a flat crank arm on a pivot pin? This determines how the actuator couples in (§6) and it's the one genuinely fiddly part of the whole project.
- **The louver's free area when fully open** — roughly, height × width of the opening. This is your exhaust aperture and it wants to be comfortably bigger than the duct you'll run.
- **Whether the metal box has any free wall area** on a face *other* than the louvered one. That's where the wall cap goes (§4).
- **Condition of the metal and the soldered seams.** Hundred-year-old galvanized has a service life, and if the box is rusting through, that's a separate conversation from this project.
- **The glass.** See §8 — worth a look while you're up there.

---

## 2. Why This Configuration Is Good News

Genuinely — this is a better starting point than the hinged-sash skylight I assumed in v1.0.

| | Hinged glazed sash (v1.0 assumption) | **Your louvered ventilator** |
|---|---|---|
| Rain when open | Real risk; rain sensor mandatory | **Not an issue** — you've tested it, and the geometry explains why |
| What moves | Heavy glazed sash, 25–40 lb | Light metal louver on a shaft |
| Weather seal risk | Moving glazing = a leak path that ages | Glass is **fixed** — nothing to fail |
| Roof penetration for duct | New cut + flashing, hire a roofer | **Vertical metal wall already above the roof** |
| Actuator sizing | Marginal; cheap units underpowered | Trivial — it's a damper |
| Fail-safe on power loss | Stays wherever it was | **Spring closes it** |

**On the rain:** your testing matches the physics. The louver is in a *vertical* face, shielded by the overhanging sloped glass above it, and louver blades are angled specifically to shed water outward. That's the same reason a louvered penthouse works as a weather-protected intake on a commercial roof. So — **no rain sensor.** In v1.0 that sensor was the most failure-prone component in the design and the reason for half the control complexity. It's simply gone.

Two small caveats I'd note and then drop, since you've tested it: hard wind-driven rain and blowing snow behave differently from ordinary rain, and snow can pack into a louver and freeze it shut. Neither is a reason to change the design. If the louver ever *does* freeze shut, §3's recommendation means your fan keeps working anyway.

---

## 3. The Design Question: Where Does the Air Terminate?

Same question as v1.0, better answer available.

**The obvious version:** fan blows up into the light well and the metal box, louver opens, air leaves. Simple, no ductwork.

**Why I'd still avoid it:** [IRC M1507.2](https://www.jaspector.com/codes/irc-2024/ch15-exhaust-systems/bathroom-exhaust-duct-requirements-irc-2024/) requires exhaust to discharge directly outdoors and prohibits discharge into interior cavities; a light well and a closed metal box qualify. But the practical argument is stronger than the code argument here:

**A sealed sheet-metal box is the worst possible condensing surface.** Uninsulated metal at outdoor temperature, shower steam pumped into it, glass on top. If the louver is shut — actuator dead, linkage frozen, breaker off — every bit of that moisture condenses on the inside of the box and runs back down the light well onto your new clear panel. You'd find out about it as drips, in January.

**But you have an option v1.0 didn't:** the box's vertical wall is already an exterior wall standing above the roof. That's a *far* better place to terminate a duct than a roof cut.

| | **Path A — Box as plenum** | **Path B — Duct to a wall cap (recommended)** |
|---|---|---|
| Fan exhausts into | The metal box, via the louver | A dedicated wall cap in the box's side wall |
| Louver motorized? | Required — sole exit | Bonus: a big passive boost vent |
| Code | Fails M1507.2 as written | Compliant |
| Fails safe? | **No** — louver shut = steam in a cold metal box | Yes — fan works regardless of the louver |
| Roof work | None | **None** — vertical penetration, not a roof cut |
| Cost | $ | $$ |

**Path B costs you a hole saw and a $20 wall cap, and it deletes the entire failure mode.** No roofer, no flashing, no roof cement. Cutting a 4" hole in a vertical metal wall above the roofline is about as low-risk as exterior penetrations get — it's a sheet-metal job, sealed with a gasket and a bead of sealant, with gravity working for you instead of against you.

**Siting the wall cap:** put it on a **different face than the louver**. If you terminate on the same face, exhaust leaving the cap gets pulled straight back in through the open louver next to it and you've built a recirculation loop. Opposite face is ideal; adjacent face is fine.

---

## 4. Recommended Design

```
                    ┌── fixed frosted glass (sloped, stays put) ──┐
                   ╱                                              │
                  ╱   ┌──────────────────────────────────────┐    │
    metal box    ╱    │                                      │    │
    above roof  ╱     │   Belimo LF24 spring-return actuator │    │
               │      │   on the louver shaft                │    │
   ┌───────────┤      │                                      ├────┤
   │ wall cap ◄├──────┼──── 4" duct up from the fan          │    │  ← louver
   │ (far face)│      │                                      │    │    (this face)
   └───────────┤      └──────────────────────────────────────┘    │
   ═════════════╧══════════════════════════════════════════════╧═══  roof line
                    │                                          │
                    │           light well / shaft             │
                    │  ┌────────────────────┐                  │
                    │  │  inline fan in an  │──── 6" duct ─────┘
                    │  │  adjacent joist bay│
                    │  └─────────┬──────────┘
   ceiling ─────────┴────────────┴──────────────────────────────────
                 clear polycarbonate panel with a small
                 off-center inlet grille
```

**Use a remote inline fan, not a ceiling-mount housing.** A conventional ceiling fan is a big white box that has to sit *in* the ceiling plane — right in the middle of your clear panel, blocking the daylight you're trying to preserve, and needing structural support the polycarbonate can't provide. An inline fan like the [Panasonic WhisperLine FV-20NLF1](https://www.amazon.com/Panasonic-FV-20NLF1-WhisperLine-Line-6-Inch/dp/B000EDUIX2) mounts in an adjacent joist bay; all that shows at the ceiling is a small grille. Its housing is insulated specifically against condensation and noise, and the motor ends up 8 feet from your head instead of directly above it.

**Sizing:** 50 CFM is the code minimum for intermittent bathroom exhaust; ~1 CFM per square foot of floor is the working rule. A light well adds volume the rule doesn't count, so **80–160 CFM** is the target. The FV-20NLF1 is 240 CFM — use a speed control, or step down to the FV-10NLF1 (160 CFM) in the same family.

**Makeup air:** undercut the bathroom door ¾" or fit a grille. Whatever you exhaust has to come back in, or the fan stalls against its own static pressure and the mirror stays foggy.

**What the louver now does for you:** it's a large-aperture passive stack vent, opened on demand. Open louver plus a running fan clears a steamy bathroom dramatically faster than the fan alone, and on a summer night you can open it manually with the chains — as you presumably already do — and get free stack ventilation with no fan at all. Keeping the manual chain operation intact is worth doing for exactly this reason.

---

## 5. The Clear Ceiling Panel

Unchanged from v1.0, and it works better here — since the glass above is fixed, the daylight path is permanent and your panel never interacts with anything that moves.

**Material: polycarbonate, not acrylic.** Acrylic is cheaper and marginally clearer but brittle, prone to crazing, and worse in a fire. Polycarbonate (Lexan) is effectively unbreakable at this thickness and shrugs off the temperature swing in a light well. Get the **UV-coated** grade — it sits under open sky all day and uncoated polycarbonate yellows.

- Home Depot stocks a [24" × 36" × 0.093" UV-coated Lexan sheet](https://www.homedepot.com/p/LEXAN-24-in-x-36-in-x-0-093-in-Shatter-Resistant-Clear-UV-Coated-Polycarbonate-Sheet-11N43101/333295898) (~$50) that covers a 24 × 32 opening with trim to spare.
- **0.093" is thin for an unsupported 32" span** — it will sag and oil-can. Step up to ¼" ([24 × 36 available](https://www.amazon.com/Clear-Polycarbonate-Lexan-Sheet-24/dp/B06XH5T1G3)) or add a frame member. ¼" is the safer call.

**Build it as a removable framed panel, not a glued-in sheet.** You will need to get back up there — to service the actuator, clear the chain, or fish a wire. Design for that on day one.

1. Rip a **1× or ¾" plywood frame** to the opening's inside dimensions, with a rabbet or stop bead to receive the panel.
2. Drop the sheet in and secure with a **removable stop bead** and screws. No adhesive, no silicone bead you'll be cutting out in three years.
3. Cut the grille hole with a hole saw at **slow speed, light pressure** — polycarbonate melts and grabs if you rush it. Leave the protective film on until the end.
4. **Oversize every screw hole by 1/16".** The sheet moves noticeably across a 32" span, and a pinned sheet cracks or bows.
5. Gasket the perimeter with closed-cell foam tape. You want the fan pulling through the *grille*, not around the panel edges.

**Design suggestions:** run a thin white or paint-matched trim ring around the grille so it reads as a deliberate fixture rather than a punched hole. Put the grille **off-center, toward the shower end** — it catches steam at the source and looks less like a bullseye. And leave the chains reachable: they'll pass through or beside the panel, so plan a grommeted slot rather than discovering the problem after the panel is cut.

---

## 6. Motorizing the Louver

This is where the new information pays off most. You're no longer trying to lift a heavy glazed sash — you're rotating a light metal damper. That's a solved problem with an off-the-shelf part.

### The right part: a spring-return damper actuator

A **[Belimo LF24](https://www.belimo.com/us/shop/en_US/p?code=LF24+US)** is a 24VAC/DC fail-safe damper actuator: 35 in-lb torque, 95° of rotation, mounting directly to a **3/8"–1/2" shaft** with a universal clamp, with a graduated position indicator.

The important property is **spring return**. Powered, it drives the damper open and holds it. Cut the power and a spring drives it closed — mechanically, with no electronics involved.

**That is precisely your stated requirement, implemented in one component.** Fan on, louver opens. Fan off, louver closes. No polarity-reversing relay, no timers, no always-on supply, no controller. And it fails closed on a power outage, a blown fuse, or a dead actuator — the safe direction.

Two variants worth knowing:
- **LF24** — plain on/off, the base part.
- **[LF24-S](https://www.fwwebb.com/product/Belimo/LF-Damper-Actuator/LF24-SUS/256371)** — adds a built-in SPDT auxiliary switch, adjustable anywhere from 0° to 95°. Worth the small upcharge: it gives you a real electrical signal that the louver actually reached open, which you can wire to an indicator lamp or a smart input. That's the difference between knowing your louver is stuck and finding out in the spring.

**Torque check:** 35 in-lb suits a small, free-moving damper. If your louver is stiff, rusty, or larger than expected, service it first, and if it's still heavy, step up to a Belimo AF24 (same family, ~180 in-lb). **Free the mechanism before you motorize it** — an actuator will happily grind itself to death against a seized pivot.

### The fiddly part: coupling to the linkage

The LF24 clamps to a round shaft. Your hundred-year-old sheet-metal louver may or may not present one.

- **If the chain sprocket rides on a round shaft of 3/8"–1/2"** — ideal. The actuator clamps straight on, and if there's enough shaft length you can leave the sprocket and chains in place for manual override. Note that with a spring-return actuator attached, manual pulls now work against the spring, so the chains become an emergency override rather than everyday operation.
- **If it's a flat crank arm on a pivot pin** — you'll need a crank-arm/linkage kit. Belimo sells universal crank arm and ball-joint linkage kits for exactly this retrofit case. Bring your photos to a controls supplier ([Kele](https://www.kele.com/product/actuators-and-dampers/spring-return/belimo/lf24), [Jackson Systems](https://jacksonsystems.com/product/belimo-lf24-us-2-position-actuator-spring-return/), or a local HVAC controls house) and let them spec the kit against what you actually have.
- **If the linkage is hopeless** — fall back position: leave the louver fully manual and put a **motorized damper in the duct run** instead. You lose the automatic louver, but Path B means the fan still works perfectly. This is a legitimate outcome, not a failure.

**Mounting:** the actuator needs a solid bracket inside the metal box, and it's now living in an unconditioned, occasionally damp outdoor enclosure. Mount it as high and as sheltered as the geometry allows, run the low-voltage wiring in with a drip loop, and check it once a year.

---

## 7. Control Wiring

With a spring-return actuator, this is the entire system:

```
   fan switch (120V switched hot) ──┬──► inline fan
                                    │
                                    └──► 24VAC transformer ──► Belimo LF24
                                         (40VA, Class 2)        (spring closes
                                                                 on loss of power)
```

That's it. One transformer, two wires to the actuator. The transformer's primary lands on the fan's switched hot, so the actuator is energized exactly when the fan is.

**Notes:**
- Use a **Class 2, 40VA 24VAC transformer** — comfortably above the LF24's draw, and Class 2 keeps the low-voltage wiring in easy territory.
- **The louver takes a moment to open; the fan starts instantly.** Under Path B this is harmless — the fan is ducted to its own wall cap and doesn't care about the louver at all. (Under Path A you'd need a delay-on-make timer; another reason to build Path B.)
- Put the transformer in a proper junction box in the joist bay, not loose in the shaft.
- If you fit the **LF24-S**, run its aux switch down to a small indicator LED by the light switch. Cheap, and it tells you at a glance whether the louver is actually opening.

### Optional: smart controls

You don't need them — the requirement is fully met above. But if you already run Home Assistant, a [Shelly Plus 1 or 1PM](https://us.shelly.com/) on the fan circuit gets you, with no change to the wiring above:

- **Humidity triggering** — fan and louver run on humidity rather than only on the wall switch. This is the upgrade you'd actually notice day to day.
- **Run-on timer** — hold both for 10 minutes after the shower to finish drying the room.
- **Cold-weather skip** — below ~35°F, run the fan without opening the louver. Path B makes that possible; the fan is fully independent of the louver.
- **Failure alerting** — pair with the LF24-S aux switch to alert when the louver doesn't reach open.

The Shelly switches line voltage on the fan circuit, which is what it's designed for — none of the 24V dry-contact complications v1.0 had to work around.

---

## 8. One Thing to Check: The Glass

While you're up there, look closely at that frosted panel. **If it's wired glass** — obscure glass with a steel mesh embedded in it, very common in this vintage of rooftop sheet-metal work — it's worth knowing that **wired glass fails the modern safety-glazing test and isn't permitted in overhead applications** except within fire-rated assemblies. Under current code, [sloped and overhead glazing](https://up.codes/s/sloped-glazing-and-skylights) has to be laminated glass, or glass with a retention film, or be protected by a screen below it, so that broken glass stays put instead of falling on whoever's below.

The counterintuitive part: the embedded wire makes wired glass *weaker* than plain annealed glass, not stronger, and when it breaks it leaves shards hanging on wire rather than falling clear.

I'm flagging it, not sounding an alarm. It's existing construction, it's been there for decades, nobody's requiring you to change it, and it's not a reason to postpone this project. But two things follow:

1. **Your clear polycarbonate panel is doing incidental safety duty.** It sits below the well and would catch anything that came down. That's a genuine (if accidental) benefit of the design — one more argument for ¼" over 0.093".
2. **If you ever do replace the glazing**, that's the moment to go to laminated. Not a today problem.

---

## 9. Condensation and Winter

- **Insulate the duct.** Any duct in the unconditioned zone between ceiling and wall cap must be insulated or it condenses internally and drips back down through the fan. This is the single most common cause of "my bathroom fan is leaking."
- **Slope the duct toward the wall cap** so any condensate that does form runs outward.
- **Expect condensation on the underside of the fixed glass** in winter regardless. It's single-pane glass on a metal box; that's just what it does. Path B means you're no longer *adding* shower steam to it, which helps a lot compared with Path A.
- **Air-seal the clear panel.** A leaky panel turns the light well into a chimney that bleeds conditioned air 24/7, fan or no fan. Between the panel gasket and a closed louver, the shaft should be reasonably sealed when nothing's running.
- **Winter louver behavior:** with Path B you can simply choose not to open the louver when it's cold (see §7, smart controls) and lose nothing but the boost.

---

## 10. Install Sequence

1. **Get up there and photograph everything** (§1) — especially the louver's shaft or linkage arm, and the free wall area on the non-louvered faces. Measure the shaft diameter.
2. **Exercise the louver by hand.** Open and close it a dozen times with the chains. Free up, clean, and lubricate the pivots. **Do this before ordering an actuator** — it determines the torque you need and whether the LF24 is enough.
3. **Take your photos to a controls supplier** and have them spec the actuator plus whatever crank-arm or linkage kit your louver needs. This is the one part worth a human's eyes on.
4. **Cut and fit the wall cap** in the metal box, on a face away from the louver. Hole saw or step drill, deburr, gasket, sealant, screw the flange, seal the perimeter.
5. **Run power to the well** — a switched leg for the fan and the transformer.
6. **Mount the inline fan** in the adjacent joist bay on its suspension brackets; don't hard-mount to framing if you can avoid it — the isolation is worth it acoustically.
7. **Duct it up**, insulated, sloped toward the cap.
8. **Mount the actuator** on the louver shaft. Set the rotation limits, then **power-cycle it a dozen times with the panel still off** so you can watch the full travel and confirm nothing binds at either end.
9. **Wire the transformer** to the fan's switched hot. Verify with a multimeter before closing anything up.
10. **Build and fit the clear panel** last, with the grille cut to match the duct drop and a grommeted slot for the chains.
11. **Commission it:** run a hot shower with the door closed. Mirror should clear in a few minutes. Kill the breaker mid-cycle and confirm the louver springs shut. Check the duct for condensation after the first cold night.

---

## 11. Things That Will Bite You

| Symptom | Likely cause |
|---|---|
| Louver opens but fan barely moves air | No makeup air — undercut the bathroom door |
| Exhaust smell or steam coming back in the louver | Wall cap sited on the same face as the louver — short-circuit recirculation |
| Water dripping from the ceiling grille | Uninsulated duct condensing, or duct sloped back toward the fan |
| Actuator buzzes and won't reach full open | Under-torqued for a stiff louver, or a pivot you didn't free up first |
| Louver doesn't close on power-off | Spring return defeated by a binding linkage — the spring only has so much authority |
| Clear panel bows or cracks near screws | Screw holes not oversized for thermal movement, or sheet too thin for the span |
| Panel fogs on the inside | Panel not sealed to the ceiling; humid air migrating around the edges into the well |
| Louver frozen shut after a snow | Packed snow in the blades. Path B means the fan doesn't care — that's the point |
| Everything works, room still steams up | Fan undersized for the added light-well volume — the space is bigger than it looks |

---

## 12. Code Notes

- **[IRC M1507.2](https://www.jaspector.com/codes/irc-2018/ch15-exhaust-systems/bathroom-exhaust-fan-outdoor-termination-irc-2018/)**: bathroom exhaust must discharge directly outdoors, never into an attic, crawl space, or other interior area. A light well and a closed metal box count as interior. This is the entire argument for Path B.
- **50 CFM** intermittent (or 20 CFM continuous) is the standard bathroom minimum.
- **Ducts in unconditioned space must be insulated** against condensation.
- **[Overhead glazing](https://up.codes/s/sloped-glazing-and-skylights)** must be laminated, filmed, or screened — see §8. Applies to replacement, not to leaving existing construction alone.
- **Philadelphia L&I**: new branch circuits generally mean a permit. Check with [Licenses & Inspections](https://www.phila.gov/departments/department-of-licenses-and-inspections/), particularly if your block is in a historic district — visible rooftop changes can draw extra review, though a small wall cap on an existing ventilator is about as low-profile as it gets. **Notably, Path B involves no new roof penetration**, which is the part that usually triggers roofing scope and a contractor.

---

## Summary Recommendation

- **Keep the ventilator.** It's well suited to this job — weather-protected vertical louver, fixed glazing with no leak path, and a free above-roof wall to terminate through. Nothing here argues for replacing it.
- **Path B:** remote inline fan in an adjacent joist bay, insulated duct up through the shaft, terminating at a **wall cap cut into a non-louvered face of the metal box**. No roof work.
- **Belimo LF24-S spring-return damper actuator** on the louver shaft, powered from a 24VAC transformer on the fan's switched hot. Fan on → louver opens. Fan off → spring closes it. That's the whole control system, and it fails closed.
- **¼" UV-coated polycarbonate panel** in a removable wood frame, small off-center grille, grommeted slot for the chains.
- **Keep the manual chains working** for summer stack ventilation and emergency override.

Estimated parts cost: **$550–850**, and — because the wall-cap termination replaces the roof penetration — plausibly **no contractor at all** if you're comfortable with a hole saw, a duct run, and a 120V circuit.

See `skylight_vent_fan_purchase_list.md` for parts and links.
