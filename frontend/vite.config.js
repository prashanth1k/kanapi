import { defineConfig } from "vite";
import { svelte } from "@sveltejs/vite-plugin-svelte";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [svelte()],
  optimizeDeps: {
    exclude: [
      "svelte-codemirror-editor",
      "codemirror",
      "@codemirror/lang-javascript",
      "@codemirror/theme-one-dark" /* ... */,
    ],
  },
});
