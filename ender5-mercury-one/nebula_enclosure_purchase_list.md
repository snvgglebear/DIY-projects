# ZeroG Nebula Enclosure — Purchase List

Parsed from ZeroG's official
[**Enclosure Docs.pdf**](https://github.com/ZeroGDesign/Nebula/blob/Nebula1/Enclosure%20Docs.pdf)
(28 pages, `ZeroGDesign/Nebula`, branch `Nebula1`). Every quantity and dimension below is taken
from that document.

Companion docs: [`quad_z_mod.md`](./quad_z_mod.md) ·
[`electronics_enclosure_options.md`](./electronics_enclosure_options.md)

> **Verification status:** The PDF was downloaded and parsed directly, and the panel drawings were
> read from rendered pages — those numbers are solid. Vendor links were found through search;
> storefronts are blocked by this environment's egress proxy, so **stock, price and kit contents
> were not confirmed**. Check what a kit actually includes before ordering.

---

## Before you buy: what Nebula actually is

The cover page reads *"A Zero G Enclosure For The Mercury One.1 and Hydra."* Two things follow:

1. **This is a frame, not a wrap-around box.** The BOM builds the whole structure out of 2020/2040
   extrusion (parts A–I). It replaces the printer frame rather than enclosing a stock Ender 5 Plus
   frame. Committing to Nebula means rebuilding the machine into it.
2. **It assumes Hydra.** Steps 11–12 assemble and locate a "Hydra sub-frame", the bottom 2040s need
   their blind cross holes facing rearward "for Hydra", and the deck panel is *"the same as stock
   Hydra"*. Without the [Hydra](https://github.com/ZeroGDesign/Hydra) Z system, parts of this BOM
   have nothing to attach to.

### Which size

| | Nebula 255 | Nebula 370 |
|---|---|---|
| Vendor naming | "255 Pro" | "370 Plus" |
| Class | Ender 5 / 5 Pro | **Ender 5 Plus** |
| Doors | 1 | 2 |
| Uprights | 2020 × 650 mm | 2020 × 800 mm |

**For an Ender 5 Plus-class machine, that's the 370.** Both columns are listed throughout so you
can confirm against your own gantry before ordering — cut extrusion isn't returnable.

---

## Fastest path: buy the kits

The frame kit is made by LDO and ships **pre-cut, tapped and anodized**, which removes the entire
extrusion-cutting problem. There is also a separate panel kit, which removes the laser-cutting
problem.

| Kit | Vendor | Link |
|---|---|---|
| Nebula frame kit (Pro/Plus, multiple colours) | KB-3D | <https://kb-3d.com/store/frame-enclosure/3971-zero-g-mercury-one1-nebula-frame-kit-multiple-sizes-colors.html> |
| Nebula frame kit (Pro/Plus, multiple colours) | West3D | <https://west3d.com/products/zero-g-mercury-one-1-nebula-frame-kit-pro-and-plus-multiple-colors> |
| Nebula frames by LDO (from ~$180) | Fabreeko | <https://www.fabreeko.com/products/zero-g-mercury-nebula-frames-by-ldo> |
| **Nebula enclosure panel kit** (multiple sizes/colours) | KB-3D | <https://kb-3d.com/store/frame-enclosure/3972-zero-g-nebula-enclosure-panel-kit-multiple-sizes-colors.html> |
| Nebula assembly kit — **370 Plus** | KB-3D | <https://kb-3d.com/store/frame-enclosure/3992-mercury-one1-zero-g-nebula-assembly-kit-370-plus-1742175112370.html> |
| Nebula assembly kit — 255 Pro | KB-3D | <https://kb-3d.com/store/frame-enclosure/3991-mercury-one1-zero-g-nebula-assembly-kit-255-pro-1742175112513.html> |
| All ZeroG products | Fabreeko | <https://www.fabreeko.com/collections/zero-g> |

**Check what the "assembly kit" contains** before also buying loose fasteners — the name suggests it
covers the hardware BOM below, but that couldn't be confirmed from here.

---

## Extrusion BOM

From p5. Note the document's own warning: *"Extrusions D and E are the same length for the Nebula
255, but not the same length for the Nebula 370."*

| ID | Qty | Nebula 255 | Nebula 370 |
|---|---|---|---|
| A | 4× | 2020 — 650 mm | 2020 — 800 mm |
| B | 2× | 2020 — 450 mm | 2020 — 600 mm |
| C | 2× | 2040 — 450 mm | 2040 — 600 mm |
| D | 2× | 2020 — 410 mm | 2020 — 560 mm |
| E | 2× | 2020 — 410 mm | 2020 — 538 mm |
| F | 4× | 2040 — 400 mm | 2040 — 510 mm |
| G | 1× | 2020 — 400 mm | 2020 — 510 mm |
| H | 2× | 2020 — 370 mm | 2020 — 498 mm |
| I | 3× | 2020 — 370 mm | 2020 — 498 mm |

**If cutting your own**, total linear metres needed (before saw waste):

| Profile | Nebula 255 | Nebula 370 |
|---|---|---|
| 2020 | 7.39 m | 9.60 m |
| 2040 | 2.50 m | 3.24 m |

---

## Hardware BOM

From p3 (255) and p4 (370).

| Part | Nebula 255 | Nebula 370 |
|---|---|---|
| 5×7×8 mm bushing (bronze) | 6 | 12 |
| 5×50 dowel pin | 3 | 6 |
| 6×3 magnet | 36 | 36 |
| 10×3 magnet | 16 | 48 |
| M3×8 FHCS | 2 | 2 |
| M3×8 BHCS | 36 | 44 |
| M3×8 SHCS | 38 | 46 |
| M3×12 SHCS | 46 | 52 |
| M3×14 SHCS | 8 | 8 |
| M3 roll-in T-nut | 106 | 118 |
| M3 × D5.0 × L4.0 heat-set insert | 26 | 34 |
| M5×6 BHCS | 8 | 8 |
| M5×10 BHCS | 66 | 66 |
| M5×10 FHCS | 6 | 6 |
| M5×12 SHCS | 4 | 4 |
| M5×18 SHCS | 14 | 14 |
| M5 roll-in T-nut | 52 | 52 |
| M5 hex nut | 4 | 4 |
| 3×5 mm foam tape | 10 m | 15 m |
| 38×19 mm rubber foot | 4 | 4 |

The fastener spec is Voron-standard — M3/M5 SHCS/BHCS/FHCS, roll-in T-nuts, and the M3×5×4 heat-set
is exactly the Voron/LDO insert. A Voron 2.4-class fastener kit gets you most of the way, but
**quantities won't match**, so expect to top up.

---

## Panels

Material spec, from pp. 14, 25–28: **3 mm or 1/8 inch recommended**; door panels also accept
**5.5–6 mm or 1/4 inch**. Frame panels can be saw cut; door and electronics-bay panels are laser or
CNC cut. The top panel of the electronics bay is *"the same as stock Hydra."*

### Nebula 255

| Panel | Qty | Size |
|---|---|---|
| Door | 1 | 435 × 644 mm |
| Side and back | 3 | 424 × 644 mm |
| Top (chamfered corners) | 1 | 424 × 424 mm |
| Electronics bay — bottom | 1 | 442.5 × 442.5 mm |
| Electronics bay — deck | 1 | 383 × 383 mm |

### Nebula 370

| Panel | Qty | Size |
|---|---|---|
| Door | 2 | 275.9 × 794 mm |
| Side | 2 | 574 × 792 mm |
| Back | 1 | 552 × 792 mm |
| Top (chamfered corners) | 1 | 542 × 574 mm |
| Electronics bay — bottom | 1 | 570.5 × 547.5 mm |
| Electronics bay — deck | 1 | 511 × 528 mm |
| Electronics bay — bottom extension | 1 | 570.5 mm wide (height not dimensioned) |

Panels carry cut-outs and mounting holes, so **order from the DXF/CAD in the Nebula repo** rather
than from these outside dimensions alone — the numbers here are for sizing stock and getting quotes.

---

## Printed parts

From <https://github.com/ZeroGDesign/Nebula>. Per the door-hinge BOM on p15 (note the ZeroG logo
option for hinges):

| Part | Nebula 255 | Nebula 370 |
|---|---|---|
| Printed base hinge | 3 | 6 |
| Printed hinge | 3 | 6 |

Plus, from pp. 19–24: bottom 2040 door catch (**left and right hand versions**), upper frame door
catch, door corner caps, 255 magnetic door catch, and the 255 / 370 door handles.

---

## Where to buy, by category

Direct product links below were surfaced by search; the generic-hardware entries are search links,
since exact listings churn.

**Fastener kits**

- Voron 2.4 fastener kit — 3D Lab Tech: <https://www.3dlabtech.ca/product/voron-2-4-fasteners-kit/>
- Voron fastener kits — DFH: <https://dfh.fm/products/voron-v0-fasteners-kit>
- All-in-one Voron-compatible screw kit — Amazon: <https://www.amazon.com/dp/B0GXFGZX9X>

**Heat-set inserts (M3 × D5 × L4 — "Voron spec")**

- CNC Kitchen: <https://cnckitchen.store/products/made-for-voron-gewindeeinsatz-threaded-insert-m3x5x4-100-stk-pcs>
- Vector 3D: <https://vector3d.shop/products/heat-set-insert-m3x5x4-voron-ldo>

**Generic hardware**

| Item | Search |
|---|---|
| M3/M5 roll-in T-nuts | <https://www.amazon.com/s?k=M5+roll+in+t-nut+2020+extrusion> |
| 6×3 mm and 10×3 mm magnets | <https://www.amazon.com/s?k=6x3mm+neodymium+disc+magnets> |
| 5×7×8 mm bronze bushing | <https://www.amazon.com/s?k=5x7x8mm+bronze+sleeve+bearing> |
| 5×50 mm dowel pin | <https://www.amazon.com/s?k=5x50mm+dowel+pin> |
| 3×5 mm foam tape | <https://www.amazon.com/s?k=3mm+x+5mm+foam+seal+tape> |
| 38×19 mm rubber feet | <https://www.amazon.com/s?k=38mm+rubber+feet+19mm> |

**Panels (acrylic / polycarbonate, cut to size)** — TAP Plastics, ePlastics, Canal Plastics,
Delvie's Plastics, or a local laser cutter. Send the repo's DXF files rather than dimensions.

**Extrusion (if not buying the LDO kit)** — Misumi, 8020.net, TNutz, or KB-3D. Specify cut-to-length
and tapped ends.

---

## Discrepancies to check before ordering

Found while parsing — none are blockers, but they change what you buy:

- **Magnet thickness.** The hardware BOM (pp. 3–4) lists **10×3** magnets, but the assembly steps
  say **10×2** three times: p20 *"Press in 3 10x2 magnets"*, p21 *"Press fit 2 10x2 magnets"*,
  p24 *"press in two 10x2mm magnets"*. Buying to the BOM (10×3) risks magnets that won't seat in
  the printed pockets. Worth asking ZeroG, or buying a few of each.
- **Extrusion E quantity.** The BOM (p5) lists **E 2×**; the assembly step (p9) lists **E 1×**.
- **Extrusion I quantity.** The BOM lists **I 3×**; step p13 installs the "final rear 2020
  extrusion" with only 2× M5×10 BHCS.
- **Door panel thickness.** 3 mm keeps the whole build one material; 6 mm doors are heavier and
  stiffer but load the printed hinges more.

---

## Next steps

- [ ] Confirm 370 (Plus) is the right size against your gantry and bed
- [ ] Decide kit vs self-source — the LDO frame kit removes the cutting and tapping entirely
- [ ] Check whether the KB-3D "assembly kit" already covers the hardware BOM
- [ ] Resolve the 10×3 vs 10×2 magnet question before ordering 48 of them
- [ ] Pull the panel DXFs from the Nebula repo for laser quotes
- [ ] Confirm Hydra is in the plan — Nebula's sub-frame and deck panel assume it

---

## Sources

- [ZeroGDesign/Nebula — `Enclosure Docs.pdf`](https://github.com/ZeroGDesign/Nebula/blob/Nebula1/Enclosure%20Docs.pdf) — the source for every BOM figure here
- [ZeroGDesign/Nebula — repo](https://github.com/ZeroGDesign/Nebula)
- [ZeroGDesign/Hydra — repo](https://github.com/ZeroGDesign/Hydra)
- [ZeroG documentation](https://docs.zerog.one/)
