# Cloudgo-start

一个简单的go语言http服务器示例，采用negroni + mux的框架组合。

## 框架选择

开发http后端可用的框架非常多，从轻量级的go语言自带`net/http`包到重量级的`beego`，琳琅满目。我在这里选择的是介于轻量级和重量级之间的negroni + mux的框架组合，它很好地取了两者的优点，规避了两者的缺点。

go语言自带包小巧、速度快、性能高，但是要实现复杂的功能开发量十分巨大；重量级框架功能丰富，但是性能较低，也难以扩展和自定义，因此选择negroni + mux，原始negroni十分接近`net/http`包，小巧高效，同时negroni又很好地支持丰富的扩展，这些扩展由社区维护，具有很强的可自定义性，使得我们能够高效的开发出各种想要的复杂功能。

## Curl 测试

采用curl对此服务器进行测试

    $ curl -v http://localhost:8080/hello/world
    *   Trying ::1...
    * Connected to localhost (::1) port 8080 (#0)
    > GET /hello/world HTTP/1.1
    > Host: localhost:8080
    > User-Agent: curl/7.47.0
    > Accept: */*
    >
    < HTTP/1.1 200 OK
    < Content-Type: application/json; charset=UTF-8
    < Date: Sun, 10 Dec 2017 03:14:41 GMT
    < Content-Length: 23
    <
    {
      "hello": "world"
    }
    * Connection #0 to host localhost left intact

    $ curl -v http://localhost:8080/hello/cloudgo
    *   Trying ::1...
    * Connected to localhost (::1) port 8080 (#0)
    > GET /hello/cloudgo HTTP/1.1
    > Host: localhost:8080
    > User-Agent: curl/7.47.0
    > Accept: */*
    >
    < HTTP/1.1 200 OK
    < Content-Type: application/json; charset=UTF-8
    < Date: Sun, 10 Dec 2017 03:15:22 GMT
    < Content-Length: 25
    <
    {
      "hello": "cloudgo"
    }
    * Connection #0 to host localhost left intact

## ab 测试

采用Apache的ab测试对服务器进行压力测试

    $ ab -n 10000 -c 1000 http://localhost:8080/hello/world
    This is ApacheBench, Version 2.3 <$Revision: 1706008 $>
    Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
    Licensed to The Apache Software Foundation, http://www.apache.org/

    Benchmarking localhost (be patient)
    Completed 1000 requests
    Completed 2000 requests
    Completed 3000 requests
    Completed 4000 requests
    Completed 5000 requests
    Completed 6000 requests
    Completed 7000 requests
    Completed 8000 requests
    Completed 9000 requests
    Completed 10000 requests
    Finished 10000 requests


    Server Software:
    Server Hostname:        localhost
    Server Port:            8080

    Document Path:          /hello/world
    Document Length:        23 bytes

    Concurrency Level:      1000
    Time taken for tests:   0.495 seconds
    Complete requests:      10000
    Failed requests:        0
    Total transferred:      1460000 bytes
    HTML transferred:       230000 bytes
    Requests per second:    20220.03 [#/sec] (mean)
    Time per request:       49.456 [ms] (mean)
    Time per request:       0.049 [ms] (mean, across all concurrent requests)
    Transfer rate:          2882.93 [Kbytes/sec] received

    Connection Times (ms)
                  min  mean[+/-sd] median   max
    Connect:        0    2   1.6      2      14
    Processing:     0    5   3.5      3      38
    Waiting:        0    4   3.5      3      37
    Total:          0    7   4.2      5      49

    Percentage of the requests served within a certain time (ms)
      50%      5
      66%      6
      75%      7
      80%      8
      90%     11
      95%     13
      98%     21
      99%     27
     100%     49 (longest request)

参数解释

- `-n` 请求数量
- `-c` 并发数量

可以看到，速度还是比较客观的
