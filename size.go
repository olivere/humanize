package humanize

import "fmt"

// Size prints a size in human readable form (as kB, MB etc.).
func Size(size int64) string {
	if size < 1024 {
		return fmt.Sprintf("%d byte", size)
	}
	if size < 1024*1024 {
		return fmt.Sprintf("%d kB", size/1024)
	}
	if size < 1024*1024*1024 {
		return fmt.Sprintf("%d MB", size/1024/1024)
	}
	return fmt.Sprintf("%d GB", size/1024/1024/1024)
}
