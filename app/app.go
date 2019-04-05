package app

import (
	b64 "encoding/base64"
	"encoding/json"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/gorilla/mux"
	"github.com/heaptracetechnology/microservice-memcached/result"
	"io/ioutil"
	"net/http"
	"os"
)

type Message struct {
	Success    string `json:"success"`
	Message    string `json:"message"`
	StatusCode int    `json:"statuscode"`
}

type ArgumentData struct {
	Key   string `json:"key"`
	Value string `json:"topic"`
}

//Set memcached
func SetMemcached(response http.ResponseWriter, request *http.Request) {

	var memcachedHost = os.Getenv("MEMCACHED_HOST")
	var memcachedPort = os.Getenv("MEMCACHED_PORT")

	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		result.WriteErrorResponse(response, err)
		return
	}
	defer request.Body.Close()

	var argumentData ArgumentData
	er := json.Unmarshal(body, &argumentData)
	if er != nil {
		result.WriteErrorResponse(response, er)
		return
	}

	URL := memcachedHost + ":" + memcachedPort
	mc := memcache.New(URL)

	setErr := mc.Set(&memcache.Item{Key: argumentData.Key, Value: []byte(argumentData.Value), Expiration: 0})
	if setErr != nil {
		result.WriteErrorResponse(response, setErr)
		return
	} else {
		message := Message{"true", "Cache is saved", http.StatusBadRequest}
		bytes, _ := json.Marshal(message)
		result.WriteJsonResponse(response, bytes, http.StatusOK)
	}

}

//Get memcached
func GetMemcached(response http.ResponseWriter, request *http.Request) {

	var memcachedHost = os.Getenv("MEMCACHED_HOST")
	var memcachedPort = os.Getenv("MEMCACHED_PORT")

	vars := mux.Vars(request)
	var key = vars["key"]

	URL := memcachedHost + ":" + memcachedPort

	mc := memcache.New(URL)
	res, err := mc.Get(key)
	if err != nil {
		result.WriteErrorResponse(response, err)
		return
	}

	if res != nil {

		if res.Value != nil {
			decode, _ := b64.StdEncoding.DecodeString(string(res.Value))
			res.Value = decode
		}

		bytes, _ := json.Marshal(res)
		result.WriteJsonResponse(response, bytes, http.StatusOK)
	}

}
