package bigquery

import (
	"github.com/aquasecurity/defsec/rules"
	"github.com/aquasecurity/defsec/rules/google/bigquery"
	"github.com/aquasecurity/tfsec/internal/app/tfsec/block"
	"github.com/aquasecurity/tfsec/internal/app/tfsec/scanner"
	"github.com/aquasecurity/tfsec/pkg/rule"
)

func init() {
	scanner.RegisterCheckRule(rule.Rule{
		BadExample: []string{`
 resource "google_bigquery_dataset" "bad_example" {
   dataset_id                  = "example_dataset"
   friendly_name               = "test"
   description                 = "This is a test description"
   location                    = "EU"
   default_table_expiration_ms = 3600000
 
   labels = {
     env = "default"
   }
 
   access {
     role          = "OWNER"
     special_group = "allAuthenticatedUsers"
   }
 
   access {
     role   = "READER"
     domain = "hashicorp.com"
   }
 }
 
 `},
		GoodExample: []string{`
 resource "google_bigquery_dataset" "good_example" {
   dataset_id                  = "example_dataset"
   friendly_name               = "test"
   description                 = "This is a test description"
   location                    = "EU"
   default_table_expiration_ms = 3600000
 
   labels = {
     env = "default"
   }
 
   access {
     role          = "OWNER"
     user_by_email = google_service_account.bqowner.email
   }
 
   access {
     role   = "READER"
     domain = "hashicorp.com"
   }
 }
 
 resource "google_service_account" "bqowner" {
   account_id = "bqowner"
 }
 `},
		Links: []string{
			"https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/bigquery_dataset#special_group",
		},
		RequiredTypes: []string{
			"resource",
		},
		RequiredLabels: []string{
			"google_bigquery_dataset",
		},
		Base: bigquery.CheckNoPublicAccess,
		CheckTerraform: func(resourceBlock block.Block, _ block.Module) (results rules.Results) {
			if specialGroupAttr := resourceBlock.GetBlock("access").GetAttribute("special_group"); specialGroupAttr.Equals("allAuthenticatedUsers") {
				results.Add("Resource has access.special_group set to allAuthenticatedUsers", resourceBlock)
			}
			return results
		},
	})
}
