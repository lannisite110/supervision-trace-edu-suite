<script setup lang="ts">
import { ref } from 'vue'
import { useLabSimulate } from '../../frontend/shared/useLabSimulate'

const campaignId = ref('DEMO-CAMPAIGN-001')
const { loading, error, result, taskStatus, taskReport, runSimulate } = useLabSimulate('edu.cn.trace.charity')

const demoEntries = [
  { id: 'DEMO-ENTRY-001', donor: '虚构捐赠者A', amount: 100, hash: 'sha256:donation-demo-001', time: '2026-06-01 12:00' },
  { id: 'DEMO-ENTRY-002', donor: '虚构捐赠者B', amount: 50, hash: 'sha256:donation-demo-002', time: '2026-06-02 15:30' },
]

const ledgerHash = 'ledger-merkle-root-demo-9b2c'
const totalAmount = demoEntries.reduce((s, e) => s + e.amount, 0)

function submit() {
  runSimulate('模拟募捐流水哈希上链', { campaign_id: campaignId.value })
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
    </div>

    <pre v-if="taskReport" class="result">{{ JSON.stringify(taskReport, null, 2) }}</pre>
    <pre v-else-if="result" class="result">{{ JSON.stringify(result, null, 2) }}</pre>
  </section>
</template>

<style scoped>
.lab-panel { padding: 1rem; }
.lab-header { display: flex; gap: 0.75rem; align-items: center; margin-bottom: 1rem; }
.lab-grid { display: grid; grid-template-columns: repeat(auto-fit, minmax(300px, 1fr)); gap: 1rem; }
.card { background: #f8fafc; border: 1px solid #e2e8f0; border-radius: 8px; padding: 1rem; }
.muted { color: #64748b; font-size: 0.875rem; }
.status { color: #2563eb; font-size: 0.875rem; margin-top: 0.5rem; }
code { font-family: monospace; font-size: 0.75rem; color: #0f766e; }
.error { color: #dc2626; }
.result { margin-top: 1rem; background: #1e293b; color: #e2e8f0; padding: 1rem; border-radius: 8px; overflow: auto; font-size: 0.8rem; }
table { width: 100%; border-collapse: collapse; font-size: 0.85rem; }
th, td { border: 1px solid #e2e8f0; padding: 0.4rem 0.5rem; text-align: left; }
input { display: block; width: 100%; margin: 0.5rem 0; padding: 0.5rem; border: 1px solid #cbd5e1; border-radius: 4px; }
button { margin-top: 0.5rem; padding: 0.5rem 1rem; background: #2563eb; color: white; border: none; border-radius: 4px; cursor: pointer; }
button:disabled { opacity: 0.6; cursor: not-allowed; }
</style>
