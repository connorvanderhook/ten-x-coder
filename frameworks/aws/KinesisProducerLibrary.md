## KPL - Kinesis Producer Library

https://docs.aws.amazon.com/streams/latest/dev/developing-producers-with-kpl.html

Acts as an intermediary between your application and the Kinesis Data Streams API.

KPL simplifes producer application development
and also building batch of aggregation of user records by increasing payload size and improve
throughput and optimize costs.

Validate the transaction by checking the successful insert into the stream by embedding automatic and confgurable retry mechanism

Kinesis Streams provide capabilities to use Future objects to validate UserRecords. No need to complicate the code by storing in memory/transient storage.

Time-to-live records need to be increases if the UserRecords could not inserted into stream in time.

Streaming data can directy be delivered into ElasticSearch Domain.

The transformation failures and delivery failures are loaded into processing-failed and errors folders in same S3 bucket.


**KPL User Record** - KPL user record is a blob of data that has meaning to the user. (ex: UI event on a website or a log entry from a web server)

`KPL user record != Kinesis Data stream record`