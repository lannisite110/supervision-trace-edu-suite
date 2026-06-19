import { ref } from 'vue'

export interface SimulateResult {
  plugin_id?: string
  evaluation?: {
    compliance_passed?: boolean
    recommended_template?: string
    audit_hints?: string[]
  }
  task?: {
    id?: string
    namespace?: string
    status?: string
    task_type?: string
  }
}

export function useLabSimulate(pluginId: string, defaultChainIds: (number | string)[] = ['fabric-local']) {
  const loading = ref(false)
  const error = ref('')
  const result = ref<SimulateResult | null>(null)
  const taskStatus = ref('')
  const taskReport = ref<Record<string, unknown> | null>(null)

  async function pollTask(taskId: string, maxAttempts = 10) {
    taskStatus.value = 'pending'
    for (let i = 0; i < maxAttempts; i++) {
      const statusRes = await fetch(`/api/v1/labs/${pluginId}/status/${taskId}`)
      const statusData = await statusRes.json()
      taskStatus.value = statusData.status ?? 'unknown'
      if (taskStatus.value === 'completed' || taskStatus.value === 'failed') {
        const reportRes = await fetch(`/api/v1/labs/${pluginId}/report/${taskId}`)
        taskReport.value = await reportRes.json()
        return
      }
      await new Promise((r) => setTimeout(r, 200))
    }
  }

  async function runSimulate(userPrompt: string, params: Record<string, unknown>) {
    loading.value = true
    error.value = ''
    result.value = null
    taskStatus.value = ''
    taskReport.value = null
    try {
      const res = await fetch(`/api/v1/labs/${pluginId}/simulate`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
          user_prompt: userPrompt,
          params,
          allowed_chain_ids: defaultChainIds,
        }),
      })
      const data = await res.json()
      if (!res.ok) throw new Error(data.error || res.statusText)
      result.value = data
      const evalBlock = typeof data.evaluation === 'string' ? JSON.parse(data.evaluation) : data.evaluation
      if (evalBlock && !evalBlock.compliance_passed) {
        throw new Error(evalBlock.rejection_reason || 'compliance check failed')
      }
      const taskId = data.task?.id
      if (taskId) await pollTask(taskId)
    } catch (e) {
      error.value = e instanceof Error ? e.message : String(e)
    } finally {
      loading.value = false
    }
  }

  return { loading, error, result, taskStatus, taskReport, runSimulate }
}
