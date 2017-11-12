## serverChecker

Checks each X seconds/minutes/hours, if the configured websites/servers, are responding or are down.

server.json:

```json
[
    {
        "url": "https://github.com"
    },
    {
        "url": "https://kaggle.com"
    },
    {
        "url": "http://fakeurlnotworking.com"
    }
]
```


ToDo:
- Some stats about the uptime and downtime
- Run in background and show desktop popups when some server goes down