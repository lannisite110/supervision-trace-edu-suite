"""医疗防篡改存证教学规则 — 病历哈希比对演示。"""

from __future__ import annotations

from plugins.rules._common import RuleInput, RuleOutput, check_compliance, fabric_hints


def evaluate(inp: RuleInput) -> RuleOutput:
    blocked = check_compliance(inp)
    if blocked:
        return blocked

    record_id = inp.params.get("record_id", "DEMO-RECORD-001")
    stored_hash = "sha256:demo-medical-hash-abc123"
    submitted = str(inp.params.get("submitted_hash", stored_hash))
    tampered = submitted != stored_hash
    text = f"{inp.user_prompt} {inp.params.get('scenario', '')}".lower()
    if any(kw in text for kw in ("等保三级", "三级等保", "生产认证")):
        return RuleOutput(
            "",
            "",
            [],
            False,
            "不得宣称等保认证，仅允许等保概念教学参考",
        )

    return RuleOutput(
        recommended_template="plugins/medical-tamper/chaincode/medical_tamper.go",
        recommended_language="go",
        audit_hints=fabric_hints([
            f"record_id={record_id}",
            f"stored_hash={stored_hash}",
            f"submitted_hash={submitted}",
            f"tamper_detected={'true' if tampered else 'false'}",
            "hash_tamper_detection_demo",
            "no_real_emr_integration",
            "CN_MEDICAL_TAMPER_DEMO",
        ]),
        compliance_passed=True,
    )
