<script>
    import Comment from './Comment.svelte'
    import AddComment from './AddComment.svelte'
    import CommentsCounter from './CommentsCounter.svelte'
    import LikesCounter from './LikesCounter.svelte'

    import {onMount} from 'svelte';

    import commentIcon from './icons/chat-fill.svg'
    
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
        <div class="comments-likes">
            <span class="comments"><CommentsCounter comments={post.CommentsCount}/> Comments</span>
            <span class="likes"><LikesCounter postId={post.Id} likes={post.Likes}/></span>
        </div>

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
    .likes {
        float: right;
    }
    .comments {
        float: left;
        font-size: 18px;
        font-weight: bold;
    }
    .comments-likes {
        margin: 10px;
    }
</style>