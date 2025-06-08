# GnoGas TODO

## 🔧 Phase 1 – Static Analysis

- [ ] Parse Gno contracts to extract call graphs
- [ ] Identify loops, recursion, and nested calls
- [ ] Estimate gas cost for state ops

## 🔍 Phase 2 – Profiler

- [ ] Wrap contract calls and log gas used
- [ ] CLI tool to show gas per function
- [ ] Compare static vs runtime cost

## ⚠️ Phase 3 – Tooling

- [ ] VS Code plugin: gas hints & warnings
- [ ] Compiler hints during build
- [ ] Auto-fix engine (redundant code, consts, etc.)
