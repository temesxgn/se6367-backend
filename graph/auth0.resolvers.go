// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
package graph

import (
	"context"
	"fmt"

	"github.com/temesxgn/se6367-backend/auth"
	ctx2 "github.com/temesxgn/se6367-backend/auth/ctx"
	"github.com/temesxgn/se6367-backend/graph/model"
)

func (r *mutationResolver) UpdateProfile(ctx context.Context, data model.UpdateAuth0Profile) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetProfile(ctx context.Context) (*model.Auth0Profile, error) {
	user := ctx2.GetUser(ctx)
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
			Connection:  uid.Connection,
			UserID:      uid.UserID,
			Provider:    uid.Provider,
			IsSocial:    uid.IsSocial,
			AccessToken: uid.AccessToken,
		}

		ids = append(ids, id)
	}

	fmt.Println(usr.Email)

	return &model.Auth0Profile{
		Email:        usr.Email,
		Nickname:     usr.Nickname,
		PhoneNumber:  usr.PhoneNumber,
		UserMetadata: usr.UserMetadata,
		AppMetadata:  usr.AppMetadata,
		Picture:      usr.Picture,
		Identities:   ids,
	}, nil
}
