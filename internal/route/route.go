package route

import (
	"net/http"

	"dift_backend_driver/driver-profile-service/internal/handler"
)

func Register(mux *http.ServeMux, h *handler.ProfileHandler) {
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})
	mux.HandleFunc("/api/v1/driver/profile/settings", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			h.GetProfileSettings(w, r)
			return
		}
		if r.Method == http.MethodPost {
			h.UpdateProfileSettings(w, r)
			return
		}
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	})
	mux.HandleFunc("/api/v1/driver/support/articles", h.GetSupportArticles)
	mux.HandleFunc("/api/v1/driver/support/incidents", h.ReportIncident)
}
