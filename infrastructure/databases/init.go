package databases

import (
	"log"

	"github.com/alhamsya/boilerplate-go/lib/helpers/database"
	"github.com/rubenv/sql-migrate"
)

func New(this *ServiceDB) *ServiceDB {
	dbCfg, ok := this.Cfg.Databases[this.Name]
	if !ok {
		log.Fatalf("[INIT] [DATABASE] [%s] config database does not exist", this.Name)
	}

	dbStore, err := database.New(dbCfg, this.Driver)
	if err != nil {
		log.Fatalf("[INIT] [DATABASE] [%s] fail connect to database: %v", this.Name, err)
	}

	numberMigration, err := migrate.Exec(
		dbStore.Master.DB,
		this.Driver,
		&migrate.FileMigrationSource{
			Dir: this.Cfg.Databases[this.Name].DirSchema,
		},
		migrate.Up,
	)
	if err != nil {
		log.Fatalf("[INIT] [DATABASE MIGRATIONS] [%s] fail %d migration database: %v", this.Name, numberMigration, err)
	} else {
		log.Printf("[IGNORE] [DATABASE MIGRATIONS] initialize")
	}

	result := &ServiceDB{
		Cfg:    this.Cfg,
		Driver: this.Driver,
		DB:     dbStore,
		Name:   this.Name,
	}
	log.Printf("[IGNORE] [DATABASE] initialize")
	return result
}
