import {get, post} from '@/lib/fetch.js'

export function newAPPRedisApi(data) {
    return post('/app/redis/add', data)
}

export function listAPPRedisApi(params) {
    return get('/app/redis/list', params)
}

export function deleteAPPRedisApi(data) {
    return post('/app/redis/delete', data)
}

export function detailAPPRedisApi(params) {
    return get('/app/redis/detail', params)
}

export function updateAPPRedisApi(data) {
    return post('/app/redis/update', data)
}
