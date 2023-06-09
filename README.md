# daphne

Daphne（ダフニー）沈丁花（月桂）

## description

high performance gin project

## DESC

1. a local map maintaining topic -> urls relations
2. a working pool for producing messages to MQ
   1. producer is not topic-combined
   2. init a bunch go-routines
3. two kinds of consumption
   1. push consumer, maintain a list of topic
   2. pull consumer, create consumer per topic
