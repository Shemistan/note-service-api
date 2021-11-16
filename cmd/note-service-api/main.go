package main

import (
	"fmt"
	"os"

	"github.com/Shemistan/note-service-api/internal/app/api"
	"github.com/Shemistan/note-service-api/internal/utills"
)

func main() {
	var err error

	swappedMap, err := utills.SwapKeyAndValue(map[int64]string{1: "one"})
	if err != nil {
		fmt.Printf("failed to swapping key and value: %s", err.Error())
		return
	}

	fmt.Println(swappedMap)

	err = OpenCloseFile("cmd/note-service-api/text.txt")
	if err != nil {
		fmt.Printf("failed to opening or closing file: %s\n\n", err.Error())
	}

	fmt.Printf("function opening and closeing file comleted\n\n")

	data := []api.Note{
		{Id: 1, UserId: 1, ClassroomId: 23, DocumentId: 6},
		{Id: 2, UserId: 2, ClassroomId: 24, DocumentId: 7},
		{Id: 3, UserId: 3, ClassroomId: 23, DocumentId: 6},
		{Id: 4, UserId: 4, ClassroomId: 24, DocumentId: 7},
		{Id: 5, UserId: 5, ClassroomId: 23, DocumentId: 6},
		{Id: 555, UserId: 6, ClassroomId: 24, DocumentId: 7},
	}

	dataMap, err := utills.ConvertSliceToMap(data)
	if err != nil {
		fmt.Printf("failed to converting slice to map: %s", err.Error())
	}

	fmt.Println(dataMap)

	splitSlice, err := utills.SplitSlice(data, 5)
	if err != nil {
		fmt.Printf("failed to spliting slice: %s", err.Error())
	}

	fmt.Println(splitSlice)

	for _, v := range data {
		v.String()
	}
}

func OpenCloseFile(file string) error {
	for i := 1; i < 5; i++ {
		err := func() error {
			data, err := os.Open(file)

			defer func(data *os.File) {
				err = data.Close()
				if err != nil {
					fmt.Printf("failed to closing file: %s", err.Error())
					return
				}

				fmt.Printf("Closed the file %v times.\n\n", i)
			}(data)

			if err != nil {
				return err
			}

			fmt.Printf("Opened the file %v times. \n", i)

			return nil
		}()

		if err != nil {
			return err
		}
	}

	return nil
}
