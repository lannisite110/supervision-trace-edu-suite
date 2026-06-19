"""食品溯源教学规则 — 批次/Merkle 证明推荐。"""

from __future__ import annotations

from plugins.rules._common import (
    FABRIC_CHAIN_ID,
    RuleInput,
    RuleOutput,
    check_compliance,
    fabric_hints,
)


def evaluate(inp: RuleInput) -> RuleOutput:
    blocked = check_compliance(inp)
    if blocked:
        return blocked

    batch_id = inp.params.get("batch_id", "DEMO-BATCH-001")
    product = inp.params.get("product_name", "有机蔬菜礼盒")

    return RuleOutput(
        recommended_template="plugins/food-trace/chaincode/food_trace.go",
        recommended_language="go",
        audit_hints=fabric_hints([
            f"batch_id={batch_id}",
            f"product={product}",
            "merkle_proof_enabled",
            "CN_FOOD_TRACE_SIM",
        ]),
        compliance_passed=True,
    )
