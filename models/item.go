// Einfache Hilfsfunktionen

package models

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// Item as type
type Item struct {
	ArticleName   string
	Category      string
	ArticleNumber string
	Supplier      string
	Manufacturer  string
	Quantity      int
	Note          string
}

type Supplier struct {
	SupplierName string
}
type Category struct {
	CategoryName string
}

func StringToInt(info string) int {
	value, _ := strconv.Atoi(strings.TrimSpace(info))
	return value
}

func IntToString(value int) string {
	return strconv.Itoa(value)
}

// file_handler
//
// getDataFromDataFile enthält Funktionen zum Lesen und Schreiben in die CSV-Datei.
func getDataFromDataFile() ([]Item, error) {
	file, err := os.Open(FileData)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	csvReader.FieldsPerRecord = reflect.TypeOf(Item{}).NumField()
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	var readItems []Item
	for _, record := range records {
		// Nutzung der angepassten Funktion zum Parsen der Zeile
		parsedItem, err := ParseItemFromCsvStringList(record)
		if err != nil {
			return nil, err
		}
		readItems = append(readItems, parsedItem) // Add parsedItem to slice
	}

	return readItems, nil
}

func updateDataInFile() error {
	file, err := os.Create(FileData)
	if err != nil {
		return err
	}
	writer := csv.NewWriter(file)

	for _, item := range items {
		itemRecord := getItemAsStringSlice(item)
		err := writer.Write(itemRecord)
		if err != nil {
			_ = file.Close()
			return err
		}
	}

	writer.Flush()
	err = file.Close()
	if err != nil {
		return err
	}

	return nil
}

// ParseItemFromCsvStringList verarbeitet eine CSV-Zeile und erstellt ein Item
func ParseItemFromCsvStringList(record []string) (Item, error) {
	parsedItem := Item{}
	numberOfReceivedElements := len(record)
	numberOfFieldsInBookStruct := reflect.TypeOf(Item{}).NumField()

	if numberOfReceivedElements != numberOfFieldsInBookStruct {
		err := fmt.Errorf("data record does not contain enough elements for parsing: received %d, expected: %d",
			numberOfReceivedElements, numberOfFieldsInBookStruct)
		return parsedItem, err
	}

	// Create new book based on parsed values
	parsedItem = Item{
		ArticleName:   strings.TrimSpace(record[0]),
		Category:      strings.TrimSpace(record[1]),
		ArticleNumber: strings.TrimSpace(record[2]),
		Supplier:      strings.TrimSpace(record[3]),
		Manufacturer:  strings.TrimSpace(record[4]),
		Quantity:      StringToInt(record[5]), // Menge als int
		Note:          strings.TrimSpace(record[6]),
	}

	return parsedItem, nil
}

func getItemAsStringSlice(item Item) []string {
	bookSerialized := []string{
		item.ArticleName,
		item.Category,
		item.ArticleNumber,
		item.Supplier,
		item.Manufacturer,
		IntToString(item.Quantity),
		item.Note,
	}

	return bookSerialized
}

// UpdateItem aktualisiert einen Artikel im Inventar
func UpdateItem(id int, updatedItem Item) error {
	if id < 0 || id >= len(items) {
		return errors.New("ungültige ID")
	}

	items[id] = updatedItem
	return updateDataInFile()
}

// getDataFromSupplierFile liest die Lieferantendaten aus der CSV-Datei
func getDataFromSupplierFile() ([]Supplier, error) {
	file, err := os.Open(FileSupplier)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	var readSuppliers []Supplier
	for _, record := range records {
		readSuppliers = append(readSuppliers, Supplier{
			SupplierName: record[0],
		})
	}

	return readSuppliers, nil
}

//
//
// file_handler //
//
//

//
//
// item_repository
//
//
// Verwaltung der Items (CRUD-Operationen) und die Dateioperationen.

const FileData = "data.csv"
const FileCategories = "categories.csv"
const FileSupplier = "supplier.csv"

var items []Item
var suppliers []Supplier
var categories []Category

// AddItem adds the passed Item to the Inventory
func AddItem(newItem Item) error {
	// Add to item Repository
	items = append(items, newItem)

	// Update data in file
	return updateDataInFile()
}

// Initialize does the initialization of the repository
func Initialize() error {
	var err error
	// Initialisieren Artikel
	items, err = getDataFromDataFile()
	if err != nil {
		return err
	}
	// Initialisieren Lieferanten
	suppliers, err = getDataFromSupplierFile()
	if err != nil {
		return err
	}
	return nil
}

// GetAllItems returns a copy of all items
func GetAllItems() []Item {
	allItems := make([]Item, len(items))
	copy(allItems, items)
	return allItems
}

func GetAllCategories() []Category {
	allCategories := make([]Category, len(categories))
	copy(allCategories, categories)
	return allCategories
}

// GetItemById retrieves an item by its index
func GetItemById(rowId int) *Item {
	if rowId < 1 || rowId >= len(items) {
		return nil // Wenn der Index ungültig ist, gebe nil zurück
	}
	return &items[rowId]
}

// RemoveItem removes the passed row ID from the library
func RemoveItem(rowId int) error {
	// input validation
	if rowId < 1 {
		fmt.Printf("row Id %d in wrong data range, value must >= 1", rowId)
	}
	// Temporary slice variable
	var tempItems []Item
	// Normalize rot ID
	rowIdNormed := rowId - 1
	//loop through existing slice and add all item except the removing one
	for index, value := range items {
		if index != rowIdNormed {
			tempItems = append(tempItems, value)
		}
	}
	// Assign temporary slice to existing package slice variable
	items = tempItems

	// Update data in file
	err := updateDataInFile()
	return err
}

//
//
// item_repository //
//
//
