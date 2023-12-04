package ultis

import (
	"encoding/json"
	"io"
	"os"
)

func ReadJSONFile(inputPath string) ([]byte, error) {
	jsonFile, err := os.Open(inputPath)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	offerInBytes, _ := io.ReadAll(jsonFile)

	return offerInBytes, nil
}

func WriteJSONFile(outputPath string, object any, fileMode os.FileMode) {
	file, _ := json.MarshalIndent(object, "", " ")
	_ = os.WriteFile(outputPath, file, fileMode)
}
