package vacancy

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	vacancy2 "upwork/internal/service/vacancy"
	"upwork/internal/usecase/vacancy"
)

type Controller struct {
	useCase vacancy.UseCase
}

func ControllerVacancy(useCase vacancy.UseCase) Controller {
	return Controller{useCase: useCase}
}

func (ct Controller) CreateTalent(c *gin.Context) {
	ctx := context.Background()

	file, errFile := c.FormFile("cv")
	if errFile != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": errFile.Error(),
			"status":  false,
		})
		return
	}

	file.Filename = strconv.Itoa(rand.Intn(1000)) + "-" + file.Filename

	var create vacancy2.Create

	create.CV = "media/cv/" + file.Filename
	errBind := c.ShouldBind(&create)

	if errBind != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": errBind.Error(),
			"status":  false,
			"err":     "err",
		})
		return
	}

	_, err := ct.useCase.Create(ctx, create)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
		return
	}

	errUpload := c.SaveUploadedFile(file, create.CV)

	if errUpload != nil {
		log.Println(errUpload.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": false,
		"data":    create,
	})
}

func (ct Controller) AllUsersForHR(c *gin.Context) {
	ctx := context.Background()

	users, count, err := ct.useCase.GetAllUsersForHR(ctx)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
		return
	}

	if len(users) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": false,
			"data":    "Candidates not found!",
			"count":   0,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": false,
		"data":    users,
		"count":   count,
	})
}

func (ct Controller) AllUsersForAdmin(c *gin.Context) {
	ctx := context.Background()

	users, count, err := ct.useCase.GetAllUsersForAdmin(ctx)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
		return
	}

	if len(users) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": false,
			"data":    "Candidates not found!",
			"count":   0,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": false,
		"data":    users,
		"count":   count,
	})
}

func (ct Controller) SetRating(c *gin.Context) {
	pk := c.Param("id")
	id, errConv := strconv.Atoi(pk)
	if errConv != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": errConv.Error(),
			"status":  false,
		})
		return
	}
	ctx := context.Background()
	var setRating vacancy2.SetRating

	errBind := c.ShouldBind(&setRating)
	if errBind != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": errBind.Error(),
		})
		return
	}

	_, err := ct.useCase.SetRating(ctx, setRating, id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": false,
		"data":    setRating,
	})

}

func (ct Controller) SuccessUser(c *gin.Context) {
	pk := c.Param("id")
	id, errConv := strconv.Atoi(pk)
	if errConv != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": errConv.Error(),
			"status":  false,
		})
		return
	}
	ctx := context.Background()

	data, err := ct.useCase.SuccessUser(ctx, id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": false,
		"data":    data,
		"success": true,
	})

}

func (ct Controller) RemoveUser(c *gin.Context) {
	pk := c.Param("id")
	id, errConv := strconv.Atoi(pk)
	if errConv != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": errConv.Error(),
			"status":  false,
		})
		return
	}
	ctx := context.Background()

	data, err := ct.useCase.RemoveUser(ctx, id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": false,
		"data":    data,
	})

}
