package functions

import (
	"context"
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

type JsonReplacer struct {
	function.Function
}

func NewJsonReplacer() function.Function {
	return &JsonReplacer{}
}

func (r *JsonReplacer) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "json_replacer"
}

func (r *JsonReplacer) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "JSON values replacer",
		MarkdownDescription: "Function provides JSON variables override by passed parameters",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:                "json",
				MarkdownDescription: "String: Serialized JSON (ex. file content)",
			},
			function.StringParameter{
				Name:				 "parameters",
				MarkdownDescription: "String: JSON values to override the original JSON",
			},
			function.BoolParameter{
				Name:				 "onlyExisting",
				MarkdownDescription: "Boolean: Override only existing fields - new fields would be ignored.",
			},
		},
		Return: function.StringReturn{},
	}
}

func (r JsonReplacer) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	
	// Pobranie inputów
	var inputJson string
	var inputParameters string
	var onlyExisting bool

	req.Arguments.Get(ctx, &inputJson, &inputParameters, &onlyExisting)

	// Przygotowanie mapy i deserializacja
	var jsonData map[string]interface{};
	var jsonDataParameters map[string]interface{};

	json.Unmarshal([]byte(inputJson), &jsonData)
	json.Unmarshal([]byte(inputParameters), &jsonDataParameters)

	// Podmiana wartości

	if (onlyExisting) {
		for key, value := range jsonDataParameters {
			var _, existingVal = jsonData[key]
			// Pomijamy, gdy klucz nie istnieje
			if !existingVal {
				continue;
			}
			jsonData[key] = value
		}
	} else {
		for key, value := range jsonDataParameters {
			jsonData[key] = value
		}
	}

	// Serializacja danych
	outputJson, _ := json.Marshal(jsonData)
	resp.Result.Set(ctx, string(outputJson))
}