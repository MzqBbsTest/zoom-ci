<template>
    <el-dialog :width="$root.DialogLargeWidth" :title="dialogTitle" :visible.sync="dialogVisible" @close="dialogCloseHandler">
        <div class="app-dialog" v-loading="dialogLoading">
            
            <div id="xtermDialog" ref="terminal" ></div>

        </div>
    </el-dialog>
</template>



<script>
import { 

} from '@/api/server'

import { Terminal } from 'xterm';
import { FitAddon } from 'xterm-addon-fit';
// import '@xterm/xterm/css/xterm.css';
import Vue from 'vue'
export default {
    data() {
        return {
            dialogVisible: false,
            dialogLoading: false,
            btnLoading: false,
            dialogTitle:""
        }
    },
    methods: {
        openXtermDialogHandler() {
            this.dialogVisible = true
            Vue.nextTick(function () {
                const term = new Terminal();
                const fitAddon = new FitAddon();
                term.loadAddon(fitAddon);
                term.open(this.$refs.terminal);
                fitAddon.fit();
            }.bind(this))
        },
        
        dialogCloseHandler() {
            this.dialogVisible = false
            this.dialogLoading = false
            this.btnLoading = false
        },
    
    },
    mounted() {
        this.$root.BindEventGlobal("openXtermDialogHandler", this.openXtermDialogHandler);
    }
}
</script>
