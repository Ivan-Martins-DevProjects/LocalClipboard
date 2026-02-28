import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import path from 'path' // Certifique-se de importar o path

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  resolve: {
    alias: {
      // Isso mapeia o @wailsjs para a pasta física
      '@wailsjs': path.resolve(__dirname, './wailsjs'),
    },
  },
})
