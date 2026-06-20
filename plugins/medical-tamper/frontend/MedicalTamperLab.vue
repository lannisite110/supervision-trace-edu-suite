<script setup lang="ts">
import { computed, ref } from 'vue'
import { useLabSimulate } from '../../frontend/shared/useLabSimulate'
import { parseHints, hintBool } from '../../frontend/shared/parseHints'

const recordId = ref('DEMO-RECORD-001')
const submittedHash = ref('sha256:demo-medical-hash-abc123')
const { loading, error, result, taskStatus, taskReport, runSimulate, parseEvaluation } =
  useLabSimulate('edu.cn.trace.medical')

const evaluation = computed(() => parseEvaluation(result.value?.evaluation))
const hints = computed(() => parseHints(evaluation.value?.audit_hints))
const ruleTampered = computed(() => hintBool(hints.value, 'tamper_detected'))

const storedHash = 'sha256:demo-medical-hash-abc123'
const isTampered = computed(() => submittedHash.value !== storedHash)

const hashChain = [
  { version: 'v1', hash: 'sha256:demo-medical-hash-abc123', action: '初始病历存证', time: '2026-05-01 09:00' },
  { version: 'v2', hash: 'sha256:demo-medical-hash-def456', action: '合法追加备注（教学）', time: '2026-05-15 14:20' },
  { version: '链上当前', hash: storedHash, action: '有效存证版本', time: '2026-06-01 08:00' },
]

const fabricSandbox = {
  chainId: 'fabric-local',
  channel: 'edu-cn-trace-sandbox',
  org: 'OrgEduDemo',
}

function submit() {
  runSimulate(
    '病历哈希防篡改存证演示',
    { record_id: recordId.value, submitted_hash: submittedHash.value },
    { taskType: 'CN_MEDICAL_TAMPER_DEMO' },
  )
}
</script>

<template>
  <section class="lab-panel">
    <header class="lab-header">
      <img src="/assets/icon.png" alt="" width="32" height="32" />
      <div>
        <h1>医疗防篡改存证 Demo</h1>
        <p class="muted">病历哈希存证数据结构演示 · 不对接真实 EMR/HIS</p>
      </div>
    </header>

    <div v-if="evaluation" class="eval-card">
      <h2>规则评估</h2>
      <p class="ok">✓ 合规通过 · {{ evaluation.recommended_language }}</p>
      <p v-if="hints.record_id"><strong>记录:</strong> {{ hints.record_id }}</p>
      <p v-if="evaluation">
        <strong>规则篡改检测:</strong>
        {{ ruleTampered ? '检测到不一致' : '哈希一致' }}
      </p>
    </div>

    <div class="lab-grid">
      <div class="card">
        <h2>存证记录</h2>
        <label>
          记录 ID
          <input v-model="recordId" />
        </label>
        <p><strong>链上哈希:</strong> <code>{{ storedHash }}</code></p>
        <label>
          待校验哈希
          <input v-model="submittedHash" />
        </label>
        <button :disabled="loading" @click="submit">
          {{ loading ? '检测中…' : '提交篡改检测实验' }}
        </button>
        <p v-if="taskStatus" class="status">任务状态: {{ taskStatus }}</p>
        <p v-if="error" class="error">{{ error }}</p>
      </div>

      <div class="card" :class="{ tampered: isTampered }">
        <h2>检测结果</h2>
        <p v-if="!isTampered" class="ok">✓ 哈希一致，未检测到篡改</p>
        <p v-else class="warn">⚠ 哈希不一致，触发异常修改检测算法（教学演示）</p>
        <p class="muted">等保概念教学参考 — 非等保认证宣称</p>
        <div class="compare">
          <div>
            <small>链上</small>
            <code>{{ storedHash.slice(0, 24) }}…</code>
          </div>
          <div>
            <small>提交</small>
            <code>{{ submittedHash.slice(0, 24) }}…</code>
          </div>
        </div>
      </div>

      <div class="card">
        <h2>版本哈希链</h2>
        <ol class="hash-chain">
          <li v-for="item in hashChain" :key="item.version">
            <strong>{{ item.version }}</strong> — {{ item.action }}
            <br /><code>{{ item.hash }}</code>
            <br /><small>{{ item.time }}</small>
          </li>
        </ol>
        <p class="muted">Fabric {{ fabricSandbox.channel }} · {{ fabricSandbox.org }}</p>
      </div>
    </div>

    <pre v-if="taskReport" class="result">{{ JSON.stringify(taskReport, null, 2) }}</pre>
    <pre v-else-if="result" class="result">{{ JSON.stringify(result, null, 2) }}</pre>
  </section>
</template>

<style scoped>
.lab-panel { padding: 1rem; }
.eval-card { background: #f0fdf4; border: 1px solid #bbf7d0; border-radius: 8px; padding: 1rem; margin-bottom: 1rem; }
.eval-card h2 { margin: 0 0 0.5rem; font-size: 1rem; }
.eval-card .ok { color: #15803d; margin: 0 0 0.5rem; }
.lab-header { display: flex; gap: 0.75rem; align-items: center; margin-bottom: 1rem; }
.lab-grid { display: grid; grid-template-columns: repeat(auto-fit, minmax(280px, 1fr)); gap: 1rem; }
.card { background: #f8fafc; border: 1px solid #e2e8f0; border-radius: 8px; padding: 1rem; }
.card.tampered { border-color: #f59e0b; background: #fffbeb; }
.muted { color: #64748b; font-size: 0.875rem; }
.status { color: #2563eb; font-size: 0.875rem; margin-top: 0.5rem; }
code { font-family: monospace; font-size: 0.8rem; color: #0f766e; }
.ok { color: #16a34a; font-weight: 600; }
.warn { color: #d97706; font-weight: 600; }
.error { color: #dc2626; }
.result { margin-top: 1rem; background: #1e293b; color: #e2e8f0; padding: 1rem; border-radius: 8px; overflow: auto; font-size: 0.8rem; }
input { display: block; width: 100%; margin: 0.5rem 0; padding: 0.5rem; border: 1px solid #cbd5e1; border-radius: 4px; }
button { margin-top: 0.5rem; padding: 0.5rem 1rem; background: #2563eb; color: white; border: none; border-radius: 4px; cursor: pointer; }
button:disabled { opacity: 0.6; cursor: not-allowed; }
.hash-chain { padding-left: 1.25rem; }
.hash-chain li { margin-bottom: 0.75rem; }
.compare { display: grid; grid-template-columns: 1fr 1fr; gap: 0.5rem; margin-top: 0.75rem; font-size: 0.8rem; }
.compare small { color: #64748b; }
</style>
