# dignore

`dignore` is a cli tool to list .dockerignore-d files

## Installation

### Install from Releases (Linux)
1. download the latest (or any) version from the releases page
2. expose the `dignore` executable to `$PATH`

### Install with Git
1. install go: https://go.dev/doc/install
2. `git clone https://github.com/ssumoo/dignore.git`
3. `cd dignore && go install`

## Usage

```bash
dignore list --help  # show helps
```

### Examples

```bash
dignore list  # in $PWD, read .dockerignore, and print all excluded files
dignore list --include  # and print all included files
dignore list --all  # print all included and excluded files
dignore list --include --quiet > included_files.txt
dignore list | fzf  # run dignore + pipe into the fzf search
```

## Limitations

- Currently no checkings is done on each lines read from the `.dockerignore` file to flag / reject invalid lines. Please
give it a valid `.dockerignore` file
    - If this feature is highly wanted please let me know! so I know it's a feature worth building for.
