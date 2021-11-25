package model

type Project struct {
	ID          int    `json:"id"`
	Name        string `json:"name,omitempty"`
	CountryCode string `json:"country_code,omitempty"`
	ManagerName string `json:"manager_name,omitempty"`
	IsActive    string `json:"-"`
}
