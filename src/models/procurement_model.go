package models

import (
	"database/sql"
	"entities"
)


type ProcurementModel struct {
	Db *sql.DB
}


func (procurementModel ProcurementModel) GetFarmers() (farmer_register []entities.Farmer_Register, err error) {

	rows, err := procurementModel.Db.Query("SELECT * from farmer_registration")

	if err != nil {
		return nil, err
	} else {

		var farmers []entities.Farmer_Register

		for rows.Next() {

			// Personal Details
		  var farmer_id           uint64    
		  var farmer_name         string    
		  var farmer_code         string   
		  var firstname           string    
		  var middlename          string    
		  var lastname            string   
		  var gender              string   
		  var farmer_photo        string   
		  var country_id          uint64   
		  var village             string  
		  var district            string   
		  var tehsil              string   
		  var state               string   
		  var birthdate           string    
		  var contact_no          string   
		  var mobile_no_1         string    
		  var mobile_no_2         string   
		  var emailaddress        string   
		  var total_landarea      uint64  
		  var land_address        string    
		  var pin_code            string   
		  var latitude            float64   
		  var longitude           float64  
		  var farmer_status       bool         // can be active or inactive
		  // Farmer Registration Details
		  var registration_date   string    
		  var transaction_date    string   
		  var doc_no              string  
		  var packhouse_no        string   
		  var documented_by       uint64   
		  var emp_name            string     // Registration done by employee
		  // Farmer Organization Details
		  var organization_id     uint64    
		  var sub_org_id          *int64   
		  var company_id          *int64 
		  // Farmer Bank Account Details
		  var bank_name           string   
		  var branch_name         string    
		  var branch_city         string   
		  var ifsc_or_iban_no     uint64   
		  var swift_code          uint64   
		  var account_type        string   
		  var account_number      uint64  
		  var acct_status         string 
		  // Farmer Identification Details
		  var identification_type string   
		  var identification_no   string   
		  var identification_doc  string   
		  // Farmer Certification Details
		  var certification_type  string  
		  var certification_no    string  
		  var certification_doc   string  


			err2 := rows.Scan(&farmer_id, &farmer_name, &farmer_code, &firstname, &middlename, &lastname, &gender, &farmer_photo, 
				&country_id, &village, &district, &tehsil, &state, &birthdate, &contact_no, &mobile_no_1, &mobile_no_2, &emailaddress, 
				&total_landarea, &land_address, &pin_code, &latitude, &longitude, &farmer_status, &registration_date, &transaction_date, 
				&doc_no, &packhouse_no, &documented_by, &emp_name, &organization_id, &sub_org_id, &company_id, &bank_name, &branch_name,
				 &branch_city, &ifsc_or_iban_no, &swift_code, &account_type, &account_number, &acct_status, &identification_type, 
				 &identification_no, &identification_doc, &certification_type, &certification_no, &certification_doc)

			if err2 != nil {
				return nil, err2
			} else {
				farmer_register := entities.Farmer_Register{
					// Personal Details
				  Farmer_id: farmer_id,
				  Farmer_name: farmer_name,
				  Farmer_code: farmer_code,
				  Firstname: firstname,
				  Middlename: middlename,
				  Lastname: lastname,
				  Gender: gender,
				  Farmer_photo: farmer_photo,
				  Country_id: country_id,
				  Village: village,
				  District: district,
				  Tehsil: tehsil,
				  State: state,
				  Birthdate: birthdate,
				  Contact_no: contact_no,
				  Mobile_no_1: mobile_no_1,
				  Mobile_no_2: mobile_no_2,
				  Emailaddress: emailaddress,
				  Total_landarea: total_landarea,
				  Land_address: land_address,
				  Pin_code: pin_code,
				  Latitude: latitude,
				  Longitude: longitude,
				  Farmer_status: farmer_status,  // can be active or inactive
				  // Farmer Registration Details
				  Registration_date: registration_date,
				  Transaction_date: transaction_date,
				  Doc_no: doc_no,
				  Packhouse_no: packhouse_no,
				  Documented_by: documented_by,
				  Emp_name: emp_name,   // Registration done by employee
				  // Farmer Organization Details
				  Organization_id: organization_id,
				  Sub_org_id: sub_org_id,
				  Company_id: company_id,
				  // Farmer Bank Account Details
				  Bank_name: bank_name,
				  Branch_name: branch_city,
				  Ifsc_or_iban_no: ifsc_or_iban_no,
				  Swift_code: swift_code,
				  Account_type: account_type,
				  Account_number: account_number,
				  Acct_status: acct_status,
				  // Farmer Identification Details
				  Identification_type: identification_type,
				  Identification_no: identification_no,
				  Identification_doc: identification_doc,
				  // Farmer Certification Details
				  Certification_type: certification_type,
				  Certification_no: certification_no,
				  Certification_doc: certification_doc}


				farmers = append(farmers, farmer_register)
			}
		}
		return farmers, nil
	}
}


func (procurementModel ProcurementModel) GetFarmer(keyword string) (farmer_register []entities.Farmer_Register, err error) {

	rows, err := procurementModel.Db.Query("SELECT * from farmer_registration where farmer_name like ?", "%"+keyword+"%")

	if err != nil {
		return nil, err
	} else {

		var farmers []entities.Farmer_Register

		for rows.Next() {

			// Personal Details
		  var farmer_id           uint64    
		  var farmer_name         string    
		  var farmer_code         string   
		  var firstname           string    
		  var middlename          string    
		  var lastname            string   
		  var gender              string   
		  var farmer_photo        string   
		  var country_id          uint64   
		  var village             string  
		  var district            string   
		  var tehsil              string   
		  var state               string   
		  var birthdate           string    
		  var contact_no          string   
		  var mobile_no_1         string    
		  var mobile_no_2         string   
		  var emailaddress        string   
		  var total_landarea      uint64  
		  var land_address        string    
		  var pin_code            string   
		  var latitude            float64   
		  var longitude           float64  
		  var farmer_status       bool         // can be active or inactive
		  // Farmer Registration Details
		  var registration_date   string    
		  var transaction_date    string   
		  var doc_no              string  
		  var packhouse_no        string   
		  var documented_by       uint64   
		  var emp_name            string     // Registration done by employee
		  // Farmer Organization Details
		  var organization_id     uint64    
		  var sub_org_id          *int64   
		  var company_id          *int64 
		  // Farmer Bank Account Details
		  var bank_name           string   
		  var branch_name         string    
		  var branch_city         string   
		  var ifsc_or_iban_no     uint64   
		  var swift_code          uint64   
		  var account_type        string   
		  var account_number      uint64  
		  var acct_status         string 
		  // Farmer Identification Details
		  var identification_type string   
		  var identification_no   string   
		  var identification_doc  string   
		  // Farmer Certification Details
		  var certification_type  string  
		  var certification_no    string  
		  var certification_doc   string  


			err2 := rows.Scan(&farmer_id, &farmer_name, &farmer_code, &firstname, &middlename, &lastname, &gender, &farmer_photo, 
				&country_id, &village, &district, &tehsil, &state, &birthdate, &contact_no, &mobile_no_1, &mobile_no_2, &emailaddress, 
				&total_landarea, &land_address, &pin_code, &latitude, &longitude, &farmer_status, &registration_date, &transaction_date, 
				&doc_no, &packhouse_no, &documented_by, &emp_name, &organization_id, &sub_org_id, &company_id, &bank_name, &branch_name, 
				&branch_city, &ifsc_or_iban_no, &swift_code, &account_type, &account_number, &acct_status, &identification_type, 
				&identification_no, &identification_doc, &certification_type, &certification_no, &certification_doc)

			if err2 != nil {
				return nil, err2
			} else {
				farmer_register := entities.Farmer_Register{

					// Personal Details
				  Farmer_id: farmer_id,
				  Farmer_name: farmer_name,
				  Farmer_code: farmer_code,
				  Firstname: firstname,
				  Middlename: middlename,
				  Lastname: lastname,
				  Gender: gender,
				  Farmer_photo: farmer_photo,
				  Country_id: country_id,
				  Village: village,
				  District: district,
				  Tehsil: tehsil,
				  State: state,
				  Birthdate: birthdate,
				  Contact_no: contact_no,
				  Mobile_no_1: mobile_no_1,
				  Mobile_no_2: mobile_no_2,
				  Emailaddress: emailaddress,
				  Total_landarea: total_landarea,
				  Land_address: land_address,
				  Pin_code: pin_code,
				  Latitude: latitude,
				  Longitude: longitude,
				  Farmer_status: farmer_status, // can be active or inactive
				  // Farmer Registration Details
				  Registration_date: registration_date,
				  Transaction_date: transaction_date,
				  Doc_no: doc_no,
				  Packhouse_no: packhouse_no,
				  Documented_by: documented_by,
				  Emp_name: emp_name,  // Registration done by employee
				  // Farmer Organization Details
				  Organization_id: organization_id,
				  Sub_org_id: sub_org_id,
				  Company_id: company_id,
				  // Farmer Bank Account Details
				  Bank_name: bank_name,
				  Branch_name: branch_city,
				  Ifsc_or_iban_no: ifsc_or_iban_no,
				  Swift_code: swift_code,
				  Account_type: account_type,
				  Account_number: account_number,
				  Acct_status: acct_status,
				  // Farmer Identification Details
				  Identification_type: identification_type,
				  Identification_no: identification_no,
				  Identification_doc: identification_doc,
				  // Farmer Certification Details
				  Certification_type: certification_type,
				  Certification_no: certification_no,
				  Certification_doc: certification_doc}


				farmers = append(farmers, farmer_register)
			}
		}
		return farmers, nil
	}
}


func (procurementModel ProcurementModel) GetMaxFarmerId() (farmer_id *int64, err error) {

	rows, err := procurementModel.Db.Query("SELECT max(farmer_id) from farmer_registration")

	if err != nil {
		return nil, err
	} else {

		for rows.Next() {

			var  farmer_id  int64
			
			err2 := rows.Scan(&farmer_id)

			if err2 != nil {
				return nil, err2
			} else {

				return &farmer_id, nil
			}
		}
		return nil, err
	}
}


func (procurementModel ProcurementModel) CreateFarmer(farmer_register *entities.Farmer_Register) (err error) {

	result, err := procurementModel.Db.Exec(`INSERT INTO farmer_registration(farmer_id, farmer_name, farmer_code, firstname, 
				middlename, lastname, gender, farmer_photo, country_id, village, district, tehsil, state, birthdate, contact_no,
				mobile_no_1, mobile_no_2, emailaddress, total_landarea, land_address, pin_code, latitude, longitude, farmer_status, 
				registration_date, transaction_date, doc_no, packhouse_no, documented_by, emp_name, organization_id, sub_org_id, 
				company_id, bank_name, branch_city, ifsc_or_iban_no, swift_code, account_type, account_number, acct_status, 
				identification_type, identification_no, identification_doc, certification_type, certification_no, certification_doc) 
				values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, 
					?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
					farmer_register.Farmer_id, farmer_register.Farmer_name, farmer_register.Farmer_code, farmer_register.Firstname, 
					farmer_register.Middlename, farmer_register.Lastname, farmer_register.Gender, farmer_register.Farmer_photo, 
					farmer_register.Country_id, farmer_register.Village, farmer_register.District, farmer_register.Tehsil,
					farmer_register.State, farmer_register.Birthdate, farmer_register.Contact_no, farmer_register.Mobile_no_1, 
					farmer_register.Mobile_no_2, farmer_register.Emailaddress, farmer_register.Total_landarea, farmer_register.Land_address,
					farmer_register.Pin_code, farmer_register.Latitude, farmer_register.Longitude, farmer_register.Farmer_status, 
					farmer_register.Registration_date, farmer_register.Transaction_date, farmer_register.Doc_no, farmer_register.Packhouse_no,
					farmer_register.Documented_by, farmer_register.Emp_name, farmer_register.Organization_id, farmer_register.Sub_org_id,
					farmer_register.Company_id, farmer_register.Bank_name, farmer_register.Branch_name, farmer_register.Ifsc_or_iban_no, 
					farmer_register.Swift_code, farmer_register.Account_type, farmer_register.Account_number, farmer_register.Acct_status, 
					farmer_register.Identification_type, farmer_register.Identification_no, farmer_register.Identification_doc, 
					farmer_register.Certification_type, farmer_register.Certification_no, farmer_register.Certification_doc)


	if err != nil {
		return  err
	} else { 

		var id int64

		id, _ = result.LastInsertId()

		farmer_register.Farmer_id = uint64(id)

		return nil
	}
}


func (procurementModel ProcurementModel) UpdateFarmer(farmer_register *entities.Farmer_Register) (int64, error) {

	result, err := procurementModel.Db.Exec(`UPDATE farmer_registration SET farmer_name = ?, farmer_code = ?, firstname = ?, 
				middlename = ?, lastname = ?, gender = ?, farmer_photo = ?, country_id = ?, village = ?, district = ?, tehsil = ?,
				state = ?, birthdate = ?, contact_no = ?, mobile_no_1 = ?, mobile_no_2 = ?, emailaddress = ?, 
				total_landarea = ?, land_address = ?, pin_code = ?, latitude = ?, longitude = ?, farmer_status = ?, 
				registration_date = ?, transaction_date = ?, doc_no = ?, packhouse_no = ?, documented_by = ?, emp_name = ?, 
				organization_id = ?, sub_org_id = ?, company_id = ?, bank_name = ?, branch_city = ?, ifsc_or_iban_no = ?, 
				swift_code = ?, account_type = ?, account_number = ?, acct_status = ?, identification_type = ?, 
				identification_no = ?, identification_doc = ?, certification_type = ?, certification_no = ?, certification_doc = ?
				WHERE farmer_id = ?`, farmer_register.Farmer_name, farmer_register.Farmer_code, farmer_register.Firstname, 
					farmer_register.Middlename, farmer_register.Lastname, farmer_register.Gender, farmer_register.Farmer_photo, 
					farmer_register.Country_id, farmer_register.Village, farmer_register.District, farmer_register.Tehsil,
					farmer_register.State, farmer_register.Birthdate, farmer_register.Contact_no, farmer_register.Mobile_no_1, 
					farmer_register.Mobile_no_2, farmer_register.Emailaddress, farmer_register.Total_landarea, farmer_register.Land_address,
					farmer_register.Pin_code, farmer_register.Latitude, farmer_register.Longitude, farmer_register.Farmer_status, 
					farmer_register.Registration_date, farmer_register.Transaction_date, farmer_register.Doc_no, farmer_register.Packhouse_no,
					farmer_register.Documented_by, farmer_register.Emp_name, farmer_register.Organization_id, farmer_register.Sub_org_id,
					farmer_register.Company_id, farmer_register.Bank_name, farmer_register.Branch_name, farmer_register.Ifsc_or_iban_no, 
					farmer_register.Swift_code, farmer_register.Account_type, farmer_register.Account_number, farmer_register.Acct_status, 
					farmer_register.Identification_type, farmer_register.Identification_no, farmer_register.Identification_doc, 
					farmer_register.Certification_type, farmer_register.Certification_no, farmer_register.Certification_doc, 
					farmer_register.Farmer_id)

	if err != nil {
		return  0, err
	} else { 

		return result.RowsAffected()
	}
}


// Delete is setting the is_active flag = false and not deleting the data
func (procurementModel ProcurementModel) DeleteFarmer(id uint64) (int64, error) {

	result, err := procurementModel.Db.Exec(`UPDATE farmer_registration SET farmer_status = false WHERE farmer_id = ?`, id)

	if err != nil {
		return  0, err
	} else { 

		return result.RowsAffected()
	}
}