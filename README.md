# zoekt-fuzzy-search

It receives the json of the zoekt search results and performs a fuzzy search.

```sh
$ curl 'localhost:6070/search?q=keyword&num=10&format=json&ctx=5' | zoekt-fuzzy-search | xargs open
```
