package tests

import ("testing"
	".."
	"encoding/json"
	"fmt"

)

func Test_ValueItem(t *testing.T)  {


	item :=go_collection_json.ValueItem{

		Item:go_collection_json.Item{
			Name:"Hi",
		},
		Value:123,


	}



	bytes,_:=json.Marshal(item)
	if (fmt.Sprintf("{\"name\":\"%v\",\"value\":%v}",item.Name,item.Value)!=string(bytes)){
		t.Error("should be %v.",string(bytes))
	}

	//fmt.Println(string(bytes))

}
func Test_Link(t *testing.T)  {


	link :=go_collection_json.Link{
		Href:"http://localhost:80/users/1",
		Rel:"hi",



	}
	bytes,_:=json.Marshal(link)
	if (fmt.Sprintf("{\"href\":\"%v\",\"rel\":\"%v\"}", link.Href, link.Rel)!=string(bytes)){
		t.Error("should be %v.",string(bytes))
	}

	fmt.Println(string(bytes))

}
func Test_Template(t *testing.T)  {


	valueItem:=go_collection_json.ValueItem{
		Item:go_collection_json.Item{
			Name:"name",

		},
		Value:"mark",
	}
	//arrayItem:=go_collection_json.ArrayItem{
	//	Item:go_collection_json.Item{
	//		Name:"friends",
	//
	//	},
	//	Array:[]string{"a1","a2"},
	//}
	var item go_collection_json.Item
	item=valueItem
	items:=[]go_collection_json.Item{valueItem}
	template:=go_collection_json.Template{
		Data:items,
	}

	//
	bytes,_:=json.Marshal(template)

	//if (fmt.Sprintf("{\"href\":\"%v\",\"rel\":\"%v\"}", link.Href, link.Rel)!=string(bytes)){
	//	t.Error("should be %v.",string(bytes))
	//}
	//
	fmt.Println(string(bytes))

}

