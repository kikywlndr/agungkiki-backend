package repository

import (
	"github.com/agungdwiprasetyo/agungkiki-backend/src/model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type invitationRepo struct {
	db *mgo.Database
}

// NewInvitationRepository create new repository
func NewInvitationRepository(db *mgo.Database) InvitationRepository {
	repo := new(invitationRepo)
	repo.db = db
	return repo
}

func (r *invitationRepo) FindAll() <-chan Result {
	output := make(chan Result)

	go func() {
		defer close(output)

		var invitations []model.Invitation
		if err := r.db.C("invitations").Find(bson.M{}).All(&invitations); err != nil {
			output <- Result{Error: err}
			return
		}

		output <- Result{Data: invitations}
	}()

	return output
}

func (r *invitationRepo) FindByEmail(email string) <-chan Result {
	output := make(chan Result)

	go func() {
		defer close(output)

		var invitation model.Invitation
		query := bson.M{"email": email}
		if err := r.db.C("invitations").Find(query).One(&invitation); err != nil {
			output <- Result{Error: err}
			return
		}

		output <- Result{Data: &invitation}
	}()

	return output
}

func (r *invitationRepo) FindByName(name string) <-chan Result {
	output := make(chan Result)

	go func() {
		defer close(output)

		var invitations []model.Invitation
		query := r.db.C("invitations").Find(bson.M{"name": bson.M{"$regex": name}})
		count, _ := query.Count()
		if err := query.All(&invitations); err != nil {
			output <- Result{Error: err}
			return
		}

		output <- Result{Count: count, Data: invitations}
	}()

	return output
}

func (r *invitationRepo) CalculateCount(isAttend bool) <-chan Result {
	output := make(chan Result)

	go func() {
		defer close(output)

		query := r.db.C("invitations").Find(bson.M{"is_attend": isAttend})
		count, err := query.Count()
		if err != nil {
			output <- Result{Error: err}
			return
		}

		output <- Result{Count: count}
	}()

	return output
}

func (r *invitationRepo) Save(obj *model.Invitation) <-chan Result {
	output := make(chan Result)

	go func() {
		defer close(output)

		obj.ID = bson.NewObjectId()
		if err := r.db.C("invitations").Insert(obj); err != nil {
			output <- Result{Error: err}
			return
		}

		output <- Result{Data: obj}
	}()

	return output
}