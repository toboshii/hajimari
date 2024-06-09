<script lang="ts">
  import Icon from "@iconify/svelte";

  import Modal from "$lib/Modal/Modal.svelte";
  import Search from "$lib/Search/Search.svelte";
  import Greeting from "$lib/Greeting/Greeting.svelte";
  import AppList from "$lib/AppList/index.svelte";
  import BookmarkList from "$lib/BookmarkList/index.svelte";
  import { apps, filteredApps, themes } from "$lib/stores";
  import type { PageData } from "./$types";
  import { onMount } from "svelte";

  export let data: PageData;
  $apps = data.apps;

  let showModal = false;

  $: darkMode = true;

  if (data.startpage.customThemes) {
    $themes.push(...(data.startpage.customThemes as Array<any>));
  }

  onMount(() => {
    darkMode =
      window.matchMedia &&
      window.matchMedia("(prefers-color-scheme: dark)").matches;

    const setValue = (property: string, value: string) => {
      if (value) {
        document.documentElement.style.setProperty(`--${property}`, value);
      }
    };

    const setTheme = (options: Record<string, string>) => {
      for (let option of Object.keys(options)) {
        const property = option;
        const value = options[option];

        setValue(property, value);
        localStorage.setItem(property, value);
      }
    };

    let restoreTheme: string;

    if (darkMode) {
      restoreTheme = data.startpage.darkTheme as string;
    } else {
      restoreTheme = data.startpage.lightTheme as string;
    }

    let themeData = $themes.find((t) => t.name === restoreTheme);

    setTheme({
      "color-background": themeData?.backgroundColor as string,
      "color-text-pri": themeData?.primaryColor as string,
      "color-text-acc": themeData?.accentColor as string,
    });
  });
</script>

<svelte:head>
  <title>{data.startpage.title}</title>
</svelte:head>

{#if showModal}
  <Modal settings={data.startpage} on:close={() => (showModal = false)} />
{/if}

{#if data.startpage.showSearch}
  <Search
    providers={data.startpage.searchProviders}
    defaultProvider={data.startpage.defaultSearchProvider}
  />
{/if}

{#if data.startpage.showGreeting}
  <Greeting name={data.startpage.name} />
{/if}

{#if data.startpage.showApps}
  <AppList
    apps={$filteredApps}
    showGroups={data.startpage.showAppGroups}
    defaultIcon={data.startpage.defaultAppIcon}
    showUrl={data.startpage.showAppUrls}
    showInfo={data.startpage.showAppInfo}
    showStatus={data.startpage.showAppStatus}
    targetBlank={data.startpage.alwaysTargetBlank}
  />
{/if}

{#if data.startpage.showBookmarks}
  <BookmarkList
    bookmarks={data.globalBookmarks}
    showGroups={data.startpage.showBookmarkGroups}
    targetBlank={data.startpage.alwaysTargetBlank}
  />
{/if}

{#if data.startpage.showGlobalBookmarks && data.slug !== ""}
  <BookmarkList
    header="Global Bookmarks"
    bookmarks={data.globalBookmarks}
    showGroups={data.startpage.showBookmarkGroups}
    targetBlank={data.startpage.alwaysTargetBlank}
  />
{/if}

<div id="modal_init">
  <a href="#modal" on:click|preventDefault={() => (showModal = true)}>
    <Icon icon="mdi:xbox-controller-menu" />
  </a>
</div>

<style>
  #modal_init a {
    z-index: 25;
    bottom: 1vh;
    color: var(--color-text-acc);
    left: 1vw;
    position: fixed;
    font-size: 2.5em;
  }

  #modal_init a:hover {
    color: var(--color-text-pri);
  }
</style>
