package compute

import (
	"encoding/base64"

	"github.com/aquasecurity/defsec/rules"
	"github.com/aquasecurity/defsec/rules/cloudstack/compute"
	"github.com/aquasecurity/tfsec/internal/app/tfsec/block"
	"github.com/aquasecurity/tfsec/internal/app/tfsec/debug"

	"github.com/aquasecurity/tfsec/internal/app/tfsec/scanner"
	"github.com/aquasecurity/tfsec/pkg/rule"
	"github.com/owenrumney/squealer/pkg/squealer"
)

func init() {
	scanner.RegisterCheckRule(rule.Rule{
		BadExample: []string{`
 resource "cloudstack_instance" "web" {
   name             = "server-1"
   service_offering = "small"
   network_id       = "6eb22f91-7454-4107-89f4-36afcdf33021"
   template         = "CentOS 6.5"
   zone             = "zone-1"
   user_data        = <<EOF
 export DATABASE_PASSWORD=\"SomeSortOfPassword\"
 EOF
 }
 `, `
 resource "cloudstack_instance" "web" {
   name             = "server-1"
   service_offering = "small"
   network_id       = "6eb22f91-7454-4107-89f4-36afcdf33021"
   template         = "CentOS 6.5"
   zone             = "zone-1"
   user_data        = "ZXhwb3J0IERBVEFCQVNFX1BBU1NXT1JEPSJTb21lU29ydE9mUGFzc3dvcmQi"
 }
 `},
		GoodExample: []string{`
 resource "cloudstack_instance" "web" {
   name             = "server-1"
   service_offering = "small"
   network_id       = "6eb22f91-7454-4107-89f4-36afcdf33021"
   template         = "CentOS 6.5"
   zone             = "zone-1"
   user_data        = <<EOF
 export GREETING="Hello there"
 EOF
 }
 `, `
 resource "cloudstack_instance" "web" {
   name             = "server-1"
   service_offering = "small"
   network_id       = "6eb22f91-7454-4107-89f4-36afcdf33021"
   template         = "CentOS 6.5"
   zone             = "zone-1"
   user_data        = "ZXhwb3J0IEVESVRPUj12aW1hY3M="
 }
 `},
		Links: []string{
			"https://registry.terraform.io/providers/hashicorp/cloudstack/latest/docs/resources/instance#",
		},
		RequiredTypes: []string{
			"resource",
		},
		RequiredLabels: []string{
			"cloudstack_instance",
		},
		Base: compute.CheckNoSensitiveInfo,
		CheckTerraform: func(resourceBlock block.Block, _ block.Module) (results rules.Results) {

			customDataAttr := resourceBlock.GetAttribute("user_data")

			if customDataAttr.IsNotNil() && customDataAttr.IsString() {
				encoded, err := base64.StdEncoding.DecodeString(customDataAttr.Value().AsString())
				if err != nil {
					debug.Log("could not decode the base64 string in the terraform, trying with the string verbatim")
					encoded = []byte(customDataAttr.Value().AsString())
				}
				if checkStringForSensitive(string(encoded)) {
					results.Add("Resource has user_data_base64 with sensitive data.", customDataAttr)
				}
			}

			return results
		},
	})
}

func checkStringForSensitive(stringToCheck string) bool {
	return squealer.NewStringScanner().Scan(stringToCheck).TransgressionFound
}
