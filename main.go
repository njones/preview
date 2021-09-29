package mainly

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello, world!"
	fmt.Printf("%#v\n", Breakdown1(str))
	fmt.Printf("%#v\n", Breakdown2(str))
	fmt.Printf("%#v\n", Breakdown3(str))
	fmt.Printf("%#v\n", Breakdown4(str))
	fmt.Printf("%#v\n", Breakdown5(str))

	fmt.Printf("%q\n", AcceptStruct1(Breakdown1(str)))
	fmt.Printf("%q\n", AcceptStruct1(Breakdown4(str)))
}

func Breakdown1(str string) struct {
	hello string
	world string
} {
	strs := strings.Split(str, " ")
	return struct {
		hello, world string
	}{
		hello: strs[0],
		world: strs[1],
	}
}

type brkdn2 = struct{ hello, world string }

func Breakdown2(str string) brkdn2 {
	strs := strings.Split(str, " ")
	return brkdn2{hello: strs[0], world: strs[1]}
}

type Brkdn3 = struct{ hello, world string }

func Breakdown3(str string) Brkdn3 {
	strs := strings.Split(str, " ")
	return Brkdn3{hello: strs[0], world: strs[1]}
}

type brkdn4 struct{ hello, world string }

func Breakdown4(str string) brkdn4 {
	strs := strings.Split(str, " ")
	return brkdn2{hello: strs[0], world: strs[1]}
}

func Breakdown5(str string) struct{ hello, world string } {
	strs := strings.Split(str, " ")
	return brkdn2{hello: strs[0], world: strs[1]}
}

func AcceptStruct1(str struct{ hello, world string }) string {
	return fmt.Sprintf("%s %s", str.hello, str.world)
}

func AcceptStruct2(str brkdn2) string {
	return fmt.Sprintf("%s %s", str.hello, str.world)
}

func AcceptStruct3(str brkdn4) string {
	return fmt.Sprintf("%s %s", str.hello, str.world)
}
