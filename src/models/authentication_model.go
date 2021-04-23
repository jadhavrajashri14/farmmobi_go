package models

import (
	"database/sql"
	"entities"
	"time"
)


type UserModel struct {
	Db *sql.DB
}

type ProfileMenuModel struct {
	Db *sql.DB
}


func (userModel UserModel) GetUsers() (user []entities.User, err error) {

	rows, err := userModel.Db.Query("SELECT * from user")

	if err != nil {
		return nil, err
	} else {

		var users []entities.User

		for rows.Next() {

			var  organization_id    *int64 
			var  sub_org_id		  	*int64
			var  company_id		  	*int64
			var  user_id  			*int64
			var  username			string
			var  password			string 	
			var  created_at		  	string 	
			var  last_login 		*string
			var  is_admin 			bool
			var  is_verified     	bool 		
			var  is_active  		bool 		
			var  is_superuser  	  	bool 		
			var  is_staff   		bool 		
			var  email  			*string
			var  date_joined  		string  	
			var  first_name  		*string
			var  last_name  		*string
			var  user_level  		*int64	
			var  department  		*string
			var  designation  		*string
			var  module_id			*int64
			var  service_id		  	*int64
			var  profile_status  	bool 		


			err2 := rows.Scan(&user_id, &username, &password, &created_at, &last_login, &is_admin, &is_verified, &is_active,
				 &is_superuser, &is_staff, &email, &date_joined, &first_name, &last_name, &user_level, &department, &designation,
				 &profile_status, &company_id, &module_id, &organization_id, &service_id, &sub_org_id )

			if err2 != nil {
				return nil, err2
			} else {

				user := entities.User{
					Organization_id: organization_id, 
					Sub_org_id: sub_org_id, 
					Company_id: company_id, 
					User_id: user_id,
					Username: username,
					Password: password,
					Created_at: created_at,
					Last_login: last_login,
					Is_admin: is_admin,
					Is_verified: is_verified,
					Is_active: is_active,
					Is_superuser: is_superuser,
					Is_staff: is_staff,
					Email: email,
					Date_joined: date_joined,
					First_name: first_name,
					Last_name: last_name,
					User_level: user_level,
					Department: department,
					Designation: designation,
					Module_id: module_id,
					Service_id: service_id,
					Profile_status: profile_status}

				users = append(users, user)
			}
		}
		return users, nil
	}
}


func (userModel UserModel) GetUserByID(id uint64) (user []entities.User, err error) {

	rows, err := userModel.Db.Query("SELECT * from user where user_id = ?", id)

	if err != nil {
		return nil, err
	} else {

		var users []entities.User

		for rows.Next() {

			var  organization_id    *int64 
			var  sub_org_id		  	*int64
			var  company_id		  	*int64
			var  user_id  			*int64
			var  username			string
			var  password			string 	
			var  created_at		  	string 	
			var  last_login 		*string
			var  is_admin 			bool
			var  is_verified     	bool 		
			var  is_active  		bool 		
			var  is_superuser  	  	bool 		
			var  is_staff   		bool 		
			var  email  			*string
			var  date_joined  		string  	
			var  first_name  		*string
			var  last_name  		*string
			var  user_level  		*int64	
			var  department  		*string
			var  designation  		*string
			var  module_id			*int64
			var  service_id		  	*int64
			var  profile_status  	bool 		


			err2 := rows.Scan(&user_id, &username, &password, &created_at, &last_login, &is_admin, &is_verified, &is_active,
				 &is_superuser, &is_staff, &email, &date_joined, &first_name, &last_name, &user_level, &department, &designation,
				 &profile_status, &company_id, &module_id, &organization_id, &service_id, &sub_org_id )

			if err2 != nil {
				return nil, err2
			} else {

				user := entities.User{
					Organization_id: organization_id, 
					Sub_org_id: sub_org_id, 
					Company_id: company_id, 
					User_id: user_id,
					Username: username,
					Password: password,
					Created_at: created_at,
					Last_login: last_login,
					Is_admin: is_admin,
					Is_verified: is_verified,
					Is_active: is_active,
					Is_superuser: is_superuser,
					Is_staff: is_staff,
					Email: email,
					Date_joined: date_joined,
					First_name: first_name,
					Last_name: last_name,
					User_level: user_level,
					Department: department,
					Designation: designation,
					Module_id: module_id,
					Service_id: service_id,
					Profile_status: profile_status}

				users = append(users, user)
			}
		}
		return users, nil
	}
}


func (userModel UserModel) GetUserByName(keyword string) (user []entities.User, err error) {

	rows, err := userModel.Db.Query("SELECT * from user where username = ?", keyword)

	if err != nil {
		return nil, err
	} else {

		var users []entities.User

		for rows.Next() {

			var  organization_id    *int64 
			var  sub_org_id		  	*int64
			var  company_id		  	*int64
			var  user_id  			*int64
			var  username			string
			var  password			string 	
			var  created_at		  	string 	
			var  last_login 		*string
			var  is_admin 			bool
			var  is_verified     	bool 		
			var  is_active  		bool 		
			var  is_superuser  	  	bool 		
			var  is_staff   		bool 		
			var  email  			*string
			var  date_joined  		string  	
			var  first_name  		*string
			var  last_name  		*string
			var  user_level  		*int64	
			var  department  		*string
			var  designation  		*string
			var  module_id			*int64
			var  service_id		  	*int64
			var  profile_status  	bool 		


			err2 := rows.Scan(&user_id, &username, &password, &created_at, &last_login, &is_admin, &is_verified, &is_active,
				 &is_superuser, &is_staff, &email, &date_joined, &first_name, &last_name, &user_level, &department, &designation,
				 &profile_status, &company_id, &module_id, &organization_id, &service_id, &sub_org_id )

			if err2 != nil {
				return nil, err2
			} else {

				user := entities.User{
					Organization_id: organization_id, 
					Sub_org_id: sub_org_id, 
					Company_id: company_id, 
					User_id: user_id,
					Username: username,
					Password: password,
					Created_at: created_at,
					Last_login: last_login,
					Is_admin: is_admin,
					Is_verified: is_verified,
					Is_active: is_active,
					Is_superuser: is_superuser,
					Is_staff: is_staff,
					Email: email,
					Date_joined: date_joined,
					First_name: first_name,
					Last_name: last_name,
					User_level: user_level,
					Department: department,
					Designation: designation,
					Module_id: module_id,
					Service_id: service_id,
					Profile_status: profile_status}

				users = append(users, user)
			}
		}
		return users, nil
	}
}


func (userModel UserModel) SearchUserEmailExists(keyword string) (user_id *int64, err error) {

	rows, err := userModel.Db.Query("SELECT user_id from user where email like ?", "%"+keyword+"%")

	if err != nil {
		return nil, err
	} else {

		for rows.Next() {

			var  user_id  int64
			
			err2 := rows.Scan(&user_id)

			if err2 != nil {
				return nil, err2
			} else {

				return &user_id, nil
			}
		}
		return nil, err
	}
}


func (userModel UserModel) GetMaxUserId() (user_id *int64, err error) {

	rows, err := userModel.Db.Query("SELECT max(user_id) from user")

	if err != nil {
		return nil, err
	} else {

		for rows.Next() {

			var  user_id  int64
			
			err2 := rows.Scan(&user_id)

			if err2 != nil {
				return nil, err2
			} else {

				return &user_id, nil
			}
		}
		return nil, err
	}
}


func (userModel UserModel) Create(user *entities.User) (err error) {


	result, err := userModel.Db.Exec(`INSERT INTO user(organization_id, sub_org_id, company_id, user_id, username, 
		password, created_at, last_login, is_admin, is_verified, is_active, is_superuser, is_staff, email, date_joined,  
		first_name, last_name, user_level, department, designation, module_id, service_id, profile_status) values(?, ?, ?, 
		?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`, user.Organization_id, user.Sub_org_id, user.Company_id, 
		user.User_id, user.Username, user.Password, time.Now().UTC().Format("2006-01-02 03:04:05"), time.Now().UTC().Format("2006-01-02 03:04:05"),
		user.Is_admin, user.Is_verified, user.Is_active, user.Is_superuser, user.Is_staff, user.Email, time.Now().UTC().Format("2006-01-02 03:04:05"), 
		user.First_name, user.Last_name, user.User_level, user.Department, user.Designation, user.Module_id, user.Service_id, user.Profile_status)

	if err != nil {
		return  err
	} else { 

		var id int64

		id, _ = result.LastInsertId()

		*user.User_id = id

		return nil
	}
}


func (userModel UserModel) UpdateFlags(id uint64) (int64, error) {

	result, err := userModel.Db.Exec(`UPDATE user SET is_verified = true, is_active = true WHERE user_id = ?`, id)

	if err != nil {
		return  0, err
	} else { 

		return result.RowsAffected()
	}
}

func (userModel UserModel) UpdatePwd(id uint64, keyword string) (int64, error) {

	result, err := userModel.Db.Exec(`UPDATE user SET password = ? where user_id = ?`, keyword, id)

	if err != nil {
		return  0, err
	} else { 

		return result.RowsAffected()
	}
}


// Delete is setting the is_active flag = false and not deleting the data
func (userModel UserModel) Delete(id uint64) (int64, error) {

	result, err := userModel.Db.Exec(`UPDATE user SET is_active = false WHERE user_id = ?`, id)

	if err != nil {
		return  0, err
	} else { 

		return result.RowsAffected()
	}
}


func (profileMenuModel ProfileMenuModel) Get_User_Menu(id int64) (profileMenu []entities.ProfileMenu, err error) {

	rows, err := profileMenuModel.Db.Query("call go_get_user_menu(?)", id)

	if err != nil {
		return nil, err
	} else {

		var profileMenu []entities.ProfileMenu

		for rows.Next() {

			var  module		  	string 	
			var  service 		string
	
			err2 := rows.Scan(&module, &service)

			if err2 != nil {
				return nil, err2
			} else {

				menuItem := entities.ProfileMenu{
					Module: module, 
					Service: service}

				profileMenu = append(profileMenu, menuItem)
			}
		}
		return profileMenu, nil
	}
}