package main

import (
	"fmt"
	"github.com/nyudlts/go-aspace"
	"io/ioutil"
	"os"
)

func main() {
	client, err := aspace.NewClient("go-aspace.yml", "stage", 20)
	if err != nil {
		panic(err)
	}

	response, err := client.JsonRequest("/plugins/reindex", "POST", "")
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
	os.Exit(0)
}
