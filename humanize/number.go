package humanize

import "fmt"

// Bytes2Human convert uint64 to string like 10.00 KB (if uint64 is 10000 and postfix is "B")
// Converting proceed in SI system
// For extended level or if we simply need "Z" and "Y" use - https://github.com/dustin/go-humanize
func Bytes2Human(n uint64, postfix string) string {
	num := float64(n)
	prefix := []string{"", "K", "M", "G", "T", "P", "E"}
	unit := ""
	for _, unt := range prefix {
		if num < 1000 {
			unit = unt
			break
		}
		num = num / 1000
	}
	return fmt.Sprintf("%3.2f %s%s", num, unit, postfix)
}


