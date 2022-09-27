<script lang="ts">
    import { createEventDispatcher, onDestroy, onMount } from "svelte";

    import { page } from "$app/stores";
    import { portal } from "svelte-portal/src/Portal.svelte";
    import Icon from "@iconify/svelte";

    import { api } from "$lib/api";
    import CodeMirror from "$lib/CodeMirror/CodeMirror.svelte";
    import yaml from "js-yaml";
    import { error } from "@sveltejs/kit";
    import ThemeSwitcher from "$lib/ThemeSwitcher/ThemeSwitcher.svelte";

    export let settings = {};

    const slug = $page.params.slug;

    let theme = "dracula";
    let editor: CodeMirror;

    onMount(async () => {
        await editor.set(yaml.dump(settings), "yaml");
    });

    const dispatch = createEventDispatcher();
    const close = () => dispatch("close");

    const handle_keydown = (e: KeyboardEvent) => {
        if (e.key === "Escape") {
            close();
            return;
        }
    };

    async function handleSave() {
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

<svelte:window on:keydown={handle_keydown} />

<div class="modal-background" on:click={close} use:portal={"body"} hidden />

<div class="modal" role="dialog" aria-modal="true" use:portal={"body"} hidden>
    <header>
        <h1>Options</h1>
        <a
            href="#close"
            title="Close"
            class="modal-close"
            on:click|preventDefault={close}
        >
            <Icon icon="mdi:close" />
        </a>
    </header>

    <!-- <section id="providers">
        <table>
            <tr>
                <th>Website</th>
                <th>Prefix</th>
            </tr>
            {{range .Providers}}
            <tr>
                <td><a href="{{.URL}}">{{.Name}}</a></td>
                <td>{{.Prefix}}</td>
            </tr>
            {{end}}
        </table>
    </section> -->
    <ThemeSwitcher bind:settings on:updateEditor={handleUpdateEditor} />
    <section class="config">
        <div class="editor">
            <CodeMirror bind:this={editor} {theme} />
            <!-- <div class="help"><u>Ctrl-F</u>: Search <u>Alt-G</u>: Go to line</div> -->
        </div>
    </section>

    <footer>
        <a href="https://github.com/toboshii/hajimari"
            ><Icon icon="mdi:github-box" /></a
        >
        <a href="https://discord.gg/NswQwRQQ"><Icon icon="mdi:discord" /></a>
        <a href="https://icon-sets.iconify.design"
            ><Icon icon="mdi:material-design" /></a
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

    /* .modal h2 {
        color: var(--color-text-acc);
        margin-top: 1.5em;
    } */

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
        color: rgba(0, 0, 0, 0.5);
    }

    .modal-close {
        color: var(--color-text-acc);
        font-size: 1.5em;
        text-align: center;
        text-decoration: none;
    }

    .modal-close:hover {
        color: var(--color-text-pri);
    }

    .editor {
        height: 35vh;
        min-height: 350px;
    }

    /* .help {
        color: var(--color-text-acc);
        font-size: 0.9em;
        margin-top: 0.4em;
    } */

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
