<script>
    import thumbUp from './icons/hand-thumbs-up.svg'
    import thumbUpFill from './icons/hand-thumbs-up-fill.svg'

    export let readOnly = false;
    export let postId;
    export let likes;
    let isLiked = false;

    const toggleLike = async () => {
        isLiked = !isLiked;

        const endpoint = isLiked ? "addLike" : "deleteLike";
        likes = isLiked ? likes + 1 : likes - 1;
        await fetch(`http://localhost:8080/posts/${postId}/${endpoint}`, {
            method: 'PUT',
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json'
            },
        })
        .catch(err =>  {
                console.log(err)
        });
    }
</script>
{#if !readOnly}
<button on:click={toggleLike} class="like-btn">
    {#if isLiked}
    <i class="bi bi-hand-thumbs-up liked"/> {likes}
    {:else}
    <i class="bi bi-hand-thumbs-up"/> {likes}
    {/if}
</button>
{:else}
    <i class="bi bi-hand-thumbs-up"/> {likes}
{/if}


<style>
    .like-btn {
        background-color: white;
        border: none;
    }
    .like-btn:hover {
        color: rgb(73, 73, 247);
        cursor: pointer;
    }
    .bi-hand-thumbs-up.liked{
        color: #0d6efd;
    }
</style>