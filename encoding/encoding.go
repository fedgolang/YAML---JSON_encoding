package encoding

import (
	"encoding/json"
	"io"
	"os"

	"github.com/Yandex-Practicum/final-project-encoding-go/models"
	"gopkg.in/yaml.v3"
)

// JSONData тип для перекодирования из JSON в YAML
type JSONData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// YAMLData тип для перекодирования из YAML в JSON
type YAMLData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// MyEncoder интерфейс для структур YAMLData и JSONData
type MyEncoder interface {
	Encoding() error
}

// Encoding перекодирует файл из JSON в YAML
func (j *JSONData) Encoding() error {
	// ниже реализуйте метод
	var obj interface{}
	JSONdata, err := os.Open("jsonInput.json")
	if err != nil {
		return err
	}
	data, err := io.ReadAll(JSONdata)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &obj)
	if err != nil {
		return err
	}
	out, err := yaml.Marshal(obj)
	if err != nil {
		return err
	}
	yamlFile, err := os.Create("yamlOutput.yml")
	if err != nil {
		return err
	}
	defer yamlFile.Close()
	_, err = yamlFile.Write(out)
	if err != nil {
		return err
	}
	return nil
}

// Encoding перекодирует файл из YAML в JSON
func (y *YAMLData) Encoding() error {
	// Ниже реализуйте метод
	var obj interface{}
	YAMLdata, err := os.Open("yamlInput.yml")
	if err != nil {
		return err
	}
	data, err := io.ReadAll(YAMLdata)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(data, &obj)
	if err != nil {
		return err
	}
	out, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	JSONFile, err := os.Create("jsonOutput.json")
	if err != nil {
		return err
	}
	defer JSONFile.Close()
	_, err = JSONFile.Write(out)
	if err != nil {
		return err
	}
	return nil
}
