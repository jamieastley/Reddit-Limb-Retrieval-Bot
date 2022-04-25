package graph

import "github.com/jamieastley/limbretrievalbot/repository"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	BannedSubreddit repository.IBannedSubreddit
}
