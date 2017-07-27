package fields

import "gopkg.in/mgo.v2/bson"

type PlaceCategory struct {
	PlaceID     bson.ObjectId `xorm:"PlaceId" bson:"placeid" json:"Idplace"`
	Name        string        `xorm:"'name'" bson:"name" json:"name"`
	Address     string        `xorm:"'formatted_address'" bson:"formattedaddress" json:"formatted_address,omitempty" jsonapi:"attr,address"`
	Phonenumber string        `xorm:"'formatted_phone_number'" bson:"formattedphonenumber" json:"formatted_phone_number,omitempty" jsonapi:"attr,phonenumbers"`
	Openhours   string        `xorm:"text 'openinghours'" bson:"-" json:"openinghours,omitempty" jsonapi:"attr,openhours"`
	Website     string        `xorm:"'website'" bson:"website" json:"website,omitempty" jsonapi:"attr,website"`
	Longitude   float64       `xorm:"longitude" bson:"-" json:"longitude" jsonapi:"attr,longitude"`
	Latitude    float64       `xorm:"latitude" bson:"-" json:"latitude" jsonapi:"attr,latitude"`
	Category    string        `xorm:"category"`
}

func init() {
	pc := PlaceCategory{}
	pqxorm.Sync2(&pc)
}

func (p *Place) DeriveCategory() []PlaceCategory {
	catmap := map[string]bool{}
	for _, t := range p.Type {
		switch string(t) {
		case "meal_delivery", "meal_takeaway", "food":
			catmap["eat"] = true
		case "drink", "cafe", "bar", "beverages":
			catmap["drink"] = true
		}
	}

	cat := []PlaceCategory{}
	for c, _ := range catmap {
		cat = append(cat, PlaceCategory{p.PlaceID, p.Name, p.FormattedAddress, p.FormattedPhoneNumber, p.OpeningHours, p.Website, p.Latitude, p.Longitude, c})
	}

	return cat
}

func (p *PlaceCategory) Put() error {
	_, err := pqxorm.Table(&PlaceCategory{}).Insert(p)
	if err != nil {
		return err
	}
	return nil
}
