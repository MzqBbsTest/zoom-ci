// clipboard.js

import Clipboard from 'clipboard';
import { Message } from 'element-ui';

export default {
  bind(el, binding) {
    el._clipboard = new Clipboard(el, {
      text: () => binding.value, // 绑定要复制的文本
    });

    el._clipboard.on('success', () => {
      if (binding.arg === 'success') {
        binding.value(); // 如果传递了 success 参数，执行成功回调
      } else {
        // 可以通过 Vue 的提示框或其他方式反馈
        Message({
            type: 'success',
            message: '复制成功！'
        });
        // alert('复制成功！'); // 这里使用 alert 提示，实际应用中可以用 Element UI 的 $message
      }
    });

    el._clipboard.on('error', () => {
      if (binding.arg === 'error') {
        binding.value(); // 如果传递了 error 参数，执行错误回调
      } else {
        Message({
            type: 'error',
            message: '复制失败，请手动复制！'
        });
        // alert('复制失败，请手动复制！'); // 这里使用 alert 提示，实际应用中可以用 Element UI 的 $message
      }
    });
  },
  update(el, binding) {
    // 当绑定值变化时，更新 Clipboard 实例的内容
    if (binding.value !== binding.oldValue) {
      el._clipboard.text = () => binding.value; // 更新要复制的文本内容
    }
  },
  unbind(el) {
    if (el._clipboard) {
      el._clipboard.destroy();
      delete el._clipboard;
    }
  }
};
