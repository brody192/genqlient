// Code generated by github.com/brody192/genqlient, DO NOT EDIT.

package test

import (
	"github.com/brody192/genqlient/graphql"
	"github.com/brody192/genqlient/internal/testutil"
)

// ListInputQueryResponse is returned by ListInputQuery on success.
type ListInputQueryResponse struct {
	// user looks up a user by some stuff.
	//
	// See UserQueryInput for what stuff is supported.
	// If query is null, returns the current user.
	User ListInputQueryUser `json:"user"`
}

// GetUser returns ListInputQueryResponse.User, and is useful for accessing the field via an interface.
func (v *ListInputQueryResponse) GetUser() ListInputQueryUser { return v.User }

// ListInputQueryUser includes the requested fields of the GraphQL type User.
// The GraphQL type's documentation follows.
//
// A User is a user!
type ListInputQueryUser struct {
	// id is the user's ID.
	//
	// It is stable, unique, and opaque, like all good IDs.
	Id testutil.ID `json:"id"`
}

// GetId returns ListInputQueryUser.Id, and is useful for accessing the field via an interface.
func (v *ListInputQueryUser) GetId() testutil.ID { return v.Id }

// __ListInputQueryInput is used internally by genqlient
type __ListInputQueryInput struct {
	Names []string `json:"names"`
}

// GetNames returns __ListInputQueryInput.Names, and is useful for accessing the field via an interface.
func (v *__ListInputQueryInput) GetNames() []string { return v.Names }

// The query or mutation executed by ListInputQuery.
const ListInputQuery_Operation = `
query ListInputQuery ($names: [String]) {
	user(query: {names:$names}) {
		id
	}
}
`

func ListInputQuery(
	client_ graphql.Client,
	names []string,
) (*ListInputQueryResponse, error) {
	req_ := &graphql.Request{
		OpName: "ListInputQuery",
		Query:  ListInputQuery_Operation,
		Variables: &__ListInputQueryInput{
			Names: names,
		},
	}
	var err_ error

	var data_ ListInputQueryResponse
	resp_ := &graphql.Response{Data: &data_}

	err_ = client_.MakeRequest(
		nil,
		req_,
		resp_,
	)

	return &data_, err_
}

