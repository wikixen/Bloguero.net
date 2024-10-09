<script lang="ts">
    import Dialog from "./dialog.svelte";
	
	let dialog: HTMLDialogElement;
	
	
</script>

<button on:click={() => dialog.showModal()}>Change Password</button>
<Dialog bind:dialog>
	<p>Click Esc to return</p>
	<form action="" method="post">
		<label for="Old Password">Old Password</label>
		<input type="text" name="Old Password" id="oldPW">
		<label for="New Password">New Password</label>
		<input type="text" name="New Password" id="newPW">
		<label for="Re-enter Password">Re-enter New Password</label>
		<input type="text" name="Re-enter Password" id="rePW">	
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
	p {
		font-size: .75rem;
	}
	form {
		display: flex;
		flex-direction: column;
		row-gap: .3rem;
	}
	.btn {
		border-radius: .5rem;
		font-size: 1rem;
		padding: .2rem .5rem;
		width: 5rem;
	}
</style>