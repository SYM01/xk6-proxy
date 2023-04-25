import { check } from 'k6'
import http from 'k6/http';
import proxy from 'k6/x/proxy';

const mockProxy = 'http://user:passwd@1.1.1.1'

export default function () {
  console.log(proxy.setProxy(mockProxy))
  const resp1 = http.get('http://httpbin.test.k6.io/get')
  check(resp1, {
    'Request with proxy': (resp) => resp.remote_ip == '1.1.1.1',
  })


  proxy.clearProxy()
  const resp2 = http.get('http://httpbin.test.k6.io/get')
  check(resp2, {
    'Request without proxy': (resp) => resp.remote_ip != '1.1.1.1',
  })
}
