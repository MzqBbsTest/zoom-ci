<template>
    <div>
        <el-popover
            placement="bottom-start"
            width="400"
            trigger="hover">
            <el-table :data="serverList">
                <el-table-column width="250" property="name" label="服务器名字"></el-table-column>
                <el-table-column width="150" property="address" label="操作">
                    <template slot-scope="scope">
                        <el-button @click="openFtpDialogHandler(scope.row)"  size="small">sftp</el-button>
                        <el-button @click="openXtermDialogHandler(scope.row)"  size="small">连接</el-button>
                    </template>
                </el-table-column>
            </el-table>
            <el-pagination
                background
                layout="prev, pager, next"
                class="app-pagination"
                @current-change="currentChangeHandler(1)"
                :current-page.sync="serverPage.Page"
                :page-size="serverPage.pageSize"
                :total="serverPage.total">
            </el-pagination>
            <el-button slot="reference">打开</el-button>
        </el-popover>
        <el-popover
            placement="bottom-start"
            width="400"
            trigger="hover">
            <el-button >新建</el-button>
            <el-table :data="cmdList">
                <el-table-column width="300" property="date" label="命令标题"></el-table-column>
                <el-table-column width="100" property="name" label="复制">
                    <template slot-scope="scope">
                        <el-button @click="handleClick(scope.row)" size="small">复制</el-button>
                    </template>
                </el-table-column>
            </el-table>
            <el-pagination
                background
                layout="prev, pager, next"
                class="app-pagination"
                @current-change="currentChangeHandler(2)"
                :current-page.sync="cmdPage.Page"
                :page-size="cmdPage.pageSize"
                :total="cmdPage.total">
            </el-pagination>
            <el-button slot="reference">命令搜藏</el-button>
        </el-popover>
        
        <el-tabs v-model="xtermTabsValue" type="card" closable @tab-remove="removeTab">
            <el-tab-pane
                v-for="(item, index) in xtermTabs"
                :key="item.name"
                :label="item.name"
                :name="item.name"
            >
                <ServerXtermPlan :server=item />
            </el-tab-pane>
        </el-tabs>
        

        <el-dialog :title="ftpDialog.title" :visible.sync="ftpDialog.dialogTableVisible" width="80%">
            <div style="margin-top: 15px;">
                <el-input placeholder="路径" v-model="ftpDialog.path" class="input-with-select">
                   
                    <el-button slot="append" @click="sftpList(ftpDialog.path)">进入</el-button>
                    <el-button slot="append" @click="sftpList(ftpDialog.path + '/..')">回退</el-button>
                    <el-button slot="append" >上传</el-button>
                    <el-button slot="append" >创建目录</el-button>
                    <el-button slot="append" @click="sftpList(ftpDialog.path)">刷新</el-button>
                </el-input>
            </div>
            <el-table :data="ftpDialog.gridData" max-height="500"   :default-sort = "{prop: 'name', order: 'descending'}">
                <el-table-column property="name" label="文件名" sortable style="font-size:12px" >
                    <template slot-scope="scope">
                        <i :class="{
                            'el-icon-folder-opened': scope.row.type == 'd' ,
                            'dir':  scope.row.type == 'd',
                             }" >
                            {{ scope.row.name }}<i v-if="scope.row.link_target" :class="{ 'link':  scope.row.link_target}">-> {{ scope.row.link_target }}</i> 
                        </i>
                    </template>
                </el-table-column>
                <el-table-column property="mode" label="权限/所有者" ></el-table-column>
                <el-table-column property="mod_time" label="修改时间" sortable ></el-table-column>
                <el-table-column property="address" width="400" label="操作">
                    <template slot-scope="scope">
                        <el-button-group>
                            <el-button @click="sftpList(scope.row.path)"  size="mini" v-if="scope.row.type == 'd'">打开</el-button>
                            <el-button @click="openXtermDialogHandler(scope.row)"  size="mini">重命名</el-button>
                            <el-button @click="openXtermDialogHandler(scope.row)"  size="mini" v-if="scope.row.type == 'd'">上传</el-button>
                            <el-button @click="openXtermDialogHandler(scope.row)"  size="mini">权限</el-button>
                            <el-button @click="openXtermDialogHandler(scope.row)"  size="mini"  v-if="scope.row.tar">压缩</el-button>
                            <el-button @click="openXtermDialogHandler(scope.row)"  size="mini"  v-if="scope.row.untar">解压</el-button>
                            <el-button @click="openXtermDialogHandler(scope.row)"  size="mini">删除</el-button>
                        </el-button-group>
                    </template>
                </el-table-column>
            </el-table>
        </el-dialog>

    </div>
</template>

<style>
    .dir{
        color:#409EFF
    }
    .link{
        color:#E6A23C
    }
    .file{
        color:#67C23A
    }
</style>

<script>
import { listServerApi, testServerConnect, listFtpApi } from '@/api/server'
import ServerXtermPlan from './ServerXtermPlan.vue'
import Vue from "vue";

export default {
    components: {
        ServerXtermPlan
    },
    data() {
        return {
            tabIndex: 0,
            xtermTabsValue:0,
            xtermTabs:[],
            cmdPage:{
                total: 0,
                page: 1,
                pageSize : 10
            },
            serverPage:{
                total: 0,
                page: 1,
                pageSize : 10
            },
            serverList:[],
            tableServerLoading:false,
            cmdList:[],
            tableCmdLoading:false,
            ftpDialog:{
                id:0,
                title:"",
                dialogTableVisible: false,
                gridData:[],
                path: ""
            }
        }
    },
    methods: {
        removeTab(name){
            this.xtermTabs = this.xtermTabs.filter(item => item.name !== name);
            Vue.nextTick(()=>{
                this.$root.EmitEventGlobal("closeXtermDialogHandler", { server: {
                    name: name
                } });
            })
        },
        sftpList(path){
            listFtpApi({id: this.ftpDialog.id, path:path}).then(res=>{
                this.ftpDialog.gridData = res.files
                this.ftpDialog.path = res.current_dir
            })
        },
        openFtpDialogHandler(row){
            this.ftpDialog.path = ""
            this.ftpDialog.dialogTableVisible = true
            this.ftpDialog.title = row.name
            this.ftpDialog.id = row.id
            this.sftpList("")
        },
        openXtermDialogHandler(row) {
            console.log(row)
            this.tabIndex++;
            let tab = {
                name: row.name + '-' + this.tabIndex,
                server_id: row.id
            }
            this.xtermTabs.push(tab);
            this.xtermTabsValue = tab.name;
            Vue.nextTick(()=>{
                this.$root.EmitEventGlobal("openXtermDialogHandler", { server: tab });
            })
        },  
        deleteHandler(row) {
            this.$root.ConfirmDelete(() => {
                deleteServerApi({ id: row.id }).then(res => {
                    this.$root.MessageSuccess()
                    this.$root.PageReset()
                    this.loadTableData()
                })
            })
        },
        currentChangeHandler(index) {
            if(index == 1) this.loadServerData()
            else  this.loadCmdData()
        },
        loadServerData() {
            this.tableServerLoading = true

            listServerApi({ keyword: this.searchInput, offset: this.$root.PageOffset(), limit: this.$root.PageSize }).then(res => {
                this.serverList = res.list
                this.serverPage.total = res.total
                this.tableServerLoading = false
                
                if(!this.$route.query.server_id){
                    return;
                    
                }
                let s = res.list.filter(item => item.id == this.$route.query.server_id)
                if(!s.length){
                    return 
                }
                this.openXtermDialogHandler(s[0])

            }).catch(err => {
                this.tableServerLoading = false
            })
        },
        loadCmdData() {
            this.tableCmdLoading = true
            listServerApi({ keyword: this.searchInput, offset: this.$root.PageOffset(), limit: this.$root.PageSize }).then(res => {
                this.cmdList = res.list
                this.cmdPage.total = res.total
                this.tableCmdLoading = false
            }).catch(err => {
                this.tableCmdLoading = false
            })
        },
    },
    mounted() {
        this.loadServerData()
        this.loadCmdData()

        window.onbeforeunload = function() {
            return '您确定要离开此页面吗？';
        };

    }
}
</script>
