<script lang="ts">
    import { enlargeImgURL } from "./stores";
    import { onMount } from "svelte";

    export let placeholder: string, src: string, alt: string;
    let intersected: boolean,
        loaded = false,
        filePath: string,
        imgElement: HTMLImageElement,
        observer: IntersectionObserver;

    // Creates the Intersection Observer
    // When the element is in the window, the real picture will load
    // and the Observer ends the observation.
    function observerCallback(
        entries: IntersectionObserverEntry[],
        observer: IntersectionObserver
    ) {
        entries.forEach((entry) => {
            if (entry.isIntersecting) {
                intersected = true;
                observer.unobserve(imgElement);
            }
        });
    }

    // If the value of intersected changed, Svelte will update the filePath accordingly
    $: filePath = intersected ? src : placeholder;

    function handleLoad() {
        if (!loaded && filePath === src) {
            loaded = true;
        }
    }

    // When the page is loaded, the observer gets attached to the image.
    onMount(() => {
        observer = new IntersectionObserver(observerCallback);
        observer.observe(imgElement);

        return () => {
            observer.unobserve(imgElement);
        };
    });
</script>

<img
    src={filePath}
    {alt}
    on:click={() => ($enlargeImgURL = `${imgElement.src.slice(0, -6)}:large`)}
    on:load={handleLoad}
    bind:this={imgElement} />

<style>
    img {
        display: block;
        cursor: pointer;
        width: 100%;
        height: auto;
        margin-bottom: 0.125rem;
    }
</style>
