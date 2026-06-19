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
  id: 'edu.cn.trace.food',
  title: 'Food Trace Simulation Lab',
  routePrefix: '/labs/edu.cn.trace.food',
  routes: [
    { path: '', component: () => import('./FoodTraceLab.vue') },
  ],
  apiBase: '/api/v1/labs/edu.cn.trace.food',
}

export default plugin
