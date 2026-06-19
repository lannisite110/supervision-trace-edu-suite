import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import { fileURLToPath, URL } from 'node:url'

export default defineConfig({
  plugins: [vue()],
  resolve: {
    alias: {
      '@plugins': fileURLToPath(new URL('../plugins', import.meta.url)),
    },
  },
  server: {
    port: 5174,
    fs: { allow: ['..'] },
    proxy: {
      '/api': { target: 'http://127.0.0.1:8080', changeOrigin: true },
    },
  },
})
