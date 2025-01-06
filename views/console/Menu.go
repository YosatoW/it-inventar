package console

import "fmt"

// ShowExecuteCommandMenu shows the menu to the console
func ShowExecuteCommandMenu() {
	fmt.Println(`
	###########################################
	#******** WELCOME TO OUR INVENTORY ********
	#******** CHOOSE YOUR OPTION BELOW ********
	# -1- Add article
	# -2- Delete article
	# -3- Article booking
	# -4- Change article information
	#
	# -9- Show active articles
	# -9D- Show deleted articles
	# -9A- Show all articles
	#
	# -C- CLEAR VIEW AND SHOW MENU
	# -Q- EXIT INVENTORY APP
	`)
}

// ShowHiddenCommandMenu shows the hidden-menu to the console
func ShowHiddenCommandMenu() {
	fmt.Println(`
	###########################################
	#****************  SERVICE *****************
	#******** CHOOSE YOUR OPTION BELOW *********
	# -1- Show suppliers
	# -2- Add supplier
	# -3- Delete supplier
	#
	# -11- Show categories
	# -12- Add category
	# -13- Delete category
	#


	# -C- SHOW MAIN MENU
	`)
}
