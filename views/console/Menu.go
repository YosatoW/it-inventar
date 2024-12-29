package console

import "fmt"

// ShowExecuteCommandMenu shows the menu to the console
func ShowExecuteCommandMenu() {
	fmt.Println(`
	###########################################
	#******** WELCOME TO OUR INVENTORY ********
	#******** CHOOSE YOUR OPTION BELOW ********
	# 1. Artikel hinzufügen
	# 2. Artikel löschen
	# 3. Artikel buchung
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
	# 1. Showing Suppliers
	# 2. Showing Existing Suppliers
	# 3. Delete a supplier
	#
	# 11. Showing Categories
	# 12. Showing Existing Categories
	# 13. Delete a category
	#
	#
	# c. SHOW MAIN MENU
	`)
}
