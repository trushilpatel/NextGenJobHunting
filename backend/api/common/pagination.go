package common

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"gorm.io/gorm"
)

// PaginationData represents the paginated data and metadata
// swagger:response PaginationData
type PaginationData struct {
	Data       []interface{} `json:"data"`
	Pagination `json:"pagination"`
}

// SaveData saves the provided data into the PaginationData structure
func (p *PaginationData) SaveData(data []interface{}) {
	interfaceData := make([]interface{}, len(data))
	for i, v := range data {
		interfaceData[i] = v
	}
	p.Data = interfaceData
}

// Pagination represents the pagination parameters
// swagger:parameters Pagination
type Pagination struct {
	// Page number, default is 1
	// in: query
	// default: 1
	Page int `form:"page,default=1" json:"page"`
	// Number of items per page, default is 10
	// in: query
	// default: 10
	Limit int `form:"limit,default=10" json:"limit"`
	// Field to sort by, default is "id"
	// in: query
	// default: id
	SortBy string `form:"sortBy,default=id" json:"sortBy"`
	// Sort order (asc or desc), default is "asc"
	// in: query
	// default: asc
	Order string `form:"order,default=asc" json:"order"`
	// Total number of items
	TotalItems int64 `json:"totalItems"`
}

// Validate validates and sets default values for pagination parameters
func (p *Pagination) Validate() {
	if p.Page <= 0 {
		p.Page = 1
	}

	if p.Limit <= 0 {
		p.Limit = 10
	}
	if p.Limit > 100 {
		p.Limit = 100
	}

	if p.SortBy == "" {
		p.SortBy = "id"
	}

	if p.Order != "asc" && p.Order != "desc" {
		p.Order = "asc"
	}
}

// Offset calculates the offset for the database query
func (p *Pagination) Offset() int {
	return (p.Page - 1) * p.Limit
}

// ApplyToDB applies the pagination parameters to the GORM database query
func (p *Pagination) ApplyToDB(db *gorm.DB) *gorm.DB {
	p.SortBy = p.toSnakeCase(p.SortBy) // Convert SortBy to snake case for table column names
	p.Validate()

	if p.Limit > 0 {
		db = db.Limit(p.Limit)
	}

	if p.Page > 0 {
		db = db.Offset(p.Offset())
	}

	if p.SortBy != "" {
		order := "asc"
		if p.Order == "desc" {
			order = "desc"
		}
		db = db.Order(p.SortBy + " " + order)
	}
	fmt.Printf("Pagination parameters: Page=%d, Limit=%d, SortBy=%s, Order=%s\n", p.Page, p.Limit, p.SortBy, p.Order)
	return db
}

// toSnakeCase converts a string to snake_case
func (p *Pagination) toSnakeCase(str string) string {
	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

// ValidateSortBy validates the SortBy field against allowed fields
func (p *Pagination) ValidateSortBy(allowedSortFields []string) error {
	p.SortBy = p.toSnakeCase(p.SortBy) // Convert SortBy to snake case for table column names
	// Convert SortBy value to lowercase for case-insensitive comparison
	sortBy := strings.ToLower(p.SortBy)

	// Iterate through the allowed fields and check if SortBy is valid
	for _, field := range allowedSortFields {
		if sortBy == field {
			return nil
		}
	}
	// If SortBy is not valid, return an error
	return errors.New(fmt.Sprintf("invalid SortBy value: %s", p.SortBy))
}
