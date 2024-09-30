package exception

type BeerCounting struct{}

func (e *BeerCounting) Error() string {
	return "Counting Item Fail !!!"
}
