<script>
  import loader from "@monaco-editor/loader";
  import { onDestroy, onMount } from "svelte";
  // import * as monaco from "monaco-editor";

  let editor;
  let monaco;

  let editorContainer;

  onMount(async () => {
    // Remove the next two lines to load the monaco editor from a CDN
    // see https://www.npmjs.com/package/@monaco-editor/loader#config
    console.log("started loading");

    // loader.config({ monaco });
    loader.config({ paths: { vs: "/node_modules/monaco-editor/min/vs" } });

    monaco = await loader.init();
    console.log("init done");

    // Your monaco instance is ready, let's display some code!
    editor = monaco.editor.create(editorContainer);
    const model = monaco.editor.createModel(
      "console.log('Hello from Monaco! (the editor, not the city...)')",
      "javascript"
    );
    editor.setModel(model);
    console.log("all done");
  });

  onDestroy(() => {
    monaco?.editor.getModels().forEach((model) => model.dispose());
    editor?.dispose();
  });
</script>

<div class="container" bind:this={editorContainer} />

<style>
  .container {
    width: 100%;
    height: 600px;
    text-align: left;
  }
</style>
