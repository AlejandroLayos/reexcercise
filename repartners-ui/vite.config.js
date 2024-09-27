import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import vuetify from 'vite-plugin-vuetify'
const path = require('path')

export default defineConfig({
  plugins: [
    vue(),
    vuetify({
      autoImport: true,
    }),
  ],
  define: { 'process.env': {} },
  resolve: {
    alias: {
      '@': path.resolve(__dirname, 'src'),
    },
  },
  server: {
    host: '0.0.0.0',
    port: 4173,
    proxy: {
      '/packs': {
        target: 'http://13.60.190.214:8080',
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/packs/, '/packs')
      }
    }
  }
})