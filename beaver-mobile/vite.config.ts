import { defineConfig } from "vite";
import { createRequire } from "node:module";
import { resolve } from "node:path"

const nodeRequire = createRequire(import.meta.url);
const uni = nodeRequire("@dcloudio/vite-plugin-uni").default;


// https://vitejs.dev/config/
export default defineConfig({
  plugins: [uni()],
  resolve: {
    alias: {
      "@": resolve(__dirname, "/")
    }
  },
  css: {
    preprocessorOptions: {
      scss: {
        api: 'modern-compiler', // or 'modern'
      },
    },
  }
});
