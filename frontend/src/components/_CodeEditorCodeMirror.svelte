<script>
  import CodeMirror from "svelte-codemirror-editor";
  import { javascript } from "@codemirror/lang-javascript";

  import { oneDark } from "@codemirror/theme-one-dark";

  let value = "console.log('Hello World')";
</script>

<div class="w-full h-full">
  <CodeMirror bind:value lang={javascript()} theme="oneDarkTheme" />
</div>
