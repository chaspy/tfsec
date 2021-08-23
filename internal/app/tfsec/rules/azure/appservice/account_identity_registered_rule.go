package appservice
// 
// // ATTENTION!
// // This rule was autogenerated!
// // Before making changes, consider updating the generator.
// 
// import (
// 	"github.com/aquasecurity/defsec/provider"
// 	"github.com/aquasecurity/defsec/result"
// 	"github.com/aquasecurity/defsec/severity"
// 	"github.com/aquasecurity/tfsec/internal/app/tfsec/block"
// 	"github.com/aquasecurity/tfsec/internal/app/tfsec/scanner"
// 	"github.com/aquasecurity/tfsec/pkg/rule"
// )
// 
// func init() {
// 	scanner.RegisterCheckRule(rule.Rule{
// 		Provider:  provider.AzureProvider,
// 		Service:   "appservice",
// 		ShortCode: "account-identity-registered",
// 		Documentation: rule.RuleDocumentation{
// 			Summary:     "Web App has registration with AD enabled",
// 			Explanation: `Registering the identity used by an App with AD allows it to interact with other services without using username and password`,
// 			Impact:      "Interaction between services can't easily be achieved without username/password",
// 			Resolution:  "Register the app identity with AD",
// 			BadExample: []string{`
// resource "azurerm_app_service" "bad_example" {
//   name                = "example-app-service"
//   location            = azurerm_resource_group.example.location
//   resource_group_name = azurerm_resource_group.example.name
//   app_service_plan_id = azurerm_app_service_plan.example.id
// }
// `},
// 			GoodExample: []string{`
// resource "azurerm_app_service" "good_example" {
//   name                = "example-app-service"
//   location            = azurerm_resource_group.example.location
//   resource_group_name = azurerm_resource_group.example.name
//   app_service_plan_id = azurerm_app_service_plan.example.id
// 
//   identity {
//     type = "UserAssigned"
//     identity_ids = "webapp1"
//   }
// }
// `},
// 			Links: []string{
// 				"https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/app_service#identity",
// 			},
// 		},
// 		RequiredTypes: []string{
// 			"resource",
// 		},
// 		RequiredLabels: []string{
// 			"azurerm_app_service",
// 		},
// 		DefaultSeverity: severity.Low,
// 		CheckTerraform: func(set result.Set, resourceBlock block.Block, module block.Module) {
// 			if identityAttr := resourceBlock.GetBlock("identity"); identityAttr.IsNil() { // alert on use of default value
// 				set.AddResult().
// 					WithDescription("Resource '%s' does not set identity", resourceBlock.FullName())
// 			}
// 		},
// 	})
// }
