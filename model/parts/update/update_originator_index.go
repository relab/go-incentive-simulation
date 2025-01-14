package update

import (
	"go-incentive-simulation/config"
	"go-incentive-simulation/model/parts/types"
	"sync/atomic"
)

// OriginatorIndex Used by the requestWorker
func OriginatorIndex(state *types.State, timeStep int) int64 {

	curOriginatorIndex := atomic.LoadInt64(&state.OriginatorIndex)

	if config.GetSameOriginator() {
		if (timeStep)%100 == 0 {
			if int(curOriginatorIndex+1) >= config.GetOriginators() {
				atomic.StoreInt64(&state.OriginatorIndex, 0)
				return 0
			} else {
				return atomic.AddInt64(&state.OriginatorIndex, 1)
			}
		}
	} else {
		if int(curOriginatorIndex+1) >= config.GetOriginators() {
			atomic.StoreInt64(&state.OriginatorIndex, 0)
			return 0
		} else {
			if config.GetSameOriginator() {
				if atomic.LoadInt64(&state.TimeStep)%100 == 0 {
					return atomic.AddInt64(&state.OriginatorIndex, 1)
				}
			} else {
				return atomic.AddInt64(&state.OriginatorIndex, 1)
			}
		}
	}
	return curOriginatorIndex
	//state.OriginatorIndex = rand.Intn(config.GetOriginators() - 1)
}
