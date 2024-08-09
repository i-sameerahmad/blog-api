package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ImranZahoor/blog-api/internal/models"
	"github.com/ImranZahoor/blog-api/pkg/util"
	"github.com/gorilla/mux"
)

// Article List Controller
func (c *Controller) ListArticleHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	articles, err := c.service.ListArticle(ctx)
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
		Data:       articles,
	})
}

// Article Find By ID Controller
func (c *Controller) GetArticleByIDHandler(w http.ResponseWriter, r *http.Request) {
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
	article, err := c.service.GetArticleByID(ctx, models.Uuid(intId))
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
		Data:       article,
	})
}

// Article Create Controller
func (c *Controller) CreateArticleHandler(w http.ResponseWriter, r *http.Request) {

	var article models.Article
	jsonDecoder := json.NewDecoder(r.Body)
	err := jsonDecoder.Decode(&article)
	if err != nil {
		util.ToJSONResponse(w, models.Status{
			Error:      err.Error(),
			Message:    ErrorInvalidPayload,
			StatusCode: http.StatusBadRequest,
		})
		return
	}
	err = c.service.CreateArticle(r.Context(), article)

	if err != nil {
		util.ToJSONResponse(w, models.Status{
			Error:      err.Error(),
			StatusCode: http.StatusBadRequest,
		})
		return
	}
	util.ToJSONResponse(w, models.Status{
		Message:    MessageSuccess,
		StatusCode: http.StatusCreated,
	})

}

// Article Delete Controller
func (c *Controller) DeleteArticleHandler(w http.ResponseWriter, r *http.Request) {
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
	err = c.service.DeleteArticle(ctx, intId)
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

// Article Update Controller
func (c *Controller) UpdateArticleHandler(w http.ResponseWriter, r *http.Request) {

	var article models.Article
	jsonDecoder := json.NewDecoder(r.Body)
	err := jsonDecoder.Decode(&article)
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
		return
	}
	ctx := r.Context()
	err = c.service.UpdateArticle(ctx, intId, article)
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
