(function(e){function t(t){for(var n,r,c=t[0],i=t[1],u=t[2],l=0,_=[];l<c.length;l++)r=c[l],o[r]&&_.push(o[r][0]),o[r]=0;for(n in i)Object.prototype.hasOwnProperty.call(i,n)&&(e[n]=i[n]);d&&d(t);while(_.length)_.shift()();return s.push.apply(s,u||[]),a()}function a(){for(var e,t=0;t<s.length;t++){for(var a=s[t],n=!0,r=1;r<a.length;r++){var c=a[r];0!==o[c]&&(n=!1)}n&&(s.splice(t--,1),e=i(i.s=a[0]))}return e}var n={},r={app:0},o={app:0},s=[];function c(e){return i.p+"js/"+({}[e]||e)+"."+{"chunk-89af92f6":"a4aa5b6e","chunk-0a853b2b":"f03209ea","chunk-0b8f8bbc":"c068027d","chunk-162ff1ff":"0a22436a","chunk-22b57838":"ddf3b8f4","chunk-22f331f4":"e1688392","chunk-232f6874":"c45db812","chunk-29553eaa":"4ebf8cb5","chunk-3d9bb2d6":"3fe8729b","chunk-5a77cfd8":"3a7ad932","chunk-768b7ab8":"2d2be15d","chunk-7ef915ac":"326a9dfa","chunk-8a44c3e0":"6e6572a3","chunk-a101e242":"0f232132","chunk-f4b66f66":"eb08a10d"}[e]+".js"}function i(t){if(n[t])return n[t].exports;var a=n[t]={i:t,l:!1,exports:{}};return e[t].call(a.exports,a,a.exports,i),a.l=!0,a.exports}i.e=function(e){var t=[],a={"chunk-0a853b2b":1,"chunk-0b8f8bbc":1,"chunk-22b57838":1,"chunk-3d9bb2d6":1,"chunk-8a44c3e0":1};r[e]?t.push(r[e]):0!==r[e]&&a[e]&&t.push(r[e]=new Promise(function(t,a){for(var n="css/"+({}[e]||e)+"."+{"chunk-89af92f6":"31d6cfe0","chunk-0a853b2b":"614e80cc","chunk-0b8f8bbc":"d93d0cd4","chunk-162ff1ff":"31d6cfe0","chunk-22b57838":"a3c25090","chunk-22f331f4":"31d6cfe0","chunk-232f6874":"31d6cfe0","chunk-29553eaa":"31d6cfe0","chunk-3d9bb2d6":"8f661264","chunk-5a77cfd8":"31d6cfe0","chunk-768b7ab8":"31d6cfe0","chunk-7ef915ac":"31d6cfe0","chunk-8a44c3e0":"1138a732","chunk-a101e242":"31d6cfe0","chunk-f4b66f66":"31d6cfe0"}[e]+".css",o=i.p+n,s=document.getElementsByTagName("link"),c=0;c<s.length;c++){var u=s[c],l=u.getAttribute("data-href")||u.getAttribute("href");if("stylesheet"===u.rel&&(l===n||l===o))return t()}var _=document.getElementsByTagName("style");for(c=0;c<_.length;c++){u=_[c],l=u.getAttribute("data-href");if(l===n||l===o)return t()}var d=document.createElement("link");d.rel="stylesheet",d.type="text/css",d.onload=t,d.onerror=function(t){var n=t&&t.target&&t.target.src||o,s=new Error("Loading CSS chunk "+e+" failed.\n("+n+")");s.request=n,delete r[e],d.parentNode.removeChild(d),a(s)},d.href=o;var p=document.getElementsByTagName("head")[0];p.appendChild(d)}).then(function(){r[e]=0}));var n=o[e];if(0!==n)if(n)t.push(n[2]);else{var s=new Promise(function(t,a){n=o[e]=[t,a]});t.push(n[2]=s);var u,l=document.createElement("script");l.charset="utf-8",l.timeout=120,i.nc&&l.setAttribute("nonce",i.nc),l.src=c(e),u=function(t){l.onerror=l.onload=null,clearTimeout(_);var a=o[e];if(0!==a){if(a){var n=t&&("load"===t.type?"missing":t.type),r=t&&t.target&&t.target.src,s=new Error("Loading chunk "+e+" failed.\n("+n+": "+r+")");s.type=n,s.request=r,a[1](s)}o[e]=void 0}};var _=setTimeout(function(){u({type:"timeout",target:l})},12e4);l.onerror=l.onload=u,document.head.appendChild(l)}return Promise.all(t)},i.m=e,i.c=n,i.d=function(e,t,a){i.o(e,t)||Object.defineProperty(e,t,{enumerable:!0,get:a})},i.r=function(e){"undefined"!==typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(e,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(e,"__esModule",{value:!0})},i.t=function(e,t){if(1&t&&(e=i(e)),8&t)return e;if(4&t&&"object"===typeof e&&e&&e.__esModule)return e;var a=Object.create(null);if(i.r(a),Object.defineProperty(a,"default",{enumerable:!0,value:e}),2&t&&"string"!=typeof e)for(var n in e)i.d(a,n,function(t){return e[t]}.bind(null,n));return a},i.n=function(e){var t=e&&e.__esModule?function(){return e["default"]}:function(){return e};return i.d(t,"a",t),t},i.o=function(e,t){return Object.prototype.hasOwnProperty.call(e,t)},i.p="/web/dist/",i.oe=function(e){throw console.error(e),e};var u=window["webpackJsonp"]=window["webpackJsonp"]||[],l=u.push.bind(u);u.push=t,u=u.slice();for(var _=0;_<u.length;_++)t(u[_]);var d=l;s.push([0,"chunk-vendors"]),a()})({0:function(e,t,a){e.exports=a("56d7")},"034f":function(e,t,a){"use strict";var n=a("27fb"),r=a.n(n);r.a},"0351":function(e,t,a){var n={"./Dashboard.vue":["9b1c","chunk-89af92f6","chunk-0a853b2b"],"./Install.vue":["f10e","chunk-89af92f6","chunk-3d9bb2d6"],"./Layer.vue":["3635","chunk-89af92f6","chunk-22b57838"],"./Login.vue":["dd1d","chunk-89af92f6","chunk-0b8f8bbc"],"./deploy/Apply.vue":["3314","chunk-89af92f6","chunk-162ff1ff"],"./deploy/Deploy.vue":["c75b","chunk-89af92f6","chunk-7ef915ac"],"./deploy/Release.vue":["24bf","chunk-89af92f6","chunk-a101e242"],"./project/Member.vue":["6737","chunk-89af92f6","chunk-22f331f4"],"./project/Project.vue":["882b","chunk-89af92f6","chunk-8a44c3e0"],"./project/Space.vue":["908b","chunk-89af92f6","chunk-5a77cfd8"],"./server/Group.vue":["a34e","chunk-89af92f6","chunk-232f6874"],"./server/Server.vue":["4341","chunk-89af92f6","chunk-f4b66f66"],"./user/Group.vue":["6283","chunk-89af92f6","chunk-29553eaa"],"./user/User.vue":["66b0","chunk-89af92f6","chunk-768b7ab8"]};function r(e){var t=n[e];return t?Promise.all(t.slice(1).map(a.e)).then(function(){var e=t[0];return a(e)}):Promise.resolve().then(function(){var t=new Error("Cannot find module '"+e+"'");throw t.code="MODULE_NOT_FOUND",t})}r.keys=function(){return Object.keys(n)},r.id="0351",e.exports=r},"27fb":function(e,t,a){},"3dfd":function(e,t,a){"use strict";var n=function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("div",{staticClass:"global-container"},[a("router-view")],1)},r=[],o=(a("034f"),a("2877")),s={},c=Object(o["a"])(s,n,r,!1,null,null,null);c.options.__file="App.vue";t["a"]=c.exports},"3fab":function(e,t,a){"use strict";t["a"]={}},4360:function(e,t,a){"use strict";a("cadf"),a("551c"),a("097d");var n=a("2b0e"),r=a("2f62"),o=a("8d81"),s=a.n(o),c={user_id:0,username:"",email:"",mobile:"",privilege:[],role_name:"",truename:""},i={isLogin:function(e){return e.user_id>0},getEmail:function(e){return e.email},getUserName:function(e){return e.username},getUserId:function(e){return e.user_id},getAvatar:function(e){return"https://www.gravatar.com/avatar/"+s()(e.email.toLowerCase())+"? s=512"},getPriv:function(e){return e.privilege?e.privilege:[]},getRoleName:function(e){return e.role_name},getMobile:function(e){return e.mobile},getTrueName:function(e){return e.truename}},u={status:function(e,t){var a=e.commit;a("setUserInfo",{user_id:t.user_id,username:t.username,email:t.email,mobile:t.mobile,privilege:t.privilege,role_name:t.role_name,truename:t.truename})},userSetting:function(e,t){var a=e.commit;a("userSetting",{mobile:t.mobile,truename:t.truename})}},l={setUserInfo:function(e,t){e.user_id=t.user_id,e.username=t.username,e.email=t.email,e.privilege=t.privilege,e.role_name=t.role_name,e.mobile=t.mobile,e.truename=t.truename},userSetting:function(e,t){e.mobile=t.mobile,e.truename=t.truename}},_={namespaced:!0,state:c,getters:i,actions:u,mutations:l};n["default"].use(r["a"]);t["a"]=new r["a"].Store({modules:{account:_}})},4678:function(e,t,a){var n={"./af":"2bfb","./af.js":"2bfb","./ar":"8e73","./ar-dz":"a356","./ar-dz.js":"a356","./ar-kw":"423e","./ar-kw.js":"423e","./ar-ly":"1cfd","./ar-ly.js":"1cfd","./ar-ma":"0a84","./ar-ma.js":"0a84","./ar-sa":"8230","./ar-sa.js":"8230","./ar-tn":"6d83","./ar-tn.js":"6d83","./ar.js":"8e73","./az":"485c","./az.js":"485c","./be":"1fc1","./be.js":"1fc1","./bg":"84aa","./bg.js":"84aa","./bm":"a7fa","./bm.js":"a7fa","./bn":"9043","./bn.js":"9043","./bo":"d26a","./bo.js":"d26a","./br":"6887","./br.js":"6887","./bs":"2554","./bs.js":"2554","./ca":"d716","./ca.js":"d716","./cs":"3c0d","./cs.js":"3c0d","./cv":"03ec","./cv.js":"03ec","./cy":"9797","./cy.js":"9797","./da":"0f14","./da.js":"0f14","./de":"b469","./de-at":"b3eb","./de-at.js":"b3eb","./de-ch":"bb71","./de-ch.js":"bb71","./de.js":"b469","./dv":"598a","./dv.js":"598a","./el":"8d47","./el.js":"8d47","./en-SG":"cdab","./en-SG.js":"cdab","./en-au":"0e6b","./en-au.js":"0e6b","./en-ca":"3886","./en-ca.js":"3886","./en-gb":"39a6","./en-gb.js":"39a6","./en-ie":"e1d3","./en-ie.js":"e1d3","./en-il":"7333","./en-il.js":"7333","./en-nz":"6f50","./en-nz.js":"6f50","./eo":"65db","./eo.js":"65db","./es":"898b","./es-do":"0a3c","./es-do.js":"0a3c","./es-us":"55c9","./es-us.js":"55c9","./es.js":"898b","./et":"ec18","./et.js":"ec18","./eu":"0ff2","./eu.js":"0ff2","./fa":"8df4","./fa.js":"8df4","./fi":"81e9","./fi.js":"81e9","./fo":"0721","./fo.js":"0721","./fr":"9f26","./fr-ca":"d9f8","./fr-ca.js":"d9f8","./fr-ch":"0e49","./fr-ch.js":"0e49","./fr.js":"9f26","./fy":"7118","./fy.js":"7118","./ga":"5120","./ga.js":"5120","./gd":"f6b4","./gd.js":"f6b4","./gl":"8840","./gl.js":"8840","./gom-latn":"0caa","./gom-latn.js":"0caa","./gu":"e0c5","./gu.js":"e0c5","./he":"c7aa","./he.js":"c7aa","./hi":"dc4d","./hi.js":"dc4d","./hr":"4ba9","./hr.js":"4ba9","./hu":"5b14","./hu.js":"5b14","./hy-am":"d6b6","./hy-am.js":"d6b6","./id":"5038","./id.js":"5038","./is":"0558","./is.js":"0558","./it":"6e98","./it-ch":"6f12","./it-ch.js":"6f12","./it.js":"6e98","./ja":"079e","./ja.js":"079e","./jv":"b540","./jv.js":"b540","./ka":"201b","./ka.js":"201b","./kk":"6d79","./kk.js":"6d79","./km":"e81d","./km.js":"e81d","./kn":"3e92","./kn.js":"3e92","./ko":"22f8","./ko.js":"22f8","./ku":"2421","./ku.js":"2421","./ky":"9609","./ky.js":"9609","./lb":"440c","./lb.js":"440c","./lo":"b29d","./lo.js":"b29d","./lt":"26f9","./lt.js":"26f9","./lv":"b97c","./lv.js":"b97c","./me":"293c","./me.js":"293c","./mi":"688b","./mi.js":"688b","./mk":"6909","./mk.js":"6909","./ml":"02fb","./ml.js":"02fb","./mn":"958b","./mn.js":"958b","./mr":"39bd","./mr.js":"39bd","./ms":"ebe4","./ms-my":"6403","./ms-my.js":"6403","./ms.js":"ebe4","./mt":"1b45","./mt.js":"1b45","./my":"8689","./my.js":"8689","./nb":"6ce3","./nb.js":"6ce3","./ne":"3a39","./ne.js":"3a39","./nl":"facd","./nl-be":"db29","./nl-be.js":"db29","./nl.js":"facd","./nn":"b84c","./nn.js":"b84c","./pa-in":"f3ff","./pa-in.js":"f3ff","./pl":"8d57","./pl.js":"8d57","./pt":"f260","./pt-br":"d2d4","./pt-br.js":"d2d4","./pt.js":"f260","./ro":"972c","./ro.js":"972c","./ru":"957c","./ru.js":"957c","./sd":"6784","./sd.js":"6784","./se":"ffff","./se.js":"ffff","./si":"eda5","./si.js":"eda5","./sk":"7be6","./sk.js":"7be6","./sl":"8155","./sl.js":"8155","./sq":"c8f3","./sq.js":"c8f3","./sr":"cf1e","./sr-cyrl":"13e9","./sr-cyrl.js":"13e9","./sr.js":"cf1e","./ss":"52bd","./ss.js":"52bd","./sv":"5fbd","./sv.js":"5fbd","./sw":"74dc","./sw.js":"74dc","./ta":"3de5","./ta.js":"3de5","./te":"5cbb","./te.js":"5cbb","./tet":"576c","./tet.js":"576c","./tg":"3b1b","./tg.js":"3b1b","./th":"10e8","./th.js":"10e8","./tl-ph":"0f38","./tl-ph.js":"0f38","./tlh":"cf75","./tlh.js":"cf75","./tr":"0e81","./tr.js":"0e81","./tzl":"cf51","./tzl.js":"cf51","./tzm":"c109","./tzm-latn":"b53d","./tzm-latn.js":"b53d","./tzm.js":"c109","./ug-cn":"6117","./ug-cn.js":"6117","./uk":"ada2","./uk.js":"ada2","./ur":"5294","./ur.js":"5294","./uz":"2e8c","./uz-latn":"010e","./uz-latn.js":"010e","./uz.js":"2e8c","./vi":"2921","./vi.js":"2921","./x-pseudo":"fd7e","./x-pseudo.js":"fd7e","./yo":"7f33","./yo.js":"7f33","./zh-cn":"5c3a","./zh-cn.js":"5c3a","./zh-hk":"49ab","./zh-hk.js":"49ab","./zh-tw":"90ea","./zh-tw.js":"90ea"};function r(e){var t=o(e);return a(t)}function o(e){var t=n[e];if(!(t+1)){var a=new Error("Cannot find module '"+e+"'");throw a.code="MODULE_NOT_FOUND",a}return t}r.keys=function(){return Object.keys(n)},r.resolve=o,e.exports=r,r.id="4678"},"4f87":function(e,t,a){},"56d7":function(e,t,a){"use strict";a.r(t),function(e){a("cadf"),a("551c"),a("097d");var t,n=a("2b0e"),r=a("5c96"),o=a.n(r),s=(a("0fae"),a("c1df")),c=a.n(s),i=a("3dfd"),u=a("a18c"),l=a("4360"),_=a("9923"),d=a("f1ed"),p=a("9607");a("4f87");e.navigator.language&&(t=e.navigator.language,t=t.toLowerCase()),0!=t.indexOf("en")&&(t="zh-cn"),c.a.locale(t),n["default"].config.productionTip=!1,n["default"].use(o.a),new n["default"]({i18n:_["a"],router:u["a"],store:l["a"],methods:d["a"],data:p["a"],render:function(e){return e(i["a"])}}).$mount("#app")}.call(this,a("c8ba"))},"6e41":function(e,t,a){"use strict";a("cadf"),a("551c"),a("097d");t["a"]={welcome_to_login_zoom:"欢迎登录",username_or_email:"用户名或邮箱",password:"密码",login:"登录",install:"安装",please_input_loginname:"请输入登录名(用户名或邮箱)",please_input_password:"请输入密码",strlen_between:"长度在 {min} 到 {max} 个字符",say_hello_to_login_user:"你好, {username}",zoom_document:"帮助文档",personal_setting:"个人设置",change_password:"修改密码",sign_out:"退出登录",dashboard:"控制台",submit_deploy_apply:"提交发布申请",deploy_manage:"发布单管理",deploying_deploy:"部署发布单",project:"项目",space_manage:"空间管理",project_manage:"项目管理",member_manage:"成员管理",user:"用户",role_manage:"角色管理",user_manage:"用户管理",server:"服务器",cluster_manage:"集群管理",server_manage:"服务器管理",add_cluster:"新增集群",please_input_keyword:"关键词搜索",please_input_keyword_id_or_name:"关键词搜索，支持ID、名称",network_error:"网络错误",unknown_error:"未知错误",operate_success:"操作成功",operate:"操作",name:"名称",edit:"编辑",delete:"删除",server_group_name:"集群名称",name_cannot_empty:"名称不能为空",cancel:"取消",enter:"确定",edit_cluster:"编辑集群",add_server:"新增服务器",edit_server:"编辑服务器",deleted:"已删除",belong_cluster:"所属集群",belong_cluster_cannot_empty:"所属集群不能为空",keyword_search:"关键字搜索",server_name:"服务器名称",please_input_server_name:"请输入服务器名称",IP_HOST_cannot_empty:"IP/HOST不能为空",please_input_server_IP_HOST:"请输入服务器IP/HOST",SSH_port:"SSH端口",SSH_port_cannot_empty:"SSH端口不能为空",edit_server_info:"编辑服务器信息",add_role:"新增角色",edit_role_info:"编辑角色信息",role_name:"角色名称",privilege_setting:"权限设置",check_all:"全选",add_user:"新增用户",cluster:"集群",role:"角色",email:"邮箱",status:"状态",last_login:"上次登录",username:"用户名",role_cannot_empty:"角色不能为空",username_cannot_empty:"用户名不能为空",please_input_username:"请输入用户名",password_cannot_empty:"密码不能为空",please_input_password_length_limit:"请输入密码, 6到20个字符",truename:"真实姓名",mobile:"手机号",allow_login:"允许登录",email_cannot_empty:"邮箱不能为空",please_input_email:"请输入邮箱",normal:"正常",locking:"锁定",last_login_time:"上次登录时间",last_login_ip:"上次登录IP",email_format_wrong:"邮箱格式错误",users_cannot_login_after_being_banned:"禁止后用户将无法登录",data_repeat_reenter_please:"数据重复，请重新输入",in_verification_please_wait:"验证中，请稍后",username_exists_please_reenter:"用户名已经存在，请重新输入",verifying_username_is_being_used_please_wait:"正在验证用户名是否被占用，请稍等",email_is_exists_please_reenter:"邮箱已经存在，请重新输入",verifying_email_is_being_used_please_wait:"正在验证邮箱是否被占用，请稍等",network_error_verify_failed:"网络错误, 校验失败",project_space:"项目空间",add_project_space:"新增项目空间",input_name_to_search:"按名称搜索",edit_project_space_info:"编辑项目空间信息",project_space_name:"项目空间名称",description:"描述信息",add_member_to_space:"添加成员到项目空间",add:"添加",project_member_select_space_tips:"成员管理需要指定项目空间，请点击 '选择项目空间' 按钮，在弹出窗口中进行选择。",prompt_message:"提示信息",select_project_space:"选择项目空间",switch_project_space:"切换项目空间",add_new_member:"添加新成员",looking_for_users:"正在查找用户...",user_not_found:"未查找到用户",search_for_users_by_username_and_email:"通过用户名、邮箱搜索用户",search_and_select_users_before_adding:"请先搜索并选择用户后再添加",member_added_successfully:"成员添加成功",member_already_exists_do_not_add_repeat:"成员已经存在，请勿重复添加",remove:"移除",project_select_space_tips:"项目管理需要指定项目空间，请点击 '选择项目空间' 按钮，在弹出窗口中进行选择。",add_project:"新增项目",project_name:"项目名称",please_input_project_name:"请输入项目名称",please_input_project_description:"请输入项目描述信息",open_audit:"开启审核",if_open_apply_need_audit:"开启后，上线单需要审核通过后才能发起上线",repo_url:"仓库地址",please_input_repo_url:"请输入仓库地址",repo_url_cannot_empty:"仓库地址不能为空",repo_branch:"指定上线分支",online_cluster:"线上集群",online_cluster_cannot_empty:"线上集群不能为空",user_cannot_empty:"用户不能为空",deploy_user:"目标机用户",path:"目录",deploy_path:"代码/包部署的目录",deploy_path_cannot_empty:"请设置代码部署目录",pre_deploy_cmd:"部署前运行命令",pre_deploy_cmd_tips:"代码部署之前在目标机运行的命令，每行一个命令",after_deploy_cmd:"部署后运行命令",after_deploy_cmd_tips:"代码部署之后在目标机运行的命令，每行一个命令",deploy_timeout:"部署超时时间(秒)",deploy_timeout_tips:"部署命令最大运行时间",project_enable:"项目启用",edit_project:"编辑项目",view:"查看",copy:"复制",project_id:"项目ID",build:"构建",build_setting:"构建设置",build_script:"构建脚本",view_build_script_eg:"查看构建脚本示例",open:"开启",close:"关闭",repo_setting:"仓库设置",if_not_need_to_assign_branch_name:"若不指定，需要在发起上线时手动填写分支名称",deploy_setting:"部署设置",selected_cluster_list:"已选集群列表",cluster_list:"集群列表",view_project_info:"查看项目信息",base_setting:"基本设置",need_audit:"需要审核",not_audit:"不需审核",audit:"审核",unaudit:"待审核",edit_build_script:"编辑构建脚本",illustrate:"说明",build_script_tips:"脚本会在代码下载完成后执行，构建脚本支持的变量",build_script_env_workspace:"代码仓库本地副本目录",build_script_env_pack_file:'打包文件绝对地址，构建完成后将需要部署到线上的代码打包到此文件中，必须使用 <span class="code">tar -zcf</span> 命令进行打包。部署模块会将此压缩包分发到目标主机并解压缩到指定目录，请按照要求打包，否则会部署失败。',project_space_cannot_empty:"项目空间不能为空",project_cannot_empty:"项目不能为空",input_apply_order:"填写上线单",input_deploy_apply:"填写上线申请",deploy_mode:"上线模式",deploy_mode_cannot_empty:"上线模式不能为空",branch_deploy:"分支上线",tag_deploy:"Tag上线",deploy_mode_tips:"测试环境推荐分支上线，生产环境推荐tag上线",apply_name:"发布单名称",apply_order:"发布单",please_input_apply_name:"请输入上线单名称",tag_name:"Tag名称",tag_name_cannot_empty:"Tag名称不能为空",please_input_tag_name:"请输入Tag名称",branch_name_cannot_empty:"分支名称不能为空",please_input_branch_name:"请输入分支名称",branch_name:"分支名称",deploy_illustrate:"上线说明",please_input_deploy_illustrate:"请详细填写上线说明",deploy_illustrate_cannot_empty:"上线说明不能为空",commit_version:"选择上线版本",version:"版本",please_input_commit_version:"输入上线版本hash",branch:"分支",deploy_apply_submit_success:"恭喜，上线申请提交成功",submit_success:"提交成功",submit_time:"提交时间",deploy_status:"上线状态",audit_status:"审核状态",select_project:"选择项目",rollback_apply:"绑定回滚单",belong_to_space:"所属空间",submiter:"提交者",pass:"通过",denied:"拒绝",wait_online:"待上线",onlineing:"上线中",success:"成功",failed:"失败",drop:"废弃",rollback:"回滚",online:"上线",space_name:"空间名称",audit_pass:"审核通过",audit_denied:"审核拒绝",deined_reason:"拒绝原因",edit_apply_order:"编辑上线单",today:"今天","7day":"7天内",within_one_month:"1个月内",within_three_months:"3个月内",within_a_year:"一年内",any_time:"时间不限",not_online:"未上线",online_success:"上线成功",online_failed:"上线失败",deprecated:"已废弃",update_success:"更新成功",audit_success:"审核成功",drop_deploy_apply_tips:"此操作将废弃该上线单, 是否继续?",have_onlined:"已上线",belong_project:"所属项目",forced_termination:"强制终止",last_buid_time:"上次构建时间",cost_time:"耗时",stopping:"正在终止",building:"构建中",build_finish:"构建成功",build_failed:"构建失败",build_ing:"构建中",unbuild:"未构建",uncreate:"未生成",build_log:"构建日志",tar_pack_path:"Tar包位置",deploy:"部署",redeploy:"重新部署",rollback_apply_order_tips:"回滚操作将以新发布单的形式对历史代码包进行重新部署",rollback_created:"回滚单已生成",rollback_unstart:"回滚未开始",rollbacking:"回滚中",rollback_success:"回滚成功",rollback_failed:"回滚失败",rollback_drop:"回滚单废弃",rollback_unknow:"回滚状态未知",click_to_view_rollback_order:"点击进入回滚单",wait_deploy:"等待部署",deploying:"部署中",deploy_success:"部署成功",deploy_failed:"部署失败",be_deined:"被终止",deploy_log:"部署日志",makesure_rollback_order:"确定要进行回滚操作吗？",build_success_and_deploy:"请构建成功后再进行部署操作",deploy_tips:"部署提示",start_build:"开始构建",i_known:"知道了",yes:"是",no:"否",have_enabled:"已启用",not_enable:"未启用",avatar:"头像",current_password:"当前密码",new_password:"新密码",current_password_cannot_empty:"当前密码不能为空",new_password_cannot_empty:"新密码不能为空",apply_type:"类型",rollback_creating:"回滚单生成中",rollback_tips:"当前是自动生成的回滚单",back_apply:"返回原发布单",email_setting:"邮箱设置",audit_notice:"审核通知",audit_notice_tips:"请输入接收审核通知的邮箱地址",audit_notice_explain:'系统通过邮件通知相关负责人审核待审单，多个邮箱地址请用 <span class="app-tag-info">,</span> 相隔',deploy_notice:"上线通知",deploy_notice_tips:"请输入接收上线通知的邮箱地址",deploy_notice_explain:'接收上线通知的邮箱地址，多个邮箱地址请用 <span class="app-tag-info">,</span> 相隔',hook:"Hook",edit_hook_script:"编辑Hook命令",build_hook_script:"构建完成后执行命令",deploy_hook_script:"部署完成后执行命令",build_hook_script_tips:"命令会在构建完成(成功或失败)后执行，支持的变量",deploy_hook_script_tips:"命令会在部署完成(成功或失败)后执行，支持的变量"}},"8e09":function(e,t,a){"use strict";a("cadf"),a("551c"),a("097d");t["a"]={DEPLOY_APPLY:1001,DEPLOY_VIEW:1002,DEPLOY_AUDIT:1003,DEPLOY_DEPLOY:1004,DEPLOY_DROP:1005,DEPLOY_EDIT:1006,PROJECT_SPACE_VIEW:2001,PROJECT_SPACE_NEW:2002,PROJECT_SPACE_EDIT:2003,PROJECT_SPACE_DEL:2004,PROJECT_USER_VIEW:2100,PROJECT_USER_NEW:2101,PROJECT_USER_DEL:2102,PROJECT_VIEW:2201,PROJECT_NEW:2202,PROJECT_EDIT:2203,PROJECT_DEL:2204,PROJECT_AUDIT:2205,PROJECT_BUILD:2206,PROJECT_HOOK:2207,PROJECT_COPY:2208,USER_ROLE_VIEW:3001,USER_ROLE_NEW:3002,USER_ROLE_EDIT:3003,USER_ROLE_DEL:3004,USER_VIEW:3101,USER_NEW:3102,USER_EDIT:3103,USER_DEL:3104,SERVER_GROUP_VIEW:4001,SERVER_GROUP_NEW:4002,SERVER_GROUP_EDIT:4003,SERVER_GROUP_DEL:4004,SERVER_VIEW:4101,SERVER_NEW:4102,SERVER_EDIT:4103,SERVER_DEL:4104,SYSTEM_STATUS:5001}},9607:function(e,t,a){"use strict";var n,r=a("bd86"),o=a("8e09");t["a"]=(n={PageSize:0,Page:0,Total:0,DialogSmallWidth:"500px",DialogNormalWidth:"750px",DialogLargeWidth:"900px",DialogNormalTop:"5vh",Priv:o["a"],BuildStatusNone:0,BuildStatusStart:1,BuildStatusSuccess:2,BuildStatusFailed:3,ApplyStatusNone:1,ApplyStatusIng:2,ApplyStatusSuccess:3,ApplyStatusFailed:4,ApplyStatusDrop:5,ApplyStatusRollback:6,DeployModeBranch:1,DeployModelTag:2},Object(r["a"])(n,"BuildStatusNone",0),Object(r["a"])(n,"BuildStatusStart",1),Object(r["a"])(n,"BuildStatusSuccess",2),Object(r["a"])(n,"BuildStatusFailed",3),Object(r["a"])(n,"DeployGroupStatusNone",0),Object(r["a"])(n,"DeployGroupStatusStart",1),Object(r["a"])(n,"DeployGroupStatusSuccess",2),Object(r["a"])(n,"DeployGroupStatusFailed",3),Object(r["a"])(n,"AuditStatusPending",1),Object(r["a"])(n,"AuditStatusOk",2),Object(r["a"])(n,"AuditStatusRefuse",3),n)},9923:function(e,t,a){"use strict";(function(e){a("cadf"),a("551c"),a("097d");var n=a("2b0e"),r=a("a925"),o=a("6e41"),s=a("3fab");n["default"].use(r["a"]);var c,i={"zh-cn":o["a"],en:s["a"]};e.navigator.language&&(c=e.navigator.language,c=c.toLowerCase()),0!=c.indexOf("en")&&(c="zh-cn");var u=new r["a"]({locale:c,messages:i});t["a"]=u}).call(this,a("c8ba"))},a18c:function(e,t,a){"use strict";a.d(t,"b",function(){return u});a("cadf"),a("551c"),a("097d");var n=a("2b0e"),r=a("8c4f"),o=a("9923"),s=a("8e09");n["default"].use(r["a"]);var c=function(e){return function(){return a("0351")("./"+e+".vue")}},i=[{path:"/install",name:"install",component:c("Install")},{path:"/login",name:"login",component:c("Login")}],u=[{path:"/",component:c("Layer"),name:"main",meta:{},redirect:{name:"dashboard"},children:[{path:"dashboard",name:"dashboard",meta:{title:o["a"].t("dashboard"),icon:"icon-dashboard",single:!0},component:c("Dashboard")}]},{path:"/deploy",name:"deploy",component:c("Layer"),meta:{title:o["a"].t("deploy"),icon:"icon-send"},children:[{path:"apply",name:"deployApply",meta:{title:o["a"].t("submit_deploy_apply"),role:[s["a"].DEPLOY_APPLY]},component:c("deploy/Apply")},{path:"deploy",name:"deployDeploy",meta:{title:o["a"].t("deploy_manage"),role:[s["a"].DEPLOY_VIEW]},component:c("deploy/Deploy")},{path:"release",name:"deployRelease",meta:{title:o["a"].t("deploying_deploy"),hide:!0},component:c("deploy/Release")}]},{path:"/project",name:"project",component:c("Layer"),meta:{title:o["a"].t("project"),icon:"icon-project"},children:[{path:"space",name:"projectSpace",meta:{title:o["a"].t("space_manage"),role:[s["a"].PROJECT_SPACE_VIEW]},component:c("project/Space")},{path:"project",name:"projectProject",meta:{title:o["a"].t("project_manage"),role:[s["a"].PROJECT_VIEW]},component:c("project/Project")},{path:"user",name:"projectUser",meta:{title:o["a"].t("member_manage"),role:[s["a"].PROJECT_USER_VIEW]},component:c("project/Member")}]},{path:"/user",name:"user",component:c("Layer"),meta:{title:o["a"].t("user"),icon:"icon-group"},children:[{path:"group",name:"userGroup",meta:{title:o["a"].t("role_manage"),role:[s["a"].USER_ROLE_VIEW]},component:c("user/Group")},{path:"list",name:"userList",meta:{title:o["a"].t("user_manage"),role:[s["a"].USER_VIEW]},component:c("user/User")}]},{path:"/server",name:"server",component:c("Layer"),meta:{title:o["a"].t("server"),icon:"icon-server"},children:[{path:"group",name:"serverGroup",meta:{title:o["a"].t("cluster_manage"),role:[s["a"].SERVER_GROUP_VIEW]},component:c("server/Group")},{path:"list",name:"serverList",meta:{title:o["a"].t("server_manage"),role:[s["a"].SERVER_VIEW]},component:c("server/Server")}]}],l=new r["a"]({routes:i.concat(u),scrollBehavior:function(){return{y:0}}});t["a"]=l},f1ed:function(e,t,a){"use strict";a("ac6a"),a("f3e2"),a("57e7"),a("6b54"),a("87b3"),a("cadf"),a("551c"),a("097d");var n=a("8d81"),r=a.n(n),o=a("a78e"),s=a.n(o),c=a("c1df"),i=a.n(c),u="_zoom_identity";t["a"]={MessageSuccess:function(e){this.$message({message:this.$t("operate_success"),type:"success",duration:1200,onClose:e})},PageInit:function(){this.$root.PageSize=7,this.$root.Page=1,this.$root.Total=0},PageReset:function(){this.$root.Total--;var e=Math.ceil(this.$root.Total/this.$root.PageSize);this.$root.Page>e&&(this.$root.Page=e),this.$root.Page<1&&(this.$root.Page=1)},PageOffset:function(){return this.$root.PageSize*(this.$root.Page-1)},ConfirmDelete:function(e,t){t||(t="此操作将永久删除该数据, 是否继续?"),this.$confirm(t,"提示",{confirmButtonText:"确定",cancelButtonText:"取消",type:"warning"}).then(function(){e()}).catch(function(e){console.log(e)})},ConfirmCopy:function(e,t){t||(t="此操作将复制一个新项目，是否继续?"),this.$confirm(t,"提示",{confirmButtonText:"确定",cancelButtonText:"取消",type:"info"}).then(function(){e()}).catch(function(e){console.log(e)})},Md5Sum:function(e){return r()(e)},FormatDateTime:function(e,t){return e?(t||(t="YYYY-MM-DD HH:mm:ss"),i.a.unix(e).format(t)):"--"},FormatDateDuration:function(e){return i.a.duration(e).humanize(!1)},FormatDateFromNow:function(e){return e?i.a.unix(e).fromNow():"--"},Substr:function(e,t){if("[object String]"!=Object.prototype.toString.call(e))return"";var a="";return e.length>t&&(a="..."),e.substr(0,t)+a},SetLoginToken:function(e){return s.a.set(u,e)},GetLoginToken:function(){return s.a.get(u)},DeleteLoginToken:function(){return s.a.remove(u)},CheckPriv:function(e){return this.$store.getters["account/getPriv"].indexOf(e)>-1},CheckPrivs:function(e){var t=this;if(!e||!e.length)return!1;var a=!1;return e.forEach(function(e){t.CheckPriv(e)&&(a=!0)}),a}}}});
//# sourceMappingURL=app.399af180.js.map