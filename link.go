package go_collection_json



//{"href" : URI, "rel" : STRING, "prompt" : STRING, "name" : STRING, "render" : "image"},
//Each has five possible properties:
// href (REQUIRED),
// rel (REQURIED),
// name (OPTIONAL),
// render (OPTIONAL),
// and prompt, (OPTIONAL).
type Link struct {
	Href string `json:"href"`
	Rel string `json:"rel"`
	Name string `json:"name,omitempty"`
	Render string `json:"render,omitempty"`
	Prompt string `json:"prompt,omitempty"`


}
