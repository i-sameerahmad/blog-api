package controller

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/ImranZahoor/blog-api/internal/models"
	"github.com/ImranZahoor/blog-api/pkg/util"
	"github.com/gorilla/mux"
)

func (c *Controller) ListCategoryHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	categories, err := c.service.ListCategory(ctx)

	if err != nil {
		log.Printf("Controller::Category::ListCategoryHandler : %s", err.Error())
		if err == io.EOF {
			util.ToJSONResponse(w, models.Status{
				Error:      err.Error(),
				Message:    ErrorNotFound,
				StatusCode: http.StatusNotFound,
			})
			return
		}
		util.ToJSONResponse(w, models.Status{
			Error:      err.Error(),
			StatusCode: http.StatusBadRequest,
		})
		return
	}

	util.JsonResponse(w, http.StatusOK, categories)
}
func (c *Controller) GetCategoryByIDHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	id := vars["id"]
	intId, err := strconv.Atoi(id)

	if err != nil {
		util.ToJSONResponse(w, models.Status{
			Error:      err.Error(),
			Message:    ErrorInvalidSearchKey,
			StatusCode: http.StatusBadRequest,
		})
		return
	}
	category, err := c.service.GetCategoryByID(ctx, models.Uuid(intId))
	if err != nil {
		util.ToJSONResponse(w, models.Status{
			Error:      err.Error(),
			Message:    ErrorNotFound,
			StatusCode: http.StatusBadRequest,
		})
		return
	}
	util.ToJSONResponse(w, models.Status{
		Message:    MessageSuccess,
		StatusCode: http.StatusOK,
		Data:       category,
	})
}
func (c *Controller) CreateCategoryHandler(w http.ResponseWriter, r *http.Request) {

	var cateory models.Category
	jsonDecoder := json.NewDecoder(r.Body)
	err := jsonDecoder.Decode(&cateory)
	if err != nil {
		log.Printf("Controller::Category::CreateCategoryHandler: %s", err.Error())
		util.ToJSONResponse(w, models.Status{
			Error:      err.Error(),
			Message:    ErrorInvalidPayload,
			StatusCode: http.StatusBadRequest,
		})
		return
	}
	err = c.service.CreateCategory(r.Context(), cateory)

	if err != nil {
		log.Printf("Controller::Category::CreateCategoryHandler: %s", err.Error())
		util.ToJSONResponse(w, models.Status{
			Error:      err.Error(),
			Message:    MessageFailure,
			StatusCode: http.StatusBadRequest,
		})
		return
	}

	util.ToJSONResponse(w, models.Status{
		Message:    MessageSuccess,
		StatusCode: http.StatusCreated,
	})
}

// Delete Category handler
func (c *Controller) DeleteCategoryHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]
	intId, err := util.ToUUID(id)
	if err != nil {
		util.ToJSONResponse(w, models.Status{
			Error:      err.Error(),
			Message:    ErrorInvalidSearchKey,
			StatusCode: http.StatusBadRequest,
		})
		return
	}

	ctx := r.Context()
	err = c.service.DeleteCategory(ctx, intId)
	if err != nil {
		util.ToJSONResponse(w, models.Status{
			Error:      err.Error(),
			Message:    MessageFailure,
			StatusCode: http.StatusBadRequest,
		})
		return
	}

	util.ToJSONResponse(w, models.Status{
		Message:    MessageSuccess,
		StatusCode: http.StatusOK,
	})
}

// Category update handler
func (c *Controller) UpdateCategoryHandler(w http.ResponseWriter, r *http.Request) {

	var category models.Category
	jsonDecoder := json.NewDecoder(r.Body)
	err := jsonDecoder.Decode(&category)
	if err != nil {
		util.ToJSONResponse(w, models.Status{
			Error:      err.Error(),
			Message:    ErrorInvalidPayload,
			StatusCode: http.StatusBadRequest,
		})
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]
	intId, err := util.ToUUID(id)
	if err != nil {
		util.ToJSONResponse(w, models.Status{
			Error:      err.Error(),
			Message:    ErrorInvalidSearchKey,
			StatusCode: http.StatusBadRequest,
		})
	}
	ctx := r.Context()
	err = c.service.UpdateCategory(ctx, intId, category)
	if err != nil {
		util.ToJSONResponse(w, models.Status{
			Error:      err.Error(),
			Message:    MessageFailure,
			StatusCode: http.StatusBadRequest,
		})
		return
	}
	util.ToJSONResponse(w, models.Status{
		Message:    MessageSuccess,
		StatusCode: http.StatusOK,
	})
}
