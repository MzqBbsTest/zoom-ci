(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-0b8f8bbc"],{"0da5":function(t,e,n){"use strict";var o=n("1057"),i=n.n(o);i.a},1057:function(t,e,n){},3680:function(t,e,n){t.exports=n.p+"img/logo.e34871a2.png"},"7ded":function(t,e,n){"use strict";n.d(e,"a",function(){return i}),n.d(e,"b",function(){return r}),n.d(e,"c",function(){return s});var o=n("ead3");function i(t){return Object(o["b"])("/login",t)}function r(){return Object(o["a"])("/login/status")}function s(){return Object(o["b"])("/logout")}},8593:function(t,e,n){"use strict";n.d(e,"a",function(){return i}),n.d(e,"b",function(){return r}),n.d(e,"c",function(){return s});n("cadf"),n("551c"),n("097d");var o=n("ead3");function i(t){return Object(o["b"])("system/install",t)}function r(){return Object(o["a"])("/system/install_status")}function s(){return Object(o["b"])("/system/status")}},dd1d:function(t,e,n){"use strict";n.r(e);var o=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",{staticClass:"app-login"},[t._m(0),n("div",{staticClass:"app-login-inner"},[n("div",{staticClass:"login-container"},[n("el-card",{staticClass:"login-box"},[n("div",{staticClass:"login-title"},[t._v(t._s(t.$t("welcome_to_login_zoom")))]),n("el-form",{ref:"loginFormRef",staticClass:"login-form",attrs:{model:t.loginForm,rules:t.loginRules,size:"medium"},nativeOn:{keyup:function(e){return"button"in e||!t._k(e.keyCode,"enter",13,e.key,"Enter")?t.loginHandler(e):null}}},[n("el-form-item",{attrs:{prop:"username"}},[n("el-input",{attrs:{placeholder:t.$t("username_or_email"),"prefix-icon":"iconfont icon-user",autocomplete:"off"},model:{value:t.loginForm.username,callback:function(e){t.$set(t.loginForm,"username",e)},expression:"loginForm.username"}})],1),n("el-form-item",{attrs:{prop:"password"}},[n("el-input",{attrs:{placeholder:t.$t("password"),"prefix-icon":"iconfont icon-lock",type:"password",autocomplete:"off"},model:{value:t.loginForm.password,callback:function(e){t.$set(t.loginForm,"password",e)},expression:"loginForm.password"}})],1),n("el-form-item",[n("el-button",{staticStyle:{width:"100%"},attrs:{type:"primary"},on:{click:t.loginHandler}},[t._v(t._s(t.$t("login")))])],1)],1)],1)],1),n("div",{staticClass:"login-cpy"},[t._v("\n            ©️ "+t._s((new Date).getFullYear())+" "),n("a",{attrs:{href:"https://github.com/zoom-ci/zoom-ci/",target:"_blank"}},[t._v("Zoom")]),t._v(". All Rights Reserved. MIT License.\n        ")])])])},i=[function(){var t=this,e=t.$createElement,o=t._self._c||e;return o("div",{staticClass:"login-logo",staticStyle:{"text-align":"center","justify-content":"center","margin-top":"20vh"}},[o("img",{staticStyle:{height:"60px"},attrs:{src:n("3680")}})])}],r=(n("cadf"),n("551c"),n("097d"),n("7ded")),s=n("8593"),a=n("1c06"),l={data:function(){return{loginLoadding:!1,loginForm:{username:"",password:""},loginRules:{username:[{required:!0,message:this.$t("please_input_loginname"),trigger:"blur"}],password:[{required:!0,message:this.$t("please_input_password"),trigger:"blur"},{min:6,max:20,message:this.$t("strlen_between",{min:6,max:20}),trigger:"blur"}]}}},computed:{},mounted:function(){this.initInstallStatus()},methods:{initInstallStatus:function(){var t=this;Object(s["b"])().then(function(e){e.is_installed||t.$router.push({name:"install"})}).catch(function(e){t.$router.push({name:"install"})})},loginHandler:function(){var t=this;this.$refs.loginFormRef.validate(function(e){if(!e)return!1;var n={username:t.loginForm.username,password:t.$root.Md5Sum(t.loginForm.password)};Object(r["a"])(n).then(function(e){t.$root.SetLoginToken(e.token),t.$router.push({name:"dashboard"})}).catch(function(e){e.code&&e.code==a["a"].CODE_ERR_LOGIN_FAILED&&t.$message.error("登录失败, 错误信息: "+e.message)})})}}},u=l,c=(n("0da5"),n("2877")),m=Object(c["a"])(u,o,i,!1,null,null,null);m.options.__file="Login.vue";e["default"]=m.exports}}]);
//# sourceMappingURL=chunk-0b8f8bbc.8ba7eb31.js.map