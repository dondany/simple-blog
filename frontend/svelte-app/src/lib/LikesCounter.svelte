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
        <img src={thumbUpFill} alt="like"/> {likes}
    {:else}
        <img src={thumbUp}/ alt="dislike"> {likes}
    {/if}
</button>
{:else}
    <img src={thumbUpFill} alt="like"/> {likes}
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
</style>