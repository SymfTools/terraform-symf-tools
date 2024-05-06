package replacer

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type JsonReplacer struct {
	datasource.DataSource
}

func modifyJsonFunc(d *schema.ResourceData) error {
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