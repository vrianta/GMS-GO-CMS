// This file is part of the Go Mod Setup package.
// used to create project and it's base template and files
package gms

import (
	"fmt"
	"os"
	"text/template"
	"time"
)

var FolderStructure = []string{
	"Views",
	"Controllers",
	"Controllers/Admin", // Default Admin Controller
	"Controllers/Api",   // Default Api Controller
	"Controllers/Home",  // Default Home Controller
	"Controllers/Login",
	"Controllers/Logout",
	"Controllers/Register",
	"Controllers/Settings",
	"Controllers/Users",    // Default Users Controller
	"Controllers/Posts",    // Default Posts Controller
	"Controllers/Comments", // Default Comments Controller
	"Models",
	"Models/Users",    // Default Users Model
	"Models/Posts",    // Default Posts Model
	"Models/Comments", // Default Comments Model
	"Components",
	"Js",
	"Css",
	"Static",
}

var DefaultFiles = []string{
	"main.go",   // Main file of the project
	"config.go", // Configuration file of the project
	"routes.go", // Router file of the project
	"Readme.md", // Readme file of the project
	"app.json",  // Application configuration file
}

func (g *GMS) NewProject() {
	// Create Folder Structure
	// first I will createa Creating the Go Mod file where replacing the @appname with the app name the user provided
	g.CreateFolderStructure()
	g.CreateBasicFiles()

	os.Chdir(g.project_path) // Change the working directory to the project path
	g.goModInit()            // Initialize the Go module for the project
	g.goModTidy()            // Tidy the Go module to remove unused dependencies
}

func (g *GMS) CreateFolderStructure() {
	if g.project_path == "" {
		g.project_path = "./" + g.project_name
		if _, err := os.Stat(g.project_path); os.IsNotExist(err) {
			if _err := os.Mkdir(g.project_path, 0755); _err != nil {
				g.Error("Error creating project folder: "+g.project_name+" Error : "+_err.Error(), true)
			} else {
				fmt.Println("Project folder created successfully at " + g.project_path)
			}
		} else {
			g.Error("A Folder Already exists with the name "+g.project_name+" in the current directory. Please choose a different name or delete the existing folder.", true)
		}
	}

	for _, folder_paths := range FolderStructure {
		_final_folder_path := g.project_path + "/" + folder_paths
		if _, err := os.Stat(_final_folder_path); os.IsNotExist(err) {
			if _err := os.Mkdir(_final_folder_path, 0755); _err != nil {
				g.Error("Error creating project folder: "+_final_folder_path+" Error : "+_err.Error(), true)
			} else {
				fmt.Println("Project folder created successfully at " + _final_folder_path)
			}
		} else {
			g.Error("A Folder Already exists with the name "+_final_folder_path+" in the current directory. Please choose a different name or delete the existing folder.", true)
		}
	}
}

// Function to Create Basic Files of the Project
// Example: main.go, config.go, etc.
// This function will create the basic files of the project like main.go, config.go, etc.
// Taking the Files from the Templates folder and copying them to the project folder
// This function will also replace the @appname with the project name in the files
func (g *GMS) CreateBasicFiles() {
	for _, file_names := range DefaultFiles {

		_final_file_path := g.project_path + "/" + file_names // final path of the file where the file will be created
		fmt.Println("Creating file: " + _final_file_path)
		// Check if the file already exists
		// If the file does not exist, then create the file
		if _, err := os.Stat(_final_file_path); os.IsNotExist(err) {

			if tpl, _err := Templates.ReadFile("templates/" + file_names + ".template"); _err != nil {
				g.Error("Error reading template file: "+file_names+" Error : "+_err.Error(), true)
			} else {
				//fmt.Println("Template file read successfully: " + string(tpl))
				// Create a template and parse the string into it
				if t, __err := template.New("").Parse(string(tpl)); __err != nil {
					g.Error("Error Persing Default Files", true)
				} else {
					// execute the template and write the output in a file
					if file, ___err := os.Create(_final_file_path); ___err != nil {
						g.Error("Error creating file: "+_final_file_path+" Error : "+___err.Error(), true)
					} else {
						defer file.Close()

						t.Execute(file, map[string]string{
							"appname": g.project_name,
							"version": "v0.0.1",
							"author":  "Your Name",
							"email":   "",
							"website": "",
							"license": "MIT",
							"year":    fmt.Sprint(time.Now().Year()),
						})
					}
				}
			}

		} else {
			g.Error("A File Already exists with the name "+_final_file_path+" in the current directory. Please choose a different name or delete the existing file.", true)
		}
	}
}
