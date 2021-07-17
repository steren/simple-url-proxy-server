package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const friendlyPackageName string = "Simple GET server"

func fetch(url string, authToken string) (string, error) {
	if url == "" {
		return "", errors.New("NO URL PROVIDED")
	}
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	if authToken != "" {
		req.Header.Set("auth-token", authToken)
	}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	defer res.Body.Close()
	responseBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return string(responseBody), nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	res, _ := fetch(os.Getenv("URL"), os.Getenv("AUTH_TOKEN"))

	fmt.Fprint(w, res)
}

func main() {
	log.Printf("%s started.", friendlyPackageName)

	url := os.Getenv("URL")
	if url == "" {
		log.Fatal("URL environment variable is not set")
		return
	}

	http.HandleFunc("/", handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
