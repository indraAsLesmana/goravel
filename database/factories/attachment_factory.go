package factories

import (
	"strings"

	"github.com/brianvoe/gofakeit/v7"
)

type AttachmentFactory struct {
}

// Definition Define the model's default state.
func (f *AttachmentFactory) Definition() map[string]any {
	book := gofakeit.Book().Title
	// extension := []string{".pdf",".xls",".docx"}
	// randIndex := rand.Intn
	return map[string]any{
		"BlogpostId": 1,
		"Title":      book,
		"Url":        "https://static.satudata." + strings.ReplaceAll(book, " ", "") + ".pdf",
	}
}
