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
                    <el-button slot="append" @click="uploadFile(ftpDialog.path)">上传</el-button>
                    <el-button slot="append" @click="ftpDialog.dir.visible=true">创建目录</el-button>
                    <el-button slot="append" @click="sftpList(ftpDialog.path)">刷新</el-button>
                </el-input>
            </div>
            <el-table :data="ftpDialog.gridData" max-height="500"   :default-sort = "{prop: 'name', order: 'descending'}">
                <el-table-column property="name" label="文件名" sortable style="font-size:12px" >
                    <template slot-scope="scope">
                        <i :class="{'el-icon-folder-opened': scope.row.type == 'd', 'dir':  scope.row.type == 'd'}" >
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
                            <el-button @click="ftpRenameShow(scope.row)"  size="mini">重命名</el-button>
                            <el-button @click="uploadFile(scope.row.path)"  size="mini" v-if="scope.row.type == 'd'">上传</el-button>
                            <el-button @click="ftpModShow(scope.row)"  size="mini">权限</el-button>
                            <el-button @click="downXtermDialogHandler(scope.row)"  size="mini" v-if="scope.row.type != 'd'">下载</el-button>
                            <el-button @click="zip(scope.row)"  size="mini"  v-if="!scope.row.tar">压缩</el-button>
                            <el-button @click="unzip(scope.row)"  size="mini"  v-if="scope.row.tar">解压</el-button>
                            <el-button @click="sftpDelete(scope.row.path)"  size="mini">删除</el-button>
                        </el-button-group>
                    </template>
                </el-table-column>
            </el-table>
        </el-dialog>

        <!-- 权限设置 -->
        <el-dialog :visible.sync="ftpDialog.mod.visible" >
            <el-form >
                <el-form-item label="权限" >
                    <el-input v-model="ftpDialog.mod.mode" autocomplete="off"></el-input>
                </el-form-item>
            </el-form>
            <div slot="footer" class="dialog-footer">
                <el-button @click="ftpDialog.mod.visible = false">取 消</el-button>
                <el-button type="primary" @click="ftpMod()">确 定</el-button>
            </div>
        </el-dialog>

        <!-- 重命名 -->
        <el-dialog :visible.sync="ftpDialog.rename.visible" >
            <el-form >
                <el-form-item label="名字" >
                    <el-input v-model="ftpDialog.rename.name" autocomplete="off"></el-input>
                </el-form-item>
            </el-form>
            <div slot="footer" class="dialog-footer">
                <el-button @click="ftpDialog.rename.visible = false">取 消</el-button>
                <el-button type="primary" @click="ftpRename()">确 定</el-button>
            </div>
        </el-dialog>

        <!-- 目录创建 -->
        <el-dialog :visible.sync="ftpDialog.dir.visible" >
            <el-form >
                <el-form-item label="新建目录" >
                    <el-input v-model="ftpDialog.dir.path" autocomplete="off"></el-input>
                </el-form-item>
            </el-form>
            <div slot="footer" class="dialog-footer">
                <el-button @click="ftpDialog.dir.visible = false">取 消</el-button>
                <el-button type="primary" @click="ftpCreateDir()">确 定</el-button>
            </div>
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
import { 
    listServerApi, listFtpApi, deleteFtpApi, createDirFtpApi, zipFtpApi, unzipFtpApi, 
    renameFtpApi, modFtpApi, downFtpApi, uploadFtpApi } from '@/api/server'
import ServerXtermPlan from './ServerXtermPlan.vue'
import util from '@/lib/util.js'
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
                sftp_upload_percentage:0,
                title:"",
                dialogTableVisible: false,
                gridData:[],
                path: "",
                dir:{
                    visible: false,
                    path: "",
                },
                rename:{
                    visible: false,
                    name: "",
                    row:null,
                },
                mod:{
                    visible: false,
                    row:null,
                    mode:"",
                }
            }
        }
    },
    methods: {
        zip(row){
            this.$confirm('确认要压缩该文件吗?', '提示', {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                type: 'warning'
            }).then(() => {
                zipFtpApi({
                    id: this.ftpDialog.id,
                    path: row.path, 
                    new_path: row.path + '.' + (new Date()).getTime(), 
                }).then(res=>{
                    this.sftpList(this.ftpDialog.path)
                    this.$message({
                        type: 'success',
                        message: '压缩成功'
                    });
                })
            })
        },
        unzip(row){
            this.$confirm('确认要压缩该文件吗?', '提示', {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                type: 'warning'
            }).then(() => {
                unzipFtpApi({
                    id: this.ftpDialog.id,
                    path: row.path, 
                    new_path: row.path + '.' + (new Date()).getTime(), 
                }).then(res=>{
                    this.sftpList(this.ftpDialog.path)
                    this.$message({
                        type: 'success',
                        message: '解压成功'
                    });
                })
            })
        },
        determineFileType(path){
            let fileN = path.split('.')
            let ext =  fileN[fileN.length - 1];
            return util.determineFileType(ext, 'compress');
        },
        uploadFile(path) {
            this.sftp_upload_percentage = 0
            let _this = this
            function upload(fileList) {
                let formData = new FormData();
                formData.append("id", _this.ftpDialog.id);
                formData.append("path", path);
                for (let i = 0; i < fileList.length; i++) {
                    formData.append("files", fileList[i]);
                }
                uploadFtpApi(formData, (progressEvent)=>{
                    const { loaded, total } = progressEvent;
                    if (!total) {
                        // 没有获取到总大小，可能是流式上传或者chunked传输
                        _this.sftp_upload_percentage = loaded;
                    } else {
                        // 计算进度，可以用 loaded / total 得到一个0到1的数字
                        _this.sftp_upload_percentage = loaded / total * 100 | 0;
                    }
                }).then(res=>{
                    this.sftpList(path)
                    this.$message({
                        type: 'success',
                        message: res.msg
                    });
                })
            }

            let fileInput = document.createElement("input");
            fileInput.type = "file";
            fileInput.multiple = true;

            fileInput.onchange = function (f) {
                let fileList = fileInput.files;
                upload(fileList);
            };
            fileInput.click();  
        },
        downXtermDialogHandler(row){
            downFtpApi({
                id: this.ftpDialog.id,
                path: row.path
            }).then((data)=>{
                let blob = new Blob([data.data], { type: 'application/x-download' });
                let a = document.createElement("a");
                a.style.display = 'none';
                let url = window.URL.createObjectURL(blob);
                a.href = url;
                a.download = row.path;
                document.body.appendChild(a);
                a.click();
                document.body.removeChild(a); 
                window.URL.revokeObjectURL(url);
            })
        },
        ftpModShow(row){
            this.ftpDialog.mod.visible=true
            this.ftpDialog.mod.row=row
            this.ftpDialog.mod.mode=util.PermissionToOctal(row.mode)
        },
        ftpRenameShow(row){
            this.ftpDialog.rename.visible=true
            this.ftpDialog.rename.row=row
        },
        ftpMod(){
            let path = this.ftpDialog.mod.row.path;
            let mode = this.ftpDialog.mod.mode;
            this.$confirm('设置目录权限'+path+'：'+mode +', 是否继续?', '提示', {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                type: 'warning'
            }).then(() => {
                modFtpApi({
                    id: this.ftpDialog.id,
                    path: path,
                    mod: mode
                }).then((res)=>{
                    this.ftpDialog.mod.row  = null
                    this.ftpDialog.mod.mode  = ""
                    this.ftpDialog.mod.visible  = false
                    this.sftpList(this.ftpDialog.path)
                    this.$message({
                        type: 'success',
                        message: '成功!'
                    });
                });
            }).catch((e) => {
                this.$message({
                    type: 'info',
                    message: '已取消'
                });          
            });
        },
        ftpRename(){
            let path = this.ftpDialog.rename.row.path;
            let newPpath = this.ftpDialog.rename.name;
            if(newPpath[0] != "/"){
                newPpath = this.ftpDialog.path + "/" + newPpath
            }
            this.$confirm('把是'+path+'修改成'+newPpath +', 是否继续?', '提示', {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                type: 'warning'
            }).then(() => {
                renameFtpApi({
                    id: this.ftpDialog.id,
                    path: path,
                    new_path: newPpath
                }).then((res)=>{
                    this.ftpDialog.rename.row  = null
                    this.ftpDialog.rename.name  = ""
                    this.ftpDialog.rename.visible  = false
                    this.sftpList(this.ftpDialog.path)
                    this.$message({
                        type: 'success',
                        message: '成功!'
                    });
                });
            }).catch((e) => {
                this.$message({
                    type: 'info',
                    message: '已取消'
                });          
            });
        },
        ftpCreateDir(){
            let path = this.ftpDialog.dir.path;
            if(path[0] != "/"){
                path = this.ftpDialog.path + "/" + path
            }
            this.$confirm('新目录是'+path+', 是否继续?', '提示', {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                type: 'warning'
            }).then(() => {
                createDirFtpApi({
                    id: this.ftpDialog.id,
                    path: path
                }).then((res)=>{
                    this.ftpDialog.dir.path = ""
                    this.sftpList(this.ftpDialog.path)
                    this.$message({
                        type: 'success',
                        message: '成功!'
                    });
                });
            }).catch((e) => {
                this.$message({
                    type: 'info',
                    message: '已取消'
                });          
            });
        },
        sftpDelete(path){
             this.$confirm('此操作将永久删除该文件, 是否继续?', '提示', {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                type: 'warning'
            }).then(() => {
                deleteFtpApi({
                    id: this.ftpDialog.id,
                    path: path
                }).then((res)=>{
                    this.sftpList(this.ftpDialog.path)
                    this.$message({
                        type: 'success',
                        message: '删除成功!'
                    });
                });

            }).catch((e) => {
                console.log(e)
                this.$message({
                    type: 'info',
                    message: '已取消删除'
                });          
            });
        },
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
                res.files.forEach(item=>item.tar = this.determineFileType(item.path))

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
