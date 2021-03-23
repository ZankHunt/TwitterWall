<script lang="ts">
	import { tweets, notFoundVisible, loadingVisible, search } from "./stores";
	import PictureContainer from "./PicturesContainer.svelte";
	import { decodeText, sendAPIRequest } from "./functions";
	import type { Tweet, TweetObject } from "./types";
	import Conversation from "./Conversation.svelte";
	export let tweet: Tweet;

	const dateOptions = { year: "numeric", month: "2-digit", day: "2-digit" };
	const date = new Date(tweet.date).toLocaleDateString("de-DE", dateOptions);

	// requests the tweets from the user
	async function userTweets() {
		try {
			$loadingVisible = true;
			const data = await sendAPIRequest(`from%3A${tweet.authorID}`);
			$loadingVisible = false;
			if (data.length > 0) {
				// adds the request to the svelte store in case the user wants to store it in the list
				$search.name = tweet.username;
				$search.searchterm = `from%3A${tweet.authorID}`;
				const tweetObject: TweetObject = {
					title: tweet.username,
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
		} catch (error) {
			console.log(error);
		}
	}

	tweet.text = decodeText(tweet.text);
</script>

<article class="tweet">
	{#if tweet.medias != null}
		<PictureContainer medias={tweet.medias} />
	{:else if tweet.refTweet != null}
		<Conversation refTweet={tweet.refTweet[0]} />
	{/if}
	<div class="tweet-content">
		<img
			src={tweet.profilePicURL}
			alt="Profile Pic"
			on:click={userTweets} />
		<header class="tweet-username">
			<div class="tweet-name-text">{tweet.name}</div>
			<div class="tweet-username-text">{tweet.username}</div>
		</header>
		<div class="tweet-date">{date}</div>
		<section class="tweet-text">{tweet.text}</section>
		<div class="tweet-metrics">❤️ {tweet.likes}</div>
		<div class="tweet-metrics">🔁 {tweet.retweets}</div>
		<div class="tweet-metrics">💬 {tweet.quotes}</div>
	</div>
</article>

<style>
	.tweet {
		background: #353535;
		border-radius: 0.25rem;
		overflow: hidden;
		display: flex;
		flex-direction: column;
		margin-bottom: 1rem;
	}

	.tweet-content {
		margin: 0.5rem;
		flex-grow: 1;
		display: grid;
		gap: 0.25rem;
		grid-template-columns: repeat(6, 1fr);
		align-content: space-evenly;
	}

	.tweet:hover {
		transform: scale(1.05);
	}

	.tweet-username {
		grid-column-end: span 4;
	}

	.tweet-date {
		font-size: 0.75rem;
		opacity: 0.5;
	}

	.tweet-text {
		grid-column-end: span 6;
		padding-left: 0.5rem;
		white-space: pre-line;
	}

	.tweet-username-text {
		font-size: 0.75rem;
		opacity: 0.5;
	}

	.tweet-metrics {
		grid-column-end: span 2;
		text-align: center;
	}

	.tweet-metrics:hover {
		transform: scale(1.1);
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
