package routers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/Kungfucoding23/portafolioBackend/models"
)

func getAbout() []models.About {
	dataAbout := make([]models.About, 0)
	raw, err := ioutil.ReadFile("db/About.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	json.Unmarshal(raw, &dataAbout)
	return dataAbout
}

//AboutSec muestra el json
func AboutSec(w http.ResponseWriter, r *http.Request) {
	about := getAbout()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	resp, err := json.Marshal(about)
	if err != nil {
		panic(err)
	}
	w.Write(resp)
}
