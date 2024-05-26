package http

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/http"

	"github.com/elazarl/goproxy"
	"golang.org/x/oauth2"

	"github.com/aplulu/iapproxy/internal/config"
	"github.com/aplulu/iapproxy/internal/util"
)

var server *http.Server

// StartServer starts the server
func StartServer(log *slog.Logger) error {
	ctx := context.Background()

	var tokenSource oauth2.TokenSource
	if config.IAPClientID() != "" {
		log.Info("IAP client ID is set,", slog.String("client_id", config.IAPClientID()))

		audience := config.IAPClientID()
		var err error
		tokenSource, err = util.GetTokenSource(ctx, audience)
		if err != nil {
			return fmt.Errorf("failed to get token source: %w", err)
		}

		for _, pattern := range config.URLPatterns() {
			log.Info("URL pattern", slog.String("pattern", pattern.String()))
		}
	} else {
		log.Info("IAP client ID is not set")
	}

	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = config.ProxyVerbose()
	proxy.Logger = util.NewProxyLogger(log)

	// KeepHeader is required to keep the Proxy-Authorization header
	proxy.KeepHeader = true
	proxy.OnRequest().HandleConnect(goproxy.AlwaysMitm)
	proxy.OnRequest().DoFunc(func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
		iapProtected := false
		if tokenSource != nil && isURLPatternMatch(r.URL.String(), config.URLPatterns()) {
			token, err := tokenSource.Token()
			if err != nil {
				log.Error(fmt.Sprintf("failed to get token: %v", err))
				return r, nil
			}

			r.Header.Set("Proxy-Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))
			iapProtected = true
		}

		log.Info(
			"Request received",
			slog.String("method", r.Method),
			slog.String("url", r.URL.String()),
			slog.Bool("iap_protected", iapProtected),
		)

		return r, nil
	})

	server = &http.Server{
		Addr:    net.JoinHostPort(config.Listen(), config.Port()),
		Handler: proxy,
	}

	listenHost := config.Listen()
	if listenHost == "" {
		listenHost = "localhost"
	}
	log.Info(fmt.Sprintf("Server started at http://%s:%s", listenHost, config.Port()))
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

// StopServer stops the server
func StopServer(ctx context.Context) error {
	return server.Shutdown(ctx)
}
