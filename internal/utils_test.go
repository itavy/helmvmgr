package utils

import (
	"bytes"
	"encoding/json"
	"gopkg.in/yaml.v2"
	"testing"
	"text/template"
)

func TestRenderOutput(t *testing.T) {
	testInputData := struct {
		Version string
		BoolField bool `yaml: "Boolean"`
		NumberField int `yaml: "Number"`
		StrField string `yaml: "String"`
	} {
		Version: "0.0.0",
		BoolField: false,
		NumberField: 123456789,
		StrField: "randomstring",

	}
	outTemplate := `Version:     {{ .Version }}
BoolField:   {{ .BoolField }}
NumberField: {{ .NumberField }}
StrField:    {{ .StrField }}`
	tmpl, _ := template.New("rendernow").Parse(outTemplate)
	var expectedRender bytes.Buffer
	tmpl.Execute(&expectedRender, testInputData)
	expectedJson, _ := json.Marshal(testInputData)
	expectedYAML, _ := yaml.Marshal(testInputData)

	tcs := [] struct{
		name string
		outputType string
		template string
		expected string
	}{
		{
			name: "Default output",
			outputType: "",
			template: outTemplate,
			expected: expectedRender.String(),
		},
		{
			name: "JSON output",
			outputType: "json",
			expected: string(expectedJson) + "\n",
		},
		{
			name: "YAML output",
			outputType: "yaml",
			expected: string(expectedYAML),
		},
	}

	for _, testData := range tcs {
		t.Run(testData.name, func(tt *testing.T) {
			var testWritter bytes.Buffer
			ttData := RenderOutputParameters{
				Data: testInputData,
				Out: &testWritter,
				Type: testData.outputType,
				Template: testData.template,
			}
			err := RenderOutput(ttData)
			if err != nil {
				t.Errorf("unexpected error, got '%v'", err)
			}

			if testWritter.String() != testData.expected {
				tt.Errorf("expected\n%q\ngot\n%q\n", testData.expected, testWritter.String())
			}
		})
	}
}