package solution

import (
	"fmt"

	"github.com/kyokomi/emoji"
)

func main() {
	fmt.Println(GetMessage())
}

func GetMessage() string {
	return emoji.Sprint("Hello :world_map:!")
}
