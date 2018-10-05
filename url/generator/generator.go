package generator

import "github.com/teris-io/shortid"

type Generator func() string

var DefaultGenerator = func() string {
	id, _ := shortid.Generate()
	return id
}
