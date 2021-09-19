# yareio-wasm-gluer

This is a faster version of [L0laapk3's yare.io-wasm](https://github.com/L0laapk3/yare.io-wasm)

## Installation

- Make sure you have go installed
- Run `go install github.com/swz-gh/yareio-wasm-gluer@latest`
- Make sure `~/go/bin` (`%USERPROFILE%/go/bin` on windows) is in your PATH

## Usage

`yareio-wasm-gluer [FILE]` and it will output the result to stdout.

To write to a file do `yareio-wasm-gluer [FILE] > [OUTPUTFILE]`

## Building

- Clone this repo
- Make sure you have go installed
- Run `go build` in this folder and you should have a yareio-wasm-gluer binary
