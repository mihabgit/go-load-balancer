package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
	"sync/atomic"
)

type Backend struct {
	URL   *url.URL
	Proxy *httputil.ReverseProxy
	Alive bool
	mu    sync.RWMutex // Mutex to protect the Alive field
}

func (b *Backend) SetAlive(alive bool) {
	b.mu.Lock()
	b.Alive = alive
	b.mu.Unlock()
}

func (b *Backend) IsAlive() (alive bool) {
	b.mu.RLock()
	alive = b.Alive
	b.mu.RUnlock()
	return
}

type LoadBalancer struct {
	backends []*Backend
	current  uint64
}

func NewLoadBalancer(backends []*Backend) *LoadBalancer {
	return &LoadBalancer{
		backends: backends,
		current:  0,
	}
}

func (lb *LoadBalancer) GetNextBackendIndex() int {
	next := atomic.AddUint64(&lb.current, 1) % uint64(len(lb.backends))
	return int(next)
}

func (lb *LoadBalancer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	backendIndex := lb.GetNextBackendIndex()
	backend := lb.backends[backendIndex]

	log.Printf("Forwarding request to backend %s\n", backend.URL)

	backend.Proxy.ServeHTTP(w, r)
}
