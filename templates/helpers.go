package templates

import (
	"fmt"
)

func formatPrice(price float64) string {
	return fmt.Sprintf("%.2f", price)
}

func formatRating(rating float64) string {
	return fmt.Sprintf("%.1f", rating)
}
