<script>
    export let group;
</script>
{#each group.apps as app}
<div class="apps_item">
    <div class="apps_icon">
        <span class="iconify icon" data-icon="mdi:{app.icon}"></span>
    </div>
    <div class="apps_text">
        <a href="{app.url}">{app.name}</a>
        <span class="app_address">{app.url}</span>
    </div>
</div>
{/each}

<style>
    .apps_icon {
        height: 64px;
        margin-right: 1em;
        padding-top: 15px;
    }

    .apps_icon span {
        font-size: 2.5em;
        line-height: 3rem;
    }

    .apps_item {
        display: flex;
        flex-direction: row;
        flex-wrap: wrap;
        height: 64px;
        margin: 0;
    }

    .apps_text {
        display: flex;
        flex-direction: column;
        justify-content: center;
        flex: 1;
        overflow: hidden;
    }

    .apps_text a {
        font-size: 1em;
        font-weight: 500;
        text-transform: uppercase;
    }

    .apps_text span {
        color: var(--color-text-acc);
        font-size: 0.8em;
        text-transform: uppercase;
    }

    @media screen and (max-width: 667px) {
        .apps_icon {
            height: 64px;
            margin-right: 0.8em;
            padding-top: 14px;
        }

        .apps_icon span {
            font-size: 2em;
            line-height: 2.5rem;
        }
    }

    @media only screen and (max-width: 400px) {
        .app_address {
            display: none;
        }
    }
</style>