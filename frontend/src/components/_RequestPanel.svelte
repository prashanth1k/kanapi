<script>
  import { Button, Dropdown, DropdownItem, Input } from "flowbite-svelte";
  import { CallApi } from "../../wailsjs/go/main/App.js";
  import { ChevronDownSolid } from "flowbite-svelte-icons";

  import ReqTab from "./_ReqTab.svelte";

  export let req = {
    method: "GET",
    url: "",
    body: "",
    headers: [],
    contentType: "application/none",
  };
  export let res;

  const dropdownOptions = ["GET", "POST", "PUT", "PATCH", "DELETE"];
  let dropdownOpen = false;

  function goCallApi() {
    console.log("calling", req.url);
    CallApi(req).then((result) => {
      res = result;
      console.log("result", res);
    });
  }
</script>

<form on:submit|preventDefault={goCallApi}>
  <div class="flex flex-col bg-gray-50 dark:bg-gray-800">
    <div class="flex h-9">
      <Button
        class="h-full border-none shadow-none w-20 mx-1 text-sm font-small text-gray-800 bg-gray-200  dark:bg-gray-700 hover:bg-gray-300 dark:hover:bg-gray-600 dark:text-gray-300"
        >{req.method}
        <ChevronDownSolid
          class="ml-2 w-3 h-5 text-gray-800 dark:text-gray-300"
          size="sm"
        /></Button
      >
      <Dropdown bind:open={dropdownOpen}>
        {#each dropdownOptions as option}
          <DropdownItem
            on:click={() => {
              req.method = option;
              dropdownOpen = false;
            }}>{option}</DropdownItem
          >
        {/each}
      </Dropdown>

      <Input
        id="req-url"
        placeholder="https://url.com/v1/posts"
        class="flex-grow px-2 h-full"
        bind:value={req.url}
      />

      <Button
        class="h-full border-none shadow-none w-18 mx-1 text-sm font-small text-gray-800 bg-gray-200 dark:bg-gray-700 dark:text-gray-300 hover:bg-gray-300 dark:hover:bg-gray-600"
        type="submit">Go</Button
      >
    </div>
    <!-- <div class="flex flex-col bg-gray-50 dark:bg-gray-800 h-fit w-fit"> -->
    <div class="flex flex-col h-[93vh] w-full">
      <ReqTab bind:body={req.body} />
    </div>
  </div>
</form>
