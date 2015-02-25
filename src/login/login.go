package login

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/sessions"
	"net/http"
	"golang.org/x/crypto/bcrypt"
)

var (
	whiteList []string
	loginUrl string = "/login"
)

func Setup(goodPaths []string) martini.Handler {
	whiteList := goodPaths

	return func(res http.ResponseWriter, r *http.Request, c martini.Context, session sessions.Session) {
		v := session.Get("user_id")

		if v == nil {
			for _, value := range whiteList {
				if value == r.URL.Path {
					return
				}
			}
			http.Redirect(res, r, loginUrl, http.StatusFound)
		}
	}
}

func EncryptPassword(password string) string {
	passwordBytes := []byte(password)

	hashedPassword, _ := bcrypt.GenerateFromPassword(passwordBytes, 10)
	return string(hashedPassword)
}

func ComparePassword(password string, hash string) bool {
	passwordBytes := []byte(password)
	hashBytes := []byte(hash)

	err := bcrypt.CompareHashAndPassword(hashBytes, passwordBytes)
	if (err == nil) {
		return true;
	} else {
		return false;
	}
}
