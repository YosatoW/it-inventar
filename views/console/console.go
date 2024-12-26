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

const (
	ExitStatusCodeNoError int = 0
	// ItemDetailsMessage or the output of article information.
	ItemDetailsMessage = "Artikel: %s (%s) - %d Stück - Notizen: %s"
)

// ShowAllItems shows all the available books in the library to the console
func ShowAllItems(items []models.Item, startIndex int) {
	// Berechnung der maximalen Länge für jede Spalte
	maxArticleNameLen := len("Artikelbezeichnung")
	maxArticleNumberLen := len("Artikelnummer")
	maxSupplierLen := len("Lieferant")
	//maxManufacturerLen := len("Hersteller")
	manQuantityLen := len("Menge [Stk]")
	maxNoteLen := len("Notizen")

	// Durchlaufen der Items, um die maximale Länge für jede Spalte zu finden
	for _, item := range items {
		if len(item.ArticleName) > maxArticleNameLen {
			maxArticleNameLen = len(item.ArticleName)
		}
		if len(item.ArticleNumber) > maxArticleNumberLen {
			maxArticleNumberLen = len(item.ArticleNumber)
		}
		if len(item.Supplier) > maxSupplierLen {
			maxSupplierLen = len(item.Supplier)
		}
		if len(fmt.Sprintf("%d", item.Quantity)) > manQuantityLen {
			manQuantityLen = len(fmt.Sprintf("%d", item.Quantity))
		}
		if len(item.Note) > maxNoteLen {
			maxNoteLen = len(item.Note)
		}
	}

	// Kopfzeile mit dynamisch berechneten Spaltenbreiten anzeigen
	fmt.Printf("%5s | %-*s | %-*s | %-*s | %-*s | %-*s |\n",
		"ID",
		maxArticleNameLen, "Artikelbezeichnung",
		maxArticleNumberLen, "Artikelnummer",
		maxSupplierLen, "Lieferant",
		manQuantityLen, "Menge [Stk]",
		maxNoteLen, "Notizen")
	fmt.Println(strings.Repeat("-", maxArticleNameLen+maxArticleNumberLen+maxSupplierLen+manQuantityLen+maxNoteLen+25)) // Dynamische Trennlinie

	// Artikel anzeigen mit fortlaufender ID
	for index, item := range items {
		fmt.Printf("%5d | %-*s | %-*s | %-*s | %-*d | %-*s |\n",
			startIndex+index+1,
			maxArticleNameLen, item.ArticleName,
			maxArticleNumberLen, item.ArticleNumber,
			maxSupplierLen, item.Supplier,
			manQuantityLen, item.Quantity,
			maxNoteLen, item.Note)
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
	fmt.Println("Bitte geben Sie die Artikeldaten ein:")
}

// ShowMessage shows the message to the console
func ShowMessage(message string) {
	fmt.Println(message)
}

func ConfirmTheArticle(item models.Item) string {
	return fmt.Sprintf(ItemDetailsMessage, item.ArticleName, item.ArticleNumber, item.Quantity, item.Note)
}

// InputC Displays the main menu and returns to it.
func InputC() bool {
	Clear()
	ShowExecuteCommandMenu()
	return true
}

// InputPageEnd Notifies the user and returns to the main menu.
func InputPageEnd() bool {
	// Wenn es keine weiteren Artikel mehr gibt
	ShowMessage("Alle Artikel wurden angezeigt.")
	ShowContinue()
	Clear()
	ShowExecuteCommandMenu()
	return true
}

// ChecksInventory whether the inventory (items) is empty.
func ChecksInventory() bool {
	if len(models.GetAllItems()) == 0 {
		ShowMessage("❌ Es sind keine Artikel im Inventar vorhanden.")
		ShowContinue()
		Clear()
		ShowExecuteCommandMenu()
		return true
	}
	return false
}

func HandleChancelAction() bool {
	ShowMessage("❌ Artikel wurde nicht geändert.")
	ShowContinue()
	Clear()
	return true
}

// PageIndexCalculate Calculates the start and end index for a page navigation, limited to the total number of elements.
func PageIndexCalculate(page, pageSize, totalItems int) (int, int) {
	start := page * pageSize
	end := start + pageSize
	if end > totalItems {
		end = totalItems
	}
	return start, end
}

// PageIndexPrompt Shows the prompt for editing, scrolling or canceling
func PageIndexPrompt() string {
	ShowMessage("Gib die ID des zu bearbeitenden Artikels ein, drücke [Enter] für nächste seite oder [c], um zum Hauptmenü zurückzukehren.")
	return AskForInput()
}

// PageIndexUserInput Processes the user input for editing, scrolling or canceling
func PageIndexUserInput(choice string, page *int, end int, items []models.Item) (bool, *models.Item, int) {
	if strings.ToLower(choice) == "c" {
		Clear()
		ShowExecuteCommandMenu()
		return true, nil, 0
	} else if strings.ToLower(choice) == "" {
		// Continue to the next page
		(*page)++
		if end == len(items) {
			InputPageEnd()
			return true, nil, 0
		}
	} else {
		// Check whether the input is a valid ID
		rowId := models.StringToInt(choice)
		if rowId <= 0 || rowId > len(items) {
			ShowMessage("❌ Ungültige ID. Bitte gib eine gültige ID ein.")
			ShowContinue()
			return false, nil, 0
		}

		// Check whether the article exists and display
		item := models.GetItemById(rowId - 1) // The index is adjusted correctly here
		if item == nil {
			ShowMessage("❌ Artikel mit dieser ID existiert nicht.")
			ShowContinue()
			return false, nil, 0
		}
		return false, item, rowId
	}
	return false, nil, 0
}

// ShowMessageData displays an error message indicating that a field cannot be empty.
func ShowMessageData(fieldName string) {
	fmt.Printf("%s darf nicht leer sein.\n", fieldName)
}

// AskForName allows input for the item name, with an optional previous value if editing
func AskForName(defaultValue string, isEditing bool) string {
	return askForInput("Artikelbezeichnung", defaultValue, isEditing, func(input string) bool {
		return input != ""
	})
}

// AskForArticleNumber allows input for the item model, with an optional previous value if editing
func AskForArticleNumber(defaultValue string, isEditing bool) string {
	return askForInput("Artikelnummer", defaultValue, isEditing, func(input string) bool {
		return input != ""
	})
}

// AskForQuantity allows input for the item quantity, with an optional previous value if editing
func AskForQuantity(defaultValue int, isEditing bool) int {
	for {
		prompt := "Menge"
		if isEditing && defaultValue >= 0 {
			ShowMessage(fmt.Sprintf("* %s [Eingegeben: %d]:", prompt, defaultValue))
		} else {
			ShowMessage(fmt.Sprintf("* %s:", prompt))
		}

		quantityInput := AskForInput()
		if quantityInput == "" && defaultValue >= 0 {
			return defaultValue // Verwende den alten Wert, wenn nichts eingegeben wurde
		}

		if strings.TrimSpace(quantityInput) == "" {
			ShowMessageData(prompt)
			continue
		}

		quantity := models.StringToInt(quantityInput)
		if quantity >= 0 {
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

// AskForStatus allows input for the item name, with an optional previous value if editing
func AskForStatus(defaultValue string, isEditing bool) string {
	return askForInput("Status", defaultValue, isEditing, func(input string) bool {
		return input != ""
	})
}

// AskForSupplier allows input for the supplier, with an optional previous value if editing
func AskForSupplier(defaultValue string, isEditing bool) string {
	return askForInput("Lieferant", defaultValue, isEditing, func(input string) bool {
		return input != ""
	})
}

// AskForManufacturer allows input for the manufacturer, with an optional previous value if editing
func AskForManufacturer(defaultValue string, isEditing bool) string {
	return askForInput("Hersteller", defaultValue, isEditing, func(input string) bool {
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
