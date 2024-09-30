package exception

import "fmt"

type BeerDelete struct {
	ItemId uint64
}

func (e *BeerDelete) Error() string {
	return fmt.Sprintf("Deleting item id : %d failed !!!", e.ItemId)
}
