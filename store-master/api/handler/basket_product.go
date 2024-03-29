package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"test/api/models"
)

// CreateBasketProduct godoc
// @Router       /basketProduct [POST]
// @Summary      Creates a new basket_product
// @Description  create a new basket_product
// @Tags         basketProduct
// @Accept       json
// @Produce      json
// @Param        basketProduct body models.CreateUser false "basketProduct"
// @Success      201  {object}  models.BasketProduct
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) CreateBasketProduct(c *gin.Context) {
	basketProduct := models.CreateBasketProduct{}

	if err := c.ShouldBindJSON(&basketProduct); err != nil {
		handleResponse(c, "error is while reading body", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storage.BasketProduct().Create(basketProduct)
	if err != nil {
		handleResponse(c, "error is while creating basket product", http.StatusInternalServerError, err)
		return
	}

	createdBasketProduct, err := h.storage.BasketProduct().GetByID(models.PrimaryKey{ID: id})
	if err != nil {
		handleResponse(c, "error is while getting by id", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusCreated, createdBasketProduct)
}

// GetBasketProduct godoc
// @Router       /basketProduct/{id} [GET]
// @Summary      Gets basketProduct
// @Description  get basketProduct by ID
// @Tags         basketProduct
// @Accept       json
// @Produce      json
// @Param        id path string true "basketProduct"
// @Success      200  {object}  models.BasketProduct
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetBasketProduct(c *gin.Context) {
	uid := c.Param("id")

	basketProduct, err := h.storage.BasketProduct().GetByID(models.PrimaryKey{ID: uid})
	if err != nil {
		handleResponse(c, "error is while getting by id", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusOK, basketProduct)
}

// GetBasketProductList godoc
// @Router       /basketProducts [GET]
// @Summary      Get basketProduct list
// @Description  get basketProduct list
// @Tags         basketProduct
// @Accept       json
// @Produce      json
// @Param        page query string false "page"
// @Param 		 limit query string false "limit"
// @Param 		 search query string false "search"
// @Success      200  {object}  models.BasketProductResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetBasketProductList(c *gin.Context) {
	var (
		page, limit int
		search      string
		err         error
	)

	pageStr := c.DefaultQuery("page", "1")
	page, err = strconv.Atoi(pageStr)
	if err != nil {
		handleResponse(c, "error is while converting page", http.StatusBadRequest, err.Error())
		return
	}

	limitStr := c.DefaultQuery("limit", "10")
	limit, err = strconv.Atoi(limitStr)
	if err != nil {
		handleResponse(c, "error is while converting page", http.StatusBadRequest, err.Error())
		return
	}

	search = c.Query("search")

	basketProducts, err := h.storage.BasketProduct().GetList(models.GetListRequest{
		Page:   page,
		Limit:  limit,
		Search: search,
	})

	handleResponse(c, "", http.StatusOK, basketProducts)
}

// UpdateBasketProduct godoc
// @Router       /basketProduct/{id} [PUT]
// @Summary      Update basketProduct
// @Description  update basketProduct
// @Tags         basketProduct
// @Accept       json
// @Produce      json
// @Param 		 id path string true "basketProduct_id"
// @Param        basketProduct body models.UpdateBasketProduct true "basketProduct"
// @Success      200  {object}  models.BasketProduct
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) UpdateBasketProduct(c *gin.Context) {
	basketProduct := models.UpdateBasketProduct{}
	uid := c.Param("id")

	if err := c.ShouldBindJSON(&basketProduct); err != nil {
		handleResponse(c, "error is while reading from body", http.StatusBadRequest, err.Error())
		return
	}

	basketProduct.ID = uid
	id, err := h.storage.BasketProduct().Update(basketProduct)
	if err != nil {
		handleResponse(c, "error is while updating basket", http.StatusInternalServerError, err.Error())
		return
	}

	updatedBasketProduct, err := h.storage.BasketProduct().GetByID(models.PrimaryKey{ID: id})
	if err != nil {
		handleResponse(c, "error is while getting by id", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusOK, updatedBasketProduct)
}

// DeleteBasketProduct godoc
// @Router       /basketProduct/{id} [DELETE]
// @Summary      Delete basketProduct
// @Description  delete basketProduct
// @Tags         basketProduct
// @Accept       json
// @Produce      json
// @Param 		 id path string true "basketProduct_id"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) DeleteBasketProduct(c *gin.Context) {
	uid := c.Param("id")

	if err := h.storage.BasketProduct().Delete(models.PrimaryKey{ID: uid}); err != nil {
		handleResponse(c, "error is while deleting", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusOK, "basket product deleted!")
}
