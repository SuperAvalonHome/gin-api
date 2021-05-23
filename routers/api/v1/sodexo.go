package v1

import (
	"net/http"

	"github.com/unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"github.com/SuperAvalonHome/gin-api/pkg/app"
	"github.com/SuperAvalonHome/gin-api/pkg/e"
	"github.com/SuperAvalonHome/gin-api/service/sodexo_service"
)

// @Summary Get a single article
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/articles/{id} [get]
func GetSodexoUser(c *gin.Context) {
	appG := app.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	sodexoService := sodexo_service.SodexoUser{Id: id}
	sodexoUser, err := sodexoService.Get()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_ARTICLE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, sodexoUser)
}


func GetSodexoTrade(c *gin.Context) {
	appG := app.Gin{C: c}
	order_id := c.Param("order_id")

	sodexoService := sodexo_service.SodexoTrade{OrderId: order_id}
	sodexoTrade, err := sodexoService.Get()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_ARTICLE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, sodexoTrade)
}


func GetSodexoRefund(c *gin.Context) {
	appG := app.Gin{C: c}
	order_id := c.Param("order_id")

	sodexoService := sodexo_service.SodexoRefund{OrderId: order_id}
	sodexoRefund, err := sodexoService.Get()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_ARTICLE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, sodexoRefund)
}
