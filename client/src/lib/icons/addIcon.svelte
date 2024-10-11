<script lang="ts">
	import BlogForm from '$lib/components/addBlog/blogForm.svelte';
	import Dialog from '$lib/components/dialog.svelte';
	
	let dialog: HTMLDialogElement;
</script>


<button class="h-add" on:click={() => dialog.showModal()}>
	<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="black" width="2rem">
	  <path d="M10.75 4.75a.75.75 0 0 0-1.5 0v4.5h-4.5a.75.75 0 0 0 0 1.5h4.5v4.5a.75.75 0 0 0 1.5 0v-4.5h4.5a.75.75 0 0 0 0-1.5h-4.5v-4.5Z" />
	</svg>
</button>

<Dialog bind:dialog>
	<BlogForm />
</Dialog>


<style>
	.h-add {
		font-size: 1.5rem;
		padding: .1rem .5rem;
		border-radius: .5rem;
	}
	.h-add:hover {
		fill: #3d3d3d;
		cursor: pointer;
	}
</style>