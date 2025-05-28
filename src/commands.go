// This file is holding all the command line instruction that might be needed or used in the application configaration or modifiocation
// Example : go mod init, go mod tidy, etc.

package gms

import (
	"fmt"
	"os/exec"
)

// function to execute go mod init command
// This function will create a go.mod file in the project folder
func (g *GMS) goModInit() {
	fmt.Println("Initializing Go module for the project...")
	cmd := exec.Command("go", "mod", "init", g.project_name)
	if err := cmd.Run(); err != nil {
		g.Error("Error executing go mod init command: "+err.Error(), true)
	} else {
		fmt.Println("Go mod file created successfully at " + g.project_path + "/go.mod")
	}
}

func (g *GMS) goModTidy() {
	fmt.Println("Running go mod tidy to clean up dependencies...")
	cmd := exec.Command("go", "mod", "tidy")
	if err := cmd.Run(); err != nil {
		g.Error("Error executing go mod tidy command: "+err.Error(), true)
	} else {
		fmt.Println("Go mod tidy executed successfully at " + g.project_path)
	}
}
