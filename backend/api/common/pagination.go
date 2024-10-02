package common

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"gorm.io/gorm"
)

type PaginationData struct {
	Data       []interface{} `json:"data"`
	Pagination `json:"pagination"`
}

func (p *PaginationData) SaveData(data []interface{}) {
	interfaceData := make([]interface{}, len(data))
	for i, v := range data {
		interfaceData[i] = v
	}
	p.Data = interfaceData
}

type Pagination struct {
	Page       int    `form:"page,default=1" json:"page"`      // Page number, default is 1
	Limit      int    `form:"limit,default=10" json:"limit"`   // Number of items per page, default is 10
	SortBy     string `form:"sortBy,default=id" json:"sortBy"` // Field to sort by, default is "id"
	Order      string `form:"order,default=asc" json:"order"`  // Sort order (asc or desc), default is "asc"
	TotalItems int64  `json:"totalItems"`                      // Total number of items
}

func (p *Pagination) Validate() {
	if p.Page <= 0 {
		p.Page = 1
	}

	if p.Limit <= 0 {
		p.Limit = 10
	}

	if p.SortBy == "" {
		p.SortBy = "id"
	}

	if p.Order != "asc" && p.Order != "desc" {
		p.Order = "asc"
	}
}

func (p *Pagination) Offset() int {
	return (p.Page - 1) * p.Limit
}

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

func (p *Pagination) toSnakeCase(str string) string {
	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

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
