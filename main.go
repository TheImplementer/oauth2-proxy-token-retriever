package main

import (
	"github.com/oauth2-proxy/oauth2-proxy/v7/pkg/apis/options"
	"github.com/oauth2-proxy/oauth2-proxy/v7/pkg/sessions"
	"net/http"
	"net/http/httptest"
	"os"
)

func main() {
	cookieValue := os.Args[0]

	sessionOptions := &options.SessionOptions{
		Type: options.RedisSessionStoreType,
		Redis: options.RedisStoreOptions{
			ConnectionURL: "redis://127.0.0.1:6379",
		},
	}

	cookieOptions := &options.Cookie{
		Name: "_oauth2_proxy",
	}

	sessionStore, _ := sessions.NewSessionStore(sessionOptions, cookieOptions)

	req := httptest.NewRequest("GET", "/test", nil)
	req.AddCookie(&http.Cookie{
		Name:  "_oauth2_proxy",
		Value: cookieValue,
	})

	_, err := sessionStore.Load(req)

	if err != nil {
		panic(err.Error())
	}
}
