<template>
  <section>
    <pre>{{ $data.pt }}</pre>
  </section>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  data() {
    return {
      pt: "",
    };
  },
  async created() {
    await this.updatePt(this.$route.params.id);
  },
  async beforeRouteUpdate(route) {
    await this.updatePt(route.params.id);
  },
  methods: {
    async updatePt(id: string | string[]) {
      const resp = await fetch(`/api/pt/${id}`);
      this.pt = await resp.text();
    },
  },
});
</script>

<style scoped lang="scss">
:deep(th) {
  font-size: 75%;
}
</style>