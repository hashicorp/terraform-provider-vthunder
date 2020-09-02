package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_FW_TCP_WINDOW_CHECK_RESOURCE = `
resource "vthunder_fw_tcp_window_check" "FwTest" {
        status = "enable" 
}
`

//Acceptance test
func TestAccFwTcpWindowCheck_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_FW_TCP_WINDOW_CHECK_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_fw_tcp_window_check.FwTest", "status", "enable"),
				),
			},
		},
	})
}
