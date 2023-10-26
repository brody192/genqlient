// Code generated by github.com/brody192/genqlient, DO NOT EDIT.

package test

import (
	"github.com/brody192/genqlient/graphql"
	"github.com/brody192/genqlient/internal/testutil"
)

// QueryWithDoubleAliasResponse is returned by QueryWithDoubleAlias on success.
type QueryWithDoubleAliasResponse struct {
	// user looks up a user by some stuff.
	//
	// See UserQueryInput for what stuff is supported.
	// If query is null, returns the current user.
	User QueryWithDoubleAliasUser `json:"user"`
}

// GetUser returns QueryWithDoubleAliasResponse.User, and is useful for accessing the field via an interface.
func (v *QueryWithDoubleAliasResponse) GetUser() QueryWithDoubleAliasUser { return v.User }

// QueryWithDoubleAliasUser includes the requested fields of the GraphQL type User.
// The GraphQL type's documentation follows.
//
// A User is a user!
type QueryWithDoubleAliasUser struct {
	// id is the user's ID.
	//
	// It is stable, unique, and opaque, like all good IDs.
	ID testutil.ID `json:"ID"`
	// id is the user's ID.
	//
	// It is stable, unique, and opaque, like all good IDs.
	AlsoID testutil.ID `json:"AlsoID"`
}

// GetID returns QueryWithDoubleAliasUser.ID, and is useful for accessing the field via an interface.
func (v *QueryWithDoubleAliasUser) GetID() testutil.ID { return v.ID }

// GetAlsoID returns QueryWithDoubleAliasUser.AlsoID, and is useful for accessing the field via an interface.
func (v *QueryWithDoubleAliasUser) GetAlsoID() testutil.ID { return v.AlsoID }

// The query or mutation executed by QueryWithDoubleAlias.
const QueryWithDoubleAlias_Operation = `
query QueryWithDoubleAlias {
	user {
		ID: id
		AlsoID: id
	}
}
`

func QueryWithDoubleAlias(
	client_ graphql.Client,
) (*QueryWithDoubleAliasResponse, error) {
	req_ := &graphql.Request{
		OpName: "QueryWithDoubleAlias",
		Query:  QueryWithDoubleAlias_Operation,
	}
	var err_ error

	var data_ QueryWithDoubleAliasResponse
	resp_ := &graphql.Response{Data: &data_}

	err_ = client_.MakeRequest(
		nil,
		req_,
		resp_,
	)

	return &data_, err_
}

