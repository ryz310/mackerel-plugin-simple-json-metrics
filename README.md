# mackerel-plugin-simple-json-metrics

Simple JSON custom metrics plugin for mackerel.io agent.
You can monitor that via server response JSON your application server real time status.

## How to use

Implements following specification API server.

- Receive header `Authorization: Bearer {api-key}`.
- Receive header `Content-Type: application/json; charset=utf-8`.
- Return status code `200`
- Response with following parameters as JSON.
  - The number of groups is arbitrary.

```json
# Response parameters
{
  "monitoring": [
    { "group": "{group 1}", "count": 123 },
    { "group": "{group 2}", "count": 456 },
    { "group": "{group 3}", "count": 789 }
  ],
  "timestamp": 1513135652
}
```

## Synopsis

```
$ mackerel-plugin-simple-json-metrics \
  -url=<Monitoring API URL> \
  -api-key=<Monitoring API Key>
```

## Example of mackerel-agent.conf

```
[plugin.metrics.simple-json-metrics]
command = "/path/to/mackerel-plugin-simple-json-metrics \
  -url=<Monitoring API URL> \
  -api-key=<Monitoring API Key>"
```
