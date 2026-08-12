# VPS ProtonVPN WireGuard Setup

Routes WireGuard client traffic through ProtonVPN on the VPS, while keeping VPS management traffic (SSH, nginx) unaffected.

## Architecture

```
Clients → wg0 (VPS WireGuard server) → protonvpn (WireGuard client) → ProtonVPN → internet
VPS management traffic (SSH, nginx) → internet directly (unaffected)
```

The gluetun/torrent machine's traffic double-tunnels through ProtonVPN (gluetun's own ProtonVPN + VPS ProtonVPN). This is a known performance tradeoff — see [Gluetun Note](#gluetun-note) below.

---

## Prerequisites

### 1. Get ProtonVPN WireGuard credentials

Log into the ProtonVPN portal → **Downloads → WireGuard configuration** → generate a config for a Linux device. You will need these values from the downloaded file:

| Field | Location in ProtonVPN config |
|---|---|
| `PrivateKey` | `[Interface]` section |
| `Address` | `[Interface]` section |
| `PublicKey` | `[Peer]` section |
| `Endpoint` | `[Peer]` section |

### 2. Enable IP forwarding

```bash
sysctl net.ipv4.ip_forward
# Should return net.ipv4.ip_forward = 1
```

If not enabled:

```bash
echo "net.ipv4.ip_forward = 1" >> /etc/sysctl.conf
sysctl -p
```

---

## Config File

Save as `/etc/wireguard/protonvpn.conf`:

```ini
[Interface]
PrivateKey = <PrivateKey from ProtonVPN config>
Address = <Address from ProtonVPN config>
Table = off

PostUp = ip route add table 200 default dev %i
PostUp = ip rule add from 10.1.0.0/24 lookup 200 priority 100
PostUp = iptables -t nat -A POSTROUTING -s 10.1.0.0/24 -o %i -j MASQUERADE
PostUp = iptables -A FORWARD -i wg0 -o %i -j ACCEPT
PostUp = iptables -A FORWARD -i %i -o wg0 -m state --state RELATED,ESTABLISHED -j ACCEPT

PreDown = ip rule del from 10.1.0.0/24 lookup 200 priority 100
PreDown = ip route del table 200 default dev %i
PreDown = iptables -t nat -D POSTROUTING -s 10.1.0.0/24 -o %i -j MASQUERADE
PreDown = iptables -D FORWARD -i wg0 -o %i -j ACCEPT
PreDown = iptables -D FORWARD -i %i -o wg0 -m state --state RELATED,ESTABLISHED -j ACCEPT

[Peer]
PublicKey = <PublicKey from ProtonVPN config>
Endpoint = <Endpoint from ProtonVPN config>
AllowedIPs = 0.0.0.0/0
PersistentKeepalive = 25
```

> **Adjust `10.1.0.0/24`** if your WireGuard client subnet differs. Check your existing server-side `wg0.conf` for the address range assigned to peers.

### Why `Table = off`

Without this, WireGuard automatically adds a default route for `AllowedIPs = 0.0.0.0/0` to the main routing table. That would send all VPS traffic — including SSH — through ProtonVPN, potentially locking you out.

`Table = off` disables automatic route management. The PostUp rules then manually add routes only to table `200`, which applies exclusively to traffic sourced from WireGuard clients.

### How the routing works

- Routing table `200` has a single default route pointing at the ProtonVPN interface
- `ip rule` directs any packet sourced from `10.1.0.0/24` (your WireGuard clients) to use table `200`
- `MASQUERADE` rewrites the source IP on outbound packets so ProtonVPN sees the VPS's assigned ProtonVPN address, not client IPs
- Return traffic is allowed back through the `RELATED,ESTABLISHED` forward rule

---

## Enable and Start

```bash
# Start now
systemctl start wg-quick@protonvpn

# Enable on boot
systemctl enable wg-quick@protonvpn

# Check status
systemctl status wg-quick@protonvpn

# Verify routing table 200 is populated after start
ip route show table 200
ip rule show
```

---

## Verify It's Working

From a connected WireGuard client, check that the exit IP is a ProtonVPN address:

```bash
curl https://api64.ipify.org
# Should return a ProtonVPN exit IP, not your VPS IP
```

---

## Gluetun Note

The machine running Transmission in a gluetun container will double-tunnel through ProtonVPN:

```
Transmission → gluetun (ProtonVPN) → wg0 (VPS) → protonvpn (VPS) → ProtonVPN servers → internet
```

This means torrent traffic traverses three tunnels (gluetun VPN + WireGuard + VPS ProtonVPN) and exits through ProtonVPN twice. It works but carries a latency and throughput penalty.

To avoid this, gluetun would need its own separate WireGuard peer on the VPS with a distinct IP, allowing a bypass routing rule to exclude it from table `200`. This requires either:
- Running gluetun as a WireGuard client to your VPS instead of ProtonVPN (loses ProtonVPN protection for torrents)
- Accepting the current performance tradeoff

For most torrent use cases the tradeoff is acceptable since throughput matters more than latency for bulk transfers, and ProtonVPN's speeds are generally sufficient.

---

## ProtonVPN Account Consideration

This setup uses two simultaneous ProtonVPN connections: one from the VPS and one from the gluetun container. Verify your plan supports multiple simultaneous connections (Plus: 10 devices, Unlimited: higher).
