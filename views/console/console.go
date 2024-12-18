package console

import (
	"bufio"
	"fmt"
	"it_inventar/models"
	"log"
	"os"
	"os/exec"
	"strings"
)

const ExitStatusCodeNoError int = 0

// ShowExecuteCommandMenu shows the menu to the console
func ShowExecuteCommandMenu() {
	fmt.Println(`
	###########################################
	#******** WELCOME TO OUR INVENTORY ********
	#******** CHOOSE YOUR OPTION BELOW ********
	# 1. Artikel hinzufügen
	# 2. Artikel löschen
	# 3. Artikel bearbeiten
	# 4. Artikelinformationen ändern
	#
	#
	# 9. Lagerbestand anzeigen
	#
	# c. CLEAR VIEW AND SHOW MENU
	# q. EXIT INVENTORY APP
	`)
}

// ShowHiddenCommandMenu shows the hidden-menu to the console
func ShowHiddenCommandMenu() {
	fmt.Println(`
	###########################################
	#****************  SERVICE *****************
	#******** CHOOSE YOUR OPTION BELOW *********
	# 1. neuen Lieferant hinzufügen
	# 2. neuen Status hinzufügen
	# 3. neue Kategorie hinzufügen
	#
	# 11. Lieferant anzeigen
	# 12. Status anzeigen
	# 13. Kategorie anzeigen
	#
	#
	# c. SHOW MAIN MENU
	`)
}

// ShowAllItems shows all the available books in the library to the console
func ShowAllItems(items []models.Item, startIndex int) {
	// Berechnung der maximalen Länge für jede Spalte
	maxNameLen := len("Artikelbezeichnung")
	maxModelLen := len("Artikelnummer")
	maxQuantityLen := len("Menge")
	maxNoteLen := len("Notizen")

	// Durchlaufen der Items, um die maximale Länge für jede Spalte zu finden
	for _, item := range items {
		if len(item.Name) > maxNameLen {
			maxNameLen = len(item.Name)
		}
		if len(item.Model) > maxModelLen {
			maxModelLen = len(item.Model)
		}
		if len(fmt.Sprintf("%d", item.Quantity)) > maxQuantityLen {
			maxQuantityLen = len(fmt.Sprintf("%d", item.Quantity))
		}
		if len(item.Note) > maxNoteLen {
			maxNoteLen = len(item.Note)
		}
	}

	// Kopfzeile mit dynamisch berechneten Spaltenbreiten anzeigen
	fmt.Printf("%5s | %-*s | %-*s | %-*s | %-*s |\n",
		"ID", maxNameLen, "Artikelbezeichnung", maxModelLen, "Artikelnummer",
		maxQuantityLen, "Menge", maxNoteLen, "Notizen")
	fmt.Println(strings.Repeat("-", maxNameLen+maxModelLen+maxQuantityLen+maxNoteLen+22)) // Dynamische Trennlinie

	// Artikel anzeigen mit fortlaufender ID
	for index, item := range items {
		fmt.Printf("%5d | %-*s | %-*s | %-*d | %-*s |\n",
			startIndex+index+1, maxNameLen, item.Name, maxModelLen, item.Model,
			maxQuantityLen, item.Quantity, maxNoteLen, item.Note)
	}
}

// AskForInput reads from console until line break is available and returns input
func AskForInput() string {
	reader := bufio.NewReader(os.Stdin)
	response, err := reader.ReadString('\n')
	checkAndHandleError(err)
	return strings.TrimSpace(response)
}

// checkAndHandleError is used to check errors and handle them accordingly. In detail:
func checkAndHandleError(err error) {
	if err != nil {
		ShowError(err)
	}
}

// ShowError shows the error to the console with the current timestamp
func ShowError(err error) {
	if err != nil {
		log.Println("error:", err.Error())
	}
}

// Clear clears the console view
func Clear() {
	c := exec.Command("cmd", "/c", "cls")
	c.Stdout = os.Stdout
	err := c.Run()
	checkAndHandleError(err)
}

// ShowContinue shows the continuation information to the console
func ShowContinue() {
	fmt.Println("Drücken Sie [Enter], um fortzufahren...")
	bufio.NewReader(os.Stdin).ReadString('\n')
}

// ShowGoodbye shows a goodbye message to the console
func ShowGoodbye() {
	fmt.Println("Goodbye!")
}

// ShutDownNormal terminates the application with exit status code 0
func ShutDownNormal() {
	os.Exit(ExitStatusCodeNoError)
}

// ShowAddItemInformation shows the information and format about adding a Item
func ShowAddItemInformation() {
	fmt.Println("Bitte geben Sie die Artikeldaten im folgenden Format ein:")
}

// ShowRemoveItemInformation shows the information and format about removing a Item
func ShowRemoveItemInformation() {
	fmt.Println("Bitte geben Sie die Zeilennummer des zu entfernenden Artikels ein:")
}

// ShowMessage shows the message to the console
func ShowMessage(message string) {
	fmt.Println(message)
}

func ShowMassageData(message string) {
	fmt.Println(message, "darf nicht leer sein. Bitte versuchen sie es erneut.")
}

/*
// AskForName allows input for the item name, with an optional previous value if editing
func AskForName(defaultValue string, isEditing bool) string {
	for {
		if isEditing {
			ShowMessage(fmt.Sprintf("* Artikelbezeichnung [Eingegeben: %s]:", defaultValue))
		} else {
			ShowMessage("* Artikelbezeichnung:")
		}
		name := AskForInput()
		if name == "" && defaultValue != "" {
			return defaultValue // Verwende den alten Wert, wenn nichts eingegeben wurde
		} else if name != "" {
			return name // Verwende den neuen eingegebenen Wert
		} else {
			ShowMessage("⚠️ Artikelbezeichnung darf nicht leer sein.")
		}
	}
}

// AskForModel allows input for the item model, with an optional previous value if editing
func AskForModel(defaultValue string, isEditing bool) string {
	for {
		if isEditing {
			ShowMessage(fmt.Sprintf("* Artikelnummer [eingegeben: %s]:", defaultValue))
		} else {
			ShowMessage("* Artikelnummer:")
		}
		itemModel := AskForInput()
		if itemModel == "" && defaultValue != "" {
			return defaultValue // Verwende den alten Wert, wenn nichts eingegeben wurde
		} else if itemModel != "" {
			return itemModel // Verwende den neuen eingegebenen Wert
		} else {
			ShowMessage("⚠️ Artikelnummer darf nicht leer sein.")
		}
	}
}

// AskForQuantity allows input for the item quantity, with an optional previous value if editing
func AskForQuantity(defaultValue int, isEditing bool) int {
	for {
		if isEditing && defaultValue > 0 {
			ShowMessage(fmt.Sprintf("* Menge [eingegeben: %d]:", defaultValue))
		} else {
			ShowMessage("* Menge:")
		}
		quantityInput := AskForInput()
		if quantityInput == "" && defaultValue > 0 {
			return defaultValue // Verwende den alten Wert, wenn nichts eingegeben wurde
		}
		if strings.TrimSpace(quantityInput) == "" {
			ShowMessage("⚠️ Menge darf nicht leer sein.")
			continue
		}
		quantity := models.StringToInt(quantityInput)
		if quantity > 0 {
			return quantity
		} else {
			ShowMessage("⚠️ Menge muss eine positive Zahl sein. Bitte versuchen Sie es erneut.")
		}
	}
}

func AskForNotes(defaultValue string, isEditing bool) string {
	if isEditing && defaultValue != "" {
		ShowMessage(fmt.Sprintf("Eingegeben: (%s)\n\"Enter\" übernehmen, \"Leertaste\" löschen\n * Notizen:", defaultValue))
	} else {
		ShowMessage("* Notizen (optional):")
	}
	note := AskForInput()

	// Wenn nur Leerzeichen eingegeben wurden, dann leeren wir das Feld
	if strings.TrimSpace(note) == "" {
		return "" // Leert das Feld, wenn nur Leerzeichen eingegeben wurden
	} else if note == "" && defaultValue != "" {
		// Wenn wirklich nichts eingegeben wurde, aber es einen Default-Wert gibt, wird dieser verwendet
		return defaultValue
	}

	// Ansonsten verwenden wir die neue Eingabe
	return note
}
*/

/* // Angepasste version von oben
// ShowMessageData displays an error message indicating that a field cannot be empty.
func ShowMessageData(fieldName string) {
	fmt.Printf("%s darf nicht leer sein. Bitte versuchen Sie es erneut.\n", fieldName)
}

// AskForName allows input for the item name, with an optional previous value if editing
func AskForName(defaultValue string, isEditing bool) string {
	return askForInput("Artikelbezeichnung", defaultValue, isEditing, func(input string) bool {
		return input != ""
	})
}

// AskForModel allows input for the item model, with an optional previous value if editing
func AskForModel(defaultValue string, isEditing bool) string {
	return askForInput("Artikelnummer", defaultValue, isEditing, func(input string) bool {
		return input != ""
	})
}

// AskForQuantity allows input for the item quantity, with an optional previous value if editing
func AskForQuantity(defaultValue int, isEditing bool) int {
	for {
		prompt := "Menge"
		if isEditing && defaultValue > 0 {
			ShowMessage(fmt.Sprintf("* %s [eingegeben: %d]:", prompt, defaultValue))
		} else {
			ShowMessage(fmt.Sprintf("* %s:", prompt))
		}

		quantityInput := AskForInput()
		if quantityInput == "" && defaultValue > 0 {
			return defaultValue // Verwende den alten Wert, wenn nichts eingegeben wurde
		}

		if strings.TrimSpace(quantityInput) == "" {
			ShowMessageData(prompt)
			continue
		}

		quantity := models.StringToInt(quantityInput)
		if quantity > 0 {
			return quantity
		} else {
			ShowMessage("⚠️ Menge muss eine positive Zahl sein. Bitte versuchen Sie es erneut.")
		}
	}
}

// AskForNotes allows input for optional notes, with an optional previous value if editing
func AskForNotes(defaultValue string, isEditing bool) string {
	if isEditing && defaultValue != "" {
		ShowMessage(fmt.Sprintf("Eingegeben: (%s)\n\"Enter\" übernehmen, \"Leertaste\" löschen\n * Notizen:", defaultValue))
	} else {
		ShowMessage("* Notizen (optional):")
	}

	note := AskForInput()
	if strings.TrimSpace(note) == "" {
		return "" // Leert das Feld, wenn nur Leerzeichen eingegeben wurden
	} else if note == "" && defaultValue != "" {
		return defaultValue // Verwende den alten Wert, wenn nichts eingegeben wurde
	}
	return note // Verwende die neue Eingabe
}

// askForInput is a generic input handler for common input logic with validation.
func askForInput(fieldName string, defaultValue string, isEditing bool, validate func(string) bool) string {
	for {
		if isEditing {
			ShowMessage(fmt.Sprintf("* %s [Eingegeben: %s]:", fieldName, defaultValue))
		} else {
			ShowMessage(fmt.Sprintf("* %s:", fieldName))
		}

		input := AskForInput()
		if input == "" && defaultValue != "" {
			return defaultValue // Verwende den alten Wert, wenn nichts eingegeben wurde
		} else if validate(input) {
			return input // Verwende die neue Eingabe, wenn sie gültig ist
		} else {
			ShowMessageData(fieldName)
		}
	}
}
*/

// ShowMessageData displays an error message indicating that a field cannot be empty.
func ShowMessageData(fieldName string) {
	fmt.Printf("%s darf nicht leer sein. Bitte versuchen Sie es erneut.\n", fieldName)
}

// AskForName allows input for the item name, with an optional previous value if editing
func AskForName(defaultValue string, isEditing bool) string {
	return askForInput("Artikelbezeichnung", defaultValue, isEditing, func(input string) bool {
		return input != ""
	})
}

// AskForModel allows input for the item model, with an optional previous value if editing
func AskForModel(defaultValue string, isEditing bool) string {
	return askForInput("Artikelnummer", defaultValue, isEditing, func(input string) bool {
		return input != ""
	})
}

// AskForQuantity allows input for the item quantity, with an optional previous value if editing
func AskForQuantity(defaultValue int, isEditing bool) int {
	for {
		prompt := "Menge"
		if isEditing && defaultValue > 0 {
			ShowMessage(fmt.Sprintf("* %s [eingegeben: %d]:", prompt, defaultValue))
		} else {
			ShowMessage(fmt.Sprintf("* %s:", prompt))
		}

		quantityInput := AskForInput()
		if quantityInput == "" && defaultValue > 0 {
			return defaultValue // Verwende den alten Wert, wenn nichts eingegeben wurde
		}

		if strings.TrimSpace(quantityInput) == "" {
			ShowMessageData(prompt)
			continue
		}

		quantity := models.StringToInt(quantityInput)
		if quantity > 0 {
			return quantity
		} else {
			ShowMessage("⚠️ Menge muss eine positive Zahl sein. Bitte versuchen Sie es erneut.")
		}
	}
}

// AskForNotes allows input for optional notes, with an optional previous value if editing
func AskForNotes(defaultValue string, isEditing bool) string {
	if isEditing && defaultValue != "" {
		ShowMessage(fmt.Sprintf("Eingegeben: (%s)\n\"Enter\" übernehmen, \"Leertaste\" löschen\n * Notizen:", defaultValue))
	} else {
		ShowMessage("* Notizen (optional):")
	}

	note := AskForInput()
	if strings.TrimSpace(note) == "" {
		return "" // Leert das Feld, wenn nur Leerzeichen eingegeben wurden
	} else if note == "" && defaultValue != "" {
		return defaultValue // Verwende den alten Wert, wenn nichts eingegeben wurde
	}
	return note // Verwende die neue Eingabe
}

// AskForSupplier allows input for the supplier, with an optional previous value if editing
func AskForSupplier(defaultValue string, isEditing bool) string {
	return askForInput("Lieferant", defaultValue, isEditing, func(input string) bool {
		return input != ""
	})
}

// AskForStatus allows input for the status, with an optional previous value if editing
func AskForStatus(defaultValue string, isEditing bool) string {
	return askForInput("Status", defaultValue, isEditing, func(input string) bool {
		return input != ""
	})
}

// AskForCategory allows input for the category, with an optional previous value if editing
func AskForCategory(defaultValue string, isEditing bool) string {
	return askForInput("Kategorie", defaultValue, isEditing, func(input string) bool {
		return input != ""
	})
}

// askForInput is a generic input handler for common input logic with validation.
func askForInput(fieldName string, defaultValue string, isEditing bool, validate func(string) bool) string {
	for {
		if isEditing {
			ShowMessage(fmt.Sprintf("* %s [Eingegeben: %s]:", fieldName, defaultValue))
		} else {
			ShowMessage(fmt.Sprintf("* %s:", fieldName))
		}

		input := AskForInput()
		if input == "" && defaultValue != "" {
			return defaultValue // Verwende den alten Wert, wenn nichts eingegeben wurde
		} else if validate(input) {
			return input // Verwende die neue Eingabe, wenn sie gültig ist
		} else {
			ShowMessageData(fieldName)
		}
	}
}
