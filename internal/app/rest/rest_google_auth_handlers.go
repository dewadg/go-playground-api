package rest

import (
	"context"
	"encoding/json"
	"math/rand"
	"net/http"
	"os"
	"time"

	"encoding/base64"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var googleOauthConfig = oauth2.Config{
	RedirectURL:  os.Getenv("APP_HOST") + "/v1/auth/google/callback",
	ClientID:     os.Getenv("GOOGLE_OAUTH_CLIENT_ID"),
	ClientSecret: os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET"),
	Scopes: []string{
		"https://www.googleapis.com/auth/userinfo.email",
		"https://www.googleapis.com/auth/userinfo.profile",
	},
	Endpoint: google.Endpoint,
}

func handleGoogleLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		oauthState, cookie := generateOauthState()

		redirectTo := googleOauthConfig.AuthCodeURL(oauthState)

		http.SetCookie(w, cookie)
		http.Redirect(w, r, redirectTo, http.StatusTemporaryRedirect)
	}
}

func generateOauthState() (string, *http.Cookie) {
	var expiration = time.Now().Add(20 * time.Minute)

	b := make([]byte, 16)
	rand.Read(b)

	state := base64.URLEncoding.EncodeToString(b)
	cookie := &http.Cookie{
		Name:    "oauthstate",
		Value:   state,
		Expires: expiration,
	}

	return state, cookie
}

type AuthTokenGenerator func(ctx context.Context, email string) (string, error)

func handleGoogleCallback(atGen AuthTokenGenerator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		oauthState, err := r.Cookie("oauthstate")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err = r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if r.FormValue("state") != oauthState.Value {
			http.Error(w, "invalid oauth2 state", http.StatusBadRequest)
			return
		}

		ctx := r.Context()
		userInfo, err := getUserFromGoogle(ctx, r.FormValue("code"))
		if err != nil {
			http.Error(w, "invalid oauth2 state", http.StatusInternalServerError)
			return
		}

		accessToken, err := atGen(ctx, userInfo.Email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		redirectTo := os.Getenv("WEB_HOST") + "/auth/success?access_token=" + accessToken
		http.Redirect(w, r, redirectTo, http.StatusTemporaryRedirect)
	}
}

type googleUserInfoResponse struct {
	Email string `json:"email"`
}

func getUserFromGoogle(ctx context.Context, authCode string) (googleUserInfoResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	token, err := googleOauthConfig.Exchange(ctx, authCode)
	if err != nil {
		return googleUserInfoResponse{}, err
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		"https://www.googleapis.com/oauth2/v2/userinfo?access_token="+token.AccessToken,
		nil,
	)
	if err != nil {
		return googleUserInfoResponse{}, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return googleUserInfoResponse{}, err
	}
	defer resp.Body.Close()

	var respBody googleUserInfoResponse
	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		return googleUserInfoResponse{}, err
	}

	return respBody, nil
}
