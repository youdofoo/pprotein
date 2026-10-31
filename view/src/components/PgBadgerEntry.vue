<template>
  <section>
    <iframe :srcdoc="$data.report" sandbox="allow-scripts"></iframe>
  </section>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  data() {
    return {
      report: "",
    };
  },
  async created() {
    await this.updateReport(this.$route.params.id);
  },
  async beforeRouteUpdate(route) {
    await this.updateReport(route.params.id);
  },
  methods: {
    async updateReport(id: string | string[]) {
      const resp = await fetch(`/api/pgbadger/${id}`);
      this.report = await resp.text();
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
