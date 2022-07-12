package graph

import "github.com/viniciussevero/graphql-aluno/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	categories []*model.Category
	courses    []*model.Course
	chapters   []*model.Chapter
}
