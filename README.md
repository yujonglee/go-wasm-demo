# Go-Wasm-Demo

## Build
### Using `official` compiler

```shell
GOOS=js GOARCH=wasm go build -o main.wasm
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
node wasm_exec.js main.wasm
```

### Using `tinygo` compiler
```shell
tinygo build -target wasm -o main.wasm --no-debug
cp "$(tinygo env TINYGOROOT)/targets/wasm_exec.js" .
node wasm_exec.js main.wasm
```

Seems like `tinygo` has some issues. Such as [this](https://github.com/tinygo-org/tinygo/issues/1140).

## Run
```shell
node wasm_exec.js main.wasm
```

or

```shell
serve
```

## Etc
`.vscode/settings.json` is there for `import "syscall/js"`.
