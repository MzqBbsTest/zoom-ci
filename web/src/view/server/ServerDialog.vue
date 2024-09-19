<template>
  <el-dialog :width="$root.DialogLargeWidth" :title="dialogTitle" :visible.sync="dialogVisible"
             @close="dialogCloseHandler">
    <div class="app-dialog" v-loading="dialogLoading">
      <el-form class="app-form" ref="dialogRef" :model="dialogForm" size="medium" label-width="120px"
               v-if="dialogName=='server'">
        <el-form-item
            :label="$t('server_name')"
            prop="name"
            :rules="[
                        { required: true, message: this.$t('name_cannot_empty'), trigger: 'blur'},
                    ]">
          <el-input :placeholder="$t('please_input_server_name')" v-model="dialogForm.name"
                    autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item
            label="IP/HOST"
            prop="ip"
            :rules="[
                        { required: true, message: this.$t('IP_HOST_cannot_empty'), trigger: 'blur'},
                    ]">
          <el-input :placeholder="$t('please_input_server_IP_HOST')" v-model="dialogForm.ip"
                    autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item
            :label="$t('SSH_port')"
            prop="ssh_port"
            :rules="[
                        { required: true, message: this.$t('SSH_port_cannot_empty'), trigger: 'blur'},
                    ]">
          <el-input maxlength=5 class="app-input-mini" v-model="dialogForm.ssh_port" autocomplete="off"></el-input>
        </el-form-item>

        <el-form-item
            :label="$t('user')"
            prop="user"
            :rules="[
                        
                    ]">
          <el-input :placeholder="$t('user_or_ssh')" v-model="dialogForm.user" autocomplete="off"></el-input>
        </el-form-item>


        <el-form-item
            :label="$t('password')"
            prop="password"
            :rules="[
                        
                    ]">
          <el-input type="password" v-model="dialogForm.password" autocomplete="off"></el-input>
        </el-form-item>


        <el-form-item
            :label="$t('server_public_sshkey')"
            prop="sshkey_id"
            :rules="[

                    ]">
          <el-select filterable :placeholder="$t('keyword_search')" style="width: 100%"
                     v-model="dialogSshKeyForm.sshkey_id">
            <el-option
                v-for="g in sshkeyList"
                :key="g.id"
                :label="g.name"
                :value="g.id">
            </el-option>
          </el-select>
        </el-form-item>
      </el-form>


      <div slot="footer" class="dialog-footer">
        <el-button size="small" @click="dialogCloseHandler">{{ $t('cancel') }}</el-button>
        <el-button :loading="btnLoading" size="small" type="primary" @click="dialogSubmitHandler">{{ $t('enter') }}
        </el-button>
      </div>


    </div>
  </el-dialog>
</template>


<script>
import {
  detailServerApi,
  listSshkeyApi,
  newServerApi,
  newServerSshkeyApi,
  updateServerApi,
  updateServerSshkeyApi
} from '@/api/server'

export default {
  data() {
    return {
      dialogVisible: false,
      dialogTitle: '',
      dialogForm: {},
      dialogSshKeyForm: {},
      dialogLoading: false,
      btnLoading: false,
      dialogName: 'server',
      sshkeyList: []
    }
  },
  methods: {

    // openBindKeyAddDialogHandler() {
    //     this.dialogName = 'sshkey'
    //     this.dialogVisible = true
    //     this.dialogTitle = this.$t('add_server_sshkey_info')
    //     this.dialogLoading = true
    //     listSshkeyApi().then(res=>this.sshkeyList=res)
    // },

    // openBindKeyEditDialogHandler() {
    //     listSshkeyApi().then(res=>this.sshkeyList=res)
    //     this.dialogName = 'sshkey'
    //     this.dialogVisible = true
    //     this.dialogTitle = this.$t('edit_server_sshkey_info')
    //     this.dialogLoading = true
    //     detailServerSshkeyApi({id: row.id}).then(res => {
    //         this.dialogLoading = false
    //         this.dialogSshKeyForm = res
    //     }).catch(err => {
    //         this.dialogCloseHandler()
    //     })
    // },
    searchSshKey() {
      listSshkeyApi().then(res => this.sshkeyList = res.list)
    },
    openAddDialogHandler() {
      this.searchSshKey()
      this.dialogName = 'server'
      this.dialogVisible = true
      this.dialogTitle = this.$t('add_server')
    },
    openEditDialogHandler(row) {
      this.searchSshKey()
      this.dialogName = 'server'
      this.dialogVisible = true
      this.dialogTitle = this.$t('edit_server_info')
      this.dialogLoading = true

      detailServerApi({id: row.id}).then(res => {
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

      this.$refs.dialogRef && this.$refs.dialogRef.resetFields();
      this.$refs.dialogSshKeyRef && this.$refs.dialogSshKeyRef.resetFields();
      this.dialogForm = {
        ssh_port: 22,
      }
    },
    dialogSubmitHandler() {
      let vm = this
      this.$refs.dialogRef.validate((valid) => {
        if (!valid) {
          return false;
        }
        this.btnLoading = true
        if (this.dialogName == "server") {
          this.dialogForm2()
        } else {
          this.dialogFormSsKey()
        }
      });
    },
    dialogFormSsKey() {
      let opFn
      if (this.dialogFormSsKey.id) {
        opFn = updateServerSshkeyApi
      } else {
        opFn = newServerSshkeyApi
      }
      opFn(this.dialogFormSsKey).then(res => {
        this.$root.MessageSuccess(() => {
          this.dialogCloseHandler()
          this.btnLoading = false
          this.$emit("load-table-data")
        })
      }).catch(err => {
        this.btnLoading = false
      })
    },
    dialogForm2() {
      let opFn
      if (this.dialogForm.id) {
        opFn = updateServerApi
      } else {
        opFn = newServerApi
      }
      opFn(this.dialogForm).then(res => {
        this.$root.MessageSuccess(() => {
          this.dialogCloseHandler()
          this.btnLoading = false
          this.$emit("load-table-data")
        })
      }).catch(err => {
        this.btnLoading = false
      })
    }

  },
  mounted() {
    this.$root.BindEventGlobal("openServerAddDialog", this.openAddDialogHandler);
    this.$root.BindEventGlobal("openServerEditDialog", this.openEditDialogHandler);
    // this.$root.BindEventGlobal("openBindKeyAddDialog", this.openBindKeyAddDialogHandler);
    // this.$root.BindEventGlobal("openBindKeyEditDialog", this.openBindKeyEditDialogHandler);
  },
  beforeDestroy() {
    this.$root.UnBindEventGlobal("openServerAddDialog");
    this.$root.UnBindEventGlobal("openServerEditDialog");
  }
}
</script>
