<template>
  <div class="container">
    <nav>
      <router-link
        v-slot="{ navigate, isActive }"
        :to="`/group/${$route.params.gid}/index/`"
        custom
      >
        <div :class="{ active: isActive }" @click="navigate">index (score: {{ score }})</div>
      </router-link>
      <router-link
        v-for="entry in $store.getters.availableEntriesByGroup(
          $route.params.gid
        )"
        v-slot="{ navigate, isActive }"
        :key="entry.Snapshot.ID"
        :to="`/group/${$route.params.gid}/${entry.Snapshot.Type}/${entry.Snapshot.ID}/`"
        custom
      >
        <div :class="{ active: isActive }" @click="navigate">
          {{ entry.Snapshot.Type }}: {{ entry.Snapshot.Label }}
        </div>
      </router-link>
    </nav>
    <router-view />
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  computed: {
    score() {
      const groupId = typeof(this.$route.params.gid) === "string" ? this.$route.params.gid : this.$route.params.gid[0];
      return this.$store.state.scores[groupId]?.score;
    }
  },
});
</script>

<style scoped lang="scss">
.container {
  display: flex;
  flex-direction: column;
  width: 100%;
  height: 100%;
}

nav {
  background-color: #555;
}
</style>
