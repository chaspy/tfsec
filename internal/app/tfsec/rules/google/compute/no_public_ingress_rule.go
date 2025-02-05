package compute

import (
	"github.com/aquasecurity/defsec/rules"
	"github.com/aquasecurity/defsec/rules/google/compute"
	"github.com/aquasecurity/tfsec/internal/app/tfsec/block"
	"github.com/aquasecurity/tfsec/internal/app/tfsec/cidr"
	"github.com/aquasecurity/tfsec/internal/app/tfsec/scanner"
	"github.com/aquasecurity/tfsec/pkg/rule"
)

func init() {
	scanner.RegisterCheckRule(rule.Rule{
		LegacyID: "GCP003",
		BadExample: []string{`
 resource "google_compute_firewall" "bad_example" {
 	source_ranges = ["0.0.0.0/0"]
 }`},
		GoodExample: []string{`
 resource "google_compute_firewall" "good_example" {
 	source_ranges = ["1.2.3.4/32"]
 }`},
		Links: []string{
			"https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/compute_firewall#source_ranges",

			"https://www.terraform.io/docs/providers/google/r/compute_firewall.html",
		},
		RequiredTypes:  []string{"resource"},
		RequiredLabels: []string{"google_compute_firewall"},
		Base:           compute.CheckNoPublicIngress,
		CheckTerraform: func(resourceBlock block.Block, _ block.Module) (results rules.Results) {
			if sourceRanges := resourceBlock.GetAttribute("source_ranges"); sourceRanges.IsNotNil() {
				if cidr.IsAttributeOpen(sourceRanges) {
					results.Add("Resource defines a fully open inbound firewall rule.", sourceRanges)
				}
			}
			return results
		},
	})
}
