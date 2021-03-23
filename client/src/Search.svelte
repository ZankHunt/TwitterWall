<script lang="ts">
    import {
        tweets,
        notFoundVisible,
        loadingVisible,
        userList,
        token,
        username,
        search,
    } from "./stores";
    import { getUserList, sendAPIRequest, updateUserList } from "./functions";
    import type { Search, List, Error, TweetObject } from "./types";
    let searchValue = "",
        checkmarkButton: HTMLButtonElement;

    // Send the search request to the backend.
    function sendSearchRequest() {
        if (searchValue != "") {
            $loadingVisible = true;
            sendingRequest();
            // adds the request to the svelte store in case the user wants to store it in the list
            $search.name = searchValue;
            $search.searchterm = searchValue;
            searchValue = "";
        }
    }

    // Returned JSON gets added to the Svelte store
    async function sendingRequest() {
        try {
            const queryValue = encodeURIComponent(searchValue);
            const data = await sendAPIRequest(queryValue);
            $loadingVisible = false;

            if (data.length > 0) {
                const tweetObject: TweetObject = {
                    title: $search.name,
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

    // adds the item to the user list
    async function addItemToList() {
        try {
            const newSearch: Search = {
                name: $search.name,
                searchterm: $search.searchterm,
            };
            // if the userList is not in the svelte store, requests it
            if ($userList == null) {
                const response = await getUserList($token);
                if (response.ok) {
                    const list: List = await response.json();
                    $userList = list;
                } else if (response.status == 404) {
                    // if the user has no list (404), creates a new empty list
                    const response = await fetch(
                        `https://${window.location.hostname}/lists`,
                        {
                            headers: {
                                "Content-Type": "application/json",
                                Authorization: `Bearer ${$token}`,
                            },
                            method: "POST",
                            body: JSON.stringify({ searchterms: [] }),
                        }
                    );
                    if (response.ok) {
                        const data: List = await response.json();
                        $userList = data;
                    } else {
                        const data: Error = await response.json();
                        console.log(data.error);
                    }
                } else {
                    const data: Error = await response.json();
                    console.log(data.error);
                }
            }
            const searchterms = [...$userList.searchterms, newSearch];
            $userList.searchterms = searchterms;

            // adds / updates the userList
            const response = await updateUserList($token, $userList);

            if (response.ok) {
                const data: List = await response.json();
                $userList = data;
            } else {
                const data: Error = await response.json();
                console.log(data.error);
            }
        } catch (error) {
            console.log(error);
        }
    }
</script>

<section>
    <form>
        <button id="search" on:click|preventDefault={sendSearchRequest}>Search</button>
        <input bind:value={searchValue} type="text" placeholder="Search" />
        <button
            id="checkmark"
            bind:this={checkmarkButton}
            disabled={$username == "" || $search.name == ""}
            on:click|preventDefault={addItemToList}>✔</button>
    </form>
</section>

<style>
    form {
        display: flex;
        justify-content: center;
    }

    input,
    button {
        padding: 0.5rem 1rem;
        margin-bottom: 1rem;
        font-size: 1rem;
        border: none;
        -webkit-appearance: none;
        float: left;
    }

    input {
        width: 12.5rem;
    }

    input:hover {
        opacity: 0.75;
    }

    #checkmark:disabled {
        opacity: 0.5;
        cursor: not-allowed;
    }

    #checkmark {
        border-top-right-radius: 0.25rem;
        border-bottom-right-radius: 0.25rem;
        width: 3rem;
        cursor: pointer;
    }

    button {
        background: #353535;
        color: #fff;
    }

    #checkmark:hover {
        color: green;
    }

    #search {
        cursor: pointer;
        border-top-left-radius: 0.25rem;
        border-bottom-left-radius: 0.25rem;
    }

    #search:hover {
        opacity: 0.75;
    }

    #search:active {
        opacity: 1;
    }
</style>
