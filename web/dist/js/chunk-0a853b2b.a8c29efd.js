(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-0a853b2b"],{"40dd":function(t,e,a){},8593:function(t,e,a){"use strict";a.d(e,"a",function(){return s}),a.d(e,"b",function(){return c}),a.d(e,"c",function(){return o});a("cadf"),a("551c"),a("097d");var n=a("ead3");function s(t){return Object(n["b"])("system/install",t)}function c(){return Object(n["a"])("/system/install_status")}function o(){return Object(n["b"])("/system/status")}},"9b1c":function(t,e,a){"use strict";a.r(e);var n=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("el-card",{staticClass:"box-card"},[a("el-row",{attrs:{gutter:30}},[a("el-col",{attrs:{span:16}},[a("el-calendar",{model:{value:t.value,callback:function(e){t.value=e},expression:"value"}})],1),a("el-col",{attrs:{span:8}},[a("el-card",{staticClass:"box-card",staticStyle:{"margin-bottom":"20px"},attrs:{shadow:"hover"}},[a("div",{staticClass:"clearfix",attrs:{slot:"header"},slot:"header"},[a("span",[t._v(t._s(t.$t("system_info")))])]),a("el-tooltip",{attrs:{effect:"dark",content:t.$t("current_config_file_path")+": "+t.currentConfigFilePath,placement:"top"}},[a("div",{staticClass:"text item",staticStyle:{"margin-bottom":"10px"}},[t._v(t._s(t.$t("local_space"))+"："+t._s(t.localSpacePath))])]),a("div",{staticClass:"text item",staticStyle:{"margin-bottom":"10px"}},[t._v(t._s(t.$t("remote_space"))+"："+t._s(t.remoteSpacePath))]),a("div",{staticClass:"text item",staticStyle:{"margin-bottom":"10px"}},[t._v(t._s(t.$t("version"))+"："+t._s(t.currentZoomVersion))])],1)],1)],1)],1)},s=[],c=a("8593"),o={data:function(){return{value:new Date,localSpacePath:"",remoteSpacePath:"",currentConfigFilePath:"",currentZoomVersion:""}},mounted:function(){this.loadSystemStatus()},methods:{loadSystemStatus:function(){var t=this;Object(c["c"])().then(function(e){e&&(t.localSpacePath=e.localSpacePath,t.remoteSpacePath=e.remoteSpacePath,t.currentConfigFilePath=e.currentConfigFilePath,t.currentZoomVersion=e.currentZoomVersion)})}}},r=o,i=(a("b162"),a("2877")),l=Object(i["a"])(r,n,s,!1,null,null,null);l.options.__file="Dashboard.vue";e["default"]=l.exports},b162:function(t,e,a){"use strict";var n=a("40dd"),s=a.n(n);s.a}}]);
//# sourceMappingURL=chunk-0a853b2b.a8c29efd.js.map