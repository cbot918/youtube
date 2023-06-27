
## Go Introduction
Golang 是一個靜態強型別語言, 特色是輕巧, 好學, 原生支援併發, 標準函式庫完整, 內建套件管理系統, 跨平台. 是Google公司的開源專案, Go也受大型開源軟體喜愛, 像是 Docker, Kubernetes, Terraform 等等知名開源. 
此系列針剛入門的朋友做介紹, 希望可以幫忙減少一點摸索時間


## Setup Golang Development
1. [install golang](https://go.dev/doc/install)
2. check version 
```bash 
go version
# go version go1.19.5 linux/amd64
```
## Hello World
1. create repo
```bash
mkdir hello
cd hello
```
2. init go mod
```bash
go mod init hello
```
3. write some code in main.go

```
package main

import (
  "fmt"
)

func main(){
  fmt.Println("hello world")
}
```

4. run code
```bash
go run .
# hello
```

專案  [Github Source Code](https://github.com/cbot918/go-tutor/tree/main/01-hello-world)
