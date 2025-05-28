package gms

import (
	"embed"
	"fmt"
	"os"

	"github.com/vrianta/GMS-GO-CMS/src/argument"
)

//go:embed templates/*
var templates embed.FS

// Object of the GMS which hold all the configaration for the application
type GMS struct {
}

// Function to start the Running the application which will handle all the command like arguments and proceses
func (g *GMS) Start() {

	arguments_lenth := len(os.Args)
	if arguments_lenth < 2 {
		g.Error("No instruction provided. Please provide a command line argument. use -h for help.", true)
	}
	// Bellow Variable storing the command line arguments
	instruction := os.Args[1] // first argument\

	switch instruction {
	case argument.HELP:
		g.Help()
		break
	case argument.NEW: // Create New Project
		if arguments_lenth < 3 {
			g.Error("No project name provided. Please provide a project name.", true)
		}
		g.NewProject()
		break
	default:
		g.Help()
	}
}

// Function to print Help Massage which print all the avaiable commands and their description
func (g *GMS) Help() {
	fmt.Println("GMS Help:")
	fmt.Println("Available commands:")
	fmt.Println("-h: Show this help message")
	fmt.Println("-v: Show the version of GMS")
	fmt.Println("-n: Create a new project")
	fmt.Println("-nc: Create a new controller")
	fmt.Println("-nm: Create a new model")
	fmt.Println("-m: Migrate the database")
	fmt.Println("For more information, visit")
}
