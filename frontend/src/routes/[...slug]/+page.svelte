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

	$: darkMode =
		window.matchMedia &&
		window.matchMedia("(prefers-color-scheme: dark)").matches;

	if (data.startpage.customThemes) {
		$themes.push(...(data.startpage.customThemes as Array<any>));
	}

	onMount(() => {
		const setValue = (property: string, value: string) => {
			if (value) {
				document.documentElement.style.setProperty(
					`--${property}`,
					value
				);
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

		console.log(restoreTheme);

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
	<!-- <style>
		:root{
			--color-background: #232530;
			--color-text-pri: #FAB795;
			--color-text-acc: #E95678;
			--font: -apple-system, BlinkMacSystemFont, Helvetica Neue, Roboto, sans-serif;
			--font-code: monospace;
			--font-mono: monospace;
		}
	</style> -->
</svelte:head>

{#if showModal}
	<Modal settings={data.startpage} on:close={() => (showModal = false)} />
{/if}

<Search />

{#if data.startpage.showGreeting}
	<Greeting name={data.startpage.name} />
{/if}

<AppList
	apps={$filteredApps}
	showGroups={data.startpage.showAppGroups}
	showUrls={data.startpage.showAppUrls}
	showInfo={data.startpage.showAppInfo}
/>

<BookmarkList
	bookmarks={data.startpage.bookmarks}
	showGroups={data.startpage.showBookmarkGroups}
/>

{#if data.startpage.showGlobalBookmarks && data.slug !== ""}
	<BookmarkList
		header="Global Bookmarks"
		bookmarks={data.globalBookmarks}
		showGroups={data.startpage.showBookmarkGroups}
	/>
{/if}

<div id="modal_init">
	<a href="#modal" on:click|preventDefault={() => (showModal = true)}>
		<Icon icon="mdi:xbox-controller-menu" />
	</a>
</div>

<style>
	#modal_init a {
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
