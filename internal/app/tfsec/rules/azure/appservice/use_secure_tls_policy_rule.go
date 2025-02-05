package appservice

import (
	"github.com/aquasecurity/defsec/rules"
	"github.com/aquasecurity/defsec/rules/azure/appservice"
	"github.com/aquasecurity/tfsec/internal/app/tfsec/block"
	"github.com/aquasecurity/tfsec/internal/app/tfsec/scanner"
	"github.com/aquasecurity/tfsec/pkg/rule"
)

func init() {
	scanner.RegisterCheckRule(rule.Rule{
		BadExample: []string{`
 resource "azurerm_app_service" "bad_example" {
   name                = "example-app-service"
   location            = azurerm_resource_group.example.location
   resource_group_name = azurerm_resource_group.example.name
   app_service_plan_id = azurerm_app_service_plan.example.id
 
   site_config {
 	  min_tls_version = "1.0"
   }
 }
 `},
		GoodExample: []string{`
 resource "azurerm_app_service" "good_example" {
   name                = "example-app-service"
   location            = azurerm_resource_group.example.location
   resource_group_name = azurerm_resource_group.example.name
   app_service_plan_id = azurerm_app_service_plan.example.id
 }
 `},
		Links: []string{
			"https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/app_service#min_tls_version",
		},
		RequiredTypes: []string{
			"resource",
		},
		RequiredLabels: []string{
			"azurerm_app_service",
		},
		Base: appservice.CheckUseSecureTlsPolicy,
		CheckTerraform: func(resourceBlock block.Block, _ block.Module) (results rules.Results) {
			if resourceBlock.MissingChild("site_config") {
				return
			}

			if minTlsVersionAttr := resourceBlock.GetBlock("site_config").GetAttribute("min_tls_version"); minTlsVersionAttr.IsNotNil() && minTlsVersionAttr.NotEqual("1.2") {
				results.Add("Resource does not have site_config.min_tls_version set to 1.2", minTlsVersionAttr)
			}
			return results
		},
	})
}
