import { svelte } from '@sveltejs/vite-plugin-svelte';
import { defineConfig } from 'vite';
import vitePluginWindicss from 'vite-plugin-windicss';
import { join, dirname } from 'path';
import { fileURLToPath } from 'url';

export default defineConfig(({ mode }) => {
  const isProduction = mode === 'production';
  return {
    optimizeDeps: {
      exclude: ['@roxi/routify'],
    },
    envDir: join(dirname(fileURLToPath(import.meta.url)), '../..'),
    plugins: [
      // uses enforce: pre
      svelte(),
      vitePluginWindicss.default({
        transformCSS: 'pre',
      }),
    ],
    build: {
      minify: isProduction,
    },
  };
});
