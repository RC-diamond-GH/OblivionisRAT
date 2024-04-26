import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

export default defineConfig({
  plugins: [react()],
  // prevent vite from obscuring rust errors
  clearScreen: false,
  // Tauri expects a fixed port, fail if that port is not available
  server: {
    port: 3000,
    strictPort: true,
  },
  // to access the Tauri environment variables set by the CLI with information about the current target
  envPrefix: ['VITE_', 'TAURI_PLATFORM', 'TAURI_ARCH', 'TAURI_FAMILY', 'TAURI_PLATFORM_VERSION', 'TAURI_PLATFORM_TYPE', 'TAURI_DEBUG'],
  build: {
    // @ts-ignore --Tauri uses Chromium on Windows and WebKit on macOS and Linux
    target: process.env.TAURI_PLATFORM == 'windows' ? 'chrome105' : 'safari13',
    // @ts-ignore --don't minify for debug builds
    minify: !process.env.TAURI_DEBUG ? 'esbuild' : false,    
    // @ts-ignore --为调试构建生成源代码映射 (sourcemap)
    sourcemap: !!process.env.TAURI_DEBUG,
  },
})