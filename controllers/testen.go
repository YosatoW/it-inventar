package controllers

import (
	"fmt"
	"it_inventar/models"
	"it_inventar/views/console"
	"strings"
)

// handleEditItem bearbeitet einen Artikel im Inventar
func handleEditItem() {
	console.Clear()
	items := models.GetAllItems() // Hole alle Artikel aus dem Inventar

	if len(items) == 0 {
		console.ShowMessage("❌ Es sind keine Artikel im Inventar vorhanden.")
		console.ShowContinue()
		console.Clear()
		console.ShowExecuteCommandMenu()
		return
	}

	page := 0
	pageSize := 20
	for {
		var isEditing bool = false
		var newName, newModel, newNotes string
		var newQuantity int

		// Berechne den Indexbereich für die aktuelle Seite
		start := page * pageSize
		end := start + pageSize
		if end > len(items) {
			end = len(items)
		}

		// Zeige die Artikel der aktuellen Seite
		console.ShowAllItems(items[start:end], start)

		// Zeige die Eingabeaufforderung zum Bearbeiten, Blättern oder Abbrechen
		console.ShowMessage("Gib die ID des zu bearbeitenden Artikels ein, drücke [Enter] für die nächste Seite oder [c], um zum Hauptmenü zurückzukehren.")
		choice := console.AskForInput()

		if choice == "c" {
			// Zurück ins Hauptmenü
			console.Clear()
			console.ShowExecuteCommandMenu()
			return
		} else if choice == "" {
			// Weiter zur nächsten Seite
			page++
			if end == len(items) {
				// Wenn es keine weiteren Artikel mehr gibt
				console.ShowMessage("Alle Artikel wurden angezeigt.")
				console.ShowContinue()
				return
			}
		} else {
			// Prüfe, ob die Eingabe eine gültige ID ist
			rowId := models.StringToInt(choice)
			if rowId <= 0 || rowId > len(items) {
				console.ShowMessage("❌ Ungültige ID. Bitte gib eine gültige ID ein.")
				console.ShowContinue()
				continue
			}

			// Überprüfen, ob der Artikel existiert und anzeigen
			item := models.GetItemById(rowId - 1) // Hier wird der Index korrekt angepasst
			if item == nil {
				console.ShowMessage("❌ Artikel mit dieser ID existiert nicht.")
				console.ShowContinue()
				continue
			}

			// Zeige aktuelle Artikelinformationen und frage nach neuen Werten
			console.ShowMessage(fmt.Sprintf("Aktuelle Informationen für Artikel: %s (%s) - %d Stück - Notizen: %s", item.Name, item.Model, item.Quantity, item.Note))
			console.ShowMessage("Geben Sie die neuen Werte ein (drücken Sie [Enter], um den aktuellen Wert beizubehalten):")

			newName = console.AskForName(item.Name, isEditing)
			newModel = console.AskForModel(item.Model, isEditing)
			newQuantity = console.AskForQuantity(item.Quantity, isEditing)
			newNotes = console.AskForNotes(item.Note, isEditing)

			// Bestätigung zum Bearbeiten des Artikels

			// Bestätigung zum Bearbeiten des Artikels
			console.ShowMessage("Bitte überprüfen Sie die neuen Daten:")
			console.ShowMessage(fmt.Sprintf("Artikelbezeichnung: %s", newName))
			console.ShowMessage(fmt.Sprintf("Artikelnummer: %s", newModel))
			console.ShowMessage(fmt.Sprintf("Menge: %d", newQuantity))
			console.ShowMessage(fmt.Sprintf("Notizen: %s", newNotes))
			console.ShowMessage("\nSind die Daten korrekt? (y/n) oder [c] um zum Hauptmenü zurückzukehren.")

			choice := console.AskForInput()
			if strings.ToLower(choice) == "y" {
				// Artikel aktualisieren
				data := models.Item{
					Name:     newName,
					Model:    newModel,
					Quantity: newQuantity,
					Note:     newNotes,
				}
				// Hier wird der Index korrekt angepasst
				err := models.UpdateItem(rowId-1, data)
				if err != nil {
					console.ShowError(err)
				} else {
					console.ShowMessage("✅ Artikel erfolgreich aktualisiert!")
				}

				return
			} else if strings.ToLower(choice) == "c" {
				//Abbrechen und zurück zum Menü
				console.ShowMessage("Vorgang abgebrochen")
				return
			} else {
				console.ShowMessage("✏️ Bitte korrigieren Sie die Daten.")
				isEditing = true // Korrekturmodus aktivieren
				//	return           // Beende die Funktion und kehre zum Hauptmenü zurück
				//} else {
				//	console.ShowMessage("❌ Änderungen wurden nicht gespeichert.")
			}
		}

	}

}
