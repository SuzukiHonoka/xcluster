package login

import (
	"fmt"
	"net/http"
	"xcluster/internal/api"
	"xcluster/internal/api/filter"
	userApi "xcluster/internal/api/user"
	"xcluster/internal/session"
	"xcluster/internal/user"
	"xcluster/pkg/utils"
)

// ServeUserLogin authenticates the requested user, assign session if authentication success
func ServeUserLogin(w http.ResponseWriter, r *http.Request) {
	if !filter.MatchMethod(w, r, http.MethodPost) {
		return
	}
	payload, ok := api.ParseJsonPayload(w, r)
	if !ok {
		return
	}
	email, ok := payload["email"].(string)
	password, ok2 := payload["password"].(string)
	if !ok || !ok2 {
		api.WriteError(w, http.StatusBadRequest, api.ErrPayloadInvalid)
		return
	}
	if filter.FieldsEmpty(w, email, password) {
		return
	}
	u, err := user.Email(email).GetUser()
	if err != nil {
		err = utils.ExtendError(api.ErrUserNotFound, err)
		api.WriteErrorAndLog(w, http.StatusForbidden, "user not found", err)
		return
	}
	ok, err = u.Password.Compare(password)
	if err != nil {
		err = utils.ExtendError(api.ErrUserWrongPassword, err)
		api.WriteErrorAndLog(w, http.StatusInternalServerError, "compare user password failed", err)
		return
	}
	if !ok {
		api.WriteErrorAndLog(w, http.StatusInternalServerError, "wrong password", api.ErrUserWrongPassword)
		return
	}
	// allocate session
	lease := session.NewLease(uint(u.ID), userApi.SessionDuration)
	userSession, err := session.NewSession(lease)
	if err != nil {
		api.WriteErrorAndLog(w, http.StatusInternalServerError, "create user session failed", err)
		return
	}
	if _, err = user.SaveSession(userSession); err != nil {
		api.WriteErrorAndLog(w, http.StatusInternalServerError, "save user session record failed", err)
		return
	}
	// set cookie
	http.SetCookie(w, &http.Cookie{
		Name:  "session",
		Value: string(userSession.ID),
		Path:  "/",
		// optional for binding domain
		//Domain:     "",
		Secure:   true,
		SameSite: http.SameSiteNoneMode, // todo: use http.SameSiteStrictMode in production build
		// let client know the session is expired
		Expires: userSession.Lease.ExpirationTime,
	})
	api.Write(w, api.NewResponse(http.StatusOK, "login success", u))
	msg := fmt.Sprintf("%s -> user: %s", userSession.ShortString(), u.Name)
	logger.Log(msg)
}
