package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct {
	text string
}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, h.text)
}

func main() {
	// http.ListenAndServe(":8081", nil)のようにhandlerの設定をしないとDefaultServerMuxが設定される
	// ServeMux は、HTTP リクエスト マルチプレクサです。各受信リクエストの URL を登録されたパターンのリストと照合し、URL に最も近いパターンのハンドラを呼び出します。
	mux := http.NewServeMux()
	mux.Handle("/api", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "api")
	}))
	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello world")
	}))

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	server.ListenAndServe()

}
