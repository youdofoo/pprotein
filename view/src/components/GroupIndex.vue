<template>
  <section>
    <GroupEntriesTable
      :group-id="groupId"
      :entries="$store.getters.entriesByGroup(groupId)"
    />
    <Score :group-id="groupId" :initialValue="score.toString()"/>
    <AddMemo :group-id="groupId" />
  </section>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import GroupEntriesTable from "./GroupEntriesTable.vue";
import AddMemo from "./AddMemo.vue";
import Score from "./Score.vue";

export default defineComponent({
  components: {
    AddMemo,
    GroupEntriesTable,
    Score,
  },
  computed: {
    groupId() {
      const gid = this.$route.params.gid;
      return typeof gid == "string" ? gid : gid[0];
    },
    score() {
      return this.$store.state.scores[this.groupId]?.score;
    }
  },
});
</script>

<style scoped lang="scss">
section {
  margin: 2em;
}
</style>
