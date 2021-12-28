package databases

import (
	"github.com/alhamsya/boilerplate-go/lib/helpers/database"
	"log"
)

func New(this *ServiceDB) *ServiceDB {
	dbCfg, ok := this.Cfg.Databases[this.Name]
	if !ok {
		log.Fatalf("[INIT] [DATABASE] [%s] config database does not exist", this.Name)
	}

	db, err := database.New(dbCfg, this.Driver)
	if err != nil {
		log.Fatalf("[INIT] [DATABASE] [%s] fail connect to database: %v", this.Name, err)
	}

	result := &ServiceDB{
		Cfg:    this.Cfg,
		Driver: this.Driver,
		DB:     db,
		Name:   this.Name,
	}
	log.Printf("[IGNORE] [DATABASE] initialize")
	return result
}