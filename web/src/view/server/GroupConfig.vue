<template>
  <div>
    <el-dialog
        :width="$root.DialogLargeWidth"
        :title="dialogTitle"
        :visible.sync="dialogTableVisible"
        @close="dialogConfigCloseHandler"
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
              >{{ $t("app_config_add") }}
              </el-button
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
            <el-table-column prop="file" :label="$t('file')"></el-table-column>
            <el-table-column :label="$t('operate')" width="380" align="right">
              <template slot-scope="scope">
                <el-button
                    v-if="$root.CheckPriv($root.Priv.SERVER_GROUP_EDIT)"
                    icon="el-icon-edit"
                    type="text"
                    @click="openEditDialogHandler(scope.row)"
                >{{ $t("edit") }}
                </el-button
                >
                <el-button
                    v-if="$root.CheckPriv($root.Priv.SERVER_GROUP_DEL)"
                    type="text"
                    icon="el-icon-delete"
                    class="app-danger"
                    @click="deleteHandler(scope.row)"
                >{{ $t("delete") }}
                </el-button
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
              :label="$t('name')"
              prop="name"
              :rules="[
              {
                required: true,
                message: $t('name_cannot_empty'),
                trigger: 'blur',
              },
            ]"
          >
            <el-input v-model="dialogForm.name" autocomplete="off"></el-input>
          </el-form-item>
          <el-form-item
              :label="$t('file_path')"
              prop="file_path"
              :rules="[
              {
                required: true,
                message: $t('file_path_cannot_empty'),
                trigger: 'blur',
              },
            ]"
          >
            <el-input v-model="dialogForm.file_path" autocomplete="off"></el-input>
          </el-form-item>

          <el-form-item
              :label="$t('file')"
              prop="file"
              :rules="[
              {
                required: true,
                message: $t('file_cannot_empty'),
                trigger: 'blur',
              },
            ]"
          >
            <el-upload
                class="upload-demo"
                action="https://jsonplaceholder.typicode.com/posts/"
                :before-remove="beforeRemove"
                :limit="3"
                :on-exceed="handleExceed"
                :file-list="fileList"
                :on-success="handleAvatarSuccess"
            >
              <el-button size="small" type="primary">点击上传</el-button>
              <div slot="tip" class="el-upload__tip">上传配置文件</div>
            </el-upload>
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
          </el-button
          >
        </div>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import {listGroupConfigApi, saveGrouConfigpApi, deleteGrouConfigpApi} from "@/api/server";

export default {
  data() {
    return {
      rules: {
        server_id: [{required: true, message: this.$t('name_cannot_empty'), trigger: 'blur',},],
        name: [{required: true, message: this.$t('name_cannot_empty'), trigger: 'blur',},],
        // path: [{required: true, message: this.$t('name_cannot_empty'), trigger: 'blur',},]
      },
      dialogForm: {
        id: 0,
        name: "",
        file_path: "",
        file: 0,
      },
      fileList: [
        //    {name: 'food2.jpeg', url: 'https://fuss10.elemecdn.com/3/63/4e7f3a15429bfda99bce42a18cdd1jpeg.jpeg?imageMogr2/thumbnail/360x360/format/webp/quality/100'}
      ],
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
  methods: {
    handleAvatarSuccess(res, file) {
      this.dialogForm.file = URL.createObjectURL(file.raw);
    },
    handleExceed(files, fileList) {
      this.$message.warning(
          `当前限制选择 1 个文件，本次选择了 ${files.length} 个文件，共选择了 ${
              files.length + fileList.length
          } 个文件`
      );
    },
    beforeRemove(file, fileList) {
      return this.$confirm(`确定移除 ${file.name}？`);
    },
    dialogSubmitHandler() {
      let vm = this;
      this.$refs.dialogRef.validate((valid) => {
        if (!valid) {
          return false;
        }
        this.btnLoading = true;
        saveGrouConfigpApi(this.dialogForm)
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
    dialogConfigCloseHandler() {
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
    openCinfigDialogHandler(group) {
      this.dialogTableVisible = true;
      this.dialogTitle = this.$t("app_path");
      this.group = group.group;
      this.dialogForm.group_id = this.group.id
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
      // this.dialogConfigCloseHandler();
      this.dialogVisible = true;
      this.dialogTitle = this.$t("app_path_add");
    },
    openEditDialogHandler(row) {
      // this.dialogConfigCloseHandler();
      this.dialogVisible = true;
      this.dialogTitle = this.$t("app_path_edit");
      this.dialogForm.path = row.path;
      this.dialogForm.server_id = row.server_id;
      this.dialogForm.id = row.id;
    },
    loadTableData() {
      this.tableLoading = true;
      listGroupConfigApi({
        group_id: this.group.id,
      })
          .then((res) => {
            console.log(res);
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
    this.$root.BindEventGlobal(
        "openCinfigDialogHandler",
        this.openCinfigDialogHandler
    );
  },
};
</script>
