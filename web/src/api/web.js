import {get, post} from '@/lib/fetch.js'

export function newAPPWebApi(data) {
    return post('/app/web/add', data)
}

export function listAPPWebApi(params) {
    return get('/app/web/list', params)
}

export function deleteAPPWebApi(data) {
    return post('/app/web/delete', data)
}

export function detailAPPWebApi(params) {
    return get('/app/web/detail', params)
}

export function updateAPPWebApi(data) {
    return post('/app/web/update', data)
}
