package handlers

import (
	"encoding/json"
	"io/ioutil"
	"joiner-management/models"
	"joiner-management/utils"
	"log"
	"net/http"
)

func (defaultHandler DefaultHandler) AddJoiner(response http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()
	body, err := ioutil.ReadAll(request.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var joiner models.Joiner
	err = json.Unmarshal(body, &joiner)
	if err != nil {
		utils.ResponseBadRequest(response)
	}

	if result := defaultHandler.DB.Create(&joiner); result.Error != nil {
		log.Fatalln(result.Error)
	}

	utils.ResponseCreated(response, joiner)
}
