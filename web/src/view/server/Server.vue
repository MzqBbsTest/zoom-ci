<template>
    <div>
        <el-card shadow="never">
            <el-row class="app-btn-group">
                <el-col :span="4">
                    <el-button v-if="$root.CheckPriv($root.Priv.SERVER_NEW)" @click="openAddDialogHandler" type="primary" size="medium" icon="iconfont left small icon-add">{{ $t('add_server') }}</el-button>&nbsp;
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
                <el-table-column prop="ip" width="200" label="IP/HOST"></el-table-column>
                <el-table-column prop="ssh_port" width="100" label="SSH Port"></el-table-column>
                <el-table-column :label="$t('operate')" width="280" align="right">
                    <template slot-scope="scope">
                        <el-button
                        v-if="$root.CheckPriv($root.Priv.SERVER_EDIT)"
                        icon="el-icon-edit"
                        type="text"
                        @click="openEditDialogHandler(scope.row)">{{ $t('edit') }}</el-button>
                        
                        <el-button
                        v-if="$root.CheckPriv($root.Priv.SERVER_EDIT)"
                        type="text"
                        icon="el-icon-ssh-key"
                        class="app-danger"
                        @click="testConnect(scope.row)">{{ $t('server_connect') }}</el-button>


                        <el-button
                        v-if="$root.CheckPriv($root.Priv.SERVER_EDIT)"
                        icon="el-icon-warning"
                        type="text"
                        @click="openXtermDialogHandler(scope.row)"
                        >{{ $t("command") }}</el-button>

                        <el-button
                        v-if="$root.CheckPriv($root.Priv.SERVER_DEL)"
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

        <ServerDialog @load-table-data="loadTableData"/>

        <ServerXterm />
    </div>
</template>

<script>
import { listServerApi, testServerConnect } from '@/api/server'
import ServerDialog from './ServerDialog.vue'
import ServerXterm from './ServerXterm.vue'

export default {
    components: {
        ServerDialog,
        ServerXterm
    },
    data() {
        return {
            searchInput: '',
            tableData: [],
            tableLoading: false,
            serverGroupList: [],
        }
    },
    methods: {
        searchHandler() {
            this.$root.PageInit()
            this.loadTableData()
        },
        testConnect(row){
            testServerConnect({id: row.id}).then(res => {
                this.$root.MessageSuccess()
            }).catch(err => {
                console.log(err)
                // this.tableLoading = false
            })
        },
        openXtermDialogHandler(row){
            this.$root.EmitEventGlobal("openXtermDialogHandler", {group:this.group, server:row});
        },
        // openBindKeyAddDialogHandler(){
        //     this.$root.EmitEventGlobal("openBindKeyAddDialog", {});
        // },
        // openBindKeyEditDialogHandler(row){
        //     this.$root.EmitEventGlobal("openBindKeyEditDialog", row);
        // },
        openAddDialogHandler() {
            this.$root.EmitEventGlobal("openServerAddDialog", {});
        },
        openEditDialogHandler(row) {
            this.$root.EmitEventGlobal("openServerEditDialog", row);
        },
        deleteHandler(row) {
            this.$root.ConfirmDelete(() => {
                deleteServerApi({id: row.id}).then(res => {
                    this.$root.MessageSuccess()
                    this.$root.PageReset()
                    this.loadTableData()
                })
            })
        },
        currentChangeHandler() {
            this.loadTableData()
        },
        loadTableData() {
            this.tableLoading = true
            listServerApi({keyword: this.searchInput, offset: this.$root.PageOffset(), limit: this.$root.PageSize}).then(res => {
                this.tableData = res.list
                this.$root.Total = res.total
                this.tableLoading = false
            }).catch(err => {
                this.tableLoading = false
            })
        },
    },
    mounted() {
        this.$root.PageInit()
        this.loadTableData()
    }
}
</script>
