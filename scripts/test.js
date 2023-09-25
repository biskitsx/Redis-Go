import http from 'k6/http'
import { check, sleep } from 'k6'

export let options = {
    vus: 5,
    duration: '5s'
}

export default function () {
    let res = http.get('http://host.docker.internal:8080/products')

    // check(res, { 'success login': (r) => r.status === 200 })

    // sleep(0.3)
}
