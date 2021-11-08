package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"app/graph/generated"
	"app/graph/model"
	"context"
)

func (r *categoryResolver) Classifieds(ctx context.Context, obj *model.Category) ([]*model.Classified, error) {

	return r.DB.GetClassifiedsByCategotyId(*obj.ID), nil
}

// quyidagi e'lonning kategoriyasini bilish
func (r *classifiedResolver) Category(ctx context.Context, obj *model.Classified) (*model.Category, error) {

	return r.DB.GetClassifiedSCategory(r.DB.GetCurrentClassifiedSCategoryId(obj.ID)), nil
	
 }

func (r *queryResolver) Classifieds(ctx context.Context) ([]*model.Classified, error) {
	
	return r.DB.GetClassifieds(), nil
}

func (r *queryResolver) Classified(ctx context.Context, id int) (*model.Classified, error) {
	
	return r.DB.GetClassified(id), nil
}

func (r *queryResolver) Categories(ctx context.Context) ([]*model.Category, error) {
	
	return r.DB.GetCategories(), nil
}

func (r *queryResolver) Category(ctx context.Context, id int) (*model.Category, error) {
	
	return r.DB.GetCategory(id), nil
}

// Category returns generated.CategoryResolver implementation.
func (r *Resolver) Category() generated.CategoryResolver { return &categoryResolver{r} }

// Classified returns generated.ClassifiedResolver implementation.
func (r *Resolver) Classified() generated.ClassifiedResolver { return &classifiedResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type categoryResolver struct{ *Resolver }
type classifiedResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
