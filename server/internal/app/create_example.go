package app

import (
	"fmt"
	"net/http"
	"server/internal/dto"
	"server/internal/entity"
	"server/pkg/util"

	"github.com/gofiber/fiber/v2"
	"github.com/riyan-eng/riyanisgood"
)

// @Summary     Create
// @Tags        Example
// @Accept		json
// @Produce		json
// @Param       body	body  dto.CreateExample	true  "body"
// @Router		/example/ [post]
// @Security ApiKeyAuth
func (s *ServiceServer) CreateExample(c *fiber.Ctx) error {
	body := new(dto.CreateExample)
	err := c.BodyParser(&body)
	util.PanicIfNeeded(err)
	err, errors := riyanisgood.NewValidation().ValidateStruct(*body)
	if err != nil {
		return util.NewResponse(c).Error(errors, util.MESSAGE_FAILED_VALIDATION, fiber.StatusBadRequest)
	}
	s.exampleService.Create(entity.CreateExampleReq{
		Nama:   body.Nama,
		Detail: body.Detail,
	})
	return util.NewResponse(c).Success(nil, nil, util.MESSAGE_OK_CREATE)
}

func (s *ServiceServer) CreateExample2(w http.ResponseWriter, req *http.Request) {
	fmt.Println("post /api")
	body := new(dto.CreateExample)
	// err := c.BodyParser(&body)
	// util.PanicIfNeeded(err)
	// err, errors := riyanisgood.NewValidation().ValidateStruct(*body)
	// if err != nil {
	// 	return util.NewResponse(c).Error(errors, util.MESSAGE_FAILED_VALIDATION, fiber.StatusBadRequest)
	// }
	// s.exampleService.Create(entity.CreateExampleReq{
	// 	Nama:   body.Nama,
	// 	Detail: body.Detail,
	// })
	// return util.NewResponse(c).Success(nil, nil, util.MESSAGE_OK_CREATE)
	fmt.Println(body)
}
