package datastruct

type Link struct {
	Original string
	Shortened string
}

func NewLink(original, shortened string) Link {
	return Link{
		Original:  original,
		Shortened: shortened,
	}
}
