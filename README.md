# taxonomy-cli

A taxonomy defines the terms and related values that need to be commonly understood and supported across the components in the system.

`taxonomy-cli` can compile taxonomy files into a json shema, and validate an existant json schema.

For more information, plaese visit [the Fybrik website](https://fybrik.io/v1.2/concepts/taxonomy/).


### Version compatibility matrix

| Fybrik           | taxonomy-cli |
| ---              | ---          |
| master           | master       |


## Usage

Simply run the `taxonomy-cli` command on the provided docker image.

For example, running `help` command:

```bash
docker run --rm -u "$(id -u):$(id -g)" ghcr.io/fybrik/taxonomy-cli:master help
```

Or compiling a `./base.json` file with a `./layer.yaml` into a `./taxonomy.json` output:

```bash
docker run --rm -u "$(id -u):$(id -g)" -v ${PWD}:/local \
        ghcr.io/fybrik/taxonomy-cli:master compile \
	-o /local/taxonomy.json \
  	-b /local/base.json \
        /local/layer.yaml
```