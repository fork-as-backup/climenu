package main

import (
	"fmt"
	"os"

	"github.com/paulrademacher/climenu"
)

func callback(id string) {
	fmt.Println("Chose item:", id)

}

func main() {
	step1 := climenu.NewButtonMenu("Welcome", "Where do you want to go today?")
	step1.AddMenuItem("SparkJob", "sparkjob")
	step1.AddMenuItem("PythonJob", "pythonjob")

	crd, escaped := step1.Run()
	if escaped {
		os.Exit(0)
	}
	fmt.Println("Selected Resource >", crd)

	step2 := climenu.NewButtonMenu("Welcome", "Action to be performed?")
	step2.AddMenuItem("Create", "create")
	step2.AddMenuItem("Delete", "delete")
	step2.AddMenuItem("List All", "list")
	step2.AddMenuItem("Inspect", "inspect")

	action, escaped := step2.Run()
	if escaped {
		os.Exit(0)
	}
	fmt.Println("action >", action)

	checkbox := climenu.NewCheckboxMenu("Which job would you like to inspect?",
		"Select options", "OK", "Cancel")
	checkbox.AddMenuItem("py-abcd", "py-abcd")
	checkbox.AddMenuItem("py-cdef", "py-cdef")
	checkbox.AddMenuItem("py-ghij", "py-ghij")

	selection, escaped := checkbox.Run()
	if escaped {
		os.Exit(0)
	}
	fmt.Println("selected >", selection)

	response := climenu.GetText("Provide a jar uri", "")
	if escaped {
		os.Exit(0)
	}

	fmt.Printf("text > \"%s\"\n", response)
}
