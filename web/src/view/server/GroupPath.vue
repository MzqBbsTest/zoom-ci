<template>
  <div>
    <el-dialog
        :width="$root.DialogLargeWidth"
        :title="dialogTitle"
        :visible.sync="dialogTableVisible"
        @close="dialogPathCloseHandler"
    >
      <div class="app-dialog" v-loading="dialogLoading">
        <el-card shadow="never">
          <el-row class="app-btn-group">
            <el-col :span="4">
              <el-button
                  v-if="$root.CheckPriv($root.Priv.SERVER_GROUP_NEW)"
                  @click="openAddDialogHandler"
                  type="primary"
                  size="medium"
                  icon="iconfont left small icon-add"
              >{{ $t("app_path_add") }}
              </el-button
              >&nbsp;
            </el-col>
          </el-row>
          <el-table
              class="app-table"
              size="medium"
              v-loading="tableLoading"
              :data="tableData">
            <el-table-column prop="id" label="ID" width="80"></el-table-column>
            <el-table-column prop="name" :label="$t('name')"></el-table-column>
            <el-table-column prop="path" :label="$t('path')"></el-table-column>
            <el-table-column prop="server_id" :label="$t('server_id')">
              <template slot-scope="scope">{{ servers[scope.row.server_id] }}</template>
            </el-table-column>
            <el-table-column :label="$t('operate')" width="380" align="right">
              <template slot-scope="scope">
                <el-button
                    v-if="$root.CheckPriv($root.Priv.SERVER_GROUP_EDIT)"
                    icon="el-icon-edit"
                    type="text"
                    @click="openEditDialogHandler(scope.row)">{{ $t("edit") }}
                </el-button>
                <el-button
                    v-if="$root.CheckPriv($root.Priv.SERVER_GROUP_DEL)"
                    type="text"
                    icon="el-icon-delete"
                    class="app-danger"
                    @click="deleteHandler(scope.row)">{{ $t("delete") }}
                </el-button>
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
              :total="$root.Total"
          >
          </el-pagination>
        </el-card>
      </div>
    </el-dialog>

    <el-dialog :width="$root.DialogLargeWidth" :title="dialogTitle" :visible.sync="dialogVisible"
               @close="dialogCloseHandler">
      <div class="app-dialog" v-loading="dialogLoading">
        <el-form
            ref="dialogRef"
            :model="dialogForm"
            size="medium"
            :rules="rules"
            label-width="80px"
        >

          <el-form-item :label="$t('server_id')" prop="server_id">
            <el-select v-model="dialogForm.server_id" placeholder="请选择">
              <el-option v-for="item in options" :key="item.id" :label="item.name" :value="item.id"></el-option>
            </el-select>
          </el-form-item>

          <el-form-item :label="$t('app_name')" prop="name">
            <el-input v-model="dialogForm.name" autocomplete="off"></el-input>
          </el-form-item>

          <el-form-item :label="$t('app_path')" prop="path">
            <el-input v-model="dialogForm.path" autocomplete="off"></el-input>
          </el-form-item>
        </el-form>
        <div slot="footer" class="dialog-footer">
          <el-button size="small" @click="dialogCloseHandler">{{
              $t("cancel")
            }}
          </el-button>
          <el-button
              :loading="btnLoading"
              size="small"
              type="primary"
              @click="dialogSubmitHandler"
          >{{ $t("enter") }}
          </el-button>
        </div>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import {deleteGroupPathApi, listGroupPathApi, saveGroupPathApi} from "@/api/server";

export default {
  data() {
    return {
      rules: {
        server_id: [{required: true, message: this.$t('server_cannot_empty'), trigger: 'blur',},],
        name: [{required: true, message: this.$t('name_cannot_empty'), trigger: 'blur',},],
        alias: [{required: true, message: this.$t('alias_cannot_empty'), trigger: 'blur',},],
        path: [{required: true, message: this.$t('path_cannot_empty'), trigger: 'blur',},]
      },
      dialogForm: {
        id: 0,
        name: "",
        alias: "",
        path: "",
        server_id: 0,
        group_id: 0,
      },
      group: null,
      dialogTableVisible: false,
      dialogVisible: false,
      dialogTitle: "",
      dialogLoading: false,
      btnLoading: false,
      tableData: [],
      tableLoading: false,
      options: [
        {
          id: 0,
          name: "全局",
        },
      ],
      servers: {"0": "全局"}
    };
  },
  computed: {},
  methods: {
    dialogSubmitHandler() {
      let vm = this;
      this.$refs.dialogRef.validate((valid) => {
        if (!valid) {
          return false;
        }
        this.btnLoading = true;
        saveGroupPathApi(this.dialogForm)
            .then((res) => {
              this.$root.MessageSuccess(() => {
                this.dialogCloseHandler();
                this.btnLoading = false;
                this.loadTableData();
              });
            })
            .catch((err) => {
              this.btnLoading = false;
            });
      });
    },
    dialogPathCloseHandler() {
      this.dialogTableVisible = false;
      this.dialogLoading = false;
      this.btnLoading = false;
    },
    dialogCloseHandler() {
      this.dialogTableVisible = true;
      this.dialogVisible = false;
      this.dialogLoading = false;
      this.btnLoading = false;
      this.dialogForm.name = ""
      this.dialogForm.id = 0
      this.dialogForm.server_id = 0
      this.dialogForm.path = ""
    },
    currentChangeHandler() {
      this.loadTableData();
    },
    openAppPathDialogHandler(group) {
      this.dialogTableVisible = true;
      this.dialogTitle = this.$t("app_path");
      this.group = group.group;
      this.dialogForm.group_id = this.group.id

      this.options = [{id: 0, name: "全局",},]
      group.servers.forEach(res => {
        this.options.push({
          id: res.key,
          name: res.label
        })
        this.servers[res.key] = res.label
      })

      this.loadTableData();
    },
    deleteHandler(row) {
      this.$root.ConfirmDelete(() => {
        deleteGroupPathApi({id: row.id}).then((res) => {
          this.$root.MessageSuccess();
          this.$root.PageReset();
          this.loadTableData();
        });
      });
    },
    openAddDialogHandler() {
      // this.dialogPathCloseHandler();
      this.dialogVisible = true;
      this.dialogTitle = this.$t("app_path_add");
    },
    openEditDialogHandler(row) {
      // this.dialogPathCloseHandler();
      this.dialogVisible = true;
      this.dialogTitle = this.$t("app_path_edit");
      this.dialogForm.path = row.path;
      this.dialogForm.name = row.name;
      this.dialogForm.server_id = row.server_id;
      this.dialogForm.id = row.id;
    },
    async loadTableData() {
      this.tableLoading = true;

      // await listServerApi({group_id: this.group.id,}).then((res) => {
      //   res.list.forEach(res=>{
      //     this.options.push({
      //       id: res.id,
      //       name: res.name
      //     })
      //     this.servers[res.id] = res.name
      //   })
      // })

      await listGroupPathApi({
        group_id: this.group.id,
        offset: this.$root.PageOffset(),
        limit: this.$root.PageSize,
      }).then((res) => {
        this.tableData = res.list;
        this.$root.Total = res.total;
        this.$root.PageInit()
        this.tableLoading = false;
      }).catch((err) => {
        this.tableLoading = false;
      });

    },
  },
  mounted() {
    this.$root.PageInit();
    this.$root.BindEventGlobal("openAppPathDialogHandler", this.openAppPathDialogHandler);
  },
  beforeDestroy() {
    this.$root.UnBindEventGlobal("openAppPathDialogHandler");
  }
};
</script>
