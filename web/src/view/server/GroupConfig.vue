<template>
    <div>
       
        <el-dialog :width="$root.DialogLargeWidth" :title="dialogTitle" :visible.sync="dialogVisible" @close="dialogCloseHandler">
            <div class="app-dialog" v-loading="dialogLoading">
                <el-form ref="dialogRef" :model="dialogForm" size="medium" label-width="80px">
                    <el-form-item 
                        :label="$t('cluster_name')"
                        prop="name"
                        :rules="[
                            { required: true, message: $t('name_cannot_empty'), trigger: 'blur'},
                        ]">
                        <el-input v-model="dialogForm.name" autocomplete="off"></el-input>
                    </el-form-item>

                    <el-form-item 
                        :label="$t('cluster_servers')"
                        prop="servers">
                        <el-transfer 
                        filterable 
                        :filter-method="filterMethod" 
                        v-model="dialogForm.servers" 
                        :data="servers"
                        target-order=""
                        :titles="[$t('servers'), $t('cluster_servers')]"
                        ></el-transfer>
                    </el-form-item>
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
import { 
    detailGrouConfigpApi,updateGrouConfigpApi
} from '@/api/server'

export default {
    data() {
        return {
            group: null,
            dialogVisible: false,
            dialogTitle: '',
            dialogForm: {},
            dialogSshKeyForm: {},
            dialogLoading: false,
            btnLoading: false,
            dialogName: 'server',

        }
    },
    methods: {
        
        openCinfigDialogHandler(group) {
            this.group = group.group;
            detailGrouConfigpApi({
                id: this.group.id
            }).then(res=>this.dialogForm=res)
            this.dialogName = 'config'
            this.dialogVisible = true
            this.dialogTitle = this.$t('app_path_config')
            this.dialogLoading = true
        },
        
        dialogCloseHandler() {
            this.dialogVisible = false
            this.dialogLoading = false
            this.btnLoading = false
            this.$refs.dialogRef.resetFields();
        },
        dialogSubmitHandler() {
            let vm = this
            this.$refs.dialogRef.validate((valid) => {
                if (!valid) {
                    return false;
                }
                this.btnLoading = true
                if (this.dialogName == "server") {
                    this.dialogForm()
                } else {
                    this.dialogFormSsKey()
                }
            });
        },
       
        dialogForm(){
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
        this.$root.BindEventGlobal("openCinfigDialogHandler", this.openCinfigDialogHandler);

    }
}
</script>
