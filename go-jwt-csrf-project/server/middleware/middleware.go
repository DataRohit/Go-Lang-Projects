package middleware

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/datarohit/go-jwt-csrf-project/db"
	"github.com/datarohit/go-jwt-csrf-project/server/middleware/myJwt"
	"github.com/datarohit/go-jwt-csrf-project/server/templates"
	"github.com/justinas/alice"
)

func NewHandler() http.Handler {
	return alice.New(recoverHandler, authHandler).ThenFunc(logicHandler)
}

func recoverHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Panicf("recovered! Panic: %+v", err)
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func authHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/restricted", "/logout", "/deleteUser":
			AuthCookie, authErr := r.Cookie("AuthToken")
			if authErr == http.ErrNoCookie {
				nullifyTokenCookies(&w, r)
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			} else if authErr != nil {
				nullifyTokenCookies(&w, r)
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}

			RefreshCookie, refreshErr := r.Cookie("RefreshToken")
			if refreshErr == http.ErrNoCookie {
				nullifyTokenCookies(&w, r)
				http.Redirect(w, r, "/login", http.StatusFound)
				return
			} else if refreshErr != nil {
				nullifyTokenCookies(&w, r)
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}

			requestCsrfToken := grabCsrfFromReq(r)
			authTokenString, refreshTokenString, csrfSecret, err := myJwt.CheckAndRefreshTokens(AuthCookie.Value, RefreshCookie.Value, requestCsrfToken)
			if err != nil {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}

			w.Header().Set("Access-Control-Allow-Origin", "*")
			setAuthAndRefreshCookies(&w, authTokenString, refreshTokenString)
			w.Header().Set("X-CSRF-Token", csrfSecret)
		}

		next.ServeHTTP(w, r)
	})
}

func logicHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/restricted":
		csrfSecret := grabCsrfFromReq(r)
		templates.RenderTemplate(w, "restricted", &templates.RestrictedPage{CsrfSecret: csrfSecret, SecretMessage: "Hello DataRohit!"})

	case "/login":
		handleLogin(w, r)

	case "/register":
		handleRegister(w, r)

	case "/logout":
		nullifyTokenCookies(&w, r)
		http.Redirect(w, r, "/login", http.StatusFound)

	case "/deleteUser":
		deleteUser(w, r)

	default:
		w.WriteHeader(http.StatusOK)
	}
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		templates.RenderTemplate(w, "login", &templates.LoginPage{BAlertUser: false, AlertMsg: ""})
	case "POST":
		r.ParseForm()
		user, uuid, loginErr := db.LogUserIn(strings.Join(r.Form["username"], ""), strings.Join(r.Form["password"], ""))
		if loginErr != nil {
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			authTokenString, refreshTokenString, csrfSecret, err := myJwt.CreateNewTokens(uuid, user.Role)
			if err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
			setAuthAndRefreshCookies(&w, authTokenString, refreshTokenString)
			w.Header().Set("X-CSRF-Token", csrfSecret)
			w.WriteHeader(http.StatusOK)
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func handleRegister(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		templates.RenderTemplate(w, "register", &templates.RegisterPage{BAlertUser: false, AlertMsg: ""})
	case "POST":
		r.ParseForm()
		_, _, err := db.FetchUserByUsername(strings.Join(r.Form["username"], ""))
		if err == nil {
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			uuid, err := db.StoreUser(strings.Join(r.Form["username"], ""), strings.Join(r.Form["password"], ""), "user")
			if err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
			authTokenString, refreshTokenString, csrfSecret, err := myJwt.CreateNewTokens(uuid, "user")
			if err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
			setAuthAndRefreshCookies(&w, authTokenString, refreshTokenString)
			w.Header().Set("X-CSRF-Token", csrfSecret)
			w.WriteHeader(http.StatusOK)
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	AuthCookie, authErr := r.Cookie("AuthToken")
	if authErr == http.ErrNoCookie {
		nullifyTokenCookies(&w, r)
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	} else if authErr != nil {
		nullifyTokenCookies(&w, r)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	uuid, uuidErr := myJwt.GrabUUID(AuthCookie.Value)
	if uuidErr != nil {
		nullifyTokenCookies(&w, r)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	db.DeleteUser(uuid)
	nullifyTokenCookies(&w, r)
	http.Redirect(w, r, "/register", http.StatusFound)
}

func nullifyTokenCookies(w *http.ResponseWriter, r *http.Request) {
	authCookie := http.Cookie{Name: "AuthToken", Value: "", Expires: time.Now().Add(-1000 * time.Hour), HttpOnly: true}
	http.SetCookie(*w, &authCookie)

	refreshCookie := http.Cookie{Name: "RefreshToken", Value: "", Expires: time.Now().Add(-1000 * time.Hour), HttpOnly: true}
	http.SetCookie(*w, &refreshCookie)

	RefreshCookie, refreshErr := r.Cookie("RefreshToken")
	if refreshErr != nil && refreshErr != http.ErrNoCookie {
		http.Error(*w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	myJwt.RevokeRefreshToken(RefreshCookie.Value)
}

func setAuthAndRefreshCookies(w *http.ResponseWriter, authTokenString, refreshTokenString string) {
	authCookie := http.Cookie{Name: "AuthToken", Value: authTokenString, HttpOnly: true}
	http.SetCookie(*w, &authCookie)

	refreshCookie := http.Cookie{Name: "RefreshToken", Value: refreshTokenString, HttpOnly: true}
	http.SetCookie(*w, &refreshCookie)
}

func grabCsrfFromReq(r *http.Request) string {
	csrfFromForm := r.FormValue("X-CSRF-Token")
	if csrfFromForm != "" {
		return csrfFromForm
	}
	return r.Header.Get("X-CSRF-Token")
}
