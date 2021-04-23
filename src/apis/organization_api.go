package organization_api

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


func GetOrganizations(response http.ResponseWriter, request *http.Request) {

  fmt.Println("API: GetOrganizations")

  db, err := config.GetDB()

  if err != nil {
  	respondWithError(response, http.StatusBadRequest, err.Error())
  } else {

    defer db.Close()

  	organizationModel := models.OrganizationModel{
  		Db: db,
  	}
  	organizations, err2 := organizationModel.GetOrganizations()

  	if err2 != nil {
  		respondWithError(response, http.StatusBadRequest, err2.Error())
  	} else {
  		respondWithJson(response, http.StatusOK, organizations)
  	}
  }
}


func Search(response http.ResponseWriter, request *http.Request) {

  fmt.Println("API: Search")

  vars := mux.Vars(request)
  keyword := vars["keyword"]

  db, err := config.GetDB()

  if err != nil {
    respondWithError(response, http.StatusBadRequest, err.Error())
  } else {

    defer db.Close()

    organizationModel := models.OrganizationModel{
      Db: db,
    }
    organizations, err2 := organizationModel.Search(keyword)

    if err2 != nil {
      respondWithError(response, http.StatusBadRequest, err2.Error())
    } else {
      respondWithJson(response, http.StatusOK, organizations)
    }
  }
}



func SearchOrgForYear(response http.ResponseWriter, request *http.Request) {

  fmt.Println("API: SearchOrgForYear")

  vars := mux.Vars(request)
  min := vars["min"]
  max := vars["max"]

  db, err := config.GetDB()

  if err != nil {
    respondWithError(response, http.StatusBadRequest, err.Error())
  } else {

    defer db.Close()

    organizationModel := models.OrganizationModel{
      Db: db,
    }
    organizations, err2 := organizationModel.SearchOrgForYear(min, max)

    if err2 != nil {
      respondWithError(response, http.StatusBadRequest, err2.Error())
    } else {
      respondWithJson(response, http.StatusOK, organizations)
    }
  }
}


func Create(response http.ResponseWriter, request *http.Request) {

  fmt.Println("API: Create")

  var organization entities.Organization

  err := json.NewDecoder(request.Body).Decode(&organization)

  db, err := config.GetDB()

  if err != nil {
    respondWithError(response, http.StatusBadRequest, err.Error())
  } else {

    defer db.Close()

    organizationModel := models.OrganizationModel{

      Db: db,
    }
    err2 := organizationModel.Create(&organization)

    if err2 != nil {
      respondWithError(response, http.StatusBadRequest, err2.Error())
    } else {
      respondWithJson(response, http.StatusCreated, organization)
    }
  }
}


func Update(response http.ResponseWriter, request *http.Request) {

  fmt.Println("API: Update")

  var organization entities.Organization

  err := json.NewDecoder(request.Body).Decode(&organization)

  db, err := config.GetDB()

  if err != nil {
    respondWithError(response, http.StatusBadRequest, err.Error())
  } else {

    defer db.Close()

    organizationModel := models.OrganizationModel{

      Db: db,
    }
    _, err2 := organizationModel.Update(&organization)

    if err2 != nil {
      respondWithError(response, http.StatusBadRequest, err2.Error())
    } else {
      respondWithJson(response, http.StatusOK, organization)
    }
  }
}


func Delete(response http.ResponseWriter, request *http.Request) {

  fmt.Println("API: Delete")

  vars := mux.Vars(request)
  sid := vars["id"]

  id, _ := strconv.ParseInt(sid, 10, 64)

  db, err := config.GetDB()

  if err != nil {
    respondWithError(response, http.StatusBadRequest, err.Error())
  } else {

    defer db.Close()
    
    organizationModel := models.OrganizationModel{
      Db: db,
    }
    RowsAffected, err2 := organizationModel.Delete(uint64(id))

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
