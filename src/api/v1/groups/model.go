package groups

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"salu2/src/orm"
)

type Group struct {
	orm.Model
	Name string
}

func NewGroup(name string) *Group {
	return &Group{Name: name}
}

var collectionName = "groups"

func (g *Group) Retrieve() {
	g.AddOne()
	g.Find()
}

func (g *Group) Find() {
	db := orm.GetDB()

	filter := bson.M{"name": g.Name}
	err := g.Read(
		context.Background(),
		db,
		collectionName,
		filter,
		&g,
	)
	if err != nil {
		return
	}
}
func (g *Group) AddOne() {
	db := orm.GetDB()
	filter := bson.M{"name": g.Name}
	var foundGroup Group
	_ = g.Read(
		context.Background(),
		db,
		collectionName,
		filter,
		&foundGroup,
	)

	if foundGroup.Name != "" {
		return
	}

	err := g.Create(
		context.Background(),
		db,
		collectionName,
		g,
	)
	if err != nil {
		return
	}
}
