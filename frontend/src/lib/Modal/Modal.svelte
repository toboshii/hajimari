<script lang="ts">
    import { createEventDispatcher, onDestroy, onMount } from "svelte";

    import { page } from "$app/stores";
    import { portal } from "svelte-portal";
    import Icon from "@iconify/svelte";

    import { api } from "$lib/api";
    import CodeMirror from "$lib/CodeMirror/CodeMirror.svelte";
    import yaml from "js-yaml";
    import { scale, slide } from "svelte/transition";
    import ThemeSwitcher from "$lib/ThemeSwitcher/ThemeSwitcher.svelte";

    export let settings = {};

    const slug = $page.params.slug;

    let theme: string = "dracula";
    let editor: CodeMirror;
    let showHelp: boolean;

    onMount(async () => {
        await editor.set(yaml.dump(settings), "yaml");
    });

    const dispatch = createEventDispatcher();
    const close = () => dispatch("close");

    const handleKeydown = (e: KeyboardEvent) => {
        if (e.key === "Escape") {
            close();
            return;
        }
    };

    const help = () => (showHelp = showHelp ? false : true);

    async function handleSave(): Promise<void> {
        let method = slug ? "put" : "post";
        let newSettings = yaml.load(editor.get());
        const response = await api(
            self.fetch,
            method,
            `startpage/${slug}`,
            newSettings as Record<string, unknown>
        );

        if (response.status >= 200 && response.status <= 299) {
            let jsonResponse = await response.json();
            location.assign(jsonResponse.id);
        } else {
            let jsonResponse = await response.json();
            // todo: pretty messages
            alert(
                `Error: ${jsonResponse.status}\n\nReason: ${jsonResponse.error}`
            );
        }
    }

    function handleUpdateEditor() {
        editor.update(yaml.dump(settings));
    }
</script>

<svelte:window on:keydown={handleKeydown} />

<div class="modal-background" on:click={close} use:portal={"body"} hidden />

<div
    class="modal"
    role="dialog"
    aria-modal="true"
    use:portal={"body"}
    hidden
    transition:slide
>
    <header>
        <h1>Settings</h1>
        <div>
            <a
                href="#close"
                title="Close"
                class="modal-close"
                on:click|preventDefault={close}
            >
                <Icon icon="mdi:close" />
            </a>
            <a
                href="#help"
                title="Help"
                class="modal-help"
                on:click|preventDefault={help}
            >
                <Icon icon="mdi:help-circle-outline" />
            </a>
        </div>
    </header>

    {#if showHelp}
        <section id="providers" transition:scale>
            <h2>Search</h2>
            <span class="help">The default provider can be changed using <pre>defaultSearchProvider</pre> in the configuration</span>
            <table>
                <tr>
                    <th>Provider</th>
                    <th>Token</th>
                </tr>
                <tr>
                    <td><Icon icon="mdi:apps" inline /> Apps</td>
                    <td>/</td>
                </tr>
                {#each settings.searchProviders as provider}
                    <tr>
                        <td
                            ><a
                                href={provider.url}
                                target="_blank"
                                rel="external noopener noreferrer"
                                ><Icon icon={provider.icon} inline />
                                {provider.name}</a
                            ></td
                        >
                        <td>@{provider.token}</td>
                    </tr>
                {/each}
            </table>
        </section>
    {/if}
    
    <section id="themes">
        <h2>Themes</h2>
        <ThemeSwitcher bind:settings on:updateEditor={handleUpdateEditor} {showHelp} />
    </section>

    <section id="config">
        <h2>Configuration</h2>
        {#if showHelp}
            <span class="help"
                ><b>Ctrl-F/G:</b> Find <b>Alt-G:</b> Goto line <b>Alt-R:</b> Fold
                group</span
            >
        {/if}
        <div class="editor">
            <CodeMirror bind:this={editor} {theme} />
        </div>
    </section>

    <footer>
        <a
            href="https://github.com/ullbergm/hajimari"
            target="_blank"
            rel="external noopener noreferrer"><Icon icon="mdi:github-box" /></a
        >
        <a
            href="https://discord.gg/HWGZSWJsA8"
            target="_blank"
            rel="external noopener noreferrer"><Icon icon="mdi:discord" /></a
        >
        <a
            href="https://icon-sets.iconify.design"
            target="_blank"
            rel="external noopener noreferrer"><Icon icon="mdi:material-design" /></a
        >
        <button
            on:click={handleSave}
            type="button"
            id="save-config"
            class="save-button">Save</button
        >
    </footer>
</div>

<style>
    .modal-background {
        position: fixed;
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
    }

    .modal {
        position: fixed;
        left: 50%;
        top: 50%;
        max-width: 982px;
        max-height: 70vh;
        overflow-y: auto;
        transform: translate(-50%, -50%);
        padding: 2em;
        background: var(--color-background);
        border: 1px solid var(--color-text-acc);
        box-shadow: 0 14px 28px rgba(0, 0, 0, 0.3),
            0 15px 12px rgba(0, 0, 0, 0.25);
        z-index: 20;
        display: flex;
        flex-direction: column;
    }

    @media screen and (max-width: 910px) {
        .modal {
            left: 0;
            top: 0;
            transform: none;
            height: 100vh;
            max-height: 100vh;
        }
    }

    .modal h1 {
        color: var(--color-text-pri);
        font-size: 2em;
    }

    .modal h2 {
        color: var(--color-text-acc);
        font-size: 1.2em;
        margin-top: 1em;
        display: inline-block;
    }

    header {
        display: flex;
        justify-content: space-between;
    }

    footer {
        margin-top: 15px;
        font-size: 2em;
    }

    footer a {
        margin-right: 0.25em;
        color: rgba(0, 0, 0, 0.35);
        transition: all 0.3s ease 0s;
    }

    footer a:hover {
        color: var(--color-text-pri);
    }

    .modal-close,
    .modal-help {
        color: var(--color-text-acc);
        font-size: 1.5em;
        text-align: center;
        text-decoration: none;
        float: right;
        margin-left: 0.8em;
    }

    .modal-close:hover,
    .modal-help:hover {
        color: var(--color-text-pri);
    }

    .editor {
        height: 35vh;
        min-height: 350px;
    }

    .help {
        color: var(--color-text-pri);
        font-size: 0.9em;
        display: inline-block;
        margin-left: 0.5em;
    }

    .help pre {
        display: inline;
    }

    /* #providers {
        margin-bottom: 2em;
    } */

    #save-config {
        float: right;
        color: var(--color-text-acc);
        text-transform: uppercase;
        padding: 5px 10px;
        background-color: var(--color-background);
        border: 1px solid var(--color-text-acc) !important;
        transition: all 0.3s ease 0s;
    }

    #save-config:hover {
        background-color: var(--color-text-acc);
        color: var(--color-background);
        border: 1px solid var(--color-background) !important;
    }
</style>
