package authorization

import (
	"github.com/aquasecurity/defsec/rules"
	"github.com/aquasecurity/defsec/rules/azure/authorization"
	"github.com/aquasecurity/tfsec/internal/app/tfsec/block"
	"github.com/aquasecurity/tfsec/internal/app/tfsec/scanner"
	"github.com/aquasecurity/tfsec/pkg/rule"
)

func init() {
	scanner.RegisterCheckRule(rule.Rule{
		BadExample: []string{`
 data "azurerm_subscription" "primary" {
 }
 
 resource "azurerm_role_definition" "example" {
   name        = "my-custom-role"
   scope       = data.azurerm_subscription.primary.id
   description = "This is a custom role created via Terraform"
 
   permissions {
     actions     = ["*"]
     not_actions = []
   }
 
   assignable_scopes = [
     "/"
   ]
 }
 `},
		GoodExample: []string{`
 data "azurerm_subscription" "primary" {
 }
 
 resource "azurerm_role_definition" "example" {
   name        = "my-custom-role"
   scope       = data.azurerm_subscription.primary.id
   description = "This is a custom role created via Terraform"
 
   permissions {
     actions     = ["*"]
     not_actions = []
   }
 
   assignable_scopes = [
     data.azurerm_subscription.primary.id,
   ]
 }
 `},
		Links: []string{
			"https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/role_definition#actions",
		},
		RequiredTypes: []string{
			"resource",
		},
		RequiredLabels: []string{
			"azurerm_role_definition",
		},
		Base: authorization.CheckLimitRoleActions,
		CheckTerraform: func(resourceBlock block.Block, _ block.Module) (results rules.Results) {
			if permissionBlock := resourceBlock.GetBlock("permissions"); permissionBlock.IsNotNil() {
				if actionsAttr := permissionBlock.GetAttribute("actions"); actionsAttr.IsNotNil() && actionsAttr.Contains("*") {
					// need more information
					if assignableScopesAttr := resourceBlock.GetAttribute("assignable_scopes"); assignableScopesAttr.IsNil() {
						results.Add("Resource has wildcard action with open assignable_scopes", resourceBlock)
					} else if assignableScopesAttr.Contains("/") {
						results.Add("Resource has wildcard action with open assignable_scopes", assignableScopesAttr)
					}
				}
			}
			return results
		},
	})
}
