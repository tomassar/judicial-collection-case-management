.PHONY: all

all:
	find ./internal/ -name '*.templ' | entr -r sh -c 'templ generate' &
	docker-compose up
