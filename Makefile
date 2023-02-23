default_goal := build
build:
	@./scripts/build
run:
	@./scripts/run 
buildall:
	@./scripts/buildall
runspecific:
	@./scripts/runspecific $(os) $(arch)
