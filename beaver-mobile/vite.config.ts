import { defineConfig } from "vite";
import { createRequire } from "node:module";
import { resolve } from "node:path"

const nodeRequire = createRequire(import.meta.url);
const uni = nodeRequire("@dcloudio/vite-plugin-uni").default;
const envConfig = nodeRequire('./env.json');
const baseUrl = envConfig.baseUrl;
const wsUrl = envConfig.wsUrl;
const friendUrl = envConfig.friendUrl;

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
  },
  server: {
	proxy: {
		'/api/auth': {
			target: baseUrl,
			changeOrigin: true,
			rewrite: (path) => path
		},
		'/api/chat': {
			target: wsUrl,
			changeOrigin: true,
			rewrite: (path) => path
		},
		'/api/friend': {
			target: friendUrl,
			changeOrigin: true,
			rewrite: (path) => path
		}
	}
  }
});
