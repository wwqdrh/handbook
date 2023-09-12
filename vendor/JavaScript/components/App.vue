<template>
  <div class="flex w-screen h-screen">
    信息聚合面板
    <div class="flex flex-col w-1/4 h-full">
      <div
        v-for="(item, i) in k8sList"
        :key="'item-' + i"
        class="cursor-pointer"
        @click="getContent(item.href)"
      >
        {{ item.title }}
      </div>
    </div>
    <div class="w-3/4 h-full overflow-y-auto">
      <MarkdownRender :fullscreen="false" :mdcontent="content" />
    </div>
  </div>
</template>

<script>
import { GetK8sArticleList, GetK8sArticleContent } from "./api/k8sorg";

import MarkdownRender from "./components/milkmd/MaskRender.vue";

export default {
  name: "App",
  components: {
    MarkdownRender,
  },
  data: () => ({
    k8sList: [],
    content: "",
  }),
  created() {
    const this_ = this;

    this.sleep(1000).then(() => {
      GetK8sArticleList()
        .then((res) => {
          this_.k8sList = res.data;
        })
        .catch((err) => {
          alert(err);
        });
    });
  },
  methods: {
    getContent(url) {
      const this_ = this;
      GetK8sArticleContent(url)
        .then((res) => {
          this_.content = res.data.content;
        })
        .catch((err) => {
          alert(err);
        });
    },
    sleep(ms) {
      return new Promise((resolve) => setTimeout(resolve, ms));
    },
  },
};
</script>
