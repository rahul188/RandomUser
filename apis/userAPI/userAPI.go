package userAPI

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/randomUser/config"
	"github.com/randomUser/entities"
	"github.com/randomUser/models"
)

// FindAll is used for finding all records
func FindAll(response http.ResponseWriter, request *http.Request) {
	db, err := config.GetDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		userModel := models.UserModel{
			Db: db,
		}
		users, err2 := userModel.FindAll()
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJSON(response, http.StatusOK, users)
		}
	}
}

// Search function used for search user name in table
func Search(response http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)
	keyword := vars["keyword"]
	db, err := config.GetDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		userModel := models.UserModel{
			Db: db,
		}
		users, err2 := userModel.Search(keyword)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJSON(response, http.StatusOK, users)
		}
	}
}

//Add Function is used for adding data to table
func Add(response http.ResponseWriter, request *http.Request) {

	decoder := json.NewDecoder(request.Body)
	var t entities.UserName
	err := decoder.Decode(&t)

	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	}

	keyword := t.Name

	db, err := config.GetDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		userModel := models.UserModel{
			Db: db,
		}
		err2 := userModel.Add(keyword)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJSON(response, http.StatusOK, "msg : 1 Record Added Successfully")
		}
	}
}

// Remove Function Will Delete records as per given ID
func Remove(response http.ResponseWriter, request *http.Request) {

	decoder := json.NewDecoder(request.Body)
	var t entities.UserID
	err := decoder.Decode(&t)

	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	}

	keyword := t.Id
	db, err := config.GetDB()

	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		userModel := models.UserModel{
			Db: db,
		}
		err2 := userModel.Remove(keyword)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJSON(response, http.StatusOK, "msg : 1 Record Deleted Successfully")
		}
	}
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}

//respondWithJSON function is used for formatting data to JSON
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
