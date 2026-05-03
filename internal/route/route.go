package route

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"

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
	mux.HandleFunc("/internal/admin/control", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		secret := strings.TrimSpace(os.Getenv("ADMIN_CONTROL_SHARED_SECRET"))
		if secret != "" && r.Header.Get("X-Admin-Secret") != secret {
			w.WriteHeader(http.StatusUnauthorized)
			_ = json.NewEncoder(w).Encode(map[string]any{"error": "unauthorized"})
			return
		}
		var req map[string]any
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(map[string]any{"error": "invalid_json"})
			return
		}
		action, _ := req["action"].(string)
		if strings.TrimSpace(action) == "" {
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(map[string]any{"error": "action_required"})
			return
		}
		w.WriteHeader(http.StatusAccepted)
		_ = json.NewEncoder(w).Encode(map[string]any{"accepted": true, "service": "driver-profile-service", "action": action})
	})
}
