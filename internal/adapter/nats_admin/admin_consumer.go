package natsadmin

import (
	"context"
	"encoding/json"
	"strings"

	"dift_backend_driver/driver-profile-service/internal/dto"
	"dift_backend_driver/driver-profile-service/internal/service"

	"github.com/nats-io/nats.go"
)

type AdminConsumer struct {
	svc *service.ProfileService
}

func NewAdminConsumer(svc *service.ProfileService) *AdminConsumer { return &AdminConsumer{svc: svc} }

func (c *AdminConsumer) Subscribe(ctx context.Context, nc *nats.Conn, subject string) error {
	_, err := nc.Subscribe(subject, func(msg *nats.Msg) {
		_ = c.handle(msg.Data)
	})
	if err != nil {
		return err
	}
	go func() {
		<-ctx.Done()
		_ = nc.Drain()
	}()
	return nil
}

func (c *AdminConsumer) handle(raw []byte) error {
	var cmd struct {
		Action  string         `json:"action"`
		Payload map[string]any `json:"payload"`
	}
	if err := json.Unmarshal(raw, &cmd); err != nil {
		return err
	}
	switch strings.TrimSpace(cmd.Action) {
	case "driver.profile.update":
		var req dto.DriverProfileSettingsResponse
		b, _ := json.Marshal(cmd.Payload)
		if err := json.Unmarshal(b, &req); err != nil {
			return err
		}
		if strings.TrimSpace(req.DriverID) == "" {
			return nil
		}
		c.svc.UpdateProfileSettings(req)
	case "driver.incident.create":
		var req dto.IncidentReportRequest
		b, _ := json.Marshal(cmd.Payload)
		if err := json.Unmarshal(b, &req); err != nil {
			return err
		}
		if strings.TrimSpace(req.DriverID) == "" {
			return nil
		}
		c.svc.ReportIncident(req)
	}
	return nil
}
