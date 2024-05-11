import {get, post} from '@/lib/fetch.js'

export function newAPPPhpApi(data) {
    return post('/app/php/add', data)
}

export function listAPPPhpApi(params) {
    return get('/app/php/list', params)
}

export function deleteAPPPhpApi(data) {
    return post('/app/php/delete', data)
}

export function detailAPPPhpApi(params) {
    return get('/app/php/detail', params)
}

export function updateAPPPhpApi(data) {
    return post('/app/php/update', data)
}
