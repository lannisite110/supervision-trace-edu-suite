export interface LabPlugin {
  id: string
  title: string
  routePrefix: string
  routes: Array<{
    path: string
    component: () => Promise<unknown>
  }>
  apiBase: string
}

export const plugin: LabPlugin = {
  id: 'edu.cn.trace.medical',
  title: 'Medical Tamper Detection Lab',
  routePrefix: '/labs/edu.cn.trace.medical',
  routes: [
    { path: '', component: () => import('./MedicalTamperLab.vue') },
  ],
  apiBase: '/api/v1/labs/edu.cn.trace.medical',
}

export default plugin
