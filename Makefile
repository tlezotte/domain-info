BRANCH	 = main
PACKAGE  = whois
VERSION	 = `date "+%Y.%m%d%"`
NEXT     = `autotag -b $(BRANCH) -n`
RELEASE_DIR  = ..
RELEASE_FILE = $(PACKAGE)-$(VERSION)

.SILENT: all
.SILENT: release-next
.SILENT: _patch


# Default target.
all:
	echo "Hello $(LOGNAME), nothing to do by default"

release:
	autotag -b $(BRANCH)
#	git tag -a v$(NEXT) -m "Release v$(NEXT)"
#	git push origin v$(NEXT)
	goreleaser release

release-next:
	echo "Next release: v$(NEXT)"

release-build:
	goreleaser release --snapshot --rm-dist

release-notag:
	goreleaser release


_patch:
	# fetch all tags and history:
	git fetch --tags --prune

	if [ `git rev-parse --abbrev-ref HEAD` != "$(BRANCH)" ]; then
		# ensure a local branch exists at 'refs/heads/master'
		git branch --track master origin/$(BRANCH)
	fi

_software:
	brew install goreleaser
	brew install nfpm
	curl -sL https://git.io/autotag-install | sh -s -- -b $HOME/bin

_packages:
	go get -u github.com/gookit/color
	go get -u github.com/likexian/whois
	go get -u github.com/likexian/whois-parser
