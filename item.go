package go_collection_json

type Item struct {
	Prompt string `json:"prompt,omitempty"`
	Name string `json:"name"`
}
type ValueItem struct {
	Item
	//Prompt string `json:"prompt,omitempty"`
	//Name string `json:"name"`
	Value interface{} `json:"value,omitempty"`
}
type ArrayItem struct {
	Item
	Array []interface{} `json:"array,omitempty"`
}
type ObjectItem struct {
	Item
	Object interface{} `json:"object,omitempty"`
}

