<template>
  <section>
    <iframe v-if="$data.src" :src="$data.src" sandbox="allow-scripts"></iframe>
  </section>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  data() {
    return {
      src: "",
    };
  },
  async created() {
    await this.updateReport(this.$route.params.id);
  },
  async beforeRouteUpdate(route) {
    await this.updateReport(route.params.id);
  },
  beforeUnmount() {
    this.revoke();
  },
  methods: {
    // srcdoc だと、レポート内の #hash のリンクが親ページの URL に解決されて真っ白になるので、blob URL で表示する
    async updateReport(id: string | string[]) {
      const resp = await fetch(`/api/pgbadger/${id}`);
      const blob = new Blob([await resp.text()], { type: "text/html" });
      this.revoke();
      this.src = URL.createObjectURL(blob);
    },
    revoke() {
      if (this.src) {
        URL.revokeObjectURL(this.src);
        this.src = "";
      }
    },
  },
});
</script>

<style scoped lang="scss">
section {
  height: 100%;
}
iframe {
  width: 100%;
  height: 100%;
  border: none;
}
</style>
