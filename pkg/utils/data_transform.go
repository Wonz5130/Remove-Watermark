package utils

import (
	"google.golang.org/protobuf/types/known/structpb"

	"encoding/json"
)

func JsonToPbList(jsonStr string) (*structpb.ListValue, error) {
	var mapV []interface{}
	if err := json.Unmarshal([]byte(jsonStr), &mapV); err != nil {
		return nil, err
	}
	listV, err := structpb.NewList(mapV)
	if err != nil {
		return nil, err
	}
	return listV, err
}

func JsonToPbStruct(jsonStr string) (*structpb.Struct, error) {
	var mapV map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &mapV); err != nil {
		return nil, err
	}
	structV, err := structpb.NewStruct(mapV)
	if err != nil {
		return nil, err
	}
	return structV, err
}

func BatchJsonToPbList(jsonList []string) (resList []*structpb.ListValue, err error) {
	for i := 0; i < len(jsonList); i++ {
		pbList, err := JsonToPbList(jsonList[i])
		if err != nil {
			return resList, err
		}
		resList = append(resList, pbList)
	}
	return resList, nil
}

func BatchJsonToPbStruct(jsonList []string) (resList []*structpb.Struct, err error) {
	for i := 0; i < len(jsonList); i++ {
		pbStruct, err := JsonToPbStruct(jsonList[i])
		if err != nil {
			return resList, err
		}
		resList = append(resList, pbStruct)
	}
	return resList, nil
}
