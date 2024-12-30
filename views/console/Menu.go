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
	# 1. Show suppliers
	# 2. Add supplier
	# 3. Delete supplier
	#
	# 11. Show categories
	# 12. Add category
	# 13. Delete category
	#
	#
	# c. SHOW MAIN MENU
	`)
}
