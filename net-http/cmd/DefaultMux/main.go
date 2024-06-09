package main

import (
	"fmt"
	"net/http"
)

func main() {
	// ①指定したパスをDefaultServeMuxに保存する
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Print("Hello World")
	}))
	// ②ListenAndServe は TCP ネットワーク アドレス addr をリッスンし、 ハンドラーを使用してServeを呼び出して、着信接続の要求を処理します。
	http.ListenAndServe(":8080", nil)
}
