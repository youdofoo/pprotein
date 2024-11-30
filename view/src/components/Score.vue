<template>
    <div class="score-container">
      <div>
        <label>
          Score:
          <input v-model="score" type="text"/>
        </label>
      </div>
      <button @click="setScore">Set Score</button>
    </div>
  </template>
  
  <script lang="ts">
  import { defineComponent } from "vue";
  
  export default defineComponent({
    props: {
      groupId: {
        type: String,
        required: true,
      },
      initialValue: {
        type: String,
        required: true,
      }
    },
    data() {
      return {
        score: this.initialValue,
      };
    },
    watch: {
      initialValue(newVal) {
        this.score = newVal;
      },
    },
    methods: {
      async setScore() {
        const v = parseInt(this.score, 10);
        if (Number.isNaN(v)) {
            alert("整数値を入力してください");
            return;
        }
        const resp = await fetch(`/api/score/${this.groupId}`, {
          method: "PUT",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({
            score: v,
          }),
        });
  
        if (!resp.ok) {
          alert(await resp.text());
        }
        this.$store.commit("saveScore", {groupId: this.groupId, score: this.score});
      },
    },
  });
  </script>
  
  <style scoped lang="scss">
  .score-container {
    margin-top: 1em;
    display: flex;
    align-items: end;
    gap: 1em;
  }
  </style>
  