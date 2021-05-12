<script>
    import Comment from './Comment.svelte'
    import AddComment from './AddComment.svelte'
    import {onMount} from 'svelte';
    
    export let post;
    let comments;

    onMount(async() => {
        await fetch(`http://localhost:8080/posts/${post.Id}/comments`)
            .then(r => r.json())
            .then(data => {
                console.log(data)
                comments = data;
            })
            .catch(err =>  {
                console.log(err)
            });
    });
</script>

<div class="post-full">
    <div class="post">
        <div class="title">
            <h1>{post.Title}</h1>
        </div>
        <div class="content">
            <p>{post.Content}</p>
        </div>
    </div>
    <div class="comment-section">
        <h3>Comments</h3>
        <AddComment postId={post.Id}/>
        {#if comments}
            {#each comments as comment}
            <ul>
                <li>
                    <Comment {comment}/>
                </li>
            </ul>
            {/each}
        {:else}
            <p class="no-comments">No comments yet.</p>
        {/if}
    </div>
</div>

<style>
    .post-full {
        width: 500px;
        margin: 1rem;
    }
    .post {
        border-bottom: 1px solid rgb(214, 214, 214);
    }
    h1 {
        font-size: 1.5em;
        margin: 0;
        display: block;
    }
    .title {
        border-bottom: 1px solid rgb(214, 214, 214);
    }
    h1 {
        font-size: 30px;
    }
    li {
        list-style-type: none;
    }
</style>