package pkg

import (
	"crawl/models"
)

type Pagination struct {
	Limit      int
	Page       int
	Sort       string
	TotalRows  int64
	TotalPages int
	ListPages  []int
	Link       string
	Rows       []models.Article
	Condition  map[string]interface{}
	HasParam   bool
}

func NewPagination() *Pagination {
	return &Pagination{}
}

func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *Pagination) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = 20
	}
	return p.Limit
}

func (p *Pagination) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *Pagination) GetSort() string {
	if p.Sort == "" {
		p.Sort = "Id desc"
	}
	return p.Sort
}

func (p *Pagination) SetListPages() {
	var pages []int
	for i := 1; i <= p.TotalPages; i++ {
		pages = append(pages, i)

	}
	p.ListPages = pages
}

func (p *Pagination) ShowPage() bool {
	if p.TotalPages-3 <= p.Page {
		return true
	}
	return false
}

func (p *Pagination) ShowPage1() bool {
	for i := 1; i <= p.TotalPages; i++ {
		if i >= p.TotalPages-4 {
			return true
		}
	}

	return false
}

func (p *Pagination) ShowPage2() bool {
	for i := 1; i <= p.TotalPages; i++ {
		if p.Page+1 == i || p.Page == i || p.Page-1 == i {
			return true
		}
	}
	return false
}

func (p *Pagination) ShowHref() string {
	if p.HasParam {
		return p.Link + "&"
	} else {
		return p.Link + "?"
	}
}
