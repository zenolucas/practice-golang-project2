package CMD

import (
	"github.com/gorilla/mux"
	"github.com/zenolucas/practice-golang-project2/PKG/routes"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))

}
