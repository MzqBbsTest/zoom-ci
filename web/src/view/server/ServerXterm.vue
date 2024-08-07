<template>
    <el-dialog width="80%" :title="dialogTitle" :visible.sync="dialogVisible"
        @close="dialogCloseHandler">
        <div class="app-dialog" v-loading="dialogLoading">
            <div id="xtermDialog" ref="terminal" style="height:600px;width: 1200px;"></div>
        </div>
        <el-button  @click="windowResize">111</el-button>
    </el-dialog>
</template>

<script>
import { serverSession, serverSessionResize } from "@/api/server";

import { Terminal } from "xterm";
import { AttachAddon } from "xterm-addon-attach";
import { FitAddon } from "xterm-addon-fit";
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
                        try{
                            _this.fitAddon.dispose();
                            _this.term.dispose();
                            _this.socket.close();
                        }catch(e){
                            console.log(e)
                        }
                    }
                    const term = new Terminal({
                        cursorBlink: true,
                        theme: {
                            background:  "#000000",
                            foreground: "#FFFFFF",
                            cursor:  "#FFFFFF",
                        },
                        fontSize: 12,
                        fontFamily: "Courier",
                        cursorStyle: "block",
                    });
                    
                    // term.prompt = () => {
                    //     // term.write("\r\n\x1b[33m$\x1b[0m ");
                    // };
                    // term.prompt();
                    term.open(_this.$refs.terminal);
                    const fitAddon = new FitAddon();
                    term.loadAddon(fitAddon);
                    _this.term = term;
                    _this.fitAddon = fitAddon;
                    _this.reconnect(term)
                    
                    
                })
            })

        },

        windowResize() {
            let _this = this
            Vue.nextTick(() => {
                let connectTabElement = this.$refs.terminal;
                if (connectTabElement === null) {
                    console.log("调整窗口大小,没有获取到dom");
                    return;
                }
                this.fitAddon.fit()
                serverSessionResize({
                    id: 2,
                    session_id: this.sessionId,
                    h: this.term.rows,
                    w: this.term.cols
                })
            });
        },

        reconnect(term){
            
            let _this = this

            // _this.socket = new WebSocket('ws://localhost:8899/api/ssh/conn?h=48&w=95&session_id=' + res + '&Authorization=111');
            this.socket = new WebSocket('ws://localhost:7002/api/ws?id=2&session_id='+this.sessionId + '&h=' + term.rows + '&w=' + term.cols);

            this.socket.onopen = function() {
                term.writeln('Connected to server'); 
                setTimeout(_this.windowResize, 1000)
            };

            this.socket.onerror = function(err) {
                term.writeln('Connected to err');
                console.log("连接错误", err)
            };

            this.socket.onclose = function () {
                console.log("WebSocket close");
                term.writeln("##  Connected,Close!  ##");
            }
            this.fitAddon.fit()
            this.term.loadAddon(new AttachAddon(this.socket));
        },

        dialogCloseHandler() {
            this.dialogVisible = false;
            this.dialogLoading = false;
            this.btnLoading = false;
            try{
                this.fitAddon.dispose();
                this.term.dispose();
                this.socket.close();
            }catch(e){
                console.log(e)
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
