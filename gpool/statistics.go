package gpool

import "time"

type Statistics struct {
	TotalTask uint64
	TotalWork uint32
	TotalWait uint64

	WorkStatisticsMap map[string]*WorkStatistics
}

type WorkStatistics struct {
	Id   string
	Name string

	TotalTask uint64
	TotalWork uint32
	TotalWait uint64

	AvgProcessingTime time.Duration
}
