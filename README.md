# Geheim

A Go binary to encrypt files

&nbsp;

## Table of contents

- [Intro](#intro)
- [Installation](#installation)
- [Configuration](#configuration)
  - [File](#file)
    - [config.yaml options](#config-yaml-options)
  - [CLI flags](#cli-flags)
  - [Environment Variables](#environment-variables)
- [Example](#example)
- [Disclaimer](#disclaimer)

&nbsp;

## Intro

The initial intent of this binary was the encryption of files before their upload to a repository.

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

&nbsp;

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

&nbsp;

## Configuration

The binary is configured using three methods, a file, CLI flags and environment variables.

Unless stated otherwise, the CLI flags take precedence over the config file, and the config file takes precedence over environment variables: `CLI flags > config.yaml > environment variable`

What can be configured:

|  | CLI flag | config.yaml | Environment variable |
| --- | --- | --- | --- |
| `secret key` | x | x | - |
| `encrypt` | x | - | - |
| `decrypt` | x | - | - |
| `files` | - | x | - |
| `log level` | - | - | x |

&nbsp;

### File

The binary looks for a config file named `config.yaml` in a folder called `.geheim` relatively to where it is ran. This takes precedence over an optional config file in `${HOME}/.geheim/config.yaml`.

If there isn't a `config.yaml` file, the secret key can be defined via CLI flag. The binary will try to work with a file called `secrets.geheim.yaml` in the same path as where the binary runs.

&nbsp;

#### config.yaml options

```yaml
---
# Should be 16 char max in length
# If it is more, it's sliced to be 16 chars in length
# If it is less, predefined chars are added to make the secret key 16 chars in length
secretkey: 'imsosecret'

# Relative to where the binary runs
# Absolute paths are not yet supported
# Defaults to ["secrets.geheim.yaml"]
files:
  - testfiles/config.json
  - testfiles/coverage_testfile.xml
  - testfiles/helpers.sh
  - testfiles/id_rsa
  - testfiles/id_rsa.pub
  - testfiles/known_hosts
  - testfiles/secrets_test.geheim.yaml
  - testfiles/simple.txt
  - testfiles/supervisor_env
```

&nbsp;

### CLI

Long flags takes precedence over shorts flags.

Flags:

- `secretkey` \ `k`
  - A key to encrypt/decrypt files. If not specified, the program will try to get one from local/global config file
- `encrypt` \ `e`
  - Whether to encrypt the files defined in the config file. Defaults to 'false'. If both encrypt and decrypt flags are set to 'true', the encrypt flag takes precedence. If both the encrypt flag and the decrypt are set to 'false', the default behavior is to encrypt any unencrypted files, ie, the encrypt flag becomes 'true'
- `decrypt` \ `d`
  - Whether to decrypt the files defined in the config file. Defaults to 'false'. If both encrypt and decrypt flags are set to 'true', the encrypt flag takes precedence. If both the encrypt flag and the decrypt are set to 'false', the default behavior is to encrypt any unencrypted files, ie, the encrypt flag becomes 'true'

Example usage:

```bash
# Encrypts
geheim \
  --secretkey 'averysecretkey'

# Decrypts
geheim \
  --secretkey 'averysecretkey' \
  -d

# Secret key is averysecretkey
geheim \
  --secretkey 'averysecretkey' \
  -k 'test'
```

&nbsp;

### Environment Variables

- `GEHEIM_LOG_LEVEL`
  - Possible values:
    - 0 -> no logs
    - 1 -> info logs (default)
    - 2 -> debug logs

&nbsp;

## Example

Assuming a file structure like:

```bash
/home/coolusername/coolproject/secrets.geheim.yaml
/home/coolusername/coolproject/.geheim/config.yaml
```

Assuming a config file like:

```yaml
---
secretkey: 'imsosecret'
files:
  - testfiles/secrets_test.geheim.yaml
```

And assuming a secrets.geheim.yaml file like:

```yaml
---
akey: a secret
anArray:
  - with secrets
aMap:
  akey: alsoWithSecrets
```

Running `geheim` in `/home/coolusername/coolproject/` would encrypt the above file to:

```txt
dodqzoudeahdidjlecc19c3eb0d22be8416553e40e823572c94d52a9b529142a9fbcbe7075effa13a3a6393274ea05c506246666b80125432750e8cebb32c2307fba80a65e0a5e5f634ab6e56665de3097c12dec77e0e430
```

&nbsp;

## Disclaimer

Beta. Use at your own risk.
