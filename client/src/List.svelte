<script lang="ts">
    import { token, userList, username, viewList } from "./stores";
    import type { List, Error } from "./types";
    import { getUserList } from "./functions";
    import ListItem from "./ListItem.svelte";

    const dateOptions = { year: "numeric", month: "2-digit", day: "2-digit" };

    function setLocalDate(date: Date) {
        return new Date(date).toLocaleDateString("de-DE", dateOptions);
    }

    // requests the userList
    async function getList() {
        try {
            const response = await getUserList($token);

            if (response.ok) {
                const list: List = await response.json();
                $userList = list;
            } else {
                const data: Error = await response.json();
                console.log(data.error);
            }
        } catch (error) {
            console.log(error);
        }
    }
    getList();
</script>

<section>
    <div class="container">
        <span class="close-button" on:click={() => ($viewList = false)}>✖</span>
        <h1>{`${$username}'s list`}</h1>
        {#if $userList != null}
            <span>
                {`Created: ${setLocalDate($userList.createdAt)}`} | {`Updated: ${setLocalDate($userList.updatedAt)}`}
            </span>
            <div>
                {#each $userList.searchterms as item}
                    <ListItem {item} />
                {/each}
            </div>
        {:else}
            <p>You have no list. Please create one.</p>
        {/if}
    </div>
</section>

<style>
    h1 {
        font-size: 2em;
        font-weight: bold;
        text-align: center;
    }

    p {
        text-align: center;
    }

    span {
        text-align: center;
        margin-bottom: 0.5rem;
    }

    section {
        position: fixed;
        top: 0;
        left: 0;
        z-index: 1;
        padding-top: 6rem;
        width: 100%;
        height: 100%;
        background-color: #000000e6;
    }

    .container {
        margin: auto;
        max-width: 25%;
        max-height: calc(100vh - 12rem);
        width: auto;
        display: flex;
        flex-direction: column;
        justify-content: center;
    }

    .container div {
        overflow-y: auto;
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
