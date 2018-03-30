package humanize

import "fmt"

// Bytes2Human convert uint64 to string like 10.00 KB (if uint64 is 10000 and postfix is "B")
// Converting proceed in SI or non-SI system depends on
// For extended level or if we simply need "Z" and "Y" use - https://github.com/dustin/go-humanize
func Bytes2Human(n uint64, postfix string, divider float64) string {
	num := float64(n)
	prefix := []string{"", "K", "M", "G", "T", "P", "E"}
	unit := ""
	for _, unt := range prefix {
		if num < divider {
			unit = unt
			break
		}
		num = num / divider
	}
	return fmt.Sprintf("%3.2f %s%s", num, unit, postfix)
}
