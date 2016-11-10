package main

import (
	"fmt"
	"net/http"
	"pos/account"
)

func Index(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "POS Route")
}

func GetAllUser(res http.ResponseWriter, req *http.Request) {
	users := account.GetAllUser(Conn)
	fmt.Fprint(res, users)
}
