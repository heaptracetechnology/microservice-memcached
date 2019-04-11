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
	os.Setenv("MEMCACHED_HOST", "192.168.0.61")
	os.Setenv("MEMCACHED_PORT", "11211")

	setMemcached := ArgumentData{Key: "Hello", Value: "World"}
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
				Expect(http.StatusOK).To(Equal(recorder.Code))
			})
		})
	})
})

//Test Set method with Invalid data
var _ = Describe("Set Memcached", func() {
	os.Setenv("MEMCACHED_HOST", "192.168.0.61")
	os.Setenv("MEMCACHED_PORT", "11211")

	setMemcached := []byte(`{Key: "Hello", Value: "World"}`)
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
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

//Test Get method with valid data
var _ = Describe("Get Memcached", func() {
	os.Setenv("MEMCACHED_HOST", "192.168.0.61")
	os.Setenv("MEMCACHED_PORT", "11211")

	getMemcached := ArgumentData{Key: "Hello"}
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
		"key": "Hello",
	}

	req = mux.SetURLVars(req, vars)
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetMemcached)
	handler.ServeHTTP(recorder, req)

	Describe("Set", func() {
		Context("Set Memcached", func() {
			It("Should result http.StatusOK", func() {
				Expect(http.StatusOK).To(Equal(recorder.Code))
			})
		})
	})
})

//Test Get method with Invalid data
var _ = Describe("Get Memcached", func() {
	os.Setenv("MEMCACHED_HOST", "192.168.0.61")
	os.Setenv("MEMCACHED_PORT", "11211")

	getMemcached := ArgumentData{Key: "Hello"}
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
