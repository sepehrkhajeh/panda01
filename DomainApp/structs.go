package DomainApp

type DomainCreateResponse struct {
	ID interface{} `json:"id"`
}

type DomainGetResponse struct {
	ID                     interface{} `json:"id"`
	Domain                 string      `json:"domain"`
	BasePricePerIdentifier uint        `json:"base_price_per_identifier"`
	DefaultTTL             uint32      `json:"default_ttl"`
	Status                 string      `json:"status"`
}


type DomainCreateRequest struct {
	Domain                 string `json:"domain"                    validate:"required,hostname"`
	BasePricePerIdentifier uint   `json:"base_price_per_identifier" validate:"required,min=1"`
	DefaultTTL             uint32 `json:"default_ttl"               validate:"required,min=1"`
	Status                 string `json:"status"                    validate:"required,oneof=active inactive"`
}