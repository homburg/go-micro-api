package services

import (
	"context"
	"fmt"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"github.com/byblix/go-micro/models"
	"google.golang.org/api/option"
)

// GetFBDataByChild ...
func GetFBDataByChild(dbRef string, childRef string) models.Profile {
	var data models.Profile
	ref, ctx := GetDBRefAndCTX(os.Getenv("ENV") + "/" + dbRef)
	if err := ref.Child(childRef).OrderByKey().Get(ctx, &data); err != nil {
		log.Fatalf("Error unmarshaling result: %s", err)
	}
	return data
}

// GetDBRefAndCTX adasd
func GetDBRefAndCTX(path string) (db.Ref, context.Context) {
	fmt.Println(path)
	ctx := context.Background()
	app := initFirebaseAdmin(ctx)
	dbclient, err := app.Database(ctx)
	if err != nil {
		log.Fatalf("Error initializing db %s\n", err)
	}
	ref := dbclient.NewRef(path)
	return *ref, ctx
}

func initFirebaseAdmin(ctx context.Context) *firebase.App {
	firebaseAdminRef := "config/fb-admin-key-" + os.Getenv("ENV") + ".json"
	opt := option.WithCredentialsFile(firebaseAdminRef)
	conf := &firebase.Config{DatabaseURL: os.Getenv("FB_DATABASE_URL")}
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		log.Fatalf("Error initializing app %s\n", err)
	}
	return app
}
