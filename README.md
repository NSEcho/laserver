# laserver
Simple tracking server for lateralus

# Features
It is used as a tracking server for [lateralusd](https://github.com/lateralusd/lateralus) tool. It just reads the user id from `id` get param and save it to bolt db with the time that the request has ocurred.

It can also be used to correlate exactly who has opened email from lateralus `json` report/
