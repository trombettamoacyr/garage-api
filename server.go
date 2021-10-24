package main

import (
	"fmt"
	"github.com/trombettamoacyr/garage-api/controller"
	"github.com/trombettamoacyr/garage-api/http"
	"net/http"
)

var (
	httpRouter = router.NewMuxRouter()
	//httpRouter = router.NewChiRouter()

	carController = controller.NewCarController()
)

func main() {
	const port string = ":8080"

	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Up and running...")
	})

	httpRouter.GET("/cars", carController.GetCars)
	httpRouter.POST("/cars", carController.CreateCar)

	httpRouter.SERVER(port)
}
