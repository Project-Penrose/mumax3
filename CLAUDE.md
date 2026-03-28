# mumax3 — CLAUDE.md

GPU-accelerated micromagnetics simulation engine (Go + CUDA). Used in Project-Penrose for spin-wave dispersion, SOT-MRAM write physics validation, Suhl instability characterization, and GGG/YIG material parameter extraction campaigns that feed SYNTH_YIG_PARAMS_V1 and SPEC-POINCARE-023/024.

## Repository Layout

| Path | Purpose |
|------|---------|
| `engine/` | Core simulation engine — LLG integrator, exchange, dipolar, Zeeman, DMI |
| `cuda/` | CUDA kernel implementations — GPU acceleration |
| `cmd/` | Command-line entrypoints — `mumax3`, `mumax3-convert` |
| `mag/` | Magnetic material definitions — GGG, YIG, CoFeB parameter presets |
| `gui/` | Web-based visualization GUI |
| `httpfs/` | HTTP filesystem for distributed job result access |
| `data/` | Simulation output data handlers |
| `doc/` | Documentation and example scripts |
| `oommf/` | OOMMF compatibility layer |
| `tests/` | Integration test suite |
| `bench/` | Performance benchmarks |
| `dump/` | Checkpoint/restart file handlers |

## Common Tasks

**Run a spin-wave dispersion simulation:**
```bash
mumax3 examples/spinwave_dispersion.mx3
```

**Run a Suhl instability threshold characterization:**
```bash
mumax3 examples/suhl_threshold.mx3
# Output P_th value must be < 0.8 mW per ADR-126 Suhl safety
```

**Convert simulation output:**
```bash
mumax3-convert -f csv output.ovf
```

**Build from source:**
```bash
# Requires NVIDIA driver, Go, CUDA (≤12.9), GCC
make build
```

**Run tests:**
```bash
go test ./...
```

## MagnonOS-Specific Usage

mumax3 is invoked by the `SimulationCampaignManager` agent (port 9360) for:
- **GGG baseline characterization** (ADR-076): Ms, α, D_ex extraction for GGG/YIG heterostructures
- **SOT-MRAM write physics** (SPEC-MAGNONIC-040): J_c0, WER<10⁻⁶, E_write=50fJ validation
- **Suhl instability verification** (ADR-126): P_th=0.64 mW hardware interlock calibration
- **Spin-wave dispersion** (SPEC-POINCARE-023): Kalinikos-Slavin group velocity v_g extraction

Simulation results are validated by the `SimulationResultValidator` agent (port 9252) against penrose-core-spec/v1 physical bounds before being accepted into the Mouseion material parameter registry.

## Key Invariants

- **Suhl safety:** Any simulation producing P_th results must compare against the 0.64 mW threshold (ADR-126). Results showing P_th < 0.64 mW trigger an immediate alert to the `PenroseSentinel` agent.
- **Result fingerprinting:** Simulation outputs must carry a SHA-256 fingerprint before being forwarded to `DeviceParameterExtractor`. Format: `SHA-256("{device_id}:{param:.4f}:{freq_ghz:.4f}")`.
- **Physical bounds validation:** All extracted parameters must pass penrose-core-spec/v1 bounds before entering the rolling 500-record deque in `MagnonDispersionAnalyzer`.
- **No public registries.** mumax3 is a private fork. Do not publish to pkg.go.dev or any public registry.

## Integration Points

| Counterpart | Direction | Purpose |
|-------------|-----------|---------|
| `agents/simulation_campaign_manager` | Orchestrator | Dispatches mumax3 jobs and collects results |
| `agents/simulation_result_validator` | Validator | Validates output against penrose-core-spec/v1 |
| `agents/magnon_dispersion_analyzer` | Consumer | Ingests spin-wave dispersion output |
| Poincaré (SYNTH_YIG_PARAMS_V1) | Upstream spec | Normative Ms/α/D_ex parameters for simulation config |
| Volterra (MCEL-1.0 characterization) | Downstream consumer | mumax3 spin-wave data feeds MCEL VNA characterization |

## Upstream Reference

Based on: https://mumax.github.io (mumax³ v3.x)
Paper: http://scitation.aip.org/content/aip/journal/adva/4/10/10.1063/1.4899186

MuMax2 is superseded — use mumax3 or mumaxplus.
