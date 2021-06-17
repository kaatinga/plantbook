package userplants

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/proplants/plantbook/internal/api/handlers"
	apimiddleware "github.com/proplants/plantbook/internal/api/middleware"
	"github.com/proplants/plantbook/internal/api/models"
	"github.com/proplants/plantbook/internal/api/restapi/operations/userplant"
	"github.com/proplants/plantbook/pkg/logging"
	"github.com/proplants/plantbook/pkg/token"
)

type deleteUserPlantImpl struct {
	storage RepoInterface
	tm      token.Manager
}

// NewDeleteUserPlantHandler builder for userplant.DeleteUserPlantHandler interface implementation.
func NewDeleteUserPlantHandler(storage RepoInterface, tm token.Manager) userplant.DeleteUserPlantHandler {
	return &deleteUserPlantImpl{storage: storage, tm: tm}
}

// Handle implementation of the refplant.userplant.DeleteUserPlantHandler interface.
func (impl *deleteUserPlantImpl) Handle(params userplant.DeleteUserPlantParams) middleware.Responder {
	log := logging.FromContext(params.HTTPRequest.Context())
	cookie, err := params.HTTPRequest.Cookie(apimiddleware.JWTCookieName)
	if err != nil {
		log.Errorf("get cookie %s error, %s", apimiddleware.JWTCookieName, err)
		return userplant.NewDeleteUserPlantDefault(http.StatusUnauthorized).
			WithPayload(&models.ErrorResponse{Message: "not token cookie"}).
			WithXRequestID(apimiddleware.GetRequestID(params.HTTPRequest))
	}
	if cookie == nil {
		return userplant.NewDeleteUserPlantDefault(http.StatusUnauthorized).
			WithPayload(&models.ErrorResponse{Message: "empty token cookie"}).
			WithXRequestID(apimiddleware.GetRequestID(params.HTTPRequest))
	}
	ok, err := impl.tm.Check(params.HTTPRequest.Context(), cookie.Value)
	if err != nil {
		log.Errorf("check token %s error, %s", cookie.Value, err)
		return userplant.NewDeleteUserPlantDefault(http.StatusUnauthorized).
			WithPayload(&models.ErrorResponse{Message: "check token error"}).
			WithXRequestID(apimiddleware.GetRequestID(params.HTTPRequest))
	}
	if !ok {
		return userplant.NewDeleteUserPlantDefault(http.StatusUnauthorized).
			WithPayload(&models.ErrorResponse{Message: "token expired"}).
			WithXRequestID(apimiddleware.GetRequestID(params.HTTPRequest))
	}

	uid, _, userRoleID, err := impl.tm.FindUserData(cookie.Value)
	if err != nil {
		log.Errorf("get user attributes from token %s error, %s", cookie.Value, err)
		return userplant.NewDeleteUserPlantDefault(http.StatusForbidden).
			WithPayload(&models.ErrorResponse{Message: "check permission error"}).
			WithXRequestID(apimiddleware.GetRequestID(params.HTTPRequest))
	}
	existingUserPlant, err := impl.storage.GetUserPlantByID(params.HTTPRequest.Context(), params.UserplantID)
	if err != nil {
		log.Infof("storage.GetUserPlantByID with id=%d error, %s", params.UserplantID, err)
		return userplant.NewDeleteUserPlantDefault(http.StatusInternalServerError).
			WithPayload(&models.ErrorResponse{Message: "db error happen"})
	}
	if existingUserPlant == nil {
		log.Infof("storage.GetUserPlantByID with id=%d not found", params.UserplantID)
		return userplant.NewDeleteUserPlantDefault(http.StatusNotFound).
			WithPayload(&models.ErrorResponse{Message: "user's plant not found"})
	}

	isAdmin := userRoleID == handlers.UserRoleAdmin
	isOwner := existingUserPlant.UserID == uid
	if !(isAdmin || isOwner) {
		log.Errorf("userID=%d, not owner and not admin try delete", uid)
		return userplant.NewDeleteUserPlantDefault(http.StatusForbidden).
			WithPayload(&models.ErrorResponse{Message: "forbidden"}).
			WithXRequestID(apimiddleware.GetRequestID(params.HTTPRequest))
	}

	err = impl.storage.DeleteUserPlant(params.HTTPRequest.Context(), params.UserplantID)
	if err != nil {
		log.Errorf("Handle DeleteUserPlant error: %s", err)
		return userplant.NewDeleteUserPlantDefault(http.StatusInternalServerError).
			WithPayload(&models.ErrorResponse{Message: "db error"}).
			WithXRequestID(apimiddleware.GetRequestID(params.HTTPRequest))
	}
	return userplant.NewDeleteUserPlantOK().WithPayload(&models.Response{Message: "User's plant deleted"}).
		WithXRequestID(apimiddleware.GetRequestID(params.HTTPRequest))
}
