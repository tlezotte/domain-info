BRANCH	 = main
PACKAGE  = whois
VERSION	 = `date "+%Y.%m%d%"`
NEXT     = `autotag -b $(BRANCH) -n`
RELEASE_DIR  = ..
RELEASE_FILE = $(PACKAGE)-$(VERSION)

.SILENT: all
.SILENT: next

# Default target.
all:
	echo "Hello $(LOGNAME), nothing to do by default"
	echo "Try 'make help'"

next:
	echo "Next release: v$(NEXT)"

release:
#	git tag -a v$(NEXT) -m "Release v$(NEXT)"
#	git push --tags origin $(BRANCH)
	goreleaser release

_software:
	brew install goreleaser
	brew install nfpm
	curl -sL https://git.io/autotag-install | sh -s -- -b $HOME/bin

_packages:
	go get -u github.com/gookit/color
	go get -u github.com/likexian/whois
	go get -u github.com/likexian/whois-parser
