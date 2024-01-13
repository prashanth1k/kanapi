<script>
  import {
    SidebarGroup,
    SidebarWrapper,
    Input,
    Tooltip,
    SidebarDropdownWrapper,
    Dropdown,
    DropdownItem,
  } from "flowbite-svelte";
  import {
    FilterOutline,
    FolderOutline,
    DotsHorizontalOutline,
  } from "flowbite-svelte-icons";

  let folders = [
    {
      name: "app1",
      isExpanded: true,
      apis: [{ name: "api11" }, { name: "api12" }],
    },
    {
      name: "app2",
      isExpanded: false,
      apis: [{ name: "api21" }, { name: "api22" }],
    },
  ];
  let showMenu = false;
</script>

<SidebarWrapper class="pl-3 pr-1 h-full bg-gray-50 dark:bg-gray-800">
  <SidebarGroup class="-mt-4">
    <Input type="text" placeholder="Filter" class="mb-2 h-9">
      <div slot="right">
        <FilterOutline
          class="w-5 h-5 text-gray-500 transition duration-75 dark:text-gray-400 group-hover:text-gray-900 dark:group-hover:text-white"
        />
      </div></Input
    >
    <Tooltip>Filter by folder / API name.</Tooltip>
  </SidebarGroup>
  <SidebarGroup>
    {#each folders as folder, i}
      <SidebarDropdownWrapper label={folder.name} isOpen={folder.isExpanded}>
        <svelte:fragment slot="icon">
          <FolderOutline
            class="w-5 h-5 text-gray-500 transition duration-75 dark:text-gray-400 group-hover:text-gray-900 dark:group-hover:text-white"
          />
        </svelte:fragment>

        {#each folder.apis as folApi, j}
          <div
            class="flex items-center justify-between h-9 p-3 text-gray-900 dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700 cursor-pointer relative"
          >
            <a
              href="#"
              class="flex items-center p-2 rounded-lg text-gray-900 dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700 group"
              >{folApi.name}</a
            >

            <div>
              <button
                class="inline-flex items-center px-2 ms-3 text-sm font-small text-gray-800 bg-gray-200 rounded-full dark:bg-gray-700 dark:text-gray-300"
                on:click|stopPropagation={() => (showMenu = !showMenu)}
              >
                <DotsHorizontalOutline />
              </button>
              <Dropdown>
                <DropdownItem>Dashboard</DropdownItem>
                <DropdownItem>Settings</DropdownItem>
                <DropdownItem>Earnings</DropdownItem>
                <DropdownItem>Sign out</DropdownItem>
              </Dropdown>
            </div>
          </div>
        {/each}
      </SidebarDropdownWrapper>
    {/each}
  </SidebarGroup>
</SidebarWrapper>
