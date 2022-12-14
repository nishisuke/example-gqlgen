package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"example/graph/app"
	"example/graph/generated"
	"example/graph/model"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.TodoEdge, error) {
	return app.CreateTodo(ctx, input)
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context, first *int, after *string) (*model.TodoConnection, error) {
	return app.QueryTodos(ctx, first, after)
}

// User is the resolver for the user field.
func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	return app.User(ctx, obj)
}

// TotalCount is the resolver for the totalCount field.
func (r *todoConnectionResolver) TotalCount(ctx context.Context, obj *model.TodoConnection) (int, error) {
	return app.TotalCount(ctx, obj)
}

// Todos is the resolver for the todos field.
func (r *userResolver) Todos(ctx context.Context, obj *model.User, first *int, after *string) (*model.TodoConnection, error) {
	return app.Todos(ctx, obj, first, after)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

// TodoConnection returns generated.TodoConnectionResolver implementation.
func (r *Resolver) TodoConnection() generated.TodoConnectionResolver {
	return &todoConnectionResolver{r}
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
type todoConnectionResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
