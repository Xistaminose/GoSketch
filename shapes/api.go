package shapes

// Functions in this file provide a simplified API for creating and drawing shapes
// They should be used by the main gosketch package for convenience functions

// CreateEllipse creates a new ellipse shape without drawing it
func CreateEllipse(x, y, rx, ry float64) *EllipseShape {
	return NewEllipse(x, y, rx, ry)
}

// CreateCircle creates a new circle shape without drawing it
func CreateCircle(x, y, radius float64) *EllipseShape {
	return NewCircle(x, y, radius)
}

// CreateRectangle creates a new rectangle shape without drawing it
func CreateRectangle(x, y, w, h float64) *RectangleShape {
	return NewRectangle(x, y, w, h)
}

// CreateSquare creates a new square shape without drawing it
func CreateSquare(x, y, size float64) *RectangleShape {
	return NewSquare(x, y, size)
}

// CreateLine creates a new line shape without drawing it
func CreateLine(x1, y1, x2, y2 float64) *LineShape {
	return NewLine(x1, y1, x2, y2)
}

// CreatePoint creates a new point shape without drawing it
func CreatePoint(x, y float64) *PointShape {
	return NewPoint(x, y)
}

// CreateTriangle creates a new triangle shape without drawing it
func CreateTriangle(x1, y1, x2, y2, x3, y3 float64) *TriangleShape {
	return NewTriangle(x1, y1, x2, y2, x3, y3)
} 