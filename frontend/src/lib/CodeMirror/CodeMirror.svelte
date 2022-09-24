<script lang="ts">
	import { onMount } from "svelte";
	import { writable } from "svelte/store";

	export let readonly: boolean = false;
	export let lineNumbers: boolean = true;
	export let tab: boolean = true;
	export let theme: string = "dracula";

	let w: number;
	let h: number;
	let code = "";
	let mode: string;
	let editor: CodeMirror.EditorFromTextArea;
	let destroyed: boolean = false;
	let CodeMirrorMounted: CodeMirror;
	let refs: { editor: HTMLTextAreaElement | null } = { editor: null };

	export function get() {
		return editor.getValue();
	}

	export async function set(new_code: string, new_mode: string) {
		if (new_mode !== mode) {
			await createEditor((mode = new_mode));
		}

		code = new_code;
		if (editor) editor.setValue(code);
	}

	export function update(new_code: string) {
		code = new_code;

		if (editor) {
			const { left, top } = editor.getScrollInfo();
			editor.setValue((code = new_code));
			editor.scrollTo(left, top);
		}
	}

	export function resize() {
		editor.refresh();
	}

	export function focus() {
		editor.focus();
	}

	export function setCursor(pos: number) {
		if (editor) editor.setCursor(pos);
	}

	export const cursorIndex = writable(0);

	const modes: { [key: string]: {} } = {
		js: {
			name: "javascript",
			json: false,
		},
		json: {
			name: "javascript",
			json: true,
		},
		yaml: {
			name: "yaml",
		},
		md: {
			name: "markdown",
		},
	};

	$: if (editor && w && h) {
		editor.refresh();
	}

	onMount(() => {
		(async () => {
			if (!CodeMirrorMounted) {
				let mod = await import("$lib/CodeMirror/codemirror");
				CodeMirrorMounted = mod.default;
			}
			await createEditor(mode || "yaml");
			if (editor) editor.setValue(code || "");
		})();

		return () => {
			destroyed = true;
			if (editor) editor.toTextArea();
		};
	});

	let first = true;

	async function createEditor(mode: string) {
		if (destroyed || !CodeMirrorMounted) return;

		if (editor) editor.toTextArea();

		const opts = {
			lineNumbers,
			lineWrapping: true,
			indentWithTabs: false,
			smartIndent: true,
			indentUnit: 2,
			tabSize: 2,
			value: "",
			mode: modes[mode] || {
				name: mode,
			},
			readOnly: readonly,
			styleActiveLine: true,
			autoCloseBrackets: true,
			autoCloseTags: true,
			extraKeys: CodeMirrorMounted.normalizeKeyMap({
				Enter: "newlineAndIndentContinueMarkdownList",
				"Ctrl-Q": function (cm: CodeMirror.Editor) {
					cm.foldCode(cm.getCursor());
				},
				"Cmd-Q": function (cm: CodeMirror.Editor) {
					cm.foldCode(cm.getCursor());
				},
				Tab: (cm: CodeMirror.Editor) => {
					if (cm.getMode().name === "null") {
						cm.execCommand("insertTab");
					} else {
						if (cm.somethingSelected()) {
							cm.execCommand("indentMore");
						} else {
							cm.execCommand("insertSoftTab");
						}
					}
				},
				"Shift-Tab": (cm: CodeMirror.Editor) =>
					cm.execCommand("indentLess"),
				// allow escaping the CodeMirror with Esc Tab
				"Esc Tab": false,
			}),
			foldGutter: true,
			gutters: [
				"CodeMirror-linenumbers",
				"CodeMirror-foldgutter",
				"CodeMirror-lint-markers",
			],
			lint: true,
			theme,
		};

		if (!tab) {
			opts.extraKeys["Tab"] = tab;
			opts.extraKeys["Shift-Tab"] = tab;
		}

		// Creating a text editor is a lot of work, so we yield
		// the main thread for a moment. This helps reduce jank
		if (first) await sleep(50);

		if (destroyed) return;

		editor = CodeMirrorMounted.fromTextArea(refs.editor, opts);

		if (first) await sleep(50);
		editor.refresh();

		first = false;
	}

	function sleep(ms: number) {
		return new Promise((fulfil) => setTimeout(fulfil, ms));
	}
</script>

<svelte:head>
	<script
		src="//cdnjs.cloudflare.com/ajax/libs/js-yaml/4.1.0/js-yaml.min.js"></script>
</svelte:head>

<div class="codemirror-container" bind:offsetWidth={w} bind:offsetHeight={h}>
	<textarea bind:this={refs.editor} readonly value={code} />

	{#if !CodeMirrorMounted}
		<pre style="position: absolute; left: 0; top: 0">{code}</pre>
	{/if}
</div>

<style>
	* {
		font-family: monospace;
	}

	.codemirror-container {
		position: relative;
		width: 100%;
		height: 100%;
		border: none;
		line-height: 1.5;
		overflow: hidden;
	}

	.codemirror-container :global(.CodeMirror) {
		height: 100%;
		font: 400 var(--font-code) / 1.7 var(--font-mono);
	}

	.codemirror-container :global(.error-line) {
		background-color: rgba(200, 0, 0, 0.05);
	}

	.codemirror-container :global(.mark-text) {
		background-color: var(--highlight);
	}

	.codemirror-container {
		border: 1px solid var(--color-text-acc) !important;
	}

	.codemirror-container :global(.CodeMirror-vscrollbar) {
		scrollbar-width: thin;
	}

	.codemirror-container :global(.CodeMirror-hscrollbar) {
		scrollbar-width: thin;
	}

	textarea {
		visibility: hidden;
	}

	pre {
		background-color: #282a36;
		position: absolute;
		width: 100%;
		height: 100%;
		top: 0;
		left: 0;
		border: none;
		padding: 4px 4px 4px 60px;
		resize: none;
		font-family: var(--font-mono);
		font-size: 13px;
		line-height: 1.7;
		user-select: none;
		pointer-events: none;
		color: var(--color-text-pri);
		tab-size: 2;
		-moz-tab-size: 2;
	}
</style>
