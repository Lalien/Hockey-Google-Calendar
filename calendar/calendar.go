package calendar

import (
	"fmt"
)

type CalendarEntry struct {
	Testing string
}

func (x CalendarEntry) Connect() {
	fmt.Println(x.Testing)
}
