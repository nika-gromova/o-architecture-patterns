package main

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os/signal"
	"strings"
	"syscall"

	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v5"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/nika-gromova/o-architecture-patterns/project/internal/api"
	"github.com/swaggest/swgui/v5emb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	grpcPort  = 50051
	httpPort  = 8080
	adminPort = 8081
)

func main() {
	s := api.Service{}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50052))
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()
	reflection.Register(server)
	s.RegisterGRPC(server)

	go func() {
		if err = server.Serve(lis); err != nil {
			panic(err)
		}
	}()

	mux := runtime.NewServeMux()
	err = s.RegisterHTTP(context.Background(), mux)
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Use(addCORS, checkToken)
	r.HandleFunc("/api/*", func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.Replace(r.URL.Path, "/api", "", -1)
		mux.ServeHTTP(w, r)
	})
	httpServer := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf(":%d", 50053),
	}
	go func() {
		if err = httpServer.ListenAndServe(); err != nil && !errors.Is(http.ErrServerClosed, err) {
			panic(err)
		}
	}()

	muxAdmin := http.NewServeMux()
	muxAdmin.Handle("/swagger.json", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "/Users/nika/Projects/o-architecture-patterns/project/pkg/rules_v1/api.swagger.json")
	}))

	h := v5emb.NewHandler("Rules", "http://localhost:50054/swagger.json", "/docs")

	muxAdmin.Handle("/docs/*", h)

	adminServer := &http.Server{
		Handler: muxAdmin,
		Addr:    fmt.Sprintf(":%d", 50054),
	}
	go func() {
		if err = adminServer.ListenAndServe(); err != nil && !errors.Is(http.ErrServerClosed, err) {
			panic(err)
		}
	}()

	<-ctx.Done()
}

func addCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Разрешаем доступ с любых доменов
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Если это preflight-запрос (OPTIONS), сразу отдаем 200 OK
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Делаем вызов к следующему хэндлеру
		next.ServeHTTP(w, r)
	})
}

func checkToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		token := r.Header.Get("Authorization")

		claims := jwt.MapClaims{}

		SecretKey := "-----BEGIN CERTIFICATE-----\n" +
			`MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAtgdDwgoU0jSCT4E8jn8iGXNbe+yXiRVh8gqpixHOeJOdllJC+Cx/p7s5HvYe8Kis70EQFGMpLFes1hMwUJ4rhjGUvnyuMejmQbu7+lGRP57E1AOZoxZ20WPto+CDZaXhRfiG2yu37KtAS8aEMIYe7qhPa8NvuZWsv/2QS+wboGiU81jdMAf36KkbrBJ9Vy+mc6GZDjwJJSwivur73Gj7VuhVptM3pt43dgPM3HL3hiaTdmHEeewl1RwkdFpBMAvfyoPmKIZRuj8Wnast8FcfxrM2bvUBFPZe0hdI91FTiAENCpsrdkFjt0K+NepfQbqxc+8j6OSHeM5l7xqpJDDTPQIDAQAB` +
			"\n-----END CERTIFICATE-----"

		key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(SecretKey))
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		t, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			// return the public key that is used to validate the token.
			return key, nil
		})
		if err != nil {
			fmt.Printf("err: %s", err)
		}
		if !t.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		for key, val := range claims {
			fmt.Printf("Key: %v, value: %v\n", key, val)
		}
		next.ServeHTTP(w, r)
	})
}
