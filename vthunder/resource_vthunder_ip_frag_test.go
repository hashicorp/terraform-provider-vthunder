package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_IP_FRAG_RESOURCE = `
resource "vthunder_ip_frag" "frag" {
	buff = 10000
	max_packets_per_reassembly = 2
	max_reassembly_sessions = 1
	timeout = 4
	sampling_enable  {
		counters1 = "all"
	}
	cpu_threshold  {
		high = 75
		low = 60
	} 
}
`

//Acceptance test
func TestAccIpFrag_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_IP_FRAG_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_ip_frag.frag", "buff", "10000"),
					resource.TestCheckResourceAttr("vthunder_ip_frag.frag", "max_packets_per_reassembly", "2"),
					resource.TestCheckResourceAttr("vthunder_ip_frag.frag", "max_reassembly_sessions", "1"),
					resource.TestCheckResourceAttr("vthunder_ip_frag.frag", "timeout", "4"),
					resource.TestCheckResourceAttr("vthunder_ip_frag.frag", "sampling_enable.0.counters1", "all"),
					resource.TestCheckResourceAttr("vthunder_ip_frag.frag", "cpu_threshold.0.high", "75"),
					resource.TestCheckResourceAttr("vthunder_ip_frag.frag", "cpu_threshold.0.low", "60"),
				),
			},
		},
	})
}
