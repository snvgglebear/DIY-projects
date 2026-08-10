# 5" Frame Part-Merging Guide (Onshape)

**Version:** 1.0
**Date:** August 2025
**Applies to:** Frame Option A from `drone_purchase_list.md` §2 (virk3D's Lumenier QAV-PRO Whoop 5" replica, Printables)
**Companion Documents:** `drone_build_guide.md`, `drone_purchase_list.md`, `drone_kit_alternatives.md`, `drone_lora_still_image_uplink.md`

---

## Why This Exists

All three verified frame designs in the purchase list (`drone_purchase_list.md` §2) are multi-part — arms, center plate, and duct sections that bolt together, not a single continuous piece. On a 600×600×600mm printer, bed size isn't the constraint; the part count comes from other design reasons (print orientation per part, support-free printing, and repairability after a crash — see the note in `drone_build_guide.md` §3).

If you'd still rather have fewer, larger printed parts, the practical path is to merge some of an existing design's pieces yourself in CAD — not find a design that's already built that way. This guide walks through doing that in Onshape.

## Before You Start: Search Findings

Before writing this guide, a search was done for a real, verified 5" **enclosed/ducted cage** design that's purpose-built with few large parts (rather than merged after the fact). Nothing matched both criteria:

- **ARS-5** (a well-known, free, printable 5" frame) keeps surfacing as a "few parts" design, but it's an **open racing frame**, not an enclosed prop-guard cage — it doesn't meet this build's "enclosed cage" spec, so it wasn't substituted in.
- No other genuinely-enclosed 5" cage design turned up that was designed from the start as fewer/larger parts.

One useful data point that did come out of this: **Option C (carnoforge3D) is already only 3 parts** — the leanest of the three verified options in the purchase list, without any merging needed. If part count matters more to you than Option A/B's confirmed gimbal-mount compatibility, Option C is worth a look before you go through the merge process below. Its downside is it's STL-only (no STEP file), so it's not a good candidate for the CAD-merge approach in this guide — a boolean union on a raw STL mesh (versus real STEP solid geometry) is a much less reliable operation.

That's why this guide uses **Option A**, which bundles a STEP file alongside the STL/DXF.

## Step-by-Step: Merging Parts in Onshape

1. **Get an Onshape account.** The free Standard plan works for this. One catch: documents on the free tier are public (anyone with the link can view/copy). Not a problem here since Option A is free/low-cost, but worth knowing if you plan to do this with paid files later.

2. **Download the STEP file** from the Printables listing: [virk3D — CineWhoop 5-Inch Frame](https://www.printables.com/model/1511864-cinewhoop-5-inch-frame-replica-of-lumenier-qav-pro). It bundles STL + STEP + DXF — grab the `.step`/`.stp`.

3. **Import it into Onshape.** Create a new document, then drag the STEP file in, or use **Create → Import**. You'll land in one of two states:
   - **One Part Studio, multiple solid bodies** — the easy case, skip to step 5.
   - **An Assembly with separate Part Studios per part** — more likely for a multi-part design like this. Handle it in step 4.

4. **If it imported as an Assembly:** open a new Part Studio, then use **Insert → Derive** to pull in each part you want to merge (e.g., one arm + the center plate) *at its assembled position* — Derive can reference the assembly's mate connectors so the geometry lands where it actually sits in the real frame, not at the part's own origin. Repeat for each part going into the merge.

5. **Check the joint before merging.** Hide everything except the two bodies you're merging and inspect where they meet:
   - **Faces genuinely overlap or are coincident** (common if the design has screws sitting in a mating pocket) — proceed straight to the Boolean.
   - **Parts only touch at a few screw points, with a gap between the main bodies** — a Boolean union will fail or silently do nothing useful here. Sketch and extrude a bridging solid (a gusset/fillet spanning the gap) *before* the union, so there's real material connecting them. This is the step where you're doing actual structural design, not just clicking a merge button — don't skip inspecting it.

6. **Run the Boolean.** In the Part Studio toolbar: **Boolean → Union**, select the bodies (plus any bridge solid from step 5), confirm.

7. **Clean up the seam.** The union leaves now-pointless screw holes and standoff bosses from the original joint sitting inside the merged part. STEP imports come in as "dumb" solids (no parametric feature history), so clean these up with new cuts/fillets on top of the result, not by editing the original features:
   - Sketch + **Extrude Cut** to remove any interior screw bosses that no longer do anything.
   - Add a **Fillet** along the internal seam edge (2–3mm+) — sharp internal transitions are exactly where FDM parts crack; this fillet is doing real structural work, not just cosmetics.

8. **Sanity-check the new part's print orientation.** Rotate it around in Onshape and think through how it'd sit on the bed — a merged arm+center-plate section may now have an overhang or bridge that didn't exist when the pieces were separate and each optimally oriented on its own.

9. **Export STL.** Right-click the merged part → **Export** → STL, resolution around 0.05–0.1mm deviation tolerance is plenty for FDM.

10. **Before you trust it structurally:** load it in your slicer and check the mesh reports as manifold/watertight (a bad boolean often leaves a non-manifold seam the slicer will complain about). Do the propeller-off motor test from `drone_build_guide.md` §8 extra carefully on the merged part, since you've changed how the frame flexes/fails at that joint compared to the tested original design.

## Safety Note

Merging parts changes the frame's engineered joint behavior — the original split points often let the frame flex slightly under impact instead of cracking. Treat a merged frame as an unvalidated design: test-fit it, run the propeller-off motor and vibration checks in `drone_build_guide.md` §8 before ever installing propellers, and fly cautiously on the first few flights.

---

**Document Version:** 1.0
**Last Updated:** August 2025
**Companion Documents:** `drone_build_guide.md`, `drone_purchase_list.md`, `drone_kit_alternatives.md`, `drone_lora_still_image_uplink.md`
