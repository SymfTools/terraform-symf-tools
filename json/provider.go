package json

import (
	"github.com/Boltairex/terraform-symf-tools/json/replacer"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
			"json_replacer": replacer.DataSourceJsonModifier(),
		},
	}
}