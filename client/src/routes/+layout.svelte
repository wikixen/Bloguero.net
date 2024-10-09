<script lang="ts">
	import SearchIcon from '$lib/icons/searchIcon.svelte'
	import AddIcon from '$lib/icons/addIcon.svelte'
    import Dialog from '$lib/components/dialog.svelte';
    import AuthForm from '$lib/components/authForm.svelte';
	import SettingsIcon from '$lib/icons/settingsIcon.svelte'
    import SettingsDialog from '$lib/components/settingsDialog.svelte';
	
	let dialog: HTMLDialogElement;
	let token = true;
</script>


<header>
	<div />
	<section>
		<a class="h-title" href="/">YouBlog</a>	
	</section>
	<div />
	<section class="h-buttons">
			<SearchIcon/>
		<AddIcon />
		{#if token !== true}
			<button class="login-btn" on:click={() => dialog.showModal()}>
				Log in
			</button>
			<Dialog bind:dialog >
				<AuthForm />
			</Dialog>
		{:else}
			<button class="login-btn" on:click={() => dialog.showModal()}>
				<SettingsIcon />
			</button>
			<Dialog bind:dialog >
				<SettingsDialog />
			</Dialog>
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
	.login-btn{
		font-size: 1.5rem;
		padding: .1rem 1.5rem;
		border-radius: .5rem;
	}
</style>