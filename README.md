# slnotify

Send notifications for a station to your phone using [Pushover](https://pushover.net/) and [sl.se](http://sl.se)'s [api](https://www.trafiklab.se/api/sl-realtidsinformation-3).

## Installation

```
go get github.com/frozzare/slnotify
```

## Usage

Edit config file with api keys and settings.

```
slnotify --config=/path/to/config.json --site-id=1002
```

The site id flag is required to find deviations for a station.

## License

MIT © [Fredrik Forsmo](https://github.com/frozzare)