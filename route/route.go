package route

import (
    "github.com/gorilla/mux"
    "github.com/heaptracetechnology/microservice-memcached/app"
    "log"
    "net/http"
)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
    Route{
        "SetMemcached",
        "POST",
        "/set",
        app.SetMemcached,
    },
    Route{
        "GetMemcached",
        "POST",
        "/get/{key}",
        app.GetMemcached,
    },
}

func NewRouter() *mux.Router {
    router := mux.NewRouter().StrictSlash(true)
    for _, route := range routes {
        var handler http.Handler
        log.Println(route.Name)
        handler = route.HandlerFunc

        router.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(handler)
    }
    return router
}
