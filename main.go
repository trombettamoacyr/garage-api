package main

import (
	"os"

	"github.com/trombettamoacyr/garage-api/controller"
	"github.com/trombettamoacyr/garage-api/http"
	"github.com/trombettamoacyr/garage-api/repository"
	"github.com/trombettamoacyr/garage-api/service"
)

var (
	httpRouter = router.NewMuxRouter()
	//httpRouter = router.NewChiRouter()
	carRepository = repository.NewFirestoreRepository()
	//carRepository = repository.NewPostgresRepo()
	insuranceService = service.NewCarInsuranceService()
	detailService    = service.NewCarDetailService()
	carService       = service.NewCarService(detailService, insuranceService, carRepository)
	carController    = controller.NewCarController(carService)
)

func main() {
	port := os.Getenv("GARAGE_API_PORT")

	httpRouter.GET("/cars", carController.GetCars)
	httpRouter.GET("/cars/{id}", carController.GetCarById)
	httpRouter.GET("/cars/detail/{id}", carController.GetCarDetailById)
	httpRouter.POST("/cars", carController.CreateCar)

	httpRouter.SERVER(port)
}
