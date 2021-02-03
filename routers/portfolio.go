package routers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/Kungfucoding23/portafolioBackend/models"
)

func getPortfolio() []models.Portfolio {
	dataPortfolio := make([]models.Portfolio, 0)
	raw, err := ioutil.ReadFile("db/portfolio.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	json.Unmarshal(raw, &dataPortfolio)
	return dataPortfolio
}

//PortfolioSection muestra el json
func PortfolioSection(w http.ResponseWriter, r *http.Request) {
	portfolio := getPortfolio()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	resp, err := json.Marshal(portfolio)
	if err != nil {
		panic(err)
	}
	w.Write(resp)
}
