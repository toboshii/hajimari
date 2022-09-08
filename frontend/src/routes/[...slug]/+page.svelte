<script lang="ts">
	import Greeting from '$lib/Greeting/Greeting.svelte';
	import AppList from '$lib/AppList/index.svelte';
	import BookmarkList from '$lib/BookmarkList/index.svelte';
	import type { PageData } from './$types';

	export let data: PageData;
</script>

<svelte:head>
	<title>{data.startpage.title}</title>
</svelte:head>

{#if data.startpage.showGreeting}
<Greeting name={data.startpage.name} />
{/if}

<AppList apps={data.apps} showGroups={data.startpage.showAppGroups} showUrls={data.startpage.showAppUrls}/>

<BookmarkList bookmarks={data.startpage.bookmarks} showGroups={data.startpage.showBookmarkGroups}/>

{#if data.startpage.showGlobalBookmarks && data.slug !== ""}
<BookmarkList header="Global Bookmarks" bookmarks={data.globalBookmarks} showGroups={data.startpage.showBookmarkGroups}/>
{/if}