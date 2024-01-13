# README

kanapi is work in progress.

## About

kanapi is a simple & light-weight API client written in Go/Wails and svelte.

## Features

- [x] Simple API client to call REST APIs
- [x] Show response in JSON format and the response headers
- [x] Support for Dark mode!

On the todo list:

- [ ] Auth and header to be implemented
- [ ] Save API request history. Planned to be Git friendly
- [ ] Make keyboard friendly
- [ ] Add ability to run entire folders/collections

## Installation

Ensure you have `go` and `node` installed.
Clone this repo.

```bash
git clone https://github.com/prashanth1k/kanapi
cd kanapi/frontend
npm i
cd ..
```

## Live Development

To run in live development mode, run `wails dev` in the project directory. This will run a Vite development server that will provide very fast hot reload of your frontend changes. If you want to develop in a browser and have access to your Go methods, there is also a dev server that runs on http://localhost:34115. Connect to this in your browser, and you can call your Go code from devtools.

## Building

To build a redistributable, production mode package, use `wails build`.

## Why do this?

Yet another experiment with Golang + Wails + a webview that enables faster frontend development.
This was a joy to write.

Additionally -

1. The program is native to the OS - thanks to Go. Is lightweight - takes < 50 MB RAM to run
1. Fast to execute

## License

MIT
