package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"golang-rest-api-clean-architecture/pkg/external-services/authenticator"
	location_verifier "golang-rest-api-clean-architecture/pkg/external-services/location-verifier"
	"golang-rest-api-clean-architecture/pkg/internal"
	"net"
	"net/http"
	"runtime/debug"
	"strings"
)

func ResponseHeadersMiddleware(headers map[string]string) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			for k, v := range headers {
				w.Header().Set(k, v)
			}
			next.ServeHTTP(w, req)
		})
	}
}

func RecovererOnPanic(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rvr := recover(); rvr != nil && rvr != http.ErrAbortHandler {
				logrus.WithField("severity", "critical").
					Errorf("%s-%s", rvr, debug.Stack())
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func JwtMiddleware(authClient authenticator.AuthenticationClient) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			authHeader := req.Header.Get("Authorization")
			if len(authHeader) < 1 {
				internal.RespondWithJSON(w, http.StatusBadRequest, "authorization header not found")
				return
			}
			if !authClient.Authenticate(authHeader) {
				internal.RespondWithJSON(w, http.StatusUnauthorized, "authorization token is not valid")
				return
			}
			next.ServeHTTP(w, req)
		})
	}
}

func LocationVerificationMiddleware(verificationClient location_verifier.LocationVerificationClient) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			ipAddress, err := getIP(req)
			if err != nil {
				internal.RespondWithJSON(w, http.StatusInternalServerError, fmt.Sprintf("something went wrong in the server: %v", err))
				return
			}
			isIpValid, err := verificationClient.Verify(ipAddress)
			if err != nil {
				internal.RespondWithJSON(w, http.StatusInternalServerError, fmt.Sprintf("something went wrong in the server: %v", err))
				return
			}
			if !isIpValid {
				internal.RespondWithJSON(w, http.StatusUnauthorized, "service is not available in your area")
				return
			}
			next.ServeHTTP(w, req)
		})
	}
}

func getIP(r *http.Request) (string, error) {
	//Get IP from the X-REAL-IP header
	ip := r.Header.Get("X-REAL-IP")
	netIP := net.ParseIP(ip)
	if netIP != nil {
		return ip, nil
	}

	//Get IP from X-FORWARDED-FOR header
	ips := r.Header.Get("X-FORWARDED-FOR")
	splitIps := strings.Split(ips, ",")
	for _, ip := range splitIps {
		netIP := net.ParseIP(ip)
		if netIP != nil {
			return ip, nil
		}
	}

	//Get IP from RemoteAddr
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "", err
	}
	netIP = net.ParseIP(ip)
	if netIP != nil {
		return ip, nil
	}
	return "", fmt.Errorf("no valid ip found")
}
