package main

import (
	"dottime.dev/room/internal/database"
	"dottime.dev/room/internal/web"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/spf13/viper"
)

func main() {

	// Logger
	log := log.New(os.Stdout, "Room:", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

	// Viper
	viper := viper.GetViper()

	viper.SetDefault("Address", "0.0.0.0")
	viper.SetDefault("Port", 56245)
	viper.SetDefault("ReadTimeout", 3000)
	viper.SetDefault("WriteTimeout", 3000)

	type webConfig struct {
		Address      string
		Port         int
		ReadTimeout  int
		WriteTimeout int
	}

	var cfg webConfig

	err := viper.Unmarshal(&cfg)

	if err != nil {
		log.Fatal("fatal on parsing config.\nProbably you add new config without default a option.")
	}

	db, _ := database.Open(log)

	router := web.NewServer(log, db)

	server := http.Server{
		Addr:              cfg.Address + ":" + strconv.Itoa(cfg.Port),
		Handler:           router,
		ReadHeaderTimeout: time.Duration(cfg.ReadTimeout),
		WriteTimeout:      time.Duration(cfg.WriteTimeout),
	}

	err = server.ListenAndServe()

	if err != nil {
		fmt.Println(err)
	}

}
