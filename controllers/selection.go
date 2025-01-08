package controllers

import (
	"it_inventar/views/console"
	"strings"
)

// executeCommand asks the user for an input command and executes a corresponding function depending on the command
func executeCommand() {
	command := strings.ToUpper(console.AskForInput())
	switch command {
	case "1":
		handleAddItem()
	case "2":
		handleRemoveItem()
	case "3":
		handleChangeQuantity()
	case "4":
		handleChanceArticleInformation()
	case "9":
		handleViewItems()
	case "4600":
		console.Clear()
		hiddenCommand()
	case "C":
		console.Clear()
		console.ShowExecuteCommandMenu()
	case "Q":
		console.Clear()
		console.ShowGoodbye()
		console.ShutDownNormal()
	default:
		console.Clear()
		console.ShowExecuteCommandMenu()
		console.ShowMessage("❌ Invalid selection. Please try again.")
	}
}

// hiddenCommand menu for administrator access with code 4600
func hiddenCommand() {
	console.ShowMessage("🔒 Enter the access code:")
	code := console.AskForInput()

	if code != "4600" {
		console.ShowMessage("❌ Incorrect access code.")
		console.Clear()
		console.ShowExecuteCommandMenu()
		return
	}
	// Menu-loop: stays in hidden command menu
	for {
		console.ShowHiddenCommandMenu()

		choice := strings.ToUpper(console.AskForInput())
		if strings.TrimSpace(choice) == "" { // Leere Eingaben ignorieren
			console.ShowMessage("⚠️ Empty input. Please choose a valid option.")
			continue
		}
		switch choice {
		case "1":
			handleShowSuppliers()
		case "2":
			handleAddSuppliers()
		case "3":
			handleDeleteSupplier()
		case "11":
			handleShowCategories()
		case "12":
			handleAddCategories()
		case "13":
			handleDeleteCategories()
		case "ID":
			handleViewDeletedItems()
		case "IA":
			handleViewAllItems()
		case "C":
			console.Clear()
			console.ShowMessage("🔙 Exiting the Hidden Command Menu...")
			console.ShowExecuteCommandMenu()
			return // Use return to exit the loop and function
		default:
			console.Clear()
			console.ShowMessage("❌ Invalid selection. Please try again.")
		}
	}
}
