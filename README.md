# FSAND

Pronounced 'F-sand'

A simple command line utility to run a script when either a file or a file within a directory changes. This is built entirely off fsnotify and beeep and I take no credit for any of those projects. This is merely the glue that binds; in the end isn't that what most software is anyway?

## Installation

To create a build, run `make build` from the root directory. If you copy this to your `/usr/local/bin` the script you specify to run will still be relative to the current working directory.


## Usage

```bash
fsand changing_dir ./to_run.sh
```

## Development

Feel free to fork this and tinker to your hearts content. Just if you somehow make money off it, let me know so I can tell my friends and family I've made it in this cruel harsh world.
