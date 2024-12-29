package controllers

import (
	"fmt"
	"it_inventar/controllers/Category"
	"it_inventar/controllers/Supplier"
	"it_inventar/models"
	"it_inventar/views/console"
	"strconv"
	"strings"
)

const (
	InitialPage = console.InitialPage
	pageSize    = console.PageSize

	messageInvalidInput         = "Ungültige Eingabe!"
	messageInvalidInputTryAgain = "Bitte wählen:\n[y] Übernehmen\n[n] abbrechen."
)

// Run does the running of the console application
func Run() {
	checkAndHandleError(models.Initialize())
	console.Clear()
	console.ShowExecuteCommandMenu()

	for {
		executeCommand()
	}
}

// checkAndHandleError Checks whether there is an error and displays it if present. Used for centralized error handling.
func checkAndHandleError(err error) {
	if err != nil {
		console.ShowError(err)
		return
	}
}

// Case 01
// handleAddItem adds a new item to the inventory.
func handleAddItem() {
	console.Clear()
	console.ShowAddItemInformation()

	var isEditing bool = false
	var articleName, chosenCategory, articleNumber, chosenSupplier, notes string
	var quantity int

	selectedCategories, err := Category.ReadCategories(models.FileCategories)
	if err != nil {
		console.ShowError(err)
		return
	}
	selectedSuppliers, err := Supplier.ReadSuppliers(models.FileSupplier)
	if err != nil {
		console.ShowError(err)
		return
	}

	for {
		articleName = console.AskForName(articleName, isEditing)
		chosenCategory = console.HandleAddSelectItem(chosenCategory, selectedCategories, "Kategorie", isEditing)
		if chosenCategory == "" {
			return
		}
		articleNumber = console.AskForArticleNumber(articleNumber, isEditing)
		chosenSupplier = console.HandleAddSelectItem(chosenSupplier, selectedSuppliers, "Lieferant", isEditing)
		if chosenSupplier == "" {
			return
		}
		quantity = console.AskForQuantity(quantity, isEditing)
		notes = console.AskForNotes(notes, isEditing)

		if handleConfirmItemDetails(articleName, chosenCategory, articleNumber, chosenSupplier, quantity, notes) {
			data := models.Item{
				ArticleName:   articleName,
				Category:      chosenCategory,
				ArticleNumber: articleNumber,
				Supplier:      chosenSupplier,
				Quantity:      quantity,
				Note:          notes,
			}
			err := models.AddItem(data)
			if err != nil {
				console.ShowError(err)
			} else {
				console.ShowMessage("✅ Artikel erfolgreich hinzugefügt!")
				console.ShowContinue()
				console.InputC()
				return
			}
		} else {
			isEditing = true
		}
	}
}

// Case 02
// handleRemoveItem Handles the removal of an item from the inventory.
func handleRemoveItem() {
	console.Clear()
	items := models.GetAllItems() // Get all items from the inventory

	if console.ChecksInventory() { // Checks inventory for content
		return
	}

	page := InitialPage
	for {
		start, end := console.PageIndexCalculate(page, pageSize, len(items))

		console.ShowAllItems(items[start:end], start)

		choice := console.PageIndexPrompt("Artikel")

		exit, item, rowId := console.PageIndexUserInput(choice, &page, end, items)
		if exit {
			return
		}
		if item != nil {
			console.ShowMessage(fmt.Sprintf("%s\nDiesen Artikel löschen? (y/n)", console.ConfirmTheArticle(*item)))

			for {
				choice = console.AskForInput()

				if strings.ToLower(choice) == "y" {
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
					console.HandleChancelAction()
					break
				} else {
					// Ungültige Eingabe, erneut fragen
					console.Clear()
					console.ShowMessage(messageInvalidInput)
					console.ShowMessage(fmt.Sprintf("Artikel: %s (%s) - %d Stück - Notizen: %s", item.ArticleName, item.ArticleNumber, item.Quantity, item.Note))
					console.ShowMessage(messageInvalidInputTryAgain)
				}
			}
		}
	}
}

// Case 03
// handleChangeQuantity bearbeitet einen Artikel im Inventar
func handleChangeQuantity() {
	console.Clear()
	items := models.GetAllItems()

	if console.ChecksInventory() {
		return
	}

	page := InitialPage
	for {
		start, end := console.PageIndexCalculate(page, pageSize, len(items))

		console.ShowAllItems(items[start:end], start)

		choice := console.PageIndexPrompt("Artikel")

		exit, item, rowId := console.PageIndexUserInput(choice, &page, end, items)
		if exit {
			return
		}
		if item != nil {
			console.ShowMessage(fmt.Sprintf("%s\nDie Mende diesen Artikel anpassen?(y/n)", console.ConfirmTheArticle(*item)))

			for {
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
				} else if strings.ToLower(choice) == "n" {
					console.HandleChancelAction()
					break

				} else {
					// Ungültige Eingabe, erneut fragen
					console.Clear()
					console.ShowMessage(messageInvalidInput)
					console.ShowMessage(fmt.Sprintf("Artikel: %s (%s) - %d Stück - Notizen: %s", item.ArticleName, item.ArticleNumber, item.Quantity, item.Note))
					console.ShowMessage(messageInvalidInputTryAgain)
				}
			}
		}
	}
}

// Case 04
// handleChanceArticleInformation bearbeitet einen Artikel im Inventar
func handleChanceArticleInformation() {
	console.Clear()
	items := models.GetAllItems()

	if console.ChecksInventory() {
		return
	}

	page := InitialPage
	for {
		var isEditing bool = false
		var NewArticleName, newCategory, newArticleNumber, newSupplier, newNotes string

		start, end := console.PageIndexCalculate(page, pageSize, len(items))

		console.ShowAllItems(items[start:end], start)

		choice := console.PageIndexPrompt("Artikel")

		exit, item, rowId := console.PageIndexUserInput(choice, &page, end, items)
		if exit {
			return
		}
		if item != nil {

			// Die Eingabewerte werden jetzt nur einmal initialisiert und bei Korrekturen wiederverwendet
			console.ShowMessage(fmt.Sprintf("Aktuelle Artikelbezeichnung: %s", item.ArticleName))
			NewArticleName = console.AskForName(item.ArticleName, isEditing)

			// Kategorien und Lieferanten laden
			selectedCategories, err := Category.ReadCategories(models.FileCategories)
			if err != nil {
				console.ShowError(err)
				return
			}
			selectedSuppliers, err := Supplier.ReadSuppliers(models.FileSupplier)
			if err != nil {
				console.ShowError(err)
				return
			}

			// Kategorie auswählen
			console.ShowMessage(fmt.Sprintf("Aktuelle Kategorie: %s", item.Category))
			newCategory = selectItemWithCancel(selectedCategories, "Kategorie")
			if newCategory == "" {
				return
			}

			console.ShowMessage(fmt.Sprintf("Aktuelle Artikelnummer: %s", item.ArticleNumber))
			newArticleNumber = console.AskForArticleNumber(item.ArticleNumber, isEditing)

			// Lieferant auswählen
			console.ShowMessage(fmt.Sprintf("Aktueller Lieferant: %s", item.Supplier))
			newSupplier = selectItemWithCancel(selectedSuppliers, "Lieferant")
			if newSupplier == "" {
				return
			}

			newQuantity := item.Quantity

			console.ShowMessage(fmt.Sprintf("Aktuelle Notizen: %s", item.Note))
			newNotes = console.AskForNotes(item.Note, isEditing)

			// Bestätigung zum Bearbeiten des Artikels
			if handleConfirmItemDetails(NewArticleName, newCategory, newArticleNumber, newSupplier, newQuantity, newNotes) {
				// Artikel aktualisieren
				data := models.Item{
					ArticleName:   NewArticleName,
					Category:      newCategory,
					ArticleNumber: newArticleNumber,
					Supplier:      newSupplier,
					Quantity:      newQuantity,
					Note:          newNotes,
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
			} else {
				isEditing = true
			}
		}
	}
}

func selectItemWithCancel(items []string, itemType string) string {
	selectedItem := console.SelectItem(items, pageSize, itemType)
	if selectedItem == "" {
		console.InputC()
	}
	return selectedItem
}

// confirmItemDetails is a method that is used to obtain confirmation from the user for the specified item details
func handleConfirmItemDetails(articleName, category, articleNumber, supplier string, quantity int, notes string) bool {
	console.Clear()
	console.ShowMessage("Bitte überprüfen Sie die neuen Daten:")
	console.ShowMessage(fmt.Sprintf("Artikelbezeichnung: %s", articleName))
	console.ShowMessage(fmt.Sprintf("Kategorie: %s", category))
	console.ShowMessage(fmt.Sprintf("Artikelnummer: %s", articleNumber))
	console.ShowMessage(fmt.Sprintf("Lieferant: %s", supplier))
	console.ShowMessage(fmt.Sprintf("Menge: %d", quantity))
	console.ShowMessage(fmt.Sprintf("Notizen: %s", notes))
	console.ShowMessage("\nSind die Daten korrekt? (y/n) oder [c] um zum Hauptmenü zurückzukehren.")

	choice := console.AskForInput()
	switch strings.ToLower(choice) {
	case "y":
		return true
	case "n":
		console.ShowMessage("✏️ Bitte korrigieren Sie die Daten.")
		return false
	case "c":
		console.InputC()
		return false
	default:
		console.ShowMessage("Ungültige Eingabe, bitte versuchen Sie es erneut.")
		return handleConfirmItemDetails(articleName, category, articleNumber, supplier, quantity, notes)
	}
}

// Case 09
// handleViewItems Displays all items in the inventory.
func handleViewItems() {
	console.Clear()
	items := models.GetAllItems()

	if console.ChecksInventory() {
		return
	}

	page := InitialPage

	for {
		start, end := console.PageIndexCalculate(page, pageSize, len(items))

		console.ShowAllItems(items[start:end], start)

		choice := console.PageIndexView()

		if choice == "c" {
			console.InputC()
			return
		} else if choice == "" {
			page++
			if end == len(items) {
				console.InputPageEnd()
				return
			}
		}
	}
}

// ShowSuppliers Case 1
// ShowSuppliers retrieves and displays the list of suppliers
func ShowSuppliers(filePath string) {
	suppliers, err := Supplier.ReadSuppliers(filePath)
	if err != nil {
		fmt.Printf("Error reading suppliers: %v\n", err)
		return
	}

	for {
		// Display the list of suppliers
		if len(suppliers) > 0 {
			fmt.Println("Available suppliers:")
			for i, supplier := range suppliers {
				fmt.Printf("%d. %s\n", i+1, supplier)
			}
		} else {
			fmt.Println("List is empty. No supplier available.")
		}

		// Provide user options
		console.ShowOption()

		break
	}
}

// AddSupplier Case 2
// AddSupplier add supplier to existing list
func AddSupplier(filePath string) {
	for { // Loop to allow multiple supplier additions
		// Display the list of existing suppliers
		suppliers, err := Supplier.ReadSuppliers(filePath)
		if err != nil {
			fmt.Printf("Error reading suppliers: %v\n", err)
			return
		}

		if len(suppliers) > 0 {
			fmt.Println("* Available suppliers:")
			for i, supplier := range suppliers {
				fmt.Printf("%d. %s\n", i+1, supplier) // Display each supplier with its index
			}
		} else {
			fmt.Println("List is empty. No supplier available.") // Message if the list is empty
		}

		// Show options to the user
		fmt.Println("Enter the name of the supplier you want to add (or 'C' to cancel):")

		// Read the supplier name input
		var supplierName string
		_, err = fmt.Scanln(&supplierName)
		if err != nil {
			return
		}

		// Check if the user wants to cancel the process
		if supplierName == "C" || supplierName == "c" {
			fmt.Println("Action canceled. Returning to the service menu ...")
			break
		}

		// Add the new supplier to the file
		err = Supplier.AddSupplierToFile(filePath, supplierName)
		if err != nil {
			fmt.Printf("Error adding supplier: %v\n", err)
		} else {
			fmt.Println("Supplier added successfully") // Confirmation message
		}

		// Reload the supplier list after adding a new one
		suppliers, err = Supplier.ReadSuppliers(filePath)
		if err != nil {
			fmt.Printf("Error reloading suppliers: %v\n", err)
			return
		}
	}
}

// DeleteSupplier Case 3
// DeleteSupplier from existing list
func DeleteSupplier(filePath string) {
	for {
		// Read the current supplier list
		suppliers, err := Supplier.ReadSuppliers(filePath)
		if err != nil {
			fmt.Printf("Error reading suppliers: %v\n", err)
			return
		}

		// Check if the list is empty
		if len(suppliers) == 0 {
			fmt.Println("No suppliers available to delete.")
			return
		}

		// Display the list of suppliers
		fmt.Println("* Available suppliers:")
		for i, supplier := range suppliers {
			fmt.Printf("%d. %s\n", i+1, supplier)
		}

		// Ask the user to select a supplier to delete or cancel
		fmt.Println("\nEnter the number of the supplier you want to delete (or 'C' to cancel):")
		var input string
		_, err = fmt.Scanln(&input)
		if err != nil {
			return
		}

		// Check if the user canceled
		if input == "C" || input == "c" {
			fmt.Println("Action canceled. Returning to the service menu...")
			return
		}

		// Convert input to integer and validate
		index, err := strconv.Atoi(input)
		if err != nil || index < 1 || index > len(suppliers) {
			fmt.Println("Invalid input. Please enter a valid supplier number.")
			continue
		}

		// Remove the selected supplier from the list
		supplierToDelete := suppliers[index-1]
		suppliers = append(suppliers[:index-1], suppliers[index:]...)

		// Overwrite the supplier CSV file with the updated list
		err = Supplier.OverwriteSupplierFile(filePath, suppliers)
		if err != nil {
			fmt.Printf("Error updating supplier file: %v\n", err)
			return
		}

		// Confirm the deletion
		fmt.Printf("Supplier '%s' deleted successfully.\n\n", supplierToDelete)
	}
}

// ShowCategories Case 11
// ShowCategories retrieves and displays the list of categories
func ShowCategories(filePath string) {
	categories, err := Category.ReadCategories(filePath)
	if err != nil {
		fmt.Printf("Error reading categoriess: %v\n", err)
		return
	}

	for {
		// Display the list of category
		if len(categories) > 0 {
			fmt.Println("Available categories:")
			for i, category := range categories {
				fmt.Printf("%d. %s\n", i+1, category)
			}
		} else {
			fmt.Println("List is empty. No categories available.")
		}

		// Provide user options
		console.ShowOption()

		break
	}
}

// AddCategory Case 2
// AddCategory add category to existing list
func AddCategory(filePath string) {
	for { // Loop to allow multiple category additions
		// Display the list of existing category
		categories, err := Category.ReadCategories(filePath)
		if err != nil {
			fmt.Printf("Error reading categories: %v\n", err)
			return
		}

		if len(categories) > 0 {
			fmt.Println("* Available categories:")
			for i, category := range categories {
				fmt.Printf("%d. %s\n", i+1, category) // Display each supplier with its index
			}
		} else {
			fmt.Println("List is empty. No category available.") // Message if the list is empty
		}

		// Show options to the user
		fmt.Println("Enter the name of the category you want to add (or 'C' to cancel):")

		// Read the category name input
		var categoryName string
		_, err = fmt.Scanln(&categoryName)
		if err != nil {
			return
		}

		// Check if the user wants to cancel the process
		if categoryName == "C" || categoryName == "c" {
			fmt.Println("Action canceled. Returning to the service menu ...")
			break
		}

		// Add the new category to the file
		err = Category.AddCategoryToFile(filePath, categoryName)
		if err != nil {
			fmt.Printf("Error adding category: %v\n", err)
		} else {
			fmt.Println("Category added successfully") // Confirmation message
		}

		// Reload the category list after adding a new one
		categories, err = Category.ReadCategories(filePath)
		if err != nil {
			fmt.Printf("Error reloading categories: %v\n", err)
			return
		}
	}
}

// DeleteCategory Case 3
// DeleteCategory from existing list
func DeleteCategory(filePath string) {
	for {
		// Read the current supplier list
		categories, err := Category.ReadCategories(filePath)
		if err != nil {
			fmt.Printf("Error reading category: %v\n", err)
			return
		}

		// Check if the list is empty
		if len(categories) == 0 {
			fmt.Println("No categories available to delete.")
			return
		}

		// Display the list of categories
		fmt.Println("* Available categories:")
		for i, category := range categories {
			fmt.Printf("%d. %s\n", i+1, category)
		}

		// Ask the user to select a category to delete or cancel
		fmt.Println("\nEnter the number of the category you want to delete (or 'C' to cancel):")
		var input string
		_, err = fmt.Scanln(&input)
		if err != nil {
			return
		}

		// Check if the user canceled
		if input == "C" || input == "c" {
			fmt.Println("Action canceled. Returning to the service menu...")
			return
		}

		// Convert input to integer and validate
		index, err := strconv.Atoi(input)
		if err != nil || index < 1 || index > len(categories) {
			fmt.Println("Invalid input. Please enter a valid category number.")
			continue
		}

		// Remove the selected category from the list
		supplierToDelete := categories[index-1]
		categories = append(categories[:index-1], categories[index:]...)

		// Overwrite the supplier CSV file with the updated list
		err = Category.OverwriteCategoryFile(filePath, categories)
		if err != nil {
			fmt.Printf("Error updating categories file: %v\n", err)
			return
		}

		// Confirm the deletion
		fmt.Printf("Category '%s' deleted successfully.\n\n", supplierToDelete)
	}
}
