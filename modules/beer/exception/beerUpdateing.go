package exception

import "fmt"

type BeerUpdateing struct {
	ItemId uint64
}

func (e *BeerUpdateing) Error() string {
	return fmt.Sprintf("Updateing item id : %d failed !!!", e.ItemId)
}
