<script>
	// import { session } from '$app/stores';
	import { group_outros } from 'svelte/internal';
import AppGroup from './AppGroup.svelte';

	export let apps;
    let config = {'groupApps': false};
</script>
<div class="apps">
    <h3>Applications</h3>
    {#if apps.length === 0}
        <p>No apps here...yet</p>
    {:else}
        <div class="apps_loop" class:grouped="{config.groupApps === true}">
            {#each apps as group}
                {#if config.groupApps === true}
                    <div class="links_item">
                        <h4>{group.name}</h4>
                        <AppGroup {group} />
                    </div>
                {:else}
                    <AppGroup {group} />
                {/if}
            {:else}
                <div>??</div>
            {/each}
        </div>
    {/if}
</div>
<style>
    .apps_loop {
        border-bottom: 0px solid var(--color-text-acc);
        display: grid;
        grid-column-gap: 0px;
        grid-row-gap: 0px;
        grid-template-columns: 1fr 1fr 1fr 1fr;
        grid-template-rows: 64px;
        padding-bottom: var(--module-spacing);
    }

    .apps_loop.grouped {
        grid-template-rows: auto;
    }

    .links_item h4 {
        color: var(--color-text-acc);
    }

    @media screen and (max-width: 1260px) {
        .apps_loop {
            grid-template-columns: 1fr 1fr 1fr;
            width: 90vw;
        }
    }

    @media screen and (max-width: 667px) {
        .apps_loop{
            grid-column-gap: 0px;
            grid-row-gap: 0px;
            grid-template-columns: 1fr 1fr;
            width: 90vw;
        }
    }
</style>
