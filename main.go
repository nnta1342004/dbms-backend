package main

import (
	"hareta/cmd"
	"time"
)

func main() {

	utcPlus7 := time.FixedZone("UTC+7", 7*60*60)

	// Set the local time zone to UTC+7
	time.Local = utcPlus7

	cmd.Execute()

}
