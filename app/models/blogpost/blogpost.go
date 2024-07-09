package blogpost

import (
	"github.com/goravel/framework/database/orm"
)

type Blogpost struct {
	orm.Model
	Title      string
	AttachFile []string
	Url        string
	orm.SoftDeletes
}

func (r *Blogpost) Connection() string {
	return "mysql"
}
