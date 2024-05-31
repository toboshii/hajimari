<script type="ts">
  import Icon from "@iconify/svelte";
  export let bookmarkGroup;
  export let targetBlank;
</script>

<div class="links_item">
  {#each bookmarkGroup.bookmarks || [] as link}
    <a
      href={link.url}
      target={link.targetBlank || targetBlank ? "_blank" : null}
      rel="external noopener noreferrer"
    >
      {#if link.icon}
        {#if link.icon.includes("//")}
          <img src={link.icon} alt="bookmark icon for {link.name}" />{link.name}
        {:else if link.icon.includes(":")}
          <Icon icon={link.icon} inline />{link.name}
        {/if}
      {:else}
        {link.name}
      {/if}
    </a>
  {/each}
</div>

<style>
  .links_item a {
    display: block;
    line-height: 2;
  }

  .links_item :global(svg) {
    width: 14px;
    height: 14px;
    margin-right: 2px;
  }

  .links_item img {
    width: 14px;
    height: 14px;
    margin-right: 2px;
    vertical-align: -0.125em;
  }
</style>
