# dignore

`dignore` is a cli tool to list .dockerignore-d files
![demo](demo.gif)

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
dignore list  # in $PWD, read .dockerignore, and print all included files
dignore list --path <my_new_path_that_is_not_cwd>
dignore list --dockerignore <my_new_dockerignore_path_that_is_not_.dockerignore>
dignore list --excluded > excluded_files.txt  # list all excluded files and write to a text file
```
