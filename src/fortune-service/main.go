package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os/exec"
	"time"
)

var port = 4711
var channel = make(chan string, 2)

func main() {
	address := fmt.Sprint("localhost:", port)

	go WriteFortune()
	http.HandleFunc("/", FortuneSelf)
	log.Printf("Starting webserver on http://%v\n", address)
	err := http.ListenAndServe(address, nil)
	log.Panic(err)
}

func FortuneExe(w http.ResponseWriter, r *http.Request) {
	fortuneRaw, _ := exec.Command("fortune").Output()

	fortuneString := string(fortuneRaw)
	fortune := map[string]string{
		"fortune": fortuneString,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(fortune)
}

func WriteFortune() {
	fortunes := []string{"This is fortune #1", "This is fortune #2", "This is no fortune!"}

	for {
		log.Println("That's a good fortune! Lets publish it...")
		channel <- fortunes[rand.Int()%len(fortunes)]
		log.Println("That was stressful. Lets sleep for a while")
		time.Sleep(5 * time.Second)
	}
}

func FortuneSelf(w http.ResponseWriter, r *http.Request) {
	fortuneString := <-channel

	fortune := map[string]string{
		"fortune": fortuneString,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(fortune)
}
