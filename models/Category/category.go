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
	// Allow variable number of fields per record
	reader.FieldsPerRecord = -1
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
func AddCategoryToFile(filePath, categoryName string) error {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
		}
	}(file)

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if err := writer.Write([]string{categoryName}); err != nil {
		return err
	}

	return nil
}

// IsValidCategoryName checks if the category name meets the validation criteria
func IsValidCategoryName(name string) bool {
	if name == "" {
		return false
	}
	for _, char := range name {
		if !(char >= 'a' && char <= 'z') && !(char >= 'A' && char <= 'Z') && !(char >= '0' && char <= '9') && char != ' ' {
			return false
		}
	}
	return true
}

// DeleteCategory from existing list
func DeleteCategory(filePath string, index int) error {
	categories, err := ReadCategories(filePath)
	if err != nil {
		return fmt.Errorf("error reading categories: %v", err)
	}

	// Remove the selected supplier
	categories = append(categories[:index], categories[index+1:]...)

	// Overwrite the supplier file
	return OverwriteCategoryFile(filePath, categories)
}

// OverwriteCategoryFile overwrites the content of the given file with the provided list of categories.
func OverwriteCategoryFile(filePath string, categories []string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	// Initialize a CSV writer to write to the file
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write each category as a new row in the CSV file
	for _, categories := range categories {
		err = writer.Write([]string{categories})
		if err != nil {
			return err
		}
	}
	return nil
}
