package stats

type StatsFactory struct {
}

func (factory *StatsFactory) GetStats(statsType string) Stats {
	_ = "STUB: not implemented"
	return *new(Stats)
}
