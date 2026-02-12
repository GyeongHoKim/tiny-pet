# Tiny Pet — Hardware (Power Supply)

KiCad project for the **3.7V Li-ion battery power supply** used when running the desk pet on battery (no USB cable).

## Block diagram

```
3.7V Li-ion 1S → Protection (1S) → Boost (3.x V → 5V) → 5V rail → [optional] LDO → 3.3V rail
                    (B+/B-/P+/P-)        (e.g. MT3608)      (J3)       (e.g. AMS1117-3.3)   (J4)
```

- **5V rail (J3):** Arduino Uno/Nano, L298N logic (VSS) and motor (VM), HC-SR04, 5V-type SSD1306. Connect to board 5V/GND or USB replacement.
- **3.3V rail (J4, optional):** Blue Pill, 3.3V OLED/IR. Fed from 5V via LDO so sensors see stable voltage even when the battery sags.

## KiCad project

| File | Description |
|------|--------------|
| `tiny-pet-power.kicad_pro` | KiCad 7/8/9 project |
| `tiny-pet-power.kicad_sch` | Schematic: J1 Battery, J2 Protection out, J3 5V out, J4 3.3V out |
| `tiny-pet-power.kicad_pcb` | PCB (minimal outline; place boost/LDO and connectors as needed) |

Open `tiny-pet-power.kicad_pro` in KiCad to edit the schematic or PCB. The schematic uses only KiCad built-in symbols (Connector_Generic:Conn_01x02, power:VCC, power:GND).

Use only KiCad default (global) symbol library table; do not add a project-specific table. If symbols show as ??, set up the global table in **Preferences → Manage Symbol Libraries** (e.g. from KiCad default template).

## Verification

- **Project structure:** `mcp_kicad_validate_project` with `project_path`: `hardware/tiny-pet-power.kicad_pro` (or full path).
- **PCB DRC:** From repo root,  
  `kicad-cli pcb drc hardware/tiny-pet-power.kicad_pcb`  
  (or use KiCad MCP `run_drc_check` with the same project path).
- **Schematic ERC:** Run ERC from KiCad GUI (Schematic Editor → Inspect → Electrical Rules Check). CLI `kicad-cli sch erc` may require the same library setup as the GUI.

## Parts

Battery power uses the same items as the main [README](../README.md) Parts list: 3.7V Li-ion 1S, 1S protection module, 5V boost converter, and optionally 5V→3.3V LDO. Input/output capacitors and connectors (e.g. JST-PH for battery) as needed.
