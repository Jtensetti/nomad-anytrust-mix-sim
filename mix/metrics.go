package mix

import "math"

// RandomGuessRecall is the expected recall if an adversary guesses targetCount
// output positions uniformly from a batch of size batchSize.
func RandomGuessRecall(batchSize, targetCount int) float64 {
	if batchSize <= 0 || targetCount <= 0 {
		return 0
	}
	return math.Min(1, float64(targetCount)/float64(batchSize))
}
