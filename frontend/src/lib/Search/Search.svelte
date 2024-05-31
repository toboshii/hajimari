<script lang="ts">
    import { appQuery } from "$lib/stores.js";
    import Icon from "@iconify/svelte";

    export let providers = [
        {
            name: "Google",
            token: "g",
            icon: "simple-icons:google",
            searchUrl: "https://www.google.com/search?q={query}",
            url: "https://www.google.com",
        },
        {
            name: "DuckDuckGo",
            token: "d",
            icon: "simple-icons:duckduckgo",
            searchUrl: "https://duckduckgo.com/?q={query}",
            url: "https://duckduckgo.com",
        },
        {
            name: "IMDB",
            token: "i",
            icon: "simple-icons:imdb",
            searchUrl: "https://www.imdb.com/find?q={query}",
            url: "https://www.imdb.com",
        },
        {
            name: "Reddit",
            token: "r",
            icon: "simple-icons:reddit",
            searchUrl: "https://www.reddit.com/search?q={query}",
            url: "https://www.reddit.com",
        },
        {
            name: "YouTube",
            token: "y",
            icon: "simple-icons:youtube",
            searchUrl: "https://www.youtube.com/results?search_query={query}",
            url: "https://www.youtube.com",
        },
        {
            name: "Spotify",
            token: "s",
            icon: "simple-icons:spotify",
            searchUrl: "https://open.spotify.com/search/{query}",
            url: "https://open.spotify.com",
        },
        {
            name: "ABC",
            token: "a",
            icon: "mdi:test-tube",
            url: "https://example.com",
        }
    ];

    export let defaultProvider = "Google";

    let query = "";
    let defaultProviderRecord = providers.find(
        (provider) => provider.name === defaultProvider
    );
    let icon = defaultProviderRecord?.icon;

    $: {
        let matches = query.match(/\/(.*)/);
        if (matches) {
            $appQuery = matches[1];
            icon = "mdi:apps";
        } else {
            $appQuery = "";
            icon = defaultProviderRecord?.icon;
        }
    }

    $: {
        let matches = query.match(/@(\w+)\s?(.*)/);
        if (matches) {
            let token = matches[1];
            let provider = providers.find(
                (provider) => provider.token === token
            );
            if (provider) {
                icon = provider.icon;
            }
        }
    }

    const handleSubmit = () => {
        query = query.replaceAll("+", "%2B");

        let matches = query.match(/@(\w+)\s?(.*)/);
        if (matches) {
            let token = matches[1];
            let queryText = matches[2];

            let provider = providers.find(
                (provider) => provider.token === token
            );

            if (provider?.searchUrl && queryText) {
                window.location.assign(
                    provider.searchUrl.replaceAll("{query}", queryText)
                );
            } else if (provider) {
                window.location.assign(provider.url);
            }
        } else if (validURL(query)) {
            if (containsProtocol(query)) {
                window.location.assign(query);
            } else {
                window.location.assign("https://" + query);
            }
        } else {
            let provider = defaultProviderRecord;
            if (provider?.searchUrl) {
                window.location.assign(
                    provider.searchUrl.replaceAll("{query}", query)
                );
            }
        }
    };

    // Source: https://stackoverflow.com/questions/5717093/check-if-a-javascript-string-is-a-url
    function validURL(str: string) {
        var pattern = new RegExp(
            "^(https?:\\/\\/)?" + // protocol
                "((([a-z\\d]([a-z\\d-]*[a-z\\d])*)\\.)+[a-z]{2,}|" + // domain name
                "((\\d{1,3}\\.){3}\\d{1,3}))" + // OR ip (v4) address
                "(\\:\\d+)?(\\/[-a-z\\d%_.~+]*)*" + // port and path
                "(\\?[;&a-z\\d%_.~+=-]*)?" + // query string
                "(\\#[-a-z\\d_]*)?$",
            "i"
        ); // fragment locator
        return !!pattern.test(str);
    }

    function containsProtocol(str: string) {
        var pattern = new RegExp("^(https?:\\/\\/){1}.*", "i");
        return !!pattern.test(str);
    }
</script>

<section id="search">
    <form on:submit|preventDefault={handleSubmit}>
        <!-- svelte-ignore a11y-autofocus -->
        <Icon {icon}/>
        <input
            bind:value={query}
            type="text"
            id="keywords"
            size="50"
            spellcheck="false"
            autofocus={true}
        />
    </form>
</section>

<style>
    #search :global(svg) {
        font-size: 1.5em;
        position: absolute;
        margin-top: 0.6em;
    }

    #search {
        margin-bottom: 3vh;
    }

    input {
        font-size: 1em;
        text-indent: 3em;
    }
</style>
