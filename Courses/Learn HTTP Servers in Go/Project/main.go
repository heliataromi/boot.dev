package main

import (
    "fmt"
    "net/http"
    "sync/atomic"
)

type apiConfig struct {
    fileserverHits atomic.Int32
}

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        cfg.fileserverHits.Add(1)
        next.ServeHTTP(w, r)
    })
}

func ReadinessHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/plain; charset=utf-8")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("OK"))
}

func (cfg *apiConfig) MetricsHandler(w http.ResponseWriter, r *http.Request) {
    hits := cfg.fileserverHits.Load()

    w.Header().Set("Content-Type", "text/html")
    w.Write([]byte(fmt.Sprintf(`<html>
  <body>
    <h1>Welcome, Chirpy Admin</h1>
    <p>Chirpy has been visited %d times!</p>
  </body>
</html>`, hits)))
}

func (cfg *apiConfig) ResetHandler(w http.ResponseWriter, r *http.Request) {
    cfg.fileserverHits = atomic.Int32{}
}

func main() {
    serveMux := http.NewServeMux()
    server := &http.Server{Handler: serveMux, Addr: ":8080"}
    apiCfg := &apiConfig{}

    serveMux.Handle("/app/", apiCfg.middlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(".")))))
    serveMux.HandleFunc("GET /api/healthz", ReadinessHandler)
    serveMux.HandleFunc("GET /admin/metrics", apiCfg.MetricsHandler)
    serveMux.HandleFunc("POST /admin/reset", apiCfg.ResetHandler)

    server.ListenAndServe()
}
