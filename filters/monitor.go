package filters

import (
	"github.com/skeyic/reverseproxy/chain"
	"log/slog"
	"net/http"
)

type MonitorFilter struct {
}

func NewMonitorFilter() chain.Filter {
	return &MonitorFilter{}
}

func (f *MonitorFilter) Handle(w http.ResponseWriter, req *http.Request) error {
	slog.Info("MonitorFilter", "URL", req.URL, "Headers", req.Header)
	return nil
}
