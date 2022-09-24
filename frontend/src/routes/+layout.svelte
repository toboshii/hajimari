<script lang="ts">
	import "../app.css";
	import { fade } from "svelte/transition";
	import { onMount } from "svelte";

	let visible = false;

	onMount(() => {
		visible = true;
	});
</script>

{#if visible}
	<main id="container" transition:fade>
		<slot />

		<!-- <section id="modules">
			<h3>Modules</h3>
			{{range .Modules}}
				<p><span style="color: var(--color-text-acc);">{{.Name}}</span> - {{.Output}}</p>
			{{end}}
		</section> -->
	</main>
{/if}

<style>
	main {
		align-items: stretch;
		display: grid;
		grid-column-gap: 20px;
		grid-row-gap: 3vh;
		grid-template-columns: 1fr;
		grid-template-rows: 8vh auto;
		justify-items: stretch;
		margin-left: auto;
		margin-right: auto;
		margin-top: 5vh;
		width: 60%;
	}

	@media screen and (max-width: 1260px) {
		main {
			align-items: stretch;
			display: grid;
			grid-column-gap: 10px;
			grid-row-gap: 0px;
			grid-template-columns: 1fr;
			grid-template-rows: 80px auto;
			justify-items: stretch;
			margin-bottom: 1vh;
			margin-left: auto;
			margin-right: auto;
			width: 90%;
		}
	}

	@media screen and (max-width: 667px) {
		main {
			align-items: stretch;
			display: grid;
			grid-column-gap: 20px;
			grid-row-gap: 0px;
			grid-template-columns: 1fr;
			grid-template-rows: 80px auto;
			justify-items: stretch;
			margin-bottom: 1vh;
			width: 90%;
		}
	}
</style>
