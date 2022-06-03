package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"joiner-management/models"
	"joiner-management/utils"
	"log"
	"net/http"
	"strconv"
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

func (defaultHandler DefaultHandler) GetJoiners(response http.ResponseWriter, request *http.Request) {
	var joiners []models.Joiner

	if result := defaultHandler.DB.Find(&joiners); result.Error != nil {
		fmt.Println(result.Error)
	}

	utils.ResponseOk(response, joiners)
}

func (defaultHandler DefaultHandler) DeleteJoiner(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, _ := strconv.Atoi(vars["id"])

	var joiner models.Joiner
	joiner = defaultHandler.getJoinerById(id)

	if (models.Joiner{}) == joiner {
		utils.ResponseNotFound(response)
	} else {
		defaultHandler.DB.Delete(&joiner)
		utils.ResponseNotContent(response)
	}
}

func (defaultHandler DefaultHandler) getJoinerById(id int) models.Joiner {
	var joiner models.Joiner
	if result := defaultHandler.DB.First(&joiner, id); result.Error != nil {
		fmt.Println(result.Error)
		return models.Joiner{}
	}
	return joiner
}
