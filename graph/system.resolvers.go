// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
package graph

import (
	"context"
	"fmt"

	"github.com/temesxgn/se6367-backend/graph/generated"
	"github.com/temesxgn/se6367-backend/graph/model"
)

func (r *healthInfoResolver) Auth0Connection(ctx context.Context, obj *model.HealthInfo) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *healthInfoResolver) DatabaseConnection(ctx context.Context, obj *model.HealthInfo) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Restart(ctx context.Context) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Health(ctx context.Context) (*model.HealthInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *Resolver) HealthInfo() generated.HealthInfoResolver { return &healthInfoResolver{r} }
func (r *Resolver) Mutation() generated.MutationResolver     { return &mutationResolver{r} }
func (r *Resolver) Query() generated.QueryResolver           { return &queryResolver{r} }

type healthInfoResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
