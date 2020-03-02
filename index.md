# srhtctl

The command structure is `srhtctl <service> <command> [FLAGS] [<args>]`.

## Global flags

- `--config <path>`: Set config.ini path, for example if you have one for
  a self hosted and one for the original instance.

## Config

- `token`: Your personal access token
- `copyToClipboard` (true, false): Whether to copy some replies (e.g.
  link to paste) instead of just printing them [default: false]
- `editor`: Which editor you want to use (e.g. for editing build manifests)
  [default: $EDITOR]
- `user`: Default user for API operations (e.g. git annotate) [default: ""]

Every service has an `url` that should point towards your service.

## Services

- [Meta](meta.md)
- [Git](git.md)
- [Paste](paste.md)
- [Builds](builds.md)
