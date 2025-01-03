package controllers

import (
	"fmt"
	"it_inventar/models"
	"it_inventar/models/Category"
	"it_inventar/models/Supplier"
	"it_inventar/views/console"
	"strconv"
	"strings"
)

const (
	InitialPage = console.InitialPage
	pageSize    = console.PageSize

	messageInvalidInput         = "Invalid input!"
	messageInvalidInputTryAgain = "Please choose:\n[y] Confirm\n[n] Cancel."
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
		chosenCategory = console.HandleAddSelectItem(chosenCategory, selectedCategories, "Category", isEditing)
		if chosenCategory == "" {
			return
		}
		articleNumber = console.AskForArticleNumber(articleNumber, isEditing)
		chosenSupplier = console.HandleAddSelectItem(chosenSupplier, selectedSuppliers, "Supplier", isEditing)
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
				console.ShowMessage("✅ Item successfully added!")
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

		choice := console.PageIndexPrompt("Item")

		exit, item, rowId := console.PageIndexUserInput(choice, &page, end, items)
		if exit {
			return
		}
		if item != nil {
			console.ShowMessage(fmt.Sprintf("%s\nDelete this item? (y/n)", console.ConfirmTheArticle(*item)))

			for {
				choice = console.AskForInput()

				if strings.ToLower(choice) == "y" {
					err := models.RemoveItem(rowId)
					if err != nil {
						console.ShowError(err)
					} else {
						console.ShowMessage("✅ Item successfully removed!")
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
					console.ShowMessage(fmt.Sprintf("Item: %s (%s) - %d pieces - Notes: %s", item.ArticleName, item.ArticleNumber, item.Quantity, item.Note))
					console.ShowMessage(messageInvalidInputTryAgain)
				}
			}
		}
	}
}

// Case 03
// handleChangeQuantity edits an item in the inventory
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

		choice := console.PageIndexPrompt("Item")

		exit, item, rowId := console.PageIndexUserInput(choice, &page, end, items)
		if exit {
			return
		}
		if item != nil {
			console.ShowMessage(fmt.Sprintf("%s\nAdjust the quantity of this item? (y/n)", console.ConfirmTheArticle(*item)))

			for {
				choice = console.AskForInput()
				if strings.ToLower(choice) == "y" {
					// Ask for adding or subtracting
					console.ShowMessage(fmt.Sprintf("[1] Add\n[2] Subtract"))
					operation := console.AskForInput()

					if strings.ToLower(operation) == "1" {
						// Ask for the quantity to add
						console.ShowMessage(fmt.Sprintf("Current stock: %d pieces", item.Quantity))
						console.ShowMessage("Enter the quantity to add:")
						quantityToAdd := console.AskForQuantity(0, false)
						item.Quantity += quantityToAdd
					} else if strings.ToLower(operation) == "2" {
						// Ask for the quantity to subtract
						console.ShowMessage(fmt.Sprintf("Current stock: %d pieces", item.Quantity))
						console.ShowMessage("Enter the quantity to subtract:")
						quantityToSubtract := console.AskForQuantity(0, false)
						if item.Quantity < quantityToSubtract {
							console.ShowMessage("❌ The quantity to subtract exceeds the available quantity.")
							console.ShowContinue()
							continue
						}
						item.Quantity -= quantityToSubtract
					} else {
						console.ShowMessage("❌ Invalid selection. Please choose '1' or '2'.")
						console.ShowContinue()
						continue
					}

					// Update item quantity
					err := models.UpdateItem(rowId-1, *item)
					if err != nil {
						console.ShowError(err)
					} else {
						console.ShowMessage(fmt.Sprintf("New stock: %d pieces", item.Quantity))
						console.ShowMessage("✅ Item quantity successfully updated!")
						console.ShowContinue()
						console.Clear()
						console.ShowExecuteCommandMenu()
						return
					}
				} else if strings.ToLower(choice) == "n" {
					console.HandleChancelAction()
					break

				} else {
					// Invalid input, ask again
					console.Clear()
					console.ShowMessage(messageInvalidInput)
					console.ShowMessage(fmt.Sprintf("Item: %s (%s) - %d pieces - Notes: %s", item.ArticleName, item.ArticleNumber, item.Quantity, item.Note))
					console.ShowMessage(messageInvalidInputTryAgain)
				}
			}
		}
	}
}

// Case 04
// handleChanceArticleInformation edits an item in the inventory
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

		choice := console.PageIndexPrompt("Item")

		exit, item, rowId := console.PageIndexUserInput(choice, &page, end, items)
		if exit {
			return
		}
		if item != nil {

			// The input values are now initialized only once and reused for corrections
			console.ShowMessage(fmt.Sprintf("Current item name: %s", item.ArticleName))
			NewArticleName = console.AskForName(item.ArticleName, isEditing)

			// Load categories and suppliers
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

			// Select category
			console.ShowMessage(fmt.Sprintf("Current category: %s", item.Category))
			newCategory = selectItemWithCancel(selectedCategories, "Category")
			if newCategory == "" {
				return
			}

			console.ShowMessage(fmt.Sprintf("Current article number: %s", item.ArticleNumber))
			newArticleNumber = console.AskForArticleNumber(item.ArticleNumber, isEditing)

			// Select supplier
			console.ShowMessage(fmt.Sprintf("Current supplier: %s", item.Supplier))
			newSupplier = selectItemWithCancel(selectedSuppliers, "Supplier")
			if newSupplier == "" {
				return
			}

			newQuantity := item.Quantity

			console.ShowMessage(fmt.Sprintf("Current notes: %s", item.Note))
			newNotes = console.AskForNotes(item.Note, isEditing)

			// Confirmation to edit the item
			if handleConfirmItemDetails(NewArticleName, newCategory, newArticleNumber, newSupplier, newQuantity, newNotes) {
				// Update item
				data := models.Item{
					ArticleName:   NewArticleName,
					Category:      newCategory,
					ArticleNumber: newArticleNumber,
					Supplier:      newSupplier,
					Quantity:      newQuantity,
					Note:          newNotes,
				}
				// Adjust the index correctly here
				err := models.UpdateItem(rowId-1, data)
				if err != nil {
					console.ShowError(err)
				} else {
					console.ShowMessage("✅ Item successfully updated!")
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

// handleConfirmItemDetails is a method that is used to obtain confirmation from the user for the specified item details
func handleConfirmItemDetails(articleName, category, articleNumber, supplier string, quantity int, notes string) bool {
	console.Clear()
	console.ShowMessage("Please review the new data:")
	console.ShowMessage(fmt.Sprintf("Item name: %s", articleName))
	console.ShowMessage(fmt.Sprintf("Category: %s", category))
	console.ShowMessage(fmt.Sprintf("Article number: %s", articleNumber))
	console.ShowMessage(fmt.Sprintf("Supplier: %s", supplier))
	console.ShowMessage(fmt.Sprintf("Quantity: %d", quantity))
	console.ShowMessage(fmt.Sprintf("Notes: %s", notes))
	console.ShowMessage("\nAre the details correct? (y/n) or [c] to return to the main menu.")

	choice := console.AskForInput()
	switch strings.ToLower(choice) {
	case "y":
		return true
	case "n":
		console.ShowMessage("✏️ Please correct the data.")
		return false
	case "c":
		console.InputC()
		return false
	default:
		console.ShowMessage("Invalid input, please try again.")
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

// handleShowSuppliers displays a list of suppliers and allows navigation or exiting
func handleShowSuppliers() {
	// Read the list of suppliers from the CSV file
	suppliers, err := Supplier.ReadSuppliers("supplier.csv")
	if err != nil {
		console.ErrorMessage(fmt.Sprintf("Error reading suppliers: %v\n", err))
		return
	}
	// Check if the list is empty
	if len(suppliers) == 0 {
		console.ShowNoSuppliersMessage()
		return
	}

	page := InitialPage
	for {
		// Calculate the start and end indices for the current page
		start := page * pageSize
		end := start + pageSize
		if end > len(suppliers) {
			end = len(suppliers)
		}

		// Display the current page of suppliers
		console.DisplaySuppliers(suppliers, start, end)

		// Check if the end of the list has been reached
		if end >= len(suppliers) {
			console.ShowEndOfSupplier()

			for {
				// Ask the user to continue or exit
				userChoice := console.ShowOnlyCancelMessage()
				if userChoice == "c" {
					break
				}
			}
			break
		}

		// For normal page navigation
		userChoice := console.GetPageInput()
		if userChoice == "c" {
			break
		} else if userChoice == "" {
			page++
		}
	}
}

// handleAddSuppliers allows the user to add new suppliers to the list with an option to cancel
func handleAddSuppliers() {
	filePath := "supplier.csv"

	for {
		// Display the list of existing suppliers
		suppliers, err := Supplier.ReadSuppliers(filePath)
		if err != nil {
			console.ErrorMessage(fmt.Sprintf("Error reading suppliers: %v", err))
			return
		}
		console.ShowSuppliersList(suppliers)

		// Prompt for new supplier
		console.ShowPrompt("Enter the name of the supplier you want to add (or 'C' to cancel):")
		supplierName := console.GetUserInput()

		if supplierName == "C" || supplierName == "c" {
			console.ShowMessage("Action canceled. Returning to the service menu...")
			return
		}

		// Add the new supplier to the file
		err = Supplier.AddSupplierToFile(filePath, supplierName)
		if err != nil {
			console.ErrorMessage(fmt.Sprintf("Error adding supplier: %v", err))
		} else {
			console.ShowMessage("Supplier added successfully")
		}
	}
}

// handleDeleteSupplier enables the deletion of suppliers from the list with input validation and cancellation
func handleDeleteSupplier() {
	filePath := "supplier.csv"

	for {
		// Read the list of suppliers from the CSV file
		suppliers, err := Supplier.ReadSuppliers(filePath)
		if err != nil {
			console.ErrorMessage(fmt.Sprintf("Error reading suppliers: %v", err))
			return
		}

		// Check if the supplier list is empty
		if len(suppliers) == 0 {
			console.ShowMessage("No suppliers available to delete.")
			return
		}

		// Display the list of existing suppliers
		console.ShowSuppliersList(suppliers)

		// Prompt user to select a supplier to delete
		console.ShowPrompt("Enter the number of the supplier you want to delete (or 'C' to cancel):")
		input := console.GetUserInput()

		// Check if the user wants to cancel
		if input == "C" || input == "c" {
			console.ShowMessage("Action canceled. Returning to the service menu...")
			return
		}

		// Convert input to integer and validate
		index, err := strconv.Atoi(input)
		if err != nil || index < 1 || index > len(suppliers) {
			console.ErrorMessage("Invalid input. Please enter a valid supplier number.")
			continue
		}

		// Perform deletion
		supplierToDelete := suppliers[index-1]
		err = Supplier.DeleteSupplier(filePath, index-1)
		if err != nil {
			console.ErrorMessage(fmt.Sprintf("Error deleting supplier: %v", err))
			return
		}

		// Confirm successful deletion
		console.ShowMessage(fmt.Sprintf("Supplier '%s' deleted successfully.", supplierToDelete))
	}
}

// handleShowCategories displays a list of categories and allows navigation or exiting
func handleShowCategories() {
	// Calculate the start and end indices for the current page
	categories, err := Category.ReadCategories("categories.csv")
	if err != nil {
		console.ErrorMessage(fmt.Sprintf("Error reading categories: %v\n", err))
		return
	}
	// Display the current page of categories
	if len(categories) == 0 {
		console.ShowNoCategoriesMessage()
		return
	}

	page := InitialPage
	for {
		// Calculate the start and end indices for the current page
		start := page * pageSize
		end := start + pageSize
		if end > len(categories) {
			end = len(categories)
		}

		console.DisplayCategories(categories, start, end)

		// Check if the end of the list has been reached
		if end >= len(categories) {
			console.ShowEndOfCategory()

			for {
				// Ask the user to continue or exit
				userChoice := console.ShowOnlyCancelMessage()
				if userChoice == "c" {
					break
				}
			}
			break
		}

		// For normal page navigation
		userChoice := console.GetPageInput()
		if userChoice == "c" {
			break
		} else if userChoice == "" {
			page++
		}
	}
}

// handleAddCategories allows the user to add new categories to the list with an option to cancel
func handleAddCategories() {
	filePath := "categories.csv"

	for {
		// Display the list of existing categories
		categories, err := Category.ReadCategories(filePath)
		if err != nil {
			console.ErrorMessage(fmt.Sprintf("Error reading categories: %v", err))
			return
		}
		console.ShowCategoriesList(categories)

		// Prompt for new supplier
		console.ShowPrompt("Enter the name of the category you want to add (or 'C' to cancel):")
		categoryName := console.GetUserInput()

		if categoryName == "C" || categoryName == "c" {
			console.ShowMessage("Action canceled. Returning to the service menu...")
			return
		}

		// Add the new category to the file
		err = Category.AddCategoryToFile(filePath, categoryName)
		if err != nil {
			console.ErrorMessage(fmt.Sprintf("Error adding category: %v", err))
		} else {
			console.ShowMessage("Category added successfully")
		}
	}
}

// handleDeleteCategories enables the deletion of categories from the list with input validation and cancellation
func handleDeleteCategories() {
	filePath := "categories.csv"

	for {
		// Read the list of categories from the CSV file
		categories, err := Category.ReadCategories(filePath)
		if err != nil {
			console.ErrorMessage(fmt.Sprintf("Error reading categories: %v", err))
			return
		}

		// Check if the category list is empty
		if len(categories) == 0 {
			console.ShowMessage("No categories available to delete.")
			return
		}

		// Display the list of existing categories
		console.ShowCategoriesList(categories)

		// Prompt user to select a category to delete
		console.ShowPrompt("Enter the number of the category you want to delete (or 'C' to cancel):")
		input := console.GetUserInput()

		// Check if the user wants to cancel
		if input == "C" || input == "c" {
			console.ShowMessage("Action canceled. Returning to the service menu...")
			return
		}

		// Convert input to integer and validate
		index, err := strconv.Atoi(input)
		if err != nil || index < 1 || index > len(categories) {
			console.ErrorMessage("Invalid input. Please enter a valid category number.")
			continue
		}

		// Perform deletion
		categoryToDelete := categories[index-1]
		err = Category.DeleteCategory(filePath, index-1)
		if err != nil {
			console.ErrorMessage(fmt.Sprintf("Error deleting category: %v", err))
			return
		}

		// Confirm successful deletion
		console.ShowMessage(fmt.Sprintf("Category '%s' deleted successfully.", categoryToDelete))
	}
}
