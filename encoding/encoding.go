package encoding

import (
	"encoding/json"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/hive-bootcamp/final-project-encoding-go/models"
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
	data, err := os.ReadFile(j.FileInput)
	if err != nil {
		fmt.Println("ошибка чтения yaml файла:", err.Error())
		return err
	}

	err = json.Unmarshal(data, &j.DockerCompose)
	if err != nil {
		fmt.Println("ошибка десериализации json файла:", err.Error())
		return err
	}

	data, err = yaml.Marshal(&j.DockerCompose)
	if err != nil {
		fmt.Println("ошибка сериализации в yaml файл:", err.Error())
		return err
	}
	err = os.WriteFile(j.FileOutput, data, 0644)
	if err != nil {
		fmt.Println("ошибка записи yaml файла:", err.Error())
		return err
	}

	return nil
}

// Encoding перекодирует файл из YAML в JSON
func (y *YAMLData) Encoding() error {
	// Ниже реализуйте метод
	data, err := os.ReadFile(y.FileInput)
	if err != nil {
		fmt.Println("ошибка чтения yaml файла:", err.Error())
		return err
	}

	err = yaml.Unmarshal(data, &y.DockerCompose)
	if err != nil {
		fmt.Println("ошибка десериализации yaml файла:", err.Error())
		return err
	}

	data, err = json.MarshalIndent(&y.DockerCompose, "", " ")
	if err != nil {
		fmt.Println("ошибка сериализации в json:", err.Error())
		return err
	}

	err = os.WriteFile(y.FileOutput, data, 0644)
	if err != nil {
		fmt.Println("ошибка записи json файла:", err.Error())
		return err
	}

	return nil
}
