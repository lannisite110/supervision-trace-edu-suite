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
  id: 'edu.cn.trace.charity',
  title: 'Charity Ledger Demo',
  routePrefix: '/labs/edu.cn.trace.charity',
  routes: [
    { path: '', component: () => import('./CharityLedgerLab.vue') },
  ],
  apiBase: '/api/v1/labs/edu.cn.trace.charity',
}

export default plugin
