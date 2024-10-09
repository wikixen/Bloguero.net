<button class="h-add">
	<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="black" width="2rem">
	  <path d="M10.75 4.75a.75.75 0 0 0-1.5 0v4.5h-4.5a.75.75 0 0 0 0 1.5h4.5v4.5a.75.75 0 0 0 1.5 0v-4.5h4.5a.75.75 0 0 0 0-1.5h-4.5v-4.5Z" />
	</svg>
</button>

<dialog>
	<form action="REST API HERE">
		<textarea id="b-text" name="blog-text" class="textform text-desc" />
		<input type="submit" value="Post" class="create-btn">
	</form>
</dialog>


<style>
	.h-add {
		background: none;
		border: none;
		text-shadow: none;
	}
	.h-add:hover {
		fill: #3d3d3d;
		cursor: pointer;
	}
	dialog {
		max-width: auto;
		width: 35rem;
		height: auto;
		padding-block: 1.5rem;
	}
	dialog form {
		display: flex;
		flex-direction: column;
		row-gap: .3rem;
	}
	.create-btn{
		width: 5rem;
	}
	.text-desc{
		height: 5rem;
		resize: none;
		max-height: 5rem;
	}
</style>