package testlib

import "fmt"
var pop map[string]string

func init() {
	// 이 package를 호출하는 곳이 있다면 프로젝트 시작시에 load된다.
	// no args, no return val
	pop = make(map[string]string)
	pop["Adele"] = "Hello"
	pop["Alicia Keys"] = "Fallin'"
	pop["John Legend"] = "All of Me"
}

// GetMusic : Popular music by singer (외부에서 호출 가능)
func GetMusic(singer string) string {
	getKeys()
	return pop[singer]
}

func getKeys() {  // 내부에서만 호출 가능
	for _, kv := range pop {
		fmt.Println(kv)
	}
}