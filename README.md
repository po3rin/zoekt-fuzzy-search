# zoekt-fuzzy-search

It receives the json of the zoekt search results and performs a fuzzy search.

# usage

```sh
$ curl 'localhost:6070/search?q=keyword&num=10&format=json&ctx=5' | zoekt-fuzzy-search | xargs open
```

For example, to define a function in zsh...

```sh
# zoekt
zs() {
  if [ -z "$1" ]; then
    echo "Usage: zs <query>"
    return 1
  fi
  curl -s "http://localhost:6070/search?q=$1&num=10&ctx=5&format=json" | zoekt-fuzzy-search | xargs open
}
```

It is recommended to set the ctx option to make the preview easier to view.

