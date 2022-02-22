# laserver
Simple tracking server for lateralus

# Features
It is used as a tracking server for [lateralusd](https://github.com/lateralusd/lateralus) tool. It just reads the user id from `id` get param and save it to bolt db with the time that the request has ocurred.

It can also be used to correlate exactly who has opened email from lateralus `json` report/

# Installation and running

```bash
$ git clone https://github.com/lateralusd/laserver.git
$ cd laserver/ && go build
$ ./laserver
```

If we make a request in another terminal as `curl 'http://localhost:3300/?id=sampleuuid'` inside the `laserver` we would see following:

```bash
2022/02/22 12:26:59 Got request from[::1]:59764
2022/02/22 12:26:59 sampleuuid saved successfully
```