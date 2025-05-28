package gms

import (
	"embed"
	"os"

	"github.com/vrianta/GMS-GO-CMS/src/arguments"
)

//go:embed templates/*
var Templates embed.FS

// Object of the GMS which hold all the configaration for the application
type GMS struct {
	project_name string
	project_path string
	Controllers  []string // names of the controllers needs to be created
	Models       []string // names of the models needs to be created
	instructions struct {
		NewApplication bool
		ProjectPath    bool
		NewController  bool
		NewModel       bool
		MigrateDb      bool
	} // instructions to be executed
}

// Function to start the Running the application which will handle all the command like arguments and proceses
func (g *GMS) Start() {
	g.SetupApplicationDetails()
	g.ExecuteInstructions()
}

// A Function to setupo the application needed details from the arguments
func (g *GMS) SetupApplicationDetails() {
	arguments_lenth := len(os.Args)
	if arguments_lenth < 2 {
		g.Error("No instruction provided. Please provide a command line arguments. use -h for help.", true)
	}
	for i := 1; i < arguments_lenth; i++ {
		switch os.Args[i] {
		case arguments.HELP, arguments.Help: // Show Help
			g.Help()
			break
		case arguments.NEW, arguments.New: // Create New Project
			if arguments_lenth < 3 || g.CheckIfTheNextArgumentIsNotInstruction(os.Args[i+1]) {
				g.Error("No project name provided. Please provide a project name.", true)
			}
			g.instructions.NewApplication = true
			g.project_name = os.Args[i+1]
			i++
			break
		case arguments.PROJECTPATH, arguments.ProjectPath: // Project Path
			if arguments_lenth < 5 || g.CheckIfTheNextArgumentIsNotInstruction(os.Args[i+1]) {
				g.Error("No project path provided or No project Name has been Provided. Please provide a project path.", true)
			}
			g.project_path = os.Args[i+1]
			i++
			break
		case arguments.NewController, arguments.NEWCONTROLLER:
			if arguments_lenth < 3 || g.CheckIfTheNextArgumentIsNotInstruction(os.Args[i+1]) {
				g.Error("No controller name provided. Please provide a controller name.", true)
			}
			g.instructions.NewController = true
			for j := i + 1; j < arguments_lenth && g.CheckIfTheNextArgumentIsNotInstruction(os.Args[j]); j++ {
				g.Controllers = append(g.Controllers, os.Args[j])
				i++
			}
			break
		case arguments.NewModel, arguments.NEWMODEL:
			if arguments_lenth < 3 || g.CheckIfTheNextArgumentIsNotInstruction(os.Args[i+1]) {
				g.Error("No model name provided. Please provide a model name.", true)
			}
			g.instructions.NewModel = true
			for j := i + 1; j < arguments_lenth && g.CheckIfTheNextArgumentIsNotInstruction(os.Args[j]); j++ {
				g.Models = append(g.Models, os.Args[j])
				i++
			}
			break
		default:
			g.Error("Invalid argument: "+os.Args[i]+". Please use -h for help.", true)
			break
		}
	}
}

// It will check if the next atgument is not any instruction from the pre made applications so that we know in case if the user do not the
// Application name after command
func (g *GMS) CheckIfTheNextArgumentIsNotInstruction(_argument string) bool {

	switch _argument {
	case arguments.HELP:
	case arguments.Help:
	case arguments.Version:
	case arguments.VERSION:
	case arguments.New:
	case arguments.NEW:
	case arguments.NewController:
	case arguments.NEWCONTROLLER:
	case arguments.NewModel:
	case arguments.NEWMODEL:
	case arguments.Migrate:
	case arguments.MIGRATE:
	case arguments.ProjectPath:
	case arguments.PROJECTPATH:
		return true
	default:
		return false
	}

	return false
}

func (g *GMS) ExecuteInstructions() {
	if g.instructions.NewApplication {
		g.NewProject()
	}

	if g.instructions.NewController {
		// g.CreateControllers()
	}

	if g.instructions.NewModel {
		// g.CreateModels()
	}

	if g.instructions.MigrateDb {
		// g.MigrateDatabase()
	}
}
