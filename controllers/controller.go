package controllers

import (
	"fmt"
	"it_inventar/models"
	"it_inventar/views/console"
	"strings"
)

// Run does the running of the console application
func Run() {
	err := models.Initialize()
	checkAndHandleError(err)

	console.Clear()
	console.ShowMenu()

	for {
		executeCommand()
	}
}

func checkAndHandleError(err error) {
	if err != nil {
		console.ShowError(err)
		return
	}
}

// executeCommand asks the user for an input command and executes a corresponding function depending on the command
func executeCommand() {
	command := console.AskForInput()
	switch command {
	case "1":
		handleAddItem()
	case "2":
		handleRemoveItem()
	case "3":
		handleEditItem()
	case "4":
		handleChanceArticleInformation()
	case "5":
		handleRemoveItem()
	case "6":
	case "7":
	case "8":
	case "9":
		handleViewItems()
	case "c":
		console.Clear()
		console.ShowMenu()
	case "q":
		console.Clear()
		console.ShowGoodbye()
		console.ShutDownNormal()
	default:
		console.ShowMessage("Unbekannter Befehl. Bitte versuchen Sie es erneut.")
	}
}

// handleAddItem  Adds a new item to the inventory.
func handleAddItem() {
	console.Clear()
	console.ShowAddItemInformation()

	var isEditing bool = false
	// Speicher der eingegebenen Werte für den Korrekturmodus
	var name, itemType, notes string
	var quantity int

	for {
		// Die Eingabewerte werden jetzt nur einmal initialisiert und bei Korrekturen wiederverwendet
		name = console.AskForName(name, isEditing)
		itemType = console.AskForType(itemType, isEditing)
		quantity = console.AskForQuantity(quantity, isEditing)
		notes = console.AskForNotes(notes, isEditing)

		// Benutzer überprüft die Eingaben
		console.Clear()
		console.ShowMessage("Bitte überprüfen Sie die eingegebenen Daten:")
		console.ShowMessage(fmt.Sprintf("Artikelbezeichnung: %s", name))
		console.ShowMessage(fmt.Sprintf("Artikelnummer: %s", itemType))
		console.ShowMessage(fmt.Sprintf("Menge: %d", quantity))
		console.ShowMessage(fmt.Sprintf("Notizen: %s", notes))
		console.ShowMessage("\nSind die Daten korrekt? (y/n) oder [c]  um zum Hauptmenü zurückzukehren.")

		choice := console.AskForInput()
		if strings.ToLower(choice) == "y" {
			// Artikel zusammenstellen
			data := models.Item{
				Name:     name,
				Model:    itemType,
				Quantity: quantity,
				Note:     notes,
			}
			// Artikel hinzufügen
			err := models.AddItem(data)
			if err != nil {
				console.ShowError(err)
			} else {
				console.ShowMessage("✅ Artikel erfolgreich hinzugefügt!")
			}
			break
		} else if strings.ToLower(choice) == "c" {
			//Abbrechen und zurück zum Menü
			console.ShowMessage("Vorgang abgebrochen")
			break
		} else {
			console.ShowMessage("✏️ Bitte korrigieren Sie die Daten.")
			isEditing = true // Korrekturmodus aktivieren
		}
	}

	console.ShowContinue()
	console.Clear()
	console.ShowMenu()
}

// handleRemoveItem Handles the removal of an item from the inventory.
func handleRemoveItem() {
	console.Clear()
	items := models.GetAllItems() // Hole alle Artikel aus dem Inventar

	if len(items) == 0 {
		console.ShowMessage("❌ Es sind keine Artikel im Inventar vorhanden.")
		console.ShowContinue()
		console.Clear()
		console.ShowMenu()
		return
	}

	page := 0
	pageSize := 20
	for {
		// Berechne den Indexbereich für die aktuelle Seite
		start := page * pageSize
		end := start + pageSize
		if end > len(items) {
			end = len(items)
		}

		// Zeige die Artikel der aktuellen Seite
		console.ShowAllItems(items[start:end], start)

		// Zeige die Eingabeaufforderung zum Weiterblättern oder Beenden
		console.ShowMessage("Drücke [Enter], um mehr zu sehen oder [c], um zum Hauptmenü zurückzukehren.")
		choice := console.AskForInput()

		if choice == "c" {
			// Zurück ins Hauptmenü
			console.Clear()
			console.ShowMenu()
			break
		} else if choice == "" {
			// Weiter zur nächsten Seite
			page++
			if end == len(items) {
				// Wenn es keine weiteren Artikel mehr gibt
				console.ShowMessage("Alle Artikel wurden angezeigt.")
				console.ShowContinue()
				console.Clear()
				break
			}
		}
	}

	// Abfrage zur Eingabe der ID des zu löschenden Artikels
	console.ShowMessage("Gib die ID des zu löschenden Artikels ein:")
	itemID := console.AskForInput()
	rowId := models.StringToInt(itemID)

	// Überprüfen, ob der Artikel existiert und anzeigen
	item := models.GetItemById(rowId - 1) // Hier wird der Index korrekt angepasst
	if item == nil {
		console.ShowMessage("❌ Artikel mit dieser ID existiert nicht.")
		console.ShowContinue()
		console.Clear()
		console.ShowMenu()
		return
	}

	// Bestätigung zum Löschen des Artikels
	console.ShowMessage(fmt.Sprintf("Artikel: %s (%s) - %d Stück - Notizen: %s", item.Name, item.Model, item.Quantity, item.Note))
	console.ShowMessage("Möchten Sie diesen Artikel wirklich löschen? (y/n)")

	confirm := console.AskForInput()
	if confirm == "y" {
		// Artikel löschen
		err := models.RemoveItem(rowId)
		if err != nil {
			console.ShowError(err)
		} else {
			console.ShowMessage("✅ Artikel erfolgreich entfernt!")
		}
	} else {
		console.ShowMessage("❌ Artikel wurde nicht gelöscht.")
	}

	console.ShowContinue()
	console.Clear()
	console.ShowMenu()
}

func handleChanceArticleInformation() {}
func handleEditItem()                 {}

//

// handleViewItems Displays all items in the inventory.
func handleViewItems() {
	console.Clear()
	items := models.GetAllItems() // Entferne "err" und den zweiten Rückgabewert

	if len(items) == 0 {
		console.ShowMessage("❌ Es sind keine Artikel im Inventar vorhanden.")
	} else {
		page := 0
		pageSize := 20
		for {
			// Berechne den Indexbereich für die aktuelle Seite
			start := page * pageSize
			end := start + pageSize
			if end > len(items) {
				end = len(items)
			}

			// Zeige die Artikel der aktuellen Seite
			console.ShowAllItems(items[start:end], start)

			// Zeige die Eingabeaufforderung zum Weiterblättern oder Beenden
			console.ShowMessage("Drücke [Enter], um mehr zu sehen oder [c], um zum Hauptmenü zurückzukehren.")
			choice := console.AskForInput()

			if choice == "c" {
				// Zurück ins Hauptmenü
				console.Clear()
				console.ShowMenu()
				break
			} else if choice == "" {
				// Weiter zur nächsten Seite
				page++
				if end == len(items) {
					// Wenn es keine weiteren Artikel mehr gibt
					console.ShowMessage("Alle Artikel wurden angezeigt.")
					console.ShowContinue()
					console.Clear()
					console.ShowMenu()
					break
				}
			}
		}
	}
}

func parseAndExecuteCommand(command string) {}
