all:
	go run draw.go image.go main.go matrix.go parser.go

run:
	display mat.ppm

clean:
	rm *.ppm
