package models

import (
	"encoding/json"
	"time"
)

type AuditEntry struct {
	ID            int64           `json:"id"`
	ServiceSource string          `json:"service_source"`
	ActorID       string          `json:"actor_id"`
	Action        string          `json:"action"`
	EntityType    string          `json:"entity_type"`
	EntityID      string          `json:"entity_id"`
	OldValues     json.RawMessage `json:"old_values"`
	NewValues     json.RawMessage `json:"new_values"`
	IPAddress     string          `json:"ip_address"`
	Status        string          `json:"status"` // Must match ENUM: 'SUCCESS' or 'FAILED'
	ErrorMessage  string          `json:"error_message"`
	RequestID     string          `json:"request_id"`
	CreatedAt     time.Time       `json:"created_at"`
}