<script>
  import { Tabs, TabItem } from "flowbite-svelte";
  import CodeEditor from "./_CodeEditor.svelte";
  import TabularData from "./_TabularData.svelte";

  export let res = {};
</script>

<Tabs
  style="underline"
  defaultClass="flex flex-wrap -mx-1 -mt-1"
  contentClass="-p-4 py-1"
  inactiveClasses="flex flex-wrap -mx-1 -mt-1"
>
  <TabItem
    open
    title="Response"
    defaultClass="text-sm font-medium text-center disabled:cursor-not-allowed"
  >
    <div class="text-left">
      <CodeEditor body={res?.body} />
    </div>
  </TabItem>
  <TabItem title="Header">
    <div class="text-left">
      {#if res?.headers?.length > 0}
        <TabularData data={res?.headers} />
      {:else}
        No header data to display.
      {/if}
    </div>
  </TabItem>
  <!-- <TabItem title="Log">
    <div class="text-left">
      {#if res?.log?.length > 0}
        <TabularData data={res?.log} />
      {:else}
        No log data to display.
      {/if}
    </div>
  </TabItem> -->
</Tabs>
