// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
package graph

import (
	"context"
	"fmt"

	"github.com/temesxgn/se6367-backend/graph/generated"
	"github.com/temesxgn/se6367-backend/graph/model"
)

func (r *mutationResolver) UpdateProfile(ctx context.Context) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetProfile(ctx context.Context) (*model.Auth0User, error) {
	panic(fmt.Errorf("not idddmplemented"))
}

func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
