package pprof

import (
	"fmt"
	"net/http"
	"testing"
)

func TestPprof(t *testing.T)  {
	ip := "0.0.0.0:6060"
	if err := http.ListenAndServe(ip, nil); err != nil {
		fmt.Printf("start pprof failed on %s\n", ip)
	}
}
