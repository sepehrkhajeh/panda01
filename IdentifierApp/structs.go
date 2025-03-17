package identifierapp

import "time"

type IdentifierCreateRequest struct {
	DomainID string `json:"domain_id"             validate:"required"`
	Pubkey   string `json:"pubkey"               validate:"required"`
	Name     string `json:"name"                  validate:"required"`
}

type IdentifierGetResponse struct {
	Name           string    `json:"name"`
	Pubkey         string    `json:"pubkey"`
	DomainID       string    `json:"domain_id"`
	ExpiresAt      time.Time `json:"expires_at"`
	FullIdentifier string    `json:"full_identifier"`
}
