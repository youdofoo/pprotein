<template>
  <section>
    <a :href="link" download>Download</a>
    <table border="1">
      <thead>
        <tr>
          <th v-for="(h, i) in header" :key="h" @click="sortBy(i)">
            {{ h }}
            <span v-if="sortColumn === i" class="sortOrder">
              {{ sortOrder == "desc" ? "▼" : "▲" }}
            </span>
          </th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="d in sortedData" :key="d.toString()">
          <td v-for="(value, i) in d" :key="header[i]">
            <div v-if="isNumeric(value)" class="numericCell">
              {{ value }}
            </div>
            <div v-else>{{ value }}</div>
          </td>
        </tr>
      </tbody>
    </table>
  </section>
</template>

<script lang="ts">
import { parse } from "papaparse";
import { defineComponent } from "vue";

export default defineComponent({
  props: {
    tsv: {
      type: String,
      required: true,
    },
    link: {
      type: String
    }
  },
  data() {
    return {
      sortColumn: -1,
      sortOrder: "desc",
    };
  },
  computed: {
    rows() {
      return parse<string[]>(this.tsv, { skipEmptyLines: true }).data;
    },
    header() {
      return this.rows[0] || [];
    },
    sortedData() {
      if (this.sortColumn === -1) {
        return this.rows.slice(1);
      }
      const data = this.rows.slice(1).sort((as, bs) => {
        const [a, b] = [as[this.sortColumn], bs[this.sortColumn]];
        if (this.isNumeric(a)) {
          return +a - +b;
        } else {
          if (a < b) {
            return -1;
          } else if (a > b) {
            return 1;
          } else {
            return 0;
          }
        }
      });

      if (this.sortOrder === "desc") {
        return data.reverse();
      } else {
        return data;
      }
    },
  },
  methods: {
    isNumeric(v: string) {
      return !isNaN(+v);
    },
    sortBy(column: number) {
      if (this.sortColumn === column) {
        this.sortOrder = this.sortOrder === "desc" ? "asc" : "desc";
      } else {
        this.sortColumn = column;
        this.sortOrder = "desc";
      }
    },
  },
});
</script>

<style scoped lang="scss">
section {
  padding: 2em;
  overflow: scroll;
}

table {
  border-collapse: collapse;
}

table tr:nth-child(odd) td {
  background-color: #f0f0f0;
}

thead {
  background-color: rgb(150, 196, 241);
}

td,
th {
  padding: 0.5em 1em;
  border: 1px solid #999;
}
th {
  cursor: pointer;
  white-space: nowrap;

  .sortOrder {
    font-size: 0.4em;
    color: gray;
  }
}
.numericCell {
  text-align: right;
  white-space: nowrap;
}

details[open] > summary {
  display: none;
}
</style>
