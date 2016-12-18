package go_collection_json

type Template struct {
	Data []Item `json:"data"`
}

func NewTemplate(data []Item) Template {
	return Template{Data:data}
}
