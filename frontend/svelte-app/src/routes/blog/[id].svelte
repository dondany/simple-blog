<script context="module">
    export async function load(ctx) {
    const id = ctx.page.params.id;
    console.log(id);
    const res = await fetch(`http://localhost:8080/posts/${id}`);

    if (res.ok) return { props: { post: await res.json() } };
    return {
      status: res.status,
      error: new Error(),
    };
  }
</script>

<script>
  import PostFull from '../../lib/PostFull.svelte'
  export let post;
</script>

<main>
  <div class="post">
    <PostFull {post}/>
  </div>  
</main>

<style>
  .post {
      position: absolute;
    	left: 50%;
    	margin-right: -50%;
    	transform: translate(-50%, 0);
  }
</style>