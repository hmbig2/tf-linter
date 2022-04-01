package test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDliQueue_basic(t *testing.T) {
	rName := "123"
	resourceName := "queue.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() {},
		ProviderFactories: nil,
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				Config: testAccDliQueue_basic(rName, 16),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "queue_type", "1"),
					resource.TestCheckResourceAttr(resourceName, "cu_count", "16"),
					resource.TestCheckResourceAttrSet(resourceName, "resource_mode"),
					resource.TestCheckResourceAttrSet(resourceName, "create_time"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"tags",
				},
			},
		},
	})
}

func testAccDliQueue_basic(rName string, cuCount int) string {
	return fmt.Sprintf(`
resource "queue" "test" {
  name     = "%s"
  cu_count = %d

  tags = {
    foo = "bar"
  }
}
`, rName, cuCount)
}
