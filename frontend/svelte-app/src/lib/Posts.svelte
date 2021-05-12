<script>
    import {onMount} from "svelte";
    import Post from "./PostCard.svelte";

    let posts;

    onMount(async() => {
        await fetch('http://localhost:8080/posts')
            .then(r => r.json())
            .then(data => {
                posts = data;
            })
            .catch(err =>  {
                console.log(err)
            });
    });
</script>

{#if posts}
    {#each posts as post}
        <ul>
            <li>
                <Post {post}/>
            </li>
        </ul>
    {/each}
{:else}
    <p class="loading">loading...</p>
{/if}

<style>
    .loading {
        opacity: 0;
        animation: 0.4s 0.8s forwards fade-in;
    }

    @keyframes fade-in {
        from {opacity: 0;}
        to {opacity: 1;}
    }

    li {
        list-style-type: none;
    }
</style>