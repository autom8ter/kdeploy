package gql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/graphikDB/kdeploy/gen/gql/go/generated"
	"github.com/graphikDB/kdeploy/gen/gql/go/model"
)

func (r *mutationResolver) CreateApp(ctx context.Context, input model.AppInput) (*model.App, error) {
	return r.client.Create(ctx, input)
}

func (r *mutationResolver) UpdateApp(ctx context.Context, input model.AppInput) (*model.App, error) {
	return r.client.Update(ctx, input)
}

func (r *queryResolver) GetApp(ctx context.Context, namespace string) (*model.App, error) {
	return r.client.Get(ctx, namespace)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }