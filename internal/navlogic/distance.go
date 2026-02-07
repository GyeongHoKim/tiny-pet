package navlogic

// Ultrasonic sensor constants.
const (
	MicrosecondsPerCm = 58
	TimeoutDistance   = -1
)

// EchoCountToDistanceCm converts ultrasonic echo loop count to distance in cm.
// Returns TimeoutDistance (-1) if echoCount is invalid.
func EchoCountToDistanceCm(echoCount int) int {
	if echoCount <= 0 {
		return TimeoutDistance
	}
	return echoCount / MicrosecondsPerCm
}

// IsWithinThreshold returns true if distance is positive and below threshold.
func IsWithinThreshold(distance, threshold int) bool {
	if distance == TimeoutDistance {
		return false
	}
	return distance < threshold
}
