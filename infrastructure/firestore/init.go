package firestore

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
)

func New(this *ServiceFirestore) *ServiceFirestore {
	this.Clients = make(map[string]*firestore.Client)
	for name, data := range this.Cfg.Firestore {
		client, err := firestore.NewClient(context.Background(), data.ProjectID)
		if err != nil {
			log.Fatalf("[INIT] [FIRESTORE] [%s] config firestore does not exist", name)
		}

		this.Clients[name] = client
	}

	result := &ServiceFirestore{
		Cfg:       this.Cfg,
		Clients:   this.Clients,
		UtilsRepo: this.UtilsRepo,
	}
	log.Printf("[IGNORE] [FIRESTORE] initialize")
	return result
}
