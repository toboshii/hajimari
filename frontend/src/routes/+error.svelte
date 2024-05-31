<script>
    import { page } from '$app/stores';
    import { dev } from '$app/environment';
    import Icon from "@iconify/svelte";
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
        
        <nav>
            <a href="/"><Icon icon="mdi:home" inline />Home</a>
        </nav>

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
                If the error persists, please drop by the <a target="_blank" rel="external noopener noreferrer" href="https://discord.gg/HWGZSWJsA8">Discord</a
                >
                and let us know, or raise an issue on
                <a target="_blank" rel="external noopener noreferrer" href="https://github.com/ullbergm/hajimari/issues/new/choose">GitHub</a>. Thanks!
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

    nav {
        font: 600 1em var(--font);
        text-transform: uppercase;
    }
    nav :global(svg) {
        margin-right: 2px;
    }

    pre {
        white-space: break-spaces;
    }

</style>