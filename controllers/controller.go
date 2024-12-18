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
	console.ShowExecuteCommandMenu()

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

// Case 01
// handleAddItem Adds a new item to the inventory.
func handleAddItem() {
	console.Clear()
	console.ShowAddItemInformation()

	var isEditing bool = false
	// Speicher der eingegebenen Werte für den Korrekturmodus
	var name, itemModel, notes string
	var quantity int

	for {
		// Die Eingabewerte werden jetzt nur einmal initialisiert und bei Korrekturen wiederverwendet
		name = console.AskForName(name, isEditing)
		itemModel = console.AskForModel(itemModel, isEditing)
		quantity = console.AskForQuantity(quantity, isEditing)
		notes = console.AskForNotes(notes, isEditing)

		// Benutzer überprüft die Eingaben
		console.Clear()
		console.ShowMessage("Bitte überprüfen Sie die eingegebenen Daten:")
		console.ShowMessage(fmt.Sprintf("Artikelbezeichnung: %s", name))
		console.ShowMessage(fmt.Sprintf("Artikelnummer: %s", itemModel))
		console.ShowMessage(fmt.Sprintf("Menge: %d", quantity))
		console.ShowMessage(fmt.Sprintf("Notizen: %s", notes))
		console.ShowMessage("\nSind die Daten korrekt? (y/n) oder [c]  um zum Hauptmenü zurückzukehren.")

		choice := console.AskForInput()
		if strings.ToLower(choice) == "y" {
			// Artikel zusammenstellen
			data := models.Item{
				Name:     name,
				Model:    itemModel,
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
			console.InputC()
			return
		} else {
			console.ShowMessage("✏️ Bitte korrigieren Sie die Daten.")
			isEditing = true // Korrekturmodus aktivieren
		}
	}

	console.ShowContinue()
	console.Clear()
	console.ShowExecuteCommandMenu()
}

// Case 02
// handleRemoveItem Handles the removal of an item from the inventory.
func handleRemoveItem() {
	console.Clear()
	items := models.GetAllItems() // Hole alle Artikel aus dem Inventar

	// Überprüft Inventar auf Inhalt
	if console.ChecksInventory() {
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

		// Zeige die Eingabeaufforderung zum Löschen, Blättern oder Abbrechen
		console.ShowMessage("Gib die ID des zu löschenden Artikels ein, drücke [Enter] für die nächste Seite oder [c], um zum Hauptmenü zurückzukehren.")
		choice := console.AskForInput()

		if strings.ToLower(choice) == "c" {
			// Zurück ins Hauptmenü
			console.Clear()
			console.ShowExecuteCommandMenu()
			return
		} else if strings.ToLower(choice) == "" {
			// Weiter zur nächsten Seite
			page++
			if end == len(items) {
				console.InputPageEnd()
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

			// Bestätigung zum Löschen des Artikels
			console.ShowMessage(fmt.Sprintf("Artikel: %s (%s) - %d Stück - Notizen: %s", item.Name, item.Model, item.Quantity, item.Note))
			console.ShowMessage("Möchten Sie diesen Artikel wirklich löschen? (y/n) oder [c], um zum Hauptmenü zurückzukehren.")

			choice = console.AskForInput()

			if strings.ToLower(choice) == "y" {
				// Artikel löschen
				err := models.RemoveItem(rowId)
				if err != nil {
					console.ShowError(err)
				} else {
					console.ShowMessage("✅ Artikel erfolgreich entfernt!")
					console.ShowContinue()
					console.Clear()
					console.ShowExecuteCommandMenu()
					return
				}
			} else if strings.ToLower(choice) == "n" {
				// Artikel nicht löschen, Abbruch
				console.ShowMessage("❌ Artikel wurde nicht gelöscht.")
				console.ShowContinue()
				console.Clear()
				break
			} else if strings.ToLower(choice) == "c" {
				console.InputC()
				return
			} else {
				// Ungültige Eingabe, erneut fragen
				console.ShowMessage(fmt.Sprintf("Ungültige Eingabe. Bitte wählen:\n[y] zum Löschen\n[n] zum Behalten\n[c] um zum Hauptmenü zurückzukehren.\n"))
			}
		}
	}
}

// Case 03
// handleChangeQuantity bearbeitet einen Artikel im Inventar
// handleChangeQuantity bearbeitet einen Artikel im Inventar
func handleChangeQuantity() {
	console.Clear()
	items := models.GetAllItems() // Hole alle Artikel aus dem Inventar

	if console.ChecksInventory() {
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

		// Zeige die Eingabeaufforderung zum Bearbeiten, Blättern oder Abbrechen
		console.ShowMessage("Gib die ID des zu bearbeitenden Artikels ein, drücke [Enter] für die nächste Seite oder [c], um zum Hauptmenü zurückzukehren.")
		choice := console.AskForInput()

		if strings.ToLower(choice) == "c" {
			console.InputC()
			return
		} else if strings.ToLower(choice) == "" {
			page++
			if end == len(items) {
				console.InputPageEnd()
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
			console.ShowMessage("Möchten Sie die Menge dieses Artikels ändern? (y/n) oder [c], um zum Hauptmenü zurückzukehren")

			choice = console.AskForInput()
			if strings.ToLower(choice) == "y" {
				// Frage nach Einbuchen oder Abbuchen
				console.ShowMessage(fmt.Sprintf("[1] Einbuchen\n[2] Ausbuchen"))
				operation := console.AskForInput()

				if strings.ToLower(operation) == "1" {
					// Frage nach der Menge zum Einbuchen
					console.ShowMessage(fmt.Sprintf("Aktuelle Bestand: %d Stück", item.Quantity))
					console.ShowMessage("Geben Sie die Menge ein, die eingebucht werden soll:")
					quantityToAdd := console.AskForQuantity(0, false)
					item.Quantity += quantityToAdd
				} else if strings.ToLower(operation) == "2" {
					// Frage nach der Menge zum Abbuchen
					console.ShowMessage(fmt.Sprintf("Aktuelle Bestand: %d Stück", item.Quantity))
					console.ShowMessage("Geben Sie die Menge ein, die abgebucht werden soll:")
					quantityToSubtract := console.AskForQuantity(0, false)
					if item.Quantity < quantityToSubtract {
						console.ShowMessage("❌ Die abzubuchende Menge überschreitet die vorhandene Menge.")
						console.ShowContinue()
						continue
					}
					item.Quantity -= quantityToSubtract
				} else {
					console.ShowMessage("❌ Ungültige Auswahl. Bitte wählen Sie '1' oder '2'.")
					console.ShowContinue()
					continue
				}

				// Artikelmenge aktualisieren
				err := models.UpdateItem(rowId-1, *item)
				if err != nil {
					console.ShowError(err)
				} else {
					console.ShowMessage(fmt.Sprintf("Neue Bestand: %d Stück", item.Quantity))
					console.ShowMessage("✅ Artikelmenge erfolgreich aktualisiert!")
					console.ShowContinue()
					console.Clear()
					console.ShowExecuteCommandMenu()
					return
				}
			} else if strings.ToLower(choice) == "c" {
				console.InputC()
				return
			} else if strings.ToLower(choice) == "n" {
				// Artikel nicht ändern, Abbruch
				console.ShowMessage("❌ Artikelmenge wurde nicht geändert.")
				console.ShowContinue()
				console.Clear()
				break
			} else {
				// Ungültige Eingabe, erneut fragen
				console.ShowMessage(fmt.Sprintf("Ungültige Eingabe. Bitte wählen:\n[y] zum Ändern\n[n] zum Behalten\n[c] um zum Hauptmenü zurückzukehren.\n"))
			}
		}
	}
}

// Case 03
// handleChanceArticleInformation bearbeitet einen Artikel im Inventar
func handleChanceArticleInformation() {
	console.Clear()
	items := models.GetAllItems() // Hole alle Artikel aus dem Inventar

	if console.ChecksInventory() {
		return
	}

	page := 0
	pageSize := 20
	for {
		var isEditing bool = false
		var newName, newModel, newNotes string

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

		if strings.ToLower(choice) == "c" {
			console.InputC()
			return
		} else if strings.ToLower(choice) == "" {
			// Weiter zur nächsten Seite
			page++
			if end == len(items) {
				console.InputPageEnd()
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
			console.ShowMessage("Möchten Sie diesen Artikel wirklich bearbeiten? (y/n) oder [c], um zum Hauptmenü zurückzukehren")

			choice = console.AskForInput()
			if strings.ToLower(choice) == "y" {
				// Die Eingabewerte werden jetzt nur einmal initialisiert und bei Korrekturen wiederverwendet
				newName = console.AskForName(item.Name, isEditing)
				newModel = console.AskForModel(item.Model, isEditing)
				newNotes = console.AskForNotes(item.Note, isEditing)

				// Bestätigung zum Bearbeiten des Artikels
				console.ShowMessage("Bitte überprüfen Sie die neuen Daten:")
				console.ShowMessage(fmt.Sprintf("Artikelbezeichnung: %s", newName))
				console.ShowMessage(fmt.Sprintf("Artikelnummer: %s", newModel))
				console.ShowMessage(fmt.Sprintf("Notizen: %s", newNotes))
				console.ShowMessage("\nSind die Daten korrekt? (y/n) oder [c] um zum Hauptmenü zurückzukehren.")

				choice = console.AskForInput()

				if strings.ToLower(choice) == "y" {
					// Artikel aktualisieren
					data := models.Item{
						Name:     newName,
						Model:    newModel,
						Note:     newNotes,
						Quantity: item.Quantity,
					}
					// Hier wird der Index korrekt angepasst
					err := models.UpdateItem(rowId-1, data)
					if err != nil {
						console.ShowError(err)
					} else {
						console.ShowMessage("✅ Artikel erfolgreich aktualisiert!")
						console.ShowContinue()
						console.Clear()
						console.ShowExecuteCommandMenu()
						return
					}
				} else if strings.ToLower(choice) == "c" {
					//Abbrechen und zurück zum Menü
					console.InputC()
					return
				} else if strings.ToLower(choice) == "n" {
					// Artikel nicht löschen, Abbruch
					console.ShowMessage("❌ Artikel wird nicht bearbeitet.")
					console.ShowContinue()
					console.Clear()
					break
				} else if strings.ToLower(choice) == "c" {
					console.ShowMessage("Vorgang abgebrochen")
					console.Clear()
					console.ShowExecuteCommandMenu()
					return
				} else {
					// Ungültige Eingabe, erneut fragen
					console.ShowMessage(fmt.Sprintf("Ungültige Eingabe. Bitte wählen:\n[y] zum Löschen\n[n] zum Behalten\n[c] um zum Hauptmenü zurückzukehren.\n"))
				}
			}
		}
	}
}

// Case 09
// handleViewItems Displays all items in the inventory.
func handleViewItems() {
	console.Clear()
	items := models.GetAllItems() // Entferne "err" und den zweiten Rückgabewert

	if console.ChecksInventory() {
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
			console.InputC()
			return
		} else if choice == "" {
			// Weiter zur nächsten Seite
			page++
			if end == len(items) {
				console.InputPageEnd()
				return
			}
		}
	}
}

//
