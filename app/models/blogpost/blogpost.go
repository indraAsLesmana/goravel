package blogpost

import (
	models "goravel/app/models/attachment"
	"goravel/database/factories"

	"github.com/goravel/framework/contracts/database/factory"
	"github.com/goravel/framework/database/orm"
)

type Blogpost struct {
	orm.Model
	Title       string
	Url         string
	Attachments []*models.Attachment
	orm.SoftDeletes
}

func (r *Blogpost) Connection() string {
	return "mysql"
}

func (r *Blogpost) Factory() factory.Factory {
	return &factories.BlogpostFactory{}
}
