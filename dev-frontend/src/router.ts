import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  { path: '/', redirect: '/labs/edu.cn.trace.food' },
  {
    path: '/labs/edu.cn.trace.food',
    component: () => import('@plugins/food-trace/frontend/FoodTraceLab.vue'),
  },
  {
    path: '/labs/edu.cn.trace.medical',
    component: () => import('@plugins/medical-tamper/frontend/MedicalTamperLab.vue'),
  },
  {
    path: '/labs/edu.cn.trace.charity',
    component: () => import('@plugins/charity-ledger/frontend/CharityLedgerLab.vue'),
  },
]

export default createRouter({
  history: createWebHistory(),
  routes,
})
