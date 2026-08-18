# Skylight-Linked Bathroom Vent Fan — Design Guide

**Version:** 2.1
**Date:** August 2026
**Subject:** 24" × 32" glazed roof ventilator with a chain-operated side louver, over a bathroom light well (Philadelphia row house)
**Goal:** Mount an exhaust fan in the ceiling opening below the skylight, fill the rest of the opening with clear plastic so daylight still comes through, and link the louver's state to the fan.
**Companion Document:** `skylight_vent_fan_purchase_list.md`

> **v2.1 changelog** — Added §7, **sensing the louver instead of moving it**. Keeping the louver manual and adding a position sensor is a legitimate alternative to motorizing, not a lesser one: it removes the project's single most uncertain component (the actuator-to-linkage coupling), costs about a sixth as much, and — used as a fan interlock — rescues the no-ductwork Path A that §3 otherwise argues against. §12's recommendation now presents this as a staged build rather than a straight either/or.
>
> **v2.0 changelog** — Revised after learning the actual configuration: a **fixed** frosted glass top on a vertical metal box, with a **hinged metal louver in the vertical side wall** as the operable element, and confirmed by testing that rain does not enter when open. This is a different (and considerably more favorable) machine than the hinged-sash skylight v1.0 assumed. Three consequences: the rain sensor is gone; the roof penetration is probably gone too; and the control system collapses to one part.

---

## Table of Contents

1. [What You Actually Have](#1-what-you-actually-have)
2. [Why This Configuration Is Good News](#2-why-this-configuration-is-good-news)
3. [The Design Question: Where Does the Air Terminate?](#3-the-design-question-where-does-the-air-terminate)
4. [Recommended Design](#4-recommended-design)
5. [The Clear Ceiling Panel](#5-the-clear-ceiling-panel)
6. [Option 1: Motorizing the Louver](#6-option-1-motorizing-the-louver)
7. [Option 2: Sensing the Louver Instead of Moving It](#7-option-2-sensing-the-louver-instead-of-moving-it)
8. [Control Wiring](#8-control-wiring)
9. [One Thing to Check: The Glass](#9-one-thing-to-check-the-glass)
10. [Condensation and Winter](#10-condensation-and-winter)
11. [Install Sequence](#11-install-sequence)
12. [Summary Recommendation](#12-summary-recommendation)
13. [Things That Will Bite You](#13-things-that-will-bite-you)
14. [Code Notes](#14-code-notes)

---

## 1. What You Actually Have

### The honest answer on identification

**I don't think you're going to find a brand and model number for this, and it probably doesn't matter.**

What you're describing is a **glazed roof ventilator** — a vertical sheet-metal box (riser/curb) standing proud of the roof, capped by a *fixed* sloped light-transmitting panel, with a hinged louver or baffle in one vertical face, chain-operated from the room below. In commercial ventilation terms this is a small **louvered penthouse gravity ventilator** with a glazed roof instead of a solid one — the same idea Greenheck and [United Enertech](https://unitedenertech.com/louvered-penthouses) still build at industrial scale, where louvered penthouses serve "as gravity ventilators, fan discharge caps, fresh air intake caps, [and] pressure relief ventilators."

In Philadelphia and Baltimore row houses this was usually **local sheet-metal shop work**, not a catalog product. Galvanized or terne-coated steel, soldered seams, wire or obscure glass, fabricated to fit whatever hole was in that particular roof. That's why it's 24 × 32 — a non-standard size matching no manufacturer's catalog. The [Philadelphia Rowhouse Manual](https://www.phila.gov/media/20190521124726/Philadelphia_Rowhouse_Manual.pdf) is the best general reference on this housing stock's roof assemblies.

**This is why my v1.0 guess (Ventarama) was wrong.** Ventarama and the VELUX/CrystaLite family all use a *hinged glazed sash* — the glass itself lifts. Yours doesn't; the glass is fixed and a separate metal louver does the work. Different machine entirely.

### What's still true from v1.0

**The two chains are almost certainly one continuous loop over a sprocket or pulley.** Pull one strand, the sprocket rotates and drives the louver's linkage arm open; pull the other, it drives closed. Confirm by grabbing both strands and pulling in opposite directions — if they fight each other, it's a loop. That detail matters for both §6 and §7: a rotating shaft is what an actuator wants to grab, and a moving arm is what a sensor wants to watch.

### What to look at before you buy anything

- **The louver's operating shaft or linkage arm.** The single most important thing to photograph. Is there a round shaft the chain sprocket rides on, and what diameter? Or is it a flat crank arm on a pivot pin? This determines the §6 coupling and it's the one genuinely fiddly part of that path. **§7 sidesteps this question entirely** — worth knowing before you spend a Saturday on it.
- **The louver's free area when fully open** — roughly height × width of the opening. This is your exhaust aperture; it wants to be comfortably bigger than any duct you run.
- **Whether the box has free wall area on a face other than the louvered one.** That's where the wall cap goes (§4).
- **Condition of the metal and the soldered seams.** Century-old galvanized has a service life. If the box is rusting through, that's a separate conversation.
- **The glass.** See §9 — worth a look while you're up there.

---

## 2. Why This Configuration Is Good News

Genuinely — a better starting point than the hinged-sash skylight v1.0 assumed.

| | Hinged glazed sash (v1.0 assumption) | **Your louvered ventilator** |
|---|---|---|
| Rain when open | Real risk; rain sensor mandatory | **Not an issue** — you've tested it, and the geometry explains why |
| What moves | Heavy glazed sash, 25–40 lb | Light metal louver on a shaft |
| Weather seal risk | Moving glazing = a leak path that ages | Glass is **fixed** — nothing to fail |
| Roof penetration for duct | New cut + flashing, hire a roofer | **Vertical metal wall already above the roof** |
| Actuator sizing | Marginal; cheap units underpowered | Trivial — it's a damper |
| Fail-safe on power loss | Stays wherever it was | **Spring closes it** (§6), or sensor blocks the fan (§7) |

**On the rain:** your testing matches the physics. The louver sits in a *vertical* face, shielded by the overhanging sloped glass, and louver blades are angled specifically to shed water outward — the same reason a louvered penthouse works as a weather-protected intake on a commercial roof. So **no rain sensor.** In v1.0 that was the most failure-prone component and the reason for half the control complexity. It's simply gone.

Two caveats I'll note once and drop, since you've tested it: hard wind-driven rain and blowing snow behave differently from ordinary rain, and snow can pack a louver and freeze it shut. Neither changes the design. If it ever *does* freeze shut, §3's recommendation means the fan keeps working anyway.

---

## 3. The Design Question: Where Does the Air Terminate?

**The obvious version:** fan blows up into the light well and the metal box, louver opens, air leaves. Simple, no ductwork.

**Why I'd normally avoid it:** [IRC M1507.2](https://www.jaspector.com/codes/irc-2024/ch15-exhaust-systems/bathroom-exhaust-duct-requirements-irc-2024/) requires exhaust to discharge directly outdoors and prohibits discharge into interior cavities; a light well and a closed metal box qualify. But the practical argument is stronger than the code argument:

**A sealed sheet-metal box is the worst possible condensing surface.** Uninsulated metal at outdoor temperature, shower steam pumped into it, glass on top. If the louver is shut — actuator dead, linkage frozen, breaker off, or you simply forgot — every bit of that moisture condenses on the inside of the box and runs back down the light well onto your new clear panel. You'd find out about it as drips, in January.

**But you have an option v1.0 didn't:** the box's vertical wall is already an exterior wall standing above the roof. That's a far better place to terminate a duct than a roof cut.

| | **Path A — Box as plenum** | **Path B — Duct to a wall cap (recommended)** |
|---|---|---|
| Fan exhausts into | The metal box, via the louver | A dedicated wall cap in the box's side wall |
| Louver's role | Sole exit — load-bearing | Bonus: a big passive boost vent |
| Code | Fails M1507.2 as written | Compliant |
| Fails safe? | **Only with a §7 sensor interlock** | Yes — fan works regardless of the louver |
| Roof work | None | **None** — vertical penetration, not a roof cut |
| Cost | $ | $$ |

**Path B costs a hole saw and a $30 wall cap, and it deletes the failure mode entirely.** No roofer, no flashing, no roof cement. Cutting a 4" hole in a vertical metal wall above the roofline is about as low-risk as exterior penetrations get — sheet-metal work, sealed with a gasket and a bead of sealant, with gravity working for you rather than against you.

**Siting the wall cap:** put it on a **different face than the louver**. Terminate on the same face and exhaust leaving the cap gets pulled straight back in through the open louver beside it — you've built a recirculation loop. Opposite face ideal, adjacent face fine.

> **§7 changes the verdict on Path A.** My objection above is that Path A fails unsafe. A position sensor wired as a fan interlock inverts that: the fan physically cannot run unless the louver is confirmed open. That makes Path A a defensible budget build, and it's covered in §7.

---

## 4. Recommended Design

```
                    ┌── fixed frosted glass (sloped, stays put) ──┐
                   ╱                                              │
                  ╱   ┌──────────────────────────────────────┐    │
    metal box    ╱    │  §6 actuator  -or-  §7 sensor        │    │
    above roof  ╱     │  on the louver shaft / arm           │    │
               │      │                                      │    │
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

**Use a remote inline fan, not a ceiling-mount housing.** A conventional ceiling fan is a big white box that has to sit *in* the ceiling plane — right in the middle of your clear panel, blocking the daylight you're trying to preserve, and needing structural support the polycarbonate can't provide. An inline fan like the [Panasonic WhisperLine FV-20NLF1](https://www.amazon.com/Panasonic-FV-20NLF1-WhisperLine-Line-6-Inch/dp/B000EDUIX2) mounts in an adjacent joist bay; all that shows at the ceiling is a small grille. Its housing is insulated specifically against condensation and noise, and the motor ends up 8 feet from your head rather than directly above it.

**Sizing:** 50 CFM is the code minimum for intermittent bathroom exhaust; ~1 CFM per square foot of floor is the working rule. A light well adds volume the rule doesn't count, so **80–160 CFM** is the target. The FV-20NLF1 is 240 CFM — use a speed control, or step down to the FV-10NLF1 (160 CFM) in the same family.

**Makeup air:** undercut the bathroom door ¾" or fit a grille. Whatever you exhaust has to come back in, or the fan stalls against its own static pressure and the mirror stays foggy.

**What the louver does for you:** a large-aperture passive stack vent. Open louver plus a running fan clears a steamy bathroom dramatically faster than the fan alone, and on a summer night you can open it and get free stack ventilation with no fan at all. Keeping the manual chains working is worth it for exactly this — and under §7 it's the whole point.

---

## 5. The Clear Ceiling Panel

Unchanged, and it works better here — since the glass above is fixed, the daylight path is permanent and your panel never interacts with anything that moves.

**Material: polycarbonate, not acrylic.** Acrylic is cheaper and marginally clearer but brittle, prone to crazing, and worse in a fire. Polycarbonate (Lexan) is effectively unbreakable at this thickness and shrugs off the temperature swing in a light well. Get the **UV-coated** grade — it sits under open sky all day and uncoated polycarbonate yellows.

- Home Depot stocks a [24" × 36" × 0.093" UV-coated Lexan sheet](https://www.homedepot.com/p/LEXAN-24-in-x-36-in-x-0-093-in-Shatter-Resistant-Clear-UV-Coated-Polycarbonate-Sheet-11N43101/333295898) (~$50) that covers a 24 × 32 opening with trim to spare.
- **0.093" is thin for an unsupported 32" span** — it will sag and oil-can. Step up to ¼" ([24 × 36 available](https://www.amazon.com/Clear-Polycarbonate-Lexan-Sheet-24/dp/B06XH5T1G3)) or add a frame member. ¼" is the safer call.

**Build it as a removable framed panel, not a glued-in sheet.** You will need to get back up there. Design for that on day one.

1. Rip a **1× or ¾" plywood frame** to the opening's inside dimensions, with a rabbet or stop bead to receive the panel.
2. Drop the sheet in and secure with a **removable stop bead** and screws. No adhesive, no silicone bead you'll be cutting out in three years.
3. Cut the grille hole with a hole saw at **slow speed, light pressure** — polycarbonate melts and grabs if you rush it. Leave the protective film on until the end.
4. **Oversize every screw hole by 1/16".** The sheet moves noticeably across a 32" span, and a pinned sheet cracks or bows.
5. Gasket the perimeter with closed-cell foam tape. You want the fan pulling through the *grille*, not around the panel edges.

**Design suggestions:** run a thin white or paint-matched trim ring around the grille so it reads as a deliberate fixture rather than a punched hole. Put the grille **off-center, toward the shower end** — it catches steam at the source and looks less like a bullseye. And leave the chains reachable: plan a grommeted slot rather than discovering the problem after the panel is cut. **Under §7 the chains stay in daily use, so make that slot genuinely comfortable to reach.**

---

## 6. Option 1: Motorizing the Louver

You're not lifting a heavy glazed sash — you're rotating a light metal damper. That's a solved problem with an off-the-shelf part.

### The right part: a spring-return damper actuator

A **[Belimo LF24](https://www.belimo.com/us/shop/en_US/p?code=LF24+US)** is a 24VAC/DC fail-safe damper actuator: 35 in-lb torque, 95° rotation, mounting directly to a **3/8"–1/2" shaft** with a universal clamp, with a graduated position indicator.

The important property is **spring return**. Powered, it drives open and holds. Cut power and a spring drives it closed — mechanically, no electronics involved.

**That's your original requirement in one component.** Fan on, louver opens. Fan off, louver closes. No polarity-reversing relay, no timers, no always-on supply, no controller. And it fails closed on a power outage, a blown fuse, or a dead actuator — the safe direction.

Two variants:
- **LF24** — plain on/off, the base part.
- **[LF24-S](https://www.fwwebb.com/product/Belimo/LF-Damper-Actuator/LF24-SUS/256371)** — adds a built-in SPDT auxiliary switch, adjustable 0° to 95°. **This is §7's position sensor, built into the actuator.** If you go this route you get the sensing for free; see §7 for what to do with it.

**Torque check:** 35 in-lb suits a small, free-moving damper. If your louver is stiff, rusty, or larger than expected, service it first; if it's still heavy, step up to a Belimo AF24 (~180 in-lb). **Free the mechanism before you motorize it** — an actuator will grind itself to death against a seized pivot.

### The fiddly part: coupling to the linkage

The LF24 clamps to a round shaft. A century-old sheet-metal louver may or may not present one.

- **Round shaft, 3/8"–1/2"** — ideal. Clamps straight on. With enough shaft length you can leave the sprocket and chains in place, though note the chains now work against the spring, so manual operation becomes an emergency override rather than something pleasant to use daily.
- **Flat crank arm on a pivot pin** — you'll need a crank-arm/linkage kit. Belimo sells universal crank arm and ball-joint kits for exactly this retrofit case. Bring photos to a controls supplier ([Kele](https://www.kele.com/product/actuators-and-dampers/spring-return/belimo/lf24), [Jackson Systems](https://jacksonsystems.com/product/belimo-lf24-us-2-position-actuator-spring-return/), or a local HVAC controls house) and let them spec it against what you actually have.
- **Linkage is hopeless** — this is where §7 stops being an alternative and becomes the answer.

**Mounting:** the actuator needs a solid bracket inside the metal box, and it now lives in an unconditioned, occasionally damp outdoor enclosure. Mount it as high and sheltered as geometry allows, run the low-voltage wiring with a drip loop, and check it annually.

---

## 7. Option 2: Sensing the Louver Instead of Moving It

**Keep the louver manual. Add a switch that reports whether it's open. Let the fan logic use that.**

This is a real alternative, not a consolation prize. It's worth taking seriously for four reasons:

1. **It deletes the project's one genuine unknown.** Everything else here is specified. The actuator-to-linkage coupling is the only item where I've had to say "photograph it and let a supplier figure it out," and it's the only one that could turn into a fabrication job. A sensor doesn't care what the linkage looks like — it just watches something move.
2. **It's roughly a sixth of the cost.** ~$40 versus $200–330.
3. **It keeps the chains pleasant to use.** A spring-return actuator means every manual pull fights the spring. Sensor-only, the louver operates exactly as it does today.
4. **Nothing moving lives on the roof.** A sealed switch in a cold damp metal box has no duty cycle, no motor, no spring, no annual check. The actuator is a maintenance item; the switch isn't.

The cost is real and should be stated plainly: **you give up automatic opening.** You pull the chain. That was your original ask, so this is a genuine reduction in scope — the question is whether it's a reduction you'd notice, given the chains are already hanging in the room.

### What the sensor does depends on which path you built

**Under Path B (ducted — recommended): the sensor is an informer.** The fan is independent of the louver, so the sensor isn't protecting anything. It gives you:
- A **reminder** — "shower's been running six minutes and the louver is shut, go pull the chain."
- A **left-open alert** — "louver open, outdoor temp 24°F," which is the failure mode that actually costs you money in a Philadelphia winter.
- An **indicator lamp** by the light switch, if you'd rather not involve software at all.
- **Data.** Logged over a season you'd learn how often it's actually left open, and whether it seals when closed.

**Under Path A (no ductwork): the sensor is a hard interlock, and it rescues the path.** Wire the fan's power through the sensor contact and the fan physically cannot run unless the louver is confirmed open. My objection to Path A in §3 was that it fails unsafe — steam pumped into a sealed cold metal box. This inverts that completely: it now fails **annoying** (no fan until you pull the chain) rather than **unsafe** (invisible moisture in the joist bay). That's a fair trade, and it makes Path A the cheapest complete build here by a wide margin — no duct, no wall cap, no actuator.

Two things I won't soften about Path A even with the interlock: it still doesn't satisfy M1507.2 on paper, since the light well isn't a duct regardless of louver position; and "the fan doesn't work when I forget" is a real usability cost that some people find maddening. But as a physical design, sensor-interlocked Path A is defensible in a way bare Path A is not.

### Wire it fail-safe — this is the part to get right

Under Path A the sensor is a safety device, so its **failure direction matters more than its accuracy.**

Use a **normally-open contact that closes only when the louver reaches open**, in series with the fan. Then a broken wire, a corroded terminal, a dead switch, or a magnet that fell off all produce the same result: **fan blocked.** Never wire it so that a failed sensor permits the fan — that's the one configuration that reproduces the exact problem you installed it to prevent.

### Trip at *fully* open, not at first crack

The classic mistake. A louver cracked an inch is not a vent, but a carelessly placed switch will happily report "open." Position the switch so it makes near the **end of travel**, and verify by watching it: pull the chain slowly and confirm the contact doesn't make until the louver is genuinely open. If you fit the LF24-S instead, its aux switch is adjustable across the full 0–95° for the same reason.

### Sensor hardware

**Go wired, not wireless.** Two reasons, both specific to your situation:
- **Battery chemistry hates a cold rooftop box.** Coin cells and lithium AAAs lose much of their capacity below freezing, and you'd be replacing them on a ladder in February.
- **A sheet-metal box is a Faraday cage.** RF out of a soldered galvanized enclosure is unreliable at best.

Put the **sensor** on the roof and the **electronics** in the conditioned space — run a pair of thermostat wire down the shaft and land it on an input indoors.

| Type | Notes |
|---|---|
| **Sealed roller-lever limit switch** *(my pick)* | Mechanical, unambiguous, immune to the surrounding steel. [IP67 SPDT roller-lever switches](https://www.amazon.com/SJZBIN-Switch-Waterproof-Sealed-Roller/dp/B0CHYC695J) are a few dollars. SPDT gives you both a normally-open and normally-closed contact, so you can wire the interlock one way and an indicator the other. Needs a small bracket positioned so the louver arm trips it at full travel |
| **Wide-gap magnetic contact** | No physical contact to wear. **Ordinary reed switches misbehave mounted on ferrous metal** — the steel distorts the field. Use contacts built for steel doors: the [Seco-Larm SM-226L-3Q](https://www.seco-larm.com/product/sm-226l-3q/) is IP66 epoxy-sealed with armored leads and a 2¾" gap, designed for exactly this abuse; the [SM-4601-L3Q](https://www.seco-larm.com/product/sm-4601-l3q/) is a 3"-gap aluminum-housed IP54 version. Both offer NO and NC in one unit |
| **Heavy-duty garage-door contact** | [Konnected's wired magnetic sensor](https://konnected.io/products/garage-door-or-gate-heavy-duty-wired-magnetic-contact-sensor) is the same idea in a hobbyist-friendly package — normally open, closing when the magnet comes within ~1" |
| **Battery Zigbee/Z-Wave door sensor** | **Avoid**, for the cold and Faraday reasons above |

### Landing the signal indoors

- **No software at all:** run the contact straight to a 24V pilot lamp beside the light switch, and (Path A) in series with the fan.
- **With Home Assistant:** land it on a [Shelly Plus i4 DC](https://us.shelly.com/products/shelly-plus-i4) — four dry-contact/12–24V DC inputs, WiFi, native HA integration. One device covers the louver sensor with three inputs to spare. Then the reminders, alerts, and logging above are a few lines of automation.

### It stacks with §6 rather than competing

If you later add the actuator, the sensor wiring isn't wasted — it becomes independent confirmation that the louver actually reached open, which is exactly what you want from a motorized damper on a roof you can't see. Or you drop it and use the LF24-S's built-in aux switch instead. **Nothing you buy in §7 is stranded by a later §6 upgrade.**

---

## 8. Control Wiring

### If you went §6 (actuator)

```
   fan switch (120V switched hot) ──┬──► inline fan
                                    │
                                    └──► 24VAC transformer ──► Belimo LF24
                                         (40VA, Class 2)        (spring closes
                                                                 on loss of power)
```

That's the whole system. One transformer, two wires. The transformer's primary lands on the fan's switched hot, so the actuator is energized exactly when the fan is.

- Use a **Class 2, 40VA 24VAC transformer** — comfortably above the LF24's draw.
- **The louver takes a moment to open; the fan starts instantly.** Under Path B this is harmless — the fan is ducted to its own cap and doesn't care. (Under Path A you'd need a delay-on-make timer; another reason to build Path B.)
- Transformer goes in a proper junction box in the joist bay, not loose in the shaft.
- With the **LF24-S**, run its aux switch to an indicator LED by the light switch.

### If you went §7 (sensor), Path B

```
   fan switch (120V switched hot) ──► inline fan        (independent)

   louver switch ──── 18/2 down the shaft ──── Shelly Plus i4 DC ──► Home Assistant
                                                (or a 24V pilot lamp)
```

The fan circuit is untouched — a completely ordinary bathroom fan install. The sensor is a parallel, low-voltage, entirely optional-to-the-fan's-operation circuit. This is the simplest wiring of any option here.

### If you went §7 (sensor), Path A — the interlock

```
   fan switch (120V hot) ──► [louver NO contact] ──► inline fan
                                    │
                                    └──► 18/2 ──► Shelly i4 input (optional, for alerts)
```

- The contact must be **normally open, closing at full louver travel** (see §7).
- **Don't switch 120V through a low-voltage reed contact.** Magnetic contacts are typically rated for small DC signal loads, not line-voltage motor current. Use the contact to drive a **24V relay**, and let the relay's contacts carry the fan. The limit-switch option can often carry the load directly, but check its rating against the fan's amperage first.
- Consider a **manual bypass switch** in a spot you have to deliberately reach for — for the day the sensor fails at 11pm and you want a shower. Label it, and don't make it convenient enough to become the default.

### Optional smart controls (either path)

Worth it if you already run Home Assistant. A [Shelly Plus 1 or 1PM](https://us.shelly.com/) on the fan circuit adds:
- **Humidity triggering** — fan runs on humidity, not just the wall switch. The upgrade you'd notice daily.
- **Run-on timer** — keep going 10 minutes after the shower.
- **Cold-weather logic** — §6: run the fan without opening the louver below ~35°F. §7: prompt you *not* to open it.
- **Failure alerting** — paired with either the LF24-S aux switch or the §7 sensor.

The Shelly switches line voltage on the fan circuit, which is what it's designed for.

---

## 9. One Thing to Check: The Glass

While you're up there, look closely at that frosted panel. **If it's wired glass** — obscure glass with steel mesh embedded, very common in this vintage of rooftop sheet-metal work — note that **wired glass fails the modern safety-glazing test and isn't permitted in overhead applications** except within fire-rated assemblies. Under current code, [sloped and overhead glazing](https://up.codes/s/sloped-glazing-and-skylights) must be laminated, or filmed, or protected by a screen below, so broken glass stays put rather than falling on whoever's underneath.

The counterintuitive part: the embedded wire makes wired glass *weaker* than plain annealed glass, and when it breaks it leaves shards hanging on wire instead of falling clear.

Flagging it, not sounding an alarm. It's existing construction, it's been there for decades, nobody requires you to change it, and it's no reason to postpone this project. Two things follow:

1. **Your clear polycarbonate panel does incidental safety duty** — it sits below the well and would catch anything coming down. A genuine, if accidental, benefit, and one more argument for ¼" over 0.093".
2. **If you ever replace the glazing**, that's the moment to go laminated. Not a today problem.

---

## 10. Condensation and Winter

- **Insulate the duct.** Any duct in the unconditioned zone between ceiling and wall cap must be insulated or it condenses internally and drips back through the fan — the most common cause of "my bathroom fan is leaking."
- **Slope the duct toward the wall cap** so condensate runs outward.
- **Expect condensation on the underside of the fixed glass** in winter regardless. Single-pane glass on a metal box does that. Path B means you're no longer *adding* shower steam to it.
- **Air-seal the clear panel.** A leaky panel turns the light well into a chimney bleeding conditioned air 24/7. Between the panel gasket and a closed louver, the shaft should be reasonably sealed when nothing's running.
- **Winter louver behavior:** under §6, don't open it when it's cold (§8, smart controls). Under §7, the left-open alert *is* the winter feature — a manually-operated louver forgotten open in January is a genuinely expensive mistake, and the sensor catches it.

---

## 11. Install Sequence

1. **Get up there and photograph everything** (§1) — the louver's shaft or crank arm, the free wall area on non-louvered faces, the metal's condition. Measure the shaft diameter with calipers.
2. **Exercise the louver by hand.** Open and close it a dozen times. Clean and lubricate the pivots. **Do this before ordering anything** — under §6 it determines torque; under §7 it determines whether daily manual operation is pleasant or a chore, which is the whole question.
3. **Decide §6 vs §7** (see §12). If §6, take your photos to a controls supplier and have them spec the actuator plus linkage kit — the one part worth a human's eyes on.
4. **Cut and fit the wall cap** in the metal box, on a face away from the louver. Hole saw or step drill, deburr, seal the cut edge, gasket, sealant, screw the flange.
5. **Run power to the well** — a switched leg for the fan, plus the transformer (§6) or sensor pair (§7).
6. **Mount the inline fan** in the adjacent joist bay on its suspension brackets; avoid hard-mounting to framing — the isolation is worth it acoustically.
7. **Duct it up**, insulated, sloped toward the cap.
8. **Fit the actuator or the sensor.**
   - *§6:* mount on the shaft, set rotation limits, then **power-cycle a dozen times with the panel still off** and watch the full travel for binding.
   - *§7:* mount the bracket, then **pull the chain slowly a dozen times** and confirm the contact makes only near full open — not at first crack.
9. **Wire it up** (§8). Verify with a multimeter before closing anything. Under Path A, deliberately break the sensor circuit and confirm **the fan refuses to run**.
10. **Build and fit the clear panel** last, with the grille cut to the duct drop and a comfortable grommeted slot for the chains.
11. **Commission it:** hot shower, door closed. Mirror should clear in a few minutes. Under §6, kill the breaker mid-cycle and confirm the louver springs shut. Check the duct for condensation after the first cold night.

---

## 12. Summary Recommendation

**Build Path B either way** — inline fan, insulated duct, wall cap in a non-louvered face of the metal box. No roof work, no roofer, and the fan is independent of everything else. That part isn't a close call.

**On §6 versus §7, the honest answer is to start with §7.**

Not because motorizing is wrong — the Belimo does exactly what you originally asked, elegantly, and if the louver turns out to have a clean 3/8"–1/2" shaft it's a genuinely satisfying afternoon. But:

- **§7 costs ~$40 and can be installed in an hour.** §6 costs $200–330 and hinges on a linkage you haven't characterized yet.
- **You don't actually know yet whether manual operation bothers you.** The chains are already there, already in reach. Live with a season of pulling them, with the sensor telling you when you've forgotten and when you've left it open in the cold. If it's fine, you've saved $300 and a rooftop maintenance item. If it drives you up the wall, you've lost nothing.
- **Nothing is stranded.** Add the actuator later and the sensor becomes its position confirmation — or you drop it for the LF24-S's built-in aux switch.
- **Step 2 of §11 tells you a lot.** If the louver moves sweetly, manual operation is a non-issue and §7 is clearly enough. If it's stiff and awkward even after servicing, that's an argument for motorizing — and also a warning that 35 in-lb may not be enough.

**The rest, unchanged:**
- ¼" UV-coated polycarbonate panel in a removable wood frame, small off-center grille, comfortable grommeted chain slot.
- Keep the manual chains fully working — essential under §7, still worth it under §6 for summer stack ventilation.
- Wire the sensor fail-safe: normally open, closing at full travel, so any failure blocks the fan rather than permitting it.

**Estimated parts cost:**

| Build | Parts |
|---|---|
| Path A + §7 sensor interlock (cheapest complete build) | **$300 – $450** |
| **Path B + §7 sensor (recommended starting point)** | **$400 – $600** |
| Path B + §6 actuator (the original ask, fully automatic) | $550 – $850 |
| *Later upgrade from §7 to §6* | +$200 – $330 |

No contractor line item in any of these, if you're comfortable with a hole saw, a duct run, and a 120V circuit.

---

## 13. Things That Will Bite You

| Symptom | Likely cause |
|---|---|
| Louver opens but fan barely moves air | No makeup air — undercut the bathroom door |
| Exhaust smell or steam coming back in the louver | Wall cap sited on the same face as the louver — recirculation loop |
| Water dripping from the ceiling grille | Uninsulated duct condensing, or duct sloped back toward the fan |
| Sensor says "open" but the room won't clear | Switch trips at first crack instead of full travel (§7) |
| Fan runs with the louver shut, on Path A | Sensor wired fail-permissive — must be NO, closing at full open |
| Wireless sensor drops out constantly | Battery in the cold, or RF trying to escape a sheet-metal box. Go wired |
| Magnetic contact reads erratically | Ordinary reed switch mounted on ferrous metal. Use a wide-gap steel-door contact |
| Actuator buzzes and won't reach full open | Under-torqued for a stiff louver, or a pivot you didn't free up first |
| Louver doesn't close on power-off (§6) | Spring return defeated by a binding linkage — the spring has limited authority |
| Chains suddenly stiff to pull | Expected after fitting a spring-return actuator; they now work against the spring |
| Clear panel bows or cracks near screws | Screw holes not oversized for thermal movement, or sheet too thin for the span |
| Panel fogs on the inside | Panel not sealed to the ceiling; humid air migrating around the edges into the well |
| Louver frozen shut after a snow | Packed snow in the blades. Path B means the fan doesn't care — that's the point |
| Everything works, room still steams up | Fan undersized for the added light-well volume — the space is bigger than it looks |

---

## 14. Code Notes

- **[IRC M1507.2](https://www.jaspector.com/codes/irc-2018/ch15-exhaust-systems/bathroom-exhaust-fan-outdoor-termination-irc-2018/)**: bathroom exhaust must discharge directly outdoors, never into an attic, crawl space, or other interior area. A light well and a closed metal box count as interior. This is the argument for Path B, and it holds even with a §7 interlock — the interlock fixes the physics, not the paperwork.
- **50 CFM** intermittent (or 20 CFM continuous) is the standard bathroom minimum.
- **Ducts in unconditioned space must be insulated** against condensation.
- **[Overhead glazing](https://up.codes/s/sloped-glazing-and-skylights)** must be laminated, filmed, or screened — see §9. Applies to replacement, not to leaving existing construction alone.
- **Philadelphia L&I**: new branch circuits generally mean a permit. Check with [Licenses & Inspections](https://www.phila.gov/departments/department-of-licenses-and-inspections/), particularly if your block is in a historic district — visible rooftop changes can draw extra review, though a small wall cap on an existing ventilator is about as low-profile as it gets. **No option here involves a new roof penetration**, which is the part that usually triggers roofing scope and a contractor.

See `skylight_vent_fan_purchase_list.md` for parts and links.
