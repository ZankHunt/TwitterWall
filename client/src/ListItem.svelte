<script lang="ts">
    import {
        loadingVisible,
        notFoundVisible,
        search,
        token,
        tweets,
        userList,
        viewList,
    } from "./stores";
    import type { List, Error, Search, TweetObject } from "./types";
    import { sendAPIRequest, updateUserList } from "./functions";
    export let item: Search;

    // removes the selected Item from the list
    async function removeItem() {
        try {
            const index = $userList.searchterms.indexOf(item);
            $userList.searchterms.splice(index, 1);
            $userList.searchterms = [...$userList.searchterms];

            // if the list is now empty, sends a DELETE request
            if ($userList.searchterms.length == 0) {
                const response = await fetch(
                    `https://${window.location.hostname}/lists`,
                    {
                        headers: { Authorization: `Bearer ${$token}` },
                        method: "DELETE",
                    }
                );

                if (response.ok) {
                    $userList = null;
                } else {
                    const data: Error = await response.json();
                    console.log(data.error);
                }
            } else {
                const response = await updateUserList($token, $userList);

                if (response.ok) {
                    const data: List = await response.json();
                    $userList = data;
                } else {
                    const data: Error = await response.json();
                    console.log(data.error);
                }
            }
        } catch (error) {
            console.log(error);
        }
    }

    // send a twitter search request with the selected item
    async function searchItem() {
        try {
            $viewList = false;
            $loadingVisible = true;
            const data = await sendAPIRequest(`${item.searchterm}`);
            $loadingVisible = false;
            if (data.length > 0) {
                $search.name = item.name;
                $search.searchterm = item.searchterm;
                const tweetObject: TweetObject = {
                    title: item.name,
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
</script>

<div class="button-container">
    <button class="button-left" on:click={searchItem}>{item.name}</button>
    <button class="button-right" on:click={removeItem}>✖</button>
</div>

<style>
    button {
        padding: 0.5rem 1rem;
        margin-bottom: 1rem;
        font-size: 1rem;
        border: none;
        -webkit-appearance: none;
        cursor: pointer;
        float: left;
        background: #353535;
        color: #fff;
        text-align: left;
    }

    button:hover {
        opacity: 0.75;
    }

    button:active {
        opacity: 1;
    }

    .button-container {
        display: flex;
        justify-content: center;
    }

    .button-left {
        border-top-left-radius: 0.25rem;
        border-bottom-left-radius: 0.25rem;
        width: calc(100% - 3rem);
    }

    .button-right {
        border-top-right-radius: 0.25rem;
        border-bottom-right-radius: 0.25rem;
        width: 3rem;
    }

    .button-right:hover {
        color: red;
    }
</style>
