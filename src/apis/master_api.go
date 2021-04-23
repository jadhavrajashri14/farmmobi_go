package master_api

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


func GetCountries(response http.ResponseWriter, request *http.Request) {

  fmt.Println("API: GetCountries")

  db, err := config.GetDB()

  if err != nil {
    respondWithError(response, http.StatusBadRequest, err.Error())
  } else {

    defer db.Close()

    masterModel := models.MasterModel{
      Db: db,
    }
    countries, err2 := masterModel.GetCountries()

    if err2 != nil {
      respondWithError(response, http.StatusBadRequest, err2.Error())
    } else {
      fmt.Println("API: GetCountries : ", countries)
      respondWithJson(response, http.StatusOK, countries)
    }
  }
}


func GetConModules(response http.ResponseWriter, request *http.Request) {

  fmt.Println("API: GetConModules")

  db, err := config.GetDB()

  if err != nil {
  	respondWithError(response, http.StatusBadRequest, err.Error())
  } else {

    defer db.Close()

  	masterModel := models.MasterModel{
  		Db: db,
  	}
  	con_modules, err2 := masterModel.GetConModules()

  	if err2 != nil {
  		respondWithError(response, http.StatusBadRequest, err2.Error())
  	} else {
  		respondWithJson(response, http.StatusOK, con_modules)
  	}
  }
}


func PostConModules(response http.ResponseWriter, request *http.Request) {

  fmt.Println("API: PostConModules")

  var con_module entities.CON_Module

  err := json.NewDecoder(request.Body).Decode(&con_module)

  db, err := config.GetDB()

  if err != nil {
    respondWithError(response, http.StatusBadRequest, err.Error())
  } else {

    defer db.Close()

    masterModel := models.MasterModel{

      Db: db,
    }
    err2 := masterModel.PostConModules(&con_module)

    if err2 != nil {
      respondWithError(response, http.StatusBadRequest, err2.Error())
    } else {
      respondWithJson(response, http.StatusCreated, con_module)
    }
  }
}


func GetConModule(response http.ResponseWriter, request *http.Request) {

  fmt.Println("API: GetConModule")
  var myerr entities.Error

  vars := mux.Vars(request)
  moduleId := vars["module_id"]

  db, err := config.GetDB()

  if err != nil {
    respondWithError(response, http.StatusBadRequest, err.Error())
  } else {

    defer db.Close()

    masterModel := models.MasterModel{
      Db: db,
    }

    id, err := strconv.ParseInt(moduleId, 10, 64)

    if err != nil {

     	myerr.Message = "Error while parsing module id"
	 	respondWithError(response, http.StatusBadRequest, myerr.Message)
    } else {

    	con_modules, err2 := masterModel.GetConModule(id)

	    if err2 != nil {
	      respondWithError(response, http.StatusBadRequest, err2.Error())
	    } else {
	      respondWithJson(response, http.StatusOK, con_modules)
	    }
    }
  }
}


func PutConModule(response http.ResponseWriter, request *http.Request) {

  fmt.Println("API: PutConModule")
 //vars := mux.Vars(request)
  //userId := vars["user_id"]

  var con_module entities.CON_Module

  err := json.NewDecoder(request.Body).Decode(&con_module)

  db, err := config.GetDB()

  if err != nil {
    respondWithError(response, http.StatusBadRequest, err.Error())
  } else {

    defer db.Close()

    masterModel := models.MasterModel{

      Db: db,
    }
    _, err2 := masterModel.PutConModule(&con_module)

    if err2 != nil {
      respondWithError(response, http.StatusBadRequest, err2.Error())
    } else {
      respondWithJson(response, http.StatusOK, con_module)
    }
  }
}


func DelConModule(response http.ResponseWriter, request *http.Request) {

  fmt.Println("API: DelConModule")

  vars := mux.Vars(request)
  moduleId := vars["module_id"]

  moduleID, _ := strconv.ParseInt(moduleId, 10, 64)

  db, err := config.GetDB()

  if err != nil {
    respondWithError(response, http.StatusBadRequest, err.Error())
  } else {

    defer db.Close()
    
    masterModel := models.MasterModel{
      Db: db,
    }
    RowsAffected, err2 := masterModel.DelConModule(moduleID)

    if err2 != nil {
      respondWithError(response, http.StatusBadRequest, err2.Error())
    } else {
      respondWithJson(response, http.StatusOK, map[string] int64{
        "RowsAffected" : RowsAffected,
      })
    }
  }
}


func GetConServices(response http.ResponseWriter, request *http.Request) {

  fmt.Println("API: GetConServices")

  db, err := config.GetDB()

  if err != nil {
    respondWithError(response, http.StatusBadRequest, err.Error())
  } else {

    defer db.Close()

    masterModel := models.MasterModel{
      Db: db,
    }
    con_services, err2 := masterModel.GetConServices()

    if err2 != nil {
      respondWithError(response, http.StatusBadRequest, err2.Error())
    } else {
      respondWithJson(response, http.StatusOK, con_services)
    }
  }
}


func PostConServices(response http.ResponseWriter, request *http.Request) {

  fmt.Println("API: PostConServices")

  var con_service entities.CON_Service

  err := json.NewDecoder(request.Body).Decode(&con_service)

  db, err := config.GetDB()

  if err != nil {
    respondWithError(response, http.StatusBadRequest, err.Error())
  } else {

    defer db.Close()

    masterModel := models.MasterModel{

      Db: db,
    }
    err2 := masterModel.PostConServices(&con_service)

    if err2 != nil {
      respondWithError(response, http.StatusBadRequest, err2.Error())
    } else {
      respondWithJson(response, http.StatusCreated, con_service)
    }
  }
}


func GetConService(response http.ResponseWriter, request *http.Request) {

  fmt.Println("API: GetConService")
  var myerr entities.Error

  vars := mux.Vars(request)
  serviceId := vars["service_id"]

  db, err := config.GetDB()

  if err != nil {
    respondWithError(response, http.StatusBadRequest, err.Error())
  } else {

    defer db.Close()

    masterModel := models.MasterModel{
      Db: db,
    }

    id, err := strconv.ParseInt(serviceId, 10, 64)

    if err != nil {

      myerr.Message = "Error while parsing service id"
    respondWithError(response, http.StatusBadRequest, myerr.Message)
    } else {

      con_services, err2 := masterModel.GetConService(id)

      if err2 != nil {
        respondWithError(response, http.StatusBadRequest, err2.Error())
      } else {
        respondWithJson(response, http.StatusOK, con_services)
      }
    }
  }
}


func PutConService(response http.ResponseWriter, request *http.Request) {

  fmt.Println("API: PutConService")
 //vars := mux.Vars(request)
  //userId := vars["user_id"]

  var con_service entities.CON_Service

  err := json.NewDecoder(request.Body).Decode(&con_service)

  db, err := config.GetDB()

  if err != nil {
    respondWithError(response, http.StatusBadRequest, err.Error())
  } else {

    defer db.Close()

    masterModel := models.MasterModel{

      Db: db,
    }
    _, err2 := masterModel.PutConService(&con_service)

    if err2 != nil {
      respondWithError(response, http.StatusBadRequest, err2.Error())
    } else {
      respondWithJson(response, http.StatusOK, con_service)
    }
  }
}


func DelConService(response http.ResponseWriter, request *http.Request) {

  fmt.Println("API: DelConService")

  vars := mux.Vars(request)
  serviceId := vars["service_id"]

  serviceID, _ := strconv.ParseInt(serviceId, 10, 64)

  db, err := config.GetDB()

  if err != nil {
    respondWithError(response, http.StatusBadRequest, err.Error())
  } else {

    defer db.Close()
    
    masterModel := models.MasterModel{
      Db: db,
    }
    RowsAffected, err2 := masterModel.DelConService(serviceID)

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
