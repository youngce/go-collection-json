package tests

import ("testing"
	".."
	"encoding/json"
	"fmt"

)

func Test_ValueItem(t *testing.T)  {

	var item go_collection_json.ValueItem

	item =go_collection_json.ValueItem{

		Value:123,


	}
	item.Name="name"

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

