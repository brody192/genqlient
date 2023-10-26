// genqlient is a GraphQL client generator for Go.
//
// To run genqlient:
//
//	go run github.com/brody192/genqlient
//
// For programmatic access, see the "generate" package, below.  For
// user documentation, see the project [GitHub].
//
// [GitHub]: https://github.com/brody192/genqlient
package main

import (
	"github.com/brody192/genqlient/generate"
)

func main() {
	generate.Main()
}
