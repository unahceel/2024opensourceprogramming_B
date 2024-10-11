package main

func main() {
	var now time.Time = time.Now()
	fmt.Printf("오늘은 %d년 %d월 %d일 입니다\n", now.Year(), int(now.Month()), now.Day())
	fmt.Printf("지금 시각은  %d시 %d분 %d초 입니다\n", now.Hour(), int(now.Minute()), now.Second())



}

package main

import (
	"fmt"
	"strings"
)

func main() {
	var army string = "우리는 !가와 !민에 충성을 다하는 대한민! 육군이다"
	armyfixed := strings.NewReplacer("!", "국")
	fmt.Println(army)
	fmt.println(armyFixed.Peplace(army))

}

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print("이름 입력 : ")
	r := buflo.NewReader(os.stdin)
	name, err := in.ReadString('\n')
	fmt.Println(name)
	fmt.prinln(err)

}

package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Print("이름 입력 : ")
	r := buflo.NewReader(os.stdin)
	name, err := in.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(name)
	}

}
