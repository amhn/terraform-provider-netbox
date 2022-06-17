package netbox

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	netboxclient "github.com/smutel/go-netbox/netbox/client"
	"github.com/smutel/go-netbox/netbox/client/circuits"
)

func dataNetboxJSONCircuitsCircuitTypesList() *schema.Resource {
	return &schema.Resource{
		Description: "Get json output from the circuits_circuit_types_list Netbox endpoint.",
		Read:        dataNetboxJSONCircuitsCircuitTypesListRead,

		Schema: map[string]*schema.Schema{
			"limit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     0,
				Description: "The max number of returned results. If 0 is specified, all records will be returned.",
			},
			"json": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "JSON output of the list of objects for this Netbox endpoint.",
			},
		},
	}
}

func dataNetboxJSONCircuitsCircuitTypesListRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*netboxclient.NetBoxAPI)

	params := circuits.NewCircuitsCircuitTypesListParams()
	limit := int64(d.Get("limit").(int))
	params.Limit = &limit

	list, err := client.Circuits.CircuitsCircuitTypesList(params, nil)
	if err != nil {
		return err
	}

	j, _ := json.Marshal(list.Payload.Results)

	d.Set("json", string(j))
	d.SetId("NetboxJSONCircuitsCircuitTypesList")

	return nil
}
