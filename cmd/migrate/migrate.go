// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package migrate

import (
	"fmt"
	"log"

	fields "github.com/Datamigration/internal/crudapi/fields"
	"github.com/Datamigration/internal/platform/db/mongo"
	"github.com/spf13/cobra"
)

var MigrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Move required data from mongoDB to Pg",
	Long:  `moved Formatted data required for Places JSon API From MongoDB to PG`,
	Run: func(cmd *cobra.Command, args []string) {
		err := PlacesMongoToPg()
		if err != nil {
			log.Fatalf("error saving to pg: %v", err)
		}
		log.Println("execution completed !!!")
	},
}

func PlacesMongoToPg() error {

	log.Print("PlacesMongoToPg begin")

	// get from mongodb
	ps, err := getPlacesFromMongo()
	if err != nil {
		return err
	}

	dedup := map[string]fields.Place{}
	for _, p := range ps {
		dedup[p.PlaceID.Hex()] = p
	}

	for _, p := range dedup {
		// save to pg: place
		log.Printf("saving %v to pg", p.Name)
		if err = p.Put(); err != nil {
			return err
		}

		for _, pc := range p.DeriveCategory() {
			log.Printf("saving %v to pg", p.Name)
			if err = pc.Put(); err != nil {
				return err
			}
		}

	}
	log.Print("PlacesMongoToPg end")
	return nil
}

func getPlacesFromMongo() ([]fields.Place, error) {

	// connect to mongo
	session, err := mongo.InitMongo()
	if err != nil {
		fmt.Println("could not connect to db ")
		return nil, err
	}
	defer session.Close()

	searchResults := []fields.Place{}
	err = session.DB("lub-db").C("places").Find(nil).All(&searchResults)
	if err != nil {
		fmt.Println("error getting place records from mongodb")
		return nil, err
	}
	return searchResults, nil
}
