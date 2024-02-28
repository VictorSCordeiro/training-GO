# Makefile

# this line disables CGO in the build process to ensure that the Go program be statically linked 
# and doesn't depend on dinamic libraries.
export CGO_ENABLED=0

# Here the Go code is compiled in the current directory and an executable name 'www' is generated without a file extension
# GOOS=linux added to enable www to run on linux docker
build:
	GOOS=linux go build -o www

#runs the tests in the directory
test:
	go test


#depends on the compilation target, ensuring that the program is created before it is executed.
#./www executes the compiled binary.
run: build
	./www


# extending the make file with docker targets
docker-build:
	docker build -t www .

docker-run:
	docker run -p 3333:3333 www