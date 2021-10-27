package main

import (
	"fmt"
	"github.com/trombettamoacyr/garage-api/controller"
	"github.com/trombettamoacyr/garage-api/http"
	"github.com/trombettamoacyr/garage-api/repository"
	"github.com/trombettamoacyr/garage-api/service"
	"net/http"
)

var (
	httpRouter = router.NewMuxRouter()
	//httpRouter = router.NewChiRouter()
	carRepository = repository.NewFirestoreRepository()
	carService    = service.NewCarService(carRepository)
	carController = controller.NewCarController(carService)
)

func main() {
	const port string = ":8080"

	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Up and running...")
	})

	httpRouter.GET("/cars", carController.GetCars)
	httpRouter.GET("/cars/{id}", carController.GetCarById)
	httpRouter.POST("/cars", carController.CreateCar)

	httpRouter.SERVER(port)
}
