package gui

import (
	"cli/http"
	"cli/models"
	"fmt"
	"github.com/go-resty/resty/v2"
	"reflect"

	"github.com/manifoldco/promptui"
)

const tagName = "promptType"

func GeneratePromptFromStruct(s interface{}) error {
	v := reflect.ValueOf(s)

	if v.Kind() != reflect.Struct {
		return fmt.Errorf("%s is not a struct", v.Kind())
	}

	for i := 0; i < v.NumField(); i++ {
		// Get the field tag value
		tag := v.Type().Field(i).Tag.Get(tagName)

		// Skip if tag is not defined or ignored
		if tag == "" || tag == "-" {
			continue
		}

		if tag == "prompt" {
			prompt := promptui.Prompt{Label: fmt.Sprintf("Enter %s", v.Type().Field(i).Name)}
			res, err := prompt.Run()
			if err != nil {
				return err
			}
		}

		if tag == "select" {


			prompt := promptui.Select{Label: fmt.Sprintf("Select %s", v.Type().Field(i).Name), Items: }
			_,res, err := prompt.Run()
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func getItems() ([]string,error) {
	c.SendMonitoringHttpRequest()
}


func getClient(source string) {
	switch source {
	case "metrics":
		c :=  http.NewRestClient[[]*models.Service{}](resty.New())
		c.SendMonitoringHttpRequest()
	default:
		return nil
	}
}
