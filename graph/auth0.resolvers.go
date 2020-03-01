// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
package graph

import (
	"context"
	"fmt"
	"github.com/temesxgn/se6367-backend/auth/ctx"

	"github.com/temesxgn/se6367-backend/auth"
	"github.com/temesxgn/se6367-backend/graph/generated"
	"github.com/temesxgn/se6367-backend/graph/model"
	auth0 "gopkg.in/auth0.v3"
)

func (r *mutationResolver) UpdateProfile(ctx context.Context) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetProfile(context context.Context) (*model.Auth0User, error) {
	user := ctx.GetUser(context)
	service, err := auth.GetAuthService(auth.AuthZeroAuthServiceType)
	if err != nil {
		fmt.Println("Error getting authentication service " + err.Error())
		return nil, err
	}

	usr, err := service.GetUser(user.UserID())
	if err != nil {
		fmt.Println("ERROR GETTING USER " + user.UserID() + "FROM AUTH0: " + err.Error())
		return nil, err
	}

	ids := make([]*model.UserIdentity, 0)
	for _, uid := range usr.Identities {
		id := &model.UserIdentity{
			Connection: uid.Connection,
			UserID:     uid.UserID,
			Provider:   uid.Provider,
			IsSocial:   uid.IsSocial,
			AccessToken: uid.AccessToken,
		}

		ids = append(ids, id)
	}

	return &model.Auth0User{
		Email:      usr.Email,
		Nickname: usr.Nickname,
		PhoneNumber: usr.PhoneNumber,
		UserMetadata: usr.UserMetadata,
		AppMetadata: usr.AppMetadata,
		Picture: usr.Picture,
		Identities: ids,
	}, nil
}

func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
