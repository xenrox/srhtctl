# srhtctl

[![builds.sr.ht status](https://builds.xenrox.net/~xenrox/srhtctl.svg)](https://builds.xenrox.net/~xenrox/srhtctl?)

srhtctl is a CLI for interacting with the sourcehut API.

A goal of this project is that you can use sourcehut from your terminal just like you would from your browser.

## Installation

Just download the source code und build it with `make`.

There is an [aur package](https://aur.archlinux.org/packages/srhtctl/) for Arch Linux.

## Usage

You have to create a `config.ini` in your `XDG_CONFIG_HOME` under the srhtctl folder.
On Darwin, your configuration directory is `~/Library/Application Support/srhtctl`.
The only necessary value is your sourcehut authentication token.
By default you will interact with https://meta.sr.ht/.

Currently implemented are parts of the git, paste and builds api.
You can for example create pastes or deploy build manifests from your command line.
As extra features you can create pastes with expiration times and edit build files on the fly with your favourite `$EDITOR`.

## Comments

When using the zsh completion, you should apply the patch under `assets` so that files will be shown where appropriate.
There is a systemd service example in `assets` for cleaning up your expired pastes.

TODO: Add checklist with implemented features.
