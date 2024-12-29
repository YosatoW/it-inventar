package controllers

/*
func handleAddItem() {
	console.Clear()
	console.ShowAddItemInformation()

	var isEditing bool = false
	var articleName, articleNumber, notes string
	var quantity int
	var chosenCategories []string
	var chosenSuppliers []string

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
		// Die Eingabewerte werden jetzt nur einmal initialisiert und bei Korrekturen wiederverwendet
		articleName = console.AskForName(articleName, isEditing)

		chosenCategory := console.SelectCategory(selectedCategories, pageSize)
		if chosenCategory == "" {
			return
		}
		chosenCategories = []string{chosenCategory}

		articleNumber = console.AskForArticleNumber(articleNumber, isEditing)

		chosenSupplier := console.SelectSupplier(selectedSuppliers, pageSize)
		if chosenSupplier == "" {
			return
		}
		chosenSuppliers = []string{chosenSupplier}

		quantity = console.AskForQuantity(quantity, isEditing)
		notes = console.AskForNotes(notes, isEditing)

		for {
			// Benutzer überprüft die Eingaben
			console.Clear()
			console.ShowMessage("Bitte überprüfen Sie die eingegebenen Daten:")
			console.ShowMessage(fmt.Sprintf("Artikel-Bez.: %s", articleName))
			console.ShowMessage(fmt.Sprintf("Kategorie: %s", chosenCategories[0]))
			console.ShowMessage(fmt.Sprintf("Artikel-Nr.: %s", articleNumber))
			console.ShowMessage(fmt.Sprintf("Lieferant: %s", chosenSuppliers[0]))
			console.ShowMessage(fmt.Sprintf("Menge: %d", quantity))
			console.ShowMessage(fmt.Sprintf("Notizen: %s", notes))
			console.ShowMessage("\nSind die Daten korrekt? (y/n) oder [c] um zum Hauptmenü zurückzukehren.")

			choice := console.AskForInput()
			if strings.ToLower(choice) == "y" {
				// Artikel zusammenstellen
				data := models.Item{
					ArticleName:   articleName,
					Category:      chosenCategories[0],
					ArticleNumber: articleNumber,
					Supplier:      chosenSuppliers[0],
					Quantity:      quantity,
					Note:          notes,
				}
				// Artikel hinzufügen
				err := models.AddItem(data)
				if err != nil {
					console.ShowError(err)
				} else {
					console.ShowMessage("✅ Artikel erfolgreich hinzugefügt!")
					console.ShowContinue()
					console.InputC()
					return
				}
			} else if strings.ToLower(choice) == "n" {
				console.ShowMessage("✏️ Bitte korrigieren Sie die Daten.")
				isEditing = true // Korrekturmodus aktivieren
				break
			} else if strings.ToLower(choice) == "c" {
				// Abbrechen und zurück zum Menü
				console.InputC()
				return
			} else {
				// Ungültige Eingabe, erneut fragen
				console.ShowMessage("Ungültige Eingabe, bitte versuchen Sie es erneut.")
			}
		}
	}
}*/
