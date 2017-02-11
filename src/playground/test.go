package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type KeyValue struct {
	Key   string
	Value string
}

func main() {
	var kvl []KeyValue
	kvl = append(kvl, KeyValue{"hi", "go"})
	for id, ele := range kvl {
		fmt.Println(id, ele)
	}
	mp := make(map[string]*os.File)
	mp["asd"] = nil
	for k, v := range mp {
		fmt.Println(k, v)
	}
	fmt.Println(mp["aaa"] == nil)
	lst := []string{"aaa", "bbb"}
	mpp := make(map[string][]string)
	mpp["hi"] = lst
	mppB, _ := json.Marshal(mpp)
	mpp2 := make(map[string][]string)
	json.Unmarshal(mppB, &mpp2)
	fmt.Println(mpp2)
	f, err := os.OpenFile("tmp", os.O_APPEND|os.O_RDWR, os.ModeAppend)
	if err != nil {
		fmt.Println(err)
		if os.IsNotExist(err) {
			f, err = os.Create("tmp")
		} else {
			panic(err)
		}
	}
	defer f.Close()
	fmt.Println(2 % 5)
	n1, err := f.Write([]byte("bbb\n"))
	fmt.Println(err, n1)
	n2, err := f.Write([]byte("ccc\n"))
	fmt.Println(err, n2)

	lst2 := make([]string, 5)
	fmt.Println(len(lst2))
	lst2 = append(lst2, "asd")
	fmt.Println(lst2)
}
