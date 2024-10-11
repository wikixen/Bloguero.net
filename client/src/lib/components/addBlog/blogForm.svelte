<script lang="ts">
	import type { Token } from "$lib/interfaces";
	let blogText: string;
	
	function idInput(e: Event) {
		e.preventDefault()
		blogText = (e.target! as HTMLTextAreaElement).value;
	}
	
	// console.log(blogText)
	
	const createBlog = () => {
		
	}
	
	
</script>

<section>
	<form action="/" method="post">
		<label for="Create Blog" class="b-title">Create Blog</label>
		<textarea name="Create Blog" class="b-text" on:input={idInput} />
		<input type="submit" value="Create" class="b-btn" on:submit={createBlog}>
	</form>
</section>


<style>
	section {
		display: flex;
		flex-direction: column;
		align-items: center;
		row-gap: 1rem;
		height: auto;
	}
	.b-title {
		font-size: 1.5rem;
		margin: 0;
	}
	form {
		display: flex;
		flex-direction: column;
		row-gap: 1rem;
		width: 20rem;
	}
	.b-text{
		min-height: 10rem;
		width: 20rem;
		resize: none;
	}	
	.b-btn{
		font-size: 1rem;
		padding: .1rem .5rem;
		border-radius: .5rem;
	}
</style>