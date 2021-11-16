package utills

import (
	"errors"
	"fmt"

	"github.com/Shemistan/note-service-api/internal/app/api"
)

var filter = []string{"a", "b", "c", "d"}

func SwapKeyAndValue(data map[int64]string) (map[string]int64, error) {
	res := make(map[string]int64)

	for key, value := range data {
		if _, found := res[value]; found {
			fmt.Println("The key exists", value)
			continue
		}

		res[value] = key
	}

	return res, nil
}

func FilterSlice(data []string) ([]string, error) {
	dataMap := make(map[string]struct{})
	var res []string

	for _, val := range data {
		dataMap[val] = struct{}{}
	}

	for _, val := range filter {
		if _, found := dataMap[val]; found {
			res = append(res, val)
		}
	}

	return res, nil
}

func ConvertSliceToMap(users []api.Note) (map[int64]api.Note, error) {
	res := make(map[int64]api.Note)
	for _, v := range users {
		res[v.Id] = v
	}

	return res, nil
}

func SplitSlice(notes []api.Note, batchSize int64) ([][]api.Note, error) {
	if batchSize <= 0 || notes == nil {
		return nil, errors.New("error input values")
	}

	if int64(len(notes)) <= batchSize {
		return [][]api.Note{notes}, nil
	}

	numBatches := int64(len(notes)) / batchSize
	if int64(len(notes))%batchSize != 0 {
		numBatches++
	}

	var end int64

	res := make([][]api.Note, 0, numBatches)

	for begin := int64(0); begin < int64(len(notes)); {
		end += batchSize
		if end > int64(len(notes)) {
			end = int64(len(notes))
		}

		res = append(res, notes[begin:end])
		begin += batchSize
	}

	return res, nil
}
