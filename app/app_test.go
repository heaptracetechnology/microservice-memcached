package app

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
)

//Set
var _ = Describe("Set Memcached", func() {
	os.Setenv("MEMCACHED_HOST", "192.168.0.61")
	os.Setenv("MEMCACHED_PORT", "11211")

	setMemcached := ArgumentData{Key: "Hello", Value: "World"}
	requestBody := new(bytes.Buffer)
	json.NewEncoder(requestBody).Encode(setMemcached)
	req, err := http.NewRequest("POST", "/set", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(SetMemcached)
	handler.ServeHTTP(recorder, req)

	Describe("Set", func() {
		Context("Set Memcached", func() {
			It("Should result http.StatusOK", func() {
				Expect(http.StatusOK).To(Equal(recorder.Code))
			})
		})
	})
})

//Get
var _ = Describe("Get Memcached", func() {
	os.Setenv("MEMCACHED_HOST", "192.168.0.61")
	os.Setenv("MEMCACHED_PORT", "11211")

	setMemcached := ArgumentData{Key: "Hello"}
	requestBody := new(bytes.Buffer)
	json.NewEncoder(requestBody).Encode(setMemcached)
	req, err := http.NewRequest("POST", "/get/key", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	vars := map[string]string{
		"key": "Hello",
	}

	req = mux.SetURLVars(req, vars)
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(SetMemcached)
	handler.ServeHTTP(recorder, req)

	Describe("Set", func() {
		Context("Set Memcached", func() {
			It("Should result http.StatusOK", func() {
				Expect(http.StatusOK).To(Equal(recorder.Code))
			})
		})
	})
})
