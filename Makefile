all:
	go run image.go main.go matrix.go parser.go

run:
	display mat.ppm

clean:
	rm *.ppm
