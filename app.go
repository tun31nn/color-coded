package main
 
import (
	"os"
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
                enableCors(&w)
                fmt.Fprintf(w, "{\"color\":\"yellow\",\"version\":\"v1.0.3\"}")
	})

	http.HandleFunc("/dashboard", func(w http.ResponseWriter, r *http.Request) {
                enableCors(&w)
		fmt.Fprintf(w, "<html> DASHBOARD Requested: %s\n </html>", r.URL.Path)
	})

	http.HandleFunc("/die", func(w http.ResponseWriter, r *http.Request) {
                enableCors(&w)
		die();
	})

	http.ListenAndServe(":8080", nil)
}

func die() {
	os.Exit(3)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
