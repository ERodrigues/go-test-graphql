package graph

import (
	"github.com/ERodrigues/go-test-graphql/graph/internal/database"
	"google.golang.org/genproto/googleapis/spanner/admin/database/v1"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	CategoryDB *database.Category
}
