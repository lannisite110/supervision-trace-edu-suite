<script setup lang="ts">
import { ref } from 'vue'
import { useLabSimulate } from '../../frontend/shared/useLabSimulate'

const batchId = ref('DEMO-BATCH-001')
const { loading, error, result, taskStatus, taskReport, runSimulate } = useLabSimulate('edu.cn.trace.food')

const demoTrace = [
  { stage: '采收', origin: '虚构农场A', hash: 'a3f2c1…demo', time: '2026-06-01 08:00' },
  { stage: '冷链运输', origin: '虚构冷链中心', hash: 'b4e3d2…demo', time: '2026-06-02 14:30' },
  { stage: '上架销售', origin: '虚构零售门店', hash: 'c5f4e3…demo', time: '2026-06-03 09:15' },
]

const merkleProof = {
  rootHash: 'merkle-root-demo-7f3a',
  leafHashes: demoTrace.map((s) => s.hash),
  proofPath: ['hash→parent→root (教学简化)'],
}

function submit() {
  runSimulate('食品批次溯源教学演示', {
    batch_id: batchId.value,
    product_name: '有机蔬菜礼盒',
  })
}
</script>

<template>
  <section class="lab-panel">
    <header class="lab-header">
      <img src="/assets/icon.png" alt="" width="32" height="32" />
      <div>
        <h1>食品溯源教学 Demo</h1>
        <p class="muted">虚构数据 · Fabric 沙箱 · 批次/Merkle 证明演示</p>
      </div>
    </header>

    <div class="lab-grid">
      <div class="card">
        <h2>批次信息</h2>
        <label>
          批次 ID
          <input v-model="batchId" />
        </label>
        <button :disabled="loading" @click="submit">
          {{ loading ? '提交中…' : '提交仿真实验' }}
        </button>
        <p v-if="taskStatus" class="status">任务状态: {{ taskStatus }}</p>
        <p v-if="error" class="error">{{ error }}</p>
      </div>

      <div class="card">
        <h2>溯源链路</h2>
        <ol class="trace-timeline">
          <li v-for="step in demoTrace" :key="step.stage">
            <strong>{{ step.stage }}</strong> — {{ step.origin }}
            <br /><span class="hash">{{ step.hash }}</span>
            <br /><small>{{ step.time }}</small>
          </li>
        </ol>
      </div>

      <div class="card">
        <h2>Merkle 证明</h2>
        <p><strong>Root:</strong> <code>{{ merkleProof.rootHash }}</code></p>
        <ul>
          <li v-for="h in merkleProof.leafHashes" :key="h"><code>{{ h }}</code></li>
        </ul>
        <p class="muted">{{ merkleProof.proofPath.join(' → ') }}</p>
      </div>
    </div>

    <pre v-if="taskReport" class="result">{{ JSON.stringify(taskReport, null, 2) }}</pre>
    <pre v-else-if="result" class="result">{{ JSON.stringify(result, null, 2) }}</pre>
  </section>
</template>

<style scoped>
.lab-panel { padding: 1rem; }
.lab-header { display: flex; gap: 0.75rem; align-items: center; margin-bottom: 1rem; }
.lab-grid { display: grid; grid-template-columns: repeat(auto-fit, minmax(260px, 1fr)); gap: 1rem; }
.card { background: #f8fafc; border: 1px solid #e2e8f0; border-radius: 8px; padding: 1rem; }
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
