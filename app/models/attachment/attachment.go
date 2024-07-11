package attachment

import (
	"github.com/goravel/framework/database/orm"
)

type Attachment struct {
	orm.Model
	BlogpostId int
	Title      string
	Url        string
	orm.SoftDeletes
}

func (r *Attachment) Connection() string {
	return "mysql"
}
