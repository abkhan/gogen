package counter

import (
	"errors"
)

// counter generator,
// details should be a map: start: int, end: int, step: int,
func Counter(ctg <-chan struct{}, data chan<- any, detailsny any) error {

	details, ok := detailsny.(map[string]int)
	if !ok {
		return errors.New("counter generator, details is not correct map type")
	}
	start := details["start"]
	end := details["end"]
	step := details["step"]

	if ctg == nil || data == nil {
		return errors.New("counter generator, channel is nil, at start")
	}
	<-ctg
	defer close(data)

	for ix := start; ix < end; ix += step {
		if ctg == nil || data == nil {
			return errors.New("counter generator, channel is nil")
		}

		<-ctg
		data <- ix
	}
	//fmt.Printf("counter is done, data channel defer close.")
	return nil
}
