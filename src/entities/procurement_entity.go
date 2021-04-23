package entities

import (
  "fmt"
)

type Farmer_Register struct {
  // Personal Details
  Farmer_id           uint64    `json:"farmer_id"`
  Farmer_name         string    `json:"farmer_name"`
  Farmer_code         string    `json:"farmer_code"`
  Firstname           string    `json:"firstname"`
  Middlename          string    `json:"middlename"`
  Lastname            string    `json:"lastname"`
  Gender              string    `json:"gender"`
  Farmer_photo        string    `json:"farmer_photo"`
  Country_id          uint64    `json:"country_id"`
  Village             string    `json:"village"`
  District            string    `json:"district"`
  Tehsil              string    `json:"tehsil"`
  State               string    `json:"state"`
  Birthdate           string    `json:"birthdate"`
  Contact_no          string    `json:"contact_no"`
  Mobile_no_1         string    `json:"mobile_no_1"`
  Mobile_no_2         string    `json:"mobile_no_2"`
  Emailaddress        string    `json:"emailaddress"`
  Total_landarea      uint64    `json:"total_landarea"`
  Land_address        string    `json:"land_address"`
  Pin_code            string    `json:"pin_code"`
  Latitude            float64   `json:"latitude"`
  Longitude           float64   `json:"longitude"`  
  Farmer_status       bool      `json:"farmer_status"`   // can be active or inactive
  // Farmer Registration Details
  Registration_date   string    `json:"registration_date"`
  Transaction_date    string    `json:"transaction_date"`
  Doc_no              string    `json:"doc_no"`
  Packhouse_no        string    `json:"packhouse_no"`
  Documented_by       uint64    `json:"documented_by"`
  Emp_name            string    `json:"emp_name"`   // Registration done by employee
  // Farmer Organization Details
  Organization_id     uint64    `json:"organization_id"`
  Sub_org_id          *int64    `json:"sub_org_id"`
  Company_id          *int64    `json:"company_id"`
  // Farmer Bank Account Details
  Bank_name           string    `json:"bank_name"`
  Branch_name         string    `json:"branch_name"`
  Branch_city         string    `json:"Branch_city"`
  Ifsc_or_iban_no     uint64    `json:"ifsc_or_iban_no"`
  Swift_code          uint64    `json:"swift_code"`
  Account_type        string    `json:"account_type"`
  Account_number      uint64    `json:"account_number"`
  Acct_status         string    `json:"acct_status"`
  // Farmer Identification Details
  Identification_type string    `json:"identification_type"`
  Identification_no   string    `json:"identification_no"`
  Identification_doc  string    `json:"identification_doc"`
  // Farmer Certification Details
  Certification_type  string    `json:"certification_type"`
  Certification_no    string    `json:"certification_no"`
  Certification_doc   string    `json:"certification_doc"`
}

func (farmer_register Farmer_Register) ToString() string {

	return fmt.Sprintf(`Farmer_id:%d\nFarmer_name:%s\nFarmer_code:%s\nFirstname:%s\nMiddlename:%s\nLastname:%s\n` +
			`Gender:%s\nFarmer_photo:%s\nCountry_id:%d\nVillage:%s\nDistrict:%s\nTehsil:%s\nState:%s\nBirthdate:%s\n` +
			`Contact_no:%s\nMobile_no_1:%s\nMobile_no_2:%s\nEmailaddress:%s\nTotal_landarea:%s\nLand_address:%s\n` +
			`Pin_code:%s\nLatitude:%s\nLongitude:%s\nFarmer_status:%s\nRegistration_date:%s\nTransaction_date:%s\n` + 
      `Doc_no:%s\nPackhouse_no:%s\nDocumented_by:%d\nEmp_name:%s\nOrganization_id:%d\nSub_org_id:%d\nCompany_id:%d\n` + 
      `Bank_name:%s\nBranch_name:%s\nBranch_city:%s\nIfsc_or_iban_no:%d\nSwift_code:%d\nAccount_type:%s\n` +
      `Account_number:%d\nAcct_status:%s\nIdentification_type:%s\nIdentification_no:%s\nIdentification_doc:%s\n` + 
      `Certification_type:%s\nCertification_no:%s\nCertification_doc:%s\n`,
			farmer_register.Farmer_id, farmer_register.Farmer_name, farmer_register.Farmer_code, farmer_register.Firstname,
      farmer_register.Middlename, farmer_register.Lastname, farmer_register.Gender, farmer_register.Farmer_photo,
      farmer_register.Country_id, farmer_register.Village, farmer_register.District, farmer_register.Tehsil,
      farmer_register.State, farmer_register.Birthdate, farmer_register.Contact_no, farmer_register.Mobile_no_1,
      farmer_register.Mobile_no_2, farmer_register.Emailaddress, farmer_register.Total_landarea, 
      farmer_register.Land_address, farmer_register.Pin_code, farmer_register.Latitude, farmer_register.Longitude,
      farmer_register.Farmer_status, farmer_register.Registration_date, farmer_register.Transaction_date,
      farmer_register.Doc_no, farmer_register.Packhouse_no, farmer_register.Documented_by, farmer_register.Emp_name, 
      farmer_register.Organization_id, farmer_register.Sub_org_id, farmer_register.Company_id, farmer_register.Bank_name, 
      farmer_register.Branch_name, farmer_register.Branch_city, farmer_register.Ifsc_or_iban_no, farmer_register.Swift_code, 
      farmer_register.Account_type, farmer_register.Account_number, farmer_register.Acct_status, farmer_register.Identification_type, 
      farmer_register.Identification_no, farmer_register.Identification_doc, farmer_register.Certification_type, 
      farmer_register.Certification_no, farmer_register.Certification_doc)
}