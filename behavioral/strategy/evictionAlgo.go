package strategy

/*具体策略*/

type evictionAlgo interface {
	evict(c *cache)
}
