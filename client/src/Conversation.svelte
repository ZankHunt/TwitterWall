<script lang="ts">
	import { tweets, notFoundVisible, loadingVisible, search } from "./stores";
	import { decodeText, sendAPIRequest } from "./functions";
	import type { RefTweet, TweetObject } from "./types";
	export let refTweet: RefTweet;

	const dateOptions = { year: "numeric", month: "2-digit", day: "2-digit" };
	const date = new Date(refTweet.date).toLocaleDateString(
		"de-DE",
		dateOptions
	);

	// requests the tweets from the refered user
	async function userTweets() {
		$loadingVisible = true;
		const data = await sendAPIRequest(`from%3A${refTweet.authorID}`);
		$loadingVisible = false;

		if (data.length > 0) {
			// adds the request to the svelte store in case the user wants to store it in the list
			$search.name = refTweet.username;
			$search.searchterm = `from%3A${refTweet.authorID}`;
			const tweetObject: TweetObject = {
				title: refTweet.username,
				tweets: data,
			};
			const storage = $tweets;
			$tweets = [tweetObject, ...storage];
		} else {
			$notFoundVisible = true;
			setTimeout(() => {
				$notFoundVisible = false;
			}, 5000);
		}
	}

	refTweet.text = decodeText(refTweet.text);
</script>

<section class="conv">
	<div class="conv-content">
		<img
			src={refTweet.profilePicURL}
			alt="Profile Pic"
			on:click={userTweets} />
		<header class="conv-username">
			<div class="conv-name-text">{refTweet.name}</div>
			<div class="conv-username-text">{refTweet.username}</div>
		</header>
		<div class="conv-date">{date}</div>
		<section class="conv-text">{refTweet.text}</section>
	</div>
	<hr />
</section>

<style>
	.conv {
		font-size: 0.75rem;
	}

	.conv-content {
		margin: 0.5rem;
		display: grid;
		gap: 0.25rem;
		grid-template-columns: repeat(6, 1fr);
		align-content: space-evenly;
	}

	.conv-username {
		grid-column-end: span 4;
	}

	.conv-date {
		font-size: 0.5rem;
		opacity: 0.5;
	}

	.conv-text {
		grid-column-end: span 6;
		padding-left: 0.5rem;
		white-space: pre-line;
	}

	.conv-username-text {
		font-size: 0.75rem;
		opacity: 0.5;
	}

	img {
		display: block;
		border-radius: 50%;
	}

	img:hover {
		transform: scale(1.1);
		cursor: pointer;
	}
</style>
