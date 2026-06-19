#!/usr/bin/env bash
# Agent-2 联合调试冒烟 — 3 个国内民生溯源插件
set -euo pipefail

ROOT="$(cd "$(dirname "$0")/.." && pwd)"
CORE="${CORE_ROOT:-$(cd "$ROOT/../web3-edu-platform-core" && pwd)}"

if [ ! -d "$CORE" ]; then
  echo "ERROR: web3-edu-platform-core not found at $CORE"
  exit 1
fi

export CORE_ROOT="$CORE"
export GATEWAY_PORT="${GATEWAY_PORT:-8080}"
export RULE_ENGINE_PORT="${RULE_ENGINE_PORT:-8081}"
export SCHEDULER_PORT="${SCHEDULER_PORT:-8082}"

cd "$CORE"
PYTHON="${CORE}/.venv/bin/python"
[ -x "$PYTHON" ] || { python3 -m venv "${CORE}/.venv" && "${CORE}/.venv/bin/pip" install -q -r rule-engine-py/requirements.txt pyyaml; }

echo "==> register plugins"
"$PYTHON" ci/register-plugins.py "$(dirname "$ROOT")"

echo "==> validate all cn-trace manifests"
for m in "$ROOT"/plugins/*/plugin.manifest.yaml; do
  MANIFEST="$m" bash ci/compliance/validate-plugin.sh
done

echo "==> compliance-check supervision-trace-edu-suite"
bash ci/compliance/compliance-check.sh "$ROOT"

STARTED=0
cleanup() {
  if [ "$STARTED" = "1" ]; then
    fuser -k "${GATEWAY_PORT}/tcp" "${RULE_ENGINE_PORT}/tcp" "${SCHEDULER_PORT}/tcp" 2>/dev/null || true
    kill $(jobs -p) 2>/dev/null || true
  fi
}
trap cleanup EXIT

if ! curl -sf "http://127.0.0.1:${RULE_ENGINE_PORT}/health" >/dev/null 2>&1; then
  echo "==> starting rule-engine + scheduler + gateway"
  STARTED=1
  (cd rule-engine-py && "$PYTHON" main.py) &
  sleep 1
  (cd control-plane-go && CORE_ROOT="$CORE" SCHEDULER_PORT="$SCHEDULER_PORT" go run ./cmd/scheduler) &
  sleep 3
  (cd api-gateway-go && CORE_ROOT="$CORE" GATEWAY_PORT="$GATEWAY_PORT" go run ./cmd/gateway) &
  sleep 5
fi

wait_health() {
  local url="$1" name="$2"
  for i in 1 2 3 4 5 6 7 8 9 10; do
    if curl -sf "$url" >/dev/null 2>&1; then
      return 0
    fi
    sleep 2
  done
  echo "ERROR: $name not healthy at $url"
  return 1
}

wait_health "http://127.0.0.1:${RULE_ENGINE_PORT}/health" "rule-engine"
curl -sf "http://127.0.0.1:${RULE_ENGINE_PORT}/health" | grep -q ok
wait_health "http://127.0.0.1:${SCHEDULER_PORT}/health" "scheduler"
curl -sf "http://127.0.0.1:${SCHEDULER_PORT}/health" | grep -q ok
wait_health "http://127.0.0.1:${GATEWAY_PORT}/health" "gateway"
curl -sf "http://127.0.0.1:${GATEWAY_PORT}/health" | grep -q ok

PLUGINS=(
  "edu.cn.trace.food|CN_FOOD_TRACE_SIM|batch_id"
  "edu.cn.trace.medical|CN_MEDICAL_TAMPER_DEMO|record_id"
  "edu.cn.trace.charity|CN_CHARITY_LEDGER_DEMO|campaign_id"
)

for entry in "${PLUGINS[@]}"; do
  IFS='|' read -r PID TASK PARAM <<<"$entry"
  echo "==> rule-engine evaluate: $PID"
  OUT=$("$PYTHON" - <<PY
import json, urllib.request
req = urllib.request.Request(
    "http://127.0.0.1:${RULE_ENGINE_PORT}/evaluate",
    data=json.dumps({
        "plugin_id": "${PID}",
        "user_prompt": "教学演示",
        "params": {},
        "allowed_chain_ids": ["fabric-local"],
    }).encode(),
    headers={"Content-Type": "application/json"},
    method="POST",
)
with urllib.request.urlopen(req) as r:
    print(r.read().decode())
PY
)
  echo "$OUT" | grep -q '"compliance_passed":true'
  echo "$OUT" | grep -q go

  echo "==> gateway simulate: $PID"
  RESP=$(curl -sf -X POST "http://127.0.0.1:${GATEWAY_PORT}/api/v1/labs/${PID}/simulate" \
    -H 'Content-Type: application/json' \
    -d "{\"params\":{},\"allowed_chain_ids\":[\"fabric-local\"]}")
  echo "$RESP" | grep -q "$PID"
  echo "$RESP" | grep -q ns-domain-cn
  echo "$RESP" | grep -q "$TASK"

  TASK_ID=$(echo "$RESP" | "$PYTHON" -c "import sys,json; print(json.load(sys.stdin)['task']['id'])")
  echo "==> status/report: $TASK_ID"
  for i in 1 2 3 4 5 6 7 8 9 10; do
    STATUS_JSON=$(curl -sf "http://127.0.0.1:${GATEWAY_PORT}/api/v1/labs/${PID}/status/${TASK_ID}")
    echo "$STATUS_JSON" | grep -q completed && break
    sleep 0.3
  done
  echo "$STATUS_JSON" | grep -q completed
  curl -sf "http://127.0.0.1:${GATEWAY_PORT}/api/v1/labs/${PID}/report/${TASK_ID}" | grep -q ns-domain-cn
done

echo "==> mainnet blocked (food-trace)"
CODE=$(curl -s -o /dev/null -w '%{http_code}' -X POST \
  "http://127.0.0.1:${GATEWAY_PORT}/api/v1/labs/edu.cn.trace.food/simulate" \
  -H 'Content-Type: application/json' \
  -d '{"params":{"target_network":"mainnet"}}')
test "$CODE" = "403"

echo "==> Agent-2 joint-debug smoke PASSED"
