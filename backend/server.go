package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	mRepo "github.com/ariefsn/ngobrol/app/message/repository/mongo"
	mSvc "github.com/ariefsn/ngobrol/app/message/service"
	rRepo "github.com/ariefsn/ngobrol/app/room/repository/mongo"
	rSvc "github.com/ariefsn/ngobrol/app/room/service"
	uRepo "github.com/ariefsn/ngobrol/app/user/repository/mongo"
	uSvc "github.com/ariefsn/ngobrol/app/user/service"
	"github.com/ariefsn/ngobrol/constants"
	"github.com/ariefsn/ngobrol/directives"
	"github.com/ariefsn/ngobrol/entities"
	"github.com/ariefsn/ngobrol/graph"
	"github.com/ariefsn/ngobrol/graph/resolvers"
	"github.com/ariefsn/ngobrol/helper"
	"github.com/ariefsn/ngobrol/logger"
	"github.com/ariefsn/ngobrol/middlewares"
	"github.com/ariefsn/ngobrol/validator"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/gorilla/websocket"
)

func init() {
	helper.InitEnv()
	logger.InitLogger()
	validator.InitValidator()
}

func main() {
	env := helper.GetEnv()

	router := chi.NewRouter()

	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		w.Header().Set("content-type", "application/json")
		res := entities.M{
			"status":  false,
			"message": "route not found",
		}
		body, _ := json.Marshal(res)
		w.Write(body)
	})

	router.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		w.Header().Set("content-type", "application/json")
		res := entities.M{
			"status":  false,
			"message": "method not allowed",
		}
		body, _ := json.Marshal(res)
		w.Write(body)
	})

	corsOpt := cors.Options{
		// AllowedOrigins: []string{
		// 	"http://localhost:5173",
		// 	"http://localhost:3000",
		// 	"http://localhost:3101",
		// 	"https://terkirim.localhost",
		// 	"https://terkirim.cloud",
		// },
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowOriginFunc: func(r *http.Request, origin string) bool {
			// fmt.Println("cors req", r.URL)
			// fmt.Println("cors org", origin)
			return true
		},
		AllowedHeaders:   []string{"Authorization", "Content-Type", "Accept", "Origin", string(constants.HeaderXTokenAccess), string(constants.HeaderXTokenRefresh), string(constants.HeaderXEmail), string(constants.HeaderXRoomCode)},
		AllowCredentials: true,
	}

	corsCfg := cors.New(corsOpt)

	// Setup db
	dbEnv := env.Mongo
	dbAddress := fmt.Sprintf("mongodb://%s:%s@%s:%s", dbEnv.User, dbEnv.Password, dbEnv.Host, dbEnv.Port)
	client, _ := helper.MongoClient(dbAddress)
	db := client.Database(dbEnv.Db)

	// Repositories
	userRepo := uRepo.NewMongoUserRepository(db)
	roomRepo := rRepo.NewMongoRoomRepository(db, userRepo)
	messageRepo := mRepo.NewMongoMessageRepository(db, roomRepo)

	// Services
	userService := uSvc.NewUserService(userRepo)
	roomService := rSvc.NewRoomService(roomRepo)
	messageService := mSvc.NewMessageService(messageRepo)

	router.Use(middlewares.Inject(*env, userService))
	router.Use(corsCfg.Handler)

	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("content-type", "application/json")
		res := entities.M{
			"status":  true,
			"message": "running",
		}
		body, _ := json.Marshal(res)
		w.Write(body)
	})

	// GraphQL Resolvers Config
	cfg := graph.Config{Resolvers: &resolvers.Resolver{
		UserService:    userService,
		RoomService:    roomService,
		MessageService: messageService,
	}}

	cfg.Directives.Protected = directives.Protected(userRepo)

	srv := handler.New(graph.NewExecutableSchema(cfg))
	srv.AroundOperations(func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
		oc := graphql.GetOperationContext(ctx)

		logger.Info(fmt.Sprintf("[%s]", strings.ToUpper(string(oc.Operation.Operation))), entities.M{
			"operationName": oc.OperationName,
			"variables":     oc.Variables,
		})

		return next(ctx)
	})
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})
	srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		InitFunc: func(ctx context.Context, initPayload transport.InitPayload) (context.Context, *transport.InitPayload, error) {
			fmt.Println(initPayload)
			return ctx, &initPayload, nil
			// return webSocketInit(ctx, initPayload)
		},
	})
	srv.Use(extension.Introspection{})

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://%s:%s/ for GraphQL playground", env.App.Host, env.App.Port)
	log.Fatal(http.ListenAndServe(":"+env.App.Port, router))
}
