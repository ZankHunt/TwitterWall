import { writable } from "svelte/store";
import type { List, Search } from "./types";

let list: List;
const searchEntry: Search = {
    name: "",
    searchterm: "",
};

export const tweets = writable([]);
export const notFoundVisible = writable(false);
export const loadingVisible = writable(false);
export const enlargeImgURL = writable("");
export const userLogsIn = writable(false);
export const viewList = writable(false);
export const userList = writable(list);
export const token = writable("");
export const username = writable("");
export const search = writable(searchEntry);