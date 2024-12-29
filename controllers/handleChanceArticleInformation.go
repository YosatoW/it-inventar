package controllers

//func handleChanceArticleInformation() {
//	console.Clear()
//	items := models.GetAllItems()
//
//	if console.ChecksInventory() {
//		return
//	}
//
//	page := InitialPage
//	for {
//		var isEditing bool = false
//		var newName, newModel, newNotes string
//
//		start, end := console.PageIndexCalculate(page, pageSize, len(items))
//
//		console.ShowAllItems(items[start:end], start)
//
//		choice := console.PageIndexPrompt("Artikel")
//
//		exit, item, rowId := console.PageIndexUserInput(choice, &page, end, items)
//		if exit {
//			return
//		}
//		if item != nil {
//			// Zeige aktuelle Artikelinformationen und frage nach neuen Werten
//			console.ShowMessage(fmt.Sprintf("%s\nDiesen Artikel bearbeiten? (y/n)", console.ConfirmTheArticle(*item)))
//
//			choice = console.AskForInput()
//			if strings.ToLower(choice) == "y" {
//				// Die Eingabewerte werden jetzt nur einmal initialisiert und bei Korrekturen wiederverwendet
//				newName = console.AskForName(item.ArticleName, isEditing)
//				newModel = console.AskForArticleNumber(item.ArticleNumber, isEditing)
//				newNotes = console.AskForNotes(item.Note, isEditing)
//
//				// Bestätigung zum Bearbeiten des Artikels
//				console.ShowMessage("Bitte überprüfen Sie die neuen Daten:")
//				console.ShowMessage(fmt.Sprintf("Artikelbezeichnung: %s", newName))
//				console.ShowMessage(fmt.Sprintf("Artikelnummer: %s", newModel))
//				console.ShowMessage(fmt.Sprintf("Notizen: %s", newNotes))
//				console.ShowMessage("\nSind die Daten korrekt? (y/n) oder [c] um zum Hauptmenü zurückzukehren.")
//
//				for {
//					choice = console.AskForInput()
//
//					if strings.ToLower(choice) == "y" {
//						// Artikel aktualisieren
//						data := models.Item{
//							ArticleName:   newName,
//							ArticleNumber: newModel,
//							Note:          newNotes,
//							Quantity:      item.Quantity,
//						}
//						// Hier wird der Index korrekt angepasst
//						err := models.UpdateItem(rowId-1, data)
//						if err != nil {
//							console.ShowError(err)
//						} else {
//							console.ShowMessage("✅ Artikel erfolgreich aktualisiert!")
//							console.ShowContinue()
//							console.Clear()
//							console.ShowExecuteCommandMenu()
//							return
//						}
//					} else if strings.ToLower(choice) == "c" {
//						console.InputC()
//						return
//					} else if strings.ToLower(choice) == "n" {
//						console.HandleChancelAction()
//						break
//					} else {
//						console.Clear()
//						console.ShowMessage(messageInvalidInput)
//						console.ShowMessage("Für folgende Artikel:")
//						console.ShowMessage(fmt.Sprintf("%s (%s)\nAnzahl: %d Stück\nNotizen: %s", newName, newModel, item.Quantity, newNotes))
//						console.ShowMessage("---------------")
//						console.ShowMessage(messageInvalidInputTryAgain)
//					}
//				}
//			}
//		}
//	}
//}
