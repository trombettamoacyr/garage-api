package main

import (
	"github.com/swaggo/http-swagger"
	_ "github.com/trombettamoacyr/garage-api/docs"

	"github.com/trombettamoacyr/garage-api/controller"
	"github.com/trombettamoacyr/garage-api/http"
	"github.com/trombettamoacyr/garage-api/repository"
	"github.com/trombettamoacyr/garage-api/service"
	"os"
)

// @title garage-api
// @version 1.0.0
// @host localhost:8080
// @BasePath /cars

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

	httpRouter.SWAGGER("/swagger/*", (httpSwagger.WrapHandler)
	httpRouter.GET("/cars", carController.GetCars)
	httpRouter.GET("/cars/{id}", carController.GetCarById)
	httpRouter.GET("/cars/detail/{id}", carController.GetCarDetailById)
	httpRouter.POST("/cars", carController.CreateCar)

	httpRouter.SERVER(port)
}
