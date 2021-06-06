package main

import (
	"flag"
	"io"
	"log"
	"net/http"
)

func proxyTorque(url string, token string) func (http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		client := &http.Client{}
		println(req.URL.String())
		reqHA, err := http.NewRequest(req.Method, url+req.URL.String(), req.Body)
		if err != nil {
			log.Println(err)
			return
		}
		reqHA.Header.Add("Authorization", "Bearer "+token)
		resHA, err := client.Do(reqHA)
		if err != nil {
			log.Println(err)
			return
		}
		defer resHA.Body.Close()

		for name, values := range resHA.Header {
			w.Header()[name] = values
		}
		w.WriteHeader(resHA.StatusCode)
		io.Copy(w, resHA.Body)
	}
}

func main() {
	url := flag.String("url", "", "Home Assistant URL")
	token := flag.String("token", "", "Home Assistant Long Lived Token")
	flag.Parse()
	if *url == "" || *token == "" {
		flag.PrintDefaults()
        panic("Please provide both url and token")
	}
	http.HandleFunc("/api/torque", proxyTorque(*url, *token))
	http.ListenAndServe(":8090", nil)
}
