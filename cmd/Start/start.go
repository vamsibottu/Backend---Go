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

package start

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/Datamigration/internal/crudapi/fields"
	"github.com/google/jsonapi"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
)

var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "shows all the places data ",
	Long:  `Displays all the places present in the postgres table accorind to the given name, placeID, Longitude and lattitude`,
	Run: func(cmd *cobra.Command, args []string) {
		r := mux.NewRouter()
		r.HandleFunc("/places", apiHandler).GetMethods()
		r.HandleFunc("/places/{Name}", placeIDhandler)
		http.ListenAndServe(":8080", r)
	},
}

func apiHandler(w http.ResponseWriter, r *http.Request) {

	pc := fields.PlacesCategory{}
	queryvalues := r.URL.Query()
	x := queryvalues["page"]
	y := queryvalues["category"]

	z := x[0]
	i, er := strconv.Atoi(z)
	if er != nil {
		os.Exit(2)
	}
	s := y[0]
	fmt.Println("page: ", z, "\n", "category:", s)
	start := i
	limit := 15

	err := pc.Find(s, limit, start)
	if err != nil {
		fmt.Println("Err: p.Find():", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	fmt.Println(len(pc), "records found in pg")

	pj := fields.PlaceCategoryJsonApi{}
	pj.Copy(pc)

	fmt.Println(len(pj), "records copied to jsonapi var")

	if err := jsonapi.MarshalPayload(w, pj); err != nil {
		fmt.Println("Err: jsonapi.MarshalPayload():", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
func placeIDhandler(wr http.ResponseWriter, rw *http.Request) {
	vars := mux.Vars(rw)
	wr.WriteHeader(http.StatusOK)
	s := vars["Name"]

	pc := fields.PlaceCategory{}
	err := pc.Get(s)
	if err != nil {
		fmt.Println("Err: p.Get():", err)
		http.Error(wr, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := jsonapi.MarshalPayload(wr, &pc); err != nil {
		fmt.Println("Err: jsonapi.MarshalPayload():", err)
		http.Error(wr, err.Error(), http.StatusInternalServerError)
	}

}
