#!/usr/bin/env bash
# Phase 0 验收：注册 + 3 插件 curl 冒烟
set -euo pipefail
ROOT="$(cd "$(dirname "$0")/.." && pwd)"
CORE="${CORE:-$ROOT/../web3-edu-platform-core}"

echo "==> register plugins"
cd "$CORE" && make register-plugins PLUGINS_DIR="$ROOT/.."

echo "==> trace plugin smoke (needs gateway :8080 + rule-engine :8081 + scheduler :8082)"
for pid in edu.cn.trace.food edu.cn.trace.medical edu.cn.trace.charity; do
  code=$(curl -sf -o /tmp/trace-smoke.json -w '%{http_code}' \
    -X POST "http://127.0.0.1:8080/api/v1/labs/${pid}/simulate" \
    -H 'Content-Type: application/json' \
    -d '{"user_prompt":"trace phase0","params":{},"allowed_chain_ids":["fabric-local"]}' || echo "000")
  if [ "$code" = "202" ] || [ "$code" = "200" ]; then
    echo "  OK $pid HTTP $code"
  else
    echo "  SKIP/FAIL $pid HTTP $code (start backend for full smoke)"
    exit 1
  fi
done

echo "==> trace phase0 smoke PASSED"