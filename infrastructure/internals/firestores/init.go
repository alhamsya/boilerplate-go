package firestores

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
)

func New(this *ServiceFirestore) *ServiceFirestore {
	this.clients = make(map[string]*firestore.Client)
	for name, data := range this.Cfg.Firestore {
		client, err := firestore.NewClient(context.Background(), data.ProjectID)
		if err != nil {
			log.Fatalf("[INIT] [FIRESTORE] [%s] %v", name, err)
		}

		this.clients[name] = client
	}

	result := &ServiceFirestore{
		Cfg:       this.Cfg,
		UtilsRepo: this.UtilsRepo,
		clients:   this.clients,
	}
	log.Printf("[IGNORE] [FIRESTORE] initialize")
	return result
}
