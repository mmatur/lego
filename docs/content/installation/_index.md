---
title: "Installation"
date: 2019-03-03T16:39:46+01:00
weight: 1
draft: false
---

## Binaries

To get the binary just download the latest release for your OS/Arch from [the release page](https://github.com/go-acme/lego/releases) and put the binary somewhere convenient.
lego does not assume anything about the location you run it from.

## From Docker

```bash
docker run goacme/lego -h
```

## From package managers

- [ArchLinux](https://archlinux.org/packages/extra/x86_64/lego/) (official):

  ```bash
  pacman -S lego
  ```

- [ArchLinux (AUR)](https://aur.archlinux.org/packages/lego-bin) (official):

  ```bash
  yay -S lego-bin
  ```

- [Snap](https://snapcraft.io/lego) (official):

  ```bash
  sudo snap install lego
  ```
  Note: The snap can only write to the `/var/snap/lego/common/.lego` directory.

- [FreeBSD (Ports)](https://www.freshports.org/security/lego) (unofficial):

  ```bash
  pkg install lego
  ```

- [Gentoo](https://gitweb.gentoo.org/repo/proj/guru.git/tree/app-crypt/lego) (unofficial):

  You can [enable GURU](https://wiki.gentoo.org/wiki/Project:GURU/Information_for_End_Users) repository and then:

  ```bash
  emerge app-crypt/lego
  ```

- [Homebrew](https://formulae.brew.sh/formula/lego) (unofficial):

  ```bash
  brew install lego
  ```

  or

  ```bash
  pkg install lego
  ```

## From sources

Requirements:

- go1.22+.
- environment variable: `GO111MODULE=on`

To install the latest version from sources, just run:

```bash
go install github.com/go-acme/lego/v4/cmd/lego@latest
```

or

```bash
git clone git@github.com:go-acme/lego.git
cd lego
make        # tests + doc + build
make build  # only build
```
