package main

import (
	"net/http"
	//"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"src/apis/organization_api"
	"apis/authentication_api"
	"apis/configuration_api"
	"apis/procurement_api"
	"apis/master_api"
	"context"
	"log"
	"fmt"
	"time"
	"os"
	"os/signal"
	"syscall"
)


func main() {
	
	log.Println("Starting development server at http://127.0.0.1:8000/")
    log.Println("Quit the server with CONTROL-C.")

	router := mux.NewRouter().StrictSlash(true)

	
	router.HandleFunc("/", homePage)

	router.HandleFunc("/api/register/", authentication_api.Signup).Methods("POST")
	router.HandleFunc("/api/login/", authentication_api.Login).Methods("POST")
	router.HandleFunc("/api/token/refresh/{token}/", authentication_api.TokenRefresh).Methods("POST")
	router.HandleFunc("/api/email-verify/", authentication_api.VerifyEmail).Methods("GET")
	router.HandleFunc("/api/request-reset-email/", authentication_api.RequestPasswordResetEmail).Methods("POST")
	router.HandleFunc("/api/password-reset/{uidb64}/{token}/", authentication_api.PasswordTokenCheck).Methods("GET")
	router.HandleFunc("/api/password-reset-complete/", authentication_api.SetNewPassword).Methods("POST")
	router.HandleFunc("/api/get_user_menu/{user_id}/", authentication_api.Get_User_Menu).Methods("GET")
	//router.HandleFunc("/protected", authentication_api.TokenVerifyMiddleWare(ProtectedEndpoint)).Methods("GET")
	router.HandleFunc("/api/user/getusers/", authentication_api.GetUsers).Methods("GET")

	router.HandleFunc("/api/config/con-service-permissions/", configuration_api.GetUserServicePermission).Methods("GET")
	router.HandleFunc("/api/config/con-service-permissions/", configuration_api.PostUserServicePermission).Methods("POST")
	router.HandleFunc("/api/config/con-service-permissions/{user_id}/", configuration_api.GetUserServicePermForUser).Methods("GET")
	router.HandleFunc("/api/config/con-service-permissions/{user_id}/", configuration_api.PutUserServicePermForUser).Methods("PUT")
	router.HandleFunc("/api/config/con-service-permissions/{user_id}/", configuration_api.DelUserServicePermForUser).Methods("DELETE")

	router.HandleFunc("/api/master/con-modules/", master_api.GetConModules).Methods("GET")
	router.HandleFunc("/api/master/con-modules/", master_api.PostConModules).Methods("POST")
	router.HandleFunc("/api/master/con-modules/{module_id}/", master_api.GetConModule).Methods("GET")
	router.HandleFunc("/api/master/con-modules/{module_id}/", master_api.PutConModule).Methods("PUT")
	router.HandleFunc("/api/master/con-modules/{module_id}/", master_api.DelConModule).Methods("DELETE")

	router.HandleFunc("/api/master/con-services/", master_api.GetConServices).Methods("GET")
	router.HandleFunc("/api/master/con-services/", master_api.PostConServices).Methods("POST")
	router.HandleFunc("/api/master/con-services/{service_id}/", master_api.GetConService).Methods("GET")
	router.HandleFunc("/api/master/con-services/{service_id}/", master_api.PutConService).Methods("PUT")
	router.HandleFunc("/api/master/con-services/{service_id}/", master_api.DelConService).Methods("DELETE")
	router.HandleFunc("/api/master/get-countries/", master_api.GetCountries).Methods("GET")

	router.HandleFunc("/api/organization/getorganizations", organization_api.GetOrganizations).Methods("GET")
	router.HandleFunc("/api/organization/search/{keyword}", organization_api.Search).Methods("GET")
	router.HandleFunc("/api/organization/searchorgforyear/{min}/{max}", organization_api.SearchOrgForYear).Methods("GET")
	router.HandleFunc("/api/organization/create", organization_api.Create).Methods("POST")
	router.HandleFunc("/api/organization/update", organization_api.Update).Methods("PUT")
	router.HandleFunc("/api/organization/delete/{id}", organization_api.Delete).Methods("DELETE")

	router.HandleFunc("/api/procurement/getfarmers", procurement_api.GetFarmers).Methods("GET")
	//router.HandleFunc("/api/procurement/getcompanyfarmers/{company_name}", procurement_api.GetCompanyFarmers).Methods("GET")
	router.HandleFunc("/api/procurement/getfarmer/{farmer_name}", procurement_api.GetFarmer).Methods("GET")
	router.HandleFunc("/api/procurement/createfarmer/{farmer_name}", procurement_api.CreateFarmer).Methods("POST")
	router.HandleFunc("/api/procurement/updatefarmer/{farmer_name}", procurement_api.UpdateFarmer).Methods("PUT")
	router.HandleFunc("/api/procurement/deletefarmer/{farmer_name}", procurement_api.DeleteFarmer).Methods("DELETE")
	

	/*router.HandleFunc("/api/course/getcourses", university_api.GetCourses).Methods("GET")
	router.HandleFunc("/api/course/getcoursebyid/{courseid}", university_api.GetCourseByID).Methods("GET")
	router.HandleFunc("/api/course/getcoursebyname/{coursename}", university_api.GetCourseByName).Methods("GET")
	router.HandleFunc("/api/course/createcourse", university_api.CreateCourse).Methods("POST")
	router.HandleFunc("/api/course/updatecourse", university_api.UpdateCourse).Methods("PUT")
	router.HandleFunc("/api/course/deletecourse/{id}", university_api.DeleteCourse).Methods("DELETE")
	router.HandleFunc("/api/course/deletecoursebyname/{keyword}", university_api.DeleteCourseByName).Methods("DELETE")

	router.HandleFunc("/api/course/getstudents", university_api.GetStudents).Methods("GET")
	router.HandleFunc("/api/course/getstudentbyname/{keyword}", university_api.GetStudentByName).Methods("GET")
	router.HandleFunc("/api/course/createstudent", university_api.AddStudentToCourse).Methods("POST")
	router.HandleFunc("/api/course/updatestudent", university_api.UpdateStudent).Methods("PUT")
	router.HandleFunc("/api/course/updatestudentbyname/{keyword}", university_api.UpdateStudentByName).Methods("PUT")
	router.HandleFunc("/api/course/deletestudent/{id}", university_api.DeleteStudent).Methods("DELETE")
	router.HandleFunc("/api/course/deletestudentbyname/{keyword}", university_api.DeleteStudentByName).Methods("DELETE")

	router.HandleFunc("/api/course/getstudentdatapercourse/{keyword}", university_api.Get_Student_data_for_Course).Methods("GET")
	
	*/



	 corsOpts := cors.New(cors.Options{
        AllowedOrigins: []string{"*"},
        AllowCredentials: true,
        AllowedHeaders: []string{"*"},
    })

	//http.Handle("/", router) // enable the router and the CORS Policy

	srv := &http.Server{
		Handler:      corsOpts.Handler(router),
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Start Server
	go func() {
		fmt.Println("Starting FarmMobi Server ----  Welcome User")
		if err := srv.ListenAndServe(); err != nil {
			fmt.Println(err)
		}
	}()

	// Graceful Shutdown
	waitForShutdown(srv)


	/*err := http.ListenAndServe(":8000", corsOpts.Handler(router))
	if  err != nil {
        log.Fatal(err)
    }
    */
}

func waitForShutdown(srv *http.Server) {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive our signal.
	<-interruptChan

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	srv.Shutdown(ctx)

	fmt.Println("FarmMobi Server Shutting down ...")
	os.Exit(0)
}

func homePage(response http.ResponseWriter, r *http.Request){
	response.Header().Set("Content-Type", "application/json")
    fmt.Fprintf(response, "Welcome to HomePage!")
    fmt.Println("Endpoint Hit: HomePage")
}