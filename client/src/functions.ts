import type { List } from "./types";

// decodes the twitter text for example <, >, etc. with textarea to block xss
export function decodeText(input: string): string {
    const textarea = document.createElement("textarea");
    textarea.innerHTML = input;
    return textarea.value;
}

// sends twitter api request
export async function sendAPIRequest(input: string): Promise<any> {
    try {
        const response = await fetch(`https://${window.location.hostname}/search?query=${input}`)
        if (!response.ok) {
            throw new Error(`HTTP error! Status: ${response.status}`);
        } else {
            return response.json();
        }
    } catch (error) {
        console.log(error);
    }
}

// requests the user list from the server
export async function getUserList(token: string) {
    try {
        return await fetch(
            `https://${window.location.hostname}/lists`,
            {
                headers: { Authorization: `Bearer ${token}` },
                method: "GET",
            }
        );
    } catch (error) {
        console.log(error);
    }
}

// updates the user list on the server
export async function updateUserList(token: string, userList: List) {
    try {
        return await fetch(
            `https://${window.location.hostname}/lists`,
            {
                headers: {
                    "Content-Type": "application/json",
                    Authorization: `Bearer ${token}`,
                },
                method: "PUT",
                body: JSON.stringify(userList),
            }
        );
    } catch (error) {
        console.log(error);
    }
}