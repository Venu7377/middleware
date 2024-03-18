package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Logger struct {
	handler http.Handler
}

func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	log.Println(r.URL.Path)

	l.handler.ServeHTTP(w, r)
	log.Printf("%v", time.Since(start))
}

func NewLogger(handlerToWrap http.Handler) *Logger {
	return &Logger{handlerToWrap}
}
func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello this is Home Page"))
}
func abstractHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello this Abstract Page"))
}
func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/abstract",abstractHandler)
	wrappedmux := NewLogger(mux)
	http.ListenAndServe(":8080", wrappedmux)
}

// basic logging middleware

// package main

// import (
// 	"log"
// 	"net/http"
// 	"time"

// 	"github.com/gorilla/mux"
// )
// func logging(f http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		start := time.Now()
// 		log.Println(r.URL.Path)
// 		f(w, r)

// 		log.Println(time.Since(start))

// 	}
// }

// func main() {

// 	mux := http.NewServeMux()

// 	mux.HandleFunc("/hello", logging(helloHandler))
// 	mux.HandleFunc("/abstract", logging(abstractHandler))

// 	http.ListenAndServe(":8081", mux)
// }

// func helloHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("hello this Home Page"))
// }
// func abstractHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("hello this Abstract Page"))
// }
