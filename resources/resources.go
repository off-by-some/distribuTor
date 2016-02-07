package resources

import (
	L "github.com/Pholey/distribuTor/logger"
	nodeResource "github.com/Pholey/distribuTor/resources/node"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {

		// Set up logging for each request
		handler := L.Logger(route.HandlerFunc, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

var routes = Routes{
	Route{
		"CreateTorConnection",
		"POST",
		"/node/create",
		nodeResource.Create,
	},
	Route{
		"GetTorConnection",
		"GET",
		"/node/{id}",
		nodeResource.One,
	},
	Route{
		"DeleteTorConnection",
		"DELETE",
		"/node/{id}",
		nodeResource.Delete,
	},
	Route{
		"RequestNewIP",
		"PATCH",
		"/node/{id}",
		nodeResource.Update,
	},
}

// TODO: Continue this later
// Preform some light reflection to grab the underlying type
// func getMemberNames(recordType interface{}) []string {
//   ffType := reflect.TypeOf(recordType)
//   memberCount := ffType.NumField()
//   var fnames = make([]string, memberCount)
//
//   for i := 0; i < memberCount; i++ {
//     fnames[i] = ffType.Field(i).Name
//   }
//
//   return fnames
// }
//
// func serializeRow(row *sq.Row, schema interface{}) {
//   keys := getMemberNames(schema)
// }
