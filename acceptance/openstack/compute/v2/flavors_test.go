// +build acceptance

package v2

import (
	"testing"

	"github.com/rackspace/gophercloud/openstack/compute/v2/flavors"
	"github.com/rackspace/gophercloud/pagination"
)

func TestListFlavors(t *testing.T) {
	client, err := newClient()
	if err != nil {
		t.Fatalf("Unable to create a compute client: %v", err)
	}

	t.Logf("ID\tRegion\tName\tStatus\tCreated")

	pager := flavors.List(client, flavors.ListFilterOptions{})
	count, pages := 0, 0
	pager.EachPage(func(page pagination.Page) (bool, error) {
		t.Logf("---")
		pages++
		flavors, err := flavors.ExtractFlavors(page)
		if err != nil {
			return false, err
		}

		for _, f := range flavors {
			t.Logf("%s\t%s\t%d\t%d\t%d", f.ID, f.Name, f.RAM, f.Disk, f.VCPUs)
		}

		return true, nil
	})

	t.Logf("--------\n%d flavors listed on %d pages.", count, pages)
}

func TestGetFlavor(t *testing.T) {
	client, err := newClient()
	if err != nil {
		t.Fatalf("Unable to create a compute client: %v", err)
	}

	choices, err := ComputeChoicesFromEnv()
	if err != nil {
		t.Fatal(err)
	}

	result, err := flavors.Get(client, choices.FlavorID)
	if err != nil {
		t.Fatalf("Unable to get flavor information: %v", err)
	}
	flavor, err := flavors.ExtractFlavor(result)
	if err != nil {
		t.Fatalf("Unable to extract flavor from GET result: %v", err)
	}

	t.Logf("Flavor: %#v", flavor)
}
