<script lang="ts">
  import { fade } from "svelte/transition";
  import { flip } from "svelte/animate";
  import Icon from "@iconify/svelte";

  export let group: any;
  export let defaultIcon: string = "mdi:application";
  export let showUrl: boolean = true;
  export let showInfo: boolean = true;
  export let showStatus: boolean = true;
  export let targetBlank: boolean = false;
</script>

{#each group.apps || [] as app (app.name)}
  <div
    class="apps_item"
    in:fade|local={{ duration: 300 }}
    animate:flip={{ duration: 300 }}
  >
    <div class="apps_icon">
      <a
        href={app.url}
        target={app.targetBlank || targetBlank ? "_blank" : null}
        rel="external noopener noreferrer"
      >
        {#if app.icon.includes("//")}
          <img src={app.icon} alt="app icon for {app.name}" />
        {:else if app.icon.includes(":") || !app.icon}
          <Icon icon={app.icon ? app.icon : defaultIcon} color="{app.iconColor}" />
        {:else}
          <!-- support old icon format to ease transition -->
          <Icon icon="mdi:{app.icon ? app.icon : defaultIcon}" color="{app.iconColor}" />
        {/if}
      </a>
      {#if app.replicas.total > 0 && showStatus}
        <hr
          class="app_status"
          style="background-image: linear-gradient(to right, var(--color-text-acc) 0 {app
            .replicas.pctReady}%, currentcolor {app.replicas.pctReady}% 100%);"
        />
      {/if}
    </div>
    <div class="apps_text">
      <a
        href={app.url}
        target={app.targetBlank || targetBlank ? "_blank" : null}
        rel="external noopener noreferrer">{app.name}</a
      >
      {#if showUrl}
        <span class="app_address">{app.url}</span>
      {/if}
      {#if showInfo}
        <span class="app_info">{app.info}</span>
      {/if}
    </div>
  </div>
{/each}

<style>
  .apps_icon {
    height: 64px;
    margin-right: 1em;
    padding-top: 15px;
  }

  .apps_icon :global(svg) {
    font-size: 2.5em;
    line-height: 3rem;
  }

  .apps_icon img {
    width: 2.5em;
    height: 2.5em;
    padding: 2px;
  }

  .apps_item {
    display: flex;
    flex-direction: row;
    flex-wrap: wrap;
    height: 64px;
    margin: 0;
    padding-right: 10px;
  }

  .apps_text {
    display: flex;
    flex-direction: column;
    justify-content: center;
    flex: 1;
    overflow: hidden;
  }

  .apps_text a {
    font-size: 1em;
    font-weight: 500;
    text-transform: uppercase;
  }

  .apps_text span {
    color: var(--color-text-acc);
    font-size: 0.8em;
    text-transform: uppercase;
  }

  .app_address {
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis ellipsis;
  }

  .app_info {
    overflow-wrap: break-word;
  }

  .app_status {
    /* background-image: linear-gradient(to right, black 0 50%, transparent 50% 100%); */
    height: 1px;
    border: 0;
  }

  @media screen and (max-width: 667px) {
    .apps_icon {
      height: 64px;
      margin-right: 0.8em;
      padding-top: 14px;
    }

    .apps_icon :global(svg) {
      font-size: 2em;
      line-height: 2.5rem;
    }
  }

  @media only screen and (max-width: 400px) {
    .app_address {
      display: none;
    }
  }
</style>
