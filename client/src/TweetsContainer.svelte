<script lang="ts">
    import TweetArrayContainer from "./TweetObjectContainer.svelte";
    import { tweets } from "./stores";
</script>

<section class="tweets-container">
    {#each $tweets as tweetsObject}
        <h2>{tweetsObject.title}</h2>
        <TweetArrayContainer tweets={tweetsObject.tweets} />
    {/each}
</section>

<style>
    h2 {
        text-align: center;
        font-weight: bold;
    }
</style>
