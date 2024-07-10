package factories

import "github.com/brianvoe/gofakeit/v7"

type BlogpostFactory struct {
}

// Definition Define the model's default state.
func (f *BlogpostFactory) Definition() map[string]any {
	return map[string]any{
		"Title": gofakeit.Book().Title,
		"Url":   gofakeit.GlobalFaker.URL(),
	}
}
