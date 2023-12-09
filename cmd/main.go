package main

import (
	"github.com/skeyic/reverseproxy/chain"
	"github.com/skeyic/reverseproxy/filters"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))

	chain.TheChain.Add(filters.NewMonitorFilter())

	http.ListenAndServe(":9090", chain.TheChain)
}
