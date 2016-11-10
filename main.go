package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"pos/config"
)

var Conn *sql.DB

func setConn(conn *sql.DB) {
	Conn = conn
}

func main() {
	mux := NewRouter()
	Conn, _ := config.LoadMysql()
	setConn(Conn)
	fmt.Println("Listen to port 9090")
	http.ListenAndServe(":9090", mux)
}
