# slnotify

Send deviation notifications for a station to your phone using [Pushover](https://pushover.net/) and [sl.se](http://sl.se)'s [api](https://www.trafiklab.se/api/sl-realtidsinformation-4).

## Installation

```
go get github.com/frozzare/slnotify
```

## Usage

Edit config file with api keys and settings. The site id flag is required to find deviations for a station.

```
slnotify --config=/path/to/config.json --site-id=1002
```

To run it more than once use cron or something like that.

## License

MIT © [Fredrik Forsmo](https://github.com/frozzare)