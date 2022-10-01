<script type="ts">
    import Icon from "@iconify/svelte";
    export let bookmarkGroup;
    export let targetBlank: boolean;
</script>

<div class="links_item">
    {#each bookmarkGroup.bookmarks as link}
        <a href={link.url} target="{link.targetBlank == "" ? (targetBlank ? "_blank" : "_self") : (link.targetBlank == "1" ? "_blank" : "_self")}" rel="noopener noreferrer"
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
