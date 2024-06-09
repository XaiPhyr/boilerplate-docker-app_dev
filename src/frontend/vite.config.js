import { defineConfig, loadEnv } from 'vite';
import vue from '@vitejs/plugin-vue';
import { fileURLToPath, URL } from 'url';

// https://vitejs.dev/config/
export default defineConfig(({ mode }) => {
  const env = loadEnv(mode, process.cwd());

  return {
    plugins: [vue()],
    define: { 'process.env': env },
    resolve: {
      alias: [
        {
          find: 'api',
          replacement: fileURLToPath(new URL('./src/api', import.meta.url)),
        },
        {
          find: 'components',
          replacement: fileURLToPath(
            new URL('./src/components', import.meta.url)
          ),
        },
      ],
    },
  };
});
