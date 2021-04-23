package entities

import (
  "fmt"
)


type CON_Service_Permission struct {
  Service_Permission_id   *int64      `json:"service_permission_id"`
  Service_id              *int64      `json:"service_id"`
  User_id                 *int64      `json:"user_id"`
  Allow_add               bool        `json:"allow_add"`
  Allow_view              bool        `json:"allow_view"`
  Allow_delete            bool        `json:"allow_delete"`
  Allow_approval          bool        `json:"allow_approval"`
  Allow_print             bool        `json:"allow_print"`
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


func (con_service_permission CON_Service_Permission) ToString() string {

  return fmt.Sprintf(`Service_Permission_id:%d\nService_id:%d\nUser_id:%d\nAllow_add:%s\nAllow_view:%s\nAllow_delete:%s\n` +
      `Allow_approval:%s\nAllow_print:%s\nActive_status:%s\nDelete_status:%s\nEntry_doneby_loginid:%d\nEntry_date:%s\n` +
      `Modify_login_id:%d\nModify_date:%s\nOrganization_id:%d\nSub_org_id:%d\nCompany_id:%d\n`,
      con_service_permission.Service_Permission_id, con_service_permission.Service_id, con_service_permission.User_id, 
      con_service_permission.Allow_add, con_service_permission.Allow_view, con_service_permission.Allow_delete,
      con_service_permission.Allow_approval, con_service_permission.Allow_print, con_service_permission.Active_status, 
      con_service_permission.Delete_status, con_service_permission.Entry_doneby_loginid, con_service_permission.Entry_date,
      con_service_permission.Modify_login_id, con_service_permission.Modify_date, con_service_permission.Organization_id,
      con_service_permission.Sub_org_id, con_service_permission.Company_id)
}

