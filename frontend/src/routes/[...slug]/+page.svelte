<script lang="ts">
	import Search from '$lib/Search/Search.svelte'
	import Greeting from '$lib/Greeting/Greeting.svelte';
	import AppList from '$lib/AppList/index.svelte';
	import BookmarkList from '$lib/BookmarkList/index.svelte';
	import { apps, filteredApps } from '$lib/stores';
	import type { PageData } from './$types';

	export let data: PageData;
	$apps = data.apps;
</script>

<svelte:head>
	<title>{data.startpage.title}</title>
</svelte:head>

<Search />

{#if data.startpage.showGreeting}
<Greeting name={data.startpage.name} />
{/if}

<AppList apps={$filteredApps} showGroups={data.startpage.showAppGroups} showUrls={data.startpage.showAppUrls} showInfo={data.startpage.showAppInfo}/>

<BookmarkList bookmarks={data.startpage.bookmarks} showGroups={data.startpage.showBookmarkGroups}/>

{#if data.startpage.showGlobalBookmarks && data.slug !== ""}
<BookmarkList header="Global Bookmarks" bookmarks={data.globalBookmarks} showGroups={data.startpage.showBookmarkGroups}/>
{/if}