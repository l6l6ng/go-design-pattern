package strategy

/*具体策略*/

import "fmt"

type lru struct {

}

func (l *lru) evict(c *cache) {
	fmt.Println("Evicting by lru strategy")
}
