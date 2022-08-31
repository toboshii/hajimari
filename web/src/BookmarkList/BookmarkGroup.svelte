<script>
    export let group;
</script>

<div class="links_item">
    {#each group.bookmarks as link}
        <a href="{link.url}" class="theme_color-border theme_text-select">{link.name}</a>
    {/each}
</div>

<style>
    .links_item a {
        display: block;
        line-height: 2;
    }
</style>