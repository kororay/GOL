package gol

// Params provides the details of how to run the Game of Life and which image to load.
type Params struct {
	Turns       int
	Threads     int
	ImageWidth  int
	ImageHeight int
}

// Run starts the processing of Game of Life. It should initialise channels and goroutines.
func Run(p Params, events chan<- Event, keyPresses <-chan rune) {

	ioCommand := make(chan ioCommand)
	ioIdle := make(chan bool)
	//added these 3
	ioFilename := make(chan string)
	outputQ := make(chan uint8)
	inputQ := make(chan uint8)

	distributorChannels := distributorChannels{
		events,
		ioCommand,
		ioIdle,
		outputQ,
		inputQ,
	}
	go distributor(p, distributorChannels)

	ioChannels := ioChannels{
		command: ioCommand,
		idle:    ioIdle,
		//these were nil before
		filename: ioFilename,
		output:   outputQ,
		input:    inputQ,
	}
	go startIo(p, ioChannels)
}
