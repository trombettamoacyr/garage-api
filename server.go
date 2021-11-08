package main

import (
	"fmt"
	"github.com/trombettamoacyr/garage-api/config"
	"github.com/trombettamoacyr/garage-api/controller"
	"github.com/trombettamoacyr/garage-api/http"
	"github.com/trombettamoacyr/garage-api/repository"
	"github.com/trombettamoacyr/garage-api/service"
	"net/http"
)

var (
	httpRouter = router.NewMuxRouter()
	//httpRouter = router.NewChiRouter()
	//carRepository    = repository.NewFirestoreRepository()
	carRepository = repository.NewPostgresRepo()
	insuranceService = service.NewCarInsuranceService()
	detailService    = service.NewCarDetailService()
	carService       = service.NewCarService(detailService, insuranceService, carRepository)
	carController    = controller.NewCarController(carService)
)

func main() {
	const port string = ":8080"

	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Up and running...")
		config.GetConnection()
	})

	httpRouter.GET("/cars", carController.GetCars)
	httpRouter.GET("/cars/{id}", carController.GetCarById)
	httpRouter.GET("/cars/detail/{id}", carController.GetCarDetailById)
	httpRouter.POST("/cars", carController.CreateCar)

	httpRouter.SERVER(port)
}
