This directory contains makelets for a bunch of different tools.
Generally speaking, anyone using **go-make** will include
`*.mk` from this directory. However, it is possible for users
to disable certain tools. For example, to disable support for **goa**,
your makefile should include the following line before including
**go-make**:

```
GOMAKE_DISABLE_goa.mk=1
```
