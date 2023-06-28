ENTRY=app.go

run: ui-build db-init
	go mod tidy 
	go run $(ENTRY)

ui-build:
	cd ui && npm install && npm run build

db-init:
	echo '* new filedb: auth.db'
	rm -f auth.db
	touch auth.db
	echo '* create table: users'
	sqlite3 auth.db "CREATE TABLE users(id INTEGER PRIMARY KEY AUTOINCREMENT, email TEXT, password TEXT)"
	echo '* insert example row'
	sqlite3 auth.db "INSERT INTO users(email, password) VALUES('sample@gmail.com', '12345');"
	echo '* select * from users: '
	echo ''
	sqlite3 auth.db "select * from users;"
	echo ''






# tutorial command
install-go:
	echo "working on"

install-gowatch:
	go install github.com/silenceper/gowatch@latest

project-init:
	go mod init test-project

ui-init_with_vite_mui:
	curl https://codeload.github.com/mui/material-ui/tar.gz/next | tar -xz --strip=2 material-ui-next/examples/vitejs
	mv vitejs ui


.PHONY: run db-init ui-init_with_vite_mui
.SILENT: run db-init ui-init_with_vite_mui

