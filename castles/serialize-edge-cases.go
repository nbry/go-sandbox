package castles

import (
	"encoding/json"
	"fmt"
)

type Fruit struct {
	Name     string
	Price    float64
	Quantity int32
}

func (f Fruit) MarshalJSON() ([]byte, error) {
	a := make([]interface{}, 3)

	a[0] = f.Name
	a[1] = f.Price
	a[2] = f.Quantity

	return json.Marshal(a)
}

func (f *Fruit) UnmarshalJSON(bs []byte) error {
	arr := []interface{}{}
	json.Unmarshal(bs, &arr)

	f.Name = arr[0].(string)
	f.Price = arr[1].(float64)
	f.Quantity = int32(arr[2].(float64))

	return nil
}

func ArrayWithDifferentTypes() {
	data := []byte(`["Apple", 0.99, 4]`)

	apple := Fruit{}
	json.Unmarshal(data, &apple)
	fmt.Println("Unmarshalled, to Fruit struct: ")
	fmt.Println(apple)

	bytes, _ := json.Marshal(apple)
	fmt.Println("\nMarshalled back to stringified list: ")
	fmt.Println(string(bytes))
}
