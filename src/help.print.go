package gms

import (
	"fmt"
	"os"
)

// Function to print Help Massage which print all the avaiable commands and their description
func (g *GMS) Help() {
	fmt.Println("GMS Help:")
	fmt.Println("--------------------------------------------------")
	fmt.Println("Available commands:")
	fmt.Println("--------------------------------------------------")
	fmt.Println("# -h / --help: Show this help message")
	fmt.Println("# -v / --version: Show the version of GMS")
	fmt.Println("# -n / --new: Create a new project | Usage: gms -n <project_name>")
	fmt.Println("# -pt / --project-path: Specify the project path | Usage: gms -n <project_name> -pt <path>")
	fmt.Println("# -nc / --new-controller: Create a new controller")
	fmt.Println("# -nm / --new-controller: Create a new model")
	fmt.Println("# -m / --migrate: Migrate the database")
	fmt.Println("--------------------------------------------------")
	fmt.Println("--------------------------------------------------")
	fmt.Println("For more information, visit")

	os.Exit(0)
}
