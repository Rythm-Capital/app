# README

### Requirements

`go 1.18`

[Wails CLI](https://github.com/wailsapp/wails)

## About

This is the official Wails React template.

You can configure the project by editing `wails.json`. More information about the project settings can be found
[here.](https://wails.io/docs/reference/project-config)

## Live Development

To run in live development mode, run `wails dev` in the project directory. This will run a Vite development
server that will provide very fast hot reload of your frontend changes. If you want to develop in a browser
and have access to your Go methods, there is also a dev server that runs on [local.](http://localhost:34115) Connect
to this in your browser, and you can call your Go code from devtools.

## Building

To build a redistributable, production mode package, use `wails build`.


## Package Structure

- __main__ desktop.rythm
  - __rythm__: desktop.rythm/core
    - Core types and methods for the indicator and signals
  - __data__: desktop.rythm/data
    - Adaptor for realtime fetching from OANDA 
  - __testData__: desktop.rythm/testData
    - Sample data for testing

## Testing

__run tests__: 
  
> `go test ./core`

__verbose__: 

> `go test ./core -v`

__coverage report__: 

> `go test ./core -cover`



GBPUSD, ASK:160.631, BID: 160.640, BARSIZE: 10, DIRECTION: UP
GBPUSD, ASK:160.631, BID: 160.640, BARSIZE: 10, DIRECTION: UP
GBPUSD, ASK:160.631, BID: 160.640, BARSIZE: 10, DIRECTION: UP
GBPUSD, ASK:160.631, BID: 160.640, BARSIZE: 10, DIRECTION: UP
GBPUSD, ASK:160.631, BID: 160.640, BARSIZE: 10, DIRECTION: UP
GBPUSD, ASK:160.631, BID: 160.640, BARSIZE: 10, DIRECTION: UP
GBPUSD, ASK:160.631, BID: 160.640, BARSIZE: 10, DIRECTION: UP
GBPUSD, ASK:160.631, BID: 160.640, BARSIZE: 10, DIRECTION: UP
GBPUSD, ASK:160.631, BID: 160.640, BARSIZE: 10, DIRECTION: UP
GBPUSD, ASK:160.631, BID: 160.640, BARSIZE: 10, DIRECTION: UP
GBPUSD, ASK:160.631, BID: 160.640, BARSIZE: 10, DIRECTION: UP
GBPUSD, ASK:160.620, BID: 160.625, BARSIZE: 10, DIRECTION: DN
GBPUSD, ASK:160.620, BID: 160.625, BARSIZE: 10, DIRECTION: DN
GBPUSD, ASK:160.620, BID: 160.625, BARSIZE: 10, DIRECTION: DN
GBPUSD, ASK:160.620, BID: 160.625, BARSIZE: 10, DIRECTION: DN
GBPUSD, ASK:160.620, BID: 160.625, BARSIZE: 10, DIRECTION: DN
GBPUSD, ASK:160.620, BID: 160.625, BARSIZE: 10, DIRECTION: DN
GBPUSD, ASK:160.620, BID: 160.625, BARSIZE: 10, DIRECTION: DN
GBPUSD, ASK:160.620, BID: 160.625, BARSIZE: 10, DIRECTION: DN
GBPUSD, ASK:160.620, BID: 160.625, BARSIZE: 10, DIRECTION: DN
GBPUSD, ASK:160.620, BID: 160.625, BARSIZE: 10, DIRECTION: DN
GBPUSD, ASK:160.631, BID: 160.640, BARSIZE: 10, DIRECTION: UP
GBPUSD, ASK:160.631, BID: 160.640, BARSIZE: 10, DIRECTION: UP
GBPUSD, ASK:160.631, BID: 160.640, BARSIZE: 10, DIRECTION: UP
GBPUSD, ASK:160.631, BID: 160.640, BARSIZE: 10, DIRECTION: UP
GBPUSD, ASK:160.631, BID: 160.640, BARSIZE: 10, DIRECTION: UP


Simple Pattern
BUY:100U, 50D
LOTSIZE: 0.01
STOPLOSS: 100D, 50D
TARGET: 100U, 50U
ACCEPTABLE SPREAD: >1.5pip


Adv Pattern
100,90,80,70,60,50,40,30,20,10
U,D,D,D,D,D,D,U,D,D,D,U,D,U