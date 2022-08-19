package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	_ "github.com/joho/godotenv/autoload"
	"github.com/rs/zerolog/log"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// @title           RFD LAUNCHER API
// @version         1.0
// @description     An API for issue tracking
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    https://linktr.ee/davegallant
// @contact.email  davegallant@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

type App struct {
	DB       *gorm.DB
	Router   *mux.Router
	BasePath string
}

func (a *App) Initialize(dbDriver string, dbURI string) {
	db, err := gorm.Open(dbDriver, dbURI)
	if err != nil {
		panic("failed to connect database")
	}
	a.DB = db
	a.BasePath = "/api/v1"

	a.DB.AutoMigrate(&Topic{})

	a.Router = mux.NewRouter().PathPrefix(a.BasePath).Subrouter()
	http.Handle("/", a.Router)

	a.initializeRoutes()
}

func (a *App) Run(httpPort string) {
	log.Info().Msgf("Serving requests on port " + httpPort)
	if err := http.ListenAndServe(fmt.Sprintf(":"+httpPort), nil); err != nil {
		panic(err)
	}
	defer a.DB.Close()
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/topics", a.listTopics).Methods("GET")
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// listtopics godoc
// @Summary      Lists all topics stored in the database
// @Description  All topics will be listed. There is currently no pagination implemented.
// @ID           list-topics
// @Router       /topics [get]
// @Success 200 {array} Topic
func (a *App) listTopics(w http.ResponseWriter, r *http.Request) {
	var topics []Topic
	a.refreshDeals()
	a.DB.Find(&topics)
	respondWithJSON(w, http.StatusOK, topics)
}

func (a *App) refreshDeals() {
	topics := a.getDeals(9, 1, 10)
	log.Debug().Msg("Dropping deals")
	a.DB.DropTable(&Topic{})
	log.Debug().Msg("Refreshing the deals")
	a.DB.CreateTable(&Topic{})
	for _, topic := range topics {
		a.DB.Create(topic)
	}
}

func (a *App) getDeals(id int, firstPage int, lastPage int) []Topic {

	var t []Topic

	for i := firstPage; i < lastPage; i++ {
		requestURL := fmt.Sprintf("https://forums.redflagdeals.com/api/topics?forum_id=%d&per_page=40&page=%d", id, i)
		res, err := http.Get(requestURL)
		if err != nil {
			log.Warn().Msgf("error fetching deals: %s\n", err)
		}
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Warn().Msgf("could not read response body: %s\n", err)
		}

		var response TopicsResponse

		err = json.Unmarshal([]byte(body), &response)

		if err != nil {
			log.Warn().Msgf("could not unmarshal response body: %s\n", err)
		}

		t = append(t, response.Topics...)
	}
	return t
}
