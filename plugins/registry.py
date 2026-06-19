"""Registry shim — 子库 plugins 包被加入 PYTHONPATH 时转发至主库 rule-engine registry。"""

from __future__ import annotations

import importlib.util
import sys
from pathlib import Path

_CORE_REG = (
    Path(__file__).resolve().parents[2]
    / "web3-edu-platform-core"
    / "rule-engine-py"
    / "plugins"
    / "registry.py"
)

_spec = importlib.util.spec_from_file_location("_core_plugins_registry", _CORE_REG)
_mod = importlib.util.module_from_spec(_spec)
assert _spec.loader is not None
sys.modules[_spec.name] = _mod
_spec.loader.exec_module(_mod)

RuleInput = _mod.RuleInput
RuleOutput = _mod.RuleOutput
load_evaluator = _mod.load_evaluator
run_plugin = _mod.run_plugin
