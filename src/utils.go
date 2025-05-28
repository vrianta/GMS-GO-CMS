// This file will contain some commond needed function and featues which is essential example read the app.json file to get information about the application
// 				// Create a file and write the template to it

package gms

import (
	"encoding/json"
	"os"
)

type AppJson struct {
	Application string `json:"Application"` // Name of the application
	Version     string `json:"Version"`     // Version of the application
}

/**
** ReadAppJson reads the app.json file and sets the project_name field in the GMS struct.
** It unmarshals the JSON data into an AppJson struct and retrieves the application name.
** If the file cannot be read or unmarshalled, it logs an error and exits the program.
** It is used to initialize the project_name field in the GMS struct based on the contents of app.json.
 */
func (g *GMS) ReadAppJson() {
	// Read the app.json file and return the data as a map
	data, err := os.ReadFile("app.json")
	if err != nil {
		g.Error("Error reading app.json file: "+err.Error(), true)
	}
	var appData AppJson
	if err := json.Unmarshal(data, &appData); err != nil {
		g.Error("Error unmarshalling app.json file: "+err.Error(), true)
	}

	g.project_name = appData.Application
}
