package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"app/graph/generated"
	"app/graph/model"
	"context"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input *model.NewUser) (*model.User, error) {
	return r.DB.CreateUser(input), nil
}

func (r *mutationResolver) NewProfession(ctx context.Context, name string) (*model.Profession, error) {
	return r.DB.CreateProfession(name), nil
}

func (r *queryResolver) Servants(ctx context.Context) ([]*model.User, error) {
	return r.GetAllServants(), nil
}

func (r *queryResolver) Professions(ctx context.Context) ([]*model.Profession, error) {
	return r.GetAllProfessions(), nil

}

func (r *queryResolver) Raqamlar(ctx context.Context) ([]int, error) {
	return []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
