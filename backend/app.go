package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	_ "github.com/joho/godotenv/autoload"
	"github.com/rs/zerolog/log"

	"github.com/gorilla/mux"
)

// @title           RFD FYI API
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
	Router        *mux.Router
	BasePath      string
	CurrentTopics []Topic
}

func (a *App) Initialize() {
	a.BasePath = "/api/v1"

	a.Router = mux.NewRouter().PathPrefix(a.BasePath).Subrouter()
	http.Handle("/", a.Router)

	a.initializeRoutes()
}

func (a *App) Run(httpPort string) {
	log.Info().Msgf("Serving requests on port " + httpPort)
	if err := http.ListenAndServe(fmt.Sprintf(":"+httpPort), nil); err != nil {
		panic(err)
	}
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
	respondWithJSON(w, http.StatusOK, topics)
}

func (a *App) refreshDeals() {
	latestTopics := a.getDeals(9, 1, 4)
	// TODO: only drop deals if a timer has been met
	log.Debug().Msg("Refreshing deals")
	a.CurrentTopics = latestTopics
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
