<script lang="ts">
	import {
		loadingVisible,
		notFoundVisible,
		enlargeImgURL,
		userLogsIn,
		viewList,
username,
	} from "./stores";
	import TweetsContainer from "./TweetsContainer.svelte";
	import LoginRegister from "./LoginRegister.svelte";
	import { fade, scale } from "svelte/transition";
	import Search from "./Search.svelte";
	import List from "./List.svelte";
</script>

{#if $enlargeImgURL != ""}
	<div id="enlargedPic" class="enlarge-container">
		<span class="close-button" on:click={() => ($enlargeImgURL = "")}>✖</span>
		<img
			class="enlarge-pic"
			src={$enlargeImgURL}
			alt="enlarged"
			in:scale={{ start: 0, duration: 1000 }} />
	</div>
{/if}
{#if $userLogsIn}
	<LoginRegister />
{:else if $viewList}
	<List />
{/if}
<main>
	{#if $username != ""}
		<button on:click={() => ($viewList = true)}>Your List</button>
	{:else}
		<button on:click={() => ($userLogsIn = true)}>Login or Register</button>
	{/if}
	<section>
		<h1>
			<img src="./img/icon.svg" alt="Icon" />
			Twitter Wall
			<img src="./img/icon.svg" alt="Icon" />
		</h1>
		<p>Display and Search your favourite tweets!</p>
		<Search />
	</section>
	{#if $loadingVisible}
		<img id="loading" src="./img/loading.svg" alt="loading" />
	{/if}
	{#if $notFoundVisible}
		<p id="noResult" transition:fade>
			No result from Twitter for that search term!
		</p>
	{/if}
	<TweetsContainer />
</main>

<style>
	main > button {
		position: absolute;
		top: 0;
		right: 0;
		padding: 0.5rem 1rem;
		font-size: 1rem;
		border: none;
		-webkit-appearance: none;
		float: left;
		background: #353535;
		color: #fff;
		cursor: pointer;
		border-radius: 0.25rem;
		margin-top: 1rem;
		margin-right: 1rem;
	}

	h1 {
		font-size: 2em;
		font-weight: bold;
		text-align: center;
	}

	h1 > img {
		padding-right: 0.5rem;
		padding-left: 0.5rem;
		width: 2.25rem;
		height: auto;
	}

	img:hover,
	button:hover {
		transform: scale(1.1);
	}

	p {
		text-align: center;
	}

	#noResult {
		margin: 0;
	}

	#loading {
		display: block;
		margin: auto;
	}

	.enlarge-container {
		position: fixed;
		z-index: 1;
		padding-top: 6rem;
		left: 0;
		top: 0;
		width: 100%;
		height: 100%;
		background-color: #000000e6;
	}

	.enlarge-pic {
		margin: auto;
		display: block;
		max-width: 95%;
		max-height: calc(100vh - 12rem);
		width: auto;
	}

	.close-button {
		position: absolute;
		top: 1rem;
		right: 1rem;
		color: #f0f0f0;
		font-size: 3rem;
		font-weight: bold;
	}

	.close-button:hover {
		cursor: pointer;
		color: #b1b1b1;
		text-decoration: none;
		transform: scale(1.1);
	}
</style>
