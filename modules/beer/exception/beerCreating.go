package exception

type BeerCreating struct{}

func (e *BeerCreating) Error() string {
	return "Creating Item Fail !!!"
}
