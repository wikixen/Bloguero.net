<script lang="ts">
    import Dialog from "./dialog.svelte";
	
	let dialog: HTMLDialogElement;
</script>

<button  on:click={() => dialog.showModal()}>Change Username</button>
<Dialog bind:dialog>
	<p>Click Esc to return</p>
	<form action="" method="post">
		<label for="New Username">New Username</label>
		<input type="text" name="New Username" id="newUS"/>
		<input type="submit" on:submit={close} value="Submit" class="btn"/>
	</form>
</Dialog>

<style>
	button{
		border-radius: .5rem;
		font-size: 1rem;
		padding: .2rem .5rem;
		width: 10rem;
	}
	form {
		display: flex;
		flex-direction: column;
		row-gap: 1rem;
	}
	p {
		font-size: .75rem;
	}
	.btn {
		border-radius: .5rem;
		font-size: 1rem;
		padding: .2rem .5rem;
		width: 5rem;
	}
</style>