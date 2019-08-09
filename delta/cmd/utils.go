package main

import (
	"fmt"
	"time"
)

func logFileName() string {
	return fmt.Sprintf("%s.log", time.Now().Format("2006-Jan-02"))
}
