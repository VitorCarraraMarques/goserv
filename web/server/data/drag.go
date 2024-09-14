package data

type Draggable struct {
	XPos   int 
	YPos   int 
	Width  int 
	Height int 
	Order  int
}

var MyDraggable = []Draggable{
	{
		XPos:   0,
		YPos:   0,
		Width:  100,
		Height: 100,
		Order:  0,
	},
	{
		XPos:   110,
        YPos:   0,
		Width:  100,
		Height: 100,
		Order:  1,
	},
}
