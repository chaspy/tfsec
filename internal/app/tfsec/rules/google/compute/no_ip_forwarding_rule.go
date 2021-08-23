package compute
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
// 		Provider:  provider.GoogleProvider,
// 		Service:   "compute",
// 		ShortCode: "no-ip-forwarding",
// 		Documentation: rule.RuleDocumentation{
// 			Summary:     "Instances should not have IP forwarding enabled",
// 			Explanation: `Disabling IP forwarding ensuresthe instance can only receive packets addressed to the instance and can only send packets with a source address of the instance.`,
// 			Impact:      "Instance can send/receive packets without the explicit instance address",
// 			Resolution:  "Disable IP forwarding",
// 			BadExample: []string{`
// resource "google_compute_instance" "bad_example" {
//   name         = "test"
//   machine_type = "e2-medium"
//   zone         = "us-central1-a"
// 
//   boot_disk {
//     initialize_params {
//       image = "debian-cloud/debian-9"
//     }
//   }
// 
//   // Local SSD disk
//   scratch_disk {
//     interface = "SCSI"
//   }
// 
//   can_ip_forward = true
// }
// `},
// 			GoodExample: []string{`
// resource "google_compute_instance" "bad_example" {
//   name         = "test"
//   machine_type = "e2-medium"
//   zone         = "us-central1-a"
// 
//   boot_disk {
//     initialize_params {
//       image = "debian-cloud/debian-9"
//     }
//   }
// 
//   // Local SSD disk
//   scratch_disk {
//     interface = "SCSI"
//   }
//   
//   can_ip_forward = false
// }
// `},
// 			Links: []string{
// 				"https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/compute_instance#can_ip_forward",
// 			},
// 		},
// 		RequiredTypes: []string{
// 			"resource",
// 		},
// 		RequiredLabels: []string{
// 			"google_compute_instance",
// 		},
// 		DefaultSeverity: severity.High,
// 		CheckTerraform: func(set result.Set, resourceBlock block.Block, _ block.Module) {
// 			if canIpForwardAttr := resourceBlock.GetAttribute("can_ip_forward"); canIpForwardAttr.IsTrue() {
// 				set.AddResult().
// 					WithDescription("Resource '%s' has can_ip_forward set to true", resourceBlock.FullName()).
// 					WithAttribute("")
// 			}
// 		},
// 	})
// }
