<script>
    export let comment;
</script>

<div class="comment-card">
    <h3>{comment.Author}</h3>
    <p>{comment.Content}</p>
</div>

<style>
    .comment-card {
        padding-left: 5px;
        border-top: 1px solid rgb(214, 214, 214);
    }
</style>