<script setup lang="ts">
import { useRoute, useRouter } from 'vue-router'

const route = useRoute()
const router = useRouter()

const labs = [
  { id: 'edu.cn.trace.food', name: '食品溯源', path: '/labs/edu.cn.trace.food' },
  { id: 'edu.cn.trace.medical', name: '医疗防篡改', path: '/labs/edu.cn.trace.medical' },
  { id: 'edu.cn.trace.charity', name: '慈善流水', path: '/labs/edu.cn.trace.charity' },
]
</script>

<template>
  <div class="layout">
    <aside class="sidebar">
      <div class="brand">
        <img src="/assets/icon.png" alt="" width="32" height="32" />
        <div>
          <div class="title">民生溯源 Demo</div>
          <div class="subtitle">Fabric 沙箱 · 开发预览</div>
        </div>
      </div>
      <nav>
        <button
          v-for="lab in labs"
          :key="lab.id"
          class="nav-item"
          :class="{ active: route.path === lab.path }"
          @click="router.push(lab.path)"
        >
          {{ lab.name }}
        </button>
      </nav>
      <p class="hint">需主库网关 :8080 已启动</p>
    </aside>
    <main class="content">
      <router-view />
    </main>
  </div>
</template>

<style scoped>
.layout { display: flex; min-height: 100vh; }
.sidebar { width: 220px; background: #0f172a; color: #e2e8f0; padding: 1rem; }
.brand { display: flex; gap: 0.5rem; align-items: center; margin-bottom: 1.5rem; }
.title { font-weight: 700; }
.subtitle { font-size: 0.75rem; color: #94a3b8; }
.nav-item {
  display: block; width: 100%; text-align: left; padding: 0.5rem 0.75rem;
  margin-bottom: 0.25rem; border: none; border-radius: 6px; background: transparent;
  color: #cbd5e1; cursor: pointer;
}
.nav-item:hover { background: #1e293b; }
.nav-item.active { background: #2563eb; color: white; }
.hint { margin-top: 2rem; font-size: 0.7rem; color: #64748b; }
.content { flex: 1; background: #f1f5f9; }
</style>
