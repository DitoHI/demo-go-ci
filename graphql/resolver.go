package graphql

import (
	"context"
	"github.com/vektah/gqlparser/gqlerror"
	graphql "gitlab.com/ditohafizh/demo-go-ci/graphql/resolvers"

	"gitlab.com/ditohafizh/demo-go-ci/api/models"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
	Users []*models.User
	Todos []*models.Todo
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateUser(ctx context.Context, username string, email string) (*models.User, error) {
	newUser := graphql.CreateUser(username, email)

	r.Users = append(r.Users, newUser)
	return newUser, nil
}
func (r *mutationResolver) CreateTodo(ctx context.Context, title string, description *string, userID string) (*models.Todo, error) {
	user := graphql.FindUserById(r.Users, userID)
	if user == nil {
		return nil, gqlerror.Errorf("User not found")
	}

	newTodo := graphql.CreateTodo(title, description, user)

	r.Todos = append(r.Todos, newTodo)
	return newTodo, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) ListUsers(ctx context.Context) ([]*models.User, error) {
	return r.Users, nil
}
func (r *queryResolver) ListTodos(ctx context.Context) ([]*models.Todo, error) {
	return r.Todos, nil
}
