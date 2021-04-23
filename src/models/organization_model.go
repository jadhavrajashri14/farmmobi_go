package models

import (
	"database/sql"
	"entities"
)


type OrganizationModel struct {
	Db *sql.DB
}


func (organizationModel OrganizationModel) GetOrganizations() (organization []entities.Organization, err error) {

	rows, err := organizationModel.Db.Query("SELECT * from organization")

	if err != nil {
		return nil, err
	} else {

		var organizations []entities.Organization

		for rows.Next() {

			var organization_id uint64
		    var org_name string 
		    var org_reg_no string 
		    var no_of_suborgs uint64 
		    var subscrn_start_date  string  
		    var subscrn_end_date  string  
		    var on_trial bool 
		    var created_on  string  
		    var no_of_licenses uint64 
		    var subdomain_prefix string 
		    var address string 
		    var city string 
		    var district string 
		    var state string 
		    var country string  
		    var contact_no string 
		    var mobile_no_1 string 
		    var mobile_no_2 string 
		    var email string 
		    var is_active bool 


			err2 := rows.Scan(&organization_id, &org_name, &org_reg_no, &no_of_suborgs, &subscrn_start_date, &subscrn_end_date,
				&on_trial, &created_on, &no_of_licenses, &subdomain_prefix, &address, &city, &district, &state, &country,
				&contact_no, &mobile_no_1, &mobile_no_2, &email, &is_active)

			if err2 != nil {
				return nil, err2
			} else {
				organization := entities.Organization{
					Organization_id: organization_id, 
					Org_name: org_name, 
					Org_reg_no: org_reg_no, 
					No_of_suborgs: no_of_suborgs,
					Subscrn_start_date: subscrn_start_date,
					Subscrn_end_date: subscrn_end_date,
					On_trial: on_trial,
					Created_on: created_on,
					No_of_licenses: no_of_licenses,
					Subdomain_prefix: subdomain_prefix,
					Address: address,
					City: city,
					District: district,
					State: state,
					Country: country,
					Contact_no: contact_no,
					Mobile_no_1: mobile_no_1,
					Mobile_no_2: mobile_no_2,
					Email: email,
					Is_active: is_active}

				organizations = append(organizations, organization)
			}
		}
		return organizations, nil
	}
}


func (organizationModel OrganizationModel) Search(keyword string) (organization []entities.Organization, err error) {

	rows, err := organizationModel.Db.Query("SELECT * from organization where org_name like ?", "%"+keyword+"%")

	if err != nil {
		return nil, err
	} else {

		var organizations []entities.Organization

		for rows.Next() {

			var organization_id uint64
		    var org_name string 
		    var org_reg_no string 
		    var no_of_suborgs uint64 
		    var subscrn_start_date  string  
		    var subscrn_end_date  string  
		    var on_trial bool 
		    var created_on  string  
		    var no_of_licenses uint64 
		    var subdomain_prefix string 
		    var address string 
		    var city string 
		    var district string 
		    var state string 
		    var country string  
		    var contact_no string 
		    var mobile_no_1 string 
		    var mobile_no_2 string 
		    var email string 
		    var is_active bool 


			err2 := rows.Scan(&organization_id, &org_name, &org_reg_no, &no_of_suborgs, &subscrn_start_date, &subscrn_end_date,
				&on_trial, &created_on, &no_of_licenses, &subdomain_prefix, &address, &city, &district, &state, &country,
				&contact_no, &mobile_no_1, &mobile_no_2, &email, &is_active)

			if err2 != nil {
				return nil, err2
			} else {
				organization := entities.Organization{
					Organization_id: organization_id, 
					Org_name: org_name, 
					Org_reg_no: org_reg_no, 
					No_of_suborgs: no_of_suborgs,
					Subscrn_start_date: subscrn_start_date,
					Subscrn_end_date: subscrn_end_date,
					On_trial: on_trial,
					Created_on: created_on,
					No_of_licenses: no_of_licenses,
					Subdomain_prefix: subdomain_prefix,
					Address: address,
					City: city,
					District: district,
					State: state,
					Country: country,
					Contact_no: contact_no,
					Mobile_no_1: mobile_no_1,
					Mobile_no_2: mobile_no_2,
					Email: email,
					Is_active: is_active}

				organizations = append(organizations, organization)
			}
		}
		return organizations, nil
	}
}


func (organizationModel OrganizationModel) SearchOrgForYear(min, max string) (organization []entities.Organization, err error) {

	rows, err := organizationModel.Db.Query("SELECT * from organization where DATE(subscrn_start_date) >= ? and DATE(subscrn_end_date) <= ?", min, max)

	if err != nil {
		return nil, err
	} else {

		var organizations []entities.Organization

		for rows.Next() {

			var organization_id uint64
		    var org_name string 
		    var org_reg_no string 
		    var no_of_suborgs uint64 
		    var subscrn_start_date  string  
		    var subscrn_end_date  string  
		    var on_trial bool 
		    var created_on  string  
		    var no_of_licenses uint64 
		    var subdomain_prefix string 
		    var address string 
		    var city string 
		    var district string 
		    var state string 
		    var country string  
		    var contact_no string 
		    var mobile_no_1 string 
		    var mobile_no_2 string 
		    var email string 
		    var is_active bool 


			err2 := rows.Scan(&organization_id, &org_name, &org_reg_no, &no_of_suborgs, &subscrn_start_date, &subscrn_end_date,
				&on_trial, &created_on, &no_of_licenses, &subdomain_prefix, &address, &city, &district, &state, &country,
				&contact_no, &mobile_no_1, &mobile_no_2, &email, &is_active)

			if err2 != nil {
				return nil, err2
			} else {
				organization := entities.Organization{
					Organization_id: organization_id, 
					Org_name: org_name, 
					Org_reg_no: org_reg_no, 
					No_of_suborgs: no_of_suborgs,
					Subscrn_start_date: subscrn_start_date,
					Subscrn_end_date: subscrn_end_date,
					On_trial: on_trial,
					Created_on: created_on,
					No_of_licenses: no_of_licenses,
					Subdomain_prefix: subdomain_prefix,
					Address: address,
					City: city,
					District: district,
					State: state,
					Country: country,
					Contact_no: contact_no,
					Mobile_no_1: mobile_no_1,
					Mobile_no_2: mobile_no_2,
					Email: email,
					Is_active: is_active}

				organizations = append(organizations, organization)
			}
		}
		return organizations, nil
	}
}


func (organizationModel OrganizationModel) Create(organization *entities.Organization) (err error) {

	result, err := organizationModel.Db.Exec(`INSERT INTO organization(organization_id, org_name, org_reg_no, no_of_suborgs, 
		subscrn_start_date, subscrn_end_date, on_trial, created_on, no_of_licenses, subdomain_prefix, address, city, district, 
		state, country, contact_no, mobile_no_1, mobile_no_2, email, is_active) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, 
			?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`, organization.Organization_id, organization.Org_name, organization.Org_reg_no, 
			organization.No_of_suborgs, organization.Subscrn_start_date, organization.Subscrn_end_date, organization.On_trial, 
		organization.Created_on, organization.No_of_licenses, organization.Subdomain_prefix, organization.Address, 
		organization.City, organization.District, organization.State, organization.Country, organization.Contact_no, 
		organization.Mobile_no_1, organization.Mobile_no_2, organization.Email, organization.Is_active)

	if err != nil {
		return  err
	} else { 

		var id int64

		id, _ = result.LastInsertId()

		organization.Organization_id = uint64(id)

		return nil
	}
}


func (organizationModel OrganizationModel) Update(organization *entities.Organization) (int64, error) {

	result, err := organizationModel.Db.Exec(`UPDATE organization SET org_name = ?, org_reg_no = ?, no_of_suborgs = ?, 
		subscrn_start_date = ?, subscrn_end_date = ?, on_trial = ?, created_on = ?, no_of_licenses = ?, subdomain_prefix = ?, address = ?, 
		city = ?, district = ?, state = ?, country = ?, contact_no = ?, mobile_no_1 = ?, mobile_no_2 = ?, email = ?, is_active = ? 
		WHERE organization_id = ?`, organization.Org_name, organization.Org_reg_no, organization.No_of_suborgs, 
		organization.Subscrn_start_date, organization.Subscrn_end_date, organization.On_trial, 
		organization.Created_on, organization.No_of_licenses, organization.Subdomain_prefix, organization.Address, 
		organization.City, organization.District, organization.State, organization.Country, organization.Contact_no, 
		organization.Mobile_no_1, organization.Mobile_no_2, organization.Email, organization.Is_active, organization.Organization_id)

	if err != nil {
		return  0, err
	} else { 

		return result.RowsAffected()
	}
}


// Delete is setting the is_active flag = false and not deleting the data
func (organizationModel OrganizationModel) Delete(id uint64) (int64, error) {

	result, err := organizationModel.Db.Exec(`UPDATE organization SET is_active = false 
		WHERE organization_id = ?`, id)

	if err != nil {
		return  0, err
	} else { 

		return result.RowsAffected()
	}
}