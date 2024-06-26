// Code generated by goa v3.11.0, DO NOT EDIT.
//
// dexWallet HTTP server
//
// Command:
// $ goa gen admin-panel/design

package server

import (
	dexwallet "admin-panel/gen/dex_wallet"
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Server lists the dexWallet service endpoint HTTP handlers.
type Server struct {
	Mounts          []*MountPoint
	ListDexWallet   http.Handler
	CreateDexWallet http.Handler
	DeleteDexWallet http.Handler
	VaultList       http.Handler
	UpdateLpWallet  http.Handler
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

// New instantiates HTTP handlers for all the dexWallet service endpoints using
// the provided encoder and decoder. The handlers are mounted on the given mux
// using the HTTP verb and path defined in the design. errhandler is called
// whenever a response fails to be encoded. formatter is used to format errors
// returned by the service methods prior to encoding. Both errhandler and
// formatter are optional and can be nil.
func New(
	e *dexwallet.Endpoints,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"ListDexWallet", "GET", "/lpnode/lpnode_admin_panel/dexWallet/list"},
			{"CreateDexWallet", "POST", "/lpnode/lpnode_admin_panel/dexWallet/create"},
			{"DeleteDexWallet", "POST", "/lpnode/lpnode_admin_panel/dexWallet/delete"},
			{"VaultList", "POST", "/lpnode/lpnode_admin_panel/dexWallet/vaultList"},
			{"UpdateLpWallet", "POST", "/lpnode/lpnode_admin_panel/dexWallet/updateLpWallet"},
		},
		ListDexWallet:   NewListDexWalletHandler(e.ListDexWallet, mux, decoder, encoder, errhandler, formatter),
		CreateDexWallet: NewCreateDexWalletHandler(e.CreateDexWallet, mux, decoder, encoder, errhandler, formatter),
		DeleteDexWallet: NewDeleteDexWalletHandler(e.DeleteDexWallet, mux, decoder, encoder, errhandler, formatter),
		VaultList:       NewVaultListHandler(e.VaultList, mux, decoder, encoder, errhandler, formatter),
		UpdateLpWallet:  NewUpdateLpWalletHandler(e.UpdateLpWallet, mux, decoder, encoder, errhandler, formatter),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "dexWallet" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.ListDexWallet = m(s.ListDexWallet)
	s.CreateDexWallet = m(s.CreateDexWallet)
	s.DeleteDexWallet = m(s.DeleteDexWallet)
	s.VaultList = m(s.VaultList)
	s.UpdateLpWallet = m(s.UpdateLpWallet)
}

// MethodNames returns the methods served.
func (s *Server) MethodNames() []string { return dexwallet.MethodNames[:] }

// Mount configures the mux to serve the dexWallet endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountListDexWalletHandler(mux, h.ListDexWallet)
	MountCreateDexWalletHandler(mux, h.CreateDexWallet)
	MountDeleteDexWalletHandler(mux, h.DeleteDexWallet)
	MountVaultListHandler(mux, h.VaultList)
	MountUpdateLpWalletHandler(mux, h.UpdateLpWallet)
}

// Mount configures the mux to serve the dexWallet endpoints.
func (s *Server) Mount(mux goahttp.Muxer) {
	Mount(mux, s)
}

// MountListDexWalletHandler configures the mux to serve the "dexWallet"
// service "listDexWallet" endpoint.
func MountListDexWalletHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/lpnode/lpnode_admin_panel/dexWallet/list", f)
}

// NewListDexWalletHandler creates a HTTP handler which loads the HTTP request
// and calls the "dexWallet" service "listDexWallet" endpoint.
func NewListDexWalletHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		encodeResponse = EncodeListDexWalletResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "listDexWallet")
		ctx = context.WithValue(ctx, goa.ServiceKey, "dexWallet")
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

// MountCreateDexWalletHandler configures the mux to serve the "dexWallet"
// service "createDexWallet" endpoint.
func MountCreateDexWalletHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/lpnode/lpnode_admin_panel/dexWallet/create", f)
}

// NewCreateDexWalletHandler creates a HTTP handler which loads the HTTP
// request and calls the "dexWallet" service "createDexWallet" endpoint.
func NewCreateDexWalletHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeCreateDexWalletRequest(mux, decoder)
		encodeResponse = EncodeCreateDexWalletResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "createDexWallet")
		ctx = context.WithValue(ctx, goa.ServiceKey, "dexWallet")
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

// MountDeleteDexWalletHandler configures the mux to serve the "dexWallet"
// service "deleteDexWallet" endpoint.
func MountDeleteDexWalletHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/lpnode/lpnode_admin_panel/dexWallet/delete", f)
}

// NewDeleteDexWalletHandler creates a HTTP handler which loads the HTTP
// request and calls the "dexWallet" service "deleteDexWallet" endpoint.
func NewDeleteDexWalletHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeDeleteDexWalletRequest(mux, decoder)
		encodeResponse = EncodeDeleteDexWalletResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "deleteDexWallet")
		ctx = context.WithValue(ctx, goa.ServiceKey, "dexWallet")
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

// MountVaultListHandler configures the mux to serve the "dexWallet" service
// "vaultList" endpoint.
func MountVaultListHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/lpnode/lpnode_admin_panel/dexWallet/vaultList", f)
}

// NewVaultListHandler creates a HTTP handler which loads the HTTP request and
// calls the "dexWallet" service "vaultList" endpoint.
func NewVaultListHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		encodeResponse = EncodeVaultListResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "vaultList")
		ctx = context.WithValue(ctx, goa.ServiceKey, "dexWallet")
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

// MountUpdateLpWalletHandler configures the mux to serve the "dexWallet"
// service "updateLpWallet" endpoint.
func MountUpdateLpWalletHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/lpnode/lpnode_admin_panel/dexWallet/updateLpWallet", f)
}

// NewUpdateLpWalletHandler creates a HTTP handler which loads the HTTP request
// and calls the "dexWallet" service "updateLpWallet" endpoint.
func NewUpdateLpWalletHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		encodeResponse = EncodeUpdateLpWalletResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "updateLpWallet")
		ctx = context.WithValue(ctx, goa.ServiceKey, "dexWallet")
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
