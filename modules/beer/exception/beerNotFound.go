package exception

import "fmt"

type BeerNotFound struct {
	ItemId uint64
}

func (e *BeerNotFound) Error() string {
	return fmt.Sprintf("Beer id : %d not found !!!", e.ItemId)
}
