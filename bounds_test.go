package geos_test

import (
	"testing"

	"github.com/alecthomas/assert/v2"

	"github.com/twpayne/go-geos"
)

func TestBounds(t *testing.T) {
	b := geos.NewBounds(1, 2, 3, 4)
	assert.True(t, b.Contains(b))
	assert.True(t, b.Contains(geos.NewBounds(1.5, 2.5, 2.5, 3.5)))
	assert.False(t, b.Contains(geos.NewBounds(1.5, 0.5, 2.5, 1.5)))
	assert.True(t, b.ContainsPoint(2, 3))
	assert.False(t, b.ContainsPoint(2, 1))
	assert.True(t, b.Equals(geos.NewBounds(1, 2, 3, 4)))
	assert.Equal(t, "POLYGON ((1.0000000000000000 2.0000000000000000, 3.0000000000000000 2.0000000000000000, 3.0000000000000000 4.0000000000000000, 1.0000000000000000 4.0000000000000000, 1.0000000000000000 2.0000000000000000))", b.Geom().ToWKT())
	assert.False(t, b.IsEmpty())
	assert.Equal(t, 2.0, b.Height())
	assert.True(t, b.Intersects(b))
	assert.True(t, b.Intersects(geos.NewBounds(1.5, 2.5, 2.5, 3.5)))
	assert.True(t, b.Intersects(geos.NewBounds(1.5, 0.5, 2.5, 3.5)))
	assert.False(t, b.Intersects(geos.NewBounds(1.5, 0.5, 2.5, 1.5)))
	assert.False(t, b.IsPoint())
	assert.Equal(t, 2.0, b.Width())
}

func TestBoundsEmpty(t *testing.T) {
	b := geos.NewBoundsEmpty()
	assert.False(t, b.Contains(b))
	assert.False(t, b.Contains(geos.NewBoundsEmpty()))
	assert.False(t, b.ContainsPoint(0, 0))
	assert.True(t, b.Equals(b)) //nolint:gocritic
	assert.Equal(t, "POINT EMPTY", b.Geom().ToWKT())
	assert.True(t, b.IsEmpty())
	assert.False(t, b.Intersects(b))
	assert.False(t, b.IsPoint())
}

func TestBoundsPoint(t *testing.T) {
	b := geos.NewBounds(0, 0, 0, 0)
	assert.True(t, b.Contains(b))
	assert.False(t, b.Contains(geos.NewBounds(1, 2, 3, 4)))
	assert.True(t, b.ContainsPoint(0, 0))
	assert.False(t, b.ContainsPoint(1, 2))
	assert.True(t, b.Equals(b)) //nolint:gocritic
	assert.False(t, b.Equals(geos.NewBounds(1, 2, 3, 4)))
	assert.False(t, b.Equals(geos.NewBoundsEmpty()))
	assert.Equal(t, "POINT (0.0000000000000000 0.0000000000000000)", b.Geom().ToWKT())
	assert.False(t, b.IsEmpty())
	assert.Equal(t, 0.0, b.Height())
	assert.True(t, b.IsPoint())
	assert.Equal(t, 0.0, b.Width())
}
