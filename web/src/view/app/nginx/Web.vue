<template>
    <div>
        <el-card shadow="never">
            <el-row class="app-btn-group">
                <el-col :span="4">
                    <el-button v-if="$root.CheckPriv($root.Priv.SERVER_WEB_NEW)" @click="openAddDialogHandler" type="primary" size="medium" icon="iconfont left small icon-add">{{ $t('add_nginx') }}</el-button>&nbsp;
                </el-col>
                <el-col :span="6" :offset="14">
                    <el-input @keyup.enter.native="searchHandler" v-model="searchInput" size="medium" :placeholder="$t('please_input_keyword_id_or_name')">
                        <el-button @click="searchHandler" slot="append" icon="el-icon-search"></el-button>
                    </el-input>
                </el-col>
            </el-row>
            <el-table
                class="app-table"
                size="medium"
                v-loading="tableLoading"
                :data="tableData">
                <el-table-column prop="id" label="ID" width="80"></el-table-column>
                <el-table-column prop="name" :label="$t('name')"></el-table-column>
                <el-table-column  :label="$t('status')">
                    <template slot-scope="scope">
                        <el-switch v-model="scope.row.status"></el-switch>
                    </template>
                </el-table-column>
                <el-table-column :label="$t('operate')" width="180" align="right">
                    <template slot-scope="scope">
                        <el-button
                        v-if="$root.CheckPriv($root.Priv.SERVER_WEB_EDIT)"
                        icon="el-icon-edit"
                        type="text"
                        @click="openEditDialogHandler(scope.row)">{{ $t('edit') }}</el-button>

                        <el-button
                        v-if="$root.CheckPriv($root.Priv.SERVER_WEB_DEL)"
                        type="text"
                        icon="el-icon-delete"
                        class="app-danger"
                        @click="deleteHandler(scope.row)">{{ $t('delete') }}</el-button>
                    </template>
                </el-table-column>
            </el-table>
            <el-pagination
                background
                layout="prev, pager, next"
                class="app-pagination"
                @current-change="currentChangeHandler"
                :current-page.sync="$root.Page"
                :page-size="$root.PageSize"
                :total="$root.Total">
            </el-pagination>
        </el-card>

        <el-dialog :width="$root.DialogSmallWidth" :title="dialogTitle" :visible.sync="dialogVisible" @close="dialogCloseHandler">
            <div class="app-dialog" v-loading="dialogLoading">
                <el-form ref="dialogRef" :model="dialogForm" size="medium" label-width="80px">
                    <el-tabs type="border-card">
                        <!-- 域名管理 -->
                        <el-tab-pane :label="$t('web_hosts')">
                            <el-table
                            :data="dialogForm.hosts"
                            style="width: 100%">
                                <el-table-column
                                    prop="date"
                                    :label="$t('host')"
                                    width="180">
                                </el-table-column>
                            </el-table>
                        </el-tab-pane>
                        <!-- 目录 -->
                        <el-tab-pane label="$t('web_path')">
                            <el-form-item :label="$t('web_work_path')">
                                <el-input v-model="dialogForm.work_path"></el-input>
                            </el-form-item>
                            <el-form-item :label="$t('web_exec_path')">
                                <el-input v-model="dialogForm.exec_path"></el-input>
                            </el-form-item>
                        </el-tab-pane>
                        <!-- 访问限制 -->
                        <el-tab-pane label="$t('web_limit')">
                            <el-form-item :label="$t('web_limit_black')">
                                <el-input v-model="dialogForm.limit_black"></el-input>
                            </el-form-item>
                        </el-tab-pane>
                        <!-- 流量限制 -->
                        <el-tab-pane label="$t('web_limit_flow')">
                            <!-- 是否开启 -->
                            <el-form-item label="$t('web_limit_flow_en')">
                                <el-switch v-model="dialogForm.web_limit_flow_en"></el-switch>
                            </el-form-item>
                            <!-- 并发 -->
                            <el-form-item label="$t('web_limit_flow_concurrent')">
                                <el-input v-model="dialogForm.web_limit_flow_concurrent"></el-input>
                            </el-form-item>
                            <!-- ip -->
                            <el-form-item label="$t('web_limit_flow_ip')">
                                <el-input v-model="dialogForm.web_limit_flow_ip"></el-input>
                            </el-form-item>
                            <!-- 流量 -->
                            <el-form-item :label="$t('web_limit_flow')">
                                <el-input v-model="dialogForm.web_limit_flow"></el-input>
                            </el-form-item>
                        </el-tab-pane>
                        <!-- 伪静态 -->
                        <el-tab-pane label="$t('web_pseudo_static')">
                            <el-form-item label="$t('web_pseudo_static_template')">
                                <el-select @change="setWebPseudoStatic" >
                                    <el-option :label="item" :value="item" v-for="item in Object.keys(pseudoStatic)"></el-option>
                                </el-select>
                            </el-form-item>
                            <el-form-item label="$t('web_pseudo_static')">
                                <el-input type="textarea"  v-model="dialogForm.web_pseudo_static"></el-input>
                            </el-form-item>
                        </el-tab-pane>
                        <!-- 默认文档 -->
                        <el-tab-pane label="$t('web_default_index')">
                            <el-form-item label="$t('web_default_index')">
                                <el-input type="textarea"  v-model="dialogForm.web_default_index"></el-input>
                            </el-form-item>
                        </el-tab-pane>
                        <!-- 配置文件 -->
                        <el-tab-pane label="$t('web_server_config')">
                            <el-form-item label="$t('web_server_config')">
                                <el-input type="textarea"  v-model="dialogForm.web_server_config"></el-input>
                            </el-form-item>
                        </el-tab-pane>

                        <!-- SSL -->
                        <el-tab-pane label="$t('web_server_ssl')">
                            <el-form-item label="$t('web_server_ssl')">
                                <el-descriptions title="">
                                    <!-- 到期时间 -->
                                    <el-descriptions-item :label="$t('web_server_ssl_end_date')">{{dialogForm.web_server_ssl.end_date}}</el-descriptions-item>
                                    <!-- 密钥(KEY) -->
                                    <el-descriptions-item :label="$t('web_server_ssl_key')">
                                        <el-input type="textarea"  v-model="dialogForm.web_server_ssl.key"></el-input>
                                    </el-descriptions-item>
                                    <!-- 证书(PEM格式) -->
                                    <el-descriptions-item :label="$t('web_server_ssl_pem')">
                                        <el-input type="textarea"  v-model="dialogForm.web_server_ssl.pem"></el-input>
                                    </el-descriptions-item>
                                </el-descriptions>
                            </el-form-item>
                        </el-tab-pane>

                        <!-- php -->
                        <el-tab-pane label="$t('web_php')">
                            <el-form-item label="$t('web_php')">
                                <el-input type="textarea"  v-model="dialogForm.web_php.pem"></el-input>
                            </el-form-item>
                        </el-tab-pane>

                        <!-- 重定向 -->
                        <el-tab-pane :label="$t('web_redirect')">
                            <el-table
                            :data="dialogForm.redirect"
                            style="width: 100%">
                                <el-table-column
                                    prop="from"
                                    :label="$t('web_redirect_from')"
                                    width="180">
                                </el-table-column>
                                <el-table-column
                                    prop="type"
                                    :label="$t('web_redirect_type')"
                                    width="180">
                                </el-table-column>
                                <el-table-column
                                    prop="to"
                                    :label="$t('web_redirect_to')"
                                    width="180">
                                </el-table-column>
                                <el-table-column
                                    prop="status"
                                    :label="$t('web_redirect_status')"
                                    width="180">
                                </el-table-column>
                                <el-table-column label="操作">
                                    <template slot-scope="scope">
                                        <el-button
                                        size="mini"
                                        @click="handleEdit(scope.$index, scope.row)">编辑</el-button>
                                        <el-button
                                        size="mini"
                                        type="danger"
                                        @click="handleDelete(scope.$index, scope.row)">删除</el-button>
                                    </template>
                                </el-table-column>
                            </el-table>
                        </el-tab-pane>

                        <!-- 反向代理 -->
                        <el-tab-pane :label="$t('web_proxy')">
                            <el-table
                            :data="dialogForm.proxy"
                            style="width: 100%">
                                <el-table-column
                                    prop="name"
                                    :label="$t('web_proxy_name')"
                                    width="180">
                                </el-table-column>
                                <el-table-column
                                    prop="type"
                                    :label="$t('web_proxy_path')"
                                    width="180">
                                </el-table-column>
                                <el-table-column
                                    prop="to"
                                    :label="$t('web_proxy_url')"
                                    width="180">
                                </el-table-column>
                                <el-table-column
                                    prop="cache"
                                    :label="$t('web_proxy_cache')"
                                    width="180">
                                </el-table-column>
                                <el-table-column
                                    prop="status"
                                    :label="$t('web_proxy_status')"
                                    width="180">
                                </el-table-column>
                                <el-table-column label="操作">
                                    <template slot-scope="scope">
                                        <el-button
                                        size="mini"
                                        @click="handleEdit(scope.$index, scope.row)">编辑</el-button>
                                        <el-button
                                        size="mini"
                                        type="danger"
                                        @click="handleDelete(scope.$index, scope.row)">删除</el-button>
                                    </template>
                                </el-table-column>
                            </el-table>
                        </el-tab-pane>

                        <!-- 防盗链 -->
                        <el-tab-pane label="$t('web_hock')">
                            <el-form-item label="$t('url_after')">
                                <el-input type="textarea"  v-model="dialogForm.web_hock_url_after"></el-input>
                            </el-form-item>
                            <el-form-item label="$t('web_hock_whost')">
                                <el-input type="textarea"  v-model="dialogForm.web_hock_whost"></el-input>
                            </el-form-item>
                            <el-form-item label="$t('web_hock_resp')">
                                <el-input type="textarea"  v-model="dialogForm.web_hock_resp"></el-input>
                            </el-form-item>

                            <el-form-item label="">
                                <!-- 启用防盗链 -->
                                <el-checkbox v-model="dialogForm.web_hock_en">112</el-checkbox>
                                <!-- 允许空HTTP_REFERER请求 -->
                                <el-checkbox v-model="dialogForm.web_hock_allow">222</el-checkbox>
                            </el-form-item>
                        </el-tab-pane>
                        
                        <!-- 网站日志 -->
                        <el-tab-pane label="$t('web_log')">
                            <template>
                                <el-tabs v-model="activeName" @tab-click="handleClick">
                                    <el-tab-pane label="$t('web_log')" name="web_log">
                                        <el-input type="textarea" v-model="dialogForm.web_log" disabled="true"></el-input>
                                    </el-tab-pane>
                                    <el-tab-pane label="$t('web_error_log')" name="web_error_log">
                                        <el-input type="textarea" v-model="dialogForm.error_log" disabled="true"></el-input>
                                    </el-tab-pane>
                                </el-tabs>
                            </template>
                        </el-tab-pane>

                    </el-tabs>
                </el-form>
                <div slot="footer" class="dialog-footer">
                    <el-button size="small" @click="dialogCloseHandler">{{ $t('cancel') }}</el-button>
                    <el-button :loading="btnLoading" size="small" type="primary" @click="dialogSubmitHandler">{{ $t('enter') }}</el-button>
                </div>
            </div>
        </el-dialog>

    </div>
</template>

<script>
import { listAPPWebApi, newAPPWebApi, updateAPPWebApi, deleteAPPWebApi, detailAPPWebApi } from '@/api/web'
import pseudoStatic from '@/lib/pseudo_static'

export default {
    data() {
        return {
            searchInput: '',
            dialogVisible: false,
            dialogTitle: '',
            pseudoStatic: pseudoStatic,
            dialogForm: {},
            dialogLoading: false,
            btnLoading: false,

            tableData: [],
            tableLoading: false,
            servers:[]
        }
    },
    methods: {
        setWebPseudoStatic(web_pseudo_static){
            this.dialogForm.web_pseudo_static = web_pseudo_static
        },
        searchHandler() {
            this.$root.PageInit()
            this.loadTableData()
        },
        openAddDialogHandler() {
            this.dialogVisible = true
            this.dialogTitle = this.$t('add_nginx')
        },
        openEditDialogHandler(row) {
            this.dialogVisible = true
            this.dialogTitle = this.$t('edit_nginx')
            this.dialogLoading = true
            detailAPPWebApi({id: row.id}).then(res => {
                this.dialogLoading = false
                this.dialogForm = res
            }).catch(err => {
                this.dialogCloseHandler()
            })
        },
        dialogCloseHandler() {
            this.dialogVisible = false
            this.dialogLoading = false
            this.btnLoading = false
            this.$refs.dialogRef.resetFields();
            this.dialogForm = {}
        },
        deleteHandler(row) {
            this.$root.ConfirmDelete(() => {
                deleteAPPWebApi({id: row.id}).then(res => {
                    this.$root.MessageSuccess()
                    this.$root.PageReset()
                    this.loadTableData()
                })
            })
        },
        currentChangeHandler() {
            this.loadTableData()
        },
        dialogSubmitHandler() {
            let vm = this
            this.$refs.dialogRef.validate((valid) => {
                if (!valid) {
                    return false;
                }
                this.btnLoading = true
                let opFn
                if (this.dialogForm.id) {
                    opFn = updateAPPWebApi
                } else {
                    opFn = newAPPWebApi
                }
                opFn(this.dialogForm).then(res => {
                    this.$root.MessageSuccess(() => {
                        this.dialogCloseHandler()
                        this.btnLoading = false
                        this.loadTableData()
                    })
                }).catch(err => {
                    this.btnLoading = false
                })
            });
        },
        loadTableData() {
            this.tableLoading = true
            listAPPWebApi({keyword: this.searchInput, offset: this.$root.PageOffset(), limit: this.$root.PageSize}).then(res => {
                this.tableData = res.list
                this.$root.Total = res.total
                this.tableLoading = false
            }).catch(err => {
                this.tableLoading = false
            })
        }
    },
    mounted() {
        this.$root.PageInit()
        this.loadTableData()
        console.log(pseudoStatic)
    }
}
</script>
