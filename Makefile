build:
	go build -o bin/command-map .

rpm: build
	fpm -s dir -t rpm \
	-p cmap-0.1.0-1-any.rpm \
	--name command-map \
	--license mit \
	--version 0.1.0 \
	--architecture all \
	--depends bash --depends lolcat \
	--description "Simple way to bookmark commands" \
	--url "" \
	--maintainer "isahmed0149@gmail.com" \
	command-map=./bin/command-map

deb: build
	fpm -s dir -t deb \
	-p cmap-0.1.0-1-any.deb \
	--name command-map \
	--license mit \
	--version 0.1.0 \
	--architecture all \
	--depends bash --depends lolcat \
	--description "Simple way to bookmark commands" \
	--url "" \
	--maintainer "isahmed0149@gmail.com" \
	command-map=./bin/command-map

install:
	go install





