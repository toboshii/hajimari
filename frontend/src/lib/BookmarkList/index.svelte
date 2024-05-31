<script lang="ts">
  import BookmarkGroup from "./BookmarkGroup.svelte";

  export let header: string = "Bookmarks";
  export let bookmarks: [];
  export let showGroups: boolean;
  export let targetBlank: boolean;
</script>

{#if bookmarks.length === 0}
  <div class="links">
    <h3>{header}</h3>
    <p>No bookmarks here...yet</p>
  </div>
{:else}
  <div class="links">
    <h3>{header}</h3>
    <div class="links_loop">
      {#each bookmarks as bookmarkGroup}
        {#if showGroups}
          <div class="links_item">
            <h4>{bookmarkGroup.group}</h4>
            <BookmarkGroup {bookmarkGroup} {targetBlank} />
          </div>
        {:else}
          <BookmarkGroup {bookmarkGroup} {targetBlank} />
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
    .links_loop {
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
