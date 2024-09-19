import md5 from 'blueimp-md5'
import Cookies from 'js-cookie'
import moment from 'moment'
import Vue from 'vue';

let loginTokenKey = '_zoom_identity'
let languageCookieKey = '_zoom_language'
const EventBus = new Vue();


export default {
    EmitEventGlobal(messageName, data){
        EventBus.$emit(messageName, data); 
    },
    BindEventGlobal(messageName, callback){
        EventBus.$on(messageName, callback);
    },
    UnBindEventGlobal(messageName){
        EventBus.$off(messageName);
    },
    MessageSuccess(cb){
        this.$message({
            message: this.$t('operate_success'),
            type: 'success',
            duration: 1200,
            onClose: cb,
        });
    },

    PageInit() {
        this.$root.PageSize = 7
        this.$root.Page = 1
        this.$root.Total = 0
    },

    PageReset() {
        this.$root.Total--
        let maxPage = Math.ceil(this.$root.Total / this.$root.PageSize)
        if (this.$root.Page > maxPage) {
            this.$root.Page = maxPage
        }
        if (this.$root.Page < 1) {
            this.$root.Page = 1
        }
    },

    PageOffset() {
        return this.$root.PageSize * (this.$root.Page - 1)
    },

    ConfirmDelete(cb, title) {
        if (!title) {
            title = '此操作将永久删除该数据, 是否继续?'
        }
        this.$confirm(title, '提示', {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'warning'
        }).then(() => {
            cb()
        }).catch((err) => {
            console.log(err)
        })
    },

    ConfirmCopy(cb, title) {
        if (!title) {
            title = '此操作将复制一个新项目，是否继续?'
        }
        this.$confirm(title, '提示', {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'info'
        }).then(() => {
            cb()
        }).catch((err) => {
            console.log(err)
        })
    },

    Md5Sum(str) {
        return md5(str);
    },

    FormatDateTime(unixtime, format) {
        if (!unixtime) {
            return '--'
        }
        if (!format) {
            format = 'YYYY-MM-DD HH:mm:ss'
        }
        return moment.unix(unixtime).format(format)
    },

    FormatDateDuration(num) {
        return moment.duration(num).humanize(false)
    },

    FormatDateFromNow(unixtime) {
        if (!unixtime) {
            return '--'
        }
        return moment.unix(unixtime).fromNow()
    },

    Substr(str, len) {
        if (Object.prototype.toString.call(str) != '[object String]') {
            return ''
        }
        let postfix = ''
        if (str.length > len) {
            postfix = "..."
        }
        return str.substr(0, len) + postfix
    },

    SetLanguage(language) {
        return Cookies.set(languageCookieKey, language)
    },

    GetLanguage() {
        return Cookies.get(languageCookieKey)
    },
    switchLanguage() {
        let localeLang
        if (this.GetLanguage() != 'en') {
            localeLang = "en"
        } else {
            localeLang = "zh-cn"
        }
        this.SetLanguage(localeLang)
        location.reload()
    },

    SetLoginToken(token) {
        return Cookies.set(loginTokenKey, token)
    },

    GetLoginToken() {
        return Cookies.get(loginTokenKey)
    },

    DeleteLoginToken() {
        return Cookies.remove(loginTokenKey)
    },

    CheckPriv(code) {
        return this.$store.getters['account/getUserId'] ==1 || this.$store.getters['account/getPriv'].indexOf(code) > -1
    },

    CheckPrivs(codeArr) {
        if (!codeArr || !codeArr.length) {
            return false
        }
        let checked = false
        codeArr.forEach(code => {
            if (this.CheckPriv(code)) {
                checked = true
            }
        })
        return checked
    },
    PermissionToOctal(permission) {
        // 文件权限 "-rw-r--r--" 转换为八进制数字 "0611"
        const permissionMap = {
          'r': 4,
          'w': 2,
          'x': 1,
          '-': 0
        };
      
       let s1 = permissionMap[permission[0]];
       let s2 = permissionMap[permission[1]] + permissionMap[permission[2]] + permissionMap[permission[3]];
       let s3 = permissionMap[permission[4]] + permissionMap[permission[5]] + permissionMap[permission[6]];
       let s4 = permissionMap[permission[7]] + permissionMap[permission[8]] + permissionMap[permission[9]];
        return s1.toString(8) + s2 + s3 + s4 ;
    },

    determineFileType: function (ext, type) {
        /**
         * @description 文件类型判断，或返回格式类型(不传入type)
         * @param {String} ext
         * @param {String} type
         * @return {Boolean|Object} 返回类型或类型是否支持
         */
        var config = {
            images: ['jpg', 'jpeg', 'png', 'bmp', 'gif', 'tiff', 'ico', 'JPG', 'webp'],
            compress: ['zip', 'rar', 'gz', 'war', 'tgz', 'tar', '7z'],
            video: ['mp4', 'mp3', 'mpeg', 'mpg', 'mov', 'avi', 'webm', 'mkv', 'mkv', 'mp3', 'rmvb', 'wma', 'wmv'],
            ont_text: ['iso', 'xlsx', 'xls', 'doc', 'docx', 'tiff', 'exe', 'so', 'bz', 'dmg', 'apk', 'pptx', 'ppt', 'xlsb', 'pdf'],
        },
        returnVal = false;
        if (type != undefined) {
            if (type == 'text') {
                config.forEach(item=>{
                    item.forEach(items=>{
                        if (items == ext) {
                            returnVal = true;
                            return false;
                        }
                    })
                })

                // $.each(config, function (key, item) {
                //     $.each(item, function (index, items) {
                //         if (items == ext) {
                //             returnVal = true;
                //             return false;
                //         }
                //     });
                // });
                returnVal = !returnVal;
            } else {
                if (typeof config[type] == 'undefined') return false;
                config[type].forEach(item=>{
                    if (item == ext) {
                        returnVal = true;
                        return false;
                    }
                })
            }


            //     $.each(config[type], function (key, item) {
            //         if (item == ext) {
            //             returnVal = true;
            //             return false;
            //         }
            //     });
            // }

        } else {
            config.forEach(item=>{
                item.forEach(items=>{
                    if (items == ext) {
                        returnVal = true;
                        return false;
                    }
                })
            })

            // $.each(config, function (key, item) {
            //     $.each(item, function (index, items) {
            //     if (items == ext) {
            //         returnVal = key;
            //         return false;
            //     }
            //     });
            // });
            if (typeof returnVal == 'boolean') returnVal = 'text';
        }
        return returnVal;
    }
    
      
}