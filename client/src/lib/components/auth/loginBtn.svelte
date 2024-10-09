<script lang="ts">	
	import Dialog from '$lib/components/dialog.svelte';
    import AuthForm from './authForm.svelte';
	
	let dialog: HTMLDialogElement;
</script>

<button on:click={() => dialog.showModal()}>
	Log in
</button>
<Dialog bind:dialog >
	<AuthForm />
</Dialog>

<style>
	button{
		font-size: 1.5rem;
		padding: .1rem 1.5rem;
		border-radius: .5rem;
	}
</style>