import {del, get, post, upload} from '@/lib/fetch.js'
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
    return get('/server/group/path', data)
}

export function saveGroupPathApi(data) {
    return post('/server/group/path', data)
}


export function deleteGroupPathApi(data) {
    return del('/server/group/path', data)
}

export function detailGroupPathApi(id) {
    return options('/server/group/path/' + id)
}


export function listGroupConfigApi(params) {
    // return Promise.resolve({
    //     "list": [
    //         {
    //             "id": 3,
    //             "name": "nginx.conf",
    //             "file": "/nginx.conf",
    //             "ctime": "2024-05-13"
    //         },
    //         {
    //             "id": 1,
    //             "name": "php.conf",
    //             "file": "/php.conf",
    //             "ctime": "2024-05-13"
    //         },
    //     ],
    // });
    return get('/server/group/config', params)
}

export function saveGrouConfigpApi(data) {
    return post('/server/group/config/update', data)
}

export function deleteGrouConfigpApi(data) {
    return del('/server/group/config/update', data)
}

export function detailGrouConfigpApi(id) {
    return get('/server/group/config/' + id)
}

export function newServerApi(data) {
    return post('/server/add', data)
}

export function updateServerApi(data) {
    return post('/server/update', data)
}

export function listServerApi(params) {
    return get('/server/list', params)
}

export function listGroupRunApi(params) {
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
    if (!params) {
        params = {limit: 20}
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

export function testServerConnect(data) {
    return get('/server/ssh/test', data)
}

export function serverSession(data) {
    // return post('http://localhost:8899/api/ssh/create_session', {"id":1,"name":"130","address":"192.168.1.130","user":"root","auth_type":"pwd","net_type":"tcp4","cert_data":"","cert_pwd":"","pwd":"Aa123456","port":22,"background":"#000000","foreground":"#FFFFFF","cursor_color":"#FFFFFF","font_family":"Courier","font_size":16,"cursor_style":"block","shell":"bash","pty_type":"xterm-256color","init_cmd":"","init_banner":""})
    return get('/server/session', data)
}

export function serverSessionResize(data) {
    return get('/server/session/resize', data)
}

export function listFtpApi(data) {
    return get('/server/sftp', data)
}


export function deleteFtpApi(data) {
    return get('/server/sftp/delete', data)
}

export function createDirFtpApi(data) {
    return get('/server/sftp/create', data)
}

export function renameFtpApi(data) {
    return get('/server/sftp/rename', data)
}

export function modFtpApi(data) {
    return get('/server/sftp/mod', data)
}

export function downFtpApi(data) {
    return axios.get('/api/server/sftp/down', {
        params: data,
        headers: {
            "Authorization": util.GetLoginToken()
        }
    })

}

export function uploadFtpApi(data, onUploadProgress) {
    return upload('/server/sftp/upload', data, {}, {}, onUploadProgress)
}

export function zipFtpApi(data) {
    return post('/server/sftp/zip', data)
}

export function unzipFtpApi(data) {
    return post('/server/sftp/unzip', data)
}


export function listCmdApi(data) {
    return get('/cmd/list', data)
}

export function deleteCmdApi(data) {
    return post('/cmd/delete', data)
}


export function createCmdApi(data) {
    return post('/cmd/add', data)
}

export function updateCmdApi(data) {
    return post('/cmd/update', data)
}



