package app

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
	"time"
	"upwork/internal/controller/user"
	vacancy2 "upwork/internal/controller/vacancy"
	token2 "upwork/internal/utils/token"
)

type Router struct {
	user    user.Controller
	vacancy vacancy2.Controller
}

func Middleware(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		roleIncome, errRole := token2.ExtractTokenRole(c)
		if strings.ToLower(role) != roleIncome || errRole != nil {
			log.Println(errRole, "err")
			c.String(http.StatusUnauthorized, "You can't do this")
			c.Abort()
			return
		}
		c.Next()
		log.Println(time.Since(t))

	}
}

func CreateRouter(user user.Controller, vacancy vacancy2.Controller) Router {
	return Router{user: user, vacancy: vacancy}
}

func (r Router) UserRouter(engine *gin.Engine) {
	users := engine.Group("/users")
	{
		users.POST("/register", r.user.Register)
	}
}

func (r Router) VacancyRouter(engine *gin.Engine) {
	vacancy := engine.Group("/vacancy")
	{
		// For Candidate
		vacancy.POST("/create", Middleware("User"), r.vacancy.CreateTalent)

		// For HR
		vacancy.GET("/users-for-hr", Middleware("HR"), r.vacancy.AllUsersForHR) // All candidates for HR
		vacancy.PUT("/set-rating/:id", Middleware("HR"), r.vacancy.SetRating)   // Set Rating only HR

		// For Admin
		vacancy.GET("/users-for-admin", Middleware("Admin"), r.vacancy.AllUsersForAdmin) // All candidates for Admin
		vacancy.PUT("/success-user/:id", Middleware("Admin"), r.vacancy.SuccessUser)     // Check rating only Admin
		vacancy.DELETE("/remove-user/:id", Middleware("Admin"), r.vacancy.RemoveUser)    // Check rating only Admin

	}
}
