package api

import (
	"github.com/infrahq/infra/internal/validate"
	"github.com/infrahq/infra/uid"
)

type Destination struct {
	ID         uid.ID                `json:"id"`
	UniqueID   string                `json:"uniqueID" form:"uniqueID" example:"94c2c570a20311180ec325fd56"`
	Name       string                `json:"name" form:"name"`
	Created    Time                  `json:"created"`
	Updated    Time                  `json:"updated"`
	Connection DestinationConnection `json:"connection"`

	Resources []string `json:"resources"`
	Roles     []string `json:"roles"`

	LastSeen  Time `json:"lastSeen"`
	Connected bool `json:"connected"`

	Version string `json:"version"`
}

type DestinationConnection struct {
	// TODO: URL is not a full url, it's set to a host:port by the connector
	URL string `json:"url" example:"aa60eexample.us-west-2.elb.amazonaws.com"`
	CA  PEM    `json:"ca" example:"-----BEGIN CERTIFICATE-----\nMIIDNTCCAh2gAwIBAgIRALRetnpcTo9O3V2fAK3ix+c\n-----END CERTIFICATE-----\n"`
}

func (r DestinationConnection) ValidationRules() []validate.ValidationRule {
	return []validate.ValidationRule{
		validate.Required("ca", r.CA),
	}
}

type ListDestinationsRequest struct {
	Name     string `form:"name"`
	UniqueID string `form:"unique_id"`
	PaginationRequest
}

func (r ListDestinationsRequest) ValidationRules() []validate.ValidationRule {
	// no-op ValidationRules implementation so that the rules from the
	// embedded PaginationRequest struct are not applied twice.
	return nil
}

type CreateDestinationRequest struct {
	UniqueID   string                `json:"uniqueID"`
	Name       string                `json:"name"`
	Version    string                `json:"version"`
	Connection DestinationConnection `json:"connection"`

	Resources []string `json:"resources"`
	Roles     []string `json:"roles"`
}

func (r CreateDestinationRequest) ValidationRules() []validate.ValidationRule {
	return []validate.ValidationRule{
		validate.Required("uniqueID", r.UniqueID),
		validateDestinationName(r.Name),
		validate.Required("name", r.Name),
	}
}

type UpdateDestinationRequest struct {
	ID         uid.ID                `uri:"id" json:"-"`
	Name       string                `json:"name"`
	UniqueID   string                `json:"uniqueID"`
	Version    string                `json:"version"`
	Connection DestinationConnection `json:"connection"`

	Resources []string `json:"resources"`
	Roles     []string `json:"roles"`
}

func (r UpdateDestinationRequest) ValidationRules() []validate.ValidationRule {
	return []validate.ValidationRule{
		validate.Required("uniqueID", r.UniqueID),
		validate.Required("id", r.ID),
		validate.Required("name", r.Name),
		validateDestinationName(r.Name),
	}
}

func (req ListDestinationsRequest) SetPage(page int) Paginatable {
	req.PaginationRequest.Page = page

	return req
}

func validateDestinationName(value string) validate.StringRule {
	rule := ValidateName(value)
	// dots are not allowed in destination name, because it would make grants
	// ambiguous. We use dots to separate destination name from resource name in
	// the Grant.Resource field.
	rule.CharacterRanges = []validate.CharRange{
		validate.AlphabetLower,
		validate.AlphabetUpper,
		validate.Numbers,
		validate.Dash, validate.Underscore,
	}
	return rule
}
