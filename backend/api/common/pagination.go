package common

import (
	"fmt"

	"gorm.io/gorm"
)

type Pagination struct {
	Page   int    `form:"page,default=1" json:"page"`       // Page number, default is 1
	Limit  int    `form:"limit,default=10" json:"limit"`    // Number of items per page, default is 10
	SortBy string `form:"sort_by,default=id" json:"sortBy"` // Field to sort by, default is "id"
	Order  string `form:"order,default=asc" json:"order"`   // Sort order (asc or desc), default is "asc"
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
