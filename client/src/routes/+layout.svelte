<script lang="ts">
	import SearchIcon from '$lib/icons/searchIcon.svelte'
	import AddIcon from '$lib/icons/addIcon.svelte'
	import SettingsIcon from '$lib/icons/settingsIcon.svelte'
    import LoginBtn from '$lib/components/auth/loginBtn.svelte';
	
	let dialog: HTMLDialogElement;
	let token = true;
</script>


<header>
	<div />
	<section>
		<a class="h-title" href="/">Bloguero</a>	
	</section>
	<div />
	<section class="h-buttons">
			<SearchIcon/>
		{#if token !== true}
			<LoginBtn />
		{:else}
			<AddIcon />
			<SettingsIcon />
		{/if}
	</section>
	<div />
</header>

<slot  class="main-body"/>

<style>
	header{
		display: grid;
		grid-auto-flow: column;
		grid-auto-columns: 1rem auto 1fr auto 1rem;
		align-items: center;
		margin-bottom: 1rem;
		border-bottom: 2px solid black;
	}
	.h-title{
		font-size: 4rem;
		font-family: 'IBM Plex Mono', monospace;
		text-decoration: none;
		font-weight: bold;
		color: black;
	}
	.h-title:hover{
		color: #3d3d3d;
	}
	.h-buttons{
		display: flex;
		column-gap: .5rem;
	}
</style>