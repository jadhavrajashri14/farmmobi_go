package models

import (
	"database/sql"
	"entities"
)


type ConfigurationModel struct {
	Db *sql.DB
}


func (configurationModel ConfigurationModel) GetUserServicePermission() (con_service_permission []entities.CON_Service_Permission, err error) {

	rows, err := configurationModel.Db.Query("SELECT * from con_service_permission")

	if err != nil {
		return nil, err
	} else {

		var con_service_permissions []entities.CON_Service_Permission

		for rows.Next() {

			var service_Permission_id   *int64 
  			var service_id              *int64   
  			var user_id                 *int64 
  			var allow_add               bool    
  			var allow_view              bool      
  			var allow_delete            bool      
  			var allow_approval          bool    
  			var allow_print             bool     
  			var active_status           bool      
  			var delete_status           bool    
  			var entry_doneby_loginid   int64     
  			var entry_date             *string    
  			var modify_login_id        *int64     
  			var modify_date            *string    
  			var organization_id        *int64    
  			var sub_org_id             *int64   
  			var company_id             *int64  


			err2 := rows.Scan(&service_Permission_id, &allow_add, &allow_view, &allow_delete,
				&allow_approval, &allow_print, &active_status, &delete_status, &entry_doneby_loginid, &entry_date, 
				&modify_login_id, &modify_date,  &company_id, &organization_id, &service_id, &sub_org_id,  &user_id)

			if err2 != nil {
				return nil, err2
			} else {
				con_service_permission := entities.CON_Service_Permission{
					Service_Permission_id: service_Permission_id, 
					Allow_add: allow_add,
					Allow_view: allow_view,
					Allow_delete: allow_delete,
					Allow_approval: allow_approval,
					Allow_print: allow_print,
					Active_status: active_status,
					Delete_status: delete_status,
					Entry_doneby_loginid: entry_doneby_loginid,
					Entry_date: entry_date,
					Modify_login_id: modify_login_id,
					Modify_date: modify_date,
					Company_id: company_id,
					Organization_id: organization_id,
					Service_id: service_id, 
					Sub_org_id: sub_org_id,
					User_id: user_id}

				con_service_permissions = append(con_service_permissions, con_service_permission)
			}
		}
		return con_service_permissions, nil
	}
}


func (configurationModel ConfigurationModel) PostUserServicePermission(con_Service_Permission *entities.CON_Service_Permission) (err error) {

	result, err := configurationModel.Db.Exec(`INSERT INTO con_service_permission(service_Permission_id, service_id, user_id,
	 allow_add, allow_view, allow_delete, allow_approval, allow_print, active_status, delete_status, entry_doneby_loginid, 
	 entry_date, modify_login_id, modify_date, organization_id, sub_org_id, company_id) values(?, ?, ?, ?, ?, ?, ?, 
			?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`, con_Service_Permission.Service_Permission_id, con_Service_Permission.Service_id, 
			con_Service_Permission.User_id, con_Service_Permission.Allow_add, con_Service_Permission.Allow_view, 
			con_Service_Permission.Allow_delete, con_Service_Permission.Allow_approval, con_Service_Permission.Allow_print, 
			con_Service_Permission.Active_status, con_Service_Permission.Delete_status, con_Service_Permission.Entry_doneby_loginid, 
			con_Service_Permission.Entry_date, con_Service_Permission.Modify_login_id, con_Service_Permission.Modify_date, 
			con_Service_Permission.Organization_id, con_Service_Permission.Sub_org_id, con_Service_Permission.Company_id)

	if err != nil {
		return  err
	} else { 

		var id int64

		id, _ = result.LastInsertId()

		*con_Service_Permission.Service_Permission_id = id

		return nil
	}
}


func (configurationModel ConfigurationModel) GetUserServicePermForUser(id int64) (con_Service_Permission []entities.CON_Service_Permission, err error) {

	rows, err := configurationModel.Db.Query("SELECT * from con_service_permission where user_id = ?", id)

	if err != nil {
		return nil, err
	} else {

		var con_service_permissions []entities.CON_Service_Permission

		for rows.Next() {

			var service_Permission_id   *int64 
  			var service_id              *int64   
  			var user_id                 *int64 
  			var allow_add               bool    
  			var allow_view              bool      
  			var allow_delete            bool      
  			var allow_approval          bool    
  			var allow_print             bool     
  			var active_status           bool      
  			var delete_status           bool    
  			var entry_doneby_loginid   int64     
  			var entry_date             *string    
  			var modify_login_id        *int64     
  			var modify_date            *string    
  			var organization_id        *int64    
  			var sub_org_id             *int64   
  			var company_id             *int64  


			err2 := rows.Scan(&service_Permission_id, &allow_add, &allow_view, &allow_delete,
				&allow_approval, &allow_print, &active_status, &delete_status, &entry_doneby_loginid, &entry_date, 
				&modify_login_id, &modify_date,  &company_id, &organization_id, &service_id, &sub_org_id,  &user_id)

			if err2 != nil {
				return nil, err2
			} else {
				con_service_permission := entities.CON_Service_Permission{
					Service_Permission_id: service_Permission_id, 
					Allow_add: allow_add,
					Allow_view: allow_view,
					Allow_delete: allow_delete,
					Allow_approval: allow_approval,
					Allow_print: allow_print,
					Active_status: active_status,
					Delete_status: delete_status,
					Entry_doneby_loginid: entry_doneby_loginid,
					Entry_date: entry_date,
					Modify_login_id: modify_login_id,
					Modify_date: modify_date,
					Organization_id: organization_id,
					Service_id: service_id,
					Sub_org_id: sub_org_id,
					User_id: user_id, 
					Company_id: company_id}

				con_service_permissions = append(con_service_permissions, con_service_permission)
			}
		}
		return con_service_permissions, nil
	}
}


func (configurationModel ConfigurationModel) PutUserServicePermForUser(con_Service_Permission *entities.CON_Service_Permission) (int64, error) {

	result, err := configurationModel.Db.Exec(`UPDATE con_service_permission SET service_id = ?, user_id = ?, allow_add = ?, 
		allow_view = ?, allow_delete = ?, allow_approval = ?, allow_print = ?, active_status = ?, delete_status = ?, entry_doneby_loginid = ?, 
		entry_date = ?, modify_login_id = ?, modify_date = ?, organization_id = ?, sub_org_id = ?, company_id = ? 
		WHERE service_permission_id = ?`, con_Service_Permission.Service_id, 
			con_Service_Permission.User_id, con_Service_Permission.Allow_add, con_Service_Permission.Allow_view, 
			con_Service_Permission.Allow_delete, con_Service_Permission.Allow_approval, con_Service_Permission.Allow_print, 
			con_Service_Permission.Active_status, con_Service_Permission.Delete_status, con_Service_Permission.Entry_doneby_loginid, 
			con_Service_Permission.Entry_date, con_Service_Permission.Modify_login_id, con_Service_Permission.Modify_date, 
			con_Service_Permission.Organization_id, con_Service_Permission.Sub_org_id, con_Service_Permission.Company_id, 
			con_Service_Permission.Service_Permission_id)

	if err != nil {
		return  0, err
	} else { 

		return result.RowsAffected()
	}
}


// Delete is setting the is_active flag = false and not deleting the data
func (configurationModel ConfigurationModel) DelUserServicePermForUser(id int64) (int64, error) {

	result, err := configurationModel.Db.Exec(`UPDATE con_service_permission SET active_status = false 
		WHERE service_permission_id = ?`, id)

	if err != nil {
		return  0, err
	} else { 

		return result.RowsAffected()
	}
}