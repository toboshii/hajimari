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

## Usage

### Ingresses

Hajimari looks for specific annotations on ingresses.

- Add the following annotations to your ingresses in order for it to be discovered by Hajimari:

| Annotation                                   | Description                                                                                                                                                 | Required |
| -------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------- | -------- |
| `hajimari.io/enable`             | Add this with value `true` to the ingress of the app you want to show in Hajimari                                                                                       | `true`   |
| `hajimari.io/icon`               | Icon name from ![MDI icons](https://materialdesignicons.com/)                                                                                                           | `false`  |
| `hajimari.io/appName`            | A custom name for your application. Use if you don't want to use the name of the ingress                                                                                | `false`  |
| `hajimari.io/group`              | A custom group name. Use if you want the application to show in a different group than the namespace it is running in                                                   | `false`  |
| `hajimari.io/instance`           | A comma separated list of name/s of the Hajimari instance/s where you want this application to appear. Use when you have multiple Hajimari instances                    | `false`  |
| `hajimari.io/url`                | A URL for the Hajimari app (This will override the ingress URL). It MUST begin with a scheme i.e., `http://` or `https://`                                              | `false`  |

### Config

Hajimari supports the following configuration options that can be modified by either ConfigMap or `values.yaml` if you are using helm

|       Field       |                                                Description                                                 |         Default         | Type              |
| :---------------: | :--------------------------------------------------------------------------------------------------------: | :---------------------: | ----------------- |
| namespaceSelector | Namespace selector which uses a combination of hardcoded namespaces as well as label selectors             |        any: true        | NamespaceSelector |
|       title       |                                     Title for the Hajimari instance                                        |        "Hajimari"       | string            |
|   instanceName    |                                      Name of the Hajimari instance                                         |           ""            | string            |
|    customApps     |                A list of custom apps that you would like to add to the Hajimari instance                   |           {}            | []CustomApp       |
|    groups         |                A list of bookmark groups to add to the Hajimari instance                                   |           {}            | []groups          |

#### NamespaceSelector

It is a selector for selecting namespaces either selecting all namespaces or a list of namespaces, or filtering namespaces through labels.

|     Field     |                                          Description                                          | Default | Type                                                                                         |
| :-----------: | :-------------------------------------------------------------------------------------------: | :-----: | -------------------------------------------------------------------------------------------- |
|      any      | Boolean describing whether all namespaces are selected in contrast to a list restricting them |  false  | bool                                                                                         |
| labelSelector |                Filter namespaces based on kubernetes metav1.LabelSelector type                |  null   | [metav1.LabelSelector](https://godoc.org/k8s.io/apimachinery/pkg/apis/meta/v1#LabelSelector) |
|  matchNames   |                                    List of namespace names                                    |  null   | []string                                                                                     |

*Note:* If you specify both `labelSelector` and `matchNames`, Hajimari will take a union of all namespaces matched and use them.

#### Custom Apps

If you want to add any apps that are not exposed through ingresses or are external to the cluster, you can use the custom apps feature. You can pass an array of custom apps inside the config.

| Field             | Description                               | Type              |
| ----------------- | ----------------------------------------- | ----------------- |
| name              | Name of the custom app                    | String            |
| icon              | URL of the icon for the custom app        | String            |
| url               | URL of the custom app                     | String            |
| group             | Group for the custom app                  | String            |

#### Bookmarks

Bookmark groups can be added by creating an array of group names and links.

| Field            | Description                               | Type              |
| -----------------| ----------------------------------------- | ----------------- |
| name             | Name of the bookmark group                | String            |
| links            | Array of links                            | Array             |

Bookmarks can be added by configuring a list of bookmarks under a group.

| Field             | Description                               | Type              |
| ----------------- | ----------------------------------------- | ----------------- |
| name              | Name of the bookmark                      | String            |
| url               | URL of the bookmark                       | String            |

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
