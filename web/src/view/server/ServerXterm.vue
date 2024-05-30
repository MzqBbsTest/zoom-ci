<template>
    <el-dialog :width="$root.DialogLargeWidth" :title="dialogTitle" :visible.sync="dialogVisible"
        @close="dialogCloseHandler">
        <div class="app-dialog" v-loading="dialogLoading">
            <div id="xtermDialog" ref="terminal"></div>
        </div>
    </el-dialog>
</template>

<script>
import { } from "@/api/server";

import { Terminal } from "xterm";
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
            command: ""
        };
    },
    methods: {
        openXtermDialogHandler(params) {
            this.dialogVisible = true;
            Vue.nextTick(
                function () {
                    let _this = this
                    if (this.term || !this.$refs.terminal) {
                        return;
                    }
                    const term = new Terminal();
                    // console.log(term.prompt)
                    term.prompt = () => {
                        term.write("\r\n\x1b[33m$\x1b[0m ");
                    };
                    term.prompt();
                    term.open(this.$refs.terminal);
                    const fitAddon = new FitAddon();
                    term.loadAddon(fitAddon);
                    fitAddon.fit()

                    this.term = term;
                    this.runFakeTerminal();
                    
                    this.socket = new WebSocket('ws://localhost:7002/api/ws?id=');
                    this.socket.onopen = function() {
                        term.write('Connected to server\\r\\n');
                    };

                    this.socket.onerror = function(err) {
                        console.log("连接错误", err)
                    };

                    this.socket.onmessage = function(event) {
                        console.log("onmessage", event.data)
                        term.write(event.data);
                    };

                    term.onData(function(data) {
                        console.log("onData", data)
                        _this.socket.send(data);
                    });

                    
                    // 内容全屏显示-窗口大小发生改变时
                    function resizeScreen(size) {
                        console.log("size", size);
                        try {
                            fitAddon.fit();

                            // 窗口大小改变时触发xterm的resize方法，向后端发送行列数，格式由后端决定
                            term.onResize(size => {
                                // _this.onSend({ Op: "resize", Cols: size.cols, Rows: size.rows });
                            });
                        } catch (e) {
                            console.log("e", e.message);
                        }
                    }
                    window.addEventListener("resize", resizeScreen);
                }.bind(this)
            );
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
            term.writeln(
                "This is Web Terminal of Modb; Good Good Study, Day Day Up."
            );
            term.prompt();
            // 添加事件监听器，支持输入方法
            term.onKey((e) => {
                const printable =
                    !e.domEvent.altKey &&
                    !e.domEvent.altGraphKey &&
                    !e.domEvent.ctrlKey &&
                    !e.domEvent.metaKey;
                if (e.domEvent.keyCode === 13) {
                    term.prompt();
                    console.log("回车，发送命令", _this.command)
                    _this.command = "";
                } else if (e.domEvent.keyCode === 8) {
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
