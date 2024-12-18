// Package oapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package oapi

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	uuid "github.com/gofrs/uuid/v5"
	"github.com/oapi-codegen/runtime"
)

const (
	AccessTokenAuthScopes = "AccessTokenAuth.Scopes"
	SkipAuthScopes        = "SkipAuth.Scopes"
)

// AccessTokenWithMe defines model for AccessTokenWithMe.
type AccessTokenWithMe struct {
	// AccessToken the access token
	AccessToken string `json:"access_token"`

	// Me the currently logged-in user info
	Me Me `json:"me"`
}

// CurrentPage defines model for CurrentPage.
type CurrentPage struct {
	// Url the url of the page currently being viewed
	Url string `json:"url"`
}

// GoogleLoginURL defines model for GoogleLoginURL.
type GoogleLoginURL struct {
	// RedirectUrl the Google's login page URL
	RedirectUrl string `json:"redirect_url"`
}

// Me defines model for Me.
type Me struct {
	Nickname string `json:"nickname"`

	// Uuid the UUID
	Uuid UUID `json:"uuid"`
}

// UUID the UUID
type UUID = uuid.UUID

// LoginWithGoogleJSONRequestBody defines body for LoginWithGoogle for application/json ContentType.
type LoginWithGoogleJSONRequestBody = CurrentPage

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /auth/google/callback)
	CallbackFromGoogle(w http.ResponseWriter, r *http.Request)

	// (POST /auth/google/link/{loginKey})
	LinkWithGoogle(w http.ResponseWriter, r *http.Request, loginKey string)

	// (POST /auth/google/login)
	LoginWithGoogle(w http.ResponseWriter, r *http.Request)

	// (POST /auth/logout)
	Logout(w http.ResponseWriter, r *http.Request)

	// (POST /auth/refresh)
	RefreshAccessToken(w http.ResponseWriter, r *http.Request)

	// (GET /me)
	GetMe(w http.ResponseWriter, r *http.Request)
}

// Unimplemented server implementation that returns http.StatusNotImplemented for each endpoint.

type Unimplemented struct{}

// (GET /auth/google/callback)
func (_ Unimplemented) CallbackFromGoogle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (POST /auth/google/link/{loginKey})
func (_ Unimplemented) LinkWithGoogle(w http.ResponseWriter, r *http.Request, loginKey string) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (POST /auth/google/login)
func (_ Unimplemented) LoginWithGoogle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (POST /auth/logout)
func (_ Unimplemented) Logout(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (POST /auth/refresh)
func (_ Unimplemented) RefreshAccessToken(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (GET /me)
func (_ Unimplemented) GetMe(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// CallbackFromGoogle operation middleware
func (siw *ServerInterfaceWrapper) CallbackFromGoogle(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	ctx = context.WithValue(ctx, SkipAuthScopes, []string{})

	r = r.WithContext(ctx)

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.CallbackFromGoogle(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// LinkWithGoogle operation middleware
func (siw *ServerInterfaceWrapper) LinkWithGoogle(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "loginKey" -------------
	var loginKey string

	err = runtime.BindStyledParameterWithOptions("simple", "loginKey", chi.URLParam(r, "loginKey"), &loginKey, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "loginKey", Err: err})
		return
	}

	ctx := r.Context()

	ctx = context.WithValue(ctx, SkipAuthScopes, []string{})

	r = r.WithContext(ctx)

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.LinkWithGoogle(w, r, loginKey)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// LoginWithGoogle operation middleware
func (siw *ServerInterfaceWrapper) LoginWithGoogle(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	ctx = context.WithValue(ctx, SkipAuthScopes, []string{})

	r = r.WithContext(ctx)

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.LoginWithGoogle(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// Logout operation middleware
func (siw *ServerInterfaceWrapper) Logout(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	ctx = context.WithValue(ctx, SkipAuthScopes, []string{})

	r = r.WithContext(ctx)

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.Logout(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// RefreshAccessToken operation middleware
func (siw *ServerInterfaceWrapper) RefreshAccessToken(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	ctx = context.WithValue(ctx, SkipAuthScopes, []string{})

	r = r.WithContext(ctx)

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.RefreshAccessToken(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// GetMe operation middleware
func (siw *ServerInterfaceWrapper) GetMe(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	ctx = context.WithValue(ctx, AccessTokenAuthScopes, []string{})

	r = r.WithContext(ctx)

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetMe(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/auth/google/callback", wrapper.CallbackFromGoogle)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/auth/google/link/{loginKey}", wrapper.LinkWithGoogle)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/auth/google/login", wrapper.LoginWithGoogle)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/auth/logout", wrapper.Logout)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/auth/refresh", wrapper.RefreshAccessToken)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/me", wrapper.GetMe)
	})

	return r
}
