# Proto files for ubrato microservices

## Generate for golang

```bash
make generate.go
```

## Generate for python

```bash
make generate.go
```

## Integration

```bash
git submodule add git submodule add git@git.ubrato.ru:ubrato/proto.git proto proto
echo "include ./proto/proto.mk" >> Makefile
echo "PROTO_PATH = ./proto" >> Makefile
echo "PROTO_OUT = ./some/path" >> Makefile
```

### For golang
```bash
echo "PROTO_OUT_MODULE = git.ubrato.ru/some/path/to/pkg" >>
```
