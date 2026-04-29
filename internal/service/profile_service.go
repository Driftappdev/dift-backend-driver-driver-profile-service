package service

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"dift_backend_driver/driver-profile-service/internal/dto"
)

type ProfileService struct {
	mu         sync.RWMutex
	profiles   map[string]dto.DriverProfileSettingsResponse
	incidents  map[string][]dto.IncidentReportRequest
	lastUpdate map[string]time.Time
}

func NewProfileService() *ProfileService {
	return &ProfileService{
		profiles:   map[string]dto.DriverProfileSettingsResponse{},
		incidents:  map[string][]dto.IncidentReportRequest{},
		lastUpdate: map[string]time.Time{},
	}
}

func (s *ProfileService) GetProfileSettings(driverID string) dto.DriverProfileSettingsResponse {
	s.mu.RLock()
	if profile, ok := s.profiles[driverID]; ok {
		s.mu.RUnlock()
		return profile
	}
	s.mu.RUnlock()
	return dto.DriverProfileSettingsResponse{
		DriverID:         driverID,
		PhoneNumber:      "+66-89-000-0001",
		Email:            fmt.Sprintf("%s@driver.drift.local", strings.ToLower(lastN(driverID, 4))),
		EmergencyContact: "+66-89-111-2222",
		Language:         "th",
		AutoAccept:       false,
		VoiceNavigation:  true,
		ShareLiveTrip:    true,
	}
}

func (s *ProfileService) UpdateProfileSettings(req dto.DriverProfileSettingsResponse) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.profiles[req.DriverID] = req
	s.lastUpdate[req.DriverID] = time.Now()
}

func (s *ProfileService) GetSupportArticles(driverID string) dto.SupportArticlesResponse {
	return dto.SupportArticlesResponse{
		DriverID: driverID,
		Items: []dto.SupportArticle{
			{ID: "support-001", Title: "How to handle rider no-shows", Summary: "Wait for the grace period, then submit a no_show cancellation so dispatch and order stay aligned.", Category: "dispatch", Priority: "normal"},
			{ID: "support-002", Title: "Cash trip settlement", Summary: "Collect payment proof before marking a cash trip complete.", Category: "payment", Priority: "high"},
			{ID: "support-003", Title: "Safety escalation", Summary: "File an incident immediately. If the network is unstable, the gateway will retry in the background.", Category: "safety", Priority: "critical"},
		},
	}
}

func (s *ProfileService) ReportIncident(req dto.IncidentReportRequest) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.incidents[req.DriverID] = append(s.incidents[req.DriverID], req)
}

func lastN(value string, n int) string {
	if len(value) <= n {
		return value
	}
	return value[len(value)-n:]
}
