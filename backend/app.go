package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/rs/zerolog/log"

	"github.com/gorilla/mux"
	"go.opentelemetry.io/otel"
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
	LastRefresh   time.Time
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

// func respondWithError(w http.ResponseWriter, code int, message string) {
// 	respondWithJSON(w, code, map[string]string{"error": message})
// }

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
	tracer := otel.Tracer(os.Getenv("OTEL_SERVICE_NAME"))
	ctx := r.Context()
	ctx, span := tracer.Start(ctx, "list-topics")
	defer span.End()

	respondWithJSON(w, http.StatusOK, a.CurrentTopics)
}

func (a *App) refreshTopics() {
	for {
		latestTopics := a.getDeals(9, 1, 6)
		latestTopics = a.updateScores(latestTopics)
		log.Debug().Msg("Refreshing topics")
		a.CurrentTopics = latestTopics
		a.LastRefresh = time.Now()
		time.Sleep(60 * time.Second)
	}
}

func (a *App) updateScores(t []Topic) []Topic {
	for i := range t {
		t[i].Score = t[i].Votes.Up - t[i].Votes.Down
	}
	return t
}

func (a *App) isSponsor(t Topic) bool {
	return strings.HasPrefix(t.Title, "[Sponsored]")
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
			log.Warn().Msgf("could not unmarshal response body: %s\n %s", err)
		}

		for _, topic := range response.Topics {
			if a.isSponsor(topic) {
				continue
			}
			t = append(t, topic)
		}
	}
	return t
}
