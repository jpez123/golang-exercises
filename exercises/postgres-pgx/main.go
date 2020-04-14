package main

import (
	"encoding/json"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/types"
	"github.com/pkg/errors"
)

//ContactFavorites - JSON type
type ContactFavorites struct {
	Colors []string `json:"colors"`
}

//Contact - model
type Contact struct {
	ID                   int
	Name, Address, Phone string
	FavoritesJSON        types.JSONText    `db:"favorites"`
	Favorites            *ContactFavorites `db:"-"`

	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}

//Creates the flag to listen to url address
//Allows the import of html template to be read
//DB Connection
var (
	connectionString = flag.String("conn", getenvWithDefault("DATABASE_URL", "postgres://jzhang:@localhost:5432/local_contacts"), "Postgres connection string")
	listenAddr       = flag.String("addr", getenvWithDefault("LISTENADR", ":8080"), "HTTP address to listen on")
	db               *sqlx.DB
	tmpl             = template.New("")
)

//Chooses the default or provided address
func getenvWithDefault(name, defaultValue string) string {
	val := os.Getenv(name)
	if val == "" {
		val = defaultValue
	}

	return val
}

func main() {
	//Parses the flag
	flag.Parse()

	//Reads HTML file in template directory
	var err error

	tmpl.Funcs(template.FuncMap{"StringsJoin": strings.Join})
	_, err = tmpl.ParseGlob(filepath.Join(".", "templates", "*.html"))
	if err != nil {
		log.Fatalf("Unable to parse templates: %v\n", err)
	}

	//Connect to DB
	if *connectionString == "" {
		log.Fatalln("Please pass the connection string using the -conn option")
	}

	db, err = sqlx.Connect("pgx", *connectionString)
	if err != nil {
		log.Fatalf("Unable to establish connection: ^v\n", err)
	}

	//Starts and logs server
	http.HandleFunc("/", handler)
	log.Printf("listening on %s\n", *listenAddr)
	http.ListenAndServe(*listenAddr, nil)
}

//Serves the page
func handler(w http.ResponseWriter, r *http.Request) {
	contacts, err := fetchContacts()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	tmpl.ExecuteTemplate(w, "index.html", struct{ Contacts []*Contact }{contacts})
}

//Fetch contacts from database
func fetchContacts() ([]*Contact, error) {
	contacts := []*Contact{}
	err := db.Select(&contacts, "select * from contacts")
	if err != nil {
		return nil, errors.Wrap(err, "Unable to fetch contacts")
	}

	for _, contact := range contacts {
		err := json.Unmarshal(contact.FavoritesJSON, &contact.Favorites)

		if err != nil {
			return nil, errors.Wrap(err, "Unable to parse JSON favorites")
		}
	}

	return contacts, nil
}
