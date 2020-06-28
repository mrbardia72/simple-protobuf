package main

import (
	"fmt"
	"log"
	"./src/complex"
	"./src/enum_example"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"./src/simple"
)

func main() {
	sm := doSimple()
	jsonDemo(sm)
	doEnum()
	doComplex()
}

func doComplex() {
	cm := complexpb.ComplexMessage{ 
		OneDummy: &complexpb.DummyMessage{
			Id:   1,
			Name: "First message",
		},
		MultipleDummy: []*complexpb.DummyMessage{
			&complexpb.DummyMessage{
				Id:   2,
				Name: "Second message",
			},
			&complexpb.DummyMessage{
				Id:   3,
				Name: "Third message",
			},
		},
	}
	fmt.Println("----------------")
	fmt.Println("size: ", proto.Size(&cm))
	fmt.Println("json data",cm)

	data, err := proto.Marshal(&cm)
    if err != nil {
        log.Fatal("marshaling error: ", err)
	}

	fmt.Println("Binary data")
	fmt.Println(data)
}

func doEnum() {
	em := enumpb.EnumMessage{
		Id:           42,
		DayOfTheWeek: enumpb.DayOfTheWeek_THURSDAY,
	}

	em.DayOfTheWeek = enumpb.DayOfTheWeek_MONDAY
	fmt.Println("----------------")
	fmt.Println("size: ", proto.Size(&em))
	fmt.Println(em)
	data, err := proto.Marshal(&em)
    if err != nil {
        log.Fatal("marshaling error: ", err)
	}
	
	fmt.Println("Binary data")
	fmt.Println(data)
}

func jsonDemo(sm proto.Message) {
	smAsString := toJSON(sm)
	fmt.Println(smAsString)

	sm2 := &simplepb.SimpleMessage{}
	fromJSON(smAsString, sm2)
	fmt.Println("----------------")
	fmt.Println("size: ", proto.Size(sm2))
	fmt.Println("Successfully created proto struct:", sm2)
	data, err := proto.Marshal(sm2)
    if err != nil {
        log.Fatal("marshaling error: ", err)
	}
	
	fmt.Println("Binary data")
	fmt.Println(data)
}

func toJSON(pb proto.Message) string {
	marshaler := jsonpb.Marshaler{}
	out, err := marshaler.MarshalToString(pb)
	if err != nil {
		log.Fatalln("Can't convert to JSON", err)
		return ""
	}
	return out
}

func fromJSON(in string, pb proto.Message) {
	err := jsonpb.UnmarshalString(in, pb)
	if err != nil {
		log.Fatalln("Couldn't unmarshal the JSON into the pb struct", err)
	}
}

func doSimple() *simplepb.SimpleMessage {
	sm := simplepb.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "My Simple Message",
		SampleList: []int32{1, 4, 7, 8},
	}
	
	fmt.Println("----------------")
	fmt.Println(sm)
	fmt.Println("size: ", proto.Size(&sm))
	data, err := proto.Marshal(&sm)
    if err != nil {
        log.Fatal("marshaling error: ", err)
	}
	
	fmt.Println("Binary data")
	fmt.Println(data)

	fmt.Println("----------------")
	sm.Name = "I renamed you"
	fmt.Println(sm)

	fmt.Println("The ID is:", sm.GetId())
	fmt.Println("size: ", proto.Size(&sm))

	return &sm
}
