package entities

import (
  "fmt"
)

type User struct {
  Organization_id     *int64   		`json:"organization_id"`
  Sub_org_id		      *int64   		`json:"sub_org_id"`
  Company_id		      *int64   	 	`json:"company_id"`
  User_id			        *int64		  `json:"user_id"`
  Username			      string 			`json:"username"`
  Password			      string 			`json:"password"`
  Created_at		      string 			`json:"created_at"`
  Last_login 		      *string  	 	`json:"last_login"`
  Is_admin			      bool 	 			`json:"is_admin"`
  Is_verified     	  bool 				`json:"is_verified"`
  Is_active  		      bool 				`json:"is_active"`
  Is_superuser  	    bool 				`json:"is_superuser"`
  Is_staff   		      bool 				`json:"is_staff"`
  Email  			        *string 		`json:"email"`
  Date_joined  		    string  		`json:"date_joined"`
  First_name  		    *string 	  `json:"first_name"`
  Last_name  		      *string	    `json:"last_name"`
  User_level  		    *int64   	  `json:"user_level"`
  Department  		    *string  	  `json:"department"`
  Designation  		    *string	    `json:"designation"`
  Module_id			      *int64 	    `json:"module_id"`
  Service_id		      *int64   	  `json:"service_id"`
  Profile_status  	  bool 				`json:"profile_status"`
}


type JWT struct {
  Token			  string 			`json:"token"`
}


type ProfileMenu struct {
  Module       string
  Service     string
}



func (user User) ToString() string {

	return fmt.Sprintf(`Organization_id:%d\nSub_org_id:%d\nCompany_id:%d\nUser_id:%d\nUsername:%s\nPassword:%s\n` +
			`Created_at:%s\nLast_login:%s\nIs_admin:%s\nIs_verified:%s\nIs_active:%s\nIs_superuser:%s\nIs_staff:%s\n` +
			`Email:%s\nDate_joined:%s\nFirst_name:%s\nLast_name:%s\nUser_level:%s\nDepartment:%s\nDesignation:%s\n` +
			`Module_id:%d\nService_id:%d\nProfile_status:%s\n`,
			user.Organization_id, user.Sub_org_id, user.Company_id, user.User_id, user.Username, user.Password,
			user.Created_at, user.Last_login, user.Is_admin, user.Is_verified, user.Is_active, user.Is_superuser,
			user.Is_staff, user.Email, user.Date_joined, user.First_name, user.Last_name, user.User_level,
			user.Department, user.Designation, user.Module_id, user.Service_id, user.Profile_status)
}


func (jwt JWT) ToString() string {

	return fmt.Sprintf(`Token:%s\n`, jwt.Token)
}


func (profileMenu ProfileMenu) ToString() string {

  return fmt.Sprintf(`Module:%s\nService:%s\n`, profileMenu.Module, profileMenu.Service)
}