package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
)

var number uint64 = 0

func main() {
	// m := sync.Mutex{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// m.Lock()
		atomic.AddUint64(&number, 1)
		// m.Unlock()
		w.Write([]byte(fmt.Sprintf("Você é o visitante número: %d\n", number)))
	})

	http.ListenAndServe(":3000", nil)
}

// Comando utilizado para testar -->  ab -n 10000 -c 100 http://localhost:3000/
// Apache Bench
