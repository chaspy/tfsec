package database

import (
	"github.com/aquasecurity/defsec/rules"
	"github.com/aquasecurity/defsec/rules/azure/database"
	"github.com/aquasecurity/tfsec/internal/app/tfsec/block"
	"github.com/aquasecurity/tfsec/internal/app/tfsec/scanner"
	"github.com/aquasecurity/tfsec/pkg/rule"
)

func init() {
	scanner.RegisterCheckRule(rule.Rule{
		BadExample: []string{`
 resource "azurerm_mssql_server_security_alert_policy" "bad_example" {
   resource_group_name        = azurerm_resource_group.example.name
   server_name                = azurerm_sql_server.example.name
   state                      = "Enabled"
   storage_endpoint           = azurerm_storage_account.example.primary_blob_endpoint
   storage_account_access_key = azurerm_storage_account.example.primary_access_key
   disabled_alerts = [
   ]
   email_account_admins = false
 }
 `},
		GoodExample: []string{`
 resource "azurerm_mssql_server_security_alert_policy" "good_example" {
   resource_group_name        = azurerm_resource_group.example.name
   server_name                = azurerm_sql_server.example.name
   state                      = "Enabled"
   storage_endpoint           = azurerm_storage_account.example.primary_blob_endpoint
   storage_account_access_key = azurerm_storage_account.example.primary_access_key
   disabled_alerts = []
 
   email_account_admins = true
 }
 `},
		Links: []string{
			"https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/mssql_server_security_alert_policy#email_account_admins",
		},
		RequiredTypes: []string{
			"resource",
		},
		RequiredLabels: []string{
			"azurerm_mssql_server_security_alert_policy",
		},
		Base: database.CheckThreatAlertEmailToOwner,
		CheckTerraform: func(resourceBlock block.Block, _ block.Module) (results rules.Results) {
			if emailAccountAdminsAttr := resourceBlock.GetAttribute("email_account_admins"); emailAccountAdminsAttr.IsNil() { // alert on use of default value
				results.Add("Resource uses default value for email_account_admins", resourceBlock)
			} else if emailAccountAdminsAttr.IsFalse() {
				results.Add("Resource has attribute email_account_admins that is false", emailAccountAdminsAttr)
			}
			return results
		},
	})
}
