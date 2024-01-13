<script>
  import { Sidebar } from "flowbite-svelte";

  import NavContent from "./components/_NavContent.svelte";
  import SideContent from "./components/_SideContent.svelte";
  import RequestPanel from "./components/_RequestPanel.svelte";
  import ResponsePanel from "./components/_ResponsePanel.svelte";
  import { onMount } from "svelte";

  let reqPane, resPane, divider;

  let req = {
    method: "GET",
    url: "",
    body: "",
    headers: [{ id: 0, key: "", value: "" }],
    contentType: "",
  };
  let res = {};

  onMount(() => {
    let mousedown = false;

    divider.addEventListener("mousedown", () => {
      mousedown = true;
    });

    window.addEventListener("mouseup", () => {
      mousedown = false;
    });

    window.addEventListener("mousemove", (e) => {
      if (!mousedown) return;
      e.preventDefault();

      let x = e.clientX;
      console.log("x: ", x);
      let leftWidth = x - reqPane.getBoundingClientRect().left;
      let rightWidth = window.innerWidth - x;

      reqPane.style.width = `${leftWidth}px`;
      resPane.style.width = `${rightWidth}px`;
    });
  });
</script>

<main class="bg-gray-100 dark:bg-gray-800">
  <NavContent />
  <div class="flex w-screen">
    <Sidebar class="w-48 min-w-48">
      <SideContent />
    </Sidebar>

    <div class="content flex flex-grow h-full">
      <div class="w-1/2 overflow-auto" bind:this={reqPane}>
        <RequestPanel bind:req bind:res />
      </div>
      <div
        id="divide"
        class="w-1 cursor-col-resize bg-gray-200 dark:bg-gray-700 shadow-sm inset-2"
        bind:this={divider}
      ></div>

      <div class="w-1/2 overflow-auto" bind:this={resPane}>
        <ResponsePanel {res} />
      </div>
    </div>
  </div>
</main>

<style>
</style>
