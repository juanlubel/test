package user

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"salu2/src/api/v1/groups"
	"salu2/src/orm"
	"time"
)

type User struct {
	orm.Model
	Email          string
	Avatar         string
	MainGroup      primitive.ObjectID
	Groups         []primitive.ObjectID
	LastConnection time.Time
}

type Users struct {
	List []User
}

var collectionName = "users"

func NewUser(email string, avatar string) *User {
	return &User{Email: email, Avatar: avatar}
}

func (u *User) setLastConnection() {
	lastConnection := time.Now()
	db := orm.GetDB()
	filter := bson.M{"email": u.Email}
	update := bson.M{"$set": bson.M{"lastconnection": lastConnection}}
	err := u.Update(context.Background(), db, collectionName, filter, update)
	if err != nil {
		return
	}
}

func (u *User) Retrieve() {
	db := orm.GetDB()
	filter := bson.M{"email": u.Email}
	_ = u.Read(
		context.Background(),
		db,
		collectionName,
		filter,
		&u,
	)
}

func (u *User) AddOne() {
	var err error
	db := orm.GetDB()
	filter := bson.M{"email": u.Email}
	var foundUser User
	_ = u.Read(
		context.Background(),
		db,
		collectionName,
		filter,
		&foundUser,
	)

	if foundUser.Email != "" {
		u.setLastConnection()
		return
	}

	group := groups.NewGroup("main")
	group.Retrieve()
	u.MainGroup = group.ID
	u.Groups = append(u.Groups, group.ID)
	err = u.Create(
		context.Background(),
		orm.GetDB(),
		collectionName,
		u,
	)
	if err != nil {
		return
	}

	if err != nil {
		return
	}
}

func (l *Users) GetByGroup(group primitive.ObjectID) {
	var err error
	db := orm.GetDB()
	filter := bson.D{{"groups", group}}
	user := User{}

	err = user.GetAll(context.Background(), db, collectionName, filter, &l.List)
	if err != nil {
		return
	}
}
