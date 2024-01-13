import { writable } from "svelte/store";

export const req = writable({ method: "GET", url: "", body: {}, headers: {} });

function reset() {
  req.set({ method: "GET", url: "", body: {}, headers: {} });
}

function setMessage() {
  //   req.update(body => return ({...req, body });
}
