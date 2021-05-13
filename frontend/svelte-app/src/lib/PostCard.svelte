<script>
    import LikesCounter from './LikesCounter.svelte'
    import CommentsCounter from './CommentsCounter.svelte'

    export let post;
</script>

<a class="post-card" href={`/blog/${post.Id}`} >
    <div class="title">
        <h1>{post.Title}</h1>
    </div>
    <div class="content">
        <p>{post.Content.substr(0,100) + "..."}</p>
    </div>
    
    <span class="likes"><LikesCounter readOnly="true" postId={post.Id} likes={post.Likes}/></span>
    <span class="comments"><CommentsCounter comments={post.CommentsCount}/></span>
</a>

<style>
    .post-card {
        border-radius: 5px;
        display: inline-block;
        height: 120px;
        width: 500px;
        
        box-shadow: 0 1px 3px rgba(0,0,0,0.30);
        padding: 0.5rem;        
    }
    .post-card:hover {
        box-shadow: 0 14px 28px rgba(0,0,0,0.25), 0 10px 10px rgba(0,0,0,0.22);
    }
    h1 {
        font-size: 1.5em;
        margin: 0;
        display: block;
    }
    .title {
        border-bottom: 1px solid rgb(214, 214, 214);
    }
    a.post-card{
        color: inherit;
        text-decoration: none;
    }
    .comments {
        float: right;
        margin-left: 2px;
        margin-right: 2px;
    }
    .likes {
        float: right;
        margin-left: 2px;
        margin-right: 2px;
    }
</style>