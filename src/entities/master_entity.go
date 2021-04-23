package entities

import (
  "fmt"
)

type Country struct {
  Country_code            string      `json:"country_code"`
  Country_name            string      `json:"country_name"`
  Code                    string      `json:"code"`
}

type CON_Module struct {
  Module_id               *int64      `json:"module_id"`
  Module_name             string      `json:"module_name"`
  Module_code             string      `json:"module_code"`
  Module_url              string      `json:"module_url"`
  Active_status           bool        `json:"active_status"`
  Delete_status           bool        `json:"delete_status"`
  Entry_doneby_loginid   int64        `json:"entry_doneby_loginid"`
  Entry_date             *string      `json:"entry_date"`
  Modify_login_id        *int64       `json:"modify_login_id"`
  Modify_date            *string      `json:"modify_date"`
  Organization_id        *int64       `json:"organization_id"`
  Sub_org_id             *int64       `json:"sub_org_id"`
  Company_id             *int64       `json:"company_id"`
}

type CON_Service struct {
  Service_id              *int64      `json:"service_id"`
  Module_id               *int64      `json:"module_id"`
  Service_name             string     `json:"service_name"`
  Service_code             string     `json:"service_code"`
  Service_url              string     `json:"service_url"`
  Is_avl_on_mob            bool       `json:"is_avl_on_mob"`
  Active_status           bool        `json:"active_status"`
  Delete_status           bool        `json:"delete_status"`
  Entry_doneby_loginid   int64        `json:"entry_doneby_loginid"`
  Entry_date             *string      `json:"entry_date"`
  Modify_login_id        *int64       `json:"modify_login_id"`
  Modify_date            *string      `json:"modify_date"`
  Organization_id        *int64       `json:"organization_id"`
  Sub_org_id             *int64       `json:"sub_org_id"`
  Company_id             *int64       `json:"company_id"`
}


func (country Country) ToString() string {

  return fmt.Sprintf(`Country_code:%s\nCountry_name:%s\nCode:%s\n`,
      country.Country_code, country.Country_name, country.Code)
}


func (con_module CON_Module) ToString() string {

  return fmt.Sprintf(`Module_id:%d\nModule_name:%s\nModule_code:%s\nModule_url:%s\nActive_status:%s\n` +
      `Delete_status:%s\nEntry_doneby_loginid:%d\nEntry_date:%s\nModify_login_id:%d\n` +
      `Modify_date:%s\nOrganization_id:%d\nSub_org_id:%d\nCompany_id:%d\n`,
      con_module.Module_id, con_module.Module_name, con_module.Module_code, con_module.Module_url, 
      con_module.Active_status, con_module.Delete_status, con_module.Entry_doneby_loginid, con_module.Entry_date,
      con_module.Modify_login_id, con_module.Modify_date, con_module.Organization_id,
      con_module.Sub_org_id, con_module.Company_id)
}


func (con_service CON_Service) ToString() string {

  return fmt.Sprintf(`Service_id:%d\nModule_id:%d\nvService_name:%s\nService_code:%s\nService_url:%s\n` +
      `Is_avl_on_mob:%s\nActive_status:%s\nDelete_status:%s\nEntry_doneby_loginid:%d\nEntry_date:%s\n` +
      `Modify_login_id:%d\nModify_date:%s\nOrganization_id:%d\nSub_org_id:%d\nCompany_id:%d\n`,
      con_service.Service_id, con_service.Module_id, con_service.Service_name, 
      con_service.Service_code, con_service.Service_url, con_service.Is_avl_on_mob,
      con_service.Active_status, con_service.Delete_status, con_service.Entry_doneby_loginid, con_service.Entry_date,
      con_service.Modify_login_id, con_service.Modify_date, con_service.Organization_id,
      con_service.Sub_org_id, con_service.Company_id)
}
