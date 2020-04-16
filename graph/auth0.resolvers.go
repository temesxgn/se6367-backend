// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
package graph

import (
	"context"
	"github.com/temesxgn/se6367-backend/auth"
	authCtx "github.com/temesxgn/se6367-backend/auth/ctx"
	"github.com/temesxgn/se6367-backend/graph/model"
)

func (r *mutationResolver) UpdateProfile(ctx context.Context, data model.UpdateAuth0Profile) (bool, error) {
	user := authCtx.GetUser(ctx)
	service, err := auth.GetAuthService(auth.AuthZeroAuthServiceType)
	if err != nil {
		return false, err
	}

	if err := service.UpdateProfile(user.UserID(), &data); err != nil {
		return false, err
	}

	return true, nil
}

func (r *queryResolver) GetProfile(ctx context.Context) (*model.Auth0Profile, error) {
	user := authCtx.GetUser(ctx)
	service, err := auth.GetAuthService(auth.AuthZeroAuthServiceType)
	if err != nil {
		return nil, err
	}

	usr, err := service.GetUser(user.UserID())
	if err != nil {
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
