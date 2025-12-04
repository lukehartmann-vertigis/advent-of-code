package collections

import (
	"errors"
)

type Matrix[T any] struct {
	Items  []T
	Width  int
	Height int
}

func NewMatrix[T any](items []T, width int) *Matrix[T] {
	return &Matrix[T]{
		Items:  items,
		Width:  width,
		Height: len(items) / width,
	}
}

func (m *Matrix[T]) At(x, y int) (T, error) {
	var val T
	if y < 0 || x < 0 || y >= m.Height || x >= m.Width {
		return val, errors.New("index out of range")
	}
	idx := y*m.Width + x
	return m.Items[idx], nil
}

func (m *Matrix[T]) Set(x, y int, val T) error {
	if y < 0 || x < 0 || y >= m.Height || x >= m.Width {
		return errors.New("index out of range")
	}
	idx := y*m.Width + x
	m.Items[idx] = val
	return nil
}

func (m *Matrix[T]) VonNeumann(x, y int) ([]T, error) {
	var zero T
	if x < 0 || y < 0 || x >= m.Width || y >= m.Height {
		return []T{}, errors.New("index out of range")
	}

	items := []T{zero, zero, zero, zero} // top, right, bottom, left

	// top
	if v, err := m.At(x, y-1); err == nil {
		items[0] = v
	}

	// right
	if v, err := m.At(x+1, y); err == nil {
		items[1] = v
	}

	// bottom
	if v, err := m.At(x, y+1); err == nil {
		items[2] = v
	}

	// left
	if v, err := m.At(x-1, y); err == nil {
		items[3] = v
	}

	return items, nil
}

func (m *Matrix[T]) Moore(x, y int) ([]T, error) {
	if x < 0 || y < 0 || x >= m.Width || y >= m.Height {
		return nil, errors.New("index out of range")
	}

	directions := [][2]int{
		{0, -1},  // top
		{1, -1},  // top right
		{1, 0},   // right
		{1, 1},   // bottom right
		{0, 1},   // bottom
		{-1, 1},  // bottom left
		{-1, 0},  // left
		{-1, -1}, // top left
	}

	items := make([]T, 8)
	var zero T

	for idx, d := range directions {
		nx := x + d[0]
		ny := y + d[1]
		v, err := m.At(nx, ny)
		if err != nil {
			items[idx] = zero
			continue
		}
		items[idx] = v
	}

	return items, nil
}
