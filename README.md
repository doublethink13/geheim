# geheim

A Go program to store secrets in repositories

## Table of contents

## Intro

The initial intent of this program was the encryption of files before their upload to a repository.

For example, imagine you have the following file:

```yaml
---
akey: a secret
anArray:
  - with secrets
aMap:
  akey: alsoWithSecrets

```

Using `geheim`, the encryption version would look like so:

```txt
dodqzoudeahdidjlea878c1ef6223ac6ccf284ec1451f9565e7f824919f68864765f1b472322e51d44f6eb4ffbb51c1bd7ecd15156adb60aef1316ac4ecc46839f41c08fb35883d432e79b53f0b19cd7d68b507b6496920a
```

## Installation

Go to the [Releases](https://github.com/doppeldenken/geheim/releases) page and find the version you are looking for.

Multiple binaries are available, built using [Go Releaser](https://goreleaser.com/) (config file [here](.goreleaser.yaml)).

For example, using the latest released version:

```bash
wget https://github.com/doppeldenken/geheim/releases/download/0.3.0-dev/geheim_0.3.0-dev_Linux_x86_64.tar.gz && \
tar -xzf geheim* && \
mv geheim /usr/local/bin && \
rm -rf geheim*
```

## Configuration

### File

### CLI

### Environment Variables

## Example

## Disclaimer

For now, this is only a learning project. Use at your own risk.
