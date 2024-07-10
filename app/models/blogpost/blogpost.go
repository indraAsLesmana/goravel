package models

import (
	"goravel/database/factories"

	"github.com/goravel/framework/contracts/database/factory"
	"github.com/goravel/framework/database/orm"
)

type Blogpost struct {
	orm.Model
	Title string
	Url   string
	orm.SoftDeletes
}

func (r *Blogpost) Connection() string {
	return "mysql"
}

func (r *Blogpost) Factory() factory.Factory {
	return &factories.BlogpostFactory{}
}
