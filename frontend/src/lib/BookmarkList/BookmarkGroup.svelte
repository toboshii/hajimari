<script>
    import Icon from "@iconify/svelte";
    export let bookmarkGroup;
</script>

<div class="links_item">
    {#each bookmarkGroup.bookmarks as link}
        <a href={link.url} class="theme_color-border theme_text-select"
            >{link.name}</a
        >
    {/each}
</div>

<style>
    .links_item a {
        display: block;
        line-height: 2;
    }
</style>
