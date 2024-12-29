package controller

import (
	"fmt"
	"net/http"

	"github.com/1067rail/portto-meme-coin/dto"
	"github.com/1067rail/portto-meme-coin/httputil"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (c *Controller) CreateCoin(ctx *gin.Context) {
	var body dto.CreateCoinReq
	if err := ctx.Bind(&body); err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	if result, err := c.coinService.CreateCoin(body); err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	} else {
		ctx.JSON(http.StatusOK, result)
	}
}

func (c *Controller) GetCoin(ctx *gin.Context) {
	var id *uuid.UUID

	if _id, err := uuid.Parse(ctx.Param("id")); err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	} else {
		id = &_id
	}

	if result, err := c.coinService.GetCoin(*id); err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	} else {
		ctx.JSON(http.StatusOK, result)
	}
}

func (c *Controller) UpdateCoin(ctx *gin.Context) {
	id := ctx.Param("id")

	var body interface{}
	if err := ctx.Bind(&body); err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	fmt.Println("Controller.GetCoin", id, body)
	c.coinService.UpdateCoin()

	ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (c *Controller) DeleteCoin(ctx *gin.Context) {
	id := ctx.Param("id")

	fmt.Println("Controller.DeleteCoin", id)
	c.coinService.DeleteCoin()

	ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (c *Controller) PokeCoin(ctx *gin.Context) {
	id := ctx.Param("id")

	fmt.Println("Controller.PokeCoin", id)
	c.coinService.PokeCoin()

	ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
}
