### Kafka Commands

Creating new topic
``` bash
kafka-topics --bootstrap-server=localhost:9092 --topic courses --create --partitions=3 --replication-factor=1
```

Start a producer
``` bash
kafka-console-producer --bootstrap-server=localhost:9092 --topic=courses 
```