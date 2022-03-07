package gpool

type Statistics struct {
	TotalWork uint64
	TotalWait uint64

	WorkStatisticsMap map[string]*WorkStatistics
}

type WorkStatistics struct {
	TotalWork uint64
	TotalWait uint64
}
