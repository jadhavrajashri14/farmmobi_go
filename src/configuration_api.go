package configuration_api

import (
  "config"
  "models"
  "entities"
  "encoding/json"
  "net/http"
  "fmt"
  "strconv"
  "github.com/gorilla/mux"
)


func GetUserServicePermission(response http.ResponseWriter, request *http.Request) {

  fmt.Println("API: GetUserServicePermission")

  db, err := config.GetDB()

  if err != nil {
  	respondWithError(response, http.StatusBadRequest, err.Error())
  } else {

    defer db.Close()

  	configurationModel := models.ConfigurationModel{
  		Db: db,
  	}
  	con_service_permissions, err2 := configurationModel.GetUserServicePermission()

  	if err2 != nil {
  		respondWithError(response, http.StatusBadRequest, err2.Error())
  	} else {
  		respondWithJson(response, http.StatusOK, con_service_permissions)
  	}
  }
}


func PostUserServicePermission(response http.ResponseWriter, request *http.Request) {

  fmt.Println("API: PostUserServicePermission")

  var con_service_permission entities.CON_Service_Permission

  err := json.NewDecoder(request.Body).Decode(&con_service_permission)

  db, err := config.GetDB()

  if err != nil {
    respondWithError(response, http.StatusBadRequest, err.Error())
  } else {

    defer db.Close()

    configurationModel := models.ConfigurationModel{

      Db: db,
    }
    err2 := configurationModel.PostUserServicePermission(&con_service_permission)

    if err2 != nil {
      respondWithError(response, http.StatusBadRequest, err2.Error())
    } else {
      respondWithJson(response, http.StatusCreated, con_service_permission)
    }
  }
}


func GetUserServicePermForUser(response http.ResponseWriter, request *http.Request) {

  fmt.Println("API: GetUserServicePermForUser")
  var myerr entities.Error

  vars := mux.Vars(request)
  userId := vars["user_id"]

  db, err := config.GetDB()

  if err != nil {
    respondWithError(response, http.StatusBadRequest, err.Error())
  } else {

    defer db.Close()

    configurationModel := models.ConfigurationModel{
      Db: db,
    }

    id, err := strconv.ParseInt(userId, 10, 64)

    if err != nil {

     	myerr.Message = "Error while getting user id"
	 	respondWithError(response, http.StatusBadRequest, myerr.Message)
    } else {

    	con_service_permissions, err2 := configurationModel.GetUserServicePermForUser(id)

	    if err2 != nil {
	      respondWithError(response, http.StatusBadRequest, err2.Error())
	    } else {
	      respondWithJson(response, http.StatusOK, con_service_permissions)
	    }
    }
  }
}


func PutUserServicePermForUser(response http.ResponseWriter, request *http.Request) {

  fmt.Println("API: PutUserServicePermForUser")
  //vars := mux.Vars(request)
  //userId := vars["user_id"]

  var con_service_permission entities.CON_Service_Permission

  err := json.NewDecoder(request.Body).Decode(&con_service_permission)

  db, err := config.GetDB()

  if err != nil {
    respondWithError(response, http.StatusBadRequest, err.Error())
  } else {

    defer db.Close()

    configurationModel := models.ConfigurationModel{

      Db: db,
    }
    _, err2 := configurationModel.PutUserServicePermForUser(&con_service_permission)

    if err2 != nil {
      respondWithError(response, http.StatusBadRequest, err2.Error())
    } else {
      respondWithJson(response, http.StatusOK, con_service_permission)
    }
  }
}


func DelUserServicePermForUser(response http.ResponseWriter, request *http.Request) {

  fmt.Println("API: DelUserServicePermForUser")

  vars := mux.Vars(request)
  userId := vars["user_id"]

  userID, _ := strconv.ParseInt(userId, 10, 64)

  db, err := config.GetDB()

  if err != nil {
    respondWithError(response, http.StatusBadRequest, err.Error())
  } else {

    defer db.Close()
    
    configurationModel := models.ConfigurationModel{
      Db: db,
    }
    RowsAffected, err2 := configurationModel.DelUserServicePermForUser(userID)

    if err2 != nil {
      respondWithError(response, http.StatusBadRequest, err2.Error())
    } else {
      respondWithJson(response, http.StatusOK, map[string] int64{
        "RowsAffected" : RowsAffected,
      })
    }
  }
}


func respondWithError(w http.ResponseWriter, code int, msg string) {

	respondWithJson(w, code, map[string]string{"error": msg})
}


func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {

		response, _ := json.Marshal(payload)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		w.Write(response)
}
