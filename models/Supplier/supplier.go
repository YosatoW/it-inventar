package Supplier

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

// ReadSuppliers reads all categories from the CSV file and returns them as a slice of strings
func ReadSuppliers(filePath string) ([]string, error) {
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

	var suppliers []string
	// Iterate over each record (row) in the CSV file
	for _, record := range records {
		if len(record) > 0 {
			suppliers = append(suppliers, strings.TrimSpace(record[0]))
		}
	}
	return suppliers, nil
}

// AddSupplierToFile adds a new category to the CSV file
func AddSupplierToFile(filePath, supplierName string) error {
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

	if err := writer.Write([]string{supplierName}); err != nil {
		return err
	}

	return nil
}

// IsValidSupplierName checks if the supplier name meets the validation criteria
func IsValidSupplierName(name string) bool {
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

// DeleteSupplier from existing list
func DeleteSupplier(filePath string, index int) error {
	suppliers, err := ReadSuppliers(filePath)
	if err != nil {
		return fmt.Errorf("error reading suppliers: %v", err)
	}

	// Remove the selected supplier
	suppliers = append(suppliers[:index], suppliers[index+1:]...)

	// Overwrite the supplier file
	return OverwriteSupplierFile(filePath, suppliers)
}

// OverwriteSupplierFile overwrites the content of the given file with the provided list of categories.
func OverwriteSupplierFile(filePath string, suppliers []string) error {
	file, err := os.Create(filePath) // Create a new file (overwrite existing)
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
	for _, supplier := range suppliers {
		err = writer.Write([]string{supplier}) // Write each supplier as a row
		if err != nil {
			return err
		}
	}
	return nil
}
