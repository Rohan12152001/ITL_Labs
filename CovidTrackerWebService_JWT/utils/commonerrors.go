package utils

import "fmt"

var (
	NoRowsFound   = fmt.Errorf("no rows found")
	NoRowsUpdated = fmt.Errorf("no rows updated")
	RowAlreadyExist = fmt.Errorf("row already exists")
)
