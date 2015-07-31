package main

import (
	"flag"
	"fmt"
	"github.com/syfaro/haste-client"
	"io/ioutil"
	"log"
	"os"
)

var hasteClient *haste.Haste

func uploadFile(name string) {
	data, err := ioutil.ReadFile(name)
	if err != nil {
		log.Printf("Unable to read file: %s\n", err.Error())
		os.Exit(2)
	}

	resp, err := hasteClient.UploadBytes(data)
	if err != nil {
		log.Printf("Error uploading: %s\n", err.Error())
		os.Exit(3)
	}

	fmt.Println(resp.GetLink(hasteClient))
}

func fetchFile(key string) {
	resp, err := hasteClient.Fetch(key)
	if err != nil {
		log.Printf("Error fetching: %s\n", err.Error())
		os.Exit(3)
	}

	fmt.Print(resp)
}

func main() {
	action := flag.String("action", "upload", "If should upload or fetch")
	host := flag.String("host", "http://hastebin.com", "Host to upload paste to")

	flag.Parse()

	if os.Getenv("HASTE_SERVER") != "" {
		hasteClient = haste.NewHaste(os.Getenv("HASTE_SERVER"))
	} else {
		hasteClient = haste.NewHaste(*host)
	}

	if len(os.Args) == 0 {
		log.Println("You need to provide a filename to upload or key to fetch!")
		os.Exit(1)
	}

	item := os.Args[len(os.Args)-1]

	if *action == "upload" {
		uploadFile(item)
	} else {
		fetchFile(item)
	}
}
