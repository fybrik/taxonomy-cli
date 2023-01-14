# taxonomy-cli

A taxonomy defines the terms and related values that need to be commonly understood and supported across the components in the system.

`taxonomy-cli` can compile taxonomy files into a json schema, and validate an existant json schema.

For more information, please visit [the Fybrik website](https://fybrik.io/v1.2/concepts/taxonomy/).


### Version compatibility matrix

| Fybrik           | taxonomy-cli |
| ---              | ---          |
| master           | main         |


## Usage

Simply run the `taxonomy-cli` command on the provided docker image.

For example, running `help` command:

```bash
docker run --rm -u "$(id -u):$(id -g)" ghcr.io/fybrik/taxonomy-cli:main help
```

Or compiling a `./base.json` file with a `./layer.yaml` into a `./taxonomy.json` output:

```bash
docker run --rm -u "$(id -u):$(id -g)" --volume ${PWD}:/local --workdir /local/ \
        ghcr.io/fybrik/taxonomy-cli:main compile \
	-o ./taxonomy.json \
  	-b ./base.json \
        ./layer.yaml
```