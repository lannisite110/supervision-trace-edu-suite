<script setup lang="ts">
import { computed, ref } from 'vue'
import { useLabSimulate } from '../../frontend/shared/useLabSimulate'
import { parseHints } from '../../frontend/shared/parseHints'

const campaignId = ref('DEMO-CAMPAIGN-001')
const { loading, error, result, taskStatus, taskReport, runSimulate, parseEvaluation } =
  useLabSimulate('edu.cn.trace.charity')

const evaluation = computed(() => parseEvaluation(result.value?.evaluation))
const hints = computed(() => parseHints(evaluation.value?.audit_hints))

const demoEntries = [
  { id: 'DEMO-ENTRY-001', donor: '虚构捐赠者A', amount: 100, hash: 'sha256:donation-demo-001', time: '2026-06-01 12:00' },
  { id: 'DEMO-ENTRY-002', donor: '虚构捐赠者B', amount: 50, hash: 'sha256:donation-demo-002', time: '2026-06-02 15:30' },
]

const fieldPolicy = [
  { field: '捐赠金额', visibility: '公开', onChain: '¥100', note: '汇总进活动账本' },
  { field: '捐赠时间', visibility: '公开', onChain: '2026-06-01T12:00:00Z', note: '审计轨迹' },
  { field: '捐赠者代号', visibility: '公开', onChain: '虚构捐赠者A', note: '非真实 PII' },
  { field: '捐赠者手机号', visibility: '隐私', onChain: 'sha256:redacted…demo', note: '仅 hash 上链' },
  { field: '支付渠道', visibility: '禁止', onChain: '—', note: '不对接微信/支付宝' },
]

const ledgerHash = 'ledger-merkle-root-demo-9b2c'
const totalAmount = demoEntries.reduce((s, e) => s + e.amount, 0)

function submit() {
  runSimulate(
    '模拟募捐流水哈希上链',
    { campaign_id: campaignId.value, entry_count: demoEntries.length, ledger_hash: ledgerHash },
    { taskType: 'CN_CHARITY_LEDGER_DEMO' },
  )
}
</script>

<template>
  <section class="lab-panel">
    <header class="lab-header">
      <img src="/assets/icon.png" alt="" width="32" height="32" />
      <div>
        <h1>慈善流水存证 Demo</h1>
        <p class="muted">模拟募捐流水哈希上链 · 不对接官方支付渠道</p>
      </div>
    </header>

    <div v-if="evaluation" class="eval-card">
      <h2>规则评估</h2>
      <p class="ok">✓ 合规通过 · {{ evaluation.recommended_language }}</p>
      <p v-if="hints.campaign_id"><strong>活动:</strong> {{ hints.campaign_id }}</p>
      <p v-if="hints.entry_count"><strong>流水笔数:</strong> {{ hints.entry_count }}</p>
      <p v-if="hints.channel"><strong>通道:</strong> {{ hints.channel }} · {{ hints.org }}</p>
    </div>

    <div class="lab-grid">
      <div class="card">
        <h2>募捐活动</h2>
        <label>
          活动 ID
          <input v-model="campaignId" />
        </label>
        <p><strong>流水汇总哈希:</strong> <code>{{ ledgerHash }}</code></p>
        <p><strong>累计金额（虚构）:</strong> ¥{{ totalAmount }}</p>
        <button :disabled="loading" @click="submit">
          {{ loading ? '提交中…' : '提交存证实验' }}
        </button>
        <p v-if="taskStatus" class="status">任务状态: {{ taskStatus }}</p>
        <p v-if="error" class="error">{{ error }}</p>
      </div>

      <div class="card">
        <h2>捐赠流水</h2>
        <table>
          <thead>
            <tr><th>捐赠者</th><th>金额</th><th>哈希</th><th>时间</th></tr>
          </thead>
          <tbody>
            <tr v-for="e in demoEntries" :key="e.id">
              <td>{{ e.donor }}</td>
              <td>¥{{ e.amount }}</td>
              <td><code>{{ e.hash }}</code></td>
              <td>{{ e.time }}</td>
            </tr>
          </tbody>
        </table>
      </div>

      <div class="card">
        <h2>公开 vs 隐私字段</h2>
        <table class="policy-table">
          <thead>
            <tr><th>字段</th><th>策略</th><th>链上示例</th></tr>
          </thead>
          <tbody>
            <tr v-for="row in fieldPolicy" :key="row.field" :class="row.visibility">
              <td>{{ row.field }}</td>
              <td><span class="badge">{{ row.visibility }}</span></td>
              <td><code>{{ row.onChain }}</code></td>
            </tr>
          </tbody>
        </table>
        <p class="muted">{{ fieldPolicy[4].note }}</p>
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
.muted { color: #64748b; font-size: 0.875rem; }
.status { color: #2563eb; font-size: 0.875rem; margin-top: 0.5rem; }
code { font-family: monospace; font-size: 0.75rem; color: #0f766e; }
.error { color: #dc2626; }
.result { margin-top: 1rem; background: #1e293b; color: #e2e8f0; padding: 1rem; border-radius: 8px; overflow: auto; font-size: 0.8rem; }
table { width: 100%; border-collapse: collapse; font-size: 0.85rem; }
th, td { border: 1px solid #e2e8f0; padding: 0.4rem 0.5rem; text-align: left; }
.policy-table tr.隐私 { background: #fef3c7; }
.policy-table tr.禁止 { background: #fee2e2; }
.badge { font-size: 0.75rem; font-weight: 600; }
input { display: block; width: 100%; margin: 0.5rem 0; padding: 0.5rem; border: 1px solid #cbd5e1; border-radius: 4px; }
button { margin-top: 0.5rem; padding: 0.5rem 1rem; background: #2563eb; color: white; border: none; border-radius: 4px; cursor: pointer; }
button:disabled { opacity: 0.6; cursor: not-allowed; }
</style>
