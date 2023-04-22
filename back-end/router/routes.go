package router

import (
	"github.com/guioliunb/Ledger-API/back-end/models"
	Home "github.com/guioliunb/Ledger-API/back-end/routes/home"
	Status "github.com/guioliunb/Ledger-API/back-end/routes/status"
)

func GetRoutes() models.Routes{
	return models.Routes{
		models.Route{Name: "Home", Method: "GET", Pattern: "/", HandlerFunc: Home.Index },
		models.Route{Name: "Status", Method: "GET", Pattern: "/status", HandlerFunc: Status.Index },
	}
}