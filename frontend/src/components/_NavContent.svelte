<script>
  import {
    Navbar,
    NavBrand,
    NavUl,
    NavHamburger,
    Button,
  } from "flowbite-svelte";
  import { SunOutline, MoonOutline } from "flowbite-svelte-icons";
  import { pref } from "../stores/store.js";

  function toggleDarkMode() {
    if (!$pref?.darkMode) document.documentElement.classList.add("dark");
    else document.documentElement.classList.remove("dark");
    pref.set({ darkMode: !$pref.darkMode });
  }
</script>

<Navbar fluid class="h-10">
  <NavBrand href="/">
    <img
      src="src/assets/images/kanapi-logo.png"
      class="me h-4 sm:h-6 dark:filter dark:invert-[100] dark:-hue-rotate-180"
      alt="kanapi logo"
    />
  </NavBrand>
  <NavHamburger />
  <NavUl class="-m-6" fluid="true">
    <Button
      class="bg-transparent text-primary-700 hover:bg-grey-500 dark:hover:bg-grey-100 dark:bg-transparent dark:text-primary-50 size-fit"
      on:click={toggleDarkMode}
    >
      {#if $pref?.darkMode}
        <SunOutline />
      {:else}
        <MoonOutline />
      {/if}
    </Button>
    <!-- <Button
      class="bg-transparent text-primary-700 hover:bg-grey-500 dark:hover:bg-grey-100 dark:bg-transparent dark:text-primary-50"
      on:click={() => console.log("clicked")}
    >
      <QuestionCircleOutline />
    </Button>
    <Button
      class="bg-transparent text-primary-700 hover:bg-grey-500 dark:hover:bg-grey-100 dark:bg-transparent dark:text-primary-50"
      on:click={() => console.log("clicked")}
    >
      <CogOutline />
    </Button> -->
  </NavUl>
</Navbar>
