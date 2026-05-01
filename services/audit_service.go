package services

import (
	"auditservice/models"
	"auditservice/repositories"
	"log"
)

type AuditService interface {
	ProcessLog(log models.AuditEntry)
}

type auditService struct {
	repo   repositories.AuditRepository
	queue  chan models.AuditEntry
}

func NewAuditService(repo repositories.AuditRepository) AuditService {
	s := &auditService{
		repo:  repo,
		queue: make(chan models.AuditEntry, 1000),
	}
	// Start background worker
	go s.worker()
	return s
}

func (s *auditService) ProcessLog(l models.AuditEntry) {
	s.queue <- l // Non-blocking push
}

func (s *auditService) worker() {
	for l := range s.queue {
		if err := s.repo.Save(l); err != nil {
			log.Printf("Worker Error: %v", err)
		}
	}
}