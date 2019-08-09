package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/graphql-go/handler"
	"github.com/tanphamhaiduong/go/delta/internal/database"
	"github.com/tanphamhaiduong/go/delta/internal/grpcservice"
	"github.com/tanphamhaiduong/go/delta/internal/modules"
	"github.com/tanphamhaiduong/go/delta/internal/modules/company"
	"github.com/tanphamhaiduong/go/delta/internal/modules/permission"
	"github.com/tanphamhaiduong/go/delta/internal/modules/role"
	"github.com/tanphamhaiduong/go/delta/internal/modules/rolepermission"
	"github.com/tanphamhaiduong/go/delta/internal/modules/user"
	"google.golang.org/grpc"
)

// Config ...
type Config struct {
	DatabasePath string
	Port         string
	GRPCPort     string
	DBDriver     string
	WriteTimeout time.Duration
	ReadTimeout  time.Duration
	IdleTimeout  time.Duration
	GracefulTime time.Duration
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		DatabasePath: "root:root@tcp(127.0.0.1:3306)/crm",
		Port:         "8000",
		GRPCPort:     ":9000",
		DBDriver:     "mysql",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		GracefulTime: time.Minute * 1,
	}
}

// ConnectDatabase ...
func ConnectDatabase(config *Config) (database.IDB, error) {
	return database.NewSQLlib(config.DBDriver, config.DatabasePath)
}

// Server ...
type Server struct {
	config         *Config
	graphQLHandler *handler.Handler
	resolvers      modules.Resolver
	handlers       modules.Handler
}

// NewServer ...
func NewServer(config *Config, resolvers modules.Resolver, handlers modules.Handler) *Server {
	return &Server{
		config:    config,
		resolvers: resolvers,
		handlers:  handlers,
	}
}

// Handler ...
func (s *Server) Handler() http.Handler {
	r := mux.NewRouter()
	r.Use(withTraceID)
	r.Use(withLogging)
	r.Use(withTimeout)
	r.Use(withAuth)
	schema, err := modules.MakeSchema(s.resolvers)
	if err != nil {
		log.Fatal(err)
	}
	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})
	// Add your routes as needed
	r.HandleFunc("/graphql", h.ServeHTTP).Methods("POST")
	r.HandleFunc("/graphql", h.ServeHTTP).Methods("GET")

	return r
}

// Run ...
func (s *Server) Run() {
	var (
		wait time.Duration
	)
	// Init HTTP Server
	flag.DurationVar(&wait, "graceful-timeout", s.config.GracefulTime, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()
	srv := &http.Server{
		Addr: ":" + s.config.Port,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: s.config.WriteTimeout,
		ReadTimeout:  s.config.ReadTimeout,
		IdleTimeout:  s.config.IdleTimeout,
		Handler:      s.Handler(), // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		log.Println(fmt.Sprintf("Listening on %v", s.config.Port))
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	// Init GRPC Server
	lis, err := net.Listen("tcp", s.config.GRPCPort)
	// creds, err := credentials.NewServerTLSFromFile("./cert/cert.pem", "./cert/key.pem")
	if err != nil {
		log.Fatal(err)
	}
	// opts := []grpc.ServerOption{grpc.Creds(creds)}
	opts := []grpc.ServerOption{}
	grpc := grpc.NewServer(opts...)
	grpcService := grpcservice.NewGrpcService(s.handlers)
	grpcService.Register(grpc)

	go func() {
		log.Println(fmt.Sprintf("Listening on %v", s.config.GRPCPort))
		if err := grpc.Serve(lis); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("graceful shutting down")
	os.Exit(0)
}

func main() {
	config := NewConfig()
	db, err := ConnectDatabase(config)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// Init All Dependencies
	// Company
	companyRepository := company.NewRepository(db)
	companyHandler := company.NewHandler(companyRepository)
	companyResolver := company.NewResolver(companyHandler)
	// Permission
	permissionRepository := permission.NewRepository(db)
	permissionHandler := permission.NewHandler(permissionRepository)
	permissionResolver := permission.NewResolver(permissionHandler)
	// Role
	roleRepository := role.NewRepository(db)
	roleHandler := role.NewHandler(roleRepository)
	roleResolver := role.NewResolver(roleHandler)
	// RolePermission
	rolepermissionRepository := rolepermission.NewRepository(db)
	rolepermissionHandler := rolepermission.NewHandler(rolepermissionRepository)
	rolepermissionResolver := rolepermission.NewResolver(rolepermissionHandler)
	// User
	userRepository := user.NewRepository(db)
	userHandler := user.NewHandler(userRepository)
	userResolver := user.NewResolver(userHandler)
	// Init Resolvers
	resolvers := modules.NewResolver(
		companyResolver,
		permissionResolver,
		roleResolver,
		rolepermissionResolver,
		userResolver,
	)

	handlers := modules.NewHandler(
		companyHandler,
		permissionHandler,
		roleHandler,
		rolepermissionHandler,
		userHandler,
	)

	server := NewServer(config, resolvers, handlers)

	server.Run()
}
