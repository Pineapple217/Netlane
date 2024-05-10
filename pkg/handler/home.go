package handler

import (
	"github.com/Pineapple217/Netlane/pkg/view"
	"github.com/labstack/echo/v4"
)

func (h *Handler) Home(c echo.Context) error {
	return render(c, view.Home())
}
