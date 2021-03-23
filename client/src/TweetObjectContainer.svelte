<script lang="ts">
    import TweetComponent from "./TweetComponent.svelte";
    import type { Tweet } from "./types";

    export let tweets: Tweet[];
</script>

<section class="tweets-container">
    {#each tweets as tweet}
        <TweetComponent {tweet} />
    {/each}
</section>

<style>
    .tweets-container {
        columns: auto 19rem;
        margin-top: 1rem;
    }
</style>
