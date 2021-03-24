About
=============

Basic Go Lang API endpoint based on the instructions found here: https://github.com/tobym/cartesian.

How to run
==========

On terminal run the following command lines:

```sh
cd /path/to/cartesian/project
go run . --port 8000
```

How to test
===========

Here's an example using HTTPie:

```sh
http -v "http://localhost:8000/api/points?x=1&y=1&distance=24"
```
