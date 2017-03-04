include github.com/KyleBanks/make/go/sanity

# Generates the README.md documentation from the source code.
docs:
	@go get -u github.com/robertkrimen/godocdown/godocdown
	@godocdown -template=".godocdown.md" . > README.md
.PHONY: docs
