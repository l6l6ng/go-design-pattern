package strategy

/*具体策略*/

import "fmt"

type lfu struct {

}

func (l *lfu) evict(c *cache) {
	fmt.Println("Evicting by lfu strategy")
}
