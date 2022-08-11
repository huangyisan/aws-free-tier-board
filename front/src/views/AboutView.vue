<template>
  <div class="about">
    <!-- <h1>This is cost page</h1> -->
    <div>
      <v-row justify="center">
        <v-col cols="12" sm="12">
          <v-date-picker full-width v-model="dates" range></v-date-picker>
        </v-col>
        <v-col cols="12" sm="12">
          <v-btn block color="primary" @click="submitDate">
            Block Button
          </v-btn>
        </v-col>
      </v-row>
      <v-chart class="chart" :option="option" />
    </div>
  </div>
</template>

<script>
// import dateFormat from "./tools.js";
import axios from "axios";
import dateformat from "dateformat";
// import VChart from "vue-echarts";

import { use } from "echarts/core";
import { CanvasRenderer } from "echarts/renderers";
import { PieChart } from "echarts/charts";
import { BarChart } from "echarts/charts";
import "echarts/lib/component/grid";
import "echarts/lib/chart/line";
import {
  TitleComponent,
  TooltipComponent,
  LegendComponent,
} from "echarts/components";
import VChart, { THEME_KEY } from "vue-echarts";

use([
  CanvasRenderer,
  PieChart,
  TitleComponent,
  TooltipComponent,
  LegendComponent,
  BarChart,
]);

export default {
  components: {
    VChart,
  },
  name: "Cost",
  methods: {
    genXAxisData(d) {
      return d.tag;
    },
    genYAxisData(d) {
      return d.value;
    },
    getAllTraffic(dates) {
      dates.sort();
      axios({
        method: "GET",
        url: "http://127.0.0.1:8888/v1/traffic",
        params: {
          start: dates[0],
          end: dates[1],
        },
      })
        .then((res) => {
          if (res.data !== null) {
            this.data = res.data;
          } else {
            this.data = [
              {
                tag: "null",
                value: 0,
              },
            ];
          }
        })
        .catch((err) => {
          console.error(err);
        });
    },

    setDatesRange(dates) {
      let today = new Date();
      today = dateformat(today, "yyyy-mm-dd");
      let tomorrow = new Date().setDate(new Date().getDate() + 1);
      tomorrow = dateformat(tomorrow, "yyyy-mm-dd");
      this.dates.push(today, tomorrow);
    },

    submitDate() {
      this.getAllTraffic(this.dates);
    },
  },
  data() {
    return {
      dates: [],
      data: [],
    };
  },
  computed: {
    dateRangeText() {
      return this.dates.join(" ~ ");
    },
    option() {
      return {
        xAxis: {
          type: "category",
          data: this.data.map((d) => d.tag),
        },
        yAxis: {
          type: "value",
        },
        series: [
          {
            data: this.data.map((d) => d.value),
            type: "bar",
          },
        ],
      };
    },
  },
  mounted() {
    this.setDatesRange(this.dates);
    this.getAllTraffic(this.dates);
  },
};
</script>

<style scoped>
.chart {
  height: 400px;
}
</style>