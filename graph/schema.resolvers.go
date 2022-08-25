package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"compada/person-service/graph/generated"
	"compada/person-service/graph/model"
	"context"
	"fmt"
	"math/rand"
)

// Persons is the resolver for the persons field.
func (r *queryResolver) Persons(ctx context.Context) ([]*model.Person, error) {
	id := fmt.Sprintf("%d", rand.Intn(100))
	name := "Simba"
	email := "simba@lionking.com"
	person := &model.Person{
		ID:    id,
		Name:  &name,
		Email: &email,
	}
	r.persons = append(r.persons, person)
	return []*model.Person{person}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
