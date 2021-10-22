package main

import (
	"fmt"
	"github.com/trombettamoacyr/garage-api/controller"
	"github.com/trombettamoacyr/garage-api/http"
	"net/http"
)

var (
	httpRouter = router.NewMuxRouter()
)

func main() {
	const port string = ":8080"

	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Up and running...")
	})

	httpRouter.GET("/cars", controller.GetCars)
	httpRouter.POST("/cars", controller.CreateCar)

	httpRouter.SERVER(port)
}
