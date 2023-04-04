package slug

import (
	"github.com/gosimple/slug"
)

func Make(text string) string {
	str := "This is a string with spaces and other characters"
	slug := slug.Make(str)
	return slug
}
