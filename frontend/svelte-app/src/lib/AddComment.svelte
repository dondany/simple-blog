
<script>
    export let postId;
    let author = '';
    let content = '';

    const handleSubmit = async () => {
		const res = await fetch(`http://localhost:8080/posts/${postId}/comments`, {
            method: 'POST',
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                postId,
                author,
                content
            })
        })
        .catch(err =>  {
                console.log(err)
        });
	}
</script>

<div class="comment-form">
    <div class="form-group">
    <form on:submit={handleSubmit}>
        <div>
            <textarea class="form-control" id="commentArea" bind:value={content} placeholder="Add comment..."></textarea>
        </div>
        <div>
            <input class="form-control" bind:value={author} placeholder="Name...">
        </div>
        
        <input type="submit" class="btn btn-primary" value="Add Comment">
    </form>
    </div>
</div>

<style>
    .input { 
        width: 100%;
        border: none;
        border-radius: 4px;
        font-family: inherit;
        font-size: 14px;
        padding: 5px 10px;
        outline: none;
        width: 100%;
        box-shadow: 0 2px 6px rgb(0 0 0 / 55%) !important;
        margin-top: 2px;
    }

    textarea.input {
        height: 80px;
    }

    .comment-form {
        margin-top: 10px;
    }

</style>