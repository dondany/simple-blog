<script context="module">
	import Posts from "../../lib/Posts.svelte"

	let posts;
</script>

<dev class="posts">
	<Posts></Posts>
</dev>


<style>
	.posts {
		position: absolute;

    	left: 50%;
    	margin-right: -50%;
    	transform: translate(-50%, 0);
	}
</style>