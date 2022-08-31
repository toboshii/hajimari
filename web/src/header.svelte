<script>
    export let greeting
    export let date
</script>

<section>
    <h1>{greeting}</h1>
    <h2>{date}</h2>
</section>

<style>
    section {
        border-bottom: 0px solid var(--color-text-acc);
        z-index: 1;
    }
</style>