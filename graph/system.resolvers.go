// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
package graph

import (
	"context"
	"fmt"

	"github.com/temesxgn/se6367-backend/graph/generated"
	"github.com/temesxgn/se6367-backend/graph/model"
)

func (r *queryResolver) Health(ctx context.Context) (*model.HealthInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
