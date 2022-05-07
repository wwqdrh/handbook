<template>
  <div>
    <div v-if="fullscreen === true">
      <transition name="milkmdmask" :duration="300">
        <div
          v-show="showModal"
          class="fixed inset-0 z-100 w-screen h-screen bg-opacity-50 bg-gray-500 flex justify-center"
        >
          <div class="h-screen w-3/4 relative">
            <div
              class="text-left w-full bg-white h-full overflow-x-auto overflow-y-auto"
            >
              <div
                id="mdcontent"
                ref="editor"
                class="p-6 px-12"
                v-html="hcontent"
              ></div>
            </div>
          </div>
        </div>
      </transition>
    </div>
    <div v-else>
      <div
        class="text-left w-full bg-white h-full overflow-x-auto overflow-y-auto"
      >
        <div
          id="mdcontent"
          ref="editor"
          class="p-6 px-12"
          v-html="hcontent"
        ></div>
      </div>
    </div>
  </div>
</template>

<script>
import MarkdownIt from "markdown-it";
import hljs from "highlight.js";
// import { base64StringToBlob } from "blob-util";

// import "./assets/mint/mint.css";
// import "highlight.js/styles/monokai-sublime.css";

const md = new MarkdownIt({
  html: true,
  linkify: true,
  typographer: true,
  highlight: function (str, lang) {
    // 此处判断是否有添加代码语言
    // console.log(str);
    if (lang && hljs.getLanguage(lang)) {
      try {
        // 得到经过highlight.js之后的html代码
        const preCode = hljs.highlight(lang, str, true).value;
        // 以换行进行分割
        const lines = preCode.split(/\n/).slice(0, -1);
        // 添加自定义行号
        let html = lines
          .map((item, index) => {
            return (
              '<li><span class="line-num" data-line="' +
              (index + 1) +
              '"></span>' +
              item +
              "</li>"
            );
          })
          .join("");
        html = "<ol>" + html + "</ol>";
        // 添加代码语言
        if (lines.length) {
          html += '<b class="name">' + lang + "</b>";
        }
        return (
          '<pre class="hljs overflow-x-auto"><code>' + html + "</code></pre>"
        );
      } catch (err) {
        console.log(err);
      }
    }
    // 未添加代码语言，此处与上面同理
    const preCode = md.utils.escapeHtml(str);
    const lines = preCode.split(/\n/).slice(0, -1);
    let html = lines
      .map((item, index) => {
        return (
          '<li><span class="line-num" data-line="' +
          (index + 1) +
          '"></span>' +
          item +
          "</li>"
        );
      })
      .join("");
    html = "<ol>" + html + "</ol>";
    return '<pre class="hljs"><code>' + html + "</code></pre>";
  },
});

export default {
  name: "HelloWorld",
  props: {
    mdcontent: {
      type: String,
      default: "",
    },
    fullscreen: {
      type: Boolean,
      default: true,
    },
  },
  data() {
    return {
      hcontent: "",
      showModal: false,
      content: "<h1>Some initial content</h1>",
      editorSettings: {
        modules: {
          imageDropAndPaste: {
            handler: this.imageHandler,
          },
          syntax: {
            highlight: (text) => hljs.highlightAuto(text).value,
          },
        },
      },
    };
  },
  watch: {
    mdcontent: {
      handler() {
        // console.log(this.fullscreen, this.mdcontent);
        this.loadMarkdownContent();
      },
      immediate: false,
    },
  },
  mounted() {
    this.show();
  },
  methods: {
    show() {
      if (this.fullscreen) {
        this.loadMarkdownContent();
        document.addEventListener("click", this._action, false);

        this.showModal = true;
      }
    },

    hide() {
      if (this.fullscreen) {
        document.removeEventListener("click", this._action, false);

        this.showModal = false;
      }
    },

    _action(e) {
      if (this.$refs.editor.contains(e.target)) {
        return false;
      }
      this.hide();
      return false;
    },

    loadMarkdownContent() {
      const attributes = document.getElementById("mdcontent").attributes;
      let mountedvalue = "";
      for (let i = 0; i < attributes.length; i++) {
        const v = attributes[i].localName; // value
        if (v.startsWith("data-v-")) {
          mountedvalue = v;
        }
      }

      let res = md.render(`${this.mdcontent}`);
      if (mountedvalue !== "") {
        // render(res)
        const newRes = [];

        let state = 0; // 状态
        for (const i in res) {
          const ch = res[i];
          if (state === 0 && ch === "<") {
            newRes.push(ch);
            state = 1;
          } else if (state === 1 && res.substr(i, 3)  === "img") {
            // img标签 跳过
            newRes.push(ch);
            state = -1;
          } else if (state === 1 && ch === "!") {
            newRes.push(ch);
            state = -1;
          } else if (state === -1 && ch === ">") {
            newRes.push(ch);
            state = 0;
          } else if (state === 1 && ch === ">") {
            newRes.push(mountedvalue + " >");
            state = 0;
          } else if (state === 1 && ch !== " ") {
            newRes.push(ch);
            state = 2;
          } else if (state === 1 && ch === "/") {
            // </a>
            newRes.push(ch);
            state = 3;
          } else if (state === 2 && ch === "/") {
            // <a />
            newRes.push(" " + mountedvalue + "/");
            state = 3;
          } else if (state === 2 && ch === ">") {
            newRes.push(" " + mountedvalue + ch);
            state = 0;
          } else if (state === 3 && ch === ">") {
            newRes.push(ch);
            state = 0;
          } else {
            newRes.push(ch);
          }
        }
        res = newRes.join("");
      }
      // 为每个dom添加mountedvalue
      this.hcontent = res;
    },

    imageHandler: function (imageDataUrl, type) {
      // base64 to blob
      // const blob = base64StringToBlob(
      //   imageDataUrl.replace(/^data:image\/\w+;base64,/, ""),
      //   type
      // );
      // const filename = [
      //   "my",
      //   "cool",
      //   "image",
      //   "-",
      //   Math.floor(Math.random() * 1e12),
      //   "-",
      //   new Date().getTime(),
      //   ".",
      //   type.match(/^image\/(\w+)$/i)[1],
      // ].join("");
    },
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
@import "./assets/mint/mint.css";
@import "./assets/monokai-sublime.css";
</style>

<style>
:root {
  --side-bar-bg-color: #d9ede5;
  --control-text-color: #6b6b6b;
  --active-file-bg-color: #ecf6f2;
  --active-file-border-color: #6b6b6b;
  --active-file-text-color: #202020;
  --table-even-row-color: #f8fcfa;
  --table-head-color: #d9ede5;
  --deep-theme-color: #4eb289;
  --code-block-bg-color: #0f111a;
}
</style>
