package main

import (
	"gopkg.in/mgo.v2/bson"
	"log"
)

type ConfCall struct {
	ConfDateStart string `bson:"ConfDateStart",omitempty`
	ConfDateEnd   string `bson:"ConfDateEnd",omitempty`
	ConfType      string `bson:"ConfType",omitempty`
}
type Prospect struct {
	ProspectID      bson.ObjectId `bson:"ProspectID"`
	Name            string        `bson:"Name",omitempty`
	ConfCalls       []ConfCall    `bson:"ConfCalls",omitempty`
	TechStack       string        `bson:"TechStack",omitempty`
	Domain          string        `bson:"Domain",omitempty`
	DesiredTeamSize int           `bson:"DesiredTeamSize",omitempty`
	ProspectNotes   string        `bson:"ProspectNotes",omitempty`
	ClientNotes     string        `bson:"ClientNotes",omitempty`
	SalesID         string        `bson:"SalesID",omitempty`
	CreateDate      string        `bson:"CreateDate",omitempty`
	StartDate       string        `bson:"StartDate",omitempty`
	BUHead          string        `bson:"BUHead",omitempty`
	TeamSize        int           `bson:"TeamSize",omitempty`
	ProspectStatus  string        `bson:"ProspectStatus",omitempty`
}

func GetAllProspects() (prospects []Prospect) {
	session := gPshServer.session.Copy()
	defer session.Close()

	collection := session.DB(kPreSalesDB).C(kProspectsTable)

	var prospect Prospect
	iter := collection.Find(bson.M{}).Iter()

	for iter.Next(&prospect) {
		prospects = append(prospects, prospect)
	}
	return
}

func (prospect *Prospect) Write() (err error) {
	session := gPshServer.session.Copy()
	defer session.Close()

	collection := session.DB(kPreSalesDB).C(kProspectsTable)

	// insert
	prospect.ProspectID = bson.NewObjectId()
	err = collection.Insert(prospect)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func (prospect *Prospect) Update() (err error) {
	session := gPshServer.session.Copy()
	defer session.Close()

	collection := session.DB(kPreSalesDB).C(kProspectsTable)
	collection.Update(bson.M{"ProspectID": prospect.ProspectID}, prospect)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
