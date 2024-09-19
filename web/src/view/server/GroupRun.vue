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
              {{ $t('start_user') }}:
              <el-input
                v-model="start_user">
              </el-input>
            </el-col>
            <el-col :span="24">
              {{ $t('start_command') }}:
              <el-input
                v-model="start_command">
              </el-input>
            </el-col>
          </el-row>
          <el-row class="app-btn-group">
            <el-col :span="8">
              <el-button
                v-if="$root.CheckPriv($root.Priv.SERVER_GROUP_NEW)"
                type="primary"
                size="medium"
                icon="iconfont left small "
                >{{ $t("start_all") }}</el-button
              >
              <el-button
                v-if="$root.CheckPriv($root.Priv.SERVER_GROUP_NEW)"

                type="danger"
                size="medium"
                icon="iconfont left small "
                >{{ $t("stop_all") }}</el-button
              >&nbsp;
            </el-col>
          </el-row>
          <el-table
            class="app-table"
            size="medium"
            v-loading="tableLoading"
            :data="tableData"
          >
            <el-table-column prop="id" label="ID" width="80"></el-table-column>
            <el-table-column prop="name" :label="$t('name')"></el-table-column>
            <el-table-column prop="status" :label="$t('status')"></el-table-column>
            <el-table-column :label="$t('operate')" width="380" align="right">
              <template slot-scope="scope">
                <el-button
                  v-if="$root.CheckPriv($root.Priv.SERVER_GROUP_EDIT)"
                  icon="el-icon-edit"
                  type="text"
                  @click="openEditDialogHandler(scope.row)"
                  >{{ $t("start_edit") }}</el-button
                >
                <el-button
                  v-if="$root.CheckPriv($root.Priv.SERVER_GROUP_EDIT)"
                  icon="el-icon-warning"
                  type="text"
                  @click="startHandler(scope.row)"
                  >{{ $t("start") }}</el-button
                >
                <el-button
                  v-if="$root.CheckPriv($root.Priv.SERVER_GROUP_DEL)"
                  type="text"
                  icon="el-icon-warning"
                  class="app-danger"
                  @click="stopHandler(scope.row)"
                  >{{ $t("stop") }}</el-button
                >
<!--                <el-button-->
<!--                  v-if="$root.CheckPriv($root.Priv.SERVER_GROUP_EDIT)"-->
<!--                  icon="el-icon-warning"-->
<!--                  type="text"-->
<!--                  @click="openXtermDialogHandler(scope.row)"-->
<!--                  >{{ $t("command") }}</el-button-->
<!--                >-->
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

    <el-dialog
      :width="$root.DialogLargeWidth"
      :title="dialogTitle"
      :visible.sync="dialogVisible"
      @close="dialogCloseHandler"
    >
      <div class="app-dialog" v-loading="dialogLoading">
        <el-form
          ref="dialogRef"
          :model="dialogForm"
          size="medium"
          label-width="80px"
        >
          <el-form-item
            :label="$t('start_user')"
            prop="start_user"
            :rules="[
              {
                required: true,
                message: $t('start_user_cannot_empty'),
                trigger: 'blur',
              },
            ]"
          >
            <el-input v-model="dialogForm.start_user" autocomplete="off"></el-input>
          </el-form-item>

          <el-form-item
            :label="$t('start_command')"
            prop="start_command"
            :rules="[
              {
                required: true,
                message: $t('start_command_cannot_empty'),
                trigger: 'blur',
              },
            ]"
          >
            <el-input v-model="dialogForm.start_command" autocomplete="off"></el-input>

          </el-form-item>
        </el-form>
        <div slot="footer" class="dialog-footer">
          <el-button size="small" @click="dialogCloseHandler">{{
            $t("cancel")
          }}</el-button>
          <el-button
            :loading="btnLoading"
            size="small"
            type="primary"
            @click="dialogSubmitHandler"
            >{{ $t("enter") }}</el-button
          >
        </div>
      </div>
    </el-dialog>

    <ServerXterm ></ServerXterm>
  </div>
</template>

<script>
import { listGroupRunApi,addGroupPathApi,updateGroupPathApi,deleteGroupPathApi } from "@/api/server";
import ServerXterm from './ServerXterm'
export default {
  components: {
    ServerXterm,
    },
  data() {
    return {
      dialogForm: {
        id: 0,
        path: "",
        server_id: 0,
      },
      group: null,
      dialogTableVisible: false,
      dialogVisible: false,
      dialogTitle: "",
      dialogLoading: false,
      btnLoading: false,
      tableData: [],
      tableLoading: false,
      start_command: "",
      start_user: "",
      options: [
        {
          id: 0,
          name: "全局",
        },
      ],
      servers:{"0": "全局"}
    };
  },
  computed: {
  },
  methods: {
    dialogSubmitHandler() {
      alert(1);
      let vm = this;
      this.$refs.dialogRef.validate((valid) => {
        if (!valid) {
          return false;
        }
        this.btnLoading = true;
        let opFn;
        if (this.dialogForm.id) {
          opFn = updateGroupPathApi;
        } else {
          opFn = addGroupPathApi ;
        }
        opFn(this.dialogForm)
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
      this.dialogForm = {};
    },
    currentChangeHandler() {
      this.loadTableData();
    },
    openRunDialogHandler(group) {
      this.dialogTableVisible = true;
      this.dialogTitle = this.$t("run");
      this.group = group.group;
      group.servers.forEach(res=>{
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
        deleteGroupPathApi({ id: row.id }).then((res) => {
          this.$root.MessageSuccess();
          this.$root.PageReset();
          this.loadTableData();
        });
      });
    },
    
    openEditDialogHandler(row) {
      // this.dialogPathCloseHandler();
      this.dialogVisible = true;
      this.dialogTitle = this.$t("run");
      this.dialogForm.path = row.path;
      this.dialogForm.server_id = row.server_id;
      this.dialogForm.id = row.id;
    },
    loadTableData() {
      this.tableLoading = true;
      listGroupRunApi({
        group_id: this.group.id
      })
        .then((res) => {
          this.tableData = res.list;
          this.$root.Total = res.total;
          this.tableLoading = false;
        })
        .catch((err) => {
          this.tableLoading = false;
        });
    },
    openXtermDialogHandler(row){

      // this.$root.EmitEventGlobal("openXtermDialogHandler", {group:this.group, server:row});
    }
  },
  mounted() {
    this.$root.PageInit();

    this.$root.BindEventGlobal(
      "openRunDialogHandler",
      this.openRunDialogHandler
    );
  },
};
</script>
