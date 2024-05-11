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
                >{{ $t("app_path_add") }}</el-button
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
            <el-table-column :label="$t('operate')" width="380" align="right">
              <template slot-scope="scope">
                <el-button
                  v-if="$root.CheckPriv($root.Priv.SERVER_GROUP_EDIT)"
                  icon="el-icon-edit"
                  type="text"
                  @click="openEditDialogHandler(scope.row)"
                  >{{ $t("edit") }}</el-button
                >
                <el-button
                  v-if="$root.CheckPriv($root.Priv.SERVER_GROUP_DEL)"
                  type="text"
                  icon="el-icon-delete"
                  class="app-danger"
                  @click="deleteHandler(scope.row)"
                  >{{ $t("delete") }}</el-button
                >
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
            :label="$t('server_id')"
            prop="server_id"
            :rules="[
              {
                required: true,
                message: $t('name_cannot_empty'),
                trigger: 'blur',
              },
            ]"
          >
            <el-select v-model="dialogForm.server_id" placeholder="请选择">
              <el-option
                v-for="item in options"
                :key="item.id"
                :label="item.name"
                :value="item.id"
              >
              </el-option>
            </el-select>
          </el-form-item>

          <el-form-item
            :label="$t('app_path')"
            prop="path"
            :rules="[
              {
                required: true,
                message: $t('name_cannot_empty'),
                trigger: 'blur',
              },
            ]"
          >
            <el-input v-model="dialogForm.path" autocomplete="off"></el-input>
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
  </div>
</template>

<script>
import { listGroupPathApi } from "@/api/server";

export default {
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
    };
  },
  computed: {
    options() {
      return [
        {
          id: 0,
          name: "全局",
        },
        {
          id: 1,
          name: "92",
        },
      ];
    },
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
          opFn = newGroupPathApi;
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
    openAppPathDialogHandler(group) {
      this.dialogTableVisible = true;
      this.dialogTitle = this.$t("app_path");
      this.group = group.group;
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
      this.dialogForm.server_id = row.server_id;
      this.dialogForm.id = row.id;
    },
    loadTableData() {
      this.tableLoading = true;
      listGroupPathApi({
        group_id: this.group.id,
        offset: this.$root.PageOffset(),
        limit: this.$root.PageSize,
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
  },
  mounted() {
    this.$root.PageInit();

    this.$root.BindEventGlobal(
      "openAppPathDialogHandler",
      this.openAppPathDialogHandler
    );
  },
};
</script>
