# Kafka and Go - Overview with Hands on Examples

## Install:

**- Install `librdkafka-devel` version `v1.1.0 or later`.**

To connect to Kafka from Go, we need to use some go-client.  There are many go-client available and we are using `confluent-kafka-go`.

Confluent go client is wrapper over `librdkafka` client which is written in C language.  So we need to install `librdkafka` to use `confluent-kafka-go`.

Please see troubleshoot section in case of issue faced to install librdkafka.

## Setup:

**- Start Kafka server / broker and Zookeeper on Docker.** 

```
# start kafka and zookeeper
$ docker-compose up -d

# list kafka and zookeeper
$ docker ps --format "table {{.ID}}\t{{.Names}}\t{{.Ports}}\t{{.Command}}"
CONTAINER ID        NAMES                  PORTS                                                COMMAND
b59aa656d5f1        go-kafka_zookeeper_1   22/tcp, 2888/tcp, 3888/tcp, 0.0.0.0:2181->2181/tcp   "/bin/sh -c '/usr/..."
80965bd9931c        go-kafka_kafka_1       0.0.0.0:9092->9092/tcp                               "start-kafka.sh"
```

Note, `KAFKA_ADVERTISED_HOST_NAME` used in docker-compose.yaml is docker host ip.  To find docker host ip run `ip a | grep docker0`

## Kafka commands:

- list topics
```
kafka-topics.sh --list --zookeeper zookeeper:2181
```

- create topic
```
kafka-topics.sh --create --zookeeper zookeeper:2181 --replication-factor 1 --partitions 1 --topic test-topic
```

- start console producer
```
kafka-console-producer.sh --broker-list localhost:9092 --topic test-topic
```

- start console consumer
```
kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic test-topic --from-beginning
```

## Troubleshoot:

**Problem1:** `rdkafka` package related error while running go program which uses `github.com/confluentinc/confluent-kafka-go/kafka` package.  Below is the sort of error:
```
$ go run main.go 
# pkg-config --cflags  -- rdkafka
Package rdkafka was not found in the pkg-config search path.
Perhaps you should add the directory containing `rdkafka.pc'
to the PKG_CONFIG_PATH environment variable
Package 'rdkafka', required by 'virtual:world', not found
pkg-config: exit status 1
```

**Solution1:** Install `librdkafka-devel` version `v1.1.0 or later`, below are the useful links, installation steps for linux.

- Issue dicussed on github and solution provided: https://github.com/confluentinc/confluent-kafka-go/issues/224#issuecomment-538889211

- Instruction for adding repo: https://docs.confluent.io/current/installation/installing_cp/rhel-centos.html

- Install `sudo dnf install librdkafka-devel`, should install below packages.

```
=============================================================================
 Package                      Architecture    Version                        
=============================================================================
Installing:
 librdkafka-devel             x86_64          1.1.0_confluent5.3.1-1.el7     
Installing dependencies:
 librdkafka1                  x86_64          1.1.0_confluent5.3.1-1.el7     

Transaction Summary
=============================================================================
Install  2 Packages
```

## References:

- Kafka Tutorial from Leaning Journal:
https://www.youtube.com/playlist?list=PLkz1SCf5iB4enAR00Z46JwY9GGkaS2NON
