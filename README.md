# dignore

`dignore` is a cli tool to list .dockerignore-d files
![demo_gif](dignore_demo.gif)

## Installation

### Install from Releases (Linux)
```bash
wget https://github.com/ssumoo/dignore/releases/download/v2.0.0/dignore -O $HOME/.local/bin/dignore && chmod +x $HOME/.local/bin/dignore
```
- in general, this should work right away
- change `v2.0.0` to other release tags if required
- make sure `$HOME/.local/bin` is in `$PATH`


## Usage

```bash
dignore list --help  # show helps
dignore list  # in $PWD, read .dockerignore, and print all included files
dignore list --path <my_new_path_that_is_not_cwd>
dignore list --dockerignore <my_new_dockerignore_path_that_is_not_.dockerignore>
dignore list --excluded > excluded_files.txt  # list all excluded files and write to a text file
```
