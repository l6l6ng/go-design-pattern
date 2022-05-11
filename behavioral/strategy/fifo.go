package strategy

/*具体策略*/

import "fmt"

type fifo struct {
}

func (*fifo) evict(c *cache) {
	fmt.Println("Eviction by fifo strategy")
}
