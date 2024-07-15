package split_test

import (
	"testing"

	"github.com/ccarcaci/learn-ai/perceptron/split"
	"github.com/stretchr/testify/assert"
)

func TestGetChoice(t *testing.T) {
	t.Run("ratio 0.81, splitRatio 0.8, random 0.89", func(t *testing.T) {
		// | training | > splitRatio * | testing |, r confirm => testing
		ratio := 0.81
		splitRatio := 0.8
		generateRandom := func() float64 { return 0.89 }
		choice := split.GetChoice(ratio, splitRatio, generateRandom)

		assert.Equal(t, split.Testing, choice)
	})

	t.Run("ratio 0.81, splitRatio 0.8, random 0.9", func(t *testing.T) {
		// | training | > splitRatio * | testing |, r deny => training
		ratio := 0.81
		splitRatio := 0.8
		generateRandom := func() float64 { return 0.9 }
		choice := split.GetChoice(ratio, splitRatio, generateRandom)

		assert.Equal(t, split.Training, choice)
	})

	t.Run("ratio 0.8, splitRatio 0.81, random 0.89", func(t *testing.T) {
		// | training | < splitRatio * | testing |, r confirm => training
		ratio := 0.8
		splitRatio := 0.81
		generateRandom := func() float64 { return 0.89 }
		choice := split.GetChoice(ratio, splitRatio, generateRandom)

		assert.Equal(t, split.Training, choice)
	})

	t.Run("ratio 0.8, splitRatio 0.81, random 0.9", func(t *testing.T) {
		// | training | < splitRatio * | testing |, r deny => testing
		ratio := 0.8
		splitRatio := 0.81
		generateRandom := func() float64 { return 0.9 }
		choice := split.GetChoice(ratio, splitRatio, generateRandom)

		assert.Equal(t, split.Testing, choice)
	})
}
