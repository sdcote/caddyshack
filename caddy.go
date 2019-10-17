package main

import (
        "github.com/caddyserver/caddy/caddy/caddymain"

        // http.prometheus
        _ "github.com/miekg/caddy-prometheus"
        // http.ipfilter
        _ "github.com/pyed/ipfilter"
)

func main() {
	caddymain.EnableTelemetry = false
	caddymain.Run()
}
