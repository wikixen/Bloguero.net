<script lang="ts">
	import '@fontsource/ibm-plex-mono/300.css';
	import '@fontsource/ibm-plex-mono/500.css';
	import '@fontsource/ibm-plex-mono/600.css';
	import sampleData from '$lib/Data/sampleData.json';
	import type { BlogType } from '$lib/interfaces';
	import BlogCard from '$lib/components/home/blogCard.svelte';

	
	const data: BlogType[] = sampleData;
</script>

<div class="grid-divider">
	<div />
	<div>
		<section class="blogs">
			{#each data as dat}
				<BlogCard 
					author={dat.author}
					content={dat.content}
					likes={dat.likes}
					dislikes={dat.dislikes}
				/>
			{/each}
		</section>
	</div>
	<div />
</div>

<style>
	.grid-divider {
		display: grid;
		grid-template-columns: 1fr 1.5fr 1fr;
	}
	.blogs{
		display: flex;
		flex-direction: column;
		row-gap: 1rem;
	}
</style>