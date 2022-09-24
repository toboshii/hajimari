<script lang="ts">
    import { createEventDispatcher, onMount } from "svelte";
    import { themes } from "$lib/stores";

    import Icon from "@iconify/svelte";
    import yaml from "js-yaml";

    export let settings: Record<string, unknown>;

    $: darkMode =
        window.matchMedia &&
        window.matchMedia("(prefers-color-scheme: dark)").matches;

    const dispatch = createEventDispatcher();

    onMount(() => {
        const setValue = (property: string, value: string) => {
            if (value) {
                document.documentElement.style.setProperty(
                    `--${property}`,
                    value
                );
            }
        };

        const setTheme = (options: Record<string, string>) => {
            for (let option of Object.keys(options)) {
                const property = option;
                const value = options[option];

                setValue(property, value);
            }
        };

        const dataThemeButtons = document.querySelectorAll("[data-theme]");

        for (let i = 0; i < dataThemeButtons.length; i++) {
            dataThemeButtons[i].addEventListener("click", () => {
                const theme: string = dataThemeButtons[i].dataset.theme;

                if (darkMode) {
                    settings.darkTheme = theme;
                } else {
                    settings.lightTheme = theme;
                }

                let themeData = $themes.find((t) => t.name === theme);

                setTheme({
                    "color-background": themeData?.backgroundColor as string,
                    "color-text-pri": themeData?.primaryColor as string,
                    "color-text-acc": themeData?.accentColor as string,
                });

                dispatch("updateEditor", yaml.dump(settings));
            });
        }
    });
</script>

<div class="theme-mode-selector">
    {#if darkMode}
        <h2>Dark Theme</h2>
    {:else}
        <h2>Light Theme</h2>
    {/if}
    <button on:click={() => (darkMode = !darkMode)}
        ><Icon icon="mdi:theme-light-dark" height="24px" /></button
    >
</div>

<div class="theme-selector">
    {#each $themes as theme}
        <button
            data-theme={theme.name}
            class="theme-button"
            style="background-color: {theme.backgroundColor}; border: 1px solid {theme.accentColor}; color: {theme.primaryColor};"
            >{theme.name}</button
        >
    {/each}
</div>

<style>

    .theme-button{
        font-size: 0.8em;
        margin: 2px;
        width:128px;
        line-height: 3em;
        text-align: center;
        text-transform: uppercase;
    }

    button {
        display: block;
    }

    .theme-mode-selector h2,
    :global(.theme-mode-selector svg) {
        display: inline;
        vertical-align: middle;
    }

    .theme-mode-selector button {
        margin-left: 10px;
        display: inline-block;
        color: var(--color-text-acc);
        text-transform: uppercase;
        padding: 5px 5px;
        background-color: var(--color-background);
        border: 1px solid var(--color-text-acc) !important;
        transition: all 0.3s ease 0s;
    }

    .theme-mode-selector button:hover {
        background-color: var(--color-text-acc);
        color: var(--color-background);
        border: 1px solid var(--color-background) !important;
    }

    /* .help {
        color: var(--color-text-acc);
        font-size: 0.9em;
        margin-top: 0.4em;
    } */

    .theme-selector {
        margin-top: 10px;
        border-bottom: 0px solid var(--color-text-acc);
        display: flex;
        flex-wrap: wrap;
        margin-bottom: 2em;
    }
</style>
