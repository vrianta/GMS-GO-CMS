// This file is part of the Go Mod Setup package.
// used to create project and it's base template and files
package gms

import (
	"fmt"
	"os"
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
	"Contrllers/Users",     // Default Users Controller
	"Controllers/Posts",    // Default Posts Controller
	"Controllers/Comments", // Default Comments Controller
	"Models",
	"Models/Users",    // Default Users Model
	"Models/Posts",    // Default Posts Model
	"Models/Comments", // Default Comments Model
	"Components",
}

func (g *GMS) NewProject() {
	// Create Folder Structure
	// first I will createa Creating the Go Mod file where replacing the @appname with the app name the user provided
	g.CreateFolderStructure()
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
