package seeders

import (
	attachmentModel "goravel/app/models/attachment"
	blogpostModel "goravel/app/models/blogpost"
	"strings"

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
	var posts []blogpostModel.Blogpost
	facades.Orm().Factory().Count(5).Make(&posts)
	//create looping for post
	for i := 0; i < len(posts); i++ {
		var attach attachmentModel.Attachment
		// facades.Orm().Factory().Make(&attach)
		title := strings.ReplaceAll(posts[i].Title, " ", "")
		attach.BlogpostId = uint(posts[i].ID)
		attach.Title = title
		attach.Url = "https://static.satudata." + title + ".pdf"
		posts[i].Attachments = append(posts[i].Attachments, &attach)

	}

	// post := models.Blogpost{
	// 	Title: "Millioner 2024",
	// 	Url:   "http://satudata.com/pengangguran2024.html",
	// }
	// return facades.Orm().Query().Create(&post)
	// return facades.Orm().Query().Create(&post)
	// Create all child associations while creating User
	// return facades.Orm().Query().Create(&posts)
	// return facades.Orm().Query().Select(orm.Association).Create(&posts)
	return facades.Orm().Query().Select("Title", "Url", "Attachments").Create(&posts)
}
