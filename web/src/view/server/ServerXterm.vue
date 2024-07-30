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
                        term.write("\r\n\x1b[33m$\x1b[0m ");
                    };
                    term.prompt();
                    term.open(_this.$refs.terminal);
                    const fitAddon = new FitAddon();
                    term.loadAddon(fitAddon);
                    fitAddon.fit()
                    _this.term = term;

                    // _this.socket = new WebSocket('ws://localhost:8899/api/ssh/conn?h=48&w=95&session_id=' + res + '&Authorization=111');
                    _this.socket = new WebSocket('ws://localhost:7002/api/ws?id=2&session_id='+_this.sessionId);

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

        runFakeTerminal() {
            let _this = this;
            let term = _this.term;
            if (term._initialized) return;
            term._initialized = true;
            term.writeln("Welcome to \x1b[1;32m麦考林\x1b[0m.");
            term.writeln( "This is Web Terminal of Modb; Good Good Study, Day Day Up." );
            term.prompt();

            // 监听键盘事件
            term.attachCustomKeyEventHandler((e) => {
                if (e.key === 'ArrowUp' || e.key === 'ArrowDown') {
                    if(e.type == "keydown"){
                        _this.socket.send(JSON.stringify({
                            id: 2,
                            action: 'Cmd',
                            uid: 0,
                            session: _this.sessionId,
                            data:e.key === 'ArrowUp' ? '\u001b[A' :  '\u001b[B' 
                        }));
                    }
                    e.preventDefault();
                    return false; // 返回 false 表示不处理这个事件
                }else if(e.key === 'ArrowLeft'){

                }else if(e.key === 'ArrowRight'){

                }
                return true; // 其他按键正常处理
            });

            // 添加事件监听器，支持输入方法
            term.onKey((e) => {
                const printable =
                    !e.domEvent.altKey &&
                    !e.domEvent.altGraphKey &&
                    !e.domEvent.ctrlKey &&
                    !e.domEvent.metaKey;
                if (e.domEvent.keyCode === 13) {
                    term.prompt();
                    if (_this.command.length == 0) term.writeln("");
                    console.log("回车，发送命令", _this.command, _this.command.length)
                    _this.command = "";
                }  else if (e.domEvent.keyCode === 8) {
                    // back 删除的情况
                    if (term._core.buffer.x > 2) {
                        term.write("\b \b");
                        _this.command = _this.command.slice(0, -1)
                    }
                } else if (printable) {
                    term.write(e.key);
                    _this.command += e.key;
                }
            });

            term.onData((key) => {
                // 粘贴的情况
                if (key.length > 1) {
                    term.write(key);
                    _this.command += key;
                }
            });

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
