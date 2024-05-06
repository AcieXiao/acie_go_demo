package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

var array []interface{}

func main() {

	//set := mapset.NewSet(1, 2, 3, 4)
	//set1 := mapset.NewSet("2", "3", "4", "5")
	//
	//haah := set.Difference(set1)
	//fmt.Println(haah)
	//startExpiry := 1671540923
	//startDuration := 5184000
	//_duration := 30 * 24 * 60 * 60
	//if int(time.Now().Unix()) > startExpiry {
	//	startExpiry = int(time.Now().Unix())
	//}
	//expiry := startExpiry + _duration
	//duration := startDuration + _duration
	//fmt.Printf("expiry = %v  duration = %v \n", expiry, duration)

	//ToDo

	var data = []byte(`{"status": 200}`)
	var result map[string]interface{}

	if err := json.Unmarshal(data, &result); err != nil {
		log.Fatalln(err)
	}

	//	ToDo

	defer func() {
		fmt.Println("recovered: ", recover())
	}()
	//panic("error")

	//fmt.Println("ok")

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	array = append(array, 1)
	array = append(array, "1")
	array = append(array, 1.0)

	fmt.Println(array)
	fmt.Printf("%T \n", array[0])
	fmt.Printf("%T \n", array[1])
	fmt.Printf("%T \n", array[2])

	//panic("sds")

	list1 := []int{1, 2, 3}
	list2 := make([]int, 3)
	list3 := list1

	copy(list2, list1)

	fmt.Println(list1)
	fmt.Println(list2)
	fmt.Println(list3)

	list1[0] = 10
	fmt.Println(list1)
	fmt.Println(list2)
	fmt.Println(list3)

	_list4 := []string{
		//"xkw_papers_download_total_",
		//"xkw_papers_download_fail_",
		//"xkw_knowledge_group_paper_total_",
		//"xkw_knowledge_group_paper_fail_",
		"xkw_raise_one_to_three_total_",
		"xkw_raise_one_to_three_fail_",
	}
	format := time.Now().Format("2006-01-02")

	for _, s := range _list4 {
		fmt.Println("SET " + s + format + " 0 NX EX 172800")
	}

}
