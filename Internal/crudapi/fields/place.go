package fields

import (
	"strings"

	"gopkg.in/mgo.v2/bson"
)

type PlacesCategoryJsonApi struct {
	PlaceID     string  `jsonapi:"primary,places"`
	Name        string  `jsonapi:"attr,name"`
	Address     string  `xorm:"'formatted_address'" bson:"formattedaddress" json:"formatted_address,omitempty" jsonapi:"attr,address"`
	Phonenumber string  `xorm:"'formatted_phone_number'" bson:"formattedphonenumber" json:"formatted_phone_number,omitempty" jsonapi:"attr,phonenumbers"`
	Openhours   string  `xorm:"text 'openinghours'" bson:"-" json:"openinghours,omitempty" jsonapi:"attr,openhours"`
	Website     string  `xorm:"'website'" bson:"website" json:"website,omitempty" jsonapi:"attr,website"`
	Longitude   float64 `xorm:"longitude" bson:"-" json:"longitude" jsonapi:"attr,longitude"`
	Latitude    float64 `xorm:"latitude" bson:"-" json:"latitude" jsonapi:"attr,latitude"`
	Category    string  `xorm:"category" bson:"types" json:"types" jsonapi:"attr,category"`
}

type Place struct {
	PlaceID              bson.ObjectId `xorm:"Id" bson:"placeid" json:"Idplace"`
	Name                 string        `xorm:"'name'" bson:"name" json:"name"`
	Icon                 string        `xorm:"'Icon'" bson:"icon" json:"icon"`
	Longitude            float64       `xorm:"longitude" bson:"-" json:"longitude" jsonapi:"attr,longitude"`
	Latitude             float64       `xorm:"latitude" bson:"-" json:"latitude" jsonapi:"attr,latitude"`
	Rating               int           `xorm:"float32 'rating'" bson:"rating" json:"rating"`
	Type                 []string      `xorm:"'Type'" bson:"types" json:"types" jsonapi:"attr,category"`
	Website              string        `xorm:"'website'" bson:"website" json:"website,omitempty" jsonapi:"attr,website"`
	FormattedPhoneNumber string        `xorm:"'formatted_phone_number'" bson:"formattedphonenumber" json:"formatted_phone_number,omitempty" jsonapi:"attr,phonenumbers"`
	FormattedAddress     string        `xorm:"'formatted_address'" bson:"formattedaddress" json:"formatted_address,omitempty" jsonapi:"attr,address"`
	OpeningHours         string        `xorm:"text 'openinghours'" bson:"-" json:"openinghours,omitempty" jsonapi:"attr,openhours"`
	Open                 float64       `xorm:"text 'open'" bson:"-" json:"opennow,omitempty"`
	OpeningHoursMongo    `xorm:"-" bson:"openinghours"`
	Geometry             `xorm:"-" bson:"geometry"`
}

type OpeningHoursMongo struct {
	Weekdaytext []string `xorm:"text '-'"  bson:"weekdaytext"`
	Opennow     float64  `xorm:"'-'"  bson:"opennow"`
}

type Geometry struct {
	Location `xorm:"text '-'" bson:"location" `
}

type Location struct {
	Lat float64 `xorm:"text '-'" bson:"lat"`
	Lng float64 `xorm:"text '-'" bson:"lng"`
}

type Places []Place
type PlaceCategoryJsonApi []*PlacesCategoryJsonApi
type PlacesCategory []PlaceCategory

func init() {
	p := Place{}
	pqxorm.Sync2(&p)

}

func (p *Place) Transform() {

	p.OpeningHours = strings.Join(p.OpeningHoursMongo.Weekdaytext, "|")
	p.Open = p.OpeningHoursMongo.Opennow
	p.Latitude = p.Geometry.Location.Lat
	p.Longitude = p.Geometry.Location.Lng
}

func (pc *PlacesCategory) Transform() []*PlaceCategory {
	p2 := []*PlaceCategory{}

	for _, p1 := range *pc {
		p2 = append(p2, &p1)
	}

	return p2
}

func (pj *PlaceCategoryJsonApi) Copy(pc PlacesCategory) {

	for _, s := range pc {
		t := PlacesCategoryJsonApi{}
		t.PlaceID = s.PlaceID.Hex()
		t.Name = s.Name
		t.Address = s.Address
		t.Phonenumber = s.Phonenumber
		t.Openhours = s.Openhours
		t.Website = s.Website
		t.Longitude = s.Longitude
		t.Latitude = s.Latitude
		t.Category = s.Category

		*pj = append(*pj, &t)
	}

}

func (pc *PlacesCategory) Find(category string, limit int, start int) error {

	err := pqxorm.Where("category = ?", category).Limit(limit, start).Find(pc)
	if err != nil {
		return err

	}

	return nil
}

func (pc *PlaceCategory) Get(Name string) error {

	_, err := pqxorm.Where("Name = ?", Name).Get(pc)
	if err != nil {
		return err

	}

	return nil
}

func (p *Place) Put() error {

	p.Transform()

	_, err := pqxorm.Table(&Place{}).Insert(p)
	if err != nil {
		return err
	}
	return nil
}
