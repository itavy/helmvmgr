package utils

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"helmvmgr/pkg"
	"io"
	"text/template"
)

type RenderOutputParameters struct {
	Out io.Writer
	Data interface{}
	Type string
	Template string
}

type RenderFunc func(rp RenderOutputParameters) error

func RenderOutput(rp RenderOutputParameters) error {
	if rp.Type == "" {
		tmpl, err := template.New("rendernow").Parse(rp.Template)
		if err != nil {
			return errors.HError("RO_01", err.Error())
		}
		err = tmpl.Execute(rp.Out, rp.Data)
		if err != nil {
			return errors.HError("RO_02", err.Error())
		}
		return nil
	}

	if rp.Type == "json" {
		err := json.NewEncoder(rp.Out).Encode(rp.Data)
		if err != nil {
			return errors.HError("RO_03", err.Error())
		}
		return nil
	}

	if rp.Type == "yaml" {
		err := yaml.NewEncoder(rp.Out).Encode(rp.Data)
		if err != nil {
			return errors.HError("RO_03", err.Error())
		}
		return nil
	}


	return errors.HError("RO_04", fmt.Sprintf("Invalid type to render: %s.", rp.Type))
}
