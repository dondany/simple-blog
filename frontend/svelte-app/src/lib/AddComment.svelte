
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
    <form on:submit={handleSubmit}>
        <textarea class="input" bind:value={content} placeholder="Add comment..."></textarea>
        <input class="input" bind:value={author} placeholder="Name...">
        <input type="submit" value="Add Comment">
    </form>
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

    input[type=submit] {
        background-color: #a5a5a5;
        border: none;
        border-radius: 4px;
        margin-top: 4px;
    }
    .input:focus, 
    input[type=submit]:hover {
        box-shadow: 0 2px 6px rgba(121, 137, 148, .55);
    }
    .comment-form {
        margin-top: 10px;
    }

</style>