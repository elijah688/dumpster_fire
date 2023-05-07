<script>
    import { onMount } from "svelte";

    onMount(async () => {
        const response = await fetch("/", {
            method: "GET",
            headers: {
                "content-type": "application/json",
            },
        });

        const total = await response.json();
        console.log(JSON.stringify(total,null, 4))
    });
</script>

<h1>Welcome to SvelteKit</h1>
<p>
    Visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to read the documentation
</p>
