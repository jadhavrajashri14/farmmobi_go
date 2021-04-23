package entities

import (
  "fmt"
)

type Organization struct {
  Organization_id     uint64    `json:"organization_id"`
  Org_name            string    `json:"org_name"`
  Org_reg_no          string    `json:"org_reg_no"`
  No_of_suborgs       uint64    `json:"no_of_suborgs"`
  Subscrn_start_date  string    `json:"subscrn_start_date"`
  Subscrn_end_date    string    `json:"subscrn_end_date"`
  On_trial            bool      `json:"on_trial"`
  Created_on          string    `json:"created_on"`
  No_of_licenses      uint64    `json:"no_of_licenses"`
  Subdomain_prefix    string    `json:"subdomain_prefix"`
  Address             string    `json:"address"`
  City                string    `json:"city"`
  District            string    `json:"district"`
  State               string    `json:"state"`
  Country             string    `json:"country"`
  Contact_no          string    `json:"contact_no"`
  Mobile_no_1         string    `json:"mobile_no_1"`
  Mobile_no_2         string    `json:"mobile_no_2"`
  Email               string    `json:"email"`
  Is_active           bool      `json:"is_active"`
}

func (organization Organization) ToString() string {

	return fmt.Sprintf(`Organization_id:%d\nOrg_name:%s\nOrg_reg_no:%s\nNo_of_suborgs:%d\n` +
			`Subscrn_start_date:%s\nSubscrn_end_date:%s\nOn_trial:%s\nCreated_on:%s\n` +
			`No_of_licenses:%d\nSubdomain_prefix:%s\nAddress:%s\nCity:%s\nDistrict:%s\n` +
			`State:%s\nCountry:%s\nContact_no:%s\nMobile_no_1:%s\nMobile_no_2:%s\nEmail:%s\nIs_active:%s\n`,
			organization.Organization_id, organization.Org_name, organization.Org_reg_no, organization.No_of_suborgs,
			organization.Subscrn_start_date, organization.Subscrn_end_date, organization.On_trial, organization.Created_on,
			organization.No_of_licenses, organization.Subdomain_prefix, organization.Address, organization.City, 
			organization.District,organization.State, organization.Country, organization.Contact_no, organization.Mobile_no_1,
			organization.Mobile_no_2, organization.Email, organization.Is_active)
}