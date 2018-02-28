package raceaverage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAvgMinutesSampleSet1(t *testing.T) {
	sampleSet1 := []string{
		"02:00 PM, DAY 19",
		"02:00 PM, DAY 20",
		"01:58 PM, DAY 20",
	}

	result, err := AvgMinutes(sampleSet1)
	assert.Nil(t, err)
	assert.Equal(t, 27239, result)
}

func TestAvgMinutesSampleSet2(t *testing.T) {
	sampleSet2 := []string{"12:00 AM, DAY 2"}

	result, err := AvgMinutes(sampleSet2)
	assert.Nil(t, err)
	assert.Equal(t, 960, result)
}

func TestAvgMinutesSampleSet3(t *testing.T) {
	sampleSet3 := []string{
		"12:00 PM, DAY 1",
		"12:01 PM, DAY 1",
	}

	result, err := AvgMinutes(sampleSet3)
	assert.Nil(t, err)
	assert.Equal(t, 241, result)
}
