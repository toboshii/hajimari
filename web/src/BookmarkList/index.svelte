<script>
	// import { session } from '$app/stores';
	import BookmarkGroup from './BookmarkGroup.svelte';

	export let bookmarks;
    let config = {'groupBookmarks': true}
</script>

{#if bookmarks.length === 0}
    <div class="links">
        <h3>Bookmarks</h3>
        <p>No bookmarks here...yet</p>
    </div>
{:else}
    <div class="links">
        <h3>Bookmarks</h3>
        <div class="links_loop">
            {#each bookmarks as group}
                {#if config.groupBookmarks === true}
                    <div class="links_item">
                        <h4>{group.name}</h4>
                        <BookmarkGroup {group} />
                    </div>
                {:else}
                    <BookmarkGroup {group} />
                {/if}
            {/each}
        </div>
    </div>
{/if}

<style>
    .links_loop {
        display: grid;
        flex-wrap: nowrap;
        grid-column-gap: 20px;
        grid-row-gap: 0px;
        grid-template-columns: 1fr 1fr 1fr 1fr;
        grid-template-rows: auto;
    }

    .links_item {
        line-height: 1.5rem;
        margin-bottom: 2em;
        webkit-font-smoothing: antialiased;
    }

    .links_item h4 {
        color: var(--color-text-acc);
    }

    @media screen and (max-width: 1260px) {
        .links_loop	{
            grid-template-columns: 1fr 1fr 1fr;
        }
    }

    @media screen and (max-width: 667px) {
        .links_loop {
            display: grid;
            flex-wrap: nowrap;
            grid-column-gap: 20px;
            grid-row-gap: 0px;
            grid-template-columns: 1fr 1fr;
            grid-template-rows: auto;
        }
    }
</style>