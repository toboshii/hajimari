<script lang="ts">
    import { appQuery } from "$lib/stores.js";
    type ProviderKey = keyof typeof providers;

    export let providers = {
        g: {
            searchUrl: "https://www.google.com/search?q={query}",
            url: "https://www.google.com",
        },
        d: {
            searchUrl: "https://duckduckgo.com/?q={query}",
            url: "https://duckduckgo.com",
        },
        i: {
            searchUrl: "https://www.imdb.com/find?q={query}",
            url: "https://www.imdb.com",
        },
        r: {
            searchUrl: "https://www.reddit.com/search?q=",
            url: "https://www.reddit.com",
        },
        y: {
            searchUrl: "https://www.youtube.com/results?search_query={query}",
            url: "https://www.youtube.com",
        },
        s: {
            searchUrl: "https://open.spotify.com/search/{query}",
            url: "https://open.spotify.com",
        },
    };
    export let defaultProvider = "g" as ProviderKey;

    let query = "";

    $: {
        let matches = query.match(/\/(.*)/);
        if (matches) {
            $appQuery = matches[1];
        } else {
            $appQuery = "";
        }
    }

    const handleSubmit = () => {
        query = query.replaceAll("+", "%2B");

        let matches = query.match(/@(\w+)\s?(.*)/);
        if (matches) {
            let token = matches[1] as ProviderKey;
            let queryText = matches[2];

            let provider = providers[token];

            if (provider && queryText) {
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
            let provider = providers[defaultProvider];
            window.location.assign(
                provider.searchUrl.replaceAll("{query}", query)
            );
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
</style>
