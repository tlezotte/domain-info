## Prep Server for Deployment

### Install Dependencies

Use make to install all macOS dependencies. `_software` installs the software listed below.

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
  cd $HOME
  curl -sL https://git.io/autotag-install | sh --
  ```

### Install Packages

Use make to install all Go packages. `_packages` installs the required Go packages.

```
make _packages
```
