"""共享合规检查与 Fabric 沙箱常量。"""

from __future__ import annotations

from dataclasses import dataclass
from typing import Any

FABRIC_CHANNEL = "edu-cn-trace-sandbox"
FABRIC_ORG = "OrgEduDemo"
FABRIC_CHAIN_ID = "fabric-local"

BLOCKED_MAINNET_IDS = {1, 56, 137, 42161, 10, 8453}

FORBIDDEN_PARAMS = {
    "target_network": {"mainnet"},
    "integration": {"his", "emr", "real_hospital", "official_payment"},
}


@dataclass
class RuleInput:
    user_prompt: str
    params: dict[str, Any]
    allowed_chain_ids: list


@dataclass
class RuleOutput:
    recommended_template: str
    recommended_language: str
    audit_hints: list[str]
    compliance_passed: bool
    rejection_reason: str | None = None


def check_compliance(inp: RuleInput) -> RuleOutput | None:
    if inp.params.get("target_network") == "mainnet":
        return RuleOutput("", "", [], False, "mainnet forbidden")

    integration = str(inp.params.get("integration", "")).lower()
    if integration in FORBIDDEN_PARAMS["integration"]:
        return RuleOutput("", "", [], False, f"real-system integration forbidden: {integration}")

    for cid in inp.allowed_chain_ids:
        if isinstance(cid, int) and cid in BLOCKED_MAINNET_IDS:
            return RuleOutput("", "", [], False, f"mainnet chainId {cid} blocked")

    return None


def fabric_hints(extra: list[str] | None = None) -> list[str]:
    hints = [
        f"chain_id={FABRIC_CHAIN_ID}",
        f"channel={FABRIC_CHANNEL}",
        f"org={FABRIC_ORG}",
        "fictional-data-only",
        "fabric-sandbox",
    ]
    if extra:
        hints.extend(extra)
    return hints
