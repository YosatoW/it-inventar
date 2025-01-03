package console

import (
	"bufio"
	"fmt"
	"it_inventar/models"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

const (
	InitialPage = 0
	PageSize    = 10

	ExitStatusCodeNoError int = 0
	// ItemDetailsMessage or the output of article information.
	ItemDetailsMessage = "Item: %s | Category: %s (%s) | %d pieces | Notes: %s"
)

// *ShowAllItems: Displays all items in the inventory with dynamically calculated column widths for better readability.
// *ShowAllItems: Zeigt alle Artikel im Inventar mit dynamisch berechneten Spaltenbreiten für bessere Lesbarkeit an.
func ShowAllItems(items []models.Item, startIndex int) {
	// Calculate the maximum length for each column
	maxArticleNameLen := len("Item Name")
	maxArticleCategoryLen := len("Category")
	maxArticleNumberLen := len("Item No.")
	maxSupplierLen := len("Supplier")
	maxQuantityLen := len("Quantity [pcs]")
	maxNoteLen := len("Notes")

	// Iterate through the items to find the maximum length for each column
	for _, item := range items {
		if len(item.ArticleName) > maxArticleNameLen {
			maxArticleNameLen = len(item.ArticleName)
		}
		if len(item.Category) > maxArticleCategoryLen {
			maxArticleCategoryLen = len(item.Category)
		}
		if len(item.ArticleNumber) > maxArticleNumberLen {
			maxArticleNumberLen = len(item.ArticleNumber)
		}
		if len(item.Supplier) > maxSupplierLen {
			maxSupplierLen = len(item.Supplier)
		}
		if len(fmt.Sprintf("%d", item.Quantity)) > maxQuantityLen {
			maxQuantityLen = len(fmt.Sprintf("%d", item.Quantity))
		}
		if len(item.Note) > maxNoteLen {
			maxNoteLen = len(item.Note)
		}
	}

	// Display header with dynamically calculated column widths
	fmt.Printf("%5s | %-*s | %-*s | %-*s | %-*s | %-*s | %-*s |\n",
		"ID",
		maxArticleNameLen, "Item Name",
		maxArticleCategoryLen, "Category",
		maxArticleNumberLen, "Item No.",
		maxSupplierLen, "Supplier",
		maxQuantityLen, "Quantity [pcs]",
		maxNoteLen, "Notes")
	ShowMessage(strings.Repeat("-", maxArticleNameLen+maxArticleCategoryLen+maxArticleNumberLen+maxSupplierLen+maxQuantityLen+maxNoteLen+25))

	// Display items with sequential ID
	for index, item := range items {
		fmt.Printf("%5d | %-*s | %-*s | %-*s | %-*s | %-*d | %-*s |\n",
			startIndex+index+1,
			maxArticleNameLen, item.ArticleName,
			maxArticleCategoryLen, item.Category,
			maxArticleNumberLen, item.ArticleNumber,
			maxSupplierLen, item.Supplier,
			maxQuantityLen, item.Quantity,
			maxNoteLen, item.Note)
	}
}

// *AskForInput: Reads user input from the console until a line break is detected and returns the input.
// *AskForInput: Liest die Benutzereingabe von der Konsole, bis ein Zeilenumbruch erkannt wird, und gibt die Eingabe zurück.
func AskForInput() string {
	reader := bufio.NewReader(os.Stdin)
	response, err := reader.ReadString('\n')
	checkAndHandleError(err)
	return strings.TrimSpace(response)
}

// *checkAndHandleError: Checks for errors and handles them by displaying an error message.
// *checkAndHandleError: Überprüft auf Fehler und behandelt sie, indem eine Fehlermeldung angezeigt wird.
func checkAndHandleError(err error) {
	if err != nil {
		ShowError(err)
	}
}

// *ShowError: Displays the error message along with the current timestamp.
// *ShowError: Zeigt die Fehlermeldung zusammen mit dem aktuellen Zeitstempel an.
func ShowError(err error) {
	if err != nil {
		log.Println("error:", err.Error())
	}
}

// *Clear Clears the console screen.
// *Clear Löscht den Konsolenbildschirm.
func Clear() {
	c := exec.Command("cmd", "/c", "cls")
	c.Stdout = os.Stdout
	err := c.Run()
	checkAndHandleError(err)
}

// *ShowContinue: Prompts the user to press Enter to continue.
// *ShowContinue:  Fordert den Benutzer auf, Enter zu drücken, um fortzufahren.
func ShowContinue() {
	ShowMessage("Press [Enter] to continue...")
	bufio.NewReader(os.Stdin).ReadString('\n')
}

// *ShowGoodbye: Displays a goodbye message.
// *ShowGoodbye: Zeigt eine Abschiedsnachricht an.
func ShowGoodbye() {
	ShowMessage("Goodbye!")
}

// *ShutDownNormal: Terminates the application with a normal exit status code.
// *ShutDownNormal: Beendet die Anwendung mit einem normalen Exit-Status-Code.
func ShutDownNormal() {
	os.Exit(ExitStatusCodeNoError)
}

// *ShowAddItemInformation: Displays instructions for adding a new item.
// *ShowAddItemInformation: Zeigt Anweisungen zum Hinzufügen eines neuen Artikels an.
func ShowAddItemInformation() {
	ShowMessage("Please enter the item details:")
}

// *ShowMessage: Displays a message on the console.
// *ShowMessage: Zeigt eine Nachricht auf der Konsole an.
func ShowMessage(message string) {
	fmt.Println(message)
}

// *ConfirmTheArticle: Returns a formatted string with the details of the specified item.
// *ConfirmTheArticle: Gibt eine formatierte Zeichenkette mit den Details des angegebenen Artikels zurück.
func ConfirmTheArticle(item models.Item) string {
	return fmt.Sprintf(ItemDetailsMessage, item.ArticleName, item.Category, item.ArticleNumber, item.Quantity, item.Note)
}

// *InputC: Clears the screen and displays the main menu.
// *InputC: Löscht den Bildschirm und zeigt das Hauptmenü an.
func InputC() bool {
	Clear()
	ShowExecuteCommandMenu()
	return true
}

// *InputPageEnd: Notifies the user that all items have been displayed and returns to the main menu.
// *InputPageEnd: Benachrichtigt den Benutzer, dass alle Artikel angezeigt wurden, und kehrt zum Hauptmenü zurück.
func InputPageEnd() bool {
	// Wenn es keine weiteren Artikel mehr gibt
	ShowMessage("All items have been displayed.")
	ShowContinue()
	Clear()
	ShowExecuteCommandMenu()
	return true
}

// *ChecksInventory Checks if the inventory is empty and returns to the main menu if it is.
// *ChecksInventory Überprüft, ob das Inventar leer ist, und kehrt zum Hauptmenü zurück, wenn es leer ist.
func ChecksInventory() bool {
	if len(models.GetAllItems()) == 0 {
		ShowMessage("❌ No items available.")
		ShowContinue()
		Clear()
		ShowExecuteCommandMenu()
		return true
	}
	return false
}

// *HandleChancelAction: Displays a message that the item was not changed and returns to the main menu.
// *HandleChancelAction: Zeigt eine Nachricht an, dass der Artikel nicht geändert wurde, und kehrt zum Hauptmenü zurück.
func HandleChancelAction() bool {
	ShowMessage("❌ Item was not changed.")
	ShowContinue()
	Clear()
	return true
}

// *PageIndexCalculate: Calculates the start and end indices for pagination based on the current page and page size.
// *PageIndexCalculate: Berechnet die Start- und Endindizes für die Seitennavigation basierend auf der aktuellen Seite und der Seitengröße.
func PageIndexCalculate(page, pageSize, totalItems int) (int, int) {
	start := page * pageSize
	end := start + pageSize
	if end > totalItems {
		end = totalItems
	}
	return start, end
}

// *PageIndexPrompt: Prompts the user to enter the ID of the item or navigate to the next page.
// *PageIndexPrompt: Fordert den Benutzer auf, die ID des Artikels einzugeben oder zur nächsten Seite zu navigieren.
func PageIndexPrompt(itemType string) string {
	fmt.Printf("Enter the ID of the %s, press [Enter] for next page or [c] to return to the main menu.\n", itemType)
	return AskForInput()
}

// *PageIndexView: Prompts the user to press Enter for the next page or 'c' to cancel.
// *PageIndexView: Fordert den Benutzer auf, Enter für die nächste Seite oder 'c' zum Abbrechen zu drücken.
func PageIndexView() string {
	ShowMessage("Press [Enter] for next page or [c] to return to the main menu.")
	return AskForInput()
}

// *MessageGeneralInvalidID: Displays a message indicating that the entered ID is invalid.
// *MessageGeneralInvalidID: Zeigt eine Nachricht an, dass die eingegebene ID ungültig ist.
func MessageGeneralInvalidID() {
	ShowMessage("❌ Invalid ID. Please enter a valid ID.")
}

// MessageGeneralNotEmpty: Displays an error message indicating that the specified field cannot be empty.
// MessageGeneralNotEmpty: Zeigt eine Fehlermeldung an, dass das angegebene Feld nicht leer sein darf.
func MessageGeneralNotEmpty(fieldName string) {
	fmt.Printf("%s cannot be empty.\n", fieldName)
}

// *PageIndexUserInput: Processes user input for editing an item, navigating pages, or canceling the action.
// *PageIndexUserInput: Verarbeitet die Benutzereingabe zum Bearbeiten eines Artikels, zum Navigieren durch Seiten oder zum Abbrechen der Aktion.
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
		if rowId < 1 || rowId > len(items) {
			MessageGeneralInvalidID()
			ShowContinue()
			return false, nil, 0
		}
		return false, &items[rowId-1], rowId
	}
	return false, nil, 0
}

// *AskForName: Prompts the user to enter the item name, with an optional default value if editing.
// *AskForName: Fordert den Benutzer auf, den Artikelnamen einzugeben, mit einem optionalen Standardwert, wenn bearbeitet wird.
func AskForName(defaultValue string, isEditing bool) string {
	return askForInput("Item name", defaultValue, isEditing, func(input string) bool {
		return input != ""
	})
}

// *AskForArticleNumber: Prompts the user to enter the category, with an optional default value if editing.
// *AskForArticleNumber: Fordert den Benutzer auf, die Kategorie einzugeben, mit einem optionalen Standardwert, wenn bearbeitet wird.
func AskForCategory(defaultValue string, isEditing bool) string {
	return askForInput("Category", defaultValue, isEditing, func(input string) bool {
		return input != ""
	})
}

// *AskForArticleNumber: Prompts the user to enter the item number, with an optional default value if editing.
// *AskForArticleNumber: Fordert den Benutzer auf, die Artikelnummer
func AskForArticleNumber(defaultValue string, isEditing bool) string {
	return askForInput("Item number", defaultValue, isEditing, func(input string) bool {
		return input != ""
	})
}

// *AskForSupplier: Prompts the user to enter the supplier, with an optional default value if editing.
// *AskForSupplier: Fordert den Benutzer auf, den Lieferanten einzugeben, mit einem optionalen Standardwert, wenn bearbeitet wird.
func AskForSupplier(defaultValue string, isEditing bool) string {
	return askForInput("Supplier", defaultValue, isEditing, func(input string) bool {
		return input != ""
	})
}

// *AskForQuantity: Prompts the user to enter the item quantity, with an optional default value if editing.
// *AskForQuantity: Fordert den Benutzer auf, die Menge des Artikels einzugeben, mit einem optionalen Standardwert, wenn bearbeitet wird.
func AskForQuantity(defaultValue int, isEditing bool) int {
	for {
		prompt := "Quantity"
		if isEditing && defaultValue >= 0 {
			ShowMessage(fmt.Sprintf("* %s [Entered: %d]:", prompt, defaultValue))
		} else {
			ShowMessage(fmt.Sprintf("* %s:", prompt))
		}

		quantityInput := AskForInput()
		if quantityInput == "" && defaultValue >= 0 {
			return defaultValue // Verwende den alten Wert, wenn nichts eingegeben wurde
		}

		if strings.TrimSpace(quantityInput) == "" {
			MessageGeneralNotEmpty(prompt)
			continue
		}

		quantity := models.StringToInt(quantityInput)
		if quantity >= 0 {
			return quantity
		} else {
			ShowMessage("⚠️ Quantity must be a positive number. Please try again.")
		}
	}
}

// *AskForNotes: Prompts the user to enter optional notes, with an optional default value if editing.
// *AskForNotes: Fordert den Benutzer auf, optionale Notizen einzugeben, mit einem optionalen Standardwert, wenn bearbeitet wird.
func AskForNotes(defaultValue string, isEditing bool) string {
	if isEditing && defaultValue != "" {
		ShowMessage(fmt.Sprintf("Entered: (%s)\n\"Enter\" to keep, \"Space\" to clear\n * Notes:", defaultValue))
	} else {
		ShowMessage("* Notes (optional):")
	}

	note := AskForInput()
	if strings.TrimSpace(note) == "" {
		return "" // Clear the field if only spaces were entered
	} else if note == "" && defaultValue != "" {
		return defaultValue // Use the old value if nothing was entered
	}
	return note // Use the new input
}

// *askForInput: A generic input handler for common input logic with validation.
// *askForInput: Ein generischer Eingabe-Handler für allgemeine Eingabelogik mit Validierung.
func askForInput(fieldName string, defaultValue string, isEditing bool, validate func(string) bool) string {
	for {
		if isEditing {
			ShowMessage(fmt.Sprintf("* %s [Entered: %s]:", fieldName, defaultValue))
		} else {
			ShowMessage(fmt.Sprintf("* %s:", fieldName))
		}

		input := AskForInput()
		if input == "" && defaultValue != "" {
			return defaultValue // Use the old value if nothing was entered
		} else if validate(input) {
			return input // Use the new input if it is valid
		} else {
			MessageGeneralNotEmpty(fieldName)
		}
	}
}

// *SelectItem: Displays a paginated list of items and returns the selected item.
// *SelectItem: Zeigt eine paginierte Liste von Artikeln an und gibt den ausgewählten Artikel zurück.
func SelectItem(items []string, pageSize int, itemType string) string {
	page := 0
	totalPages := (len(items) + pageSize - 1) / pageSize

	for {
		start, end := PageIndexCalculate(page, pageSize, len(items))

		fmt.Printf("Please select a %s from the list:\n", itemType)
		for i := start; i < end; i++ {
			fmt.Printf("%d: %s\n", i+1, items[i])
		}

		if totalPages > 1 {
			fmt.Printf("Page %d of %d.\n", page+1, totalPages)
		}

		input := PageIndexPrompt(itemType)
		input = strings.TrimSpace(input)

		switch strings.ToLower(input) {
		case "c":
			ShowMessage("Action canceled. Press [Enter] to continue...")
			return ""
		case "":
			page = (page + 1) % totalPages
		default:
			choice, err := strconv.Atoi(input)
			if err == nil && choice > 0 && choice <= len(items) {
				return items[choice-1]
			}
			MessageGeneralInvalidID()
			ShowContinue()
		}
	}
}

//// SelectCategory zeigt die Kategorienauswahl in Seiten an und gibt die ausgewählte Kategorie zurück
//func SelectCategory(categories []string, pageSize int) string {
//	return SelectItem(categories, pageSize, "Kategorie")
//}
//
//// SelectSupplier zeigt die Lieferantenauswahl in Seiten an und gibt den ausgewählten Lieferanten zurück
//func SelectSupplier(suppliers []string, pageSize int) string {
//	return SelectItem(suppliers, pageSize, "Lieferant")
//}

// *HandleAddSelectItem: Checks if the user is in edit mode and displays the current selection before allowing a new selection.
// *HandleAddSelectItem: Überprüft, ob der Benutzer im Bearbeitungsmodus ist, und zeigt die aktuelle Auswahl an, bevor eine neue Auswahl getroffen wird.
func HandleAddSelectItem(currentItem string, items []string, itemType string, isEditing bool) string {
	if !isEditing {
		return SelectItem(items, PageSize, itemType)
	} else {
		ShowMessage(fmt.Sprintf("Current: %s", currentItem))
		newItem := SelectItem(items, PageSize, itemType)
		if newItem != "" {
			return newItem
		}
		return currentItem
	}
}

// DisplaySuppliers displays a paginated list of suppliers
func DisplaySuppliers(suppliers []string, start, end int) {
	fmt.Println("* Available suppliers:")
	for i := start; i < end && i < len(suppliers); i++ {
		fmt.Printf("%d. %s\n", i+1, suppliers[i]) // Add 1 to i for correct numbering
	}
}

// ShowSuppliersList shows the list of items with their index and details
func ShowSuppliersList(suppliers []string) {
	fmt.Println("* Showing Existing Suppliers *")
	for i, supplier := range suppliers {
		fmt.Printf("%d. %s\n", i+1, supplier)
	}
}

// ShowNoSuppliersMessage displays a message when no suppliers are available
func ShowNoSuppliersMessage() {
	fmt.Println("List is empty. No supplier available.")
}

// ShowEndOfSupplier indicate the end of the displayed list
func ShowEndOfSupplier() {
	fmt.Println("End of supplier list reached.")
}

// DisplayCategories displays a paginated list of categories
func DisplayCategories(suppliers []string, start, end int) {
	fmt.Println("* Available Categories:")
	for i := start; i < end && i < len(suppliers); i++ {
		fmt.Printf("%d. %s\n", i+1, suppliers[i]) // Add 1 to i for correct numbering
	}
}

// ShowCategoriesList shows the list of items with their index and details
func ShowCategoriesList(Categories []string) {
	fmt.Println("* Showing Existing Categories *")
	for i, category := range Categories {
		fmt.Printf("%d. %s\n", i+1, category)
	}
}

// ShowNoCategoriesMessage displays a message when no category are available
func ShowNoCategoriesMessage() {
	fmt.Println("List is empty. No Category available.")
}

// ShowEndOfCategory indicate the end of the displayed list
func ShowEndOfCategory() {
	fmt.Println("End of Category list reached.")
}

// ErrorMessage displays an error message
func ErrorMessage(message string) {
	fmt.Println("Error:", message)
}

// GetPageInput waits for the user to press Enter or type 'c' to cancel
func GetPageInput() string {
	fmt.Print("Press Enter to continue or 'c' to cancel:\n")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// ShowOnlyCancelMessage displays a message indicating that only 'c' can be pressed to return to the service menu
func ShowOnlyCancelMessage() string {
	fmt.Println("Press 'c' to continue to the service menu:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// ShowPrompt displays a message to prompt the user for input
func ShowPrompt(message string) {
	fmt.Println(message)
}

// GetUserInput reads and trims the user input from the console
func GetUserInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}
