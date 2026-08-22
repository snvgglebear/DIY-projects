# 📡 Rooftop Meshtastic Mesh System - Bill of Materials

**Last Updated:** August 2026
**Configuration:** Flat Roof Mounted | 18" Vertically Separated Antennas | 6 ft Total Mast Height | 1W Radios | Solar + Battery | PoE Options
**Prepared For:** Carl (Software Developer)

---

## Executive Summary

| Configuration | Total Estimated Cost | Best For |
|---------------|---------------------|----------|
| **Budget** | ~$280-$380 | Entry deployments, testing, hobbyist projects |
| **Premium** | ~$550-$750 | Production deployments, maximum range/reliability |
| **Balanced (Recommended)** | ~$332 | Most users seeking performance/value balance |

Add **~$40-$120** to any tier to include a digital TV antenna on the same mast (see [Digital TV Antenna](#digital-tv-antenna-optional-add-on)).

---

## Component Categories

1. [Radio Units](#radio-units)
2. [Antennas](#antennas)
3. [Digital TV Antenna (Optional Add-On)](#digital-tv-antenna-optional-add-on)
4. [Power System](#power-system)
5. [Mounting & Cabling](#mounting--cabling)
6. [Installation Notes](#installation-notes)

---

## Radio Units

### Budget Configuration
| Item | Description | Price | Sources |
|------|-------------|-------|---------|
| Heltec WiFi LoRa 32 V3 | 915MHz, ESP32-S3, ~27dBm output | $30-$40 each | [Amazon](https://www.amazon.com/s?k=heltec+lora+v3) |
| Heltec Mesh Node T114 | nRF52840 ultra-low power alternative | $15-$40 each | [Heltec Store](https://heltec.org/project/mesh-node-t114/) |

**Notes:**
- T114 offers superior power efficiency (23μA sleep current)
- V3 lacks 1W output but adequate for most urban/suburban deployments
- Both support Meshtastic firmware out of the box

### Premium Configuration
| Item | Description | Price | Sources |
|------|-------------|-------|---------|
| LILYGO T-Beam 1W | ESP32-S3, 32dBm (1W), cooling fan | $55-$60 each | [Amazon](https://www.amazon.com/LILYGO-Meshtastic-Development-CH9102F-Soldered/dp/B0B63FV7FR) |
| T-Beam Supreme | GPS, 6-axis IMU, upgraded specs | $50-$70 each | [LILYGO](https://meshtastic.org/docs/hardware/devices/lilygo/tbeam) |

**Notes:**
- 1W output dramatically improves range in challenging terrain
- Active cooling required for sustained high-power transmission
- Built-in GPS essential for position reporting features

---

## Antennas

### Regular Gain Options
| Item | Gain | Connector | Price | Sources |
|------|------|-----------|-------|---------|
| RAK Flexible Whip 915MHz | 2 dBi | SMA-Male | $8 each | [RAK Store](https://store.rakwireless.com/products/sma-male-915mhz-868mhz-whip-antenna-20cm) |
| Generic 3.5 dBi RP-SMA | 3.5 dBi | RP-SMA | $13-$15 each | [AliExpress](https://www.aliexpress.com/item/1005001874331313.html) |
| Lora Long Range (eBay) | 8 dBi | SMA | $30-$35 each | [eBay](https://www.ebay.com/itm/127535860992) |

### High Gain Options
| Item | Gain | Connector | Price | Sources |
|------|------|-----------|-------|---------|
| Rokland 8 dBi Omni | 8 dBi | N-Male, 45" fiberglass | $30-$35 each | [Rokland](https://store.rokland.com/products/meshtastic-compatible-6-dbi-n-female-omni-outdoor-915-mhz-antenna-kit-with-6-10-or-15-cable-choice) |
| Rokland 10 dBi Omni | 10 dBi | N-Male, extended | $50 each | [Rokland](https://store.rokland.com/collections/all/products/10-dbi-rokland-backcountry-rural-n-male-omni-outdoor-helium-915-mhz-antenna-45-for-rak-nebra-bobcat-sensecap-m1-hotspots) |
| Amazon 10 dBi LoRa | 10 dBi | Various | $40-$45 each | [Amazon](https://www.amazon.com/Meshnology-Meshtastic-10dBi-Development-Extension/dp/B0FJ5G1N3H) |

**Vertical Separation Requirements:**
- **18 inches (0.46 m)** vertical separation between the Meshtastic and Mesh Core antennas — enough to decouple the two omnis without needing the full 6-12 ft spacing some guides recommend, given the total mast height budget below
- Both should be vertically polarized omni-directional
- **Total mast height: 6 feet above the roof mount point.** Stack the two mesh antennas (and the optional TV antenna, see below) within that 6 ft envelope rather than maximizing height — this keeps wind load, guying and roof penetration requirements manageable on a flat roof
- A practical stacking order, bottom to top on a 6 ft mast: TV antenna (if used) → Mesh Core antenna → 18" gap → Meshtastic antenna at the top, so the highest-gain/most range-critical antenna gets the clearest sky view

---

## Digital TV Antenna (Optional Add-On)

Adding an over-the-air (OTA) digital TV antenna to the same mast is a common ask once you're already running coax up to the roof. Because it operates on VHF/UHF broadcast bands rather than 915MHz, it won't interfere with the Meshtastic radios as long as it's mounted with reasonable separation (the 18" spacing budget above is sufficient) and run on its own RG6 coax rather than sharing the LMR-400 feedline.

### Omnidirectional / Multi-Directional Options
Best when broadcast towers are spread across multiple directions from the roof.

| Item | Range | Price | Sources |
|------|-------|-------|---------|
| ClearStream 2MAX (omni-capable multi-directional) | ~60 miles | $70-$90 | [Antennas Direct](https://www.antennasdirect.com/) |
| RCA Outdoor Omnidirectional | ~70 miles | $60-$80 | [Amazon](https://www.amazon.com/s?k=rca+outdoor+omnidirectional+tv+antenna) |
| Channel Master Omni+ | ~60 miles | $80-$100 | [Channel Master](https://www.channelmaster.com/) |

### Directional (Yagi) Options
Higher gain toward a single direction — use when most/all broadcast towers are clustered behind one bearing.

| Item | Range | Price | Sources |
|------|-------|-------|---------|
| Channel Master CM-4228HD | ~80 miles | $90-$120 | [Channel Master](https://www.channelmaster.com/) |
| Antennas Direct DB8e | ~70 miles | $90-$110 | [Antennas Direct](https://www.antennasdirect.com/) |

### Amplified / Preamp Options
Worth adding if the coax run from mast to TV/tuner exceeds ~50 ft, or measured signal is weak — don't add gain you don't need, as over-amplifying can overload nearby tuners.

| Item | Notes | Price | Sources |
|------|-------|-------|---------|
| Channel Master CM-7777HD Preamp | Mast-mounted, adjustable gain | $40-$60 | [Channel Master](https://www.channelmaster.com/) |
| Antop Antenna In-Line Amplifier | Simple inline booster | $20-$35 | [Amazon](https://www.amazon.com/s?k=antop+tv+antenna+amplifier) |

**Notes:**
- Use **RG6 coax** (not LMR-400) for the TV antenna run — it's the standard impedance/connector for ATSC tuners and much cheaper per foot
- Ground the TV antenna mount the same as the mesh antennas (see [Weatherproofing Checklist](#weatherproofing-checklist))
- Use [rabbitears.info](https://www.rabbitears.info/) or a similar broadcast-tower lookup for your address before choosing omni vs. directional
- Mounting the TV antenna lowest on the mast keeps its (typically bulkier) wind load closer to the roof and out of the way of the mesh antennas' fresnel zone

---

## Power System

### DIY Configuration
| Item | Specs | Price | Sources |
|------|-------|-------|---------|
| 12V PoE Injector | Passive PoE, 12V output | $25-$35 | [Amazon](https://www.amazon.com/s?k=12v+poe+injector) |
| 12V 20Ah LiFePO4 Battery | BMS protected | $120-$160 | [Battle Born](https://battlebornbatteries.com/) |
| 100W Solar Panel | Monocrystalline + mount | $80-$120 | [Renogy](https://www.renogy.com/) |
| 10A MPPT Charge Controller | Temperature sensor included | $60-$90 | [HQST](https://hqstsolar.com/) |

### Integrated Solutions
| Item | Specs | Price | Sources |
|------|-------|-------|---------|
| RevoPower Solar PoE | 12V, battery included, weatherproof | $79 | [RevoPower](https://revopower.us/solar-poe-injector-with-battery) |
| LINOVISION 5-in-1 | 100W panel + 12V 40Ah + PoE | $250-$350 | [Linovision](https://linovision.com/products/solar-s1240p100poe-fully-integrated-5-in-1-smart-solar-power-supply-system-for-security-cameras-and-iot-device-support-3-ports-poe-output) |
| 12V 40Ah LiFePO4 | IP66 enclosure, cloud monitoring | $220-$260 | [Amazon](https://www.amazon.com/LINOVISION-Controller-Monitoring-Waterproof-Security/dp/B0DT9M568Z) |
| 200W Solar Panel | With pole mount kit | $170-$230 | [Off Grid Stores](https://offgridstores.com/) |

**Power Consumption Estimates:**

| Component | Typical Draw | Daily Consumption |
|-----------|--------------|-------------------|
| Heltec V3 (sleep mode) | 23μA | ~0.55mAh/day |
| Heltec V3 (transmitting) | 90-120mA | Variable |
| T-Beam 1W (idle) | ~90mA | ~2.16Ah/day |
| T-Beam 1W (transmit) | ~500mA | Variable |
| **Recommended Battery** | **12V 20-40Ah** | **2-3 days autonomy** |

**⚠️ Important Note:** Meshtastic radios typically run on 5V USB. PoE systems deliver 48V or 12V:
- Use a **PoE splitter** to convert to 5V USB for the radios
- Alternative: 12V PoE with DC-DC buck converter to 5V
- Some newer Meshtastic devices accept 9-36V directly (verify your hardware)

---

## Mounting & Cabling

### Mounting Hardware
| Item | Specs | Price | Sources |
|------|-------|-------|---------|
| Mast Mount Bracket | 1.75" mast capacity, swivel | $4-$15 each | [Data Alliance](https://www.data-alliance.net/pole-mounts) |
| Stainless Steel Mast Clamp | 4" pipe, 304 SS, heavy duty | $18-$25 each | [Walmart](https://www.walmart.com/c/kp/antenna-brackets-mounts) |
| Universal Clamp Set | Wall or roof mounting | $20-$30 | [Tessco](https://www.tessco.com/catalog/Infrastructure-Hardware/Structural-Support/Antenna-Mounts/c/30108) |

### Coaxial Cabling
| Item | Length | Price | Sources |
|------|--------|-------|---------|
| LMR-400 Custom Cut | — | $0.99-$1.30/ft | [eBay](https://www.ebay.com/itm/146372673136) |
| LMR-400 (4ft assembly) | 4 ft | ~$68 | [L-com](https://www.l-com.com/low-loss-n-male-sma-male-cable-assembly-using-lmr-400-coax-4-ft-times-microwave-components-lcca30289-ft4) |
| LMR-400 Bulk (1000ft) | Per foot | $1.32/ft | [Digi-Key](https://www.digikey.com/en/product-highlight/a/amphenol-times-microwave/lmr-400-coaxial-cable-connectors-and-accessories) |
| Weatherproof Ethernet | Cat8, shielded outdoor | $15-$25 | [Amazon](https://www.amazon.com/s?k=cat8+weatherproof+ethernet) |
| RG6 Coax (for TV antenna) | Per foot | $0.30-$0.50/ft | [Amazon](https://www.amazon.com/s?k=rg6+coax+outdoor+cable) |

### Connectors & Protection
| Item | Type | Price |
|------|------|-------|
| N-Male to SMA-Male Adapter | RF adapter | $6.97 |
| RF Surge Protector | N-type lightning arrestor | $15-$30 |
| Weatherproof Connectors | N-type bulkhead adapters | $8-$12 each |

---

## Installation Notes

### Weatherproofing Checklist
- [ ] N-type connectors for all outdoor connections
- [ ] Self-amalgamating tape on all joints
- [ ] Silicone sealant at penetration points
- [ ] UV-rated cable ties for strain relief
- [ ] Lightning arrestor if in strike-prone area

### Cable Loss Considerations (915MHz)

| Cable Type | Loss per 10ft @ 915MHz | Max Practical Length |
|------------|----------------------|---------------------|
| LMR-240 | ~3.0 dB | 10-15 ft |
| LMR-400 | ~1.5 dB | 25-30 ft |
| LMR-600 | ~0.8 dB | 40-50 ft |

**Recommendation:** Keep antenna-to-radio cable runs under 20ft to minimize signal degradation. With a 6 ft mast, an 8-10 ft LMR-400 run comfortably covers either mesh antenna position with slack for routing.

### Roof Mounting Best Practices

1. **Avoid penetrations** when possible — use weighted ballast mounts
2. **Ground the mast** to reduce lightning risk
3. **Leave maintenance access** — you'll need to adjust alignment occasionally
4. **Consider wind load** — a 6 ft mast carrying two (or three, with a TV antenna) antennas is a lighter, easier-to-guy load than a taller mast, but still plan for guy wires or a wider base plate on a flat roof
5. **Test before permanent install** — temporary setup helps optimize placement

### Firmware Configuration Tips

1. Enable **long-range mode** (SF12) for maximum distance
2. Disable **Bluetooth** to reduce power consumption
3. Set **TX power** appropriately (don't max out unless needed)
4. Configure **role** correctly (ROUTER vs CLIENT vs REPLAYER)
5. Enable **position sharing** at reasonable intervals to conserve battery

---

## Recommended Balanced Setup (~$332)

For most users, this configuration provides optimal performance-to-cost ratio:

| Qty | Item | Unit Price | Subtotal |
|-----|------|------------|----------|
| 2 | LILYGO T-Beam 1W | $60 | $120 |
| 1 | 3.5 dBi Antenna | $15 | $15 |
| 1 | 8 dBi High-Gain Antenna | $35 | $35 |
| 1 | RevoPower Solar PoE Kit | $79 | $79 |
| 2 | 15ft LMR-400 Cables | $45 total | $45 |
| 2 | Mast Mount Brackets | $15 each | $30 |
| 1 | Surge Protector | $20 | $20 |
| **TOTAL** | | | **~$344** |

**Why This Configuration:**
- ✓ Dual antennas allow comparison of gain impact
- ✓ 1W radios provide solid range without excessive power draw
- ✓ Integrated power system reduces installation complexity
- ✓ Room to upgrade (swap in larger battery/solar if needed)
- ✓ Fits within the 6 ft mast / 18" antenna-spacing envelope with room left over for an optional TV antenna

**Optional add-on:**

| Qty | Item | Unit Price | Subtotal |
|-----|------|------------|----------|
| 1 | ClearStream 2MAX Omni TV Antenna | $80 | $80 |
| 1 | RG6 Coax (25ft) | $10 | $10 |
| **TOTAL w/ TV antenna** | | | **~$434** |

---

## References & Further Reading

- [Meshtastic Hardware Documentation](https://meshtastic.org/docs/hardware/)
- [Meshtastic Antenna Selection Guide](https://meshtastic.org/docs/hardware/antennas/lora-antenna)
- [Meshtastic Power Management](https://meshtastic.org/docs/configuration/power/)
- [LoRa Range Calculator](https://www.loraserver.io/lora-range-calculator/)
- [RabbitEars — OTA broadcast tower lookup](https://www.rabbitears.info/)

---

## Document Metadata

| Field | Value |
|-------|-------|
| Created | January 2024 |
| Updated | August 2026 — 18" antenna spacing, 6 ft mast height, digital TV antenna options |
| Version | 1.1 |
| Author | Lumo (Proton AI Assistant) |
| Format | Markdown (.md) |
| License | Public Domain / CC0 |

---

*End of Document*
