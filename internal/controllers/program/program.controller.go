package program

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"net/http"
	"program_service/internal/controllers"
	"program_service/internal/core/program/service"
	"program_service/libs"
	"program_service/utils"
)

type ProgramController struct {
	service         *service.ProgramService
	logger          libs.Logger
	externalService *service.ProgramExternalService
}

func New(dbClient *sqlx.DB, logger libs.Logger, url string) *ProgramController {
	return &ProgramController{
		service:         service.New(dbClient),
		externalService: service.NewProgramExternalService(url),
		logger:          logger,
	}
}

func (c *ProgramController) GetGroup() string {
	return "/programs"
}

func (c *ProgramController) GetHandlers() []controllers.ControllerHandler {
	return []controllers.ControllerHandler{
		&controllers.Handler{
			Method:  "POST",
			Path:    "/",
			Handler: c.create,
		},
		&controllers.Handler{
			Method:  "GET",
			Path:    "/{id}",
			Handler: c.get,
		},
		&controllers.Handler{
			Method:  "GET",
			Path:    "/external/{id}",
			Handler: c.externalGet,
		},
		&controllers.Handler{
			Method:  "PUT",
			Path:    "/{id}/update",
			Handler: c.update,
		},
		&controllers.Handler{
			Method:  "DELETE",
			Path:    "/{id}/delete",
			Handler: c.delete,
		},
	}
}

func (c *ProgramController) create(ctx echo.Context) error {
	return nil
}

func (c *ProgramController) get(ctx echo.Context) error {
	return nil
}

func (c *ProgramController) externalGet(ctx echo.Context) error {
	id := ctx.Param("id")
	if id == "" {
		return ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse(http.StatusBadRequest, "id is required"))
	}

	secret := ctx.QueryParam("secret")

	uuid, err := uuid.Parse(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse(http.StatusBadRequest, "invalid id"))
	}
	program, err := c.externalService.Get(uuid, secret)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, utils.NewErrorResponse(http.StatusInternalServerError, err.Error()))
	}
	return ctx.JSON(http.StatusOK, utils.NewSuccessResponse(*program))
}

func (c *ProgramController) update(ctx echo.Context) error {
	return nil
}

func (c *ProgramController) delete(ctx echo.Context) error {
	return nil
}
