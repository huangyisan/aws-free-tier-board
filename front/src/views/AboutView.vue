<template>
  <div class="about">
    <h1>This is cost page</h1>
    <v-chart class="chart" :option="option" />
  </div>
</template>

<script>
import axios from "axios";
// import VChart from "vue-echarts";

import { use } from "echarts/core";
import { CanvasRenderer } from "echarts/renderers";
import { PieChart } from "echarts/charts";
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
]);

export default {
  components: {
    VChart,
  },
  name: "Cost",
  data() {
    return {
      data: [],
    };
  },
  computed: {
    option() {
      return {
        xAxis: {
          type: "category",
          data: this.data.map((d) => d.date),
        },
        yAxis: {
          type: "value",
        },
        series: [
          {
            data: this.data.map((d) => d.usage),
            type: "line",
          },
        ],
      };
    },
  },
  mounted() {
    axios
      .get("http://127.0.0.1:8888/v1/cost")
      .then((res) => {
        console.log(res.data);
        this.data = res.data;
      })
      .catch((err) => {
        console.error(err);
      });
  },
};
</script>

<style scoped>
.chart {
  height: 400px;
}
</style>