<script>
  import { Tabs, TabItem } from "flowbite-svelte";
  import CodeEditor from "./_CodeEditor.svelte";

  export let body = {};
</script>

<Tabs
  style="underline"
  defaultClass="flex flex-wrap -mx-3 -my-1"
  contentClass="-p-4 py-1"
  inactiveClasses="flex flex-wrap -mx-3 -my-1"
>
  <TabItem
    open
    title="JSON"
    defaultClass="text-sm font-medium text-center disabled:cursor-not-allowed"
  >
    <div class="text-left h-full w-full">
      <CodeEditor bind:body />
    </div>
  </TabItem>
  <TabItem title="Auth">
    <p class="text-sm text-gray-500 dark:text-gray-400">
      <b>Dashboard:</b>
      Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor
      incididunt ut labore et dolore magna aliqua.
    </p>
  </TabItem>
  <TabItem title="Header">
    <p class="text-sm text-gray-500 dark:text-gray-400">
      <b>Dashboard:</b>
      Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor
      incididunt ut labore et dolore magna aliqua.
    </p>
  </TabItem>
</Tabs>
