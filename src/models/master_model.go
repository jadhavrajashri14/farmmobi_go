package models

import (
	"database/sql"
	"entities"
)


type MasterModel struct {
	Db *sql.DB
}


func (masterModel MasterModel) GetCountries() (country []entities.Country, err error) {

	rows, err := masterModel.Db.Query("SELECT * from country")

	if err != nil {
		return nil, err
	} else {

		var countries []entities.Country

		for rows.Next() {

			 var country_code            string 
			 var country_name            string      
			 var code   		         string     
			 

			err2 := rows.Scan(&country_code, &country_name, &code)

			if err2 != nil {
				return nil, err2
			} else {
				country := entities.Country{
					Country_code: country_code, 
					Country_name: country_name,
					Code: code}

				countries = append(countries, country)
			}
		}
		return countries, nil
	}
}


func (masterModel MasterModel) GetConModules() (con_module []entities.CON_Module, err error) {

	rows, err := masterModel.Db.Query("SELECT * from con_module")

	if err != nil {
		return nil, err
	} else {

		var con_modules []entities.CON_Module

		for rows.Next() {

			 var module_id               *int64 
			 var module_name             string      
			 var module_code             string     
			 var module_url              string      
			 var active_status           bool        
			 var delete_status           bool       
			 var entry_doneby_loginid   int64       
			 var entry_date             *string     
			 var modify_login_id        *int64      
			 var modify_date            *string      
			 var organization_id        *int64      
			 var sub_org_id             *int64      
			 var company_id             *int64    


			err2 := rows.Scan(&module_id, &module_name, &module_code, &module_url, &active_status, 
				&delete_status, &entry_doneby_loginid, &entry_date, &modify_login_id, &modify_date, 
				&organization_id, &sub_org_id, &company_id)

			if err2 != nil {
				return nil, err2
			} else {
				con_module := entities.CON_Module{
					Module_id: module_id, 
					Module_name: module_name,
					Module_code: module_code,
					Module_url: module_url,
					Active_status: active_status,
					Delete_status: delete_status,
					Entry_doneby_loginid: entry_doneby_loginid,
					Entry_date: entry_date,
					Modify_login_id: modify_login_id,
					Modify_date: modify_date,
					Organization_id: organization_id,
					Sub_org_id: sub_org_id,
					Company_id: company_id}

				con_modules = append(con_modules, con_module)
			}
		}
		return con_modules, nil
	}
}


func (masterModel MasterModel) PostConModules(con_module *entities.CON_Module) (err error) {

	result, err := masterModel.Db.Exec(`INSERT INTO con_module(module_id, module_name, module_code, module_url, 
	 active_status, delete_status, entry_doneby_loginid, entry_date, modify_login_id, modify_date, 
	 organization_id, sub_org_id, company_id) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`, 
	 con_module.Module_id, con_module.Module_name, con_module.Module_code, con_module.Module_url, 
			con_module.Active_status, con_module.Delete_status, con_module.Entry_doneby_loginid, 
			con_module.Entry_date, con_module.Modify_login_id, con_module.Modify_date, 
			con_module.Organization_id, con_module.Sub_org_id, con_module.Company_id)

	if err != nil {
		return  err
	} else { 

		var id int64

		id, _ = result.LastInsertId()

		*con_module.Module_id = id

		return nil
	}
}


func (masterModel MasterModel) GetConModule(id int64) (con_module []entities.CON_Module, err error) {

	rows, err := masterModel.Db.Query("SELECT * from con_module where module_id = ?", id)

	if err != nil {
		return nil, err
	} else {

		var con_modules []entities.CON_Module

		for rows.Next() {

			 var module_id               *int64 
			 var module_name             string      
			 var module_code             string     
			 var module_url              string      
			 var active_status           bool        
			 var delete_status           bool       
			 var entry_doneby_loginid   int64       
			 var entry_date             *string     
			 var modify_login_id        *int64      
			 var modify_date            *string      
			 var organization_id        *int64      
			 var sub_org_id             *int64      
			 var company_id             *int64    

			err2 := rows.Scan(&module_id, &module_name, &module_code, &module_url, &active_status, 
				&delete_status, &entry_doneby_loginid, &entry_date, &modify_login_id, &modify_date, 
				&organization_id, &sub_org_id,  &company_id)

			if err2 != nil {
				return nil, err2
			} else {

				con_module := entities.CON_Module{
					Module_id: module_id, 
					Module_name: module_name,
					Module_code: module_code,
					Module_url: module_url,
					Active_status: active_status,
					Delete_status: delete_status,
					Entry_doneby_loginid: entry_doneby_loginid,
					Entry_date: entry_date,
					Modify_login_id: modify_login_id,
					Modify_date: modify_date,
					Organization_id: organization_id,
					Sub_org_id: sub_org_id,
					Company_id: company_id}

				con_modules = append(con_modules, con_module)
			}
		}
		return con_modules, nil
	}
}


func (masterModel MasterModel) PutConModule(con_module *entities.CON_Module) (int64, error) {

	result, err := masterModel.Db.Exec(`UPDATE con_module SET module_id = ?, module_name = ?, module_code = ?, 
		module_url = ?, active_status = ?, delete_status = ?, entry_doneby_loginid = ?, 
		entry_date = ?, modify_login_id = ?, modify_date = ?, organization_id = ?, sub_org_id = ?, company_id = ? 
		WHERE module_id = ?`, con_module.Module_id, 
			con_module.Module_name, con_module.Module_code, con_module.Module_url, 
			con_module.Active_status, con_module.Delete_status, con_module.Entry_doneby_loginid, 
			con_module.Entry_date, con_module.Modify_login_id, con_module.Modify_date, 
			con_module.Organization_id, con_module.Sub_org_id, con_module.Company_id, con_module.Module_id)

	if err != nil {
		return  0, err
	} else { 

		return result.RowsAffected()
	}
}


// Delete is setting the is_active flag = false and not deleting the data
func (masterModel MasterModel) DelConModule(id int64) (int64, error) {

	result, err := masterModel.Db.Exec(`UPDATE con_module SET active_status = false WHERE module_id = ?`, id)

	if err != nil {
		return  0, err
	} else { 

		return result.RowsAffected()
	}
}


func (masterModel MasterModel) GetConServices() (con_service []entities.CON_Service, err error) {

	rows, err := masterModel.Db.Query("SELECT * from con_service")

	if err != nil {
		return nil, err
	} else {

		var con_services []entities.CON_Service

		for rows.Next() {
			 var service_id              *int64  
			 var module_id               *int64 
			 var service_name             string      
			 var service_code             string     
			 var service_url              string   
			 var is_avl_on_mob            bool          
			 var active_status           bool        
			 var delete_status           bool       
			 var entry_doneby_loginid   int64       
			 var entry_date             *string     
			 var modify_login_id        *int64      
			 var modify_date            *string      
			 var organization_id        *int64      
			 var sub_org_id             *int64      
			 var company_id             *int64    


			err2 := rows.Scan(&service_id, &service_name, &service_code, &service_url, &is_avl_on_mob,
			 &active_status, &delete_status, &entry_doneby_loginid, &entry_date, &modify_login_id, &modify_date, 
				&company_id,  &module_id, &organization_id, &sub_org_id)

			if err2 != nil {
				return nil, err2
			} else {

				con_service := entities.CON_Service{
					Service_id: service_id,
					Module_id: module_id, 
					Service_name: service_name,
					Service_code: service_code,
					Service_url: service_url,
					Is_avl_on_mob: is_avl_on_mob,
					Active_status: active_status,
					Delete_status: delete_status,
					Entry_doneby_loginid: entry_doneby_loginid,
					Entry_date: entry_date,
					Modify_login_id: modify_login_id,
					Modify_date: modify_date,
					Organization_id: organization_id,
					Sub_org_id: sub_org_id,
					Company_id: company_id}

				con_services = append(con_services, con_service)
			}
		}
		return con_services, nil
	}
}


func (masterModel MasterModel) PostConServices(con_service *entities.CON_Service) (err error) {

	result, err := masterModel.Db.Exec(`INSERT INTO con_service(service_id, module_id, service_name, service_code, 
		service_url, is_avl_on_mob, active_status, delete_status, entry_doneby_loginid, entry_date, modify_login_id, 
		modify_date, organization_id, sub_org_id, company_id) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`, 
	 con_service.Service_id, con_service.Module_id, con_service.Service_name, con_service.Service_code, 
	 con_service.Service_url, con_service.Is_avl_on_mob, con_service.Active_status, con_service.Delete_status, 
	 con_service.Entry_doneby_loginid, con_service.Entry_date, con_service.Modify_login_id, con_service.Modify_date, 
	 con_service.Organization_id, con_service.Sub_org_id, con_service.Company_id)

	if err != nil {
		return  err
	} else { 

		var id int64

		id, _ = result.LastInsertId()

		*con_service.Service_id = id

		return nil
	}
}


func (masterModel MasterModel) GetConService(id int64) (con_service []entities.CON_Service, err error) {

	rows, err := masterModel.Db.Query("SELECT * from con_service where service_id = ?", id)

	if err != nil {
		return nil, err
	} else {

		var con_services []entities.CON_Service

		for rows.Next() {

			 var service_id              *int64  
			 var module_id               *int64 
			 var service_name             string      
			 var service_code             string     
			 var service_url              string   
			 var is_avl_on_mob            bool      
			 var active_status           bool        
			 var delete_status           bool       
			 var entry_doneby_loginid   int64       
			 var entry_date             *string     
			 var modify_login_id        *int64      
			 var modify_date            *string      
			 var organization_id        *int64      
			 var sub_org_id             *int64      
			 var company_id             *int64    

			err2 := rows.Scan(&service_id, &service_name, &service_code, &service_url, &is_avl_on_mob,
			 &active_status, &delete_status, &entry_doneby_loginid, &entry_date, &modify_login_id, &modify_date, 
				&company_id,  &module_id, &organization_id, &sub_org_id)

			if err2 != nil {
				return nil, err2
			} else {

				con_service := entities.CON_Service{
					Service_id: service_id,
					Module_id: module_id, 
					Service_name: service_name,
					Service_code: service_code,
					Service_url: service_url,
					Is_avl_on_mob: is_avl_on_mob,
					Active_status: active_status,
					Delete_status: delete_status,
					Entry_doneby_loginid: entry_doneby_loginid,
					Entry_date: entry_date,
					Modify_login_id: modify_login_id,
					Modify_date: modify_date,
					Organization_id: organization_id,
					Sub_org_id: sub_org_id,
					Company_id: company_id}

				con_services = append(con_services, con_service)
			}
		}
		return con_services, nil
	}
}


func (masterModel MasterModel) PutConService(con_service *entities.CON_Service) (int64, error) {

	result, err := masterModel.Db.Exec(`UPDATE con_service SET module_id = ?, service_name = ?, 
		service_code = ?, service_url = ?, is_avl_on_mob = ?, active_status = ?, delete_status = ?, 
		entry_doneby_loginid = ?, entry_date = ?, modify_login_id = ?, modify_date = ?, organization_id = ?, 
		sub_org_id = ?, company_id = ? WHERE service_id = ?`, con_service.Module_id, con_service.Service_name, 
		con_service.Service_code, con_service.Service_url, con_service.Is_avl_on_mob, con_service.Active_status, 
		con_service.Delete_status, con_service.Entry_doneby_loginid, con_service.Entry_date, 
		con_service.Modify_login_id, con_service.Modify_date, con_service.Organization_id, 
		con_service.Sub_org_id, con_service.Company_id, con_service.Service_id)

	if err != nil {
		return  0, err
	} else { 

		return result.RowsAffected()
	}
}


// Delete is setting the is_active flag = false and not deleting the data
func (masterModel MasterModel) DelConService(id int64) (int64, error) {

	result, err := masterModel.Db.Exec(`UPDATE con_service SET active_status = false WHERE service_id = ?`, id)

	if err != nil {
		return  0, err
	} else { 

		return result.RowsAffected()
	}
}