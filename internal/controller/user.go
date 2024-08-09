package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/ImranZahoor/blog-api/internal/models"
	"github.com/ImranZahoor/blog-api/pkg/util"
	"github.com/gorilla/mux"
)

// List Users Controller
func (c *Controller) ListUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	users, err := c.service.ListUsers(ctx)
	if err != nil {
		util.ToJSONResponse(w, models.Status{
			Error:      ErrorNotFound,
			Message:    "User not found",
			StatusCode: http.StatusBadRequest,
		})
		return
	}

	util.ToJSONResponse(w, models.Status{
		Data:       users,
		Message:    MessageSuccess,
		StatusCode: http.StatusOK,
	})
}

// Find User By ID Users Controller
func (c *Controller) GetUserByIDHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	vars := mux.Vars(r)
	id := vars["id"]
	intID, err := strconv.Atoi(id)

	if err != nil {
		log.Println(err)
		util.ToJSONResponse(w, models.Status{
			Error:      err.Error(),
			Message:    "wrong user search key",
			StatusCode: http.StatusBadRequest,
		})
		return
	}

	user, err := c.service.GetUserByID(ctx, models.Uuid(intID))
	if err != nil {
		util.ToJSONResponse(w, models.Status{
			Error:      err.Error(),
			Message:    "user not found",
			StatusCode: http.StatusBadRequest,
		})
		return
	}

	util.ToJSONResponse(w, models.Status{
		Data:       user,
		StatusCode: http.StatusOK,
	})
}

// Create User Controller
func (c *Controller) CreateUserHandler(w http.ResponseWriter, r *http.Request) {

	var user models.User
	jsonDecoder := json.NewDecoder(r.Body)
	err := jsonDecoder.Decode(&user)
	if err != nil {
		util.ToJSONResponse(w, models.Status{
			Error:      err.Error(),
			Message:    ErrorInvalidPayload,
			StatusCode: http.StatusBadRequest,
		})
		return
	}
	ctx := r.Context()
	err = c.service.CreateUser(ctx, user)
	log.Println("Controller::CreateUserHandler")
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
		StatusCode: http.StatusCreated,
	})
}

// Delete User Controller
func (c *Controller) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	intID, err := util.ToUUID(id)
	if err != nil {
		util.ToJSONResponse(w, models.Status{
			Error:      err.Error(),
			Message:    ErrorInvalidSearchKey,
			StatusCode: http.StatusBadRequest,
		})
		return
	}

	ctx := r.Context()
	err = c.service.DeleteUser(ctx, intID)
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

// Update User Controller
func (c *Controller) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	jsonDecoder := json.NewDecoder(r.Body)
	err := jsonDecoder.Decode(&user)
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
	intID, err := util.ToUUID(id)
	if err != nil {
		util.ToJSONResponse(w, models.Status{
			Error:      err.Error(),
			Message:    ErrorInvalidSearchKey,
			StatusCode: http.StatusBadRequest,
		})
		return
	}

	ctx := r.Context()
	err = c.service.UpdateUser(ctx, intID, user)
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
