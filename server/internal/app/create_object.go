package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/infrastructure"
	"server/internal/entity"
	"server/util"

	"github.com/gofiber/fiber/v2"
	"github.com/thedevsaddam/govalidator"
)

// @Summary     Create
// @Tags        Object
// @Accept		json
// @Produce		json
// @Param       body	body  dto.CreateExample	true  "body"
// @Router		/object/ [post]
// @Security ApiKeyAuth
func (s *ServiceServer) CreateObject(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	util.PanicIfNeeded(err)
	fileMeta := util.NewFile().SaveLocal(c, file, "default_private")
	s.objectService.Create(entity.CreateObjectReq{
		UUID:     fileMeta.UUID,
		Bukcet:   "default_private",
		Nama:     fileMeta.Nama,
		Size:     fileMeta.Size,
		MimeType: fileMeta.MimeType,
		Url:      fileMeta.Url,
		Path:     fileMeta.Path,
	})
	data := fiber.Map{
		"url": fileMeta.Url,
	}
	return util.NewResponse(c).Success(data, nil, infrastructure.Localize("OK_CREATE"))
}

func (s *ServiceServer) CreateObjectHttp(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "Hello World!")
	rules := govalidator.MapData{
		"file:photo": []string{"ext:jpg,png", "size:100000", "required"},
	}

	messages := govalidator.MapData{
		"file:photo": []string{"ext:Only jpg/png is allowed", "required:Photo is required"},
	}

	opts := govalidator.Options{
		Request:  r,     // request object
		Rules:    rules, // rules map,
		Messages: messages,
	}

	v := govalidator.New(opts)
	e := v.Validate()
	fmt.Println(len(e))
	err := map[string]interface{}{"validationError": e}
	w.Header().Set("Content-type", "applciation/json")

	file, header, errr := r.FormFile("photo")
	fmt.Println(file)
	fmt.Println(header.Filename)
	fmt.Println(errr)
	json.NewEncoder(w).Encode(err)
}
