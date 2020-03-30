package categories

import (
	"strings"

	"github.com/sachinkapalidigi/backend-expense-manager/utils/errors"
)

// Category : category of expense
type Category struct {
	ID           int64  `json:"id"`
	CategoryName string `json:"category_name"`
	Description  string `json:"description"`
	CreatedAt    string `json:"created_at"`
}

// Categories : Slice of categories
type Categories []Category

// Validate : validate category
func (c *Category) Validate() *errors.RestErr {
	c.CategoryName = strings.TrimSpace(c.CategoryName)
	if c.CategoryName == "" {
		return errors.NewBadRequestError("Invalid category name")
	}

	return nil
}
