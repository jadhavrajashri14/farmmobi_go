package procurement_api

import (
  "config"
  "entities"
  "github.com/gorilla/mux"
  "encoding/json"
  "models"
	"net/http"
  "fmt"
  "strconv"
)


func GetFarmers(response http.ResponseWriter, request *http.Request) {

  fmt.Println("API: GetFarmers")

  db, err := config.GetDB()

  if err != nil {
  	respondWithError(response, http.StatusBadRequest, err.Error())
  } else {

    defer db.Close()

  	procurementModel := models.ProcurementModel{
  		Db: db,
  	}
  	farmers, err2 := procurementModel.GetFarmers()

  	if err2 != nil {
  		respondWithError(response, http.StatusBadRequest, err2.Error())
  	} else {
  		respondWithJson(response, http.StatusOK, farmers)
  	}
  }
}


func GetFarmer(response http.ResponseWriter, request *http.Request) {

  fmt.Println("API: GetFarmer")

  vars := mux.Vars(request)
  farmerName := vars["farmer_name"]

  db, err := config.GetDB()

  if err != nil {
    respondWithError(response, http.StatusBadRequest, err.Error())
  } else {

    defer db.Close()

    procurementModel := models.ProcurementModel{
      Db: db,
    }
    farmer, err2 := procurementModel.GetFarmer(farmerName)

    if err2 != nil {
      respondWithError(response, http.StatusBadRequest, err2.Error())
    } else {
      respondWithJson(response, http.StatusOK, farmer)
    }
  }
}


func CreateFarmer(response http.ResponseWriter, request *http.Request) {

  fmt.Println("API: CreateFarmer")

  var farmer_register entities.Farmer_Register

  err := json.NewDecoder(request.Body).Decode(&farmer_register)

  db, err := config.GetDB()

  if err != nil {
    respondWithError(response, http.StatusBadRequest, err.Error())
  } else {

    defer db.Close()

    procurementModel := models.ProcurementModel{
      Db: db,
    }
    err2 := procurementModel.CreateFarmer(&farmer_register)

    if err2 != nil {
      respondWithError(response, http.StatusBadRequest, err2.Error())
    } else {
      respondWithJson(response, http.StatusCreated, farmer_register)
    }
  }
}


func UpdateFarmer(response http.ResponseWriter, request *http.Request) {

  fmt.Println("API: UpdateFarmer")

  var farmer_register entities.Farmer_Register

  err := json.NewDecoder(request.Body).Decode(&farmer_register)

  db, err := config.GetDB()

  if err != nil {
    respondWithError(response, http.StatusBadRequest, err.Error())
  } else {

    defer db.Close()

    procurementModel := models.ProcurementModel{
      Db: db,
    }
    _, err2 := procurementModel.UpdateFarmer(&farmer_register)

    if err2 != nil {
      respondWithError(response, http.StatusBadRequest, err2.Error())
    } else {
      respondWithJson(response, http.StatusOK, farmer_register)
    }
  }
}


func DeleteFarmer(response http.ResponseWriter, request *http.Request) {

  fmt.Println("API: Deletefarmer")

  vars := mux.Vars(request)
  sid := vars["farmer_name"]

  id, _ := strconv.ParseInt(sid, 10, 64)

  db, err := config.GetDB()

  if err != nil {
    respondWithError(response, http.StatusBadRequest, err.Error())
  } else {

    defer db.Close()
    
    procurementModel := models.ProcurementModel{
      Db: db,
    }
    RowsAffected, err2 := procurementModel.DeleteFarmer(uint64(id))

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
