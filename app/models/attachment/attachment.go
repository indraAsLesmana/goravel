package attachment

import (
	"goravel/database/factories"

	"github.com/goravel/framework/contracts/database/factory"
	"github.com/goravel/framework/database/orm"
)

type Attachment struct {
	orm.Model
	BlogpostId uint
	Title      string
	Url        string
	orm.SoftDeletes
}

func (r *Attachment) Connection() string {
	return "mysql"
}

func (r *Attachment) Factory() factory.Factory {
	return &factories.AttachmentFactory{}
}
