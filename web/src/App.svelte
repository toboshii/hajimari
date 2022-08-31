<script>
	import Header from './header.svelte';
	import AppList from './AppList/index.svelte';
	import BookmarkList from './BookmarkList/index.svelte';

	let greeting = 'Good evening, Toboshii!';
	let date = 'fri, aug 20';

	import { onMount } from "svelte";
	import { apiData, name, apps, bookmarks } from './store.js';

	onMount(async () => {
		fetch("http://localhost:3000/apps")
		.then(response => response.json())
		.then(data => {
			console.log(data);
			apps.set(data);
		}).catch(error => {
			console.log(error);
			return [];
		});

		fetch("http://localhost:3000/bookmarks")
		.then(response => response.json())
		.then(data => {
			console.log(data);
			bookmarks.set(data);
		}).catch(error => {
			console.log(error);
			return [];
		});
	});
</script>

<main id="container" class="fade">
	<section id="search">
		<input name="keywords" type="text" id="keywords" size="50" spellcheck="false" autofocus="true" onkeydown="handleKeyPress(event)">
	</section>

	<Header {greeting} {date} />
	
	<AppList apps = {$apps} />

	<BookmarkList bookmarks = {$bookmarks} />

	<!-- <section id="modules">
		<h3>Modules</h3>
		{{range .Modules}}
			<p><span style="color: var(--color-text-acc);">{{.Name}}</span> - {{.Output}}</p>
		{{end}}
	</section> -->

	<!-- <section id="apps">
		<h3>Applications</h3>
		<div id="apps_loop">
			{{range .Apps}}
				<div class="apps_item">
					<div class="apps_icon">
						<span class="iconify icon" data-icon="mdi-{{or .Icon "application"}}"></span>
					</div>
					<div class="apps_text">
						<a href="{{.URL}}">{{.Name}}</a>
						<span id="app-address">{{.URL}}</span>
					</div>
				</div>
			{{end}}
		</div>
	</section>

	 <section id="links">
		<h3>Bookmarks</h3>
		<div id="links_loop">
			{{range .Groups}}
				<div id="links_item">
					<h4>{{.Name}}</h4>
					{{range .Links}}
						<a href="{{.URL}}" class="theme_color-border theme_text-select">{{.Name}}</a>
					{{end}}
				</div>
			{{end}}
		</div>
	</section> -->
</main>

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

	@media screen and (max-width: 1260px)
	{
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

	@media screen and (max-width: 667px)
	{
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