package controller

import (
	"go-calendar-practice/go-calendar-practice/go-calendar/entity"
	"go-calendar-practice/go-calendar-practice/go-calendar/service"

	"github.com/gin-gonic/gin"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) entity.Video
}

type controller struct {
	service service.VideoService
}

func New(service service.VideoService) VideoController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) entity.Video {
	var video entity.Video

	// if e := ctx.BindJSON(&video); e != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{
	// 		"errors": "someError/nHere/nAndYeatOneError",
	// 	})
	// }

	c.service.Save(video)
	return video
}
