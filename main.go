package main

import (
	"io"
	"log"
	"net/http"
)

func proxyTorque(w http.ResponseWriter, req *http.Request) {
	client := &http.Client{}
	println(req.URL.String())
	reqHA, err := http.NewRequest(req.Method, "https://xxx"+req.URL.String(), req.Body)
	if err != nil {
		log.Println(err)
		return
	}
	reqHA.Header.Add("Authorization", "Bearer xxx")
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

func main() {
	http.HandleFunc("/api/torque", proxyTorque)
	http.ListenAndServe(":8090", nil)
}
