package replacer

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceJsonModifier() *schema.Resource {
	return &schema.Resource{
		Create: resourceJsonModifierCreate,
		Read:   resourceJsonModifierRead,
		Update: resourceJsonModifierUpdate,
		Delete: resourceJsonModifierDelete,
		Schema: map[string]*schema.Schema{
			"input_json": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"values_to_replace": &schema.Schema{
				Type:     schema.TypeMap,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
			"output_json": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceJsonModifierCreate(d *schema.ResourceData, m interface{}) error {
	return modifyJson(d)
}

func resourceJsonModifierRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceJsonModifierUpdate(d *schema.ResourceData, m interface{}) error {
	return modifyJson(d)
}

func resourceJsonModifierDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}

func modifyJson(d *schema.ResourceData) error {
	inputJson := d.Get("input_json").(string)
	valuesToReplace := d.Get("values_to_replace").(map[string]interface{})

	var jsonData map[string]interface{}
	json.Unmarshal([]byte(inputJson), &jsonData)

	for key, value := range valuesToReplace {
		jsonData[key] = value
	}

	outputJson, _ := json.Marshal(jsonData)
	d.Set("output_json", string(outputJson))

	return nil
}