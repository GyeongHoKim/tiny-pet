package navlogic

const (
	MicrosecondsPerCm = 58
	TimeoutDistance   = -1
)

// EchoCountToDistanceCm converts echo loop count to distance in cm; returns TimeoutDistance (-1) if invalid.
func EchoCountToDistanceCm(echoCount int) int {
	if echoCount <= 0 {
		return TimeoutDistance
	}
	return echoCount / MicrosecondsPerCm
}

// EchoMicrosecondsToDistanceCm converts echo pulse width (µs) to distance in cm; returns TimeoutDistance (-1) if invalid.
func EchoMicrosecondsToDistanceCm(us int) int {
	if us <= 0 {
		return TimeoutDistance
	}
	return us / MicrosecondsPerCm
}

// IsWithinThreshold reports whether distance is valid and below threshold.
func IsWithinThreshold(distance, threshold int) bool {
	if distance == TimeoutDistance {
		return false
	}
	return distance < threshold
}
