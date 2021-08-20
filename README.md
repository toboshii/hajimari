<img src="https://raw.githubusercontent.com/toboshii/hajimari/main/docs/static/img/logo.png" align="left" height="144px"/>

# Hajimari :sunrise:
*...The beginning of a pleasant experience*

<br />
<br />

![Hajimari](https://raw.githubusercontent.com/toboshii/hajimari/main/docs/static/img/screen01.png)

## Installation

### Helm

`helm repo add hajimari https://hajimari.io`

`helm repo update`

`helm install hajimari hajimari/hajimari`

![Helm docs](charts/hajimari)

### Locally

Clone the repo and run the following command to generate the `hajimari` binary:

```bash
make build
```

You will need to have `go` installed.

#### Usage

Copy the binary and edit `config.yaml`. Then run:

```bash
./hajimari
```

Please note there is no authentication. You might want to run this behind a web server with reverse proxy capabilities.

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

Run `make help` for information on linting, tests, etc.

## Thank you / dependencies

- [SUI](https://github.com/jeroenpardon/sui) For the great startpage template
- [Forecastle](https://github.com/stakater/Forecastle) Ideas for integrating k8s ingress

## License
[Apache-2.0](https://choosealicense.com/licenses/apache-2.0/)
