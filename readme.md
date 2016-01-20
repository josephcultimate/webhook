# webhook-api

----
## Subscriptions
Post JSON data: https://webhook-api.com/api/subscription

* app: the name of the application subscribing
* eventtype: namespace.nameofevent
	* namespace.* - for all events in namespace
	* \* - for all events
* url: url to post data to

```sample
{
"app":"hrtech",
"eventtype":"spark.answer",
"url": "https:hrtech.ultilabs.xyz"
}
```

----
## Event Publishing
Post JSON data: https://webhook-api.com/api/event

* type: event name - namespace.nameofevent
* payload: string of data. Could be json object unstructured

```sample
{
"type":"spark.answer",
"payload":"{id: 'some uuid', quiz: 'My Quiz', question: 2, answer: 3}"
}
```
