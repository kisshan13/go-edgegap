package edgegap

import "fmt"

// Utility function to get query paramters for pagination parameters.
func (pp *PaginationParams) GetParams() string {
	return fmt.Sprintf("?page=%d&limit=%d", pp.Page, pp.Size)
}
