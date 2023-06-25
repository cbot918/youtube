# Auth
A Learning Purpose Project To Help Entry Level Go-React Developer Quick Learning

# Test Environment
- Linux Ubuntu 20.04 ( MacOS/WSL2 maybe run well also)


# Pre-Requesties
- [Golang](https://go.dev/)
- [Makefile](https://github.com/michaelfromyeg/makefiles)


# Stack
- Server: Golang
- Client: React (Vite/MaterialUI)
- DB: Sqlite3


# Run App
```bash
make run
```
then browse localhost:3005


# Dev Info
- server config is in main.go
- main feature in server.go and dao.go


# Repo Structure
in server folder
- server.go: server main
- util.go: include a allow cors setup function
- model.go: model
- dao.go: db operation
- config.go: a struct to load config 


# Deploy
local
```
docker build -t _iname .
docker run -it -p 3005:3005 --name _cname _iname
```
cloud: [Zeabur](https://zeabur.com/)