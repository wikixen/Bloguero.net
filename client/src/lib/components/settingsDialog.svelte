<script lang="ts">
    import { redirect } from "@sveltejs/kit";
	import ChangePw from "./changePW.svelte";
    import ChangeUs from "./changeUS.svelte";
    
    const onSubmit = () => {
    	redirect
    }
</script>

<section class="settings-menu">
		<h1>Account</h1>
		<p>Click Esc to return</p>
		<section class="choices">
			<ChangeUs />
			<ChangePw />
			<button on:submit={onSubmit} >Delete Account</button>
		</section>
</section>

<style>
	.settings-menu{
		display: flex;
		flex-direction: column;
		align-items: center;
		width: 15rem;
		height: 12rem;
	}
	h1{
		margin: 0;
	}
	p {
		font-size: .75rem;
	}
	.choices {
		display: flex;
		flex-direction: column;
		row-gap: 1rem;
	}
	button{
		border-radius: .5rem;
		font-size: 1rem;
		padding: .2rem .5rem;
		width: 10rem;
		border: 1px solid red;
		color: white;
		background-color: red;
	}
	button:hover{
		background-color: darkred;
	}
</style>