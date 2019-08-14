all: promptpath

promptpath: Dockerfile *.cr
	img=$$(docker build -q .) \
		&& docker run --rm $$img > $@ \
		&& chmod 744 $@
