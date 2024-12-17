package models

import (
	"encoding/csv"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

const FileName = "data.csv"

// Item as type
type Item struct {
	Name     string
	Model    string
	Quantity int
	Note     string
}

// Liste der Items, die mit der Item-Struktur arbeitet
var items []Item

// Initialize does the initialization of the repository
func Initialize() error {
	var err error
	items, err = getDataFromFile()
	if err != nil {
		return err
	}
	return nil
}

func getDataFromFile() ([]Item, error) {
	file, err := os.Open(FileName)
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
	//
	parsedItem = Item{
		Name:     strings.TrimSpace(record[0]),
		Model:    strings.TrimSpace(record[1]),
		Quantity: StringToInt(record[2]), // Menge als int
		Note:     strings.TrimSpace(record[3]),
	}

	return parsedItem, nil
}

// GetAllItems returns a copy of all items
func GetAllItems() []Item {
	allItems := make([]Item, len(items))
	copy(allItems, items)
	return allItems
}

// AddItem adds the passed Item to the Inventory
func AddItem(newItem Item) error {
	// Add to item Repository
	items = append(items, newItem)

	// Update data in file
	return updateDataInFile()
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

	// Remove book from repository

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

func updateDataInFile() error {
	file, err := os.Create(FileName)
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

func getItemAsStringSlice(item Item) []string {
	bookSerialized := []string{
		item.Name,
		item.Model,
		IntToString(item.Quantity),
		item.Note,
	}

	return bookSerialized
}

func StringToInt(info string) int {
	value, _ := strconv.Atoi(strings.TrimSpace(info))
	return value
}

func IntToString(value int) string {
	return strconv.Itoa(value)
}
