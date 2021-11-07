package graph

// nima qanday ishlaydi degan ma'limotni yozamiz

import (
	"app/graph/generated"
	"app/graph/model"
	"context"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	newToto := model.Todo{
		ID: "X",
		Text: input.Text,
		User: &model.User {
			ID: input.UserID,
		},
	}

	r.todos = append(r.todos, &newToto)

	return &newToto, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {

	return r.todos, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
