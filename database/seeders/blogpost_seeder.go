package seeders

import (
	models "goravel/app/models/blogpost"

	"github.com/goravel/framework/facades"
)

type BlogpostSeeder struct {
}

// Signature The name and signature of the seeder.
func (s *BlogpostSeeder) Signature() string {
	return "BlogpostSeeder"
}

// Run executes the seeder logic.
func (s *BlogpostSeeder) Run() error {
	post := models.Blogpost{
		Title: "Millioner 2024",
		Url:   "http://satudata.com/pengangguran2024.html",
	}
	return facades.Orm().Query().Create(&post)
}
