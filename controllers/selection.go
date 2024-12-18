package controllers

import (
	"it_inventar/views/console"
	"strings"
)

// executeCommand asks the user for an input command and executes a corresponding function depending on the command
func executeCommand() {
	command := console.AskForInput()
	switch command {
	case "1":
		handleAddItem()
	case "2":
		handleRemoveItem()
	case "3":
		handleChangeQuantity()
	case "4":
		handleChanceArticleInformation()
	case "5":
		// TODO: Funktionalität implementieren
	case "6":
		// TODO: Funktionalität implementieren
	case "7":
		// TODO: Funktionalität implementieren
	case "8":
		// TODO: Funktionalität implementieren
	case "9":
		handleViewItems()
	case "4600":
		console.Clear()
		hiddenCommand()
	case "c":
		console.Clear()
		console.ShowExecuteCommandMenu()
	case "q":
		console.Clear()
		console.ShowGoodbye()
		console.ShutDownNormal()
	default:
		console.Clear()
		console.ShowExecuteCommandMenu()
		console.ShowMessage("❌ Ungültige Auswahl. Bitte versuchen Sie es erneut.")
	}
}

// hiddenCommand menu for administrator access with code 4600
func hiddenCommand() {
	console.ShowMessage("🔒 Geben Sie den Zugangscode ein:")
	code := console.AskForInput()

	if code != "4600" {
		console.ShowMessage("❌ Falscher Zugangscode.")
		console.Clear()
		console.ShowExecuteCommandMenu()
		return
	}
	// Menü-Schleife: bleibt im Hidden Command Menü
	for {
		console.ShowHiddenCommandMenu()

		choice := console.AskForInput()
		if strings.TrimSpace(choice) == "" { // Leere Eingaben ignorieren
			console.ShowMessage("⚠️ Leere Eingabe. Bitte eine gültige Option wählen.")
			continue
		}
		switch choice {
		case "1":
			console.AskForSupplier("", false)
		case "2":
			console.AskForStatus("", false)
		case "3":
			console.AskForCategory("", false)
		case "4":
			// TODO: Funktionalität implementieren
		case "5":
			// TODO: Funktionalität implementieren
		case "6":
			// TODO: Funktionalität implementieren
		case "c":
			console.ShowMessage("🔙 Verlassen des Hidden Command Menüs...")
			console.ShowExecuteCommandMenu()
			return // Verwende return, um die Schleife und Funktion zu verlassen
		default:
			console.ShowMessage("❌ Ungültige Auswahl. Bitte versuchen Sie es erneut.")
		}
	}
}
