build:
	go build -o bin/cmap .

rpm: build
	fpm -s dir -t rpm \
	-p cmap-0.1.0-1-any.rpm \
	--name cmap \
	--license mit \
	--version 0.1.0 \
	--architecture all \
	--depends bash --depends lolcat \
	--description "Simple way to bookmark commands" \
	--url "" \
	--maintainer "isahmed0149@gmail.com" \
	cmap=./bin/cmap

deb: build
	fpm -s dir -t deb \
	-p cmap-0.1.0-1-any.deb \
	--name cmap \
	--license mit \
	--version 0.1.0 \
	--architecture all \
	--depends bash --depends lolcat \
	--description "Simple way to bookmark commands" \
	--url "" \
	--maintainer "isahmed0149@gmail.com" \
	cmap=./bin/cmap

install:
	go install





