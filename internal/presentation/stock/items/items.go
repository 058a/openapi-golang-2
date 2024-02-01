package items

import (
	oapicodegen "openapi/internal/infrastructure/oapicodegen/stock"

	"github.com/labstack/echo/v4"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

type Api struct {
}

func New() *Api {
	return &Api{}
}

func RegisterHandlers(e *echo.Echo, api *Api) {
	oapicodegen.RegisterHandlers(e, api)
}

func (a *Api) PostStockItem(ctx echo.Context) error {
	return PostStockItem(ctx)
}

func (a *Api) DeleteStockItem(ctx echo.Context, stockitemId openapi_types.UUID) error {
	return DeleteStockItem(ctx, stockitemId)
}

func (a *Api) PutStockItem(ctx echo.Context, stockitemId openapi_types.UUID) error {
	return PutStockItem(ctx, stockitemId)
}
