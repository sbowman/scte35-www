# SCTE 35 WWW

This is a little demonstration of using the `github.com/Comcast/scte35-go` library in WASM.
This app compiles the WASM file and runs it in a simple static HTML file.

## Build and Run

Run `make run` to compile the WASM file, copy over your Go `wasm_exec.js` file (the only 
dependency), and launch a Docker nginx instance to host the `index.html` file.

Once started, visit http://localhost:8080 to try the decoder.
