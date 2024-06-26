// Code generated by goa v3.11.0, DO NOT EDIT.
//
// relayAccount HTTP server
//
// Command:
// $ goa gen admin-panel/design

package server

import (
	relayaccount "admin-panel/gen/relay_account"
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Server lists the relayAccount service endpoint HTTP handlers.
type Server struct {
	Mounts          []*MountPoint
	ListAccount     http.Handler
	RegisterAccount http.Handler
	DeleteAccount   http.Handler
}

// MountPoint holds information about the mounted endpoints.
type MountPoint struct {
	// Method is the name of the service method served by the mounted HTTP handler.
	Method string
	// Verb is the HTTP method used to match requests to the mounted handler.
	Verb string
	// Pattern is the HTTP request path pattern used to match requests to the
	// mounted handler.
	Pattern string
}

// New instantiates HTTP handlers for all the relayAccount service endpoints
// using the provided encoder and decoder. The handlers are mounted on the
// given mux using the HTTP verb and path defined in the design. errhandler is
// called whenever a response fails to be encoded. formatter is used to format
// errors returned by the service methods prior to encoding. Both errhandler
// and formatter are optional and can be nil.
func New(
	e *relayaccount.Endpoints,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"ListAccount", "GET", "/lpnode/lpnode_admin_panel/relayAccount/list"},
			{"RegisterAccount", "POST", "/lpnode/lpnode_admin_panel/relayAccount/register"},
			{"DeleteAccount", "POST", "/lpnode/lpnode_admin_panel/relayAccount/delete"},
		},
		ListAccount:     NewListAccountHandler(e.ListAccount, mux, decoder, encoder, errhandler, formatter),
		RegisterAccount: NewRegisterAccountHandler(e.RegisterAccount, mux, decoder, encoder, errhandler, formatter),
		DeleteAccount:   NewDeleteAccountHandler(e.DeleteAccount, mux, decoder, encoder, errhandler, formatter),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "relayAccount" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.ListAccount = m(s.ListAccount)
	s.RegisterAccount = m(s.RegisterAccount)
	s.DeleteAccount = m(s.DeleteAccount)
}

// MethodNames returns the methods served.
func (s *Server) MethodNames() []string { return relayaccount.MethodNames[:] }

// Mount configures the mux to serve the relayAccount endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountListAccountHandler(mux, h.ListAccount)
	MountRegisterAccountHandler(mux, h.RegisterAccount)
	MountDeleteAccountHandler(mux, h.DeleteAccount)
}

// Mount configures the mux to serve the relayAccount endpoints.
func (s *Server) Mount(mux goahttp.Muxer) {
	Mount(mux, s)
}

// MountListAccountHandler configures the mux to serve the "relayAccount"
// service "listAccount" endpoint.
func MountListAccountHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/lpnode/lpnode_admin_panel/relayAccount/list", f)
}

// NewListAccountHandler creates a HTTP handler which loads the HTTP request
// and calls the "relayAccount" service "listAccount" endpoint.
func NewListAccountHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		encodeResponse = EncodeListAccountResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "listAccount")
		ctx = context.WithValue(ctx, goa.ServiceKey, "relayAccount")
		var err error
		res, err := endpoint(ctx, nil)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountRegisterAccountHandler configures the mux to serve the "relayAccount"
// service "registerAccount" endpoint.
func MountRegisterAccountHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/lpnode/lpnode_admin_panel/relayAccount/register", f)
}

// NewRegisterAccountHandler creates a HTTP handler which loads the HTTP
// request and calls the "relayAccount" service "registerAccount" endpoint.
func NewRegisterAccountHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeRegisterAccountRequest(mux, decoder)
		encodeResponse = EncodeRegisterAccountResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "registerAccount")
		ctx = context.WithValue(ctx, goa.ServiceKey, "relayAccount")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountDeleteAccountHandler configures the mux to serve the "relayAccount"
// service "deleteAccount" endpoint.
func MountDeleteAccountHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/lpnode/lpnode_admin_panel/relayAccount/delete", f)
}

// NewDeleteAccountHandler creates a HTTP handler which loads the HTTP request
// and calls the "relayAccount" service "deleteAccount" endpoint.
func NewDeleteAccountHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeDeleteAccountRequest(mux, decoder)
		encodeResponse = EncodeDeleteAccountResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "deleteAccount")
		ctx = context.WithValue(ctx, goa.ServiceKey, "relayAccount")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}
