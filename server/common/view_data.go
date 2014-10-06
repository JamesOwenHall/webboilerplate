package common

// ViewData is a type that is passed to the template engine.
type ViewData struct {
	Title  string
	Values map[string]interface{}
}

// NewViewData is the preferred way to instantiate a ViewData.
func NewViewData() ViewData {
	return ViewData{
		Values: make(map[string]interface{}),
	}
}
