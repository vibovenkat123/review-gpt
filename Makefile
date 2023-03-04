default_goal := build
build:
	@./scripts/build
run:
	@./scripts/run 
buildall:
	@./scripts/buildall
runspecific:
	@./scripts/runspecific $(os) $(arch)
buildman:
	@./scripts/buildman
showman:
	@./scripts/showman $(pager)
outputman:
	@./scripts/man $(pager)
