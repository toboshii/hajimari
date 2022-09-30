<script lang="ts">
    import { createEventDispatcher, onMount } from "svelte";
    import { themes } from "$lib/stores";

    import Icon from "@iconify/svelte";
    import yaml from "js-yaml";

    export let settings: Record<string, unknown>;
    export let showHelp: boolean;

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
{#if showHelp}
    <span class="help">Setting theme for <b>{darkMode ? 'dark' : 'light'}</b> mode. Active mode will be chosen based on the browser's preference</span>
{/if}
<button
    class="theme-mode"
    class:dark={darkMode}
    on:click={() => (darkMode = !darkMode)}
    ><Icon icon="mdi:theme-light-dark" inline /> {darkMode ? 'dark' : 'light'}</button
>
<div class="theme-selector">
    {#each $themes as theme}
        <button
            data-theme={theme.name}
            style="--color-background: {theme.backgroundColor}; --color-text-acc: {theme.accentColor}; --color-text-pri: {theme.primaryColor};"
            >{theme.name}</button
        >
    {/each}
</div>

<style>
    .theme-mode {
        font: var(--font);
        color: var(--color-text-acc);
        text-transform: uppercase;
        padding: 5px 5px;
        background-color: var(--color-background);
        border: 1px solid var(--color-text-acc) !important;
        transition: all 0.3s ease 0s;
        float: right;
        font-size: 1em;
        margin: 1em 0.125em 0 0;
    }

    .theme-mode:hover {
        background-color: var(--color-text-acc);
        color: var(--color-background);
        border: 1px solid var(--color-background) !important;
    }

    .theme-mode.dark {
        color: var(--color-background);
        background-color: var(--color-text-acc);
    }

    .theme-selector {
        border-bottom: 0px solid var(--color-text-acc);
        display: flex;
        flex-wrap: wrap;
    }

    .theme-selector button {
        font-family: var(--font);
        font-weight: 700;
        font-size: 0.8em;
        margin: 2px;
        width:128px;
        line-height: 3em;
        text-align: center;
        text-transform: uppercase;
        color: var(--color-text-pri);
        border: 1px solid var(--color-text-acc);
        background-color: var(--color-background);
        transition: all 0.3s ease 0s;
    }

    .theme-selector button:hover {
        background-color: var(--color-text-acc);
        color: var(--color-background);
        border: 1px solid var(--color-background) !important;
    }

    .help {
        color: var(--color-text-pri);
        font-size: 0.9em;
        display: inline-block;
        margin-left: 0.5em;
    }
</style>
