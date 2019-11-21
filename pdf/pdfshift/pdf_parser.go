package pdfshift

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	pdfshiftURL = "https://api.pdfshift.io/v2/convert"
)

type PDF struct {
	Html []byte
}

// DrawHeader draw a text in header format with a specific size, this size comes
// from the standar definition of the six different headers in markdown
func (p PDF) DrawHeader(text string, size float64) {
	// Not implemented
}

// DrawText draw a text and move to the next line, text variable is the value of
// the text and the style can be "B" (bold), "I" (italic), "U" (underscore), "S" (strike-out)
// or any combination. The default value (specified with an empty string) is
// regular. Bold and italic styles do not apply to Symbol and ZapfDingbats.
func (p PDF) DrawText(text string, style string) {
	// Not Implemented
}

// OutputFile just call the specific call from the library for output and create the pdf file
// this method also calls the pdfshift API to get the pdf converted
func (p PDF) OutputFile(fileName string) error {
	message := map[string]interface{}{
		"source": string(p.Html),
	}

	bytesRepresentation, err := json.Marshal(message)
	if err != nil {
		return err
	}

	client := http.Client{}
	request, err := http.NewRequest("POST", pdfshiftURL, bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(request)
	if err != nil {
		return err
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		ioutil.WriteFile(fileName, body, 0644)
	} else {
		var result map[string]interface{}

		json.NewDecoder(resp.Body).Decode(&result)

		log.Println(result)
		log.Println(result["data"])
	}
	return nil
}
