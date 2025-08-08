package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	backendURLs := []string{
		"http://localhost:8081",
		"http://localhost:8082",
		"http://localhost:8083",
	}

	var backends []*Backend

	for _, backendURL := range backendURLs {
		parsedURL, err := url.Parse(backendURL)
		if err != nil {
			log.Fatalf("Failed to parse backend URL: %s: %v", backendURL, err)
		}

		proxy := httputil.NewSingleHostReverseProxy(parsedURL)

		backend := &Backend{
			URL:   parsedURL,
			Proxy: proxy,
			Alive: true,
		}

		backends = append(backends, backend)
	}

	lb := NewLoadBalancer(backends)

	for i, _ := range backendURLs {
		go func(port int) {
			server := newBackendServer(port)
			log.Printf("Starting backend server on port %d...", port)
			if err := server.ListenAndServe(); err != nil {
				log.Fatalf("Backend server failed on port %d : %v", port, err)
			}
		}(8081 + i)
	}

	LoadBalancerPort := 8080
	log.Printf("Starting load balancer on port %d...", LoadBalancerPort)
	loadBalancerServer := &http.Server{
		Addr:    ":8080",
		Handler: lb,
	}

	if err := loadBalancerServer.ListenAndServe(); err != nil {
		log.Fatalf("Load balancer failed: %v", err)
	}

}
