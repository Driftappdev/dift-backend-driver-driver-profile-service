package dto

type DriverProfileSettingsResponse struct {
	DriverID         string `json:"driver_id"`
	PhoneNumber      string `json:"phone_number"`
	Email            string `json:"email"`
	EmergencyContact string `json:"emergency_contact"`
	Language         string `json:"language"`
	AutoAccept       bool   `json:"auto_accept"`
	VoiceNavigation  bool   `json:"voice_navigation"`
	ShareLiveTrip    bool   `json:"share_live_trip"`
}

type SupportArticle struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Summary  string `json:"summary"`
	Category string `json:"category"`
	Priority string `json:"priority"`
}

type SupportArticlesResponse struct {
	DriverID string           `json:"driver_id"`
	Items    []SupportArticle `json:"items"`
}

type IncidentReportRequest struct {
	DriverID    string `json:"driver_id"`
	Category    string `json:"category"`
	Description string `json:"description"`
	RouteID     string `json:"route_id,omitempty"`
}
