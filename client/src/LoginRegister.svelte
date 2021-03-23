<script lang="ts">
    import { userLogsIn, username, token } from "./stores";
    import type { Error } from "./types";
    let isRegistering = false,
        form: HTMLFormElement,
        password: HTMLInputElement,
        vPassword: HTMLInputElement,
        userForm: HTMLInputElement,
        h1: HTMLHeadingElement,
        user: string,
        iden: string;

    // logs in the user
    async function login() {
        try {
            const formData = new FormData(form);
            const login = {
                identity: formData.get("identity").toString(),
                password: formData.get("password").toString(),
            };
            const response = await fetch(
                `https://${window.location.hostname}/user/login`,
                {
                    headers: { "Content-Type": "application/json" },
                    method: "POST",
                    body: JSON.stringify(login),
                }
            );

            if (response.ok) {
                type Login = {
                    token: string;
                    username: string;
                };

                let data: Login = await response.json();
                $username = data.username;
                $token = data.token;
                $userLogsIn = false;
            } else {
                let data: Error = await response.json();
                h1.textContent = data.error;
                h1.style.color = "red";
                setTimeout(() => {
                    h1.textContent = "Login";
                    h1.style.color = "";
                }, 5000);
            }
        } catch (error) {
            console.log(error);
            isRegistering = false;
            $userLogsIn = false;
        }
    }

    // registers the user
    async function register() {
        try {
            const formData = new FormData(form);
            const register = {
                username: formData.get("username").toString(),
                email: formData.get("email").toString(),
                password: formData.get("password").toString(),
            };
            const response = await fetch(
                `https://${window.location.hostname}/user/register`,
                {
                    headers: { "Content-Type": "application/json" },
                    method: "POST",
                    body: JSON.stringify(register),
                }
            );

            if (response.ok) {
                user = register.username;
                isRegistering = false;
                $userLogsIn = true;
                iden = user;
            } else {
                let data: Error = await response.json();
                h1.textContent = data.error;
                h1.style.color = "red";
                setTimeout(() => {
                    h1.textContent = "Register";
                    h1.style.color = "";
                }, 5000);
            }
        } catch (error) {
            console.log(error);
            isRegistering = false;
            $userLogsIn = false;
        }
    }

    // validates the user password while they type
    function validatePassword() {
        if (password.value != vPassword.value) {
            vPassword.setCustomValidity("Password doesn't match!");
            vPassword.style.color = "red";
        } else {
            vPassword.setCustomValidity("");
            vPassword.style.color = "";
        }
    }
</script>

<section>
    {#if !isRegistering}
        <h1 bind:this={h1}>Login</h1>
        <form bind:this={form}>
            <input
                bind:value={iden}
                name="identity"
                type="text"
                placeholder="Username or E-Mail" />
            <input name="password" type="password" placeholder="Password" />
            <button on:click|preventDefault={login}>Login</button>
            <button on:click|preventDefault={() => (isRegistering = true)}>Register</button>
            <button on:click|preventDefault={() => ($userLogsIn = false)}>Close</button>
        </form>
    {:else}
        <h1 bind:this={h1}>Register</h1>
        <form bind:this={form}>
            <input
                bind:this={userForm}
                name="username"
                type="text"
                placeholder="Username" />
            <input name="email" type="text" placeholder="E-Mail" />
            <input
                bind:this={password}
                name="password"
                type="password"
                placeholder="Password" />
            <input
                bind:this={vPassword}
                name="verifyPassword"
                type="password"
                placeholder="Password"
                on:input={validatePassword} />
            <button on:click|preventDefault={register}>Register</button>
            <button on:click|preventDefault={() => (isRegistering = false)}>Return</button>
            <button on:click|preventDefault={() => ($userLogsIn = false)}>Close</button>
        </form>
    {/if}
</section>

<style>
    h1 {
        font-size: 2em;
        font-weight: bold;
        text-align: center;
    }

    section {
        position: fixed;
        top: 0;
        left: 0;
        z-index: 1;
        padding-top: calc(25vh);
        width: 100%;
        height: 100%;
        background-color: #000000e6;
    }

    form {
        margin: auto;
        display: block;
        max-width: 25%;
        max-height: calc(100vh - 12rem);
        width: auto;
        display: flex;
        flex-direction: column;
        justify-content: center;
    }

    input,
    button {
        padding: 0.5rem 1rem;
        margin-bottom: 1rem;
        font-size: 1rem;
        border-radius: 0.25rem;
        border: none;
        -webkit-appearance: none;
        float: left;
    }

    input:hover {
        opacity: 0.75;
    }

    button {
        background: #353535;
        color: #fff;
        cursor: pointer;
    }

    button:hover {
        opacity: 0.75;
    }

    button:active {
        opacity: 1;
    }
</style>
