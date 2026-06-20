<script setup lang="ts">
import { computed, ref } from 'vue'
import { useLabI18n } from '@/composables/useLabI18n'
import { useLabSimulate } from '../../frontend/shared/useLabSimulate'
import { parseHints } from '../../frontend/shared/parseHints'

const { t } = useLabI18n('edu.cn.trace.charity')

const campaignId = ref('DEMO-CAMPAIGN-001')
const { loading, error, result, taskStatus, taskReport, runSimulate, parseEvaluation } =
  useLabSimulate('edu.cn.trace.charity')

const evaluation = computed(() => parseEvaluation(result.value?.evaluation))
const hints = computed(() => parseHints(evaluation.value?.audit_hints))

const demoEntries = computed(() => [
  { id: 'DEMO-ENTRY-001', donor: t('donor_a'), amount: 100, hash: 'sha256:donation-demo-001', time: '2026-06-01 12:00' },
  { id: 'DEMO-ENTRY-002', donor: t('donor_b'), amount: 50, hash: 'sha256:donation-demo-002', time: '2026-06-02 15:30' },
])

const fieldPolicy = computed(() => [
  { field: t('policy_amount'), visibility: t('public'), cssClass: '公开', onChain: '¥100', note: '汇总进活动账本' },
  { field: t('policy_time'), visibility: t('public'), cssClass: '公开', onChain: '2026-06-01T12:00:00Z', note: '审计轨迹' },
  { field: t('policy_donor'), visibility: t('public'), cssClass: '公开', onChain: t('donor_a'), note: '非真实 PII' },
  { field: t('policy_phone'), visibility: t('private'), cssClass: '隐私', onChain: 'sha256:redacted…demo', note: '仅 hash 上链' },
  { field: t('policy_channel'), visibility: t('forbidden'), cssClass: '禁止', onChain: '—', note: '不对接微信/支付宝' },
])

const ledgerHash = 'ledger-merkle-root-demo-9b2c'
const totalAmount = computed(() => demoEntries.value.reduce((s, e) => s + e.amount, 0))

function submit() {
  runSimulate(
    '模拟募捐流水哈希上链',
    { campaign_id: campaignId.value, entry_count: demoEntries.value.length, ledger_hash: ledgerHash },
    { taskType: 'CN_CHARITY_LEDGER_DEMO' },
  )
}
</script>

<template>
  <section class="lab-panel">
    <header class="lab-header">
      <img src="/assets/icon.png" alt="" width="32" height="32" />
      <div>
        <h1>{{ t('title') }}</h1>
        <p class="muted">{{ t('subtitle') }}</p>
      </div>
    </header>

    <div v-if="evaluation" class="eval-card">
      <h2>{{ t('ruleEval') }}</h2>
      <p class="ok">{{ t('compliancePass') }} · {{ evaluation.recommended_language }}</p>
      <p v-if="hints.campaign_id"><strong>{{ t('campaign') }}:</strong> {{ hints.campaign_id }}</p>
      <p v-if="hints.entry_count"><strong>{{ t('entryCount') }}:</strong> {{ hints.entry_count }}</p>
      <p v-if="hints.channel"><strong>{{ t('channel') }}:</strong> {{ hints.channel }} · {{ hints.org }}</p>
    </div>

    <div class="lab-grid">
      <div class="card">
        <h2>{{ t('campaignSection') }}</h2>
        <label>
          {{ t('campaignId') }}
          <input v-model="campaignId" />
        </label>
        <p><strong>{{ t('ledgerHash') }}:</strong> <code>{{ ledgerHash }}</code></p>
        <p><strong>{{ t('totalAmount') }}:</strong> ¥{{ totalAmount }}</p>
        <button :disabled="loading" @click="submit">
          {{ loading ? t('submitting') : t('submitLedger') }}
        </button>
        <p v-if="taskStatus" class="status">{{ t('taskStatus') }}: {{ taskStatus }}</p>
        <p v-if="error" class="error">{{ error }}</p>
      </div>

      <div class="card">
        <h2>{{ t('donations') }}</h2>
        <table>
          <thead>
            <tr><th>{{ t('donor') }}</th><th>{{ t('amount') }}</th><th>{{ t('hash') }}</th><th>{{ t('time') }}</th></tr>
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
        <h2>{{ t('fieldPolicy') }}</h2>
        <table class="policy-table">
          <thead>
            <tr><th>{{ t('field') }}</th><th>{{ t('visibility') }}</th><th>{{ t('onChainCol') }}</th></tr>
          </thead>
          <tbody>
            <tr v-for="row in fieldPolicy" :key="row.field" :class="row.cssClass">
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
