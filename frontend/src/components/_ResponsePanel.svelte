<script>
  import { Badge, Button, Dropdown, DropdownItem } from "flowbite-svelte";

  import { ChevronDownSolid, ClockSolid } from "flowbite-svelte-icons";
  import ResTab from "./_ResTab.svelte";

  export let res;
</script>

<div class="flex flex-col bg-gray-50 dark:bg-gray-800">
  <div class="flex h-9 px-1 align-middle">
    {#if res?.status}
      <Badge color="green" border class="mx-2 text-[70%] font-bold text-nowrap"
        >{res.status}</Badge
      >
      <Badge color="dark" border class="mx-2 text-[70%] font-bold text-nowrap"
        ><ClockSolid class="w-2.5 h-2.5 me-1.5" />20s</Badge
      >
    {/if}
    <div class="flex-grow px-2 h-full" />

    <Button
      class="h-full border-none shadow-none w-24 text-sm font-small text-gray-800 bg-gray-200 dark:bg-gray-700 dark:text-gray-300 hover:bg-gray-300 dark:hover:bg-gray-600"
      >GET <ChevronDownSolid
        class="ml-6 w-3 h-5  text-gray-800  dark:text-gray-300"
        size="sm"
      /></Button
    >
    <Dropdown>
      <DropdownItem>Dashboard</DropdownItem>
      <DropdownItem>Settings</DropdownItem>
      <DropdownItem>Earnings</DropdownItem>
      <DropdownItem>Sign out</DropdownItem>
    </Dropdown>
  </div>
  <div class="flex flex-col h-[93vh] w-full">
    <!-- cannot figure out the issue with editor height/width/scroll - needs to be fixed -->
    <ResTab {res} />
  </div>
</div>
