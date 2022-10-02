<script>
    import { page } from '$app/stores';
    import { dev } from '$app/environment';
    // we don't want to use <svelte:window bind:online> here,
    // because we only care about the online state when
    // the page first loads
    const online = typeof navigator !== 'undefined' ? navigator.onLine : true;
  </script>

<svelte:head>
	<title>{$page.status}</title>
</svelte:head>

{#if online}
    <section id="error">
        <h1>Yikes!</h1>

        {#if $page.error.message}
            <p class="error">{$page.status}: {$page.error.message}</p>
        {:else}
            <p class="error">Encountered a {$page.status} error</p>
        {/if}

        {#if dev && $page.error.stack}
            <pre>{$page.error.stack}</pre>
        {:else}
            {#if $page.status >= 500}
                <p>Please try reloading the page.</p>
            {/if}

            <p>
                If the error persists, please drop by the <a target="_blank" rel="noopener noreferrer" href="https://discord.gg/NswQwRQQ">Discord</a
                >
                and let us know, or raise an issue on
                <a target="_blank" rel="noopener noreferrer" href="https://github.com/toboshii/hajimari/issues/new/choose">GitHub</a>. Thanks!
            </p>
        {/if}
    </section>
{:else}
    <section>
        <h1>It looks like you're offline</h1>

        <p>Reload the page once you've found the internet.</p>
    </section>
{/if}

<style>
	.error {
		background-color: var(--color-text-pri);
		color: var(--color-text-acc);
		padding: 15px;
		font: 600 1.5em var(--font);
	}

    pre {
        white-space: break-spaces;
    }

</style>