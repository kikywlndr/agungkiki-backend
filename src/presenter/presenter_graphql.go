package presenter

import (
	"github.com/graphql-go/graphql"
)

// GetAll graphql query
func (p *InvitationPresenter) getAll(params graphql.ResolveParams) (interface{}, error) {
	_, data := p.invitationUsecase.GetAll()
	return data, nil
}

// GetByEmail graphql query
func (p *InvitationPresenter) getByEmail(params graphql.ResolveParams) (interface{}, error) {
	email, _ := params.Args["email"].(string)
	data := p.invitationUsecase.GetByEmail(email)
	return data, nil
}

// GetByName graphql query
func (p *InvitationPresenter) getByName(params graphql.ResolveParams) (interface{}, error) {
	name, _ := params.Args["name"].(string)
	_, data := p.invitationUsecase.GetByName(name)
	return data, nil
}

// getTotalPresent graphql query
func (p *InvitationPresenter) getCount(params graphql.ResolveParams) (interface{}, error) {
	isAttend, _ := params.Args["is_attend"].(bool)
	return p.invitationUsecase.GetCount(isAttend)
}