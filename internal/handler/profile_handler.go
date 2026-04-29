package handler

import (
	"encoding/json"
	"net/http"

	response "github.com/driftappdev/libpackage/contracts/response"

	"dift_backend_driver/driver-profile-service/internal/dto"
	"dift_backend_driver/driver-profile-service/internal/service"
)

type ProfileHandler struct{ svc *service.ProfileService }

func NewProfileHandler(svc *service.ProfileService) *ProfileHandler { return &ProfileHandler{svc: svc} }

func (h *ProfileHandler) GetProfileSettings(w http.ResponseWriter, r *http.Request) {
	driverID := r.URL.Query().Get("driver_id")
	if driverID == "" {
		writeJSON(w, http.StatusBadRequest, response.Envelope[any]{Error: &response.AppError{Code: "bad_request", Message: "driver_id required"}})
		return
	}
	writeJSON(w, http.StatusOK, response.Envelope[any]{Data: h.svc.GetProfileSettings(driverID)})
}

func (h *ProfileHandler) UpdateProfileSettings(w http.ResponseWriter, r *http.Request) {
	var req dto.DriverProfileSettingsResponse
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, response.Envelope[any]{Error: &response.AppError{Code: "bad_request", Message: err.Error()}})
		return
	}
	if req.DriverID == "" {
		writeJSON(w, http.StatusBadRequest, response.Envelope[any]{Error: &response.AppError{Code: "bad_request", Message: "driver_id required"}})
		return
	}
	h.svc.UpdateProfileSettings(req)
	w.WriteHeader(http.StatusAccepted)
}

func (h *ProfileHandler) GetSupportArticles(w http.ResponseWriter, r *http.Request) {
	driverID := r.URL.Query().Get("driver_id")
	if driverID == "" {
		writeJSON(w, http.StatusBadRequest, response.Envelope[any]{Error: &response.AppError{Code: "bad_request", Message: "driver_id required"}})
		return
	}
	writeJSON(w, http.StatusOK, response.Envelope[any]{Data: h.svc.GetSupportArticles(driverID)})
}

func (h *ProfileHandler) ReportIncident(w http.ResponseWriter, r *http.Request) {
	var req dto.IncidentReportRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, response.Envelope[any]{Error: &response.AppError{Code: "bad_request", Message: err.Error()}})
		return
	}
	if req.DriverID == "" || req.Category == "" || req.Description == "" {
		writeJSON(w, http.StatusBadRequest, response.Envelope[any]{Error: &response.AppError{Code: "bad_request", Message: "driver_id category description required"}})
		return
	}
	h.svc.ReportIncident(req)
	w.WriteHeader(http.StatusAccepted)
}
