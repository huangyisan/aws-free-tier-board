<template>
  <GChart :type="type" :data="data" :options="options" />
</template>

<script>
import { GChart } from "vue-google-charts/legacy";

import { chartType, chartData, chartOptions } from "./GoogleChart.js";

export default {
  name: "GoogleChart",
  components: {
    GChart,
  },
  data() {
    return {
      type: chartType,
      data: chartData,
      options: chartOptions,
    };
  },
  mounted() {
    onMount: {
      console.log(this.data.push(["2022", 2000, 550]));
    }
  },
};
</script>
