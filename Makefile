include github.com/KyleBanks/make/go/sanity
include github.com/KyleBanks/make/git/precommit

# Generates the README.md documentation from the source code.
docs:
	@go get -u github.com/robertkrimen/godocdown/godocdown
	@godocdown -template=".godocdown.md" . > README.md
.PHONY: docs
