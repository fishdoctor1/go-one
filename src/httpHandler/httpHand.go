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
	case "/price":
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

func (iv inventory) price(w http.ResponseWriter, r *http.Request) {
	searchItem := r.URL.Query().Get("item")
	price, ok := iv[searchItem]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no item: %s", searchItem)
		return
	}
	fmt.Fprintf(w, "%.2f", price)
}

func (iv inventory) items(w http.ResponseWriter, r *http.Request) {
	for k, v := range iv {
		fmt.Fprintf(w, "%s: % .2f", k, v)
	}
}

func (iv inventory) notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "sorry page not found: %s", r.URL)
}

func main() {
	log.SetFlags(0)
	log.Println(os.Getpid())
	inven := inventory{
		"apple":  1.25,
		"orange": 0.99,
	}
	// http.ListenAndServe(":8080", http.HandlerFunc(inventoryFunc)) //EZ start
	// http.ListenAndServe(":8080", inven) //normal start
	// mux := http.NewServeMux()
	// mux.Handle("/items", http.HandlerFunc(inven.items))
	// mux.Handle("/price", http.HandlerFunc(inven.price))
	// mux.Handle("/", http.HandlerFunc(inven.notFound))
	// http.ListenAndServe(":8080", mux) //mux start

	//mux inside
	http.HandleFunc("/items", inven.items)
	http.HandleFunc("/price", inven.price)
	http.HandleFunc("/", inven.notFound)
	http.ListenAndServe(":8080", nil) //mux inside start

}

func inventoryFunc(w http.ResponseWriter, r *http.Request) {
	//EZ template
	log.Println(r.URL.Path)
	switch r.URL.Path {
	case "/items":
		w.Write([]byte("items"))
	case "price":
		w.Write([]byte("price"))
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "sorry imfrom head2 %s", r.URL)

	}
}
