package main

import (
	"context"
	"feedPost-service/pkg/models"
	"flag"
	"fmt"
	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
	"log"
	"net/http"
	"os"
	"time"

	"feedPost-service/pkg/models/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	comments    *mongodb.CommentModel
	contents    *mongodb.ContentModel
	feedPosts   *mongodb.FeedPostModel
	posts       *mongodb.PostsModel
	locations   *mongodb.LocationModel
}

func main() {

	fmt.Printf("Found multiple documents (array of pointers): %+v\n")

	// Define command-line flags
	serverAddr := flag.String("serverAddr", "", "HTTP server network address")
	serverPort := flag.Int("serverPort", 4000, "HTTP server network port")
	mongoURI := flag.String("mongoURI", "mongodb://localhost:27017", "Database hostname url")
	mongoDatabse := flag.String("mongoDatabse", "feedPosts", "Database name")
	enableCredentials := flag.Bool("enableCredentials", false, "Enable the use of credentials for mongo connection")
	flag.Parse()

	// Create logger for writing information and error messages.
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Create mongo client configuration
	co := options.Client().ApplyURI(*mongoURI)
	if *enableCredentials {
		co.Auth = &options.Credential{
			Username: os.Getenv("MONGODB_USERNAME"),
			Password: os.Getenv("MONGODB_PASSWORD"),
		}
	}


	// Establish database connection
	client, err := mongo.NewClient(co)
	if err != nil {
		errLog.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		errLog.Fatal(err)
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	infoLog.Printf("Database connection established")

	// Initialize a new instance of application containing the dependencies.
	app := &application{
		infoLog:  infoLog,
		errorLog: errLog,
		contents: &mongodb.ContentModel{
			C: client.Database(*mongoDatabse).Collection("contents"),
		},
		comments: &mongodb.CommentModel{
			C: client.Database(*mongoDatabse).Collection("comments"),
		},
		locations: &mongodb.LocationModel{
			C: client.Database(*mongoDatabse).Collection("locations"),
		},
		posts: &mongodb.PostsModel{
			C: client.Database(*mongoDatabse).Collection("posts"),
		},
		feedPosts: &mongodb.FeedPostModel{
			C: client.Database(*mongoDatabse).Collection("feedPosts"),
		},
	}

	// Initialize a new http.Server struct.
	serverURI := fmt.Sprintf("%s:%d", *serverAddr, *serverPort)
	srv := &http.Server{
		Addr:         serverURI,
		ErrorLog:     errLog,
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	collection := client.Database(*mongoDatabse).Collection("contents")
	content := models.Content{uuid.UUID{},"Image", "Video"}
	insertResult, err := collection.InsertOne(context.TODO(), content)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)



	infoLog.Printf("Starting server on %s", serverURI)
	err = srv.ListenAndServe()
	errLog.Fatal(err)

}
