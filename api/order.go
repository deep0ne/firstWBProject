package api

import (
	"database/sql"
	"html/template"
	"net/http"

	"github.com/deep0ne/firstWBProject/utils"
	"github.com/gin-gonic/gin"
)

type getOrder struct {
	OrderUID string `uri:"id"`
}

func (server *Server) getOrder(ctx *gin.Context) {
	var req getOrder
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	order, err := server.store.SelectOrder(req.OrderUID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	t, _ := template.New("").Funcs(template.FuncMap{
		"unmarshalDelivery": utils.UnmarshalDelivery,
		"unmarshalItems":    utils.UnmarshalItems,
		"parseTime":         utils.ParseUnix,
	}).ParseFiles("templates/template.html")

	ctx.Header("Content-Type", "text/html; charset=utf-8")
	err = t.ExecuteTemplate(ctx.Writer, "template.html", order)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
}
