<script lang="ts">
	import DislikeIcon from "$lib/icons/dislikeIcon.svelte";
	import LikeIcon from "$lib/icons/likeIcon.svelte";
	
	export let author: string, content: string, likes: number, dislikes: number;
	
	let contentLen = 300;
</script>

<article class="b-card">	
	<a href="/{author}"class="b-author">{author}</a>
	{#if content.length > contentLen}
		<p class="b-content">{content.slice(0,300)+"..."}
		<button class="readmore" on:click={()=>contentLen = content.length}>Read More</button>	
		</p>
	{:else}
		<p class="b-content">{content}</p>
	{/if}
	<div class="b-btn-grp">
		<LikeIcon likes={likes}  />
		<DislikeIcon dislikes={dislikes}/>
	</div>
</article>

<style>
	.b-card{
		background-color: #d9d9d9;
		display: flex;
		flex-direction: column;
		padding: .5rem 2rem;
	}
	.b-title, .b-author, .b-content {
		margin:0;
	}
	.b-title {
		font-size: 1.5rem;
		font-weight: bold;
		margin-top: 1rem;
	}
	.b-author {
		margin-bottom: .75rem;
		font-size: .9rem;
		text-decoration: none;
	}
	.b-content {
		margin-bottom: 1rem;
	}
	.readmore {
		cursor: pointer;
		background: none;
		border: none;
		color: #525151;
	}
	.readmore:hover {
		text-decoration: underline;
	}
	.b-btn-grp {
		display: flex;
		column-gap: .5rem;
	}
</style>