package bot

type Strategy interface {
	ShouldBuy(info Info) (bool, float64, error)
	ShouldSell(info Info) (bool, float64, error)
}
