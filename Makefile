.PHONY: example
examples := $(shell ls -rtd example/*)
example:
	@echo "+ $@"
	@$(foreach dir,$(examples),./gqldoc -s $(dir)/schema.graphql -o $(dir);)
