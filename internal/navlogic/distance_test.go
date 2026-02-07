package navlogic

import "testing"

func TestEchoCountToDistanceCm(t *testing.T) {
	tests := []struct {
		name      string
		echoCount int
		want      int
	}{
		// Normal cases
		{"58 counts = 1 cm", 58, 1},
		{"116 counts = 2 cm", 116, 2},
		{"580 counts = 10 cm", 580, 10},
		{"1160 counts = 20 cm (obstacle threshold)", 1160, 20},
		{"5800 counts = 100 cm", 5800, 100},

		// Edge cases
		{"0 counts = timeout", 0, TimeoutDistance},
		{"negative counts = timeout", -1, TimeoutDistance},
		{"57 counts = 0 cm (below 1 cm)", 57, 0},
		{"59 counts = 1 cm", 59, 1},

		// Boundary values
		{"1 count = 0 cm", 1, 0},
		{"29 counts = 0 cm (half of threshold)", 29, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := EchoCountToDistanceCm(tt.echoCount)
			if got != tt.want {
				t.Errorf("EchoCountToDistanceCm(%d) = %d, want %d", tt.echoCount, got, tt.want)
			}
		})
	}
}

func TestIsWithinThreshold(t *testing.T) {
	const obstacleThreshold = 20 // cm

	tests := []struct {
		name      string
		distance  int
		threshold int
		want      bool
	}{
		// Obstacle detected (within threshold)
		{"5 cm < 20 cm threshold", 5, obstacleThreshold, true},
		{"19 cm < 20 cm threshold", 19, obstacleThreshold, true},
		{"1 cm < 20 cm threshold", 1, obstacleThreshold, true},

		// No obstacle (at or beyond threshold)
		{"20 cm = 20 cm threshold (not within)", 20, obstacleThreshold, false},
		{"21 cm > 20 cm threshold", 21, obstacleThreshold, false},
		{"100 cm > 20 cm threshold", 100, obstacleThreshold, false},

		// Timeout case
		{"timeout = no obstacle", TimeoutDistance, obstacleThreshold, false},

		// Edge cases
		{"0 cm < 20 cm threshold", 0, obstacleThreshold, true},
		{"0 cm with 0 threshold", 0, 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsWithinThreshold(tt.distance, tt.threshold)
			if got != tt.want {
				t.Errorf("IsWithinThreshold(%d, %d) = %v, want %v",
					tt.distance, tt.threshold, got, tt.want)
			}
		})
	}
}

// Benchmark for performance-critical distance calculation
func BenchmarkEchoCountToDistanceCm(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EchoCountToDistanceCm(1160) // typical obstacle detection range
	}
}
