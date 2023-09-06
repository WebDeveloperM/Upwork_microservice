package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"upwork/app"
	userContorller "upwork/internal/controller/user"
	vacancyController "upwork/internal/controller/vacancy"
	postgresql2 "upwork/internal/pkg/postgresql"
	repoUser "upwork/internal/repository/postgresql/user"
	repoVacancy "upwork/internal/repository/postgresql/vacancy"
	serviceUser "upwork/internal/service/user"
	serviceVacancy "upwork/internal/service/vacancy"
	usecaseUser "upwork/internal/usecase/user"
	usecaseVacancy "upwork/internal/usecase/vacancy"
)

func main() {
	server := gin.New()

	server.Use(gin.Logger())
	server.Use(gin.Recovery())

	server.StaticFS("/media", http.Dir("media"))

	db := postgresql2.ConnectPostgresql()

	repoUser := repoUser.RepositoryUser(db)
	repoVacancy := repoVacancy.RepositoryVacancy(db)

	serviceUser := serviceUser.ServiceUser(repoUser)
	serviceVacancy := serviceVacancy.ServiceVacancy(repoVacancy)

	useCaseUser := usecaseUser.UserUseCase(serviceUser)
	useCaseVacancy := usecaseVacancy.VacancyUseCase(serviceVacancy)

	uController := userContorller.ControllerUser(useCaseUser)
	vContorller := vacancyController.ControllerVacancy(useCaseVacancy)

	router := app.CreateRouter(uController, vContorller)
	router.UserRouter(server)
	router.VacancyRouter(server)

	err := server.Run(":8000")

	if err != nil {
		log.Fatal(err)
	}
}
