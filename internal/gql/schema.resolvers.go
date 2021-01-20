package gql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/autom8ter/meshpaas/gen/gql/go/generated"
	"github.com/autom8ter/meshpaas/gen/gql/go/model"
	meshpaaspb "github.com/autom8ter/meshpaas/gen/grpc/go"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func (r *mutationResolver) DelProject(ctx context.Context, input *string) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateApp(ctx context.Context, input model.AppInput) (*model.App, error) {
	app, err := r.client.CreateApp(ctx, toApp(input))
	if err != nil {
		return nil, &gqlerror.Error{
			Message: err.Error(),
			Path:    graphql.GetPath(ctx),
		}
	}
	return fromApp(app), nil
}

func (r *mutationResolver) UpdateApp(ctx context.Context, input model.AppInput) (*model.App, error) {
	app, err := r.client.UpdateApp(ctx, toApp(input))
	if err != nil {
		return nil, &gqlerror.Error{
			Message: err.Error(),
			Path:    graphql.GetPath(ctx),
		}
	}
	return fromApp(app), nil
}

func (r *mutationResolver) DelApp(ctx context.Context, input model.Ref) (*string, error) {
	_, err := r.client.DeleteApp(ctx, &meshpaaspb.Ref{
		Name: input.Name,
	})
	if err != nil {
		return nil, &gqlerror.Error{
			Message: err.Error(),
			Path:    graphql.GetPath(ctx),
		}
	}
	return nil, nil
}

func (r *mutationResolver) CreateTask(ctx context.Context, input model.TaskInput) (*model.Task, error) {
	task, err := r.client.CreateTask(ctx, toTask(input))
	if err != nil {
		return nil, &gqlerror.Error{
			Message: err.Error(),
			Path:    graphql.GetPath(ctx),
		}
	}
	return fromTask(task), nil
}

func (r *mutationResolver) UpdateTask(ctx context.Context, input model.TaskInput) (*model.Task, error) {
	task, err := r.client.UpdateTask(ctx, toTask(input))
	if err != nil {
		return nil, &gqlerror.Error{
			Message: err.Error(),
			Path:    graphql.GetPath(ctx),
		}
	}
	return fromTask(task), nil
}

func (r *mutationResolver) DelTask(ctx context.Context, input model.Ref) (*string, error) {
	_, err := r.client.DeleteApp(ctx, &meshpaaspb.Ref{
		Name: input.Name,
	})
	if err != nil {
		return nil, &gqlerror.Error{
			Message: err.Error(),
			Path:    graphql.GetPath(ctx),
		}
	}
	return nil, nil
}

func (r *queryResolver) GetApp(ctx context.Context, input model.Ref) (*model.App, error) {
	app, err := r.client.GetApp(ctx, &meshpaaspb.Ref{
		Name: input.Name,
	})
	if err != nil {
		return nil, &gqlerror.Error{
			Message: err.Error(),
			Path:    graphql.GetPath(ctx),
		}
	}
	return fromApp(app), nil
}

func (r *queryResolver) ListApps(ctx context.Context, input *string) ([]*model.App, error) {
	apps, err := r.client.ListApps(ctx, &empty.Empty{})
	if err != nil {
		return nil, &gqlerror.Error{
			Message: err.Error(),
			Path:    graphql.GetPath(ctx),
		}
	}
	var toReturn []*model.App
	for _, a := range apps.GetApplications() {
		toReturn = append(toReturn, fromApp(a))
	}
	return toReturn, nil
}

func (r *queryResolver) GetTask(ctx context.Context, input model.Ref) (*model.Task, error) {
	app, err := r.client.GetTask(ctx, &meshpaaspb.Ref{
		Name: input.Name,
	})
	if err != nil {
		return nil, &gqlerror.Error{
			Message: err.Error(),
			Path:    graphql.GetPath(ctx),
		}
	}
	return fromTask(app), nil
}

func (r *queryResolver) ListTasks(ctx context.Context, input *string) ([]*model.Task, error) {
	apps, err := r.client.ListTasks(ctx, &empty.Empty{})
	if err != nil {
		return nil, &gqlerror.Error{
			Message: err.Error(),
			Path:    graphql.GetPath(ctx),
		}
	}
	var toReturn []*model.Task
	for _, a := range apps.GetTasks() {
		toReturn = append(toReturn, fromTask(a))
	}
	return toReturn, nil
}

func (r *subscriptionResolver) StreamLogs(ctx context.Context, input model.Ref) (<-chan string, error) {
	stream, err := r.client.StreamLogs(ctx, &meshpaaspb.Ref{
		Name: input.Name,
	})
	if err != nil {
		return nil, &gqlerror.Error{
			Message: err.Error(),
			Path:    graphql.GetPath(ctx),
		}
	}
	ch := make(chan string)
	go func() {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		for {
			select {
			case <-ctx.Done():
				close(ch)
				return
			default:
				msg, err := stream.Recv()
				if err != nil {
					ch <- err.Error()
				}
				ch <- msg.GetMessage()
			}
		}
	}()
	return ch, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
