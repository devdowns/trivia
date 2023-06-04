package types

type Turn struct {
	PlayerName string
	next       *Turn
	TurnNumber int
}

var playerTurn *Turn

func AssignPlayerTurns(values []string) *Turn {
	if len(values) == 0 {
		return nil
	}
	
	// Create the first node
	first := &Turn{PlayerName: values[0], TurnNumber: 1}
	current := first
	
	// Create the remaining nodes and link them
	for i := 1; i < len(values); i++ {
		next := &Turn{PlayerName: values[i]}
		current.next = next
		current = next
	}
	
	// Make the last node point back to the first node
	current.next = first
	
	return first
}

// Next returns the next node in the cyclical list and increments the turn counter
func (n *Turn) Next() *Turn {
	n.TurnNumber++
	return n.next
}
