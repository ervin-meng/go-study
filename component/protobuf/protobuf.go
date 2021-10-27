package protobuf

import (
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
)

func StudyProtobuf() {
	studyRequest := StudyRequest{
		Name:     "aaa",
		Location: &Location{Longitude: 1.11, Latitude: 1.11},
		Path: []*Location{
			{
				Longitude: 1.12,
				Latitude:  1.12,
			}, {
				Longitude: 1.13,
				Latitude:  1.13,
			},
		},
		Sys: Sys_Android,
	}

	fmt.Println(&studyRequest)

	b, err := proto.Marshal(&studyRequest)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%X\n", b)

	var studyRequest1 StudyRequest

	err = proto.Unmarshal(b, &studyRequest1)

	if err != nil {
		panic(err)
	}

	fmt.Println(&studyRequest1)

	studyRequestJson, err := json.Marshal(&studyRequest1)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(studyRequestJson))
}
