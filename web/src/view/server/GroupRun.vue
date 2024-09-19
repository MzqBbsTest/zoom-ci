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
          <!--          <el-row class="app-btn-group">-->
          <!--            <el-col :span="4">-->
          <!--              {{ $t('start_user') }}:-->
          <!--              <el-input-->
          <!--                  v-model="start_user">-->
          <!--              </el-input>-->
          <!--            </el-col>-->
          <!--            <el-col :span="24">-->
          <!--              {{ $t('start_command') }}:-->
          <!--              <el-input-->
          <!--                  v-model="start_command">-->
          <!--              </el-input>-->
          <!--            </el-col>-->
          <!--          </el-row>-->
          <el-row class="app-btn-group">
            <el-col :span="8">
              <el-button
                  @click="openAddDialogHandler()"
                  v-if="$root.CheckPriv($root.Priv.SERVER_GROUP_NEW)"
                  type="primary"
                  size="medium"
                  icon="iconfont left small ">{{ $t("start_edit") }}
              </el-button>
              <!--              <el-button-->
              <!--                  v-if="$root.CheckPriv($root.Priv.SERVER_GROUP_NEW)"-->
              <!--                  type="primary"-->
              <!--                  size="medium"-->
              <!--                  icon="iconfont left small ">{{ $t("start_all") }}-->
              <!--              </el-button>-->
              <!--              <el-button-->
              <!--                  v-if="$root.CheckPriv($root.Priv.SERVER_GROUP_NEW)"-->
              <!--                  type="danger"-->
              <!--                  size="medium"-->
              <!--                  icon="iconfont left small ">{{ $t("stop_all") }}-->
              <!--              </el-button>&nbsp;-->
            </el-col>
          </el-row>
          <el-table
              class="app-table"
              size="medium"
              v-loading="tableLoading"
              :data="tableData">
            <el-table-column prop="id" label="ID" width="80"></el-table-column>
            <el-table-column prop="name" :label="$t('name')"></el-table-column>
            <el-table-column prop="start_command" :label="$t('start_command')"></el-table-column>
            <el-table-column :label="$t('operate')" width="380" align="right">
              <template slot-scope="scope">
                <el-button
                    v-if="$root.CheckPriv($root.Priv.SERVER_GROUP_EDIT)"
                    icon="el-icon-edit"
                    type="text"
                    @click="openEditDialogHandler(scope.row)">{{ $t("edit") }}
                </el-button>
                <el-button
                    v-if="$root.CheckPriv($root.Priv.SERVER_GROUP_EDIT)"
                    icon="el-icon-warning"
                    type="text"
                    @click="startHandler(scope.row)">{{ $t("start") }}
                </el-button>
                <!--                <el-button-->
                <!--                    v-if="$root.CheckPriv($root.Priv.SERVER_GROUP_DEL)"-->
                <!--                    type="text"-->
                <!--                    icon="el-icon-warning"-->
                <!--                    class="app-danger"-->
                <!--                    @click="stopHandler(scope.row)">{{ $t("stop") }}-->
                <!--                </el-button>-->
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
            :rules="rules"
            label-width="80px"
        >
          <el-form-item :label="$t('start_user')" prop="start_user">
            <el-input v-model="dialogForm.start_user" autocomplete="off"></el-input>
          </el-form-item>

          <el-form-item :label="$t('server_id')" prop="server_id">
            <el-select v-model="dialogForm.server_id" placeholder="请选择" @change="changeServer">
              <el-option v-for="item in options" :key="item.id" :label="item.name" :value="item.id"></el-option>
            </el-select>
          </el-form-item>

          <el-form-item :label="$t('group_path_alias')" prop="server_id">
            <el-select v-model="dialogForm.group_path_alias" placeholder="请选择" @change="generateCommand">
              <el-option v-for="item in group_path_alias_list" :key="item.alias" :label="item.alias"
                         :value="item.alias"></el-option>
            </el-select>
          </el-form-item>

          <el-form-item :label="$t('group_config_alias')" prop="server_id">
            <el-select v-model="dialogForm.group_config_alias" placeholder="请选择" @change="generateCommand">
              <el-option v-for="item in group_config_alias_list" :key="item.alias" :label="item.alias"
                         :value="item.alias"></el-option>
            </el-select>
          </el-form-item>

          <el-form-item :label="$t('start_command')" prop="start_command">
            <el-input type="textarea" v-model="dialogForm.start_command" autocomplete="off"></el-input>
            <p>###应用别名### + 服务器，会定位对应服务器的应用路径，然后使用真实的路径替换掉别名，并在对应服务器上执行。</p>
          </el-form-item>

        </el-form>
        <div slot="footer" class="dialog-footer">
          <el-button size="small" @click="dialogCloseHandler">{{ $t("cancel") }}</el-button>
          <el-button :loading="btnLoading" size="small" type="primary" @click="dialogSubmitHandler">{{ $t("enter") }}
          </el-button>
        </div>
      </div>
    </el-dialog>

    <!--    <ServerXterm ></ServerXterm>-->
  </div>
</template>

<script>
import {listGroupConfigAliasApi, listGroupPathAliasApi, listGroupRunApi, saveGroupRunApi, startGroupRunApi} from "@/api/server";
// import ServerXterm from './ServerXterm'
export default {
  components: {
    // ServerXterm,
  },
  data() {
    return {
      rules: {
        start_command: [{required: true, message: this.$t('start_command_cannot_empty'), trigger: 'blur',},],
        start_user: [{required: true, message: this.$t('start_user_cannot_empty'), trigger: 'blur',},],
        server_id: [{required: true, message: this.$t('server_cannot_empty'), trigger: 'blur',},],
      },
      dialogForm: {
        id: 0,
        start_user: "root",
        start_command: "",
        server_id: 0,
        group_id: 0,
        group_config_alias: "",
        group_path_alias: "",
      },
      group: null,
      dialogTableVisible: false,
      dialogVisible: false,
      dialogTitle: "",
      dialogLoading: false,
      btnLoading: false,
      tableData: [],
      tableLoading: false,
      group_path_alias_list: [],
      group_config_alias_list: [],
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
    generateCommand() {
      this.dialogForm.start_command = ""
      if (this.dialogForm.group_path_alias.length) {
        this.dialogForm.start_command += `###${this.dialogForm.group_path_alias}### `
      }
      if (this.dialogForm.group_config_alias.length) {
        this.dialogForm.start_command += `###${this.dialogForm.group_config_alias}###`
      }
    },
    changeServer() {
      // 别名
      listGroupConfigAliasApi({
        group_id: this.dialogForm.group_id,
        server_id: this.dialogForm.server_id,
      }).then((res) => {
        this.group_config_alias_list = res.list
      })

      listGroupPathAliasApi({
        group_id: this.dialogForm.group_id,
        server_id: this.dialogForm.server_id,
      }).then((res) => {
        this.group_path_alias_list = res.list
      })

      this.group_config_alias = ""
      this.group_path_alias = ""
    },

    dialogSubmitHandler() {
      let vm = this;
      this.$refs.dialogRef.validate((valid) => {
        if (!valid) {
          return false;
        }
        this.btnLoading = true;

        saveGroupRunApi(this.dialogForm).then((res) => {
          this.$root.MessageSuccess(() => {
            this.dialogCloseHandler();
            this.btnLoading = false;
            this.loadTableData();
          });
        }).catch((err) => {
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
      this.dialogForm.group_id = group.group.id
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

      this.dialogVisible = true;
      this.dialogTitle = this.$t("run");
      this.changeServer();

    },

    openEditDialogHandler(row) {
      // this.dialogPathCloseHandler();
      this.dialogVisible = true;
      this.dialogTitle = this.$t("run");


      this.dialogForm.start_user = row.start_user;
      this.dialogForm.start_command = row.start_command;
      this.dialogForm.server_id = row.server_id;
      this.dialogForm.group_id = row.group_id;
      this.dialogForm.group_config_alias = row.group_config_alias;
      this.dialogForm.group_path_alias = row.group_path_alias;
      this.dialogForm.id = row.id;

      this.changeServer();
    },

    loadTableData() {

      this.tableLoading = true;
      listGroupRunApi({
        group_id: this.group.id
      }).then((res) => {
        this.tableData = res.list;
        this.$root.Total = res.total;
        this.$root.PageInit()
        this.tableLoading = false;
      }).catch((err) => {
        this.tableLoading = false;
      });

    },

    openXtermDialogHandler(row) {
      // this.$root.EmitEventGlobal("openXtermDialogHandler", {group:this.group, server:row});
    },
    startHandler(row) {
      startGroupRunApi({
        id: row.id
      }).then((res) => {
        this.$root.MessageSuccess();
      });
    }
  },
  mounted() {
    this.$root.PageInit();

    this.$root.BindEventGlobal("openRunDialogHandler", this.openRunDialogHandler);
  },
  beforeDestroy() {
    this.$root.UnBindEventGlobal("openRunDialogHandler");
  }
};
</script>
