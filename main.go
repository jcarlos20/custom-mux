package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

type CustomServeMux struct {

}

func (p *CustomServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request){
	if r.URL.Path == "/" {
		giveRandom(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

func giveRandom(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Your random number is: %f", rand.Float64() )
}

func main() {
	// mux := &CustomServeMux{}
	// http.ListenAndServe(":8000", mux)
	newMux := http.NewServeMux()
	
	newMux.HandleFunc("/randomFloat", func (w http.ResponseWriter, r *http.Request){
		fmt.Fprintln(w, rand.Float64())
	})
	
	newMux.HandleFunc("/randomInt", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, rand.Int())
	})
	http.ListenAndServe(":8000", newMux)

}