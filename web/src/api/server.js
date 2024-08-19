import {get, post, upload} from '@/lib/fetch.js'
import axios from "axios";
import util from '@/lib/util.js'
export function newGroupApi(data) {
    return post('/server/group/add', data)
}

export function listGroupApi(params) {
    return get('/server/group/list', params)
}

export function deleteGroupApi(data) {
    return post('/server/group/delete', data)
}

export function detailGroupApi(params) {
    return get('/server/group/detail', params)
}

export function updateGroupApi(data) {
    return post('/server/group/update', data)
}

export function listGroupPathApi(data) {
    console.log('/server/group/path')
    return Promise.resolve({
        "list": [
            {
                "id": 2,
                "group_id": 2,

                "server_id":0,
                "name": "全局",
                "path":"",
                "ctime": 1713167856
            },
            {
                "id": 3,
                "group_id": 2,
      
                "server_id":1,
                "name": "92",
                "path":"/",
                "ctime": 1713167856
            }
        ],
        "total": 3
    });
    return get('/server/group/path', data)
}

export function appGroupPathApi(data) {
    return post('/server/group/path/add', data)
}

export function updateGroupPathApi(data) {
    return post('/server/group/path/update', data)
}

export function deleteGroupPathApi(data) {
    return post('/server/group/path/delete', data)
}

export function listGroupConfigApi(params) {
    return Promise.resolve({
        "list": [
            {
                "id": 3,
                "name": "nginx.conf",
                "file": "/nginx.conf",
                "ctime": "2024-05-13"
            },
            {
                "id": 1,
                "name": "php.conf",
                "file": "/php.conf",
                "ctime": "2024-05-13"
            },
        ],
    });
    return get('/server/group/config/list', params)
}

export function updateGrouConfigpApi(data) {
    return post('/server/group/config/update', data)
}


export function newServerApi(data) {
    return post('/server/add', data)
}

export function updateServerApi(data) {
    return post('/server/update', data)
}

export function listServerApi(params) {
    return Promise.resolve({
        "list": [
            {
                "id": 3,
                "group_id": 0,
                "group_name": "",
                "name": "101",
                "ip": "192.168.1.101",
                "ssh_port": 22,
                "ctime": 1713167856
            },
            {
                "id": 2,
                "group_id": 4,
                "group_name": "web-nginx-test",
                "name": "130",
                "ip": "192.168.1.130",
                "ssh_port": 22,
                "ctime": 1713162409
            },
            {
                "id": 1,
                "group_id": 3,
                "group_name": "api-nginx-test",
                "name": "102",
                "ip": "192.167.1.102",
                "ssh_port": 22,
                "ctime": 1706180377
            }
        ],
        "total": 3
    })
    return get('/server/list', params)
}

export function listGroupRunApi(params){
    return Promise.resolve({
        "list": [
            {
                "id": 3,
                "name": "nginx.conf",
                "status": "1",
                "ctime": "2024-05-13"
            },
            {
                "id": 1,
                "name": "php.conf",
                "status": "2",
                "ctime": "2024-05-13"
            },
        ],
    });
    return get('/server/group/config/list', params)
}

export function deleteServerApi(data) {
    return post('/server/delete', data)
}

export function detailServerApi(params) {
    return get('/server/detail', params)
}

export function newSshkeyApi(data) {
    return post('/sshkey/add', data)
}

export function updateSshkeyApi(data) {
    return post('/sshkey/update', data)
}

export function listSshkeyApi(params) {
    if(!params){
        params = {limit:20}
    }
    return get('/sshkey/list', params)
}

export function deleteSshkeyApi(data) {
    return post('/sshkey/delete', data)
}

export function detailSshkeyApi(params) {
    return get('/sshkey/detail', params)
}

export function newServerSshkeyApi(data) {
    return post('/server/sshkey/add', data)
}

export function updateServerSshkeyApi(data) {
    return post('/server/sshkey/update', data)
}

export function deleteServerSshkeyApi(data) {
    return post('/server/sshkey/delete', data)
}

export function detailServerSshkeyApi(data) {
    return get('/server/sshkey/detail', data)
}

export function testServerConnect(data){
    return get('/server/ssh/test', data)
}

export function serverSession(data){
    // return post('http://localhost:8899/api/ssh/create_session', {"id":1,"name":"130","address":"192.168.1.130","user":"root","auth_type":"pwd","net_type":"tcp4","cert_data":"","cert_pwd":"","pwd":"Aa123456","port":22,"background":"#000000","foreground":"#FFFFFF","cursor_color":"#FFFFFF","font_family":"Courier","font_size":16,"cursor_style":"block","shell":"bash","pty_type":"xterm-256color","init_cmd":"","init_banner":""})
    return get('/server/session', data)
}

export function serverSessionResize(data){
    return get('/server/session/resize', data)
}

export function listFtpApi(data){
    return get('/server/sftp', data)
}


export function deleteFtpApi(data){
    return get('/server/sftp/delete', data)
}

export function createDirFtpApi(data){
    return get('/server/sftp/create', data)
}

export function renameFtpApi(data){
    return get('/server/sftp/rename', data)
}

export function modFtpApi(data){
    return get('/server/sftp/mod', data)
}

export function downFtpApi(data){
    return axios.get('/api/server/sftp/down', {
        params: data,
        headers:{
            "Authorization": util.GetLoginToken()
        }
    }) 
    
}

export function uploadFtpApi(data, onUploadProgress){
    return upload('/server/sftp/upload', data, {}, {}, onUploadProgress) 
}

export function zipFtpApi(data){
    return post('/server/sftp/zip', data) 
}

export function unzipFtpApi(data){
    return post('/server/sftp/unzip', data) 
}

