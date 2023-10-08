# README

## About

A sample todo demo app built using official Wails Vue template.

To know more, see the post on [building a todo app using Wails & Vue on techformist.com](https://techformist.com/go-webview-experiments)

## Live Development

You need to have Go, Wails and Node installed on your system to run this project.
Just clone this repo and run `wails dev` in the project directory.

`wails dev` in the project directory will -
- run a Vite development server that will provide very fast hot reload of your frontend changes.
- a golang backend

If you want to develop in a browser and have access to your Go methods, there is also a dev server that runs on http://localhost:34115. Connect to this in your browser, and you can call your Go code from devtools.

## Building

To build a redistributable, production mode package, use `wails build`.
