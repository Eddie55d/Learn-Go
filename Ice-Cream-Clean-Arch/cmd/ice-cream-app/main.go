package main

import (
	deleteicecream "ice-cream-app/internal/controller/http/v1/delete-ice-cream"
	getallicecreams "ice-cream-app/internal/controller/http/v1/get-all-ice-creams"
	geticecreambyid "ice-cream-app/internal/controller/http/v1/get-ice-cream-by-id"
	posticecream "ice-cream-app/internal/controller/http/v1/post-ice-cream"
	puticecream "ice-cream-app/internal/controller/http/v1/put-ice-cream"
	icecreamrepo "ice-cream-app/internal/controller/repositories/ice-cream-repo"
	router "ice-cream-app/internal/controller/router"
	services "ice-cream-app/internal/controller/services"
	database "ice-cream-app/internal/db"

	"github.com/pkg/errors"
)

func main() {

	DB, err := database.ConnectDB()
	if err != nil {
		errors.Wrap(err, "Database connection fail")
		return
	}
	defer DB.Close()

	IceCreamPgSQLRepo := icecreamrepo.NewIceCreamDbStore(DB)
	IceCreamService := services.NewIceCreamService(*IceCreamPgSQLRepo)
	IceCreamsHandler := getallicecreams.NewIceCreamHandler(IceCreamService)
	IceCreamByIdHandler := geticecreambyid.NewIceCreamHandler(IceCreamService)
	DeleteIceCreamHandler := deleteicecream.NewIceCreamHandler(IceCreamService)
	PostIceCreamHandler := posticecream.NewIceCreamHandler(IceCreamService)
	PutIceCreamHandler := puticecream.NewIceCreamHandler(IceCreamService)

	r := router.New(router.Handlers{
		GetAllIceCreamHandler:  *IceCreamsHandler,
		GetIceCreamByIDHandler: *IceCreamByIdHandler,
		DeleteIceCreamHandler:  *DeleteIceCreamHandler,
		PostIceCreamHandler:    *PostIceCreamHandler,
		PutIceCreamHandler:     *PutIceCreamHandler,
	})

	r.Run("0.0.0.0:8080")
}
