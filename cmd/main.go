package main

import (
	"github.com/rs/zerolog/log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stphanvebrian/go-simple-chat/usecase/repository"
)

func main() {
	db, err := repository.NewDatabase()
	if err != nil {
		log.Fatal().Err(err)
	}

	// do migration
	repository.RunMigration(db)

	log.Info().Msg("running application...")

	err = startServer(db)
	if err != nil {
		log.Fatal().Err(err)
	}
}
