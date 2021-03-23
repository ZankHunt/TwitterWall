<script lang="ts">
    import Picture from "./Picture.svelte";
    import type { Media } from "./types";
    export let medias: Media[];

    // Creates a svg placeholder for the image, because the real
    // images are only loaded, when they are in the window
    function setPlaceholder(width: number, height: number): string {
        const svg = encodeURIComponent(
            `<svg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 ${width} ${height}' height='${height}' width='${width}'><rect fill='#DDD' width='${width}' height='${height}' /></svg>`
        );
        return `data:image/svg+xml;charset=UTF-8,${svg}`;
    }
</script>

<div class="pic">
    {#each medias as media}
        <Picture
            src={media.type == "photo" ? media.url : media.previewImageURL}
            alt={media.type}
            placeholder={setPlaceholder(media.width, media.height)} />
    {/each}
</div>

<style>
    .pic {
        background-color: #202020;
    }
</style>
