<script lang="ts">
    import BlogCard from "$lib/components/home/blogCard.svelte";
    import sampleData from '$lib/Data/sampleData.json';
    import type { BlogType } from "$lib/interfaces";

    const data: BlogType = sampleData[0];
</script>

<div class="grid-divider">
    <div />
    <section>
        <BlogCard {...data}/>
    </section>
    <div />
</div>

<style>
    .grid-divider {
        display: grid;
        grid-template-columns: 1fr 2fr 1fr;
    }
</style>