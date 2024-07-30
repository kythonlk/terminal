# Lern how to create webassembly app with go

## Setup

- init go app 
```
go mod init github.com/kythonlk/example & go mod tidy
```

- add js varibale creating package syscall/js
```
go get syscall/js
```

- create main.go and download wasm_exec.js
```
touch main.go
```
https://github.com/golang/go/blob/master/misc/wasm/wasm_exec.js


## Run 

- Create index.html and add wasm_exec.js
```
touch index.html && echo '<script src="wasm_exec.js"></script>' >> index.html
```

- Run
```
GOOS=js GOARCH=wasm go build -o main.wasm main.go
```

- add main.wasm file in to index.html to use wasm 
```
<script src="main.wasm"></script>
```

- run http using python for local otherwise use github pages
```
python -m http.server
```
