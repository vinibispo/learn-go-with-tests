package blogposts

import (
	"bufio"
	"io"
	"strings"
)

type Post struct {
	Title       string
	Description string
}

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
)

func newPost(postBody io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postBody)

	readMetaLine := func(tagName string) string {
		scanner.Scan()
    return strings.TrimPrefix(scanner.Text(), tagName)
	}

  return Post{
    Title:       readMetaLine(titleSeparator),
    Description: readMetaLine(descriptionSeparator),
  }, nil
}
