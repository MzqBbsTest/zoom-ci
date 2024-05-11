import {get, post} from '@/lib/fetch.js'

export function newAPPNginxApi(data) {
    return post('/app/nginx/add', data)
}

export function listAPPNginxApi(params) {
    return get('/app/nginx/list', params)
}

export function deleteAPPNginxApi(data) {
    return post('/app/nginx/delete', data)
}

export function detailAPPNginxApi(params) {
    return get('/app/nginx/detail', params)
}

export function updateAPPNginxApi(data) {
    return post('/app/nginx/update', data)
}
