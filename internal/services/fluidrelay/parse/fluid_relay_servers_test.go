package parse

// NOTE: this file is generated via 'go:generate' - manual changes will be overwritten

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.Id = FluidRelayServersId{}

func TestFluidRelayServersIDFormatter(t *testing.T) {
	actual := NewFluidRelayServersID("00000000-0000-0000-0000-000000000000", "rg1", "server1").ID()
	expected := "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.FluidRelay/fluidRelayServers/server1"
	if actual != expected {
		t.Fatalf("Expected %q but got %q", expected, actual)
	}
}

func TestFluidRelayServersID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *FluidRelayServersId
	}{

		{
			// empty
			Input: "",
			Error: true,
		},

		{
			// missing SubscriptionId
			Input: "/",
			Error: true,
		},

		{
			// missing value for SubscriptionId
			Input: "/subscriptions/",
			Error: true,
		},

		{
			// missing ResourceGroup
			Input: "/subscriptions/00000000-0000-0000-0000-000000000000/",
			Error: true,
		},

		{
			// missing value for ResourceGroup
			Input: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/",
			Error: true,
		},

		{
			// missing FluidRelayServerName
			Input: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.FluidRelay/",
			Error: true,
		},

		{
			// missing value for FluidRelayServerName
			Input: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.FluidRelay/fluidRelayServers/",
			Error: true,
		},

		{
			// valid
			Input: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.FluidRelay/fluidRelayServers/server1",
			Expected: &FluidRelayServersId{
				SubscriptionId:       "00000000-0000-0000-0000-000000000000",
				ResourceGroup:        "rg1",
				FluidRelayServerName: "server1",
			},
		},

		{
			// upper-cased
			Input: "/SUBSCRIPTIONS/00000000-0000-0000-0000-000000000000/RESOURCEGROUPS/RG1/PROVIDERS/MICROSOFT.FLUIDRELAY/FLUIDRELAYSERVERS/SERVER1",
			Error: true,
		},
	}

	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := FluidRelayServersID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %s", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SubscriptionId != v.Expected.SubscriptionId {
			t.Fatalf("Expected %q but got %q for SubscriptionId", v.Expected.SubscriptionId, actual.SubscriptionId)
		}
		if actual.ResourceGroup != v.Expected.ResourceGroup {
			t.Fatalf("Expected %q but got %q for ResourceGroup", v.Expected.ResourceGroup, actual.ResourceGroup)
		}
		if actual.FluidRelayServerName != v.Expected.FluidRelayServerName {
			t.Fatalf("Expected %q but got %q for FluidRelayServerName", v.Expected.FluidRelayServerName, actual.FluidRelayServerName)
		}
	}
}
