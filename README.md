# k6-proxy

A [k6](https://k6.io/) that allow you to dynamic change the proxy settings


## Build

To build a `k6` binary with this extension, first ensure you have the prerequisites:

- [Go toolchain](https://go101.org/article/go-toolchain.html)
- Git

Then:

1. Download the `xk6`:

```bash
go install go.k6.io/xk6/cmd/xk6@latest
```

2. Build your `k6` binnary.

```bash
xk6 build --with github.com/sym01/k6-proxy
```

## Example

```js
import http from 'k6/http';
import proxy from 'k6/x/proxy';

const YOUR_PROXY = 'http://user:passwd@proxy.hostname'

export default function () {
  proxy.setProxy(YOUR_PROXY)

  const resp = http.get('http://httpbin.test.k6.io/get')

  proxy.clearProxy()
}
```

For a full example, can refer to [here](./examples/dynamic-proxy.js)

Example output:

```txt

          /\      |‾‾| /‾‾/   /‾‾/
     /\  /  \     |  |/  /   /  /
    /  \/    \    |     (   /   ‾‾\
   /          \   |  |\  \ |  (‾)  |
  / __________ \  |__| \__\ \_____/ .io

  execution: local
     script: examples/dynamic-proxy.js
     output: -

  scenarios: (100.00%) 1 scenario, 1 max VUs, 10m30s max duration (incl. graceful stop):
           * default: 1 iterations for each of 1 VUs (maxDuration: 10m0s, gracefulStop: 30s)


running (00m01.0s), 1/1 VUs, 0 complete and 0 interrupted iterations
default   [   0% ] 1 VUs  00m01.0s/10m0s  0/1 iters, 1 per VU

running (00m02.0s), 1/1 VUs, 0 complete and 0 interrupted iterations
default   [   0% ] 1 VUs  00m02.0s/10m0s  0/1 iters, 1 per VU

     ✓ Request with proxy
     ✓ Request without proxy

     checks.........................: 100.00% ✓ 2        ✗ 0
     data_received..................: 6.8 kB  2.9 kB/s
     data_sent......................: 741 B   313 B/s
     http_req_blocked...............: avg=529.8ms  min=275.86ms med=313.71ms max=999.84ms p(90)=862.61ms p(95)=931.22ms
     http_req_connecting............: avg=351.06ms min=275.48ms med=309.85ms max=467.85ms p(90)=436.25ms p(95)=452.05ms
     http_req_duration..............: avg=257.99ms min=10.52ms  med=249.71ms max=513.75ms p(90)=460.94ms p(95)=487.34ms
       { expected_response:true }...: avg=381.73ms min=249.71ms med=381.73ms max=513.75ms p(90)=487.34ms p(95)=500.54ms
     http_req_failed................: 33.33%  ✓ 1        ✗ 2
     http_req_receiving.............: avg=60.66µs  min=53µs     med=60µs     max=69µs     p(90)=67.2µs   p(95)=68.1µs
     http_req_sending...............: avg=411.33µs min=27µs     med=91µs     max=1.11ms   p(90)=911µs    p(95)=1.01ms
     http_req_tls_handshaking.......: avg=177.3ms  min=0s       med=0s       max=531.91ms p(90)=425.52ms p(95)=478.71ms
     http_req_waiting...............: avg=257.52ms min=9.33ms   med=249.56ms max=513.66ms p(90)=460.84ms p(95)=487.25ms
     http_reqs......................: 3       1.268011/s
     iteration_duration.............: avg=2.36s    min=2.36s    med=2.36s    max=2.36s    p(90)=2.36s    p(95)=2.36s
     iterations.....................: 1       0.42267/s
     vus............................: 1       min=1      max=1
     vus_max........................: 1       min=1      max=1


running (00m02.4s), 0/1 VUs, 1 complete and 0 interrupted iterations
default ✓ [ 100% ] 1 VUs  00m02.4s/10m0s  1/1 iters, 1 per VU
```