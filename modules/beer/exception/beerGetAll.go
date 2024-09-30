package exception

type BeerGetAll struct{}

func (e *BeerGetAll) Error() string {
	return "Get All Item Fail !!!"
}
