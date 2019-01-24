package mock

type Retriever struct {
	Constents string
}

func (r Retriever) Get(url string) string {
	return r.Constents
}