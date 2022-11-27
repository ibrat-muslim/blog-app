package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ibrat-muslim/blog-app/api/models"
	"github.com/ibrat-muslim/blog-app/storage/repo"
)

// @Router /categories [post]
// @Summary Create a category
// @Description Create a category
// @Tags category
// @Accept json
// @Produce json
// @Param category body models.CreateCategoryRequest true "Category"
// @Success 201 {object} models.Category
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) CreateCategory(ctx *gin.Context) {

	var req models.CreateCategoryRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	resp, err := h.storage.Category().Create(&repo.Category{
		Title: req.Title,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, models.Category{
		ID:        resp.ID,
		Title:     resp.Title,
		CreatedAt: resp.CreatedAt,
	})
}

// @Router /categories/{id} [get]
// @Summary Get a category by id
// @Description Get a category by id
// @Tags category
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} models.Category
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) GetCategory(ctx *gin.Context) {

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	resp, err := h.storage.Category().Get(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Category{
		ID:        resp.ID,
		Title:     resp.Title,
		CreatedAt: resp.CreatedAt,
	})
}
