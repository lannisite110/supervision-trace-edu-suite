"""慈善流水存证教学规则 — 模拟募捐哈希上链。"""

from __future__ import annotations

from plugins.rules._common import RuleInput, RuleOutput, check_compliance, fabric_hints


def evaluate(inp: RuleInput) -> RuleOutput:
    blocked = check_compliance(inp)
    if blocked:
        return blocked

    text = f"{inp.user_prompt} {inp.params.get('scenario', '')}".lower()
    if any(kw in text for kw in ("官方支付", "真实募捐", "微信支付对接", "支付宝对接")):
        return RuleOutput(
            "",
            "",
            [],
            False,
            "不得对接官方募捐支付，仅允许模拟流水哈希存证",
        )

    campaign_id = inp.params.get("campaign_id", "DEMO-CAMPAIGN-001")
    entry_count = inp.params.get("entry_count", 2)
    ledger_hash = str(inp.params.get("ledger_hash", "ledger-merkle-root-demo-9b2c"))

    return RuleOutput(
        recommended_template="plugins/charity-ledger/chaincode/charity_ledger.go",
        recommended_language="go",
        audit_hints=fabric_hints([
            f"campaign_id={campaign_id}",
            f"entry_count={entry_count}",
            f"ledger_hash={ledger_hash}",
            "simulated_donation_ledger",
            "audit_trail_hash_only",
            "CN_CHARITY_LEDGER_DEMO",
        ]),
        compliance_passed=True,
    )
