<script lang="ts">
	type btnString = "Log In" | "Sign Up";
	
	let btnVal:btnString = "Log In";
	let otherLink: btnString = "Sign Up";
	
	const onClick = () => {
		btnVal = (btnVal === "Log In") ? "Sign Up": "Log In";
		otherLink = (btnVal === "Log In") ? "Sign Up": "Log In";
	}
</script>


<form action="">
	<label for="">Username</label>
	<input type="text" class="textInput"/>
	<label for="">Password</label>
	<input type="text" class="textInput"/>
	<input type="submit" value={btnVal} class="submit-btn"/>
	<button on:click={onClick}>{otherLink}</button>
</form>

<style>
	form {
		display: flex;
		flex-direction: column;
		row-gap: .5rem;
		padding: 1rem;
		align-items: start;
	}
	.textInput {
		height: 1.3rem;
		min-width: 15rem;
		border-radius: .3rem;
	}
	.submit-btn {
		padding: .1rem 1rem;
		border-radius: .5rem;
	}
	button {
		background: none;
		border: none;
		cursor: pointer;
		
	}
	button:hover {
		text-decoration: underline;
		color: grey;
	}
</style>