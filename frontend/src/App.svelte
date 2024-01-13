<script>
  import { onMount, onDestroy } from "svelte";
  import { pref } from "./stores/store.js";

  import { Pane, Splitpanes } from "svelte-splitpanes";

  import { GetPref } from "../wailsjs/go/main/App.js";

  import NavContent from "./components/_NavContent.svelte";
  // import SideContent from "./components/_SideContent.svelte";
  import RequestPanel from "./components/_RequestPanel.svelte";
  import ResponsePanel from "./components/_ResponsePanel.svelte";

  let req = {
    method: "GET",
    url: "",
    body: "",
    headers: [{ id: 0, key: "", value: "" }],
    contentType: "",
  };
  let res = {};

  //  set preferences on app start
  onMount(async () => {
    const prefLoc = await GetPref();
    pref.set(prefLoc);
    setLocalPref();
  });

  const setLocalPref = () => {
    if ($pref?.darkMode) {
      document.documentElement.classList.add("dark");
    } else {
      document.documentElement.classList.remove("dark");
    }
  };
</script>

<main class="bg-gray-50 dark:bg-gray-800">
  <NavContent />

  <Splitpanes>
    <!-- <Pane size={18} maxSize={24}><SideContent /></Pane> -->
    <Pane>
      <RequestPanel bind:req bind:res />
    </Pane>
    <Pane>
      <ResponsePanel {res} />
    </Pane>
  </Splitpanes>
</main>

<style>
  :global(.splitpanes__pane) {
    position: relative;
    overflow-y: auto !important;
    overflow-x: auto !important;
  }
  :global(
      .default-theme.splitpanes--vertical > .splitpanes__splitter,
      .default-theme .splitpanes--vertical > .splitpanes__splitter
    ) {
    width: 5px !important;
    background-color: darkgrey !important;
  }
</style>
