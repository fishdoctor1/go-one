package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type inventory map[string]float64

func (iv inventory) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	log.Println(r.URL.Path)
	switch r.URL.Path {
	case "/items":
		w.Write([]byte("items"))
	case "price":
		/*api/price?item="sdsdsd"*/
		searchItem := r.URL.Query().Get("item")
		price, ok := iv[searchItem]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "no item:%s", searchItem)
			return
		}
		// w.Write([]byte("price"))
		fmt.Fprintf(w, "%.2f\n", price)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "sorry %s", r.URL)

	}
}

func main() {
	log.SetFlags(0)
	log.Println(os.Getpid())
	inven := inventory{
		"apple":  1.25,
		"orange": 0.99,
	}
	// http.ListenAndServe(":8080", http.HandlerFunc(inventoryFunc))
	http.ListenAndServe(":8080", inven)
}

func inventoryFunc(w http.ResponseWriter, r *http.Request) {
	//EZ
	log.Println(r.URL.Path)
	switch r.URL.Path {
	case "/items":
		w.Write([]byte("items"))
	case "price":
		w.Write([]byte("price"))
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "sorry %s", r.URL)

	}
}
