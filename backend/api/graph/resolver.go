package graph

import "github.com/jamieastley/limbretrievalbot/repository"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	BannedSubredditHandler repository.IBannedSubreddit
	IgnoredUserHandler     repository.IIgnoredUser
}
