<script lang="ts">
    import App from './AppGroup.svelte';

	export let apps: any;
    export let showGroups: boolean;
    export let defaultIcon: string = 'mdi:application';
    export let showUrls: boolean = true;
</script>

{#if apps.length === 0}
    <div class="apps">
        <h3>Applications</h3>
        <p>No apps here...yet</p>
    </div>
{:else}
    <div class="apps">
        <h3>Applications</h3>
        <div class="apps_loop" class:grouped="{showGroups === true}">
            {#each apps as group}
                {#if showGroups === true}
                    <div class="links_item">
                        <h4>{group.group}</h4>
                        <div class="apps_group">
                            <App {group} showUrl={showUrls} defaultIcon={defaultIcon}/>
                        </div>
                    </div>
                {:else}
                    <App {group} showUrl={showUrls} defaultIcon={defaultIcon}/>
                {/if}
            {/each}
        </div>
    </div>
{/if}

<style>
    .apps_loop {
        display: grid;
        grid-column-gap: 0px;
        grid-row-gap: 0px;
        grid-template-columns: 1fr 1fr 1fr 1fr;
        grid-template-rows: 64px;
        padding-bottom: var(--module-spacing);
    }

    .apps_loop.grouped {
        grid-template-columns: 1fr 1fr;
        grid-template-rows: auto;
    }

    .apps_group {
        display: grid;
        grid-template-columns: 1fr 1fr 1fr;
    }

    .links_item h4 {
        color: var(--color-text-acc);
    }

    @media screen and (max-width: 1260px) {
        .apps_loop {
            grid-template-columns: 1fr 1fr 1fr;
            width: 90vw;
        }

        .apps_group {
            grid-template-columns: 1fr 1fr;
        }
    }

    @media screen and (max-width: 667px) {
        .apps_loop{
            grid-column-gap: 0px;
            grid-row-gap: 0px;
            grid-template-columns: 1fr 1fr;
            width: 90vw;
        }

        .apps_group {
            grid-template-columns: 1fr;
        }
    }
</style>