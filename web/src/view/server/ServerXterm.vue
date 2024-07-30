<template>
    <el-dialog :width="$root.DialogLargeWidth" :title="dialogTitle" :visible.sync="dialogVisible"
        @close="dialogCloseHandler">
        <div class="app-dialog" v-loading="dialogLoading">
            <div id="xtermDialog" ref="terminal"></div>
        </div>
    </el-dialog>
</template>

<script>
import { serverSession } from "@/api/server";

import { Terminal } from "xterm";
import { FitAddon } from "xterm-addon-fit";
import { AttachAddon } from "xterm-addon-attach";
import "xterm/css/xterm.css";
import Vue from "vue";
export default {
    data() {
        return {
            socket: null,
            dialogVisible: false,
            dialogLoading: false,
            btnLoading: false,
            dialogTitle: "",
            command: "",
            sessionId :0,
        };
    },
    methods: {
        openXtermDialogHandler(params) {
            
            serverSession({
                id: 2,
                session_id: this.sessionId
            }).then((res)=>{
            
                this.sessionId =  res.session_id
                // this.sessionId = res // res.session_id
                this.dialogVisible = true;
                let _this = this    
                Vue.nextTick(()=>{
                    if (_this.term || !_this.$refs.terminal) {
                        return;
                    }
                    const term = new Terminal();
                    term.prompt = () => {
                        // term.write("\r\n\x1b[33m$\x1b[0m ");
                    };
                    term.prompt();
                    term.open(_this.$refs.terminal);
                    const fitAddon = new FitAddon();
                    term.loadAddon(fitAddon);
                    fitAddon.fit()
                    _this.term = term;
                
                    // _this.socket = new WebSocket('ws://localhost:8899/api/ssh/conn?h=48&w=95&session_id=' + res + '&Authorization=111');
                    _this.socket = new WebSocket('ws://localhost:7002/api/ws?id=2&session_id='+_this.sessionId + '&h=' + term.rows + '&w=' + term.cols);

                    _this.socket.onopen = function() {
                        term.writeln('Connected to server');  
                    };

                    _this.socket.onerror = function(err) {
                        term.writeln('Connected to err');
                        console.log("连接错误", err)
                    };

                    _this.socket.onclose = function () {
                        console.log("WebSocket close");
                        term.writeln("##  Connected,Close!  ##");
                    }

                    term.loadAddon(new AttachAddon(_this.socket));


                })
            })

        },

        dialogCloseHandler() {
            this.dialogVisible = false;
            this.dialogLoading = false;
            this.btnLoading = false;
            if(this.socket){
                this.socket.close();
                this.socket = null
            }
        },

    },
    mounted() {
        this.$root.BindEventGlobal(
            "openXtermDialogHandler",
            this.openXtermDialogHandler
        );
    },
};
</script>
