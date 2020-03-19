package models

import "github.com/temesxgn/se6367-backend/integration/integrationtype"

type Calendar struct {
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Color string `json:"color,omitempty"`
	Integration integrationtype.ServiceType `json:"integration,omitempty"`
}
