<script>
  import "../assets/css/prism-vsc-dark-plus.css"; // Import the CSS here
  import Prism from "prismjs";
  import { CodeJar } from "@novacbn/svelte-codejar";

  export let syntax = "javascript"; // Export the syntax variable

  export let body;
  let value = "";
  $: {
    value = body ? body : "";
  }

  let highlight = (code, syntax) =>
    Prism.highlight(code, Prism.languages[syntax], syntax);
</script>

<div class="w-full h-full">
  <CodeJar {syntax} {highlight} withLineNumbers {value} />
</div>

<style>
  ::-webkit-scrollbar {
    width: 3px;
  }

  ::-webkit-scrollbar-track {
    background: #54525200;
  }

  ::-webkit-scrollbar-thumb {
    background: #888;
  }

  ::-webkit-scrollbar-thumb:hover {
    background: #555;
  }
</style>
