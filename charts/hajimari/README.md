# hajimari

![Version: 1.2.0](https://img.shields.io/badge/Version-1.2.0-informational?style=flat-square) ![AppVersion: v0.2.0](https://img.shields.io/badge/AppVersion-v0.2.0-informational?style=flat-square)

Hajimari is a beautiful & customizable browser startpage/dashboard with
Kubernetes application discovery

**Homepage:** <https://github.com/toboshii/hajimari/tree/master/charts/hajimari>

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| toboshii | toboshii@users.noreply.github.com |  |

## Example values to get started:

```yaml
hajimari:
  namespaceSelector:
    matchNames:
    - downloads
    - media
  customApps:
  - name: Test
    url: https://example.com
    icon: test-tube
  groups:
  - name: Communicate
    links:
    - name: Discord
      url: 'https://discord.com'
    - name: Gmail
      url: 'http://gmail.com'
    - name: Slack
      url: 'https://slack.com/signin'
  - name: Cloud
    links:
    - name: Box
      url: 'https://box.com'
    - name: Dropbox
      url: 'https://dropbox.com'
    - name: Drive
      url: 'https://drive.google.com'
  - name: Design
    links:
    - name: Awwwards
      url: 'https://awwwards.com'
    - name: Dribbble
      url: 'https://dribbble.com'
    - name: Muz.li
      url: 'https://medium.muz.li/'
  - name: Dev
    links:
    - name: Codepen
      url: 'https://codepen.io/'
    - name: Devdocs
      url: 'https://devdocs.io'
    - name: Devhints
      url: 'https://devhints.io'
  - name: Lifestyle
    links:
    - name: Design Milk
      url: 'https://design-milk.com/category/interior-design/'
    - name: Dwell
      url: 'https://www.dwell.com/'
    - name: Freshome
      url: 'https://www.mymove.com/freshome/'
  - name: Media
    links:
    - name: Spotify
      url: 'http://browse.spotify.com'
    - name: Trakt
      url: 'http://trakt.tv'
    - name: YouTube
      url: 'https://youtube.com/feed/subscriptions'
  - name: Reading
    links:
    - name: Instapaper
      url: 'https://www.instapaper.com/u'
    - name: Medium
      url: 'http://medium.com'
    - name: Reddit
      url: 'http://reddit.com'
  - name: Tech
    links:
    - name: TheNextWeb
      url: 'https://thenextweb.com/'
    - name: The Verge
      url: 'https://theverge.com/'
    - name: MIT Technology Review
      url: 'https://www.technologyreview.com/'
ingress:
  main:
    enabled: true
    hosts:
      - host: hajimari.domain.tld
        paths:
          - path: /
            pathType: Prefix
persistence:
  data:
    enabled: true
    accessMode: ReadWriteOnce
    size: 1Gi
```

## Source Code

* <https://github.com/toboshii/hajimari>

## Requirements

Kubernetes: `>=1.16.0-0`

| Repository | Name | Version |
|------------|------|---------|
| https://library-charts.k8s-at-home.com | common | 4.0.0 |

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| env | object | See below | environment variables. |
| env.TZ | string | `"UTC"` | Set the container timezone |
| hajimari | object | See below | Configures Hajimari settings for this instance. |
| hajimari.customApps | list | `[]` | Add custom applications to the discovered application list |
| hajimari.defaultEnable | bool | `false` | Set to true to show all discovered applications by default. |
| hajimari.groups | list | `[]` | Set default bookmarks |
| hajimari.instanceName | string | `nil` | The name of this instance, this allows running multiple  instances of Hajimari on the same cluster |
| hajimari.name | string | `"You"` | Default name for welcome message |
| hajimari.namespaceSelector | object | `{"matchNames":["media"]}` | Namespace selector to use for discovering applications |
| hajimari.title | string | `nil` | Override the title of the Hajimari pages |
| image.pullPolicy | string | `"IfNotPresent"` | image pull policy |
| image.repository | string | `"ghcr.io/toboshii/hajimari"` | image repository |
| image.tag | string | `"v0.1.0"` | image tag |
| ingress.main | object | See values.yaml | Enable and configure ingress settings for the chart under this key. |
| persistence | object | See values.yaml | Configure persistence settings for the chart under this key. |
| service | object | See values.yaml | Configures service settings for the chart. |
| serviceAccount | object | See below | Configures service account needed for reading k8s ingress objects |
| serviceAccount.create | bool | `true` | Create service account |
