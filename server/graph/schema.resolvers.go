package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"math/rand"

	"github.com/martel-darin/graphql_app/db"
	"github.com/martel-darin/graphql_app/graph/generated"
	"github.com/martel-darin/graphql_app/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo := model.Todo{
		ID:   rand.Intn(1000000),
		Text: input.Text,
		Done: false,
		User: &model.User{
			ID: input.UserID,
		},
	}
	if err := db.InsertTodo(ctx, &todo); err != nil {
		return nil, err
	}
	return &todo, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	user := &model.User{
		ID:   rand.Intn(1000000),
		Name: input.Name,
	}
	if err := db.InsertUser(ctx, user); err != nil {
		return nil, err
	}
	return user, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	todos, err := db.FetchTodos(ctx)
	if err != nil {
		return nil, err
	}

	return todos, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	users, err := db.FetchUsers(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
