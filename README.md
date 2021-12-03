## Prep for Deployment

### Install Software

Use `make` to install all macOS system software.  Installs the software listed below.

```
brew install make
make _software
```

- Install GoReleaser

  ```
  brew install goreleaser
  ```

- Install nFPM

  ```
  brew install nfpm
  ```

- Install AutoTag

  ```
  curl -sL https://git.io/autotag-install | sh -s -- -b $HOME/bin
  ```

### Install Packages

Use `make` to install or update all Go packages.

```
make _packages
```
