<script setup lang="ts">
import { computed, ref } from 'vue'
import { useLabI18n } from '@/composables/useLabI18n'
import { useLabSimulate } from '../../frontend/shared/useLabSimulate'
import { parseHints } from '../../frontend/shared/parseHints'
import LabAssistDrawer from '@/components/LabAssistDrawer.vue'

const { t } = useLabI18n('edu.cn.trace.food')

const batchId = ref('DEMO-BATCH-001')
const { loading, error, result, taskStatus, taskReport, runSimulate, parseEvaluation } =
  useLabSimulate('edu.cn.trace.food')

const evaluation = computed(() => parseEvaluation(result.value?.evaluation))
const hints = computed(() => parseHints(evaluation.value?.audit_hints))

const assistParams = computed(() => ({
  batch_id: batchId.value,
  product_name: '有机蔬菜礼盒',
}))

const demoTrace = computed(() => [
  { stage: t('stage_harvest'), origin: t('origin_a'), hash: 'a3f2c1…demo', time: '2026-06-01 08:00' },
  { stage: t('stage_cold'), origin: t('origin_cold'), hash: 'b4e3d2…demo', time: '2026-06-02 14:30' },
  { stage: t('stage_retail'), origin: t('origin_store'), hash: 'c5f4e3…demo', time: '2026-06-03 09:15' },
])

const fabricSandbox = {
  chainId: 'fabric-local',
  channel: 'edu-cn-trace-sandbox',
  org: 'OrgEduDemo',
  chaincode: 'plugins/food-trace/chaincode/food_trace.go',
}

const merkleProof = computed(() => ({
  rootHash: 'merkle-root-demo-7f3a',
  leafHashes: demoTrace.value.map((s) => s.hash),
  proofPath: [t('proofPath')],
}))

function submit() {
  runSimulate(
    '食品批次溯源教学演示',
    { batch_id: batchId.value, product_name: '有机蔬菜礼盒' },
    { taskType: 'CN_FOOD_TRACE_SIM' },
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
      <p class="ok">{{ t('evalDetail', { lang: evaluation.recommended_language, template: evaluation.recommended_template }) }}</p>
      <p v-if="hints.channel"><strong>{{ t('channel') }}:</strong> {{ hints.channel }} · <strong>{{ t('org') }}:</strong> {{ hints.org }}</p>
      <p v-if="hints.batch_id"><strong>{{ t('batch') }}:</strong> {{ hints.batch_id }}</p>
    </div>

    <div class="lab-grid">
      <div class="card">
        <h2>{{ t('batchInfo') }}</h2>
        <label>
          {{ t('batchId') }}
          <input v-model="batchId" />
        </label>
        <button :disabled="loading" @click="submit">
          {{ loading ? t('submitting') : t('submit') }}
        </button>
        <p v-if="taskStatus" class="status">{{ t('taskStatus') }}: {{ taskStatus }}</p>
        <p v-if="error" class="error">{{ error }}</p>
      </div>

      <div class="card">
        <h2>{{ t('traceChain') }}</h2>
        <ol class="trace-timeline">
          <li v-for="step in demoTrace" :key="step.stage">
            <strong>{{ step.stage }}</strong> — {{ step.origin }}
            <br /><span class="hash">{{ step.hash }}</span>
            <br /><small>{{ step.time }}</small>
          </li>
        </ol>
      </div>

      <div class="card">
        <h2>{{ t('merkleProof') }}</h2>
        <p><strong>Root:</strong> <code>{{ merkleProof.rootHash }}</code></p>
        <ul>
          <li v-for="h in merkleProof.leafHashes" :key="h"><code>{{ h }}</code></li>
        </ul>
        <p class="muted">{{ merkleProof.proofPath.join(' → ') }}</p>
      </div>
      <div class="card fabric-card">
        <h2>{{ t('fabricSandbox') }}</h2>
        <p><strong>chainId:</strong> <code>{{ fabricSandbox.chainId }}</code></p>
        <p><strong>{{ t('channel') }}:</strong> <code>{{ fabricSandbox.channel }}</code></p>
        <p><strong>{{ t('org') }}:</strong> {{ fabricSandbox.org }}</p>
        <p class="muted"><code>{{ fabricSandbox.chaincode }}</code></p>
      </div>
    </div>

    <pre v-if="taskReport" class="result">{{ JSON.stringify(taskReport, null, 2) }}</pre>
    <pre v-else-if="result" class="result">{{ JSON.stringify(result, null, 2) }}</pre>

    <LabAssistDrawer
      plugin-id="edu.cn.trace.food"
      :params="assistParams"
      :audit-hints="evaluation?.audit_hints"
    />
  </section>
</template>

<style scoped>
.lab-panel { padding: 1rem; }
.eval-card { background: #f0fdf4; border: 1px solid #bbf7d0; border-radius: 8px; padding: 1rem; margin-bottom: 1rem; }
.eval-card h2 { margin: 0 0 0.5rem; font-size: 1rem; }
.eval-card .ok { color: #15803d; margin: 0 0 0.5rem; }
.lab-header { display: flex; gap: 0.75rem; align-items: center; margin-bottom: 1rem; }
.lab-grid { display: grid; grid-template-columns: repeat(auto-fit, minmax(260px, 1fr)); gap: 1rem; }
.card { background: #f8fafc; border: 1px solid #e2e8f0; border-radius: 8px; padding: 1rem; }
.card.fabric-card { background: #eff6ff; border-color: #bfdbfe; }
.muted { color: #64748b; font-size: 0.875rem; }
.status { color: #2563eb; font-size: 0.875rem; margin-top: 0.5rem; }
.hash, code { font-family: monospace; font-size: 0.8rem; color: #0f766e; }
.trace-timeline { padding-left: 1.25rem; }
.trace-timeline li { margin-bottom: 0.75rem; }
.error { color: #dc2626; }
.result { margin-top: 1rem; background: #1e293b; color: #e2e8f0; padding: 1rem; border-radius: 8px; overflow: auto; font-size: 0.8rem; }
input { display: block; width: 100%; margin: 0.5rem 0; padding: 0.5rem; border: 1px solid #cbd5e1; border-radius: 4px; }
button { margin-top: 0.5rem; padding: 0.5rem 1rem; background: #2563eb; color: white; border: none; border-radius: 4px; cursor: pointer; }
button:disabled { opacity: 0.6; cursor: not-allowed; }
</style>
