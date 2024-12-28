package Category

import (
	"encoding/csv"
	"os"
)

func ReadCategories(filePath string) ([]string, error) {
	// Open the CSV file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err // Return an error if the file cannot be opened
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file) // Ensure the file is closed when the function exits

	// Create a CSV reader and read all rows from the file
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err // Return an error if there is an issue reading the file
	}

	var suppliers []string
	// Iterate over each record (row) in the CSV file
	for _, record := range records {
		if len(record) > 0 {
			suppliers = append(suppliers, record[0]) // Append the supplier name (first column) to the list
		}
	}
	return suppliers, nil // Return the list of suppliers and no error
}

// AddCategoryToFile adds a new supplier to the CSV file
func AddCategoryToFile(filePath string, categoryName string) error {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
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

	return writer.Write([]string{categoryName})
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
