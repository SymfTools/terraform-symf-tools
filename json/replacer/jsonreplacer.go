package replacer

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceJsonModifier() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceJsonModifierRead,
		Schema: map[string]*schema.Schema{
			"json": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeMap,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
			"result": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceJsonModifierRead(d *schema.ResourceData, m interface{}) error {
	return modifyJson(d)
}

func modifyJson(d *schema.ResourceData) error {
	inputJson := d.Get("json").(string)
	valuesToReplace := d.Get("parameters").(map[string]interface{})

	var jsonData map[string]interface{}
	json.Unmarshal([]byte(inputJson), &jsonData)

	for key, value := range valuesToReplace {
		jsonData[key] = value
	}

	outputJson, _ := json.Marshal(jsonData)
	d.Set("result", string(outputJson))

	return nil;
}