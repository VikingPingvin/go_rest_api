package router

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"vikingpingvin/restpractice/article"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

func InitRouter() (router *mux.Router) {
	router = mux.NewRouter()

	apiv1 := router.PathPrefix("/api/v1").Subrouter()
	apiv1.HandleFunc("", apiv1get).Methods(http.MethodGet)
	apiv1.HandleFunc("", apiv1post).Methods(http.MethodPost)
	apiv1.HandleFunc("", apiv1put).Methods(http.MethodPut)
	apiv1.HandleFunc("", apiv1delete).Methods(http.MethodDelete)
	apiv1.HandleFunc("", apiv1notFound)

	apiArticles := apiv1.PathPrefix("/articles").Subrouter()
	apiArticles.HandleFunc("", listArticles).Methods(http.MethodGet)
	apiArticles.HandleFunc("/{id:[0-9]+}", returnArticle).Methods(http.MethodGet)
	apiArticles.HandleFunc("", addArticle).Methods(http.MethodPost)

	router.HandleFunc("/", homeHandler)
	router.HandleFunc("/articles", articlesHandler)
	router.HandleFunc("/api", restHandler)
	return router
}
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}

func articlesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Articles!")
}

func apiv1get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`"message":"GET"`))
}
func apiv1post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`"message":"POST"`))
}
func apiv1put(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(`"message":"PUT"`))
}
func apiv1delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`"message":"DELETE"`))
}
func apiv1notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`"message":"not found"`))
}

func restHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`"message":"LEGACY API"`))
}

func listArticles(w http.ResponseWriter, r *http.Request) {
	log.Info().Msg(fmt.Sprintf("API:%v", r))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(article.ReturnAllArticles())
}

func returnArticle(w http.ResponseWriter, r *http.Request) {
	log.Info().Msg(fmt.Sprintf("API:\n%v", r.URL))
	vars := mux.Vars(r)
	id := vars["id"]
	i, err := strconv.Atoi(id)
	if err != nil {
		log.Warn().Msg("Parameter ID is not an integer!")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(article.ReturnArticle(i))
}
func addArticle(w http.ResponseWriter, r *http.Request) {
	log.Info().Msg(fmt.Sprintf("API:%v", r))

	// Parse body to JSON
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Info().Msg("Cant read POST body")
		return
	}
	var jsonBody article.Article
	json.Unmarshal([]byte(body), &jsonBody)

	w.WriteHeader(http.StatusCreated)
	article.AddArticle(jsonBody)
}
