package repositories

import (
	"auditservice/models"
	"database/sql"
)

type AuditRepository interface {
	Save(log models.AuditEntry) error
}

type auditRepo struct {
	Db *sql.DB
}

func NewAuditRepository(db *sql.DB) AuditRepository {
	return &auditRepo{Db: db}
}

func (r *auditRepo) Save(l models.AuditEntry) error {
	query := `INSERT INTO audit_logs (
		service_source, actor_id, action, entity_type, entity_id, 
		old_values, new_values, ip_address, status, error_message, request_id
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	_, err := r.Db.Exec(query, 
		l.ServiceSource, 
		l.ActorID, 
		l.Action, 
		l.EntityType, 
		l.EntityID, 
		l.OldValues, 
		l.NewValues, 
		l.IPAddress, 
		l.Status, 
		l.ErrorMessage, 
		l.RequestID,
	)
	return err
}