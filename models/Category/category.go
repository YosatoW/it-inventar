package Category

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

// ReadCategories reads all categories from the CSV file and returns them as a slice of strings
func ReadCategories(filePath string) ([]string, error) {
	// Open the CSV file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer func() {
		closeErr := file.Close()
		if closeErr != nil {
			fmt.Printf("Error closing file: %v\n", closeErr)
		}
	}()

	// Create a CSV reader and read all rows from the file
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1 // Allow variable number of fields per record
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var categories []string
	// Iterate over each record (row) in the CSV file
	for _, record := range records {
		if len(record) > 0 {
			categories = append(categories, strings.TrimSpace(record[0]))
		}
	}
	return categories, nil
}

// AddCategoryToFile adds a new category to the CSV file
func AddCategoryToFile(filePath string, categoryName string) error {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
		}
	}(file)

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write supplier name in a new line
	if err := writer.Write([]string{categoryName}); err != nil {
		return err
	}

	return nil
}

// DeleteCategory from existing list
func DeleteCategory(filePath string, index int) error {
	suppliers, err := ReadCategories(filePath)
	if err != nil {
		return fmt.Errorf("error reading suppliers: %v", err)
	}

	// Remove the selected supplier
	suppliers = append(suppliers[:index-1], suppliers[index:]...)

	// Overwrite the supplier file
	return OverwriteCategoryFile(filePath, suppliers)
}

func OverwriteCategoryFile(filePath string, categories []string) error {
	file, err := os.Create(filePath) // Create a new file (overwrite existing)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, categories := range categories {
		err = writer.Write([]string{categories}) // Write each category as a row
		if err != nil {
			return err
		}
	}
	return nil
}
