package main

import(
    "net/http"

    "github.com/gorilla/mux"
)

type Route struct {
    Name string
    Method string
    Path string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
    Route{
        "Index",
        "GET",
        "/",
        Index,
    },
    Route{
        "GetAllUser",
        "GET",
        "/users",
        GetAllUser,
    },
}

func NewRouter() *mux.Router {
    router := mux.NewRouter().StrictSlash(true)
    
    for _, route := range routes {
        router.
            Methods(route.Method).
            Path(route.Path).
            Name(route.Name).
            Handler(route.HandlerFunc)
    }
    
    return router
}
