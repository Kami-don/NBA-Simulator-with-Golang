package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Furkan-Gulsen/NBA-Simulator-with-Golang/cmd/app/simulation"
	"github.com/Furkan-Gulsen/NBA-Simulator-with-Golang/config"
	"github.com/Furkan-Gulsen/NBA-Simulator-with-Golang/internal/domain/game"
	"github.com/Furkan-Gulsen/NBA-Simulator-with-Golang/internal/domain/player"
	playergame "github.com/Furkan-Gulsen/NBA-Simulator-with-Golang/internal/domain/player_game"
	"github.com/Furkan-Gulsen/NBA-Simulator-with-Golang/internal/domain/team"
	database "github.com/Furkan-Gulsen/NBA-Simulator-with-Golang/internal/infra/db"
	"github.com/Furkan-Gulsen/NBA-Simulator-with-Golang/internal/infra/shutdown"
	"github.com/gin-gonic/gin"
)

func loadConfiguration() *config.Config {
	productionMode := flag.Bool("dev", false, "dev mode")
	flag.Parse()

	isDevMode := false
	if *productionMode {
		isDevMode = true
	}

	if isDevMode {
		cfg, err := config.LocalLoadConfig()
		if err != nil {
			log.Fatalf("Failed to load configuration: %s", err)
		}
		return &cfg
	} else {
		cfg, err := config.LoadConfig()
		if err != nil {
			log.Fatalf("Failed to load configuration: %s", err)
		}

		return &cfg
	}
}

func initRouter(db *database.Database, dbName string) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())

	r.Use(gin.Recovery())
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	gameRepo := game.NewGameRepository(db.Collection(dbName, "games"))
	gameService := game.NewService(gameRepo)

	playerRepo := player.NewPlayerRepository(db.Collection(dbName, "players"))
	playerService := player.NewService(playerRepo)

	teamRepo := team.NewTeamRepository(db.Collection(dbName, "teams"))
	teamService := team.NewService(teamRepo)

	playerGameRepo := playergame.NewPlayerGameRepository(db.Collection(dbName, "player_games"))
	playerGameService := playergame.NewService(playerGameRepo)

	r.GET("/games", func(c *gin.Context) {
		games, err := gameService.GetAll(context.Background())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}
		c.JSON(http.StatusOK, games)
	})

	// simulation
	simulate := simulation.NewSimulate(context.Background(), db, gameService, playerService, teamService, playerGameService)
	go simulate.Run()

	return r
}

func initDatabase(cfg *config.Config) *database.Database {
	database, err := database.Connect(context.Background(), fmt.Sprintf("mongodb://%s:%s", cfg.MongoHostname, cfg.MongoPort))
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err)
	}
	return database
}

func startServer(cfg *config.Config, router *gin.Engine) {
	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", cfg.ServerAddress, cfg.ServerPort),
		Handler: router,
	}

	go shutdown.Graceful(server, 5*time.Second)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Failed to start server: %s", err)
	}
}

func main() {
	cfg := loadConfiguration()
	fmt.Println("cfg: ", cfg)
	db := initDatabase(cfg)
	defer db.Disconnect()

	r := initRouter(db, cfg.DatabaseName)
	startServer(cfg, r)
}
