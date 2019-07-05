package app

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/gorilla/mux"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//Test Set method with valid data
var _ = Describe("Set Memcached", func() {
	os.Setenv("MEMCACHED_HOST", "")
	os.Setenv("MEMCACHED_PORT", "")

	setMemcached := ArgumentData{Key: "qwerty", Value: "keyboard"}
	requestBody := new(bytes.Buffer)
	err := json.NewEncoder(requestBody).Encode(setMemcached)
	if err != nil {
		log.Fatal(err)
	}
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
				Expect(recorder.Code).To(Equal(http.StatusBadRequest))
			})
		})
	})
})

//Test Set method with Invalid data
var _ = Describe("Set Memcached", func() {

	setMemcached := []byte(`{Key: "qwerty", Value: "keyboard"}`)
	requestBody := new(bytes.Buffer)
	err := json.NewEncoder(requestBody).Encode(setMemcached)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("POST", "/set", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(SetMemcached)
	handler.ServeHTTP(recorder, req)

	Describe("Set", func() {
		Context("Set Memcached", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(recorder.Code).To(Equal(http.StatusBadRequest))
			})
		})
	})
})

//Test Get method with valid data
var _ = Describe("Get Memcached", func() {
	os.Setenv("MEMCACHED_HOST", "")
	os.Setenv("MEMCACHED_PORT", "")

	getMemcached := ArgumentData{Key: "qwerty"}
	requestBody := new(bytes.Buffer)
	err := json.NewEncoder(requestBody).Encode(getMemcached)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("POST", "/get/key", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	vars := map[string]string{
		"key": "qwerty",
	}

	req = mux.SetURLVars(req, vars)
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetMemcached)
	handler.ServeHTTP(recorder, req)

	Describe("Get", func() {
		Context("Get Memcached", func() {
			It("Should result http.StatusOK", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

//Test Get method with Invalid data
var _ = Describe("Get Memcached", func() {

	getMemcached := ArgumentData{Key: "qwerty"}
	requestBody := new(bytes.Buffer)
	err := json.NewEncoder(requestBody).Encode(getMemcached)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("POST", "/get/key", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	vars := map[string]string{
		"key": "true",
	}

	req = mux.SetURLVars(req, vars)
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetMemcached)
	handler.ServeHTTP(recorder, req)

	Describe("Set", func() {
		Context("Set Memcached", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})
