package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {

	// Test endpoint
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Method: "+r.Method+"\n")
		fmt.Fprintf(w, "URL: "+r.URL.Host+"\n")
		fmt.Fprintf(w, "Host: "+r.Host+"\n")
	})

	// Teapot IOT endpoint
	http.HandleFunc("/teapot", func(w http.ResponseWriter, r *http.Request) {
		response := []byte("I'm pretending to be a teapot. For the time being. Sorry for any inconvenience.")
		w.WriteHeader(418) // I'm a teapot
		w.Write(response)
	})

	// Puns endpoint
	http.HandleFunc("/pun", func(w http.ResponseWriter, r *http.Request) {
		rand.Seed(time.Now().Unix())
		jokes := []string{
			"Q: Why did the scarecrow win an award?\nA: Because he was outstanding in his field",
			"Q: What did the inspector molecule say to the suspect molecule?\nA: \"I've got my ion you\"",
			"Q: Did you hear about the cheese factory that exploded in France?\nA: There was nothing left but de brie",
			"I just watched a documentary on beavers. It was the best dam program I've ever seen.",
			"A prisoner asked to take his own mugshot. They called it a cellfie.",
			"I'm not addicted to brake fluid, I can stop whenever I want.",
			"Q: What do you call cheese by itself?\nA: Provolone",
			"Q: Why do crabs never give to charity?\nA: Because they're shellfish",
			"Don't trust atoms. They make up everything.",
			"Q: Why can't you hear a pterodactyl go to the bathroom?\nA: Because the P is silent.",
		}

		// pick a random pun index
		n := rand.Int() % len(jokes)

		fmt.Fprintf(w, jokes[n])
	})

	// Start serving
	log.Fatal(http.ListenAndServe(":2983", nil))
}
